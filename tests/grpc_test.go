package tests

import (
	"context"
	gp "github.com/makubit/grpc-vs-rest/grpc"
	pb "github.com/makubit/grpc-vs-rest/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func grpcSendToClient(client pb.SortingServiceClient, t *testing.B) {
	_, err := client.Sort(context.Background(), &pb.SortRequest{TableToSort: []int32{4,3,2,4,3,2}})
	if err != nil {
		log.Fatalf("Could not sort: %v", err)
	}

	//log.Println("Sorted: ", r.SortedTable)
}

func init() {
	go gp.StartGRPC()
	time.Sleep(time.Second * 2)
}

func BenchmarkGRPC(t *testing.B) {
	serive, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial: %v", err)
	}

	defer serive.Close()
	client := pb.NewSortingServiceClient(serive)

	for i := 0; i< t.N; i++ {
		grpcSendToClient(client, t)
	}
}