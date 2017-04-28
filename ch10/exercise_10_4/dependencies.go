package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
)

type Package struct {
	Name         string   `json:"Name"`
	ImportPath   string   `json:"ImportPath"`
	Dependencies []string `json:"Deps"`
}

// getWorkspacePackages returns a slice of packages in the specified directory.
func getPackages(dir string) ([]string, error) {
	// get list of packages in the workspace and save output in buf
	d := filepath.Clean(dir) + string(filepath.Separator) + "..."
	cmd := exec.Command("go", "list", d)
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error getting packages in %q: %v", dir, err)
	}

	// read packages in the workspace line by line
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)
	packages := []string{}
	for scanner.Scan() {
		packages = append(packages, scanner.Text())
	}
	return packages, nil
}

// getPackageDependencies return a slice of pkgDeps which contain packages
// and their dependencies.
func getAllPackageDependencies(packages []string) ([]Package, error) {
	// get list of dependencies for package and save output in buf
	cmdArgs := append([]string{"list", "-json"}, packages...)
	//cmd := exec.Command("go", "list", "-json", strings.Join(packages, " "))
	cmd := exec.Command("go", cmdArgs...)
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	var pd []Package
	decoder := json.NewDecoder(buf)
	for decoder.More() {
		var pkg Package
		err := decoder.Decode(&pkg)
		if err != nil {
			return nil, err
		}
		pd = append(pd, pkg)
	}
	return pd, nil
}
