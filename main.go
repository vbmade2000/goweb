/*
 * Goweb chat demo for Gin + Melody
 *
 * compile with 'go build -o goweb main.go' and
 * run with './goweb'
 */

// package declaration
package main

// imports
import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"net/http"
)

// main function
func main(){
	router:=gin.Default()	// gin router
	mel:=melody.New()		// melody middleware

	// static route (for CSS/JS)
	router.Static("/public","static/public")

	// root route ('/')
	router.GET("/",func(ctx *gin.Context){
		// serve a static HTML file
		http.ServeFile(ctx.Writer,ctx.Request,"static/index.html")
	})

	// websocket route
	router.GET("/ws",func(ctx *gin.Context){
		// handle request with Melody
		mel.HandleRequest(ctx.Writer,ctx.Request)
	})

	// Melody message handler
	mel.HandleMessage(func(ses *melody.Session,msg []byte){
		// broadcast message to connected sockets
		mel.Broadcast(msg)
	})

	// run app on port 3000
	router.Run(":3000")
}