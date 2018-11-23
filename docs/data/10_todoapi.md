## 3. TODOリストAPI

こんな感じのAPI作る

```console
$ curl -XPOST localhost:8080/todolist/ -F "text=My Task1"
{"id":"0","done":false,"text":"My Task1"}

$ curl -XPOST localhost:8080/todolist/ -F "text=My Task2"
{"id":"0","done":false,"text":"My Task2"}

$ curl localhost:8080/todolist/
[{"id":"0","done":false,"text":"My Task1"},{"id":"1","done":false,"text":"My Task2"}]

$ curl localhost:8080/todolist/0
{"id":"0","done":false,"text":"My Task1"}

$ curl -XDELETE localhost:8080/todolist/0
{"id":"0","done":false,"text":"My Task1"}

$ curl -XPOST localhost:8080/todolist/1/done
{"id":"1","done":true,"text":"My Task2"}

$ curl -XPOST localhost:8080/todolist/1/undone
{"id":"1","done":false,"text":"My Task2"}

$ curl -XPOST localhost:8080/todolist/1/text -F "text=New My Task2"
{"id":"1","done":false,"text":"New My Task2"}
```

~~~

### サンプル

`samples/todo/main.go`