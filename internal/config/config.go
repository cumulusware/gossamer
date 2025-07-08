// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

// Package config provides the configuration information for a new project.
package config

import "time"

type ProjectConfig struct {
	Name         string
	EnvName      string
	ModulePath   string
	IncludeHTMX  bool
	IncludeAPI   bool
	DatabaseType string
	Author       string
	Description  string
	Year         int
}

func NewProjectConfig() ProjectConfig {
	return ProjectConfig{
		DatabaseType: "postgresql", // Default to PostgreSQL
		Year:         time.Now().Year(),
	}
}

func (pc *ProjectConfig) GetDependencies() []string {
	deps := []string{
		"github.com/jackc/pgx/v5",
		"github.com/joho/godotenv",
		"github.com/justinas/nosurf",
		"github.com/pressly/goose/v3",
		"github.com/google/uuid",
		"golang.org/x/crypto",
	}

	return deps
}
