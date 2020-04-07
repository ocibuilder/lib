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

package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"

	docker "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/ocibuilder/lib/clients/types"
	"github.com/sirupsen/logrus"
)

// Client is the client used for building with Docker using Beval
type Client struct {
	APIClient client.APIClient
	Logger    *logrus.Logger
}

// ImageBuild conducts an image build with Docker using Beval
func (cli Client) ImageBuild(options types.BuildOptions) (types.BuildResponse, error) {
	apiCli := cli.APIClient
	res, err := apiCli.ImageBuild(options.Ctx, options.Context, options.ImageBuildOptions)
	if err != nil {
		return types.BuildResponse{}, err
	}
	return types.BuildResponse{
		ImageBuildResponse: res,
	}, nil
}

// ImagePull conducts an image pull with Docker using Beval
func (cli Client) ImagePull(options types.PullOptions) (types.PullResponse, error) {
	apiCli := cli.APIClient
	res, err := apiCli.ImagePull(options.Ctx, options.Ref, options.ImagePullOptions)
	if err != nil {
		return types.PullResponse{}, err
	}
	return types.PullResponse{
		Body: res,
	}, nil
}

// ImagePush conducts an image push with Docker using Beval
func (cli Client) ImagePush(options types.PushOptions) (types.PushResponse, error) {
	apiCli := cli.APIClient
	res, err := apiCli.ImagePush(options.Ctx, options.Ref, options.ImagePushOptions)
	if err != nil {
		return types.PushResponse{}, err
	}
	return types.PushResponse{
		Body: res,
	}, nil
}

// ImageRemove conducts an image remove with Docker using Beval
func (cli Client) ImageRemove(options types.RemoveOptions) (types.RemoveResponse, error) {
	apiCli := cli.APIClient
	res, err := apiCli.ImageRemove(options.Ctx, options.Image, options.ImageRemoveOptions)
	if err != nil {
		return types.RemoveResponse{}, err
	}
	return types.RemoveResponse{
		Response: res,
	}, nil
}

// ImageInspect conducts an inspect of a built image with Docker using Beval
func (cli Client) ImageInspect(imageId string) (docker.ImageInspect, error) {
	apiCli := cli.APIClient
	res, _, err := apiCli.ImageInspectWithRaw(context.Background(), imageId)
	if err != nil {
		return docker.ImageInspect{}, err
	}
	return res, nil
}

// ImageHistory conducts an image history call of a built image with Docker using Beval
func (cli Client) ImageHistory(imageId string) ([]image.HistoryResponseItem, error) {
	apiCli := cli.APIClient
	res, err := apiCli.ImageHistory(context.Background(), imageId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RegistryLogin conducts a registry login with Docker using Beval
func (cli Client) RegistryLogin(options types.LoginOptions) (types.LoginResponse, error) {
	apiCli := cli.APIClient
	res, err := apiCli.RegistryLogin(options.Ctx, options.AuthConfig)
	if err != nil {
		return types.LoginResponse{}, err
	}
	return types.LoginResponse{
		AuthenticateOKBody: res,
	}, nil
}

// GenerateAuthRegistryString generates the auth registry string for pushing and pulling images targeting the Docker daemon
func (cli Client) GenerateAuthRegistryString(auth docker.AuthConfig) string {
	encodedJSON, err := json.Marshal(auth)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error trying to marshall auth config")
	}
	return base64.URLEncoding.EncodeToString(encodedJSON)
}
