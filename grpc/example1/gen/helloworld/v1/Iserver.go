package v1

import "golang.org/x/net/context"

type server interface {
	SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error)
}
