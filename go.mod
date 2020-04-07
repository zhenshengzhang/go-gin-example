module github.com/zzs/go-gin-example

go 1.13

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20200322073714-2b22b57dfce9
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200406155108-e3b113bbe6a4 // indirect
)

replace (
	github.com/zzs/go-gin-example/conf => ./pkg/conf
	github.com/zzs/go-gin-example/middleware => ./pkg/middleware
	github.com/zzs/go-gin-example/models => ./pkg/models
	github.com/zzs/go-gin-example/pkg/e => ./pkg/e
	github.com/zzs/go-gin-example/pkg/setting => ./pkg/setting
	github.com/zzs/go-gin-example/pkg/util => ./pkg/util
	github.com/zzs/go-gin-example/routers => ./routers
)
