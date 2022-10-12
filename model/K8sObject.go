package model

type Namespaces struct {
	Items []NameSpaceItem `json:"items"`
}

type NameSpaceItem struct {
	Metadata Namesmetadata   `json:"metadata"`
	Status   NameSpaceStatus `json:"status"`
}
type Namesmetadata struct {
	Name string `json:"name"`
}

type NameSpaceStatus struct {
	Phase string `json:"phase"`
}
