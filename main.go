package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

// RequestInfoはリクエスト情報を保持する構造体
type RequestInfo struct {
	Method string              `json:"method"`
	URL    string              `json:"url"`
	Header map[string][]string `json:"header"`
	Body   json.RawMessage     `json:"body"`
}

func main() {

	/*
		// 第一引数のURLパターンと第二引数の関数を紐づける
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("request /")
			fmt.Fprintf(w, "Hello, World! 4")
		})
		http.HandleFunc("/sample1", sample1)
		http.HandleFunc("/sample2", sample2)
		http.HandleFunc("/sample3", sample3)
		http.HandleFunc("/sample4", sample4)
		http.HandleFunc("/linecallback", linecallback)
	*/

	// ミドルウェア用にデフォルトのマルチプレクサを取得
	mux := http.NewServeMux()

	// 第一引数のURLパターンと第二引数の関数を紐づける
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! 4")
	})
	mux.HandleFunc("/sample1", sample1)
	mux.HandleFunc("/sample2", sample2)
	mux.HandleFunc("/sample3", sample3)
	mux.HandleFunc("/sample4", sample4)
	mux.HandleFunc("/sample5", sample5)
	mux.HandleFunc("/sample6", sample6)
	mux.HandleFunc("/sample7", sample7)
	mux.HandleFunc("/sample8", sample8)
	mux.HandleFunc("/sample9", sample9)
	mux.HandleFunc("/sample10", sample10)
	mux.HandleFunc("/linecallback", linecallback)
	mux.HandleFunc("/linecallback2", linecallback2)

	// ミドルウェアを追加
	wrappedMux := middleware(mux)

	port := os.Getenv("PORT") // 環境変数を取得
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil)) // Webサーバーを作成
	log.Fatal(http.ListenAndServe(":"+port, wrappedMux)) // Webサーバーを作成（ミドルウェア指定）
}

// ミドルウェア
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// リクエストURLをログに記録
		log.Println("Request", r.Method, r.URL.String())
		// 次のハンドラを呼び出す
		next.ServeHTTP(w, r)
	})
}

/**
 * 固定のパラメータを出力
 */
func sample1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hoge", r.FormValue("hoge"))
	fmt.Println("foo", r.FormValue("foo"))
}

/**
 * 任意のパラメータを出力
 */
func sample2(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("Parse error")
	}

	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

/**
 * リクエスト情報をJSON形式で出力
 */
func sample3(w http.ResponseWriter, r *http.Request) {
	log.Printf("Body=", r.Body)
	// リクエストボディを読み取る
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Body=", body)

	// リクエスト情報を構造体に格納する
	requestInfo := RequestInfo{
		Method: r.Method,
		URL:    r.URL.String(),
		Header: r.Header,
		Body:   json.RawMessage(body),
	}

	// 構造体をJSONにエンコードする
	jsonResponse, err := json.Marshal(requestInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSONをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

/**
 * リクエストBODY情報をTEXT形式で出力
 */
func sample4(w http.ResponseWriter, r *http.Request) {
	// リクエストボディを読み取る
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 読み取ったボディを文字列として表示
	log.Println("Request Body:", string(body))

	// クライアントにレスポンスを返す
	w.Header().Set("Content-Type", "text/plain")
	w.Write(body)
}

/**
 * GETパラメータを取得
 */
func sample5(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	// レスポンスとしてパラメータを表示
	fmt.Fprintf(w, "Name: %s, Age: %s\n", name, age)
}

/**
 * POSTパラメータ（フォームデータ）を取得
 */
func sample6(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// 入力フォームを返す
		t, _ := template.ParseFiles("public/form.html")
		t.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		// POSTパラメータを解析
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		fmt.Println("name:", r.Form)

		// フォームパラメータを取得
		username := r.FormValue("username")
		password := r.FormValue("password")
		getParam := r.FormValue("getParam") // GETパラメータも取得される
		fmt.Printf("Username: %s, Password: %s, getParam: %s\n", username, password, getParam)

		username = r.Form.Get("username")
		password = r.Form.Get("password")
		getParam = r.Form.Get("getParam") // GETパラメータも取得される
		fmt.Printf("Username: %s, Password: %s, getParam: %s\n", username, password, getParam)

		username = r.PostForm.Get("username")
		password = r.PostForm.Get("password")
		getParam = r.PostForm.Get("getParam") // GETパラメータは取得されない
		fmt.Printf("Username: %s, Password: %s, getParam: %s\n", username, password, getParam)

		mf := r.MultipartForm
		if mf != nil {
			// 通常のリクエスト
			for k, v := range mf.Value {
				fmt.Printf("%v : %v", k, v)
			}
		}

		// すべてのPOSTパラメータを取得して表示
		for key, values := range r.PostForm {
			for _, value := range values {
				fmt.Printf("Key: %s, Value: %s\n", key, value)
			}
		}

		http.Redirect(w, r, "/sample6", 301)
		// レスポンスとしてパラメータを表示
		//fmt.Fprintf(w, "Username: %s, Password: %s, getParam: %s\n", username, password, getParam)
	}
}

/**
 * POST（JSON）パラメータを取得
 */
func sample7(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// JSONリクエストボディをデコード
		var params map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&params)
		if err != nil {
			http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
			return
		}

		param1 := params["param1"]
		param2 := params["param2"]
		log.Println("param1:", param1)
		log.Println("param2:", param2)

		// データを再度JSON形式の文字列にエンコード
		responseJSON, err := json.Marshal(params) // []byte
		if err != nil {
			http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
			return
		}
		log.Printf("Received JSON: %s\n", responseJSON)

		// すべてのJSONパラメータを取得して表示
		for key, value := range params {
			fmt.Fprintf(w, "Key: %s, Value: %v\n", key, value)
		}
	} else {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

/**
 * 環境変数(.env)の取得
 */
func sample8(w http.ResponseWriter, r *http.Request) {
	// 環境変数を取得
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	token := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	// レスポンスとしてパラメータを表示
	fmt.Fprintf(w, "secret: %s, token: %s\n", secret, token)
}

/**
 * 環境変数の取得
 */
func sample9(w http.ResponseWriter, r *http.Request) {
	// 環境変数を取得
	secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	token := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	// レスポンスとしてパラメータを表示
	fmt.Fprintf(w, "secret: %s, token: %s\n", secret, token)
}

/**
 * MySQLの操作（1レコード取得）
 */
func sample10(w http.ResponseWriter, r *http.Request) {
	// 環境変数を取得
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbProtocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("DB_USER", dbUser)
	fmt.Println("DB_PASS", dbPass)
	fmt.Println("DB_PROTOCOL", dbProtocol)
	fmt.Println("DB_NAME", dbName)

	// レコード定義
	type Item struct {
		code   string
		name   string
		imgUrl string
	}
	var item Item

	// DB接続
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true", dbUser, dbPass, dbProtocol, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB Open：", err)
		panic(err.Error())
	}
	defer db.Close()

	code := "001"
	if err := db.QueryRow("SELECT code, name, img_url FROM items WHERE code = ? LIMIT 1", code).Scan(&item.code, &item.name, &item.imgUrl); err != nil {
		log.Fatal(err)
	}
	fmt.Println(item.code, item.name, item.imgUrl)

	// レスポンスとしてパラメータを表示
	fmt.Fprintf(w, "code: %s, name: %v, imgUrl: %v", item.code, item.name, item.imgUrl)
}

/**
 * LINE BOT
 */
func linecallback(w http.ResponseWriter, r *http.Request) {
	// 環境変数を取得
	secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	token := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	// BOTを初期化
	bot, botErr := linebot.New(secret, token)
	if botErr != nil {
		log.Fatal(botErr)
	}

	// リクエストからBOTのイベントを取得
	events, parseErr := bot.ParseRequest(r)
	if parseErr != nil {
		if parseErr == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		// イベントがメッセージ受信だった場合
		if event.Type == linebot.EventTypeMessage {
			message := event.Message.(*linebot.TextMessage).Text
			log.Println(message)
			message = "こんにちは"
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
			if err != nil {
				log.Print(err)
			}
		}
	}
}

/**
 * LINE BOT(.env使用)
 */
func linecallback2(w http.ResponseWriter, r *http.Request) {
	log.Println("linecallback2 1")
	// 環境変数を取得
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	token := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	log.Println("linecallback2 2")
	// BOTを初期化
	bot, botErr := linebot.New(secret, token)
	if botErr != nil {
		log.Fatal(err)
	}

	log.Println("linecallback2 3")
	// リクエストからBOTのイベントを取得
	events, parseErr := bot.ParseRequest(r)
	if parseErr != nil {
		if parseErr == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	log.Println("linecallback2 4")
	for _, event := range events {
		// イベントがメッセージ受信だった場合
		if event.Type == linebot.EventTypeMessage {
			message := event.Message.(*linebot.TextMessage).Text
			log.Println(message)
			message = "こんにちは"
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
			if err != nil {
				log.Print(err)
			}
		}
	}
	log.Println("linecallback2 5")
}
