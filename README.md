# golagn-cookie
GoでCookie操作

## cookieの値
```go:main.go
cookie := &http.Cookie{
		Name:  "first",
		Value: "valueだぞ",
	}
```
## http.HandleFunc の1mmわかる解説
http.HandleFunc("/cookie", showCookie)
```go:main.go
	http.HandleFunc( url string, hander func(ResponseWriter, *Request))
```
第一引数にpathを文字列で指定
第二引数に実行したい関数を渡す(handlerを登録)