## 2.1 少しリッチなhello world

<small>外部パッケージを取得して使ってみる</small>

オプション指定機能付きのCLIを作る

```
$ ./richhello
hello
$ ./richhello --times 3
hello
hello
hello
$ ./richhello -n 2
hello
hello
$ ./richhello -n 3 "foo"
foo
foo
foo
```

~~~

### パッケージの取得

```console
$ go get github.com/urfave/cli
```

取得したものは以下の場所に展開されます。

```console
ls $(go env GOPATH)/src/github.com/urfave/cli
```

ちなみに、goファイル群、別に特別なことはなく、同じような内容です。

```console
$ head $(go env GOPATH)/src/github.com/urfave/cli/app.go
```

gitリポジトリとしてアクセスできるようにすれば、パッケージとして公開できる!

[awesome-go](https://github.com/avelino/awesome-go)とか見てみると楽しいよ！

~~~

### サンプル

`samples/richhello/main.go`

~~~

暇だったら色つけたり、動かしたりしてみても面白いよ！

- 色: https://github.com/ttacon/chalk
- プログレスバー: https://github.com/schollz/progressbar
- ターミナルUI: https://github.com/gizak/termui