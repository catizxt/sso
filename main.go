package main

import (
	//"github.com/micro/go-log"
	"github.com/micro/go-micro"
	//"github.com/cicdi-go/sso/subscriber"
	"os"
	"log"	
	"github.com/micro/go-micro/service/grpc"
	sso "github.com/cicdi-go/sso/proto/sso"
	"github.com/cicdi-go/sso/handler"
)

func init() {
        file := "/usr/local/logfile/golangweb-log/" +"loginweb"+ ".log"
        logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
        if err != nil {
                panic(err)
        }
        log.SetOutput(logFile) // 将文件设置为log输出的文件
        log.SetPrefix("[qSkipTool]")
        log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
        return
}

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.sso"),
		micro.Version("v1.0.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	sso.RegisterSsoHandler(service.Server(), new(handler.Sso))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.sso", service.Server(), new(subscriber.Sso))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.sso", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
