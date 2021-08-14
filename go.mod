module entgo.io/quynguyen-todo

go 1.16

require (
	entgo.io/contrib v0.0.0-20210709131656-a2b04157b717
	entgo.io/ent v0.9.0
	github.com/99designs/gqlgen v0.13.0
	github.com/AlekSi/pointer v1.1.0
	github.com/alecthomas/kong v0.2.11
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/arsmn/fastgql v0.14.0
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.2
	github.com/masseelch/elk v0.3.2
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.28.0
	github.com/vektah/gqlparser/v2 v2.1.0
	github.com/vmihailenco/msgpack/v5 v5.0.0-beta.9
	go.uber.org/zap v1.19.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

// replace github.com/mailru/easyjson => ./patches/github.com/mailru/easyjson
// replace github.com/masseelch/elk => ./patches/github.com/masseelch/elk
replace github.com/masseelch/elk => /Volumes/Project/go-workspace/src/elk/
