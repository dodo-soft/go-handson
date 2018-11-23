## 2.2 少しリッチなhello world - サブコマンド

git add, git commitのようなサブコマンド形式にしてみます。

```go
$ ./richhello times -n 3 hello
hello
hello
hello
$ ./richhello reverse hello
olleh
```

~~~

### サンプル

`samples/richhello-sub/main.go`

~~~

暇だったら弊社だいすきExcelとか使ってみても面白いかもしれない

- https://github.com/360EntSecGroup-Skylar/excelize