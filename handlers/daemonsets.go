package handlers

import (
	"context"
	"fmt"
	"k8s/pkg/config"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DaemonSetsHandler(w http.ResponseWriter, r *http.Request) {
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
	namespace := r.URL.Path[len("/daemonsets/"):]

	daemonsets, err := clientset.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	for _, daemonset := range daemonsets.Items {
		fmt.Fprintf(w, "Deployment: %s\n", daemonset.Name)
	}
}
