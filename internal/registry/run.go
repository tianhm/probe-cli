package registry

//
// Registers the `run' experiment.
//

import (
	"github.com/ooni/probe-cli/v3/internal/engine/experiment/run"
	"github.com/ooni/probe-cli/v3/internal/model"
)

func init() {
	allexperiments["run"] = &Factory{
		build: func(config interface{}) model.ExperimentMeasurer {
			return run.NewExperimentMeasurer(
				*config.(*run.Config),
			)
		},
		config:      &run.Config{},
		inputPolicy: model.InputStrictlyRequired,
	}
}