## 2.3 少しリッチなhello world - 設定ファイル

YAML形式でコマンドの設定をファイル化してみます。

config.yml

```yaml
message: hello
times: 3
```

```console
$ ./richhello -c config.yml
```

~~~

### パッケージ取得

```console
$ go get gopkg.in/yaml.v2
```

~~~

### サンプル

`samples/richhello-config/main.go`

試しに何か設定項目を追加して、プログラムの挙動変えてみて！
