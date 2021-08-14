patch:
	mkdir -p vendor/github.com/mailru/easyjson
	cp -R patches/github.com/mailru/easyjson/* vendor/github.com/mailru/easyjson/
gen-ent:
	go generate ./ent
gen-gql:
	go generate -mod=mod ./...
air:
	air
serve:
	go run server.go
docker:
	dockdocker-compose up -d