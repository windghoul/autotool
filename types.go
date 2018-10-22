package main

//easyjson:json
type Headcommit struct {
	ID        string   `json:"id"`
	Timestamp string   `json:"timestamp"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

//easyjson:json
type GitJSON struct {
	Ref        string     `json:"ref"`
	Headcommit Headcommit `json:"head_commit"`
}
