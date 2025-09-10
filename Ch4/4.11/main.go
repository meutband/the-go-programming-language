package main

import (
	"bufio"
	"fmt"
	"gobook/Ch4/4.11/github"
	"os"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("incorrect usage\n./main <create/update/close/get> <owner> <repo> <issue_num>")
		os.Exit(1)
	}

	// ./main <create/edit/close/get> <owner> <repo> <issue_num>
	mode := os.Args[1]
	owner := os.Args[2]
	repo := os.Args[3]
	issue_num := os.Args[4]

	switch mode {
	case "get":

		issue, err := github.GetIssue(owner, repo, issue_num)
		if err != nil {
			fmt.Println("err getting issue", err)
			os.Exit(1)
		}
		fmt.Printf("Issue: %+v\n", issue)

	case "close":

		err := github.CloseIssue(owner, repo, issue_num)
		if err != nil {
			fmt.Println("err closing issue", err)
			os.Exit(1)
		}
		fmt.Println("issue closed")

	case "create":

		fmt.Print("Please enter a title <Press Enter, then <CTRL+D>: ")
		title, err := scan()
		if err != nil {
			fmt.Println("err creating issue title", err)
			os.Exit(1)
		}

		fmt.Print("Please enter some text for body <Press Enter, then <CTRL+D>: ")
		body, err := scan()
		if err != nil {
			fmt.Println("err creating issue body", err)
			os.Exit(1)
		}

		err = github.CreateIssue(owner, repo, title, body)
		if err != nil {
			fmt.Println("err creating issue", err)
			os.Exit(1)
		}
		fmt.Println("issue created")

	case "update":

		fmt.Print("Please enter some text for the body <Press Enter, then <CTRL+D>: ")
		body, err := scan()
		if err != nil {
			fmt.Println("err creating issue text", err)
			os.Exit(1)
		}

		err = github.UpdateIssue(owner, repo, issue_num, body)
		if err != nil {
			fmt.Println("err updating issue", err)
			os.Exit(1)
		}
		fmt.Println("issue updated")

	default:
		fmt.Println("incorrect mode (<create/update/close/get>)")
		os.Exit(1)
	}

}

func scan() (string, error) {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return text, nil
}
