package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/senforsce/tndr"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		expected     string
		expectedCode int
	}{
		{
			name:         "no args prints usage",
			args:         []string{},
			expected:     usageText,
			expectedCode: 0,
		},
		{
			name:         `"t1 help" prints help`,
			args:         []string{"t1", "help"},
			expected:     usageText,
			expectedCode: 0,
		},
		{
			name:         `"t1 --help" prints help`,
			args:         []string{"t1", "--help"},
			expected:     usageText,
			expectedCode: 0,
		},
		{
			name:         `"t1 version" prints version`,
			args:         []string{"t1", "version"},
			expected:     t1.Version() + "\n",
			expectedCode: 0,
		},
		{
			name:         `"t1 --version" prints version`,
			args:         []string{"t1", "--version"},
			expected:     t1.Version() + "\n",
			expectedCode: 0,
		},
		{
			name:         `"t1 fmt --help" prints usage`,
			args:         []string{"t1", "fmt", "--help"},
			expected:     fmtUsageText,
			expectedCode: 0,
		},
		{
			name:         `"t1 generate --help" prints usage`,
			args:         []string{"t1", "generate", "--help"},
			expected:     generateUsageText,
			expectedCode: 0,
		},
		{
			name:         `"t1 lsp --help" prints usage`,
			args:         []string{"t1", "lsp", "--help"},
			expected:     lspUsageText,
			expectedCode: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := bytes.NewBuffer(nil)
			actualCode := run(actual, test.args)

			if actualCode != test.expectedCode {
				t.Errorf("expected code %v, got %v", test.expectedCode, actualCode)
			}
			if diff := cmp.Diff(test.expected, actual.String()); diff != "" {
				t.Error(diff)
				t.Error("expected:")
				t.Error(test.expected)
				t.Error("actual:")
				t.Error(actual.String())
			}
		})
	}
}
