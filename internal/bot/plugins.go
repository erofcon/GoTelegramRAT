package bot

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func info() (string, error) {

	response := ""

	client := http.Client{Timeout: 6 * time.Second}

	res, err := client.Get("https://geolocation-db.com/json")

	if err != nil {
		fmt.Println(err)
		return response, err
	}

	for {
		buff := make([]byte, 1024)

		n, err := res.Body.Read(buff)
		response += string(buff[:n])

		if n == 0 || err != nil {
			break
		}

	}

	return response, nil
}

func pwd() (string, error) {

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path, nil

}

func cd(text string) (string, error) {

	dir := strings.Split(text, " ")

	if len(dir) != 2 {
		return "", fmt.Errorf("wrong format")
	}

	err := os.Chdir(dir[1])

	if err != nil {
		return "", fmt.Errorf("error to change directory " + err.Error())
	}

	return fmt.Sprintf("directory changed to '%s'", dir[1]), nil

}

func ls() (string, error) {

	pathEntries := ""
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	if len(entries) == 0 {
		pathEntries = "directory is empty"
	} else {
		pathEntries += "directory files:\n"
		for _, e := range entries {

			pathEntries += "\n" + e.Name()
		}
	}

	return pathEntries, nil
}
