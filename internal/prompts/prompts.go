// Package prompts provide the prompts to the user when initializing a new project.
package prompts

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"

	"github.com/cumulusware/gossamer/internal/config"
)

func GetProjectConfig(projectName string) (*config.ProjectConfig, error) {
	cfg := config.NewProjectConfig()

	fmt.Println(color.CyanString("🔧 Let's configure your new Go web application!"))
	fmt.Println()

	// If project name not provided, prompt for it
	if projectName == "" {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What is the name of your project?").
					Description("This will be used as the directory name and default module path").
					Placeholder("my-web-app").
					Value(&cfg.Name).
					Validate(validateProjectName),
			),
		)

		if err := form.Run(); err != nil {
			return nil, err
		}
	} else {
		cfg.Name = projectName
		if err := validateProjectName(cfg.Name); err != nil {
			return nil, err
		}
	}

	cfg.EnvName = strings.ToUpper(cfg.Name)

	// Module path
	defaultModule := strings.ReplaceAll(cfg.Name, "-", "")
	cfg.ModulePath = defaultModule

	log.Printf("Module path before user input = %s", cfg.ModulePath)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the Go module path?").
				Description("This will be used in go.mod (e.g., github.com/username/project-name)").
				Value(&cfg.ModulePath).
				Validate(validateModulePath),

			huh.NewInput().
				Title("Brief description of your project (optional):").
				Description("This will be used in README.md and comments").
				Placeholder("A modern Go web application").
				Value(&cfg.Description),

			huh.NewInput().
				Title("Author name (optional):").
				Description("This will be used in generated files and documentation").
				Value(&cfg.Author),
		),
	)

	log.Printf("Module path after user input = %s", cfg.ModulePath)

	if err := form.Run(); err != nil {
		return nil, err
	}

	// Set default description if empty
	if cfg.Description == "" {
		cfg.Description = "A modern Go web application"
	}

	fmt.Println(color.MagentaString("\n🎛️ Feature Selection:"))

	// Feature selection form
	featureForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Include HTMX for dynamic interactions?").
				Description("HTMX allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML").
				Value(&cfg.IncludeHTMX).
				Affirmative("Yes").
				Negative("No"),

			huh.NewConfirm().
				Title("Include REST API endpoints?").
				Description("Creates /api/v1/* endpoints for building APIs alongside your web interface").
				Value(&cfg.IncludeAPI).
				Affirmative("Yes").
				Negative("No"),

			huh.NewSelect[string]().
				Title("Choose database:").
				Description("Currently only PostgreSQL is supported").
				Options(
					huh.NewOption("PostgreSQL", "postgresql"),
				).
				Value(&cfg.DatabaseType),
		),
	)

	if err := featureForm.Run(); err != nil {
		return nil, err
	}

	fmt.Println(color.GreenString("\n📋 Configuration Summary:"))
	fmt.Printf("  Project: %s\n", cfg.Name)
	fmt.Printf("  Module:  %s\n", cfg.ModulePath)
	fmt.Printf("  HTMX:    %s\n", boolToYesNo(cfg.IncludeHTMX))
	fmt.Printf("  API:     %s\n", boolToYesNo(cfg.IncludeAPI))
	fmt.Printf("  Database: %s\n", cfg.DatabaseType)

	// Confirmation
	confirmed := true
	confirmForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Create project with these settings?").
				Value(&confirmed).
				Affirmative("Yes, create project").
				Negative("No, cancel"),
		),
	)

	if err := confirmForm.Run(); err != nil {
		return nil, err
	}

	if !confirmed {
		return nil, fmt.Errorf("project creation cancelled")
	}

	return &cfg, nil
}

func validateProjectName(val string) error {
	if val == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	// Check for valid directory name
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`, val)
	if !matched {
		return fmt.Errorf(
			"project name must start with letter/number and contain only letters, numbers, hyphens, and underscores",
		)
	}

	// Check for reserved names
	reserved := []string{"go", "test", "main", "src", "pkg", "cmd", "internal"}
	for _, r := range reserved {
		if strings.EqualFold(val, r) {
			return fmt.Errorf("'%s' is a reserved name", val)
		}
	}

	return nil
}

func validateModulePath(val string) error {
	if val == "" {
		return fmt.Errorf("module path cannot be empty")
	}

	// Basic validation for Go module path
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9][a-zA-Z0-9._/-]*[a-zA-Z0-9]$`, val)
	if !matched {
		return fmt.Errorf("invalid module path format")
	}

	return nil
}

func boolToYesNo(b bool) string {
	if b {
		return color.GreenString("Yes")
	}
	return color.RedString("No")
}
