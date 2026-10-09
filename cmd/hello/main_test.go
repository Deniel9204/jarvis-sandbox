package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunDefault(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run(nil, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if got, want := stdout.String(), "Hello, world!\n"; got != want {
		t.Errorf("stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunWithName(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"separate value", []string{"-name", "Bob"}, "Hello, Bob!\n"},
		{"equals form", []string{"-name=Bob"}, "Hello, Bob!\n"},
		{"double dash", []string{"--name", "Bob"}, "Hello, Bob!\n"},
		{"spaces", []string{"-name", "Ada Lovelace"}, "Hello, Ada Lovelace!\n"},
		{"non-ASCII", []string{"-name", "Zoë"}, "Hello, Zoë!\n"},
		{"empty is literal", []string{"-name", ""}, "Hello, !\n"},
		{"extra positional args ignored", []string{"-name", "Bob", "extra"}, "Hello, Bob!\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			code := run(tt.args, &stdout, &stderr)

			if code != 0 {
				t.Fatalf("exit code = %d, want 0 (stderr: %q)", code, stderr.String())
			}
			if got := stdout.String(); got != tt.want {
				t.Errorf("stdout = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRunUnknownFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-nme", "Bob"}, &stdout, &stderr)

	if code != 2 {
		t.Errorf("exit code = %d, want 2", code)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), "flag provided but not defined: -nme") {
		t.Errorf("stderr = %q, want unknown-flag error", stderr.String())
	}
	if !strings.Contains(stderr.String(), "-name") {
		t.Errorf("stderr = %q, want usage mentioning -name", stderr.String())
	}
}

func TestRunHelp(t *testing.T) {
	for _, arg := range []string{"-h", "-help"} {
		t.Run(arg, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			code := run([]string{arg}, &stdout, &stderr)

			if code != 0 {
				t.Errorf("exit code = %d, want 0", code)
			}
			if stdout.Len() != 0 {
				t.Errorf("stdout = %q, want empty", stdout.String())
			}
			if !strings.Contains(stderr.String(), "-name") {
				t.Errorf("stderr = %q, want usage mentioning -name", stderr.String())
			}
		})
	}
}
