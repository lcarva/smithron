package forger

import (
	"context"

	v1alpha "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

type ArgoWorkflowsForger struct{}

func (ArgoWorkflowsForger) Forge(ctx context.Context, plan Plan) ([]byte, error) {
	var env []corev1.EnvVar

	for name, value := range plan.Env {
		env = append(env, corev1.EnvVar{
			Name:  name,
			Value: value,
		})
	}

	container := corev1.Container{
		Image: plan.Image,
		Command: plan.Command,
		Env:    env,
	}
	template := v1alpha.Template{
		Name: "run",
		Container: &container,
	}

	wft := v1alpha.WorkflowTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "argoproj.io/v1alpha1",
			Kind:       "WorkflowTemplate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: plan.Name,
		},
		Spec: v1alpha.WorkflowSpec{
			Templates: []v1alpha.Template{template},
		},
	}

	// wft.SetDefaults(ctx)

	return yaml.Marshal(wft)
}

func init() {
	registerForger("argo-workflows", func() Forger { return ArgoWorkflowsForger{} })
}
