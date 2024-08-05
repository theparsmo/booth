package main


import (
    "html/template"
    "net/http"
    "bufio"
    "os"
)



type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
	Title string
    PageTitle string
    Todos     []Todo
}
  
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    line := scanner.Text()
    tmpl := template.Must(template.ParseFiles("./main.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
			Title: line,
            PageTitle: "My TODO list",
            Todos: []Todo{	
                {Title: "Booth", Done: false},
                {Title: "bome", Done: false},
                {Title: "Bo", Done: true},
                {Title: "oooo", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":8080", nil)

}