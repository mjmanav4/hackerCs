package main

import (
	"context"
	"flag"
	"fmt"
	"hackerCs/feedImpl"
	ph "hackerCs/panicHandler"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	client "hackerCs/clients"
	"hackerCs/feed"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcPort int
	httpPort int
)

var grpcServer *grpc.Server
var panicHandler *ph.PanicHandler

func init() {
	flag.IntVar(&grpcPort, "grpcPort", 50052, "The server port")
	flag.IntVar(&httpPort, "httpPort", 50051, "The server port")
	flag.Parse()
}
func initGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}

	uInt := grpc.UnaryInterceptor(panicHandler.UnaryPanicHandler)
	sInt := grpc.StreamInterceptor(panicHandler.StreamPanicHandler)

	grpcServer = grpc.NewServer(uInt, sInt)
	feed.RegisterFeedServiceServer(grpcServer, feedImpl.NewUserServer())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Fatalf("Grpc server failed to listen : %v", err)
	}
	glog.Infof("GrpcServer started at port: %v", grpcServer)

}

func initHTTPServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := feed.RegisterFeedServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("0.0.0.0:%d", grpcPort), opts)
	if err != nil {
		panic(err)
	}
	glog.Infof("Starting user http server on port %d", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux); err != nil {
		panic(err)
	}
}
func initState() {
	//dao.InitMysqlConnection()
	//helper.InitFileStore()
	client.InitializeIdeaClient()
}

func closeState() {
	//dao.TearDownMysqlConnection()
}

func main() {

	go initHTTPServer()

	initState()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigs
		glog.Infof("Server shutting down. Signal: %v", sig)
		grpcServer.GracefulStop()
		closeState()
		done <- true
	}()
	//	glog.Infof("hellop")

	initGrpcServer()
	<-done
}
