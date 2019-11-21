VERSION?="0.0.1"
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')


default: bin


dep:
	go mod download

fmt:
	@gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

bin: fmtcheck test
	sh -c "'$(CURDIR)/scripts/build.sh'"

test: fmtcheck
	sh -c "'$(CURDIR)/scripts/test.sh'"

cover:
	sh -c "'$(CURDIR)/scripts/cover.sh'"

psql:
	sh -c "docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres"

.NOTPARALLEL:

.PHONY: bin fmt fmtcheck test cover dep