package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/andygrunwald/go-jira"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Search(jql string, options *jira.SearchOptions) ([]jira.Issue, *jira.Response, error) {
	args := m.Called(jql, options)
	return args.Get(0).([]jira.Issue), args.Get(1).(*jira.Response), args.Error(2)
}

func TestMainFunction(t *testing.T) {
	mockClient := new(MockClient)

	mockIssues := []jira.Issue{
		{
			Key: "JIRA-1",
			Fields: &jira.IssueFields{
				Summary:     "Card 1",
				Description: "Description of Card 1",
			},
		},
		{
			Key: "JIRA-2",
			Fields: &jira.IssueFields{
				Summary:     "Card 2",
				Description: "Description of Card 2",
			},
		},
	}

	mockResponse := &jira.Response{
		Total: len(mockIssues),
	}

	mockClient.On("Search", mock.Anything, mock.Anything).Return(mockIssues, mockResponse, nil)

	output := bytes.NewBuffer(nil)
	log.SetOutput(output)

	// Update the below values with your desired test values
	testUsername := "testuser"
	testAPIKey := "testapitoken"
	testMonths := 6
	testURL := "https://your-jira-instance-url.com"

	err := runCLI(testUsername, testAPIKey, testMonths, testURL)
	if err != nil {
		t.Errorf("Failed to run CLI: %v", err)
	}

	expectedOutput := fmt.Sprintf(`Title: Card 1
Text: Description of Card 1
--------------------------
Title: Card 2
Text: Description of Card 2
--------------------------
`)

	if output.String() != expectedOutput {
		t.Errorf("Unexpected output. Expected:\n%s\nGot:\n%s", expectedOutput, output.String())
	}

	mockClient.AssertExpectations(t)
}
