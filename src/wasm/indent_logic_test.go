package main

import (
	"strings"
	"testing"
)

func TestIndentMakeScript(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "simple if block",
			input: "{{if(40.value; \"a\"; \"b\")}}",
			expected: `{{
  if(
    40.value;
    "a";
    "b"
  )
}}`,
		},
		{
			name:  "Switch statement with multiple cases",
			input: "{{switch(1; 2; 3)}}",
			expected: `{{
  switch(
    1;
    2;
    3
  )
}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IndentMakeScript(tt.input)
			if result != tt.expected {
				t.Errorf("IndentMakeScript() = %q, want %q", result, tt.expected)
				t.Logf("Input: %q", tt.input)
				t.Logf("Expected lines:")
				for i, line := range strings.Split(tt.expected, "\n") {
					t.Logf("  %d: %q", i, line)
				}
				t.Logf("Actual lines:")
				for i, line := range strings.Split(result, "\n") {
					t.Logf("  %d: %q", i, line)
				}
			}
		})
	}
}
