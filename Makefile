patch:
	# mkdir -p vendor/github.com/masseelch/elk-fiber
	# cp -R patches/github.com/masseelch/elk/* vendor/github.com/masseelch/elk-fiber/
	mkdir -p vendor/github.com/mailru/easyjson
	cp -R patches/github.com/mailru/easyjson/* vendor/github.com/mailru/easyjson/
link:
	# mkdir -p /Users/pq/go/pkg/mod/github.com/masseelch/elk@v0.3.3
	sudo cp -R patches/github.com/masseelch/elk/* /Users/pq/go/pkg/mod/github.com/masseelch/elk@v0.3.2
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