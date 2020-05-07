package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GitClone clones a repo with depth 1
func GitClone(repo string, outDir string) (err error) {
	os.MkdirAll(filepath.Dir(outDir), os.ModePerm)
	cmd := exec.Command("git", "clone", "git@github.com:"+repo+".git", outDir, "--depth=1")
	fmt.Printf("Running git clone and waiting for it to finish...\n")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
		return
	}
	err = os.RemoveAll(filepath.Join(outDir, ".git"))
	return
}
