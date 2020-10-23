package uninstall

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"

	"github.com/kyma-project/cli/internal/cli"
	"github.com/kyma-project/cli/internal/kube"
	"github.com/spf13/cobra"

	"github.com/kyma-incubator/hydroform/installation-poc/pkg/installation"
)

type command struct {
	opts *Options
	cli.Command
}

//NewCmd creates a new kyma command
func NewCmd(o *Options) *cobra.Command {

	cmd := command{
		Command: cli.Command{Options: o.Options},
		opts:    o,
	}

	cobraCmd := &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstalls Kyma on a running Kubernetes cluster.",
		Long:    `Use this command to uninstall Kyma on a running Kubernetes cluster.`,
		RunE:    func(_ *cobra.Command, _ []string) error { return cmd.Run() },
		Aliases: []string{"i"},
	}

	cobraCmd.Flags().StringVarP(&o.ComponentsYaml, "components", "c", "", "Path to a YAML file with component list to override.")
	return cobraCmd
}

//Run runs the command
func (cmd *command) Run() error {
	var err error
	if cmd.K8s, err = kube.NewFromConfig("", cmd.KubeconfigPath); err != nil {
		return errors.Wrap(err, "Could not initialize the Kubernetes client. Make sure your kubeconfig is valid")
	}

	var componentsContent string
	if cmd.opts.ComponentsYaml != "" {
		data, err := ioutil.ReadFile(cmd.opts.ComponentsYaml)
		if err != nil {
			return fmt.Errorf("Failed to read installation CR file: %v", err)
		}
		componentsContent = string(data)
	}

	installer, err := installation.NewInstallation(componentsContent, "", "dsadas")
	if err != nil {
		return err
	}

	err = installer.StartKymaUninstallation(cmd.K8s.RestConfig())
	if err != nil {
		return err
	}

	log.Println("Kyma uninstalled!")

	return nil
}
