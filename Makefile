BUILD_DIR := build

all: backend frontend containers

backend:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./$(BUILD_DIR)/k8sdemo-backend ./k8sdemo-backend

frontend:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./$(BUILD_DIR)/k8sdemo-frontend ./k8sdemo-frontend

containers:
	docker build -f Docker/backend.dockerfile -t k8sdemo_backend .
	docker build -f Docker/frontend.dockerfile -t k8sdemo_frontend .

clean:
	rm -rf $(BUILD_DIR)
