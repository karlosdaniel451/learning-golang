package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type GitHubRepository struct {
	Id               int
	Name             string
	Description      string
	Fork             bool
	StarGazersCount  int `json:"stargazers_count"`
	WatchersCount    int `json:"watchers_count"`
	Language         string
	ForksCount       int
	Topics           []string
	Forks            int
	OpenIssues       int `json:"open_issues"`
	SubscribersCount int `json:"subscribers_count"`
}

var RepositoryRequestUrl = flag.String("url", "", "Repository Request URL")

func main() {
	flag.Parse()

	if *RepositoryRequestUrl == "" {
		fmt.Fprintf(os.Stderr, "error: undefined repository url")
		flag.Usage()
		os.Exit(1)
	}

	repository, err := requestGitHubRepository(*RepositoryRequestUrl)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	repositoryString, err := json.MarshalIndent(*repository, "", "  ")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("repository data:\n%s\n", repositoryString)
}

func requestGitHubRepository(repositoryURL string) (*GitHubRepository, error) {
	response, err := http.Get(repositoryURL)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("request failed with status %d", response.StatusCode)
	}

	var repository GitHubRepository

	err = json.NewDecoder(response.Body).Decode(&repository)
	if err != nil {
		response.Body.Close()
		return nil, err
	}

	response.Body.Close()
	return &repository, nil
}
