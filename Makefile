fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

run: test
	go run ./...

clean:
	find . -type d -name ".idea" -exec rm -rf \{\} \;
	find . -type f -name "[0-9][0-9][0-9]" -exec rm \{\} \;
	find . -type f -name "cpu.prof" -exec rm \{\} \;

# Targets that do not represent actual files
.PHONY: clean fmt run test vet
