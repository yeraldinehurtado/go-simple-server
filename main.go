package main

import "net/http"

func main() {

	// rutas - routes
	http.HandleFunc("/", homeHandler) // ruta inicial

	// inicia servidor - start the server
	http.ListenAndServe(":3000", nil) // puerto 3000 y nil significa que no hay error.
}

func homeHandler(w http.ResponseWriter, r *http.Request) { // r es la petición , w es la respuesta
	w.Write([]byte("Hello World")) //arreglo de un byte
}
