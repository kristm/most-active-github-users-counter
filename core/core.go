package core

type RepoResponse struct {
  Repo string `json:"name"`
  Watchers int `json:"stargazers_count"`
  Fork bool `json:"fork"`
}
