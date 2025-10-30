fmt:
	find . -type f -name "???.go" -execdir go fmt ./... \;

vet: fmt
	find . -type f -name "???.go" -execdir go vet ./... \;

test: vet
	find . -type f -name "???.go" -execdir go test -vet=all ./... \;

run: test
	go run ./...

clean:
	go clean -testcache
	find . -type f -name "[0-9][0-9][0-9]" -exec rm \{\} \;
	find . -type f -name "cpu.prof" -exec rm \{\} \;
	find . -type d -name ".idea" -exec rm -rf \{\} \;

# Targets that do not represent actual files
.PHONY: clean fmt run test vet
