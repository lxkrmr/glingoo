package cmd

import (
	"testing"
)

func TestPoFilename(t *testing.T) {
	tests := []struct {
		lang     string
		expected string
	}{
		{"de_DE", "de.po"},
		{"it_IT", "it.po"},
		{"fr_FR", "fr.po"},
		{"de_CH", "de.po"},
	}

	for _, tt := range tests {
		t.Run(tt.lang, func(t *testing.T) {
			got := poFilename(tt.lang)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestParseExportArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantAddon string
		wantLang  string
		wantDir   string
		wantErr   bool
	}{
		{
			name:      "valid",
			args:      []string{"my_addon", "de_DE", "/path/to/i18n"},
			wantAddon: "my_addon",
			wantLang:  "de_DE",
			wantDir:   "/path/to/i18n",
		},
		{
			name:    "missing output dir",
			args:    []string{"my_addon", "de_DE"},
			wantErr: true,
		},
		{
			name:    "missing all",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"my_addon", "de_DE", "/path", "extra"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseExportArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.addon != tt.wantAddon {
				t.Errorf("addon: expected %q, got %q", tt.wantAddon, got.addon)
			}
			if got.lang != tt.wantLang {
				t.Errorf("lang: expected %q, got %q", tt.wantLang, got.lang)
			}
			if got.outputDir != tt.wantDir {
				t.Errorf("outputDir: expected %q, got %q", tt.wantDir, got.outputDir)
			}
		})
	}
}

func TestBuildExportResult(t *testing.T) {
	input := exportInput{addon: "my_addon", lang: "de_DE", outputDir: "/path/i18n"}
	result := buildExportResult(input, "/path/i18n/de.po", true)

	if result["addon"] != "my_addon" {
		t.Errorf("addon: expected %q, got %v", "my_addon", result["addon"])
	}
	if result["lang"] != "de_DE" {
		t.Errorf("lang: expected %q, got %v", "de_DE", result["lang"])
	}
	if result["path"] != "/path/i18n/de.po" {
		t.Errorf("path: expected %q, got %v", "/path/i18n/de.po", result["path"])
	}
	if result["created"] != true {
		t.Errorf("created: expected true, got %v", result["created"])
	}
}
