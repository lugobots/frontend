module github.com/lugobots/frontend

go 1.16

replace github.com/lugobots/lugo4go/v2 => /home/rubens/projects/lugo/lugo4go

replace bitbucket.org/makeitplay/lugo-server => /home/rubens/projects/lugo/lugo-server

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.0
	github.com/lugobots/lugo4go/v2 v2.0.0-alpha.10
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/rubens21/srvmgr v0.0.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.13.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/grpc v1.31.0
)
