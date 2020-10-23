package kyma

import (
	alphaInstall "github.com/kyma-project/cli/cmd/kyma/alpha/install"
	alphaUninstall "github.com/kyma-project/cli/cmd/kyma/alpha/uninstall"
	"github.com/kyma-project/cli/cmd/kyma/apply"
	"github.com/kyma-project/cli/cmd/kyma/completion"
	"github.com/kyma-project/cli/cmd/kyma/console"
	"github.com/kyma-project/cli/cmd/kyma/create"
	initial "github.com/kyma-project/cli/cmd/kyma/init"
	"github.com/kyma-project/cli/cmd/kyma/install"
	"github.com/kyma-project/cli/cmd/kyma/test"
	"github.com/kyma-project/cli/cmd/kyma/test/definitions"
	del "github.com/kyma-project/cli/cmd/kyma/test/delete"
	"github.com/kyma-project/cli/cmd/kyma/test/list"
	"github.com/kyma-project/cli/cmd/kyma/test/logs"
	"github.com/kyma-project/cli/cmd/kyma/test/run"
	"github.com/kyma-project/cli/cmd/kyma/test/status"
	"github.com/kyma-project/cli/cmd/kyma/version"

	"github.com/kyma-project/cli/cmd/kyma/alpha"
	"github.com/kyma-project/cli/cmd/kyma/upgrade"
	"github.com/kyma-project/cli/internal/cli"
	"github.com/spf13/cobra"
)

//NewCmd creates a new kyma CLI command
func NewCmd(o *cli.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kyma",
		Short: "Controls a Kyma cluster.",
		Long: `Kyma is a flexible and easy way to connect and extend enterprise applications in a cloud-native world.
Kyma CLI allows you to install, test, and manage Kyma.

`,
		// Affects children as well
		SilenceErrors: false,
		SilenceUsage:  true,
	}

	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "Displays details of actions triggered by the command.")
	cmd.PersistentFlags().BoolVar(&o.NonInteractive, "non-interactive", false, "Enables the non-interactive shell mode.")
	cmd.PersistentFlags().BoolVar(&o.CI, "ci", false, "Enables the CI mode to run on CI/CD systems.")
	// Kubeconfig env var and default paths are resolved by the kyma k8s client using the k8s defined resolution strategy.
	cmd.PersistentFlags().StringVar(&o.KubeconfigPath, "kubeconfig", "", `Specifies the path to the kubeconfig file. By default, Kyma CLI uses the KUBECONFIG environment variable or "/$HOME/.kube/config" if the variable is not set.`)
	cmd.PersistentFlags().BoolP("help", "h", false, "Displays help for the command.")

	alphaCmd := alpha.NewCmd()
	alphaCmd.AddCommand(alphaInstall.NewCmd(alphaInstall.NewOptions(o)))
	alphaCmd.AddCommand(alphaUninstall.NewCmd(alphaUninstall.NewOptions(o)))

	cmd.AddCommand(
		version.NewCmd(version.NewOptions(o)),
		completion.NewCmd(),
		install.NewCmd(install.NewOptions(o)),
		alphaCmd,
		console.NewCmd(console.NewOptions(o)),
		upgrade.NewCmd(upgrade.NewOptions(o)),
		create.NewCmd(o),
	)

	testCmd := test.NewCmd()
	testRunCmd := run.NewCmd(run.NewOptions(o))
	testStatusCmd := status.NewCmd(status.NewOptions(o))
	testDeleteCmd := del.NewCmd(del.NewOptions(o))
	testListCmd := list.NewCmd(list.NewOptions(o))
	testDefsCmd := definitions.NewCmd(definitions.NewOptions(o))
	testLogsCmd := logs.NewCmd(logs.NewOptions(o))
	testCmd.AddCommand(testRunCmd, testStatusCmd, testDeleteCmd, testListCmd, testDefsCmd, testLogsCmd)
	cmd.AddCommand(testCmd)

	cmd.AddCommand(
		initial.NewCmd(o),
		apply.NewCmd(o),
	)

	return cmd
}
