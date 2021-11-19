
fmt:	
	@echo "go fmt lfu package"	
	go fmt ./...

vet:
	@echo "go vet lfu package"	
	go vet ./...

lint:
	@echo "go lint lfu package"	
	golint ./...

test:
	@echo "Testing lfu package"	
	go test -v -race --cover ./...
