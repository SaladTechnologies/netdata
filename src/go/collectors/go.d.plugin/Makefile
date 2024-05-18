build: clean
	CGO_ENABLED=0 GOOS=linux go build -o go.d.plugin cmd/godplugin/main.go

clean:
	rm -f go.d.plugin

lint:
	golangci-lint run modules/salad/...
