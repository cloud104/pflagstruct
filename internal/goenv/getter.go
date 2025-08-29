package goenv

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Get retrieves the value of a Go environment variable
func Get(key string) (string, error) {
	out, err := exec.Command("go", "env", key).Output()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get Go env variable %q", key)
	}

	value := strings.TrimSpace(string(out))
	if value == "" {
		return "", errors.Errorf("Go env variable %q is empty", key)
	}

	return value, nil
}
