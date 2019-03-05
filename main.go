package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

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
	app, err := InitializeApp()
	if err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}
	app.Start()
}

type App struct {
	HTTP http.Handler
	GRPC grpc.Handler
}

func NewApp(h http.Handler, g grpc.Handler) App {
	return App{HTTP: h, GRPC: g}
}

func (e App) Start() {
	go func() {
		e.HTTP.Instance.Start(viper.GetString(`server.address`))
	}()

	go func() {
		lis, err := net.Listen("tcp", viper.GetString(`grpc.address`))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := e.GRPC.Instance.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server")

	// give n seconds for server to shutdown gracefully
	duration := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err := e.HTTP.Instance.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shut down server gracefully: %s", err)
	}

	e.GRPC.Instance.GracefulStop()
	fmt.Printf("Server shutted down")
}
