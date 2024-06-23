package routes

import (
	"k8s/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/clusterinfo", handlers.ClusterInfoHandler)
	http.HandleFunc("/namespaces", handlers.NamespacesHandler)
	http.HandleFunc("/nodes", handlers.NodesHandler)
	http.HandleFunc("/pods/", handlers.PodsHandler)
	http.HandleFunc("/deployments/", handlers.DeploymentsHandler)
}
