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
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/ocibuilder/gofeas"
	"github.com/ocibuilder/lib/store"
	"github.com/ocibuilder/lib/util"
	"github.com/ocibuilder/ocibuilder/pkg/apis/ocibuilder/v1alpha1"
	"github.com/ocibuilder/ocibuilder/testing/dummy"
	"github.com/stretchr/testify/assert"
)

func TestMetadataWriter_ParseResponseMetadata(t *testing.T) {
	dataTypes := []v1alpha1.MetadataType{
		v1alpha1.MetadataType("image"),
	}

	mw := MetadataWriter{
		Metadata: &v1alpha1.Metadata{
			StoreConfig: &v1alpha1.StoreConfig{},
			Key:         nil,
			Hostname:    "",
			Data:        dataTypes,
		},
		Logger: util.Logger,
	}
	err := mw.ParseMetadata("test-image", testClientMetadata{}, &v1alpha1.BuildProvenance{})
	assert.Equal(t, nil, err)

	record := mw.records[0]
	fingerprint := record.DerivedImage.DerivedImage.Fingerprint
	args := record.DerivedImage.DerivedImage.LayerInfo[0].Arguments
	cmd := *record.DerivedImage.DerivedImage.LayerInfo[0].Directive

	assert.Equal(t, expectedRecord.DerivedImage.DerivedImage.Fingerprint, fingerprint)
	assert.Equal(t, gofeas.ADD_LayerDirective, cmd)
	assert.Equal(t, "file:0eb5ea35741d23fe39cbac245b3a5d84856ed6384f4ff07d496369ee6d960bad", args)
}

func TestCreateAttestation(t *testing.T) {
	mw := MetadataWriter{
		Metadata: &v1alpha1.Metadata{
			StoreConfig: &v1alpha1.StoreConfig{},
			Key: &v1alpha1.SignKey{
				PlainPrivateKey: dummy.TestPrivKey,
				PlainPublicKey:  dummy.TestPubKey,
				Passphrase:      "",
			},
		},
		Logger: util.Logger,
	}
	record, err := mw.createAttestationRecord("image-digest")
	assert.Equal(t, nil, err)
	assert.True(t, strings.HasPrefix(record.Attestation.Attestation.PgpSignedAttestation.Signature, expectedPrefix))
}

func TestParseCreatedBy(t *testing.T) {
	testRunCmd := "/bin/sh -c mkdir /certs /certs/client && chmod 1777 /certs /certs/client"
	testStdCmd := "/bin/sh -c #(nop)  ENV _BUILDAH_STARTED_IN_USERNS= BUILDAH_ISOLATION=chroot"

	runLayer := parseCreatedBy(testRunCmd)
	assert.Equal(t, gofeas.RUN_LayerDirective, *runLayer.Directive)
	assert.Equal(t, "mkdir /certs /certs/client && chmod 1777 /certs /certs/client", runLayer.Arguments)

	envLayer := parseCreatedBy(testStdCmd)
	assert.Equal(t, gofeas.ENV_LayerDirective, *envLayer.Directive)
	assert.Equal(t, "_BUILDAH_STARTED_IN_USERNS= BUILDAH_ISOLATION=chroot", envLayer.Arguments)
}

var expectedRecord = store.Record{
	DerivedImage: &gofeas.V1beta1imageDetails{
		DerivedImage: &gofeas.ImageDerived{
			Fingerprint: &gofeas.ImageFingerprint{
				V1Name: "sha256-imageid",
				V2Blob: []string{"sha256-imageid2", "sha256-imageid"},
			},
			LayerInfo: []gofeas.ImageLayer{{
				Arguments: "ADD ./test .",
			}},
		},
	},
}

var expectedPrefix = `-----BEGIN PGP SIGNATURE-----`

func (t testClientMetadata) ImageBuild(options v1alpha1.OCIBuildOptions) (v1alpha1.OCIBuildResponse, error) {
	return v1alpha1.OCIBuildResponse{}, nil
}

func (t testClientMetadata) ImagePull(options v1alpha1.OCIPullOptions) (v1alpha1.OCIPullResponse, error) {
	return v1alpha1.OCIPullResponse{}, nil
}

func (t testClientMetadata) ImagePush(options v1alpha1.OCIPushOptions) (v1alpha1.OCIPushResponse, error) {
	return v1alpha1.OCIPushResponse{}, nil
}

func (t testClientMetadata) ImageRemove(options v1alpha1.OCIRemoveOptions) (v1alpha1.OCIRemoveResponse, error) {
	return v1alpha1.OCIRemoveResponse{}, nil
}

func (t testClientMetadata) ImageInspect(imageId string) (types.ImageInspect, error) {
	return types.ImageInspect{
		ID:          "sha256-imageid",
		RepoDigests: []string{"sha25-de3gie3st"},
	}, nil

}

func (t testClientMetadata) ImageHistory(imageId string) ([]image.HistoryResponseItem, error) {
	return []image.HistoryResponseItem{{
		CreatedBy: "/bin/sh -c #(nop) ADD file:0eb5ea35741d23fe39cbac245b3a5d84856ed6384f4ff07d496369ee6d960bad",
		ID:        "sha256-imageid2",
	}}, nil
}

func (t testClientMetadata) RegistryLogin(options v1alpha1.OCILoginOptions) (v1alpha1.OCILoginResponse, error) {
	return v1alpha1.OCILoginResponse{}, nil
}

func (t testClientMetadata) GenerateAuthRegistryString(auth types.AuthConfig) string {
	return ""
}

type testClientMetadata struct {
}
