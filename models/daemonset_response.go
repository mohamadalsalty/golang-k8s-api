package models

type DaemonSet struct {
	Name string `json:"name"`
}

type DaemonSetListResponse struct {
	DaemonSets []DaemonSet `json:"daemonSets"`
}
