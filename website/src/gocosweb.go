package GocosWeb

import(
	"fmt"
	"net/http"
	"html/template"
)
var (
	templates = template.Must(template.ParseFiles(
		"index.html",
	))
)

func init(){
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request){
	url := fmt.Sprintf("http://%s/index.html", r.Host)
	templates.ExecuteTemplate(w, "index.html", url)
}

