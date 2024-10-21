package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubActivity struct {
	Type      string     `json:"type"`
	Repo      Repository `json:"repo"`
	CreatedAt string     `json:"created_at"`
	Payload   Payload    `json:"payload"`
}

type Repository struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	Action  string    `json:"action"`
	Ref     string    `json:"ref"`
	RefType string    `json:"ref_type"`
	Commits []Commits `json:"commits"`
}

type Commits struct {
	Message string `json:"message"`
}

func FetchGithubActivity(username string) ([]GithubActivity, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("User not found")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error in fetching data from server")
	}

	var activities []GithubActivity

	if err := json.NewDecoder(response.Body).Decode(&activities); err != nil {
		return nil, err
	}

	return activities, nil
}

func ShowActivities(activities []GithubActivity) error {

	if len(activities) == 0 {
		return fmt.Errorf("No activities here")
	}

	for _, event := range activities {
		var action string
		repoName := event.Repo.Name

		switch event.Type {
		case "PushEvent":
			commitsCount := len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commit(s) to %s", commitsCount, repoName)
		case "IssuesEvent":
			fmt.Sprintf("%s an issue in %s", event.Payload.Action, repoName)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", repoName)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", repoName)
		case "CreateEvent":
			action = fmt.Sprintf("Created %s in %s", event.Payload.RefType, repoName)
		default:
			action = fmt.Sprintf("%s in %s", event.Type, repoName)
		}
		fmt.Println(fmt.Sprintf("- %s", action))
	}
	return nil
}
