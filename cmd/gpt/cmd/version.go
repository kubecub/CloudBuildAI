// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/kubecub/CloudBuildAI/pkg/version"
)

var (
	shortPrint bool
	output     string
)
var kubecubErr error

func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Print version info",
		Args:    cobra.NoArgs,
		Example: `kubecub version`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate validates the provided options.
			if output != "" && output != "yaml" && output != "json" {
				return fmt.Errorf("output format must be yaml or json")
			}
			if shortPrint {
				fmt.Println(version.Get().String())
				return nil
			}
			return PrintInfo()
		},
	}
	versionCmd.Flags().BoolVar(&shortPrint, "short", false, "If true, print just the version number.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "yaml", "choose `yaml` or `json` format to print version info")
	return versionCmd
}

func PrintInfo() error {
	OutputInfo := &version.Output{}
	OutputInfo.kubecubVersion = version.Get()

	if err := PrintToStd(OutputInfo); err != nil {
		return err
	}
	return nil
}

func PrintToStd(OutputInfo *version.Output) error {
	var (
		marshalled []byte
		err        error
	)
	switch output {
	case "yaml":
		marshalled, err = yaml.Marshal(&OutputInfo)
		if err != nil {
			return fmt.Errorf("fail to marshal yaml: %w", err)
		}
		fmt.Println(string(marshalled))
	case "json":
		marshalled, err = json.Marshal(&OutputInfo)
		if err != nil {
			return fmt.Errorf("fail to marshal json: %w", err)
		}
		fmt.Println(string(marshalled))
	default:
		// There is a bug in the program if we hit this case.
		// However, we follow a policy of never panicking.
		return fmt.Errorf("versionOptions were not validated: --output=%q should have been rejected", output)
	}
	return kubecubErr
}
