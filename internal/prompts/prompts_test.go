package prompts

import (
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"valid name", "my-web-app", false},
		{"valid with numbers", "app123", false},
		{"valid with underscores", "my_app", false},
		{"empty name", "", true},
		{"starts with hyphen", "-invalid", true},
		{"starts with underscore", "_invalid", true},
		{"reserved name", "go", true},
		{"reserved name case insensitive", "GO", true},
		{"reserved name test", "test", true},
		{"contains spaces", "my app", true},
		{"contains special chars", "my@app", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProjectName(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf(
					"validateProjectName(%q) error = %v, wantError %v",
					tt.input,
					err,
					tt.wantError,
				)
			}
		})
	}
}

func TestValidateModulePath(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"valid github path", "github.com/user/repo", false},
		{"valid with subdirectory", "github.com/user/repo/subdir", false},
		{"valid simple name", "mymodule", false},
		{"valid with version", "github.com/user/repo/v2", false},
		{"empty path", "", true},
		{"starts with special char", "@invalid", true},
		{"ends with special char", "invalid-", true},
		{"contains spaces", "github.com/user/my repo", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModulePath(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf(
					"validateModulePath(%q) error = %v, wantError %v",
					tt.input,
					err,
					tt.wantError,
				)
			}
		})
	}
}

func TestBoolToYesNo(t *testing.T) {
	// Note: This function returns colored output, so we test the logical behavior
	result := boolToYesNo(true)
	if result == "" {
		t.Error("boolToYesNo(true) returned empty string")
	}

	result = boolToYesNo(false)
	if result == "" {
		t.Error("boolToYesNo(false) returned empty string")
	}

	// The actual output contains color codes, so we can't test exact string equality
	// but we can verify non-empty output
}
