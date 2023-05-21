// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kubecub/CloudBuildAI/pkg/common"

	"github.com/kubecub/CloudBuildAI/pkg/log"
	"github.com/kubecub/CloudBuildAI/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type rootOpts struct {
	cfgFile              string
	debugModeOn          bool
	hideLogTime          bool
	hideLogPath          bool
	logToFile            bool
	colorMode            string
	remoteLoggerURL      string
	remoteLoggerTaskName string
}

var rootOpt rootOpts

const (
	colorModeNever  = "never"
	colorModeAlways = "always"
)

var longRootCmdDescription = `kubecub automatic Kubernetes yaml and dockerfile generation via chatgpt
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "kubecub",
	Short:         "A tool to build, share and run any distributed applications.",
	Long:          longRootCmdDescription,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("kubecub-%s: %v", version.GetSingleVersion(), err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(NewGenDocCommand(), NewVersionCmd())

	rootCmd.PersistentFlags().StringVar(&rootOpt.cfgFile, "config", "", "config file of kubecub tool (default is $HOME/.kubecub.json)")
	rootCmd.PersistentFlags().BoolVarP(&rootOpt.debugModeOn, "debug", "d", false, "turn on debug mode")
	rootCmd.PersistentFlags().BoolVarP(&rootCmd.SilenceUsage, "quiet", "q", false, "silence the usage when fail")
	rootCmd.PersistentFlags().BoolVar(&rootOpt.hideLogTime, "hide-time", false, "hide the log time")
	rootCmd.PersistentFlags().BoolVar(&rootOpt.hideLogPath, "hide-path", false, "hide the log path")
	rootCmd.PersistentFlags().BoolVar(&rootOpt.logToFile, "log-to-file", true, "write log message to disk")
	rootCmd.PersistentFlags().StringVar(&rootOpt.colorMode, "color", colorModeAlways, fmt.Sprintf("set the log color mode, the possible values can be %v", supportedColorModes))
	rootCmd.PersistentFlags().StringVar(&rootOpt.remoteLoggerURL, "remote-logger-url", "", "remote logger url, if not empty, will send log to this url")
	rootCmd.PersistentFlags().StringVar(&rootOpt.remoteLoggerTaskName, "task-name", "", "task name which will embedded in the remote logger header, only valid when --remote-logger-url is set")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.DisableAutoGenTag = true

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if rootOpt.cfgFile == "" {
		// Find home directory.
		rootOpt.cfgFile = filepath.Join(common.GetHomeDir(), ".kubecub.json")
	}
	// Use config file from the flag.
	// if not set config file, Search config in home directory with name ".kubecub.json" (without extension).
	//viper.AddConfigPath(home)
	viper.SetConfigFile(rootOpt.cfgFile)

	viper.AutomaticEnv() // read in environment variables that match

	if err := log.Init(log.LogOptions{
		LogToFile:            rootOpt.logToFile,
		Verbose:              rootOpt.debugModeOn,
		RemoteLoggerURL:      rootOpt.remoteLoggerURL,
		RemoteLoggerTaskName: rootOpt.remoteLoggerTaskName,
		DisableColor:         rootOpt.colorMode == colorModeNever,
	}); err != nil {
		panic(fmt.Sprintf("failed to init logger: %v\n", err))
	}
}
