package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"

	"github.com/rizkix/wired/delivery/grpc"
	"github.com/rizkix/wired/delivery/http"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func main() {
	e, _ := InitializeEvent()
	e.Start()
}

type Event struct {
	HTTP http.Handler
	GRPC grpc.Handler
}

func NewEvent(h http.Handler, g grpc.Handler) Event {
	return Event{HTTP: h, GRPC: g}
}

func (e Event) Start() {
	go e.HTTP.Instance.Start(viper.GetString(`server.address`))

	lis, err := net.Listen("tcp", viper.GetString(`grpc.address`))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := e.GRPC.Instance.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
