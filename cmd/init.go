// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cumulusware/gossamer/internal/config"
	"github.com/cumulusware/gossamer/internal/generator"
	"github.com/cumulusware/gossamer/internal/prompts"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Go web application project",
	Long: `Initialize a new Go web application project with clean architecture.

This command will create a new directory with the project name and scaffold
a complete Go web application with the following features:
- Clean architecture with domain-driven design
- User authentication and authorization
- Database integration with PostgreSQL
- Optional HTMX for dynamic interactions
- Optional REST API endpoints
- Development tooling (Air, Justfile, Docker Compose)
- Security best practices (CSRF, password hashing, sessions)`,
	Args: cobra.MaximumNArgs(1),
	Run:  runInit,
}

var (
	flagForce bool
	flagDry   bool
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolVarP(&flagForce, "force", "f", false, "Force creation even if directory exists")
	initCmd.Flags().BoolVar(&flagDry, "dry-run", false, "Show what would be created without actually creating files")
}

func runInit(cmd *cobra.Command, args []string) {
	var projectName string

	if len(args) > 0 {
		projectName = args[0]
	}

	// Get project configuration through interactive prompts
	projectConfig, err := prompts.GetProjectConfig(projectName)
	if err != nil {
		fmt.Fprintf(os.Stderr, color.RedString("Error getting project configuration: %v\n"), err)
		os.Exit(1)
	}

	// Validate project directory
	projectPath := filepath.Join(".", projectConfig.Name)
	if !flagForce && !flagDry {
		if _, err := os.Stat(projectPath); err == nil {
			fmt.Fprintf(os.Stderr, color.RedString("Error: Directory '%s' already exists. Use --force to overwrite.\n"), projectPath)
			os.Exit(1)
		}
	}

	// Generate project
	gen := generator.New(projectConfig)

	if flagDry {
		fmt.Println(color.YellowString("🔍 Dry run mode - showing what would be created:"))
		files := gen.GetFileList()
		for _, file := range files {
			fmt.Printf("  📄 %s\n", file)
		}
		fmt.Printf("\n📊 Total files: %d\n", len(files))
		return
	}

	fmt.Println(color.CyanString("🚀 Creating new Go web application project..."))

	if err := gen.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, color.RedString("Error generating project: %v\n"), err)
		os.Exit(1)
	}

	// Print success message and next steps
	printSuccessMessage(projectConfig)
}

func printSuccessMessage(config *config.ProjectConfig) {
	fmt.Printf(color.GreenString("\n✅ Project '%s' created successfully!\n"), config.Name)

	fmt.Println(color.CyanString("\n📁 Project structure:"))
	fmt.Printf("  %s/\n", config.Name)
	fmt.Println("  ├── cmd/server/          # Application entry point")
	fmt.Println("  ├── internal/            # Private application code")
	fmt.Println("  │   ├── app/            # Application setup")
	fmt.Println("  │   ├── domain/         # Business logic")
	fmt.Println("  │   ├── infrastructure/ # External concerns")
	fmt.Println("  │   └── adapters/       # Interface adapters")
	fmt.Println("  ├── static/             # Static assets")
	fmt.Println("  └── config/             # Configuration")

	fmt.Println(color.MagentaString("\n🌐 Features included:"))
	fmt.Println("  ✅ User authentication and authorization")
	fmt.Println("  ✅ PostgreSQL database with migrations")
	fmt.Println("  ✅ Clean architecture with domain-driven design")
	if config.IncludeHTMX {
		fmt.Println("  ✅ HTMX for dynamic interactions")
	}
	if config.IncludeAPI {
		fmt.Println("  ✅ REST API endpoints")
	}
	fmt.Printf("  ✅ Security best practices (CSRF, sessions, password hashing)\n")
	fmt.Printf("  ✅ Development tooling (Air, Justfile, Docker Compose)\n")
	fmt.Printf("  ✅ Tailwind CSS for styling\n")
	fmt.Printf("  ✅ Example tests\n")

	fmt.Println(color.BlueString("\n📚 Useful commands:"))
	fmt.Println("  just                     # Show all available tasks")
	fmt.Println("  just dev                 # Start development server")
	fmt.Println("  just test                # Run tests")
	fmt.Println("  just db-migrate          # Run database migrations")

	fmt.Println(color.YellowString("\n🚀 Next steps:"))
	fmt.Printf("  1. cd %s\n", config.Name)
	fmt.Println("  2. just tidy")
	fmt.Println("  3. git init")
	fmt.Println("  4. cp config/env.template config/.env")
	fmt.Println("  5. Edit config/.env with your settings")
	fmt.Println("  6. just db-up")
	fmt.Println("  7. just db-migrate")
	fmt.Println("  8. just dev")

	fmt.Println(color.GreenString("\n🎉 Happy coding!"))
	fmt.Println(color.GreenString("\n✨ Notice how both Gossamer and your generated project use Justfile for consistency!"))
}
