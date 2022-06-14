## time_api
- Udemy講座「REST based microservices API development in Golang」の課題1

### go runの前に実行したコマンド
```
// Gorilla Muxを使う
go mod init time_api
go get -u github.com/gorilla/mux
```

### 参考
- go mod関連
  - https://casualdevelopers.com/tech-tips/lets-create-rest-api-with-gorilla-mux-on-golang/
- time.Location
  - https://qiita.com/yyoshiki41/items/3acfe3c03b5a3a1e7592
- w.WriteHeaderの例
  - https://qiita.com/gold-kou/items/99507d33b8f8ddd96e3a#%E3%83%91%E3%82%B9%E3%83%91%E3%83%A9%E3%83%A1%E3%83%BC%E3%82%BF
- クエリパラメータの使い方
  - https://zetcode.com/golang/gorilla-mux/
