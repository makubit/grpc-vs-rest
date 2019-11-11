package grpc

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"

	pb "github.com/makubit/grpc-vs-rest/grpc/proto"
	q "github.com/makubit/grpc-vs-rest/quickSort"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port=":50051"
)

type repository interface {
	Sort(request *pb.SortRequest) ([]int32, error)
}

type Repository struct {
	mu sync.RWMutex
	sortedTableRequest *pb.SortRequest
}

func (repo *Repository)Sort(sortRequest *pb.SortRequest) ([]int32, error) {
	var sortedTable []int32
	repo.mu.Lock()
	sortedTable, _ = q.QuickSort(sortRequest.GetTableToSort())
	repo.mu.Unlock()

	return sortedTable, nil
}

type service struct {
	repo repository
}

func (s *service)Sort(ctx context.Context, req *pb.SortRequest) (*pb.Response, error)  {
	sortedTable, err := s.repo.Sort(req)

	//sortedTable,err := q.QuickSort(req.GetTableToSort())
	if err != nil {
		return nil, err
	}

	return &pb.Response{Sorted: true, SortedTable: sortedTable}, nil
}

func StartGRPC() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSortingServiceServer(s, &service{repo})
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serce: %v", err)
	}
}
