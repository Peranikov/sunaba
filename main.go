package main

import (
	"github.com/peranikov/sunaba/graphql"
	"github.com/peranikov/sunaba/grpc"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	eg := errgroup.Group{}
	eg.Go(grpc.Run)
	eg.Go(graphql.Run)

	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}
}
