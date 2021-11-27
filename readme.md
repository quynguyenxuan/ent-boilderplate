#add model
go run -mod=mod entgo.io/ent/cmd/ent init User
#add enum fields
func (User) Fields() []ent.Field {

}

go generate ./ent
go generate ./...

edit only
ent/schema
ent/entc.go
ent/generate.go
graphql

ngrok to publish local server
