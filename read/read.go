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

package read

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/artbegolli/yenv"
	"github.com/ghodss/yaml"
	"github.com/ocibuilder/api/apis/beval/v1alpha1"
	"github.com/ocibuilder/lib/overlay"
	"github.com/ocibuilder/lib/validate"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/sjson"
)

// Reader is the struct for the common reader library
type Reader struct {
	// Logger is the logrus logger pass to the reader
	Logger *logrus.Logger
}

// Read is responsible for reading in the specification files, either
// combined in beval.yaml or separated in login.yaml, build.yaml and push.yaml.
// The passed in OCIBuilderSpec reference is populated
// If a filepath is not specified the current working directory is used
func (r Reader) Read(spec *v1alpha1.OCIBuilderSpec, overlayPath string, filepaths ...string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	filepath := strings.Join(filepaths[:], "/")
	if filepath != "" {
		dir = filepath
	}
	r.Logger.WithField("filepath", dir+"/beval.yaml").Debugln("looking for beval.yaml")
	file, err := ioutil.ReadFile(dir + "/beval.yaml")
	if err != nil {
		r.Logger.Infoln("beval.yaml file not found, looking for individual specifications...")
		if err := r.readIndividualSpecs(spec, dir); err != nil {
			return errors.Wrap(err, "failed to read individual specs")
		}
	}

	if overlayPath != "" {
		r.Logger.WithField("overlayPath", overlayPath).Debugln("overlay path not empty - looking for overlay file")
		file, err = applyOverlay(file, overlayPath)
		if err != nil {
			return errors.Wrap(err, "failed to apply overlay to spec at path")
		}
	}

	if err = yaml.Unmarshal(file, spec); err != nil {
		return errors.Wrap(err, "failed to unmarshal spec at directory")
	}

	if err := validate.Validate(spec); err != nil {
		return errors.Wrap(err, "failed to validate spec at directory")
	}

	if err = yaml.Unmarshal(file, spec); err != nil {
		return errors.Wrap(err, "failed to unmarshal spec at directory")
	}

	if err := validate.Validate(spec); err != nil {
		return errors.Wrap(err, "failed to validate spec at directory")
	}

	if err = yaml.Unmarshal(file, spec); err != nil {
		return errors.Wrap(err, "failed to unmarshal spec at directory")
	}

	if err := validate.Validate(spec); err != nil {
		return errors.Wrap(err, "failed to validate spec at directory")
	}

	if spec.Params != nil {
		if err = r.applyParams(file, spec); err != nil {
			return errors.Wrap(err, "failed to apply params to spec")
		}
	}

	return nil
}

// readIndividualSpecs reads the individual specifications if a global
// beval.yaml is not found
func (r Reader) readIndividualSpecs(spec *v1alpha1.OCIBuilderSpec, path string) error {
	var loginSpec []v1alpha1.LoginSpec
	var buildSpec *v1alpha1.BuildSpec
	var pushSpec []v1alpha1.PushSpec

	r.Logger.Debugln("attempting to read individual specs as spec.yaml as not found")
	if file, err := ioutil.ReadFile(path + "/login.yaml"); err == nil {
		if err := yaml.Unmarshal(file, &loginSpec); err != nil {
			return errors.Wrap(err, "failed to unmarshal login.yaml")
		}
		spec.Login = loginSpec
	}
	if file, err := ioutil.ReadFile(path + "/build.yaml"); err == nil {
		if err := yaml.Unmarshal(file, &buildSpec); err != nil {
			return errors.Wrap(err, "failed to unmarshal build.yaml")
		}
		spec.Build = buildSpec
	}
	if file, err := ioutil.ReadFile(path + "/push.yaml"); err == nil {
		if err := yaml.Unmarshal(file, &pushSpec); err != nil {
			return errors.Wrap(err, "failed to unmarshal push.yaml")
		}
		spec.Push = pushSpec
	}
	return nil
}

// applyOverlay applys a ytt overlay to the specification
func applyOverlay(yamlTemplate []byte, overlayPath string) ([]byte, error) {

	yttOverlay := overlay.YttOverlay{
		Spec: yamlTemplate,
		Path: overlayPath,
	}

	overlayedSpec, err := yttOverlay.Apply()
	if err != nil {
		return nil, errors.Wrap(err, "unable to apply overlay to spec")
	}

	return overlayedSpec, nil
}

func (r Reader) applyParams(yamlObj []byte, spec *v1alpha1.OCIBuilderSpec) error {
	log := r.Logger

	// Yenv applies all environment variables in the form ${ENV_VARIABLE_HERE}
	yamlObj, err := yenv.ApplyEnvValues(yamlObj)
	if err != nil {
		return err
	}

	specJSON, err := yaml.YAMLToJSON(yamlObj)
	if err != nil {
		return err
	}

	log.WithField("number", len(spec.Params)).Debugln("found custom params in spec.yaml")
	for _, param := range spec.Params {
		if param.Value != "" {
			log.WithFields(logrus.Fields{
				"value": param.Value,
				"dest":  param.Dest,
			}).Debugln("setting param value at destination")

			if err := validate.ValidateParams(specJSON, param.Dest); err != nil {
				log.WithError(err).WithField("dest", param.Dest).Errorln("Error validating params, check that your param dest is valid")
				return err
			}

			tmp, err := sjson.SetBytes(specJSON, param.Dest, param.Value)
			if err != nil {
				return err
			}
			specJSON = tmp
		}
		if param.ValueFromEnvVariable != "" {
			log.WithFields(logrus.Fields{
				"value": param.Value,
				"dest":  param.Dest,
			}).Debugln("setting param value at destination")

			val := os.Getenv(param.ValueFromEnvVariable)
			if val == "" {
				log.Warn("env variable ", param.ValueFromEnvVariable, " is empty")
			}

			if err := validate.ValidateParams(specJSON, param.Dest); err != nil {
				log.WithError(err).WithField("dest", param.Dest).Errorln("Error validating params, check that your param dest is valid")
				return err
			}
			tmp, err := sjson.SetBytes(specJSON, param.Dest, val)
			if err != nil {
				return err
			}
			specJSON = tmp
		}
	}

	if err := json.Unmarshal(specJSON, spec); err != nil {
		return err
	}
	return nil
}
