all:
	CGO_ENABLED=0 go build .
	docker build -t jimminter.azurecr.io/pingpong:latest .
	az acr login -n jimminter
	docker push jimminter.azurecr.io/pingpong:latest

.PHONY: all