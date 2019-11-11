package rest

import (
	"encoding/json"
	"log"
	"net/http"

	q "github.com/makubit/grpc-vs-rest/quickSort"
)

func StartREST() {
	http.HandleFunc("/", Sort)
	log.Println(http.ListenAndServe(":50052", nil))
}

type SortRequest struct {
	TableToSort []int32 `json:"tableToSort"`
}

type Response struct {
	Sorted      bool    `json:"sorted"`
	SortedTable []int32 `json:"tableToSort"`
}

func Sort(req http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var unsortedTable []int32
	_ = decoder.Decode(&unsortedTable)
	defer r.Body.Close()

	req.Header().Set("Content-Type", "application/json")

	sortedTable, _ := q.QuickSort(unsortedTable)

	_ = json.NewEncoder(req).Encode(Response{
		Sorted:      true,
		SortedTable: sortedTable,
	})
}

