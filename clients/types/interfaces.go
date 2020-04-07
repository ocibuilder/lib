package types

import (
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
)

// Overlay is the overlay interface for handling of overylaying onto specification files
type Overlay interface {
	// Apply applies an overlay in your implementing struct
	Apply() ([]byte, error)
}

// ContextReader provides an interface for reading from multiple different build contexts
type ContextReader interface {
	Read() (io.ReadCloser, error)
}

// SpecGenerator provides an interface for spec generation for ocibuilder.yaml specification files
type SpecGenerator interface {
	Generate() ([]byte, error)
}

// BuilderClient is the client interface for the ocibuilder
type BuilderClient interface {
	ImageBuild(options BuildOptions) (BuildResponse, error)
	ImagePull(options PullOptions) (PullResponse, error)
	ImagePush(options PushOptions) (PushResponse, error)
	ImageRemove(options RemoveOptions) (RemoveResponse, error)
	ImageInspect(imageId string) (types.ImageInspect, error)
	ImageHistory(imageId string) ([]image.HistoryResponseItem, error)
	RegistryLogin(options LoginOptions) (LoginResponse, error)
	GenerateAuthRegistryString(auth types.AuthConfig) string
}
