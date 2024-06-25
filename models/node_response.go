package models

type Node struct {
	Name string `json:"name"`
}

type NodeListResponse struct {
	Nodes []Node `json:"nodes"`
}
