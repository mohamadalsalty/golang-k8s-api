package handlers

import (
	"context"
	"fmt"
	"k8s/pkg/config"
	"net/http"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// Extract namespace from URL path, or set to empty string if not provided
	parts := strings.Split(r.URL.Path, "/")
	var namespace string
	if len(parts) >= 3 {
		namespace = parts[2]
	}

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing Deployments: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	for _, deployment := range deployments.Items {
		fmt.Fprintf(w, "Deployment: %s\n", deployment.Name)
	}
}
