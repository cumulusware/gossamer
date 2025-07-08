// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package config

import (
	"testing"
)

func TestNewProjectConfig(t *testing.T) {
	cfg := NewProjectConfig()

	if cfg.DatabaseType != "postgresql" {
		t.Errorf("Expected default DatabaseType to be 'postgresql', got '%s'", cfg.DatabaseType)
	}

	// Other fields should be zero values initially
	if cfg.Name != "" {
		t.Errorf("Expected Name to be empty initially, got '%s'", cfg.Name)
	}

	if cfg.ModulePath != "" {
		t.Errorf("Expected ModulePath to be empty initially, got '%s'", cfg.ModulePath)
	}
}

func TestGetDependencies(t *testing.T) {
	cfg := NewProjectConfig()
	deps := cfg.GetDependencies()

	expectedDeps := []string{
		"github.com/jackc/pgx/v5",
		"github.com/joho/godotenv",
		"github.com/justinas/nosurf",
		"github.com/pressly/goose/v3",
		"github.com/google/uuid",
		"golang.org/x/crypto",
	}

	if len(deps) != len(expectedDeps) {
		t.Fatalf("Expected %d dependencies, got %d", len(expectedDeps), len(deps))
	}

	for i, expected := range expectedDeps {
		if deps[i] != expected {
			t.Errorf("Expected dependency %d to be '%s', got '%s'", i, expected, deps[i])
		}
	}
}
