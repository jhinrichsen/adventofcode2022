.PHONY: all
all: lint test

.PHONY: bench
bench:
	CGO_ENABLED=0 go test -bench=. -run="" -benchmem

.PHONY: lint
lint:
	CGO_ENABLED=0 go vet
	CGO_ENABLED=0 staticcheck

.PHONY: test
test:
	CGO_ENABLED=0 go test -coverprofile=coverage.txt -covermode count gitlab.com/jhinrichsen/adventofcode2022
	gocover-cobertura < coverage.txt > coverage.xml

prof:
	go test -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	go tool pprof cpu.profile
	# go tool pprof mem.profile
