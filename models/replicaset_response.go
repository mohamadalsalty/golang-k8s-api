package models

type ReplicaSet struct {
	Name string `json:"name"`
}

type ReplicaSetListResponse struct {
	ReplicaSets []ReplicaSet `json:"replicaSets"`
}
