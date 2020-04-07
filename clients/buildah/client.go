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

package buildah

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	docker "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/ocibuilder/lib/clients/types"
	"github.com/ocibuilder/lib/command"
	"github.com/sirupsen/logrus"
)

// Client is the client used for building with Buildah using beval
type Client struct {
	Logger *logrus.Logger
}

// ImageBuild conducts an image build with Buildah using beval
func (cli Client) ImageBuild(options types.BuildOptions) (types.BuildResponse, error) {

	buildFlags := []command.Flag{
		{Name: "f", Value: options.Dockerfile, Short: true, OmitEmpty: true},
		{Name: "storage-driver", Value: options.StorageDriver, Short: false, OmitEmpty: true},
		{Name: "t", Value: options.Tags[0], Short: true, OmitEmpty: true},
	}

	if options.NoCache {
		buildFlags = append(buildFlags, command.Flag{Name: "no-cache", Value: "", Short: false, OmitEmpty: true})
	}

	for _, l := range options.Labels {
		buildFlags = append(buildFlags, command.Flag{Name: "label", Value: l, Short: false, OmitEmpty: true})
	}

	cmd := command.Builder("buildah").Command("bud").Flags(buildFlags...).Args(options.ContextPath).Build()
	cli.Logger.WithField("cmd", cmd).Debugln("executing build with command")

	stdout, stderr, err := execute(&cmd)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error building image...")
		return types.BuildResponse{}, err
	}
	return types.BuildResponse{
		ImageBuildResponse: docker.ImageBuildResponse{
			Body: stdout,
		},
		Exec:   &cmd,
		Stderr: stderr,
	}, nil
}

// ImagePull conducts an image pull with Buildah using beval
func (cli Client) ImagePull(options types.PullOptions) (types.PullResponse, error) {

	pullFlags := []command.Flag{
		// Buildah registry auth in format username[:password]
		{Name: "creds", Value: options.RegistryAuth, Short: false, OmitEmpty: true},
	}

	cmd := command.Builder("buildah").Command("pull").Flags(pullFlags...).Args(options.Ref).Build()
	cli.Logger.WithField("cmd", cmd).Debugln("executing pull with command")

	stdout, stderr, err := execute(&cmd)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error building image...")
		return types.PullResponse{}, err
	}
	return types.PullResponse{
		Body:   stdout,
		Exec:   &cmd,
		Stderr: stderr,
	}, nil
}

// ImagePush conducts an image push with Buildah using beval
func (cli Client) ImagePush(options types.PushOptions) (types.PushResponse, error) {

	pushFlags := []command.Flag{
		// Buildah registry auth in format username[:password]
		{Name: "creds", Value: options.RegistryAuth, Short: false, OmitEmpty: true},
	}

	cmd := command.Builder("buildah").Command("push").Flags(pushFlags...).Args(options.Ref).Build()
	cli.Logger.WithField("cmd", cmd).Debugln("executing push with command")

	stdout, stderr, err := execute(&cmd)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error building image...")
		return types.PushResponse{}, err
	}
	return types.PushResponse{
		Body:   stdout,
		Exec:   &cmd,
		Stderr: stderr,
	}, nil
}

// ImageRemove conducts an image remove with Buildah using beval
func (cli Client) ImageRemove(options types.RemoveOptions) (types.RemoveResponse, error) {

	cmd := command.Builder("buildah").Command("rmi").Args(options.Image).Build()
	cli.Logger.WithField("cmd", cmd).Debugln("executing remove with command")

	_, _, err := execute(&cmd)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error building image...")
		return types.RemoveResponse{}, err
	}
	return types.RemoveResponse{
		Response: []docker.ImageDeleteResponseItem{
			{
				Deleted: options.Image,
			},
		},
		Exec: &cmd,
	}, nil
}

// @TBC
// ImageInspect conducts an inspect of a build image with Buildah using beval
//
// The type currently returned by running buildah inspect varies greatly to docker image inspect
// issues have been created to manage this and have a more representative mapping between buildah and docker
func (cli Client) ImageInspect(imageId string) (docker.ImageInspect, error) {

	inspectFlag := command.Flag{
		Name: "type", Value: "image", Short: false, OmitEmpty: false,
	}

	cmd := command.Builder("buildah").Command("inspect").Flags(inspectFlag).Args(imageId).Build()
	cli.Logger.WithField("cmd", cmd).Debugln("executing remove with command")

	stdout, _, err := execute(&cmd)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error building image...")
		return docker.ImageInspect{}, err
	}
	var imageInspect docker.ImageInspect
	res, err := ioutil.ReadAll(stdout)
	if err != nil {
		return docker.ImageInspect{}, err
	}
	if err := json.Unmarshal(res, &imageInspect); err != nil {
		return docker.ImageInspect{}, err
	}
	return docker.ImageInspect{}, nil
}

// ImageHistory is TBC for buildah client
func (cli Client) ImageHistory(imageId string) ([]image.HistoryResponseItem, error) {
	return nil, nil
}

// RegistryLogin conducts a registry login with Buildah using beval
func (cli Client) RegistryLogin(options types.LoginOptions) (types.LoginResponse, error) {

	loginFlags := []command.Flag{
		{Name: "u", Value: options.Username, Short: true, OmitEmpty: true},
		{Name: "p", Value: options.Password, Short: true, OmitEmpty: true},
	}

	cmd := command.Builder("buildah").Command("login").Flags(loginFlags...).Args(options.ServerAddress).Build()
	cli.Logger.WithField("cmd", cmd).Debugln("executing login with command")

	_, _, err := execute(&cmd)
	if err != nil {
		cli.Logger.WithError(err).Errorln("error building image...")
		return types.LoginResponse{}, err
	}

	return types.LoginResponse{
		AuthenticateOKBody: registry.AuthenticateOKBody{
			Status: "login completed",
		},
		Exec: &cmd,
	}, nil
}

// GenerateAuthRegistryString generates the auth registry string for pushing and pulling images targeting Buildah
func (cli Client) GenerateAuthRegistryString(auth docker.AuthConfig) string {
	return fmt.Sprintf("%s:%s", auth.Username, auth.Password)
}

// Execute executes the buildah command. This function is mocked in buildah client tests.
var execute = func(cmd *command.Command) (stdout io.ReadCloser, stderr io.ReadCloser, err error) {
	return cmd.Exec()
}
