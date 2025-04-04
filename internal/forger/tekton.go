package forger

import (
	"context"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

type TektonForger struct{}

func (TektonForger) Forge(ctx context.Context, plan Plan) ([]byte, error) {
	var env []corev1.EnvVar

	for name, value := range plan.Env {
		env = append(env, corev1.EnvVar{
			Name:  name,
			Value: value,
		})
	}

	step := v1.Step{
		Name:    "run",
		Image:   plan.Image,
		Command: plan.Command,
		Env:     env,
	}

	t := v1.Task{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "tekton.dev/v1",
			Kind:       "Task",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: plan.Name,
		},
		Spec: v1.TaskSpec{
			Steps: []v1.Step{step},
		},
	}

	t.SetDefaults(ctx)

	return yaml.Marshal(t)
}

func init() {
	registerForger("tekton", func() Forger { return TektonForger{} })
}
