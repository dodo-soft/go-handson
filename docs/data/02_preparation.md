## 前準備

SDKと開発環境を整えます。

<small>
※SDKはインストーラでもインストールできますが、  
みんなanyenvが好きだと思うのでgoenv使います。
</small>

---

## goenv

`anyenv`インストール

```console
$ git clone https://github.com/riywo/anyenv ~/.anyenv
$ echo '# anyenv' >> ~/.bashrc
$ echo 'export PATH="$HOME/.anyenv/bin:$PATH"' >> ~/.bashrc
$ echo 'eval "$(anyenv init - bash)"' >> ~/.bashrc
$ exec $SHELL -l
```

<small>※↑シェル変えてる人は適宜変更して</small>

`goenv`、Go 1.11.1をインストール

```console
$ anyenv install goenv
$ goenv install 1.11.1
$ goenv global 1.11.1
```

---

## Visual Studio Code(任意)

以下２つをインストール

1. VSCode本体  
   https://code.visualstudio.com/Download
2. Extension  
   https://code.visualstudio.com/docs/languages/go

<img src="data/images/vscode-extension.png" style="width:400px">

---

## Hello World実行

```go
package main

import "fmt"

func main(){
    fmt.Println("hello world")
}
```

実行

```console
$ go run main.go
```

~~~

<img src="data/images/helloworld.png" style="width:700px">

~~~

## 今日のサンプルコード取得

```console
$ go get github.com/dodo-soft/go-handson/...
$ code $(go env GOPATH)/dodo-soft/go-handson/samples
```
