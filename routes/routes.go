package routes

import (
	"k8s/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/clusterinfo", enableCors(handlers.ClusterInfoHandler))
	http.HandleFunc("/namespaces", enableCors(handlers.NamespacesHandler))
	http.HandleFunc("/nodes", enableCors(handlers.NodesHandler))
	http.HandleFunc("/pods/", enableCors(handlers.PodsHandler))
	http.HandleFunc("/deployments/", enableCors(handlers.DeploymentsHandler))
	http.HandleFunc("/replicasets/", enableCors(handlers.ReplicaSetsHandler))
	http.HandleFunc("/daemonsets/", enableCors(handlers.DaemonSetsHandler))
}

// enableCors is a middleware function to enable CORS for handlers
func enableCors(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	}
}
