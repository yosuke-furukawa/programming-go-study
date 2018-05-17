package gh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Owner string
	Repo  string
	Token string
}

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state,omitempty"`
}

type IssueResult struct {
	Id     json.Number
	Number json.Number
	Url    string `json:"html_url"`
	Title  string
	Body   string
	State  string
}

const GithubIssueURL = "https://api.github.com/repos/%s/%s/issues"

func (repository Repository) Create(issue Issue) (*IssueResult, error) {
	url := fmt.Sprintf(GithubIssueURL, repository.Owner, repository.Repo)
	data, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", url, buf)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+repository.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result IssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository Repository) Read(id int) (*IssueResult, error) {
	url := fmt.Sprintf(GithubIssueURL+"/%d", repository.Owner, repository.Repo, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+repository.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result IssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository Repository) Edit(id int, issue Issue) (*IssueResult, error) {
	url := fmt.Sprintf(GithubIssueURL+"/%d", repository.Owner, repository.Repo, id)
	data, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)

	req, err := http.NewRequest("PATCH", url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+repository.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result IssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository Repository) Close(id int) (*IssueResult, error) {
	iss, err := repository.Read(id)
	if err != nil {
		return nil, err
	}
	issue := Issue{
		iss.Title,
		iss.Body,
		"closed",
	}
	return repository.Edit(id, issue)
}
