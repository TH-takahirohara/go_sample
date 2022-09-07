## docker-gs-ping
- [Go言語 - イメージのビルド](https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/)
- [Go言語 - コンテナーの実行](https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/run-containers/)

### 実行したこと
- [サイト](https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/)のmain.goのコードをコピー
- 下記コマンドを実行してローカルで動作を確認
```
go mod init docker-gs-ping
go mod tidy
go run main.go
```
- [サイト](https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/)内の記述に従い、イメージを作成
- 下記コマンドでコンテナを起動し、動作を確認
```
docker run -p 8080:8080 docker-gs-ping
```
- 下記コマンドでデタッチモードでコンテナを起動し、動作を確認
```
docker run -d -p 8080:8080 docker-gs-ping
```
