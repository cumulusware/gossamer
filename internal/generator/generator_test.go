// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cumulusware/gossamer/internal/config"
)

func TestGeneratorBasicProject(t *testing.T) {
	// Create a test config
	cfg := &config.ProjectConfig{
		Name:         "test-project",
		ModulePath:   "github.com/test/test-project",
		Description:  "Test project",
		Author:       "Test Author",
		IncludeHTMX:  false,
		IncludeAPI:   false,
		DatabaseType: "postgresql",
	}

	// Create generator
	gen := New(cfg)

	// Get file list
	files := gen.GetFileList()

	// Basic checks
	if len(files) == 0 {
		t.Fatal("Generator should produce files")
	}

	// Check that essential files are present
	essentialFiles := []string{
		"go.mod",
		"README.md",
		"Justfile",
		"cmd/server/main.go",
		"internal/app/app.go",
	}

	for _, essential := range essentialFiles {
		found := false
		for _, file := range files {
			if file == essential {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Essential file %s not found in generated files", essential)
		}
	}
}

func TestGeneratorWithHTMX(t *testing.T) {
	cfg := &config.ProjectConfig{
		Name:         "test-htmx",
		ModulePath:   "github.com/test/test-htmx",
		Description:  "Test HTMX project",
		Author:       "Test Author",
		IncludeHTMX:  true,
		IncludeAPI:   false,
		DatabaseType: "postgresql",
	}

	gen := New(cfg)
	files := gen.GetFileList()

	// Check HTMX files are included
	htmxFiles := []string{
		"internal/adapters/handlers/web/htmx_handler.go",
		"internal/infrastructure/web/templates/partials/user_info.gohtml",
	}

	for _, htmxFile := range htmxFiles {
		found := false
		for _, file := range files {
			if file == htmxFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("HTMX file %s not found when IncludeHTMX=true", htmxFile)
		}
	}
}

func TestGeneratorWithAPI(t *testing.T) {
	cfg := &config.ProjectConfig{
		Name:         "test-api",
		ModulePath:   "github.com/test/test-api",
		Description:  "Test API project",
		Author:       "Test Author",
		IncludeHTMX:  false,
		IncludeAPI:   true,
		DatabaseType: "postgresql",
	}

	gen := New(cfg)
	files := gen.GetFileList()

	// Check API files are included
	apiFiles := []string{
		"internal/adapters/handlers/api/handlers.go",
		"internal/adapters/handlers/api/user_handler.go",
	}

	for _, apiFile := range apiFiles {
		found := false
		for _, file := range files {
			if file == apiFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("API file %s not found when IncludeAPI=true", apiFile)
		}
	}
}

func TestConditionalFileInclusion(t *testing.T) {
	tests := []struct {
		name         string
		config       *config.ProjectConfig
		shouldHave   []string
		shouldntHave []string
	}{
		{
			name: "minimal config",
			config: &config.ProjectConfig{
				Name:         "minimal",
				ModulePath:   "github.com/test/minimal",
				IncludeHTMX:  false,
				IncludeAPI:   false,
				DatabaseType: "postgresql",
			},
			shouldHave: []string{
				"go.mod",
				"internal/app/app.go",
			},
			shouldntHave: []string{
				"internal/adapters/handlers/web/htmx_handler.go",
				"internal/adapters/handlers/api/handlers.go",
			},
		},
		{
			name: "full features",
			config: &config.ProjectConfig{
				Name:         "full",
				ModulePath:   "github.com/test/full",
				IncludeHTMX:  true,
				IncludeAPI:   true,
				DatabaseType: "postgresql",
			},
			shouldHave: []string{
				"go.mod",
				"internal/app/app.go",
				"internal/adapters/handlers/web/htmx_handler.go",
				"internal/adapters/handlers/api/handlers.go",
			},
			shouldntHave: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := New(tt.config)
			files := gen.GetFileList()

			// Check files that should be present
			for _, shouldHave := range tt.shouldHave {
				found := false
				for _, file := range files {
					if file == shouldHave {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected file %s not found", shouldHave)
				}
			}

			// Check files that shouldn't be present
			for _, shouldntHave := range tt.shouldntHave {
				for _, file := range files {
					if file == shouldntHave {
						t.Errorf("Unexpected file %s found", shouldntHave)
					}
				}
			}
		})
	}
}

// Integration test that actually creates files
func TestGeneratorFileCreation(t *testing.T) {
	// Skip in CI unless we have proper temp directory setup
	if testing.Short() {
		t.Skip("Skipping file creation test in short mode")
	}

	cfg := &config.ProjectConfig{
		Name:         "test-creation",
		ModulePath:   "github.com/test/test-creation",
		Description:  "Test file creation",
		Author:       "Test Author",
		IncludeHTMX:  true,
		IncludeAPI:   true,
		DatabaseType: "postgresql",
	}

	// Create temp directory
	tempDir := t.TempDir() // Go 1.15+ method that automatically cleans up

	// Change to temp directory
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	gen := New(cfg)

	// Generate project
	if err := gen.Generate(); err != nil {
		t.Fatalf("Failed to generate project: %v", err)
	}

	// Check that files were actually created
	projectPath := filepath.Join(tempDir, cfg.Name)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		t.Fatal("Project directory was not created")
	}

	// Check a few key files
	keyFiles := []string{
		"go.mod",
		"cmd/server/main.go",
		"internal/app/app.go",
		"static/css/app.css", // static file
	}

	for _, keyFile := range keyFiles {
		filePath := filepath.Join(projectPath, keyFile)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Key file %s was not created", keyFile)
		}
	}

	// Verify template processing worked by checking go.mod content
	goModPath := filepath.Join(projectPath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}

	if !strings.Contains(string(content), cfg.ModulePath) {
		t.Errorf("go.mod doesn't contain expected module path %s", cfg.ModulePath)
	}
}

func TestTemplateFileVsStaticFile(t *testing.T) {
	// Test that our type system correctly identifies template vs static files
	projectFiles := GetProjectFiles()

	templateCount := 0
	staticCount := 0

	for _, file := range projectFiles {
		if file.IsTemplate() {
			templateCount++
			// Template files should have .gotmpl or .gohtml extensions or be known template files
			sourcePath := file.GetSourcePath()
			if !strings.HasSuffix(sourcePath, ".gotmpl") &&
				!strings.HasSuffix(sourcePath, ".gohtml") &&
				!isKnownTemplateFile(sourcePath) {
				t.Errorf("Template file %s has unexpected extension", sourcePath)
			}
		} else {
			staticCount++
			// Static files should not have template extensions
			sourcePath := file.GetSourcePath()
			if strings.HasSuffix(sourcePath, ".gotmpl") || strings.HasSuffix(sourcePath, ".gohtml") {
				t.Errorf("Static file %s has template extension", sourcePath)
			}
		}
	}

	if templateCount == 0 {
		t.Error("No template files found - this seems wrong")
	}
	if staticCount == 0 {
		t.Error("No static files found - this seems wrong")
	}

	t.Logf("Found %d template files and %d static files", templateCount, staticCount)
}

func isKnownTemplateFile(path string) bool {
	// Some files might be templates even without .gotmpl extension
	knownTemplates := []string{
		"web-templates/base.gohtml",
		"web-templates/home.gohtml",
		"web-templates/login.gohtml",
		"web-templates/register.gohtml",
		"web-templates/dashboard.gohtml",
	}

	for _, known := range knownTemplates {
		if strings.HasSuffix(path, known) {
			return true
		}
	}
	return false
}
