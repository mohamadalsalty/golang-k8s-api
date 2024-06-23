# Golang Kubernetes API

Este é um projeto simples em Golang para criar um API básico que lista informações sobre um cluster Kubernetes usando a API HTTP.

## Funcionalidades

- **/clusterinfo**: Retorna informações sobre o cluster, incluindo nome e versão.
- **/namespaces**: Lista todos os namespaces no cluster.
- **/nodes**: Lista todos os nós no cluster.
- **/pods/{namespace}**: Lista todos os pods no namespace especificado (padrão: '*').
- **/deployments/{namespace}**: Lista todos os deployments no namespace especificado (padrão: '*').

## Requisitos

- Go 1.22.4 ou superior
- Kubernetes
- Minikube (para desenvolvimento local)
- Cliente kubectl (opcional, para gerenciar o cluster)

## Instalação

1. Clone o repositório
2. Construa o projeto:

   ```
   make build
   ```

3. Configure as variáveis de ambiente no arquivo `.env`:

   ```
   K8S_API_HOST=https://<ip-do-k8s>:<porta>
   K8S_TOKEN=<SEU_TOKEN>
   ```

4. Execute a API:

   ```
   make run
   ```

   A API estará disponível em `http://localhost:8080`.

## Testes

Para executar os testes, utilize o seguinte comando:

```
make test
```

## Limpeza

Para limpar o diretório de compilação, utilize o seguinte comando:

```
make clean
```
