package handlers

import (
	"fmt"
	"k8s/helpers"
	"k8s/models"
	"k8s/pkg/config"
	"net/http"

	"k8s.io/client-go/kubernetes"
)

type ClusterInfoResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func ClusterInfoHandler(w http.ResponseWriter, r *http.Request) {
	config, err := config.GetKubernetesConfig()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating client config: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating clientset: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	info, err := clientset.Discovery().ServerVersion()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting cluster info: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := models.ClusterInfoResponse{
		Name:    info.String(),
		Version: info.String(),
	}

	helpers.JSONResponse(w, http.StatusOK, response)
}
