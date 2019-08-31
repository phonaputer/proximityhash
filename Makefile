DEPS_DIR=./vendor

clean:
	rm -rf $(DEPS_DIR)

deps:
	dep ensure

format:
	go fmt ./...
