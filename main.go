package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// cookieの設定を行う
func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "first",
		Value: "valueだぞ",
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Cookieの設定ができたよ")
}

// cookieの取得を行い、htmlを返す
func showCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("first")

	if err != nil {
		log.Fatal("Cookie: ", err)
	}

	tmpl := template.Must(template.ParseFiles("./cookie.html"))
	tmpl.Execute(w, cookie)

}

// メイン
func main() {
	http.HandleFunc("/", setCookies)
	http.HandleFunc("/cookie", showCookie)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
