package handlers

import (
	"context"
	"fmt"
	"net/http"

	"k8s/helpers"
	"k8s/models"
	"k8s/pkg/config"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NodesHandler(w http.ResponseWriter, r *http.Request) {
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

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing nodes: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var response models.NodeListResponse
	for _, node := range nodes.Items {
		nodeItem := models.Node{
			Name: node.Name,
		}
		response.Nodes = append(response.Nodes, nodeItem)
	}

	helpers.JSONResponse(w, http.StatusOK, response)

}
