package dto

type DroneWebHookEntity struct {
	Event              string `json:"event,omitempty"`
	Owner              string
	A                  string
	B                  string
	Link               string
	CommitAuthor       string
	CommitAuthorEmail  string
	CommitAuthorName   string
	CommitAuthorAvatar string
	CommitMessage      string
	FailStages         string
	FailStep           string
	Target             string
	Repo               string
	Status             string
	Message            string
}
