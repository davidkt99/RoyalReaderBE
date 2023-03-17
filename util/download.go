package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Download(url string) []byte {
	resp, loadErr := http.Get(url)
	if loadErr != nil {
		fmt.Println("Error: ", loadErr)
		os.Exit(1)
	}
	// time.Sleep(1000 * time.Millisecond)

	// io.Copy(os.Stdout, resp.Body)

	dataInBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Error: ", readErr)
		os.Exit(1)
	}

	return dataInBytes
}
