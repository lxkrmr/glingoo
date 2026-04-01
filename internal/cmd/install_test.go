package cmd

import (
	"testing"
)

func TestParseInstallArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantLang string
		wantErr  bool
	}{
		{
			name:     "valid",
			args:     []string{"de_DE"},
			wantLang: "de_DE",
		},
		{
			name:    "missing lang",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"de_DE", "extra"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInstallArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.lang != tt.wantLang {
				t.Errorf("lang: expected %q, got %q", tt.wantLang, got.lang)
			}
		})
	}
}

func TestBuildInstallResult(t *testing.T) {
	result := buildInstallResult("de_DE")
	if result["lang"] != "de_DE" {
		t.Errorf("lang: expected %q, got %v", "de_DE", result["lang"])
	}
	if result["result"] != "installed" {
		t.Errorf("result: expected %q, got %v", "installed", result["result"])
	}
}
