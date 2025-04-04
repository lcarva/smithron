package forger

import (
	"context"

	yaml "sigs.k8s.io/yaml/goyaml.v2"
)

type GitHubForger struct{}

func (GitHubForger) Forge(_ context.Context, plan Plan) ([]byte, error) {
	env := map[string]string{}
	for name, value := range plan.Env {
		env[name] = value
	}

	var entrypoint string
	var args []string
	for i, v := range plan.Command {
		if i == 0 {
			entrypoint = v
		} else {
			args = append(args, v)
		}
	}

	action := gitHubAction{
		Name: plan.Name,
		Runs: gitHubActionRuns{
			Using:      "docker",
			Image:      plan.Image,
			Env:        env,
			Entrypoint: entrypoint,
			Args:       args,
		},
	}

	return yaml.Marshal(action)
}

// https://docs.github.com/en/actions/sharing-automations/creating-actions/metadata-syntax-for-github-actions
type gitHubAction struct {
	Name string           `json:"name"`
	Runs gitHubActionRuns `json:"runs"`
}

type gitHubActionRuns struct {
	Using      string            `json:"using"`
	Image      string            `json:"image"`
	Entrypoint string            `json:"entrypoint"`
	Args       []string          `json:"args"`
	Env        map[string]string `json:"env"`
}

func init() {
	registerForger("github", func() Forger { return GitHubForger{} })
}
