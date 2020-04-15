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

package oci

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	docker "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/ocibuilder/api/apis/beval/v1alpha1"
	"github.com/ocibuilder/lib/clients/types"
	"github.com/ocibuilder/lib/fake/testing/dummy"
	"github.com/ocibuilder/lib/util"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Build(t *testing.T) {
	builder := Builder{
		Logger:     util.GetLogger(true),
		Client:     testClient{},
		Provenance: []*v1alpha1.BuildProvenance{},
	}

	res := make(chan types.BuildResponse)
	errChan := make(chan error)
	finished := make(chan bool)

	defer func() {
		close(res)
		close(errChan)
		close(finished)
	}()

	go builder.Build(dummy.Spec, res, errChan, finished)

	for {
		select {
		case err := <-errChan:
			{
				assert.Equal(t, nil, err)
				return
			}
		case buildResponse := <-res:
			{
				res <- buildResponse
			}
		case fin := <-finished:
			{
				assert.True(t, fin, "expecting finished to be reached without an error on the error channel")
				return
			}
		}
	}

}

func TestBuilder_Build2(t *testing.T) {
	exists := true
	if _, err := os.Stat("./ocib"); os.IsNotExist(err) {
		exists = false
	}
	assert.False(t, exists, "There should be no context directory (./ocib) after a build has finished executing")
}

func TestBuilder_Pull(t *testing.T) {
}

func TestBuilder_Push(t *testing.T) {
}

func TestBuilder_Login(t *testing.T) {
}

func TestBuilder_Clean(t *testing.T) {
}

func TestBuilder_Purge(t *testing.T) {

}

func (t testClient) ImageBuild(options types.BuildOptions) (types.BuildResponse, error) {
	body := ioutil.NopCloser(strings.NewReader("image build response"))
	return types.BuildResponse{
		ImageBuildResponse: docker.ImageBuildResponse{
			Body:   body,
			OSType: "",
		},
	}, nil
}

func (t testClient) ImagePull(options types.PullOptions) (types.PullResponse, error) {
	return types.PullResponse{}, nil
}
func (t testClient) ImagePush(options types.PushOptions) (types.PushResponse, error) {
	return types.PushResponse{}, nil
}
func (t testClient) ImageRemove(options types.RemoveOptions) (types.RemoveResponse, error) {
	return types.RemoveResponse{}, nil
}
func (t testClient) ImageInspect(imageId string) (docker.ImageInspect, error) {
	return docker.ImageInspect{}, nil
}
func (t testClient) ImageHistory(imageId string) ([]image.HistoryResponseItem, error) {
	return nil, nil
}
func (t testClient) RegistryLogin(options types.LoginOptions) (types.LoginResponse, error) {
	return types.LoginResponse{}, nil
}
func (t testClient) GenerateAuthRegistryString(auth docker.AuthConfig) string {
	return ""
}

type testClient struct {
}
