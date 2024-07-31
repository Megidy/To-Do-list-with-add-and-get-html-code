package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type List struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func HandleAddToDoList(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/add-task.html")
	Data := Task{
		Title:       r.FormValue("Title"),
		Description: r.FormValue("Description"),
	}
	if Data.Description == "" {
		return
	}
	var NewTask List
	filepath := "static/data.json"

	jsonFile, err := os.Open(filepath)
	CheckError(err)

	ByteValue, err := ioutil.ReadAll(jsonFile)
	CheckError(err)

	err = json.Unmarshal(ByteValue, &NewTask)
	CheckError(err)

	NewTask.Tasks = append(NewTask.Tasks, Data)
	updatedData, err := json.Marshal(NewTask)
	CheckError(err)

	err = ioutil.WriteFile(filepath, updatedData, 0644)
	CheckError(err)
}

func HandleGetAllToDoLists(w http.ResponseWriter, r *http.Request) {
	var Data List
	filepath := "static/data.json"
	jsonFile, err := os.Open(filepath)
	CheckError(err)

	ByteValue, err := ioutil.ReadAll(jsonFile)
	CheckError(err)

	err = json.Unmarshal(ByteValue, &Data)
	CheckError(err)

	tmpl, err := template.ParseFiles("static/get-list.html")
	CheckError(err)

	err = tmpl.Execute(w, Data)
	CheckError(err)
}

func main() {
	http.HandleFunc("/add-task", HandleAddToDoList)
	http.HandleFunc("/list", HandleGetAllToDoLists)
	http.ListenAndServe(":8181", nil)

}
