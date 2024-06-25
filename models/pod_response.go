package models

type Pod struct {
	Name string `json:"name"`
}

type PodListResponse struct {
	Pods []Pod `json:"pods"`
}
