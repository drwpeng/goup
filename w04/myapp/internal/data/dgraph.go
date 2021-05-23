package data

import (
	"log"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

func NewDgraph(url string) *dgo.Dgraph {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}