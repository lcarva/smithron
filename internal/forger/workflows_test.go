package forger

import (
	"context"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestWorkflows(t *testing.T) {
	p := Plan{
		Name:    "hello-world",
		Image:   "docker.io/busybox:latest",
		Command: []string{"echo", "Hello, World!"},
		Env: map[string]string{
			"GREETING": "Hello",
		},
	}

	ctx := context.Background()

	tf, err := GetForgerForTarget("argo-workflows")
	require.NoError(t, err)
	require.IsType(t, ArgoWorkflowsForger{}, tf)

	got, err := tf.Forge(ctx, p)
	require.NoError(t, err)
	snaps.MatchYAML(t, got)
}
