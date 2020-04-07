package docker

import (
	"context"
	"io"
	"testing"

	docker "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/ocibuilder/lib/clients/types"
	"github.com/ocibuilder/lib/util"
	"github.com/stretchr/testify/assert"
)

func TestClient_ImageBuild(t *testing.T) {
	_, err := cli.ImageBuild(types.BuildOptions{})
	assert.Equal(t, nil, err)
}

func TestClient_ImagePull(t *testing.T) {
	_, err := cli.ImagePull(types.PullOptions{})
	assert.Equal(t, nil, err)
}

func TestClient_ImagePush(t *testing.T) {
	_, err := cli.ImagePush(types.PushOptions{})
	assert.Equal(t, nil, err)
}

func TestClient_ImageRemove(t *testing.T) {
	_, err := cli.ImageRemove(types.RemoveOptions{})
	assert.Equal(t, nil, err)
}

func TestClient_RegistryLogin(t *testing.T) {
	_, err := cli.RegistryLogin(types.LoginOptions{})
	assert.Equal(t, nil, err)
}

func TestClient_GenerateAuthRegistryString(t *testing.T) {
	authString := cli.GenerateAuthRegistryString(authConfig)
	assert.Equal(t, "eyJ1c2VybmFtZSI6InVzZXIiLCJwYXNzd29yZCI6InBhc3MifQ==", authString)
}

var cli = Client{

	Logger:    util.GetLogger(true),
	APIClient: testClient{},
}

var authConfig = docker.AuthConfig{
	Username: "user",
	Password: "pass",
}

func (t testClient) ImageBuild(ctx context.Context, context io.Reader, options docker.ImageBuildOptions) (docker.ImageBuildResponse, error) {
	return docker.ImageBuildResponse{
		Body:   nil,
		OSType: "",
	}, nil
}

func (t testClient) ImagePull(ctx context.Context, ref string, options docker.ImagePullOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (t testClient) ImagePush(ctx context.Context, ref string, options docker.ImagePushOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (t testClient) ImageRemove(ctx context.Context, image string, options docker.ImageRemoveOptions) ([]docker.ImageDeleteResponseItem, error) {
	return nil, nil
}

func (t testClient) RegistryLogin(ctx context.Context, auth docker.AuthConfig) (registry.AuthenticateOKBody, error) {
	return registry.AuthenticateOKBody{
		IdentityToken: "",
		Status:        "",
	}, nil
}

type testClient struct {
	client.APIClient
}
