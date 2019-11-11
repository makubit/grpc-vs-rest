package tests

import (
	"bytes"
	"encoding/json"
	"github.com/makubit/grpc-vs-rest/rest"
	"net/http"
	"testing"
	"time"
)

func init() {
	go rest.StartREST()
	time.Sleep(time.Second * 2)
}

func BenchmarkREST(t *testing.B) {
	client := &http.Client{}

	for i:=0; i< t.N; i++ {
		postToREST(client, t)
	}
}

func postToREST(client *http.Client, t *testing.B) {
	sort := &rest.SortRequest{
		TableToSort: []int32{1, 3, 4, 2, 3, 1},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(sort)

	res, err := client.Post("http://127.0.0.1:50052/", "application/json", buf)
	if err != nil {
		t.Fatalf("http request failed: %v", err)
	}

	defer res.Body.Close()

	var response rest.Response
	decodeErr := json.NewDecoder(res.Body).Decode(&response)
	if decodeErr != nil {
		t.Fatalf("unable to decode json: %v", decodeErr)
	}

	if response.Sorted != true {
		t.Fatalf("hhtp wrong response: %v", res)
	}
}