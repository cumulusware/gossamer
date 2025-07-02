// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gossamer",
	Short: "A Go web application scaffolding tool",
	Long: color.New(color.FgCyan).Sprint(`
   ____
  / ___| ___  ___ ___  __ _ _ __ ___   ___ _ __
 | |  _ / _ \/ __/ __|/ _` + "`" + ` | '_ ` + "`" + ` _ \ / _ \ '__|
 | |_| | (_) \__ \__ \ (_| | | | | | |  __/ |
  \____|\___/|___/___/\__,_|_| |_| |_|\___|_|

Gossamer is a CLI tool for scaffolding modern Go web applications
with clean architecture, built-in authentication, and optional features
like HTMX and REST APIs.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, color.RedString("Error: %v\n"), err)
		os.Exit(1)
	}
}
