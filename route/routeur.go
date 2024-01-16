package route

import (
	"fmt"
	"net/http"
	"os"
	controller "piscine/controller"
)

func InitRoute() {
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/save", controller.Save)
	http.HandleFunc("/display", controller.Display)
	http.HandleFunc("/delete", controller.DeleteCharacter)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/index) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Server closed")
}
