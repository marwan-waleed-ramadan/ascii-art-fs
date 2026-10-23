package main

import (
	"strings"
	"testing"
)

func TestLoadBanner(t *testing.T) {
	alphabet, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(alphabet) != 95 {
		t.Errorf("expected 95 characters, got %d", len(alphabet))
	}

	hLines, ok := alphabet['H']
	if !ok {
		t.Errorf("expected 'H' to be in alphabet")
	}
	if len(hLines) != 8 {
		t.Errorf("expected 'H' to have 8 lines, got %d", len(hLines))
	}
}

func TestRenderSingleChar(t *testing.T) {
	alphabet, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	out, err := renderASCII("A", alphabet)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

func TestRenderMultiLine(t *testing.T) {
	alphabet, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	out, err := renderASCII("A\nB", alphabet)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")

	expectedLines := 16
	if len(lines) != expectedLines {
		t.Errorf("expected %d lines, got %d", expectedLines, len(lines))
	}
}
