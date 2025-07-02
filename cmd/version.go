// Copyright (c) 2024 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Long:  `Display Gossamer's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gossamer v0.0.1")
	},
}
