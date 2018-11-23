## Goの文法

今日必要な文法を紹介

だいたい雰囲気でいけそう？

~~~

### 変数定義

```go
var i int = 10
var i = 10
i := 10
```

~~~

### for

```go
for i := 0; i < 10; i++ {
	// ...
}
for {

}
for i < 10 {

}
```

- whileはない
- ()は不要

~~~

### if

```go
if i < 10 {
}
```

~~~

### switch

```go
switch s {
    case "a":
        // break不要
    case "b":
    default:
        fmt.Println("aaa")
}
```

~~~

### 関数

```go
func main(){
    fmt.Println(f(2))
}
func f(n int) int {
    return n * 2
}
```

~~~

### エラー処理

```go
func main(){
    if err := func1(); err != nil {
        fmt.Println("failed")
    }
}
func func1() error {
    return errors.New("an error")
}
```

```go
func main(){
    s, err := func2()
    if err == nil {
        fmt.Println(s)
    } else {
        fmt.Println("failed")
    }
}
func func2() (string, error) {
    return "result", nil
}
```