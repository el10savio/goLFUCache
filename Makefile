
fmt:	
	@echo "go fmt lfu package"	
	go fmt ./...

vet:
	@echo "go vet lfu package"	
	go vet ./...

lint:
	@echo "go lint lfu package"	
	golint ./...
