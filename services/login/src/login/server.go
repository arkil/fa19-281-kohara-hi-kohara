/*
	Login Microservice
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// MongoDB Config

var mongodb_server = "mongodb://localhost:27017"
var mongodb_database = "user"
var mongodb_collection = "user"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	if len(os.Getenv("MONGO")) != 0 {
		mongodb_server = os.Getenv("MONGO")
	}
	fmt.Println("Running Database on ",mongodb_server)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/login", login(formatter)).Methods("POST")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Login is alive!"})
	}
}

// API to add a new user
func login(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusBadRequest, struct{ Message string}{"Can not connect to mongo DB"})
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var requestBody Login_req
		_ = json.NewDecoder(req.Body).Decode(&requestBody)

		var post bson.M
		err = c.Find(bson.M{"Username": requestBody.Username, "Password": requestBody.Password}).
					Select(bson.M{"_id" : 1,"Username": 2, "Firstname" :3, "Lastname": 4}).
					One(&post)

		if err != nil {
			formatter.JSON(w, http.StatusBadRequest, struct{ Message string }{"Something went wrong"})
			panic(err)
		}
		formatter.JSON(w, http.StatusOK, post)
	}
}
