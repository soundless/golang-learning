GOCMD=go
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: deps
	$(GOTEST) -cover ./...
deps:
	$(GOGET) github.com/google/go-cmp/cmp
test:
	$(GOTEST) -v ./...
debug:
	$(GOTEST) -v ./...
bench:
	$(GOTEST) -v -bench . -benchmem ./...
git-%:
	git add . && git commit -m "$(@:git-%=%)" && git push origin master
