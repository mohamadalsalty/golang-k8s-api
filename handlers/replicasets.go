package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"k8s/helpers"
	"k8s/models"
	"k8s/pkg/config"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ReplicaSetsHandler(w http.ResponseWriter, r *http.Request) {
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

	parts := strings.Split(r.URL.Path, "/")
	var namespace string
	if len(parts) >= 3 {
		namespace = parts[2]
	}

	replicasets, err := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing ReplicaSets: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var response models.ReplicaSetListResponse
	for _, rs := range replicasets.Items {
		replicaSet := models.ReplicaSet{
			Name: rs.Name,
		}
		response.ReplicaSets = append(response.ReplicaSets, replicaSet)
	}
	helpers.JSONResponse(w, http.StatusOK, response)

}
