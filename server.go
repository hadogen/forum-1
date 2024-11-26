package main

import (
	"fmt"
	"net/http"

	creations "main/creations"
	handler "main/handler"
	reactions "main/reactions"
	userData "main/userData"
)

var port = "7070"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/login", userData.Login)
	http.HandleFunc("/guest", handler.Guest)
	http.HandleFunc("/register", userData.HandleRegistration)
	http.HandleFunc("/forum", handler.Forum) // this is where the forum would be handled after the login
	http.HandleFunc("/create_post", creations.CreatePost)
	http.HandleFunc("/InsertPost", creations.InsertPost)
	http.HandleFunc("/Likes", reactions.Like)
	http.HandleFunc("/api/likes", reactions.LikesCounterWithApi)
	http.HandleFunc("/Dislikes", reactions.Dislikes)
	http.HandleFunc("/logout", userData.Logout)
	http.HandleFunc("/style/", handler.Style)
	http.HandleFunc("/create_comment", creations.CreatCommnet)
	http.HandleFunc("/showcomments", creations.ShowComments)
	fmt.Println("Server started on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}