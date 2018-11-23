## 1.足し算
<small>コーディング、実行、テスト、ビルドを一通り体験する</small>

標準出力から受け取ったスペース区切りの２つの整数を足した結果をコンソールに出力

```sh
$ ./add 1 2
3
```

- `os.Args`
- `strconv.Atoi(s) (int, error)`
- `fmt.Println(s)`
- `fmt.Printf(format, args...)`

~~~

### 作業ディレクトリ作成

作業ディレクトリは、今日はとりあえずGOPATH直下にします。

```console
$ cd $(go env GOPATH)
$ mkdir -p github.com/dodosoft/samples/add
$ code add
```

~~~

### サンプル

samples/add/main.go

~~~

### 実行

`go run`コマンドで、ビルドと実行を同時に行う。  
VSCode下部のターミナルで、以下を実行：

```console
$ cd samples/add
$ go run main.go 1 2
3
```

~~~

### 関数化

`main.go`を改変して、２つの文字列を受け取り足し合わせる関数を抽出する。  
戻り値は足した結果と、不正文字列だった場合のエラー

```go
package main

...

func add(s1 string, s2 string) (int, error) {
    ...
    return num1 + num2, nil
}
```

~~~

### テスト作成

`main.go`を開いて、Ctrl+Shift+P、`Go: Generate Unit Tests for Function`を選択

~~~

### テスト作成

ここを

```go
// TODO: Add test cases.
```

こんな感じに

```go
{
    name: "simple",
    args: args{
        s1: "1",
        s2: "2",
    },
    want: 3,
},
```

~~~

### テスト実行

- `run test`でその関数のテスト
- `run package tests`でそのパッケージ内のすべてのテスト実行＆カバレッジ測定

~~~

### ビルド/実行

```console
$ go build -o add
$ ./add 1 2
3
```

~~~

### クロスプラットフォームビルド

環境変数GOARCH/GOOSを指定することで、異なるプラットフォーム向けのバイナリがビルドできる。

例

```console
$ GOARCH=386 GOOS=windows go build -o add.exe
$ GOARCH=amd64 GOOS=darwin go build -o add
$ GOARCH=arm GOOS=linux go build -o add
```
