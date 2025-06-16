package git

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func ListBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "-a")
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
func ReadReadmesFromBranches(branches []string) (map[string]string, error) {
	readmes := make(map[string]string)
	currentBranch, err := getCurrentBranch()
	if err != nil {
		return nil, err
	}

	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if branch == "" || branch == currentBranch {
			continue
		}
		// Cambiar a la rama
		cmd := exec.Command("git", "checkout", branch)
		if err := cmd.Run(); err != nil {
			readmes[branch] = "No se pudo cambiar a la rama"
			continue
		}
		// Leer README.md
		content, err := ioutil.ReadFile("README.md")
		if err != nil {
			readmes[branch] = "README.md no encontrado"
		} else {
			readmes[branch] = string(content)
		}
	}

	// Volver a la rama original
	exec.Command("git", "checkout", currentBranch).Run()
	return readmes, nil
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func CheckoutAllRemoteBranches() error {
	cmd := exec.Command("git", "branch", "-r")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	remotes := strings.Split(out.String(), "\n")
	localBranches, _ := ListBranches()
	localSet := make(map[string]struct{})
	for _, b := range localBranches {
		localSet[strings.TrimSpace(b)] = struct{}{}
	}
	for _, remote := range remotes {
		remote = strings.TrimSpace(remote)
		if remote == "" || strings.Contains(remote, "->") || strings.HasSuffix(remote, "/HEAD") {
			continue
		}
		parts := strings.Split(remote, "/")
		if len(parts) < 2 {
			continue
		}
		branch := parts[len(parts)-1]
		if branch == "main" {
			continue
		}
		if _, exists := localSet[branch]; exists {
			continue // Ya existe localmente
		}
		// Crear rama local desde la remota
		cmd := exec.Command("git", "checkout", "-b", branch, remote)
		if err := cmd.Run(); err != nil {
			fmt.Printf("No se pudo crear la rama local %s desde %s: %v\n", branch, remote, err)
		}
	}
	// Volver a main
	exec.Command("git", "checkout", "main").Run()
	return nil
}
func DeleteAllLocalBranchesExceptMain() error {
	branches, err := ListBranches()
	if err != nil {
		return err
	}
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if branch == "" || branch == "main" {
			continue
		}
		exec.Command("git", "branch", "-D", branch).Run()
	}
	return nil
}
