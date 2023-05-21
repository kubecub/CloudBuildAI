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

type rootOpts struct {

}

var rootOpt rootOpts

const (
	colorModeNever  = "never"
	colorModeAlways = "always"
)

var longRootCmdDescription = `cba automatic Kubernetes yaml and dockerfile generation via chatgpt
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "cba",
	Short:         "A tool to build, share and run any distributed applications.",
	Long:          longRootCmdDescription,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("cba-%s: %v", version.GetSingleVersion(), err)
		os.Exit(1)
	}
}


func init() {

}
