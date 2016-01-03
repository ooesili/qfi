VERSION = `git describe --tags`
LDFLAGS = -X main.version=${VERSION}
OS = darwin freebsd linux netbsd openbsd

dev: generate
	go build -ldflags "${LDFLAGS}"

cross: generate
	gox -ldflags "${LDFLAGS}" -os "${OS}" \
	  -output="build/${VERSION}_{{.OS}}_{{.Arch}}/{{.Dir}}"
	test -d dist || mkdir dist
	for build in build/${VERSION}*; do \
	  zip -jr dist/qfi_`basename $$build`.zip $$build; \
	done

generate:
	go generate ./...

clean:
	rm -rf qfi dist build

.PHONY: dev cross generate clean
