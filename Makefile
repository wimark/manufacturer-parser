PROJECT_PREFIX = registry.wimark.tk/wne/
SERVICE_NAME = manufacturer-parser

SERVICE_VERSION := $$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
SERVICE_COMMIT := $$(git show -s --pretty=%H)
SERVICE_BUILD := 0


build:
	env CGO_ENABLED=0 go build -ldflags \
		"-X main.Version=$(SERVICE_VERSION) -X main.Commit=$(SERVICE_COMMIT) -X main.Build=$(SERVICE_BUILD)" \
		-o ./bin/$(SERVICE_NAME)

clean:
	rm ./bin/$(SERVICE_NAME)
