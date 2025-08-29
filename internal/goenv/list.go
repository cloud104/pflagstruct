package goenv

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// GoListDir runs "go list -f '{{.Dir}}' <pkg>" and returns the source path.
func GoListDir(pkg string) (string, error) {
	cmd := exec.Command("go", "list", "-f", "{{.Dir}}", pkg)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", errors.Errorf("failed to run go list: %v, output: %s", err, out.String())
	}

	return strings.TrimSpace(out.String()), nil
}
