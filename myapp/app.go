package app

import(
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render
type Todo struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"createdat"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request){
	list := []*Todo{}
	for _, v := range todoMap{
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTestTodos(){

	todoMap[1] = &Todo{1, "a",false,time.Now()}
	todoMap[2] = &Todo{2, "bb",false,time.Now()}
	todoMap[3] = &Todo{3, "bb",true,time.Now()}
}

func addTodoHandler(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	id := len(todoMap) + 1
	todo := &Todo{id, name , false, time.Now()}
	//id 넣기
	todoMap[id] = todo
	rd.JSON(w, http.StatusOK, todo)
}

func MakeHandler() http.Handler{
	todoMap = make(map[int]*Todo)
	//testdate
	addTestTodos()
	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/",indexHandler)
	return r
}
