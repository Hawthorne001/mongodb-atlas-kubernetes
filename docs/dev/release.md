# Atlas Operator Release Instructions

For the various PRs involved, seek at least someone else to approve. In case of doubts, engage the team member(s) who might be able to clarify and seek their review as well.

## Prerequisites

To get PRs to be auto-committed for RedHat community & Openshift you need to make sure you are listed in the [team members ci.yaml list for community-operators](https://github.com/k8s-operatorhub/community-operators/blob/main/operators/mongodb-atlas-kubernetes/ci.yaml) and [team members ci.yaml list for community-operators-prod](https://github.com/redhat-openshift-ecosystem/community-operators-prod/blob/main/operators/mongodb-atlas-kubernetes/ci.yaml).

This is not required for [Certified Operators](https://github.com/redhat-openshift-ecosystem/certified-operators/blob/main/operators/mongodb-atlas-kubernetes/ci.yaml).

Finally, make sure you have a "RedHat Connect" account and are a [team member with org administrator role in the team list](https://connect.redhat.com/account/team-members).

### Tools

Most tools are automatically installed for you. Most of them are Go binaries and use `go install`. There are a few that might cause issues and you might want to pre-install manually:

- [devbox](https://www.jetify.com/devbox) to be able to enter a sandbox development environment that includes necessary tools for the release process.
- [Docker](https://www.docker.com/) to be able to deal with containers.

## Release preparations (minimum n-1 weeks before the actual release)

At least **one** (1) week before the release the Kubernetes Version testing matrix has to be updated both in this repository and the CLI repository https://github.com/mongodb/mongodb-atlas-cli.

Please refer to the [CI documentation](ci.md#kubernetes-version-matrix) and submit a pull request, example: https://github.com/mongodb/mongodb-atlas-kubernetes/pull/2161 or https://github.com/mongodb/mongodb-atlas-kubernetes/pull/2082.

## Release Notes

- Create a draft of the release notes in a Google Document and share with Product and the Docs team.
  - In confluence, look for the `AKO Release Internal details` page for more details.
- Ensure as well that supporting documents for new features are in review.
- Wait for approval of the release notes and availability of the associated documents.

**DO NOT** start the release process until the release notes are approved and associated documentation is, at least, in review state. Always seek explicit approval by Product and/or Management.

The reason for this preparatory step is to avoid customers getting new or breaking changes before their supporting documentation.

## Create the Release

Once release notes and documentation are approved, trigger the [`release-image.yml`](../../.github/workflows/release-image.yml) workflow.

You will be prompted to enter:

| Input       | Description                                                                                         | Required | Default  | Example                               |
|-------------|-----------------------------------------------------------------------------------------------------|----------|----------|---------------------------------------|
| `version`   | The version to be released, including the `v` prefix                                                | Yes      | None     | `v1.10.3`                             |
| `authors`   | A comma-separated list of MongoDB email addresses responsible for the release                       | Yes      | None     | `alice@mongodb.com,bob@mongodb.com`   |
| `image_sha` | The 7-character Git commit SHA used for the promoted image, or `'latest'` for the most recent       | No       | `latest` | `3e79a3f`.                            |

The inputs `version` and `authors` must be filled out every time you trigger the release workflow. The `image_sha` is optional and defaults to `latest` if left empty.

The `image_sha` corresponds exactly to the 7-character Git commit SHA used to build the operator image; for example, `image_sha: 3e79a3f` means the image was built from Git commit `3e79a3f`. Using `latest` as the `image_sha` means the workflow will release the most recently promoted and tested operator image—not necessarily the latest Git commit—and when `latest` is used, the workflow will echo the corresponding Git commit during the internal steps so the user knows exactly which source is being released.

### Example Release Input

```yaml
version: v1.10.3
authors: alice@mongodb.com,bob@mongodb.com
image_sha: 3e79a3f
```

or

```yaml
version: v1.10.3
authors: alice@mongodb.com,bob@mongodb.com
image_sha: latest
```

### What Happens Next

Once triggered:

- A release PR is created that adds a new `release/<version>` directory (containing `deploy/`, `helm-charts/`, and `bundle/` directories)
- A Git tag of the form `v<version>` is created and pushed on GitHub
- A GitHub release is published with:
  - Zipped `all-in-one.yml`
  - SDLC-compliant artifacts: SBOMs and compliance reports

The only manual step is to **review and merge** the release PR. This PR does **not** re-run any of the expensive tests on cloud-qa.

**Note:** this directory-based approach avoids merge conflicts entirely. Because each release introduces a clean, isolated `release/<version>` folder, it can be merged directly into `main` without conflicting with prior or future releases. This enables a linear and conflict-free release history while maintaining clear traceability for each version.

---

## Image Promotion

The `image_sha` used in a release must already be tested and promoted via CI. Promotion can occur in one of three ways:

- A scheduled CI run on the `main` branch
- A merge into `main` that includes production code changes
- A manual dispatch of the `tests.yml` workflow with the `promote` flag enabled

### How Promotion Works

During promotion, the operator image used in Helm-based E2E tests is first built and published as a dummy image in `ghcr.io`. Once **all** tests—including the cloud-based Helm scenarios—complete successfully, the [`promote-image.yml`](../../.github/workflows/promote-image.yml) workflow is triggered.

This workflow:

- Verifies that all required tests succeeded
- Moves the image from `ghcr.io` to official prerelease registries in `docker.io` and `quay.io`
- Tags the image in the official prerelease registires as:
  - `promoted-<git_sha>` — uniquely maps the image to the source Git commit
  - `promoted-latest` — always points to the most recent image that passed all tests

The `promoted-<git_sha>` builds the one-to-one correspondence between the 7-character Git commit and the `image_sha`. For the correspondence between the 7-character Git commit and `image_sha: latest`, we internally store a label within the image `promoted-latest` that has the exact git commit used for that image. Moreover, the `promoted-latest` tag is only updated by events that run on the main branch—whether triggered by a schedule, a merge, or a workflow dispatch. Manual promotions on any other branch will never overwrite this tag.

One can find promoted images by checking the [`promote-image.yml`](../../.github/workflows/promote-image.yml) workflow runs in GitHub Actions, or by browsing the prerelease Docker registries at:

- Docker Hub: `mongodb/mongodb-atlas-kubernetes-prerelease`
- Quay.io: `mongodb/mongodb-atlas-kubernetes-prerelease`

**Note:** When releasing, you omit the `promoted-` prefix and specify only the image SHA or `latest`. The `promoted-` prefix is used internally to organize images in the registries.

### Best Practice

Releases should generally use `latest` as the `image_sha`. This ensures that you are releasing the most recently tested and CI-verified image.

## Manual SSDLC steps

### Process Overview

The SSDLC process requirements are as follows:

1. Sign our images with a MongoDB owned signature.
1. Produce SBOM (Software Bill Of Materials) for each platform we support (`linux-amd64` and `linux-arm64`).
1. Upload the plain SBOMs to a MongoDB internal Kondukto service instance.
1. Produce the augmented SBOMS, including vulnerability metadata, from using Silkbomb 2.0.
1. Store both sets of SBOM files for internal reference.

The first two steps are semi-automated as documented here. The rest is fully manual.

Right now we are only using **one Kondukto branch per platform**:
- `main-linux-amd64`
- `main-linux-arm64`

This means only the latest version is tracked by Kondukto. Note each upload will replace the SBOM document tracked on each asset group.

For more details about credentials required, to to `MongoDB Confluence` and look for page:
`Kubernetes Atlas Operator SSDLC Compliance Manual`

What follows is a quick reference of the make rules involved, assuming the credential setup is already completed and the process is already familiar.

### Upload SBOMs to Kondukto and Augment SBOMs with Kondukto Scan results

Make sure that you have the credentials configured to handle SBOM artifacts.
Read through the wiki page "Kubernetes Atlas Operator SSDLC Compliance Manual" on how to retrieve them.

Update the local `main` branch to point to the commit which includes the merged SSDLC files from the previous step:

```shell
$ git checkout main
$ git pull
```

```shell
$ make augment-sbom SBOM_JSON_FILE="docs/releases/v${VERSION}/linux_amd64.sbom.json"
$ make augment-sbom SBOM_JSON_FILE="docs/releases/v${VERSION}/linux_arm64.sbom.json"
```

### Register SBOMs internally

To be able to store SBOMs in S3, you need special credentials.
Please advise the Wiki page "Kubernetes Atlas Operator SSDLC Compliance Manual".

```shell
$ make store-augmented-sboms VERSION=${VERSION} TARGET_ARCH=amd64
$ make store-augmented-sboms VERSION=${VERSION} TARGET_ARCH=arm64
```

## Edit the Release Notes and publish the release

Follow the format described in the [release-notes-template.md](../release-notes/release-notes-template.md) file.
Paste the release notes content approved before the release was started.
Once the image is out, publish the release notes draft as soon as possible.

## Synchronize configuration changes with the Helm Charts

Go to the [helm-chart repo](https://github.com/mongodb/helm-charts) and locate the [Pull Request](https://github.com/mongodb/helm-charts/pulls)
that is being automatically generated by the [GitHub "Create PR with Atlas Operator Release" action](https://github.com/mongodb/helm-charts/actions/workflows/post-atlas-operator-release.yaml).
It is named "Release Atlas Operator x.y.z.".

The will update two Helm charts:
* [atlas-operator-crds](https://github.com/mongodb/helm-charts/tree/main/charts/atlas-operator-crds)
* [atlas-operator](https://github.com/mongodb/helm-charts/tree/main/charts/atlas-operator)
    
Merge the PR - the chart will get released automatically.

## Create Pull Requests to publish OLM bundles

All bundles/package manifests for Operators for operatorhub.io reside in the following repositories:
* https://github.com/k8s-operatorhub/community-operators - Kubernetes Operators that appear on [OperatorHub.io](https://operatorhub.io/)
* https://github.com/redhat-openshift-ecosystem/community-operators-prod - Kubernetes Operators that appear on [OpenShift](https://openshift.com/) and [OKD](https://www.okd.io/)
* https://github.com/redhat-openshift-ecosystem/certified-operators - Red Hat certified Kubernetes Operators

### Fork/Update the community operators repositories

**Note**: this has to be done once only. 

First ensure your SSH keys in [https://github.com/settings/keys] are authorized for `mongodb-forks` MongoDB SSO.

Execute the following steps:

1. Clone each of the above forked OLM repositories from https://github.com/mongodb-forks
2. Add `upstream` remotes
3. Export each cloned repository directory in environment variables

#### community-operators repository
```
git clone git@github.com:mongodb-forks/community-operators.git
git remote add upstream https://github.com/k8s-operatorhub/community-operators.git
export RH_COMMUNITY_OPERATORHUB_REPO_PATH=$PWD/community-operators
```
#### community-operators-prod repository
```
git clone git@github.com:mongodb-forks/community-operators-prod.git
git remote add upstream https://github.com/redhat-openshift-ecosystem/community-operators-prod.git
export RH_COMMUNITY_OPENSHIFT_REPO_PATH=$PWD/community-operators-prod
```
#### certified-operators repository
```
git clone git@github.com:mongodb-forks/certified-operators.git
git remote add upstream https://github.com/redhat-openshift-ecosystem/certified-operators
export RH_CERTIFIED_OPENSHIFT_REPO_PATH=$PWD/certified-operators
```

### Create a Pull Request for the `community-operators` repository

1. Ensure the `RH_COMMUNITY_OPERATORHUB_REPO_PATH` environment variable is set.
2. Invoke the following script with `<version>` set to `1.0.0` (don't use a `v` prefix):
```
./scripts/release-redhat.sh <version>
```

You can see an [example fixed PR here on Community Operators for version 1.9.1](https://github.com/k8s-operatorhub/community-operators/pull/3457).

Create the PR to the main repository and wait until CI jobs get green. 
After the PR is approved and merged - it will soon get available on https://operatorhub.io

### Create a Pull Request for the `community-operators-prod` repository

1. Ensure the `RH_COMMUNITY_OPENSHIFT_REPO_PATH` environment variable is set.
2. Invoke the following script with `<version>` set to `1.0.0` (don't use a `v` prefix):
```
./scripts/release-redhat-openshift.sh <version>
```

Submit the PR to the upstream repository and wait until CI jobs get green.

**Note**: It is required that the PR consists of only one commit - you may need to do
`git rebase -i HEAD~2; git push origin +mongodb-atlas-operator-community-<version>` if you need to squash multiple commits into one and perform force push)

After the PR is approved it will soon appear in the [Atlas Operator openshift cluster](https://console-openshift-console.apps.atlas.operator.mongokubernetes.com)

### Create a Pull Request for the `certified-operators` repository

This is necessary for the Operator to appear on "operators" tab in Openshift clusters in the "certified" section.
Ensure the `RH_CERTIFIED_OPENSHIFT_REPO_PATH` environment variable is set.

Invoke the following script and ensure to have the `VERSION` variable set from above:
```
./scripts/release-redhat-certified.sh
```

Then go the GitHub and create a PR
from the `mongodb-fork` repository to https://github.com/redhat-openshift-ecosystem/certified-operators (`origin`).

Note: For some reason, the certified OpenShift metadata does not use the multi arch image reference at all, and only understand direct architecture image references.

You can see an [example fixed PR here for certified version 1.9.1](https://github.com/redhat-openshift-ecosystem/certified-operators/pull/3020).

After the PR is approved it will soon appear in the [Atlas Operator openshift cluster](https://console-openshift-console.apps.atlas.operator.mongokubernetes.com)

# Post install hook release

If changes have been made to the post install hook (mongodb-atlas-kubernetes/cmd/post-install/main.go).
You must also release this image. Run the "Release Post Install Hook" workflow manually specifying the desired 
release version.

# Post Release actions

If the release is a new minor version, then the CLI must be updated with the new version (and any new CRDs) [here](https://github.com/mongodb/atlas-cli-plugin-kubernetes/blob/main/internal/kubernetes/operator/features/crds.go).

If necessary, a CLI plugin release can be created as detailed [here](https://github.com/mongodb/atlas-cli-plugin-kubernetes/blob/main/RELEASING.md).

# Updating the ROSA cluster

For the Openshift upgrade tests we rely on a service account to be present in the OpenShift cluster and its login token to be present in CI.

## Setup Kubectl against the new cluster

1. Go to https://console.redhat.com/openshift
1. Use your RedHat account credentials to log in, see Pre-requisites on the RedHat Connect account you need to setup before this.
1. Form the list of Clusters, click of the name of the one to be used now.
1. CLick the `Open Console` in the top right of the page.
1. Use the cluster `htpasswd` credentials you should have been given beforehand to login to the cluster itself.
1. On the landing page, click the account drop down on the top right corner if the page and click on `Copy login command` there.
1. Login again with the `htpasswd`credentials.
1. On the white page click `Display token`.
1. Copy the `oc` command there and run it. You need to have [oc installed](https://docs.openshift.com/container-platform/4.8/cli_reference/openshift_cli/getting-started-cli.html) for this step to work.

After that if you do `kubectl config current-context` it should display you are connected to your new cluster.

## Create the cluster managing service account

Using the kubectl context against the new cluster, create the service account and its token:

```shell
$ kubectl create ns atlas-upgrade-test-tokens
$ kubectl -n atlas-upgrade-test-tokens create serviceaccount atlas-operator-upgrade-test
$ oc create token --duration=87600h -n atlas-upgrade-test-tokens atlas-operator-upgrade-test >token.txt
```

Give this service account enough permissions, currently this is cluster-admin:

```shell
$ oc adm policy add-cluster-role-to-user cluster-admin system:serviceaccount:atlas-upgrade-test-tokens:atlas-operator-upgrade-test
```

Copy & Paste token.txt into the `OPENSHIFT_UPGRADE_TOKEN` secret in Github Actions.

Run `kubectl cluster-info` Eg:

```shell
% kubectl cluster-info
Kubernetes control plane is running at https://***somehostname***.com:6443
...
```

And use the URL there to set `OPENSHIFT_UPGRADE_SERVER_API` so that openshift upgrade tests to run successfully.

## Troubleshooting

### Major version issues when executing the "Create Release Branch" workflow

The release creation will fail if the file `major-version` contents does not match the major version to be released. This file explicitly means the upcoming release is for a particular major version, with potential breaking changes. This allows us to:

1. Notice if we forgot to update the `major-version` file before releasing the next major version.
2. Notice if we tried to re-release an older major version when the code is already prepared for the next major version.
3. Skip some tests, like `helm update`, when crossing from one major version to the next, as such test is not expected to work across incompatible major version upgrades.

If the create release branch job fails due an error such as `Bad major version for X... expected Y..`, review whether or not the `major-version` file was updated as expected. Check as well you are not trying to release a patch for the older major version from the new major version codebase.
