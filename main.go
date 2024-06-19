package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Get environment variables
	tokenPath := os.Getenv("K8S_TOKEN_PATH")
	if tokenPath == "" {
		tokenPath = "secret.txt"
	}

	host := os.Getenv("K8S_API_HOST")
	if host == "" {
		host = "https://192.168.49.2:8443"
	}

	// Read the service account token
	token, err := os.ReadFile(tokenPath)
	if err != nil {
		fmt.Printf("Error reading token: %s\n", err.Error())
		return
	}

	// Create a new config using the token
	config := &rest.Config{
		Host:        host,
		BearerToken: string(token),
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true, // Disable SSL verification (use only for testing)
		},
	}
	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating clientset: %s\n", err.Error())
		os.Exit(1)
	}

	// HTTP handler function to get cluster information
	http.HandleFunc("/clusterinfo", func(w http.ResponseWriter, r *http.Request) {
		// Get cluster info
		info, err := clientset.Discovery().ServerVersion()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting cluster info: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// Print cluster info
		fmt.Fprintf(w, "Cluster Info:\n")
		fmt.Fprintf(w, "  Name: %s\n", info.String())
		fmt.Fprintf(w, "  Version: %s\n", info.String())
	})

	// HTTP handler function to list namespaces
	http.HandleFunc("/namespaces", func(w http.ResponseWriter, r *http.Request) {
		// List namespaces
		namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error listing namespaces: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// Print namespace names
		for _, ns := range namespaces.Items {
			fmt.Fprintf(w, "Namespace: %s\n", ns.Name)
		}
	})

	// HTTP handler function to list nodes
	http.HandleFunc("/nodes", func(w http.ResponseWriter, r *http.Request) {
		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error listing nodes: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		for _, node := range nodes.Items {
			fmt.Fprintf(w, "Node: %s\n", node.Name)
		}
	})

	// HTTP handler function to list pods
	http.HandleFunc("/pods/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the namespace from the request URL
		namespace := r.URL.Path[len("/pods/"):]
		if namespace == "" {
			namespace = "default" // Default to "default" namespace if no namespace is provided
		}

		// List pods in the specified namespace
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error listing pods in namespace %s: %s", namespace, err.Error()), http.StatusInternalServerError)
			return
		}

		// Print pod names
		for _, pod := range pods.Items {
			fmt.Fprintf(w, "Pod: %s\n", pod.Name)
		}
	})

	// HTTP handler function to list deployments
	http.HandleFunc("/deployments/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the namespace from the request URL
		namespace := r.URL.Path[len("/deployments/"):]
		if namespace == "" {
			namespace = "default" // Default to "default" namespace if no namespace is provided
		}

		// List deployments in the specified namespace
		deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error listing deployments in namespace %s: %s", namespace, err.Error()), http.StatusInternalServerError)
			return
		}

		// Print deployment names
		for _, deployment := range deployments.Items {
			fmt.Fprintf(w, "Deployment: %s\n", deployment.Name)
		}
	})

	// Start the HTTP server
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %s\n", err.Error())
		os.Exit(1)
	}
}
