package main

import (
	"flag"
	"fmt"
	"log"

	jira "github.com/andygrunwald/go-jira"
)

func runCLI(username, apiToken string, months int, instanceURL string) error {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: apiToken,
	}

	client, err := jira.NewClient(tp.Client(), instanceURL)
	if err != nil {
		return err
	}

	jql := fmt.Sprintf("assignee = currentUser() AND created >= -%dM", months)

	options := &jira.SearchOptions{
		Fields:     []string{"summary", "description"},
		MaxResults: 100,
	}

	issues, _, err := client.Issue.Search(jql, options)
	if err != nil {
		return err
	}

	for _, issue := range issues {
		fmt.Printf("Title: %s\n", issue.Fields.Summary)
		fmt.Printf("Text: %s\n", issue.Fields.Description)
		fmt.Println("--------------------------")
	}

	return nil
}

func main() {
	var (
		username    string
		apiToken    string
		months      int
		instanceURL string
	)

	flag.StringVar(&username, "username", "", "JIRA username")
	flag.StringVar(&apiToken, "apitoken", "", "JIRA API token")
	flag.IntVar(&months, "months", 6, "Number of months to search")
	flag.StringVar(&instanceURL, "url", "https://your-jira-instance-url.com", "JIRA instance URL")
	flag.Parse()

	err := runCLI(username, apiToken, months, instanceURL)
	if err != nil {
		log.Fatal(err)
	}
}
