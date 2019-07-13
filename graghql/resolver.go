package graghql

import (
	"context"
	pb "github.com/peranikov/sunaba/grpc/lib"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Greet(ctx context.Context, persons []*Person) (*Greet, error) {
	conn, err := grpc.DialContext(ctx, ":50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	cli := pb.NewGreeterClient(conn)
	srm, err := cli.SayHelloStream(ctx)
	if err != nil {
		return nil, err
	}

	eg := errgroup.Group{}
	for _, p := range persons {
		p := p // https://golang.org/doc/faq#closures_and_goroutines
		eg.Go(func() error {
			return srm.Send(&pb.HelloRequest{ Name: p.Name })
		})
	}

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	rep, err := srm.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &Greet{
		Message: rep.Message,
	}, nil
}
