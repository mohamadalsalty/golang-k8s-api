package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"

	"k8s/handlers"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}
func TestClusterInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/clusterinfo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ClusterInfoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedPattern := `Cluster Info:\n\s+Name: .*\n\s+Version: .*\n`
	matched, err := regexp.MatchString(expectedPattern, rr.Body.String())
	if err != nil {
		t.Fatalf("error matching regex: %v", err)
	}

	if !matched {
		t.Errorf("handler returned unexpected body: got %v want match for pattern %v",
			rr.Body.String(), expectedPattern)
	}

}

func TestNodesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nodes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.NodesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestPodsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/pods/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PodsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestNamespacesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/namespaces", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.NamespacesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeploymentsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/deployments/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeploymentsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestReplicaSetsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/replicasets/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeploymentsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDaemonSets(t *testing.T) {
	req, err := http.NewRequest("GET", "/daemonsets/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeploymentsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
