.PHONY: fmt

fmt:
	for go in examples/*/*.go *.go; do \
		go fmt $$go; \
	done; \
