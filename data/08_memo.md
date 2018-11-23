## 3. メモAPI

POSTで送ったメッセージを、GETで取得できるAPI

```console
$ curl -XPOST localhost:8080/memo -Fmessage=hello
hello
$ curl localhost:8080/memo
hello
```

~~~

### echoの利用

```console
$ go get -u github.com/labstack/echo/...
```

- echo.New(): echoインスタンス作成
- (*echo.Echo) POST(path, handlerFunc): POSTリクエストハンドラを設定
- (*echo.Echo) GET(path, handlerFunc): GETリクエストハンドラを設定
- (*echo.Echo) Start(address): echoによるHTTPサーバーの開始
- (echo.Context) String(code, s): HTTPレスポンスコード`code`で文字列のレスポンスを返す。

~~~

### サンプル

`samples/memo/main.go`