package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// GITHUB API: https://api.github.com/repo/owner/repo/issues

var token = os.Getenv("GITHUB_TOKEN")

type Issue struct {
	Number     int       `json:"number"`
	Title      string    `json:"title"`
	State      string    `json:"state"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateedAt time.Time `json:"updated_at"`
	HTMLURL    string    `json:"html_url"`
}

type IssuePatch struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

type IssueCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateIssue(owner, repo, title, body string) error {

	var ic IssueCreate
	ic.Title = title
	ic.Body = body

	blob, err := json.Marshal(ic)
	if err != nil {
		return err
	}

	fmt.Println(string(blob))

	buf := bytes.NewBuffer(blob)
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	fmt.Println("req Body", req.Body)

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	rBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(rBody), err)

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("http status not created")
	}

	return nil
}

func UpdateIssue(owner, repo, issue_num, body string) error {

	iss, err := GetIssue(owner, repo, issue_num)
	if err != nil {
		return err
	}

	var ip IssuePatch
	ip.Title = iss.Title
	ip.Body = body
	ip.State = iss.State

	blob, err := json.Marshal(ip)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(blob)

	req, err := request(buf, "PATCH", owner, repo, issue_num)
	if err != nil {
		return err
	}

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status not ok")
	}

	return nil
}

func CloseIssue(owner, repo, issue_num string) error {

	iss, err := GetIssue(owner, repo, issue_num)
	if err != nil {
		return err
	}

	var ip IssuePatch
	ip.Title = iss.Title
	ip.Body = iss.Body
	ip.State = "closed"

	blob, err := json.Marshal(ip)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(blob)

	req, err := request(buf, "PATCH", owner, repo, issue_num)
	if err != nil {
		return err
	}

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status not ok")
	}

	return nil
}

func GetIssue(owner, repo, issue_num string) (*Issue, error) {

	req, err := request(nil, "GET", owner, repo, issue_num)
	if err != nil {
		return nil, err
	}

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status not ok")
	}

	blob, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issue Issue
	err = json.Unmarshal(blob, &issue)

	return &issue, err
}

func request(reader io.Reader, method, owner, repo, issue_num string) (*http.Request, error) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, issue_num)

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req, nil
}
