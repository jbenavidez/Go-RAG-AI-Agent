RAG_SERVICE=RagAgent
CLI_SERVICE=CliService

up_build:  build_rag_agent build_cli
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Starting Docker images..."
	docker-compose up 
	@echo "Docker images started!"


build_rag_agent:
	@echo "Building RAG IA Agent  binary..."
	cd ../rag-service && env GOOS=linux CGO_ENABLED=0 go build -o ${RAG_SERVICE} ./cmd/api
	@echo "Done!"


build_cli:
	@echo "Building clip service  binary..."
	cd ../cli-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CLI_SERVICE} ./cmd/cli
	@echo "Done!"