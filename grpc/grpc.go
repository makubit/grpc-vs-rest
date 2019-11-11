package grpc

import(
	//"errors"
	//"log"
	"math/rand"
	"sync"

	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	pb "github.com/makubit/grpc-vs-rest/grpc/proto"
)

type repository interface {
	Sort(request *pb.SortRequest) (*pb.SortRequest, error)
}

type Repository struct {
	mu sync.RWMutex
	sortedTableRequest *pb.SortRequest
}

func (repo *Repository)Sort(sortRequest *pb.SortRequest) (*pb.SortRequest, error) {
	var sortedTable []int32
	repo.mu.Lock()
	sortedTable, _ = quickSort(sortRequest.GetTableToSort())
	repo.sortedTableRequest.TableToSort = sortedTable
	repo.mu.Unlock()

	return sortRequest, nil
}

type service struct {
	repo repository
}

func (s *service)Sort(ctx context.Context, req *pb.SortRequest) (*pb.Response, error)  {
	sortedTableRequest, err := s.repo.Sort(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Sorted: true, SortedTable: sortedTableRequest.TableToSort}, nil
}

//---------

func quickSort(unsortedTable []int32) ([]int32, error) {
	if len(unsortedTable) < 2 {
		return unsortedTable, nil
	}

	left, right := 0, len(unsortedTable)-1

	pivot := rand.Int() % len(unsortedTable)

	unsortedTable[pivot], unsortedTable[right] = unsortedTable[right], unsortedTable[pivot]

	for i, _ := range unsortedTable {
		if unsortedTable[i] < unsortedTable[right] {
			unsortedTable[left], unsortedTable[i] = unsortedTable[i], unsortedTable[left]
			left++
		}
	}

	unsortedTable[left], unsortedTable[right] = unsortedTable[right], unsortedTable[left]

	quickSort(unsortedTable[:left])
	quickSort(unsortedTable[left+1:])

	return unsortedTable, nil
}