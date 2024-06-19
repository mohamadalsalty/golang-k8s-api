# Golang Kubernetes API

Este é um projeto simples em Golang para criar um api básico que lista informações sobre um cluster Kubernetes usando a API HTTP.

## Funcionalidades

- **/clusterinfo**: Retorna informações sobre o cluster, incluindo nome e versão.
- **/namespaces**: Lista todos os namespaces no cluster.
- **/nodes**: Lista todos os nós no cluster.
- **/pods/{namespace}**: Lista todos os pods no namespace especificado (padrão: 'default').
- **/deployments/{namespace}**: Lista todos os deployments no namespace especificado (padrão: 'default').

## Requisitos

- Go 1.22.4 ou superior
- Kubernetes
- Minikube (para desenvolvimento local)
- Cliente kubectl (opcional, para gerenciar o cluster)

## Instalação

1. Clone o repositório


2. Construa o projeto:

   ```
   go build -o k8s-api
   ```

3. Execute a api:

   ```
   ./k8s-api
   ```

   A api estará disponível em `http://localhost:8080`.

## Configuração do Kubeconfig

Por padrão, a api usa o arquivo de configuração do Kubernetes localizado em `$HOME/.kube/config`. Você pode especificar um caminho diferente usando a flag `-kubeconfig`.

Exemplo:

```
./k8s-api -kubeconfig=/caminho/para/seu/kubeconfig
```

