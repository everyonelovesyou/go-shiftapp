package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler)
	router.StaticFS("/assets", http.Dir("./src/public"))

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("起動失敗しちゃった", err)
	}
}

func handler(ctx *gin.Context) {
	data := []string{"mon", "tue", "wed", "thu", "fri"}
	t := template.Must(template.ParseFiles(
		// これ関数にしたい
		"./src/tmpl/layout.html",
		"./src/tmpl/index.html",
		"./src/tmpl/ctt/form.html",
	))
	t.Execute(ctx.Writer, data)
}
