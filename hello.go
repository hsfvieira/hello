package main

import (
  "fmt"
  "./aluno"
  "net/http"
  "os"
)

func handle(w http.ResponseWriter, r *http.Request) {
  a1 := aluno.Aluno{Nome: "Henrique", Idade: 22}
  fmt.Fprintf(w, a1.Nome)
}

func main() {
  http.HandleFunc("/", handle)
  http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
