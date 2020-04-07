/*
Copyright 2019 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	ctx "context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/ocibuilder/lib/command"
)

// BuildOptions are the build options for an ocibuilder build
type BuildOptions struct {
	// ImageBuildOptions are standard Docker API image build options
	types.ImageBuildOptions `json:"imageBuildOptions,inline"`
	// ContextPath is the path to the raw build context, used for Buildah builds
	ContextPath string `json:"contextPath"`
	// Ctx is the goroutine context
	Ctx ctx.Context `json:"ctx"`
	// Context is the docker tared build context
	Context io.Reader `json:"context"`
	// StorageDriver is a buildah flag for storage driver e.g. vfs
	StorageDriver string `json:"storageDriver"`
}

// BuildResponse is the build response from an ocibuilder build
type BuildResponse struct {
	// ImageBuildResponse is standard build response from the Docker API
	types.ImageBuildResponse `json:"imageBuildResponse,inline"`
	// Exec is part of the response for Buildah command executions
	Exec *command.Command `json:"exec,inline"`
	// Stderr is the stderr output stream used to stream buildah response
	Stderr io.ReadCloser `json:"stderr,inline"`
	// Finished is the flag to determine that the response has finished being read
	Finished bool
}

// PullOptions are the pull options for an ocibuilder pull
type PullOptions struct {
	// ImagePullOptions are the standard Docker API pull options
	types.ImagePullOptions `json:"imagePullOptions,inline"`
	// Ref is the reference image name to pull
	Ref string `json:"ref,inline"`
	// Ctx is the goroutine context
	Ctx ctx.Context `json:"ctx,inline"`
}

// PullResponse is the pull response from an ocibuilder pull
type PullResponse struct {
	// Body is the body of the response from an ocibuilder pull
	Body io.ReadCloser `json:"body,inline"`
	// Exec is part of the response for Buildah command executions
	Exec *command.Command `json:"exec,inline"`
	// Stderr is the stderr output stream used to stream buildah response
	Stderr io.ReadCloser `json:"stderr,inline"`
}

// PushOptions are the pull options for an ocibuilder push
type PushOptions struct {
	// ImagePushOptions are the standard Docker API push options
	types.ImagePushOptions `json:"imagePushOptions,inline"`
	// Ref is the reference image name to push
	Ref string `json:"ref,inline"`
	// Ctx is the goroutine context
	Ctx ctx.Context `json:"ctx,inline"`
}

// PushResponse is the push response from an ocibuilder push
type PushResponse struct {
	// Body is the body of the response from an ocibuilder push
	Body io.ReadCloser `json:"body,inline"`
	// Exec is part of the response for Buildah command executions
	Exec *command.Command `json:"exec,inline"`
	// Stderr is the stderr output stream used to stream buildah response
	Stderr io.ReadCloser `json:"stderr,inline"`
	// Finished is the flag to determine that the response has finished being read
	Finished bool
}

// RemoveOptions are the remove options for an ocibuilder remove
type RemoveOptions struct {
	// ImageRemoveOptions are the standard Docker API remove options
	types.ImageRemoveOptions `json:"imageRemoveOptions,inline"`
	// Image is the name of the image to remove
	Image string `json:"image,inline"`
	// Ctx is the goroutine context
	Ctx ctx.Context `json:"ctx,inline"`
}

// RemoveResponse is the response from an ocibuilder remove
type RemoveResponse struct {
	// Response are the responses from an image delete
	Response []types.ImageDeleteResponseItem `json:"response,inline"`
	// Exec is part of the response for Buildah command executions
	Exec *command.Command `json:"exec,inline"`
	// Stderr is the stderr output stream used to stream buildah response
	Stderr io.ReadCloser `json:"stderr,inline"`
}

// LoginOptions are the login options for an ocibuilder login
type LoginOptions struct {
	// AuthConfig is the standard auth config for the Docker API
	types.AuthConfig `json:"authConfig,inline"`
	// Ctx is the goroutine context
	Ctx ctx.Context `json:"ctx,inline"`
}

// LoginResponse is the login response from an ocibuilder login
type LoginResponse struct {
	// AuthenticateOKBody is the standar login response from the Docker API
	registry.AuthenticateOKBody
	// Exec is part of the response for Buildah command executions
	Exec *command.Command `json:"exec,inline"`
	// Stderr is the stderr output stream used to stream buildah response
	Stderr io.ReadCloser `json:"stderr,inline"`
}
