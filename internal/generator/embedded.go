// Copyright (c) 2024-2025 The gossamer developers. All rights reserved.
// Project site: https://github.com/cumulusware/gossamer
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package generator

import (
	"embed"
	"io/fs"
)

// Embed all template files into the binary
//
//go:embed templates/*
var templatesFS embed.FS

// GetTemplatesFS returns the embedded templates filesystem
func GetTemplatesFS() fs.FS {
	// Return the templates subdirectory
	templatesSubFS, err := fs.Sub(templatesFS, "templates")
	if err != nil {
		panic("failed to get templates sub-filesystem: " + err.Error())
	}
	return templatesSubFS
}

// TemplateFile represents a file that needs Go template processing
type TemplateFile struct {
	SourcePath      string      // Path in templates filesystem
	DestinationPath string      // Path in generated project
	Permissions     fs.FileMode // File permissions
	Conditional     string      // Condition for including this file (e.g., "IncludeHTMX", "IncludeAPI")
}

// StaticFile represents a file that should be copied as-is without processing
type StaticFile struct {
	SourcePath      string      // Path in templates filesystem
	DestinationPath string      // Path in generated project
	Permissions     fs.FileMode // File permissions
	Conditional     string      // Condition for including this file
}

// ProjectFile is a union type for either template or static files
type ProjectFile interface {
	GetSourcePath() string
	GetDestinationPath() string
	GetPermissions() fs.FileMode
	GetConditional() string
	IsTemplate() bool
}

// GetSourcePath Implement ProjectFile interface for TemplateFile
func (tf TemplateFile) GetSourcePath() string       { return tf.SourcePath }
func (tf TemplateFile) GetDestinationPath() string  { return tf.DestinationPath }
func (tf TemplateFile) GetPermissions() fs.FileMode { return tf.Permissions }
func (tf TemplateFile) GetConditional() string      { return tf.Conditional }
func (tf TemplateFile) IsTemplate() bool            { return true }

// GetSourcePath Implement ProjectFile interface for StaticFile
func (sf StaticFile) GetSourcePath() string       { return sf.SourcePath }
func (sf StaticFile) GetDestinationPath() string  { return sf.DestinationPath }
func (sf StaticFile) GetPermissions() fs.FileMode { return sf.Permissions }
func (sf StaticFile) GetConditional() string      { return sf.Conditional }
func (sf StaticFile) IsTemplate() bool            { return false }

// GetProjectFiles returns the list of all files to process (both templates and static)
func GetProjectFiles() []ProjectFile {
	return []ProjectFile{
		// Root template files
		TemplateFile{
			SourcePath:      "base/gitignore.gotmpl",
			DestinationPath: ".gitignore",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/gomod.gotmpl",
			DestinationPath: "go.mod",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/readme.gotmpl",
			DestinationPath: "README.md",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/justfile.gotmpl",
			DestinationPath: "Justfile",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/compose.gotmpl",
			DestinationPath: "compose.yaml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/compose.override.gotmpl",
			DestinationPath: "compose.override.yaml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/compose.production.gotmpl",
			DestinationPath: "compose.production.yaml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/air.gotmpl",
			DestinationPath: ".air.toml",
			Permissions:     0644,
		},

		// Config template files
		TemplateFile{
			SourcePath:      "base/env_template.gotmpl",
			DestinationPath: "config/env.template",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "base/env_dev.gotmpl",
			DestinationPath: "config/env.dev",
			Permissions:     0644,
		},

		// Main application templates
		TemplateFile{
			SourcePath:      "cmd/main.gotmpl",
			DestinationPath: "cmd/server/main.go",
			Permissions:     0644,
		},

		// App layer templates
		TemplateFile{
			SourcePath:      "internal/app/app.gotmpl",
			DestinationPath: "internal/app/app.go",
			Permissions:     0644,
		},

		// Domain layer templates
		TemplateFile{
			SourcePath:      "internal/domain/user/entity.gotmpl",
			DestinationPath: "internal/domain/user/entity.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/domain/user/repository.gotmpl",
			DestinationPath: "internal/domain/user/repository.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/domain/user/service.gotmpl",
			DestinationPath: "internal/domain/user/service.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/domain/auth/entity.gotmpl",
			DestinationPath: "internal/domain/auth/entity.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/domain/auth/service.gotmpl",
			DestinationPath: "internal/domain/auth/service.go",
			Permissions:     0644,
		},

		// Infrastructure layer templates
		TemplateFile{
			SourcePath:      "internal/infrastructure/config/config.gotmpl",
			DestinationPath: "internal/infrastructure/config/config.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/infrastructure/database/postgres.gotmpl",
			DestinationPath: "internal/infrastructure/database/postgres.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/infrastructure/web/server.gotmpl",
			DestinationPath: "internal/infrastructure/web/server.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/infrastructure/web/router.gotmpl",
			DestinationPath: "internal/infrastructure/web/router.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/infrastructure/web/middleware/auth.gotmpl",
			DestinationPath: "internal/infrastructure/web/middleware/auth.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/infrastructure/web/middleware/csrf.gotmpl",
			DestinationPath: "internal/infrastructure/web/middleware/csrf.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/infrastructure/web/middleware/logging.gotmpl",
			DestinationPath: "internal/infrastructure/web/middleware/logging.go",
			Permissions:     0644,
		},

		// Adapters layer templates
		TemplateFile{
			SourcePath:      "internal/adapters/repository/user_postgres.gotmpl",
			DestinationPath: "internal/adapters/repository/user_postgres.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/adapters/repository/session_postgres.gotmpl",
			DestinationPath: "internal/adapters/repository/session_postgres.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/adapters/handlers/web/handlers.gotmpl",
			DestinationPath: "internal/adapters/handlers/web/handlers.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/adapters/handlers/web/auth_handler.gotmpl",
			DestinationPath: "internal/adapters/handlers/web/auth_handler.go",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "internal/adapters/handlers/web/home_handler.gotmpl",
			DestinationPath: "internal/adapters/handlers/web/home_handler.go",
			Permissions:     0644,
		},

		// HTMX-specific template files
		TemplateFile{
			SourcePath:      "internal/adapters/handlers/web/htmx_handler.gotmpl",
			DestinationPath: "internal/adapters/handlers/web/htmx_handler.go",
			Permissions:     0644,
			Conditional:     "IncludeHTMX",
		},
		TemplateFile{
			SourcePath:      "web-templates/partials/user_info.gotmpl",
			DestinationPath: "internal/infrastructure/web/templates/partials/user_info.gohtml",
			Permissions:     0644,
			Conditional:     "IncludeHTMX",
		},

		// API-specific template files
		TemplateFile{
			SourcePath:      "internal/adapters/handlers/api/handlers.gotmpl",
			DestinationPath: "internal/adapters/handlers/api/handlers.go",
			Permissions:     0644,
			Conditional:     "IncludeAPI",
		},
		TemplateFile{
			SourcePath:      "internal/adapters/handlers/api/user_handler.gotmpl",
			DestinationPath: "internal/adapters/handlers/api/user_handler.go",
			Permissions:     0644,
			Conditional:     "IncludeAPI",
		},

		// Web HTML templates
		TemplateFile{
			SourcePath:      "web-templates/base.gotmpl",
			DestinationPath: "internal/infrastructure/web/templates/base.gohtml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "web-templates/home.gotmpl",
			DestinationPath: "internal/infrastructure/web/templates/home.gohtml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "web-templates/login.gotmpl",
			DestinationPath: "internal/infrastructure/web/templates/login.gohtml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "web-templates/register.gotmpl",
			DestinationPath: "internal/infrastructure/web/templates/register.gohtml",
			Permissions:     0644,
		},
		TemplateFile{
			SourcePath:      "web-templates/dashboard.gotmpl",
			DestinationPath: "internal/infrastructure/web/templates/dashboard.gohtml",
			Permissions:     0644,
		},

		// JavaScript templates (need project name)
		TemplateFile{
			SourcePath:      "static/js/app.js.gotmpl",
			DestinationPath: "static/js/app.js",
			Permissions:     0644,
		},

		// Static files (copied as-is)
		StaticFile{
			SourcePath:      "static/app.css",
			DestinationPath: "static/css/app.css",
			Permissions:     0644,
		},
		StaticFile{
			SourcePath:      "static/favicon.ico",
			DestinationPath: "static/favicon.ico",
			Permissions:     0644,
		},
		StaticFile{
			SourcePath:      "static/robots.txt",
			DestinationPath: "static/robots.txt",
			Permissions:     0644,
		},
		StaticFile{
			SourcePath:      "internal/infrastructure/database/migrations/001_users.sql",
			DestinationPath: "internal/infrastructure/database/migrations/001_create_users_table.sql",
			Permissions:     0644,
		},
		StaticFile{
			SourcePath:      "internal/infrastructure/database/migrations/002_sessions.sql",
			DestinationPath: "internal/infrastructure/database/migrations/002_create_sessions_table.sql",
			Permissions:     0644,
		},
	}
}
