// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package finalizer_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/finalizer"
)

func TestUnsetFinalizers(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	obj := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-obj",
			Finalizers: []string{"finalizer1", "finalizer2"},
		},
	}
	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(obj).Build()
	ctx := context.TODO()

	err := finalizer.UnsetFinalizers(ctx, fakeClient, obj, "finalizer1")

	require.NoError(t, err)
	patched := corev1.Pod{}
	require.NoError(t, fakeClient.Get(ctx, client.ObjectKeyFromObject(obj), &patched))
	assert.NotContains(t, patched.GetFinalizers(), "finalizer1")
}

func TestEnsureFinalizers(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	obj := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-obj",
			Finalizers: []string{},
		},
	}
	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(obj).Build()
	ctx := context.TODO()

	err := finalizer.EnsureFinalizers(ctx, fakeClient, obj, "finalizer1")

	require.NoError(t, err)
	patched := corev1.Pod{}
	require.NoError(t, fakeClient.Get(ctx, client.ObjectKeyFromObject(obj), &patched))
	require.Contains(t, patched.GetFinalizers(), "finalizer1")
}
