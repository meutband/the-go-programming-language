// Issues prints a table of Github issues matching the search terms
package main

import (
	"fmt"
	"gobook/Ch4/4.10/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %55.55s %.25s\n",
			item.Number, item.User.Login, item.Title, getLabel(item.CreatedAt))
	}
}

func getLabel(prev time.Time) string {

	var label string
	curr := time.Now()
	monthAgo := curr.AddDate(0, -1, 0)
	yearAgo := curr.AddDate(-1, 0, 0)
	if prev.After(monthAgo) {
		label = "less than month old"
	} else if prev.Before(yearAgo) {
		label = "more than a year old"
	} else {
		label = "less than a year old"
	}
	return label
}
