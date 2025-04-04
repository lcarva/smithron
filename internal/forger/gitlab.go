package forger

import (
	"context"

	"sigs.k8s.io/yaml"
)

type GitLabForger struct{}

func (GitLabForger) Forge(_ context.Context, plan Plan) ([]byte, error) {
	j := map[string]gitLabJob{
		plan.Name: {
			Image: gitLabImage{
				Name:       plan.Image,
				Entrypoint: plan.Command,
			},
			Variables: plan.Env,
		},
	}

	return yaml.Marshal(j)
}

// https://docs.gitlab.com/ci/yaml/
type gitLabJob struct {
	Image     gitLabImage       `json:"image"`
	Variables map[string]string `json:"variables"`
}

type gitLabImage struct {
	Name       string   `json:"name"`
	Entrypoint []string `json:"entrypoint"`
}

func init() {
	registerForger("gitlab", func() Forger { return GitLabForger{} })
}
