patch:
	mkdir -p vendor/github.com/mailru/easyjson
	cp -R patches/github.com/mailru/easyjson/* vendor/github.com/mailru/easyjson/
gen-ent:
	go generate ./ent
gen-gql:
	go generate -mod=mod ./...
gen-doc:
	swag init
get-schema:
	gq http://localhost:8060/v1/graphql --introspect > schema.graphql
diagram: 
	enter ./ent/schema
air:
	air
serve:
	go run server.go
docker:
	dockdocker-compose up -d