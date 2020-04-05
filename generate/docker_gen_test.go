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

package generate

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerGenerator_Generate(t *testing.T) {
	dockerGen := DockerGenerator{
		ImageName: "test-image",
		Tag:       "v0.1.0",
		Filepath:  "../../testing/dummy/Dockerfile_Test",
	}
	_, err := ioutil.ReadFile("../../testing/dummy/spec_docker_gen_test.yaml")
	assert.Equal(t, nil, err)

	_, err = dockerGen.Generate()
	assert.Equal(t, nil, err)
}
