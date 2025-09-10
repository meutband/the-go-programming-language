// This file download movie posters

// NOTE: create and set up a API Key for OMDBAPI.COM
// https://www.omdbapi.com/apikey.aspx

// > export OMDB_API_KEY=<key from above>

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const baseURL string = "https://omdbapi.com"
const posterDir string = "posters"

type movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	Rating     string `json:"imdbRating"`
	Votes      string `json:"imdbVotes"`
	ID         string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

const usage string = `omdb <movie name...>`

func getMovie(name string) (movie, error) {

	apiKey := os.Getenv("OMDB_API_KEY")

	url := fmt.Sprintf("%s/?t=%s&apikey=%s", baseURL, url.QueryEscape(name), apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return movie{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return movie{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	blob, err := io.ReadAll(resp.Body)
	if err != nil {
		return movie{}, err
	}

	var m movie
	err = json.Unmarshal(blob, &m)
	return m, err
}

func downloadPoster(m movie) error {

	resp, err := http.Get(m.Poster)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	fName := fmt.Sprintf("%s/%s.jpg", posterDir, m.Title)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Incorrect number of arguments. Usage:")
		fmt.Println(usage)
		os.Exit(0)
	}

	name := os.Args[1:]
	movName := strings.Join(name, " ")

	mov, err := getMovie(movName)
	if err != nil {
		panic(err)
	}

	err = downloadPoster(mov)
	if err != nil {
		panic(err)
	}
}
