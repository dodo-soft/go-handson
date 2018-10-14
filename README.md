# dodosoft Goハンズオン資料

## reveal.jsサーバー起動

```console
$ docker-compose up -d
```

## 公開

```console
$ docker-compose run --rm revealjs build
$ git add .
$ git commit -m "Updated foo"
```