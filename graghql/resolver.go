package graghql

import (
	"context"
	pb "github.com/peranikov/sunaba/grpc/lib"
	"google.golang.org/grpc"
	"golang.org/x/sync/errgroup"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Greet(ctx context.Context, persons []*Person) (*Greet, error) {
	conn, err := grpc.DialContext(ctx, ":50051")
	if err != nil {
		return nil, err
	}

	cli := pb.NewGreeterClient(conn)
	srm, err := cli.SayHelloStream(ctx)
	if err != nil {
		return nil, err
	}

	eg := errgroup.Group{}
	for _, person := range persons {
		err := srm.Send(&pb.HelloRequest{ Name: person.Name })
		if err != nil {
			return nil, err
		}
	}

	err = srm.CloseSend()
	if err != nil {
		return nil, err
	}
}
