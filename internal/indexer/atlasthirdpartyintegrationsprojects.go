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

//nolint:dupl
package indexer

import (
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/client"

	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
)

const (
	AtlasThirdPartyIntegrationByProjectIndex = "atlasThirdPartyIntegration.spec.projectRef"
)

type AtlasThirdPartyIntegrationByProjectIndexer struct {
	AtlasReferrerByProjectIndexerBase
}

func NewAtlasThirdPartyIntegrationByProjectIndexer(logger *zap.Logger) *AtlasThirdPartyIntegrationByProjectIndexer {
	return &AtlasThirdPartyIntegrationByProjectIndexer{
		AtlasReferrerByProjectIndexerBase: *NewAtlasReferrerByProjectIndexer(
			logger,
			AtlasThirdPartyIntegrationByProjectIndex,
		),
	}
}

func (*AtlasThirdPartyIntegrationByProjectIndexer) Object() client.Object {
	return &akov2.AtlasThirdPartyIntegration{}
}
