package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		args       []string
		wantErr    bool
		wantOutput string
	}{
		{
			name:       "version flag prints version",
			args:       []string{"switchboard", "--version"},
			wantOutput: "dev",
		},
		{
			name:       "no args prints version",
			args:       []string{"switchboard"},
			wantOutput: "dev",
		},
		{
			name:    "unknown flag returns error",
			args:    []string{"switchboard", "--bogus"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			err := run(&buf, tt.args)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := buf.String()
			if !strings.Contains(got, tt.wantOutput) {
				t.Errorf("output %q does not contain %q", got, tt.wantOutput)
			}
		})
	}
}

func TestVersionNonEmpty(t *testing.T) {
	t.Parallel()

	if version == "" {
		t.Fatal("version must not be empty")
	}
}

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func TestRun_WriteError(t *testing.T) {
	t.Parallel()

	err := run(failWriter{}, []string{"switchboard", "--version"})
	if err == nil {
		t.Fatal("expected error from failing writer")
	}
}
