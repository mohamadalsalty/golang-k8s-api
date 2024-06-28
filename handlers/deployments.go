// handlers/deployments_handler.go

package handlers

import (
	"context"
	"fmt"
	"k8s/helpers"
	"k8s/pkg/config"
	"net/http"
	"strings"

	"k8s/models"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeploymentsHandler(w http.ResponseWriter, r *http.Request) {
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

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing Deployments: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var response models.DeploymentListResponse
	for _, deployment := range deployments.Items {
		dep := models.Deployment{
			Name:              deployment.Name,
			Replicas:          *deployment.Spec.Replicas,
			Image:             deployment.Spec.Template.Spec.Containers[0].Image,
			AvailableReplicas: deployment.Status.AvailableReplicas,
		}
		response.Deployments = append(response.Deployments, dep)
	}
	helpers.JSONResponse(w, http.StatusOK, response)
}
