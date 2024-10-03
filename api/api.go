package api

import "net/http"

func Run() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/kill", kill)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("template/assets")))) // Corrigido o caminho do assets
	http.ListenAndServe(":9339", nil)
}
