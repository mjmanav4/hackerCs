package panicHandler

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// TODO WRITE ON OWN ALL THE FUNCTIONS

type PanicHandler struct{}

func handleCrash() {
	if r := recover(); r != nil {
		glog.Errorf("Panic recovery: %v %v", codes.Internal, r)
	}
}

// UnaryPanicHandler to handle unary intercerption
func (p *PanicHandler) UnaryPanicHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	defer handleCrash()

	return handler(ctx, req)
}

func (p *PanicHandler) StreamPanicHandler(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {

	defer handleCrash()

	return handler(srv, stream)

}
