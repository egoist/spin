package main

import (
	"fmt"

	"github.com/egoist/spin/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	repo := kingpin.Arg("repo", "Repository slug").Required().String()
	target := kingpin.Arg("out-dir", "Output directory").Default(".").String()

	kingpin.Parse()

	if utils.FileExists(*target) {
		isEmpty, _ := utils.IsEmpty(*target)
		if !isEmpty {
			fmt.Println("Output directory already exists and is not empty")
			return
		}
	}

	url := fmt.Sprintf("https://github.com/%s/archive/master/archive.zip", *repo)

	outFile, err := utils.DownloadFile(url)

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("Creating new project in %s\n", *target)
		utils.Unzip(outFile.Name(), *target)
		fmt.Println("Success!")
	}
}
