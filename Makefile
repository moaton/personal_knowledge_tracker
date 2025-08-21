.PHONY: docs
docs:
	swag i -g handler.go -dir internal/controller/http/v1 --instanceName v1 --parseDependency
