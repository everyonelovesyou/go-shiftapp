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
	router.StaticFS("/assets", http.Dir("./public"))

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("起動失敗しちゃった", err)
	}
}

func handler(ctx *gin.Context) {
	data := []string{"mon", "tue", "wed", "thu", "fri"}
	t := template.Must(template.ParseFiles("tpl/layout.html", "tpl/index.html", "tpl/component/day.html"))
	t.Execute(ctx.Writer, data)
}
