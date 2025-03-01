package main

import (
	"html/template"
	"log"
	"net/http"
)


var todolist []string

// 表示
func handleTodo(w http.ResponseWriter, r *http.Request){
	// テンプレートからHTML生成
	t,_ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todolist)
}

// 追加
func handleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := r.Form.Get("todo")
	todolist = append(todolist, todo)
	// リダイレクト
	http.Redirect(w,r,"/todo", 303)
}


func main(){
	todolist = append(todolist, "顔を洗う", "朝食を食べる", "歯を磨く")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/todo", handleTodo)
	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}