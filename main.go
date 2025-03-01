package main

import (
	"html"
	"html/template"
	"log"
	"net/http"
	"strings"
)


var todolists = make(map[string][]string)

// セッションIDに紐づくToDoリストを取得する．
func getTodoList(sessionId string) []string{
	todos, ok := todolists[sessionId]
	if !ok {
		todos = []string{}
		todolists[sessionId] = todos
	}

	return todos
}

// ToDoリストを返す
func handleTodo(w http.ResponseWriter, r *http.Request){
	sessionId, err := ensureSession(w, r)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}  
  todos := getTodoList(sessionId)

	// テンプレートからHTML生成
	t,_ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todos)
}

// 追加
func handleAdd(w http.ResponseWriter, r *http.Request) {
	sessionId, err := ensureSession(w, r)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}  
  todos := getTodoList(sessionId)

	r.ParseForm()
	todo := strings.TrimSpace(html.EscapeString(r.Form.Get("todo")))	
	if todo != ""{
		todolists[sessionId] = append(todos, todo)
	}
	// リダイレクト
	http.Redirect(w,r,"/todo", 303)
}


func main(){

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/todo", handleTodo)
	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}