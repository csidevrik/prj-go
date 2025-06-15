package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func ListBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	branches := strings.Split(out.String(), "\n")
	for i, b := range branches {
		branches[i] = strings.TrimSpace(strings.TrimPrefix(b, "*"))
	}
	return branches, nil
}
