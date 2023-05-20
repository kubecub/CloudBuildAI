// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"os"

	"github.com/kubecub/CloudBuildAI/pkg/version"
	"github.com/kubecub/log"
	"github.com/spf13/cobra"
)

var longRootCmdDescription = `sealer is a tool to seal application's all dependencies and Kubernetes
into sealer image by Kubefile, distribute this application anywhere via sealer image, 
and run it within any cluster with Clusterfile in one command.
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "sealer",
	Short:         "A tool to build, share and run any distributed applications.",
	Long:          longRootCmdDescription,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("sealer-%s: %v", version.GetSingleVersion(), err)
		os.Exit(1)
	}
}
