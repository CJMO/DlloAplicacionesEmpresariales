package main

import (
	//"fmt"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
	//"gopkg.in/unrolled/render.v1"
)

type Pagina struct{
  Titulo string
  Cuerpo []byte
}

var plantillas = template.Must(template.ParseFiles("plantillas/index.html", "plantillas/post.html"))

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controladorInicio)
	r.HandleFunc("/noticias", controladorNoticias)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	//http.Handle("/css/", http.FileServer(http.Dir("/css")))
	//http.Handle("/js", http.FileServer(http.Dir("/js")))
	//http.Handle("/img", http.FileServer(http.Dir("/img")))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("plantillas/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("plantillas/js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("plantillas/img"))))
	http.Handle("/scss/", http.StripPrefix("/scss/", http.FileServer(http.Dir("plantillas/scss"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("plantillas/vendor"))))

	http.ListenAndServe(":9000", nil)

	// render
	/*r2 := render.New(render.Options{})
	mux := http.NewServeMux()
	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
        // Assumes you have a template in ./templates called "example.tmpl".
        // $ mkdir -p templates && echo "<h1>Hello HTML world.</h1>" > templates/example.tmpl
        r2.HTML(w, http.StatusOK, "noticias", nil)
    })
	http.ListenAndServe(":9000", mux)*/
	

	/*http.HandleFunc("/", controladorInicio)
	fmt.Println("El servidor se encuentra en ejecución")
	http.ListenAndServe(":9000", nil)*/

	/*server := &http.Server{
		Addr:    ":9000",
		Handler: http.HandlerFunc(controladorNoticias),
	}

	fmt.Println("El servidor se encuentra en ejecución")
	server.ListenAndServe()*/
}

func controladorInicio(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w,"Hola, %s, ¡este es un servidor!", r.URL.Path)
	//w.Write([]byte("Página inicial!\n"))

	titulo := "Página inicial"
	p := &Pagina{Titulo:titulo, Cuerpo: []byte("")}
	cargarPlantilla(w, "index", p)
}

func controladorNoticias(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("Página de noticias!\n"))

	titulo := "Página de noticias"
	p := &Pagina{Titulo:titulo, Cuerpo: []byte("")}
	cargarPlantilla(w, "post", p)
}

//Carga las plantillas HTML
func cargarPlantilla(w http.ResponseWriter, nombre_plantilla string, pagina *Pagina){
  plantillas.ExecuteTemplate(w, nombre_plantilla + ".html", pagina)
}