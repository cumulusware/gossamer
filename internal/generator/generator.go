// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

// Package generator generates the files for the project initialization.
package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/fatih/color"

	"github.com/cumulusware/gossamer/internal/config"
)

type Generator struct {
	config      *config.ProjectConfig
	templatesFS fs.FS
}

func New(config *config.ProjectConfig) *Generator {
	return &Generator{
		config:      config,
		templatesFS: GetTemplatesFS(),
	}
}

func (g *Generator) Generate() error {
	projectPath := filepath.Join(".", g.config.Name)

	// Create project directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Get all project files (both templates and static)
	projectFiles := GetProjectFiles()

	for _, file := range projectFiles {
		// Check if this file should be included based on conditions
		if !g.shouldIncludeFile(file) {
			continue
		}

		fullPath := filepath.Join(projectPath, file.GetDestinationPath())

		// Create directory if it doesn't exist
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Generate file content
		content, err := g.generateFileContent(file)
		if err != nil {
			return fmt.Errorf(
				"failed to generate content for %s: %w",
				file.GetDestinationPath(),
				err,
			)
		}

		// Write file
		if err := os.WriteFile(fullPath, []byte(content), file.GetPermissions()); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}

		fmt.Printf("  📄 %s\n", color.GreenString(file.GetDestinationPath()))
	}

	return nil
}

func (g *Generator) GetFileList() []string {
	projectFiles := GetProjectFiles()
	var result []string

	for _, file := range projectFiles {
		if g.shouldIncludeFile(file) {
			result = append(result, file.GetDestinationPath())
		}
	}

	return result
}

func (g *Generator) shouldIncludeFile(file ProjectFile) bool {
	conditional := file.GetConditional()
	if conditional == "" {
		return true
	}

	// Use reflection to check the config field
	configValue := reflect.ValueOf(g.config).Elem()
	field := configValue.FieldByName(conditional)

	if !field.IsValid() {
		// If field doesn't exist, include the file
		return true
	}

	// If it's a boolean field, return its value
	if field.Kind() == reflect.Bool {
		return field.Bool()
	}

	// For other types, include if not zero value
	return !field.IsZero()
}

func (g *Generator) generateFileContent(file ProjectFile) (string, error) {
	// Read the file
	content, err := fs.ReadFile(g.templatesFS, file.GetSourcePath())
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", file.GetSourcePath(), err)
	}

	// If it's a static file, return content as-is
	if !file.IsTemplate() {
		return string(content), nil
	}

	// Process as Go template
	tmpl, err := template.New(filepath.Base(file.GetSourcePath())).Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", file.GetSourcePath(), err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, g.config); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", file.GetSourcePath(), err)
	}

	return buf.String(), nil
}
