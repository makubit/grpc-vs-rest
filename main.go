package main

import (
	"bytes"
	"context"
	"encoding/json"
	//gp "github.com/makubit/grpc-vs-rest/grpc"
	r "github.com/makubit/grpc-vs-rest/rest"
	pb "github.com/makubit/grpc-vs-rest/grpc/proto"
	"github.com/makubit/grpc-vs-rest/rest"
	"google.golang.org/grpc"
	"log"
	"net/http"
	//"time"
)

func grpcClient() {
	serive, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial: %v", err)
	}

	defer serive.Close()
	client := pb.NewSortingServiceClient(serive)

	r, err := client.Sort(context.Background(), &pb.SortRequest{TableToSort: []int32{4,3,2,4,3,2}})
	if err != nil {
		log.Fatalf("Could not sort: %v", err)
	}

	log.Println("Sorted: ", r.SortedTable)
}

func restClient() {
	client := &http.Client{}

	sort := &rest.SortRequest{
		TableToSort: []int32{1, 3, 4, 2, 3, 1},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(sort)

	res, err := client.Post("http://127.0.0.1:50051/", "application/json", buf)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}

	defer res.Body.Close()

	var response http.Response
	decodeErr := json.NewDecoder(res.Body).Decode(response)
	if decodeErr != nil {
		log.Fatalf("unable to decode json: %v", decodeErr)
	}

	if response.StatusCode != 200 {
		log.Fatalf("hhtp wrong response: %v", res)
	}
}

func main() {
	/*go gp.StartGRPC()
	time.Sleep(time.Second*3)
	grpcClient()*/

	r.StartREST()
}