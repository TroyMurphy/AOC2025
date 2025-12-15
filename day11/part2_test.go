package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart2_Sample(t *testing.T) {
	b, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("read sample.txt: %v", err)
	}
	content := strings.TrimSpace(string(b))
	lines := strings.Split(content, "\n")

	got := Part2(lines)
	const want = 2
	if got != want {
		t.Fatalf("Part2(sample) = %d, want %d", got, want)
	}
}
