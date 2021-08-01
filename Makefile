gen-ent:
	go generate ./ent
gen-gql:
	go generate ./...
serve-hotreload:
	air
serve:
	go run server.go