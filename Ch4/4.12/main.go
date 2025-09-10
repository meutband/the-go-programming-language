// This file downloads comic book informattion and searches it's transcript for specific terms

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"sate_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	ImgLink    string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const baseURL = "http://xkcd.com"
const infoLink = "info.0.json"
const comicsDir = "comics"

const usage = `xkcd get <idx>
xkcd search <term>`

func getComic(idx string) (comic, error) {

	url := fmt.Sprintf("%s/%s/%s", baseURL, idx, infoLink)
	resp, err := http.Get(url)
	if err != nil {
		return comic{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return comic{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	blob, err := io.ReadAll(resp.Body)
	if err != nil {
		return comic{}, err
	}

	var c comic
	err = json.Unmarshal(blob, &c)
	return c, err
}

func saveComic(idx string, com comic) error {

	f := fmt.Sprintf("comics/%s.json", idx)

	file, err := os.Create(f)
	if err != nil {
		return err
	}

	blob, err := json.Marshal(com)
	if err != nil {
		return err
	}

	_, err = file.Write(blob)
	return err
}

func getFiles() ([]string, error) {

	var files []string
	err := filepath.WalkDir(comicsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.EqualFold(filepath.Ext(path), ".json") {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func searchFiles(f, term string) (comic, bool, error) {

	file, err := os.Open(f)
	if err != nil {
		return comic{}, false, err
	}

	blob, err := io.ReadAll(file)
	if err != nil {
		return comic{}, false, err
	}

	var c comic
	if strings.Contains(strings.ToLower(string(blob)), strings.ToLower(term)) {
		err = json.Unmarshal(blob, &c)
		return c, true, err
	}

	return comic{}, false, err
}

func searchComics(term string) ([]comic, error) {

	files, err := getFiles()
	if err != nil {
		return nil, err
	}

	var comics []comic

	for _, f := range files {

		com, found, err := searchFiles(f, term)
		if err != nil {
			fmt.Printf("err searching file %s: %v", f, err)
			continue
		}

		if found {
			comics = append(comics, com)
		}
	}

	return comics, nil
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Incorrect number of arguments. Usage:")
		fmt.Println(usage)
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "get":
		idx := os.Args[2]
		com, err := getComic(idx)
		if err != nil {
			panic(err)
		}
		err = saveComic(idx, com)
		if err != nil {
			fmt.Println("error saving comic", err)
		}
	case "search":
		term := os.Args[2]
		comics, err := searchComics(term)
		if err != nil {
			fmt.Println("error searching file", err)
			os.Exit(0)
		}
		for _, c := range comics {
			fmt.Println("Title:", c.Title)
			fmt.Printf("URL: %s/%v/%s\n", baseURL, c.Num, infoLink)
			fmt.Println("Transcript", c.Transcript)
			fmt.Println()
		}
	default:
		fmt.Println("wrong command. Usage:")
		fmt.Println(usage)
	}

}
