package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Content string `json:"content"`
}

func pageHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!!!")
	fmt.Println("Endpoint: Home")
}

func allPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: allPosts")
	json.NewEncoder(w).Encode(Posts)
}

func singlePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//fmt.Fprintf(w, "Key: "+key)

	// Posts is the array that is being iteratted over
	for _, value := range Posts {
		if value.Id == key {
			fmt.Println("Endpoint: returning " + key)
			json.NewEncoder(w).Encode(value)
		}
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	rBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(rBody))

	// x can be any name but Post is from the Global Array
	var x Post

	// Unmarshal this into a new struct
	json.Unmarshal(rBody, &x)

	// Appending value from above to the Global Posts array, both Posts have to match
	Posts = append(Posts, x)

	json.NewEncoder(w).Encode(x)
	fmt.Println("Endpoint: Post " + key + " Created")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	booty := mux.Vars(r)
	id := booty["id"]

	for d, a := range Posts {
		if a.Id == id {
			Posts = append(Posts[:d], Posts[d+1:]...)
		}
	}
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	booty := mux.Vars(r)
	id := booty["id"]

	var z Post

	rBody, _ := ioutil.ReadAll(r.Body)

	// Unmarshal this into a new struct
	json.Unmarshal(rBody, &z)

	for i, value2 := range Posts {
		if value2.Id == id {
			value2.Title = z.Title
			value2.Author = z.Author
			value2.Content = z.Content
			Posts[i] = value2
			json.NewEncoder(w).Encode(value2)
		}
	}

	fmt.Println("Endpoint: Post updated")
}

// Global Posts array that will be populated later
var Posts []Post

func handleFunc() {
	r := mux.NewRouter()

	//http.HandleFunc("/", pageHome)
	//http.HandleFunc("/posts", allPosts)
	//log.Fatal(http.ListenAndServe(":8000", nil))

	r.HandleFunc("/", pageHome)
	// maps this route to the allPosts function
	r.HandleFunc("/posts", allPosts)
	r.HandleFunc("/post", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE") // Needs to be before here
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", singlePost)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	// Simulates a DB
	Posts = []Post{
		Post{Id: "1", Title: "Hello", Author: "John Doe", Content: "Post 1"},
		Post{Id: "2", Title: "Hello World", Author: "Jane Doe", Content: "Post 2"},
	}

	handleFunc()
}
