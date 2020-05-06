package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/egoist/spin/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	repo := kingpin.Arg("repo", "Repository slug").Required().String()
	target := kingpin.Arg("out-dir", "Output directory").Default(".").String()

	kingpin.Parse()

	if utils.PathExists(*target) {
		isEmpty, _ := utils.IsEmpty(*target)
		if !isEmpty {
			fmt.Println("Output directory already exists and is not empty")
			return
		}
	}

	homedir, _ := os.UserHomeDir()
	cacheDir := filepath.Join(homedir, fmt.Sprintf(".spin/templates/github/%s", *repo))

	if !utils.PathExists(cacheDir) {
		url := fmt.Sprintf("https://github.com/%s/archive/master/archive.zip", *repo)

		zipFile, err := utils.DownloadFile(url)

		if err != nil {
			fmt.Println("error", err)
			return
		}

		utils.Unzip(zipFile.Name(), cacheDir)
	}

	if err := utils.Copy(cacheDir, *target); err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Printf("Creating new project in %s\n", *target)
	fmt.Println("Success!")
}
