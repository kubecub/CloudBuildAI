package cmd

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/kubecub/CloudBuildAI/pkg/version"
)

func TestVersionCommandIncludesProjectVersion(t *testing.T) {
	originalShort, originalOutput, originalStdout := shortPrint, output, os.Stdout
	t.Cleanup(func() { shortPrint, output, os.Stdout = originalShort, originalOutput, originalStdout })
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	defer writer.Close()
	os.Stdout = writer
	command := NewVersionCmd()
	command.SetArgs([]string{"--output=json"})
	if err := command.Execute(); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	var parsed struct {
		KubecubVersion version.Info `json:"kubecubVersion"`
	}
	if err := json.Unmarshal(got, &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if parsed.KubecubVersion.GitVersion != version.Get().GitVersion {
		t.Fatalf("project version missing from actual command output: %s", got)
	}
}
