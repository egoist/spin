package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(url string) (f *os.File, err error) {
	// Create the file
	f, _ = ioutil.TempFile("", "spin-repo.*.zip")

	// Get the data
	var resp *http.Response
	resp, err = http.Get(url)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var bodyBytes []byte
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = errors.New(string(bodyBytes))
		return
	}

	if err != nil {
		return
	}
	defer f.Close()

	// Write the body to file
	_, err = io.Copy(f, resp.Body)
	return
}
