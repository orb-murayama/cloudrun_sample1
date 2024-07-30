package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
		// 第一引数のURLパターンと第二引数の関数を紐づける
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    port := os.Getenv("PORT")		// 環境変数を取得
    if port == "" {
        port = "8080"
    }

    log.Printf("Listening on port %s", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))		// Webサーバーを作成
}
