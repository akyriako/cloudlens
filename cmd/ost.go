package cmd

import (
	"github.com/one2nc/cloudlens/internal"
	"github.com/spf13/cobra"
)

func ostCommand() *cobra.Command {

	command := cobra.Command{
		Use:   "ost",
		Short: "Select Openstack",
		Long:  "Selects Openstack as default cloud",
		Run: func(cmd *cobra.Command, args []string) {
			selectOST()
		},
	}

	command.Flags().StringVarP(&clouds, "clouds", "c", "$HOME/.config/openstack/clouds.yaml", "Read Openstack profile")
	//command.Flags().StringVarP(&region, "region", "r", "", "Read aws region")
	//
	//command.Flags().BoolVarP(&useLocalStack, "localstack", "l", false, "Use localsatck instead of AWS")
	//command.Flags().StringVarP(&localStackPort, "port", "", "4566", "Read localstack port")

	return &command
}

func selectOST() {
	cloudConfig.SelectedCloud = internal.OST
	cloudConfig.OSTConfig.CloudsFilePath = clouds
	//cloudConfig.AWSConfig.Region = region
	//cloudConfig.AWSConfig.UseLocalStack = useLocalStack
	//cloudConfig.AWSConfig.LocalStackPort = localStackPort

	//os.Setenv(internal.LOCALSTACK_PORT, cloudConfig.LocalStackPort)
	initView()
}
