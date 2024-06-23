// tests/handlers_test.go

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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// Handle error loading .env file
		panic(err)
	}

	// Run tests
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

	// Use regular expression to match anything after "Name: " and "Version: "
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

	// Add more assertions for the NodesHandler response if needed
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

	// Add more assertions for the PodsHandler response if needed
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

	// Add more assertions for the NamespacesHandler response if needed
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

	// Add more assertions for the DeploymentsHandler response if needed
}
