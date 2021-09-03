package model

type GitlabProject struct {
	Id                  int64  `json:"id"`
	Description         string `json:"description"`
	Name                string `json:"name"`
	Name_with_namespace string `json:"name_with_namespace"`
	Path                string `json:"path"`
	Created_at          string `json:"created_at"`
	Ssh_url_to_repo     string `json:"ssh_url_to_repo"`
	Http_url_to_repo    string `json:"http_url_to_repo"`
	Web_url             string `json:"web_url"`
	Readme_url          string `json:"readme_url"`
	Avatar_url          string `json:"avatar_url"`
	Forks_count         int64  `json:"forks_count"`
	Star_count          int64  `json:"star_count"`
	Last_activity_at    string `json:"last_activity_at"`
	///
	Namespace GitlabNamespace `json:"namespace"`
}

type GitlabNamespace struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"Kind"`
	Full_path string `json:"Full_path"`
	Parent_id int64  `json:"parent_id"`
	Web_url   string `json:"web_url"`
}

type GitlabBranch struct {
	Name   string             `json:"name"`
	Commit GitlabBranchCommit `json:"commit"`
}
type GitlabBranchCommit struct {
	Title          string `json:"title"`
	Message        string `json:"Message"`
	Author_name    string `json:"author_name"`
	Committer_name string `json:"committer_name"`
}
