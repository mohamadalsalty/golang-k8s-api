package models

type Namespace struct {
	Name string `json:"name"`
}

type NamespaceListResponse struct {
	Namespaces []Namespace `json:"namespaces"`
}
