package models

type Deployment struct {
	Name string `json:"name"`
}

type DeploymentListResponse struct {
	Deployments []Deployment `json:"deployments"`
}
