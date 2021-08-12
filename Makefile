gen-ent:
	go generate ./ent
gen-gql:
	go generate ./...
serve-hotreload:
	air
serve:
	go run server.go
docker:
	dockdocker-compose up -d