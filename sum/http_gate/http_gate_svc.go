package http_gate

import (
	"context"
	"fmt"
	"net/http"
	"server/sum/common/discov"
)

func HttpServer() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("http listen on 8080")
	discov.Register(context.Background(), "http_gate", "127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
