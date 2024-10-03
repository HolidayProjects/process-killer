package api

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"path"
	"fmt"
	"time"

	"github.com/PyMarcus/process-killer/internal"
)

type Items struct {
	ID       int
	Name     string
	Memory   float32
	CPU      float32
	ImageUrl string
	Message  string 
}

// kill: Captura o id do processo e o mata
func kill(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	exec.Command("sudo", "kill", "-9", id)

	message := fmt.Sprintf("Process: %s removed!", id)
	http.SetCookie(w, &http.Cookie{
		Name:    "message",
		Value:   message,
		Path:    "/",
		Expires: time.Now().Add(5 * time.Minute), // Expira em 5 minutos
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// renderTemplate: carrega o template html
func renderTemplate(w http.ResponseWriter, r *http.Request) {
	templatePath := path.Join("template", "index.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Fail to load template %s", err)
		return
	}

	cookie, err := r.Cookie("message")
	var message string
	if err == nil {
		message = cookie.Value
		http.SetCookie(w, &http.Cookie{
			Name:   "message",
			Value:  "",
			Path:   "/",
			Expires: time.Now().Add(-1 * time.Hour), // Expira imediatamente
		})
	}

	imgUrl := "/assets/excluir.png"
	var items []Items
	for _, p := range internal.Run() {
		items = append(items, Items{
			ID:       p.Id,
			Name:     p.Name,
			Memory:   p.Memory,
			CPU:      p.CPU,
			ImageUrl: imgUrl,
			Message:  message,
		})
	}

	data := struct {
        Items  []Items
        Message string
    }{
        Items:  items,
        Message: message, 
    }

    tmpl.Execute(w, data)
}
