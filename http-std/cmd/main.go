package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/michal-laskowski/wax"
)

func main() {
	renderer := wax.New(wax.NewFsViewResolver(os.DirFS("./views/")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "Hello", "John")
	})

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
