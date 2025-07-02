// Copyright (c) 2024 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/cumulusware/gossamer/templates"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [type] [directory]",
	Short: "Initialize a new site of the given type in the given directory.",
	Long: `Initialize a new site of the given type in the given directory.
The new site can be of the type static, app, or api.
If the given directory already exists, gossamer will quit with an error.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		siteType := args[0]
		if siteType != "static" && siteType != "app" && siteType != "api" {
			fmt.Printf("The site type must be static, app, or api not %s.", siteType)
			os.Exit(1)
		}
		dir := args[1]
		err := os.Mkdir(dir, 0750)
		if errors.Is(err, fs.ErrExist) {
			fmt.Printf("The directory '%s' already exists.", dir)
			os.Exit(1)
		}
		if err != nil {
			log.Fatalf("Error creating directory '%s': %s", dir, err)
		}

		switch siteType {
		case "static":
			initStatic(dir)
		case "app":
			initApp(dir)
		case "api":
			initAPI(dir)
		}
	},
}

func initStatic(dir string) {
}

func initApp(dir string) {
	copyFileFS(templates.FS, "base/gitignore.gotxt", path.Join(dir, ".gitignore"))
	copyFileFS(templates.FS, "app/Justfile.gotxt", path.Join(dir, "Justfile"))
}

func initAPI(dir string) {
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copyFileFS(fileSystem fs.FS, input string, output string) {
	data, err := fs.ReadFile(fileSystem, input)
	checkError(err)
	err = os.WriteFile(output, data, 0644)
	checkError(err)
}
