package models

type Deployment struct {
	Name              string `json:"name"`
	Replicas          int32  `json:"replicas"`
	Image             string `json:"image"`
	AvailableReplicas int32  `json:"availableReplicas"`
}

type DeploymentListResponse struct {
	Deployments []Deployment `json:"deployments"`
}
