module bitbucket.org/makeitplay/lugo-frontend

go 1.14

replace github.com/lugobots/lugo4go/v2 => /home/rubens/go/src/github.com/lugobots/lugo4go

replace bitbucket.org/makeitplay/lugo-server => /home/rubens/go/src/bitbucket.org/makeitplay/lugo-server

require (
	bitbucket.org/makeitplay/lugo-server v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.0
	github.com/lugobots/lugo4go/v2 v2.0.0-alpha.10
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.13.0
	google.golang.org/grpc v1.31.0
)
