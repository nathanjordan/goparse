PKGS ?= $(shell glide novendor)
PKG_FILES ?= *.go

goparse: $(wildcard *.go) $(wildcard **/*.go)
	go build -o $@

.PHONY: test
test: goparse
	go test -race $(PKGS)

.PHONY: cover
cover: goparse
	go test -cover -coverprofile cover.out -race $(PKGS)

.PHONY: coveralls
coveralls: cover
	goveralls -service=travis-ci || echo "Coveralls failed"
