GO_BIN := bin
GQL_CONFIG_ROOT := schema/gqlgen.yml

$(GO_BIN):
	@mkdir -p $(GO_BIN)

build: $(GO_BIN)
	GO111MODULE=on go build -o $(GO_BIN)/sample-server main.go

run: build
	./bin/sample-server --config-files-dir ./config/yaml

clean:
	@rm -rf $(GO_BIN)
