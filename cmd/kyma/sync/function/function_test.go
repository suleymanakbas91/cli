package function

import (
	"github.com/kyma-project/cli/internal/cli"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestUpgradeFlags ensures that the provided command flags are stored in the options.
func TestUpgradeFlags(t *testing.T) {
	o := NewOptions(&cli.Options{})
	c := NewCmd(o)

	// test default flag values
	require.Empty(t, o.Name, "Default value for the --name flag not as expected")
	require.Empty(t, o.Namespace, "Default value for the --namespace flag not as expected.")
	require.Equal(t, "", o.OutputPath, "Default value for the --output flag not as expected.")

	// test passing flags
	err := c.ParseFlags([]string{
		"--output", "/fakepath",
		"--name", "test-name",
		"--namespace", "test-namespace",
	})

	require.NoError(t, err, "Parsing flags should not return an error")
	require.Equal(t, "/fakepath", o.OutputPath, "The parsed value for the --output flag not as expected.")
	require.Equal(t, "test-name", o.Name, "The parsed value for the --name flag not as expected.")
	require.Equal(t, "test-namespace", o.Namespace, "The parsed value for the --namespace flag not as expected.")
}