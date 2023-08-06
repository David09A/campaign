package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func handle404(w http.ResponseWriter, r *http.Request) {
	// Obtener la ruta completa de 404.html
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "public", "404.html")

	// Leer el contenido del archivo 404.html
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error al leer el archivo 404.html", http.StatusInternalServerError)
		return
	}

	// Establecer el código de estado 404
	w.WriteHeader(http.StatusNotFound)

	// Establecer el tipo de contenido como "text/html"
	w.Header().Set("Content-Type", "text/html")
	// Escribir el contenido en la respuesta
	_, _ = w.Write(content)
}

func handleBiografia(w http.ResponseWriter, r *http.Request) {
	// Obtener la ruta completa de biografia.html
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "public", "biografia.html")

	// Leer el contenido del archivo biografia.html
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error al leer el archivo biografia.html", http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido como "text/html"
	w.Header().Set("Content-Type", "text/html")
	// Escribir el contenido en la respuesta
	_, _ = w.Write(content)
}

func handlePrincipalesLogros(w http.ResponseWriter, r *http.Request) {
	// Obtener la ruta completa de principales_logros.html
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "public", "principales_logros.html")

	// Leer el contenido del archivo principales_logros.html
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error al leer el archivo principales_logros.html", http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido como "text/html"
	w.Header().Set("Content-Type", "text/html")
	// Escribir el contenido en la respuesta
	_, _ = w.Write(content)
}

func handlePlanDeGobierno(w http.ResponseWriter, r *http.Request) {
	// Obtener la ruta completa de plan_de_gobierno.html
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "public", "plan_de_gobierno.html")

	// Leer el contenido del archivo plan_de_gobierno.html
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error al leer el archivo plan_de_gobierno.html", http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido como "text/html"
	w.Header().Set("Content-Type", "text/html")
	// Escribir el contenido en la respuesta
	_, _ = w.Write(content)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Obtener la ruta completa de index.html
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "public", "index.html")

	// Leer el contenido del archivo index.html
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error al leer el archivo index.html", http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido como "text/html"
	w.Header().Set("Content-Type", "text/html")
	// Escribir el contenido en la respuesta
	_, _ = w.Write(content)
}

func main() {
	// Manejador para la página de inicio ("/")
	http.HandleFunc("/", handleIndex)

	// Manejador para la ruta "/carlosnavas"
	http.HandleFunc("/carlos_navas", func(w http.ResponseWriter, r *http.Request) {
		// Redirigir a la página de inicio
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// Manejador para la página de biografía ("/biography")
	http.HandleFunc("/biografia", handleBiografia)

	// Manejador para la página de biografía ("/principales_logros")
	http.HandleFunc("/principales_logros", handlePrincipalesLogros)

	// Manejador para la página de plan_gobierno ("/biography")
	http.HandleFunc("/plan_de_gobierno", handlePlanDeGobierno)

	// Manejador para cualquier otra ruta que no sea "/carlosnavas"
	http.HandleFunc("/404", handle404)

	// Manejador para los archivos estáticos (imágenes, CSS, etc.)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Iniciar el servidor en el puerto 8080
	http.ListenAndServe(":8080", nil)
}
