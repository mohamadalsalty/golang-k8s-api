package config

import (
	"os"

	"k8s.io/client-go/rest"
)

func GetToken() string {
	return os.Getenv("K8S_TOKEN")
}

func GetAPIHost() string {
	host := os.Getenv("K8S_API_HOST")
	return host
}

func GetKubernetesConfig() (*rest.Config, error) {
	token := GetToken()
	host := GetAPIHost()

	config := &rest.Config{
		Host:        host,
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	return config, nil
}
