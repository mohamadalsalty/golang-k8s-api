package handlers

import (
	"context"
	"fmt"
	"k8s/helpers"
	"k8s/pkg/config"
	"net/http"

	"k8s/models"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NamespacesHandler(w http.ResponseWriter, r *http.Request) {
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

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing namespaces: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var response models.NamespaceListResponse
	for _, ns := range namespaces.Items {
		namespace := models.Namespace{
			Name: ns.Name,
		}
		response.Namespaces = append(response.Namespaces, namespace)
	}

	helpers.JSONResponse(w, http.StatusOK, response)

}
