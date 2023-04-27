// @Version 1.0.0
// @Title Backend API
// @Description Rest API cloned from TutorialEdge/create-rest-api-in-go-tutorial
// @ContactName SPears
// @ContactEmail abce@email.com
// @ContactURL http://www.example.com
// @TermsOfServiceUrl http://www.example.com
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Server http://localhost:5000 Local
// @Server http://www.example-dev.com Development
// @Server http://www.example.com Production
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader http bearer Input your token
// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

/**
* Types
**/

// Article - Struct representation of an article
type Article struct {
    Id      string `json:"id" example:"123"`
    Title   string `json:"title" example:"Reverend"`
    Desc    string `json:"desc" example:"article description"`
    Content string `json:"content" example:"Article content goes here"`
}

/**
* Variables
**/

var Articles []Article


/**
* Request handlers
**/

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

// @Title Get an article by ID.
// @Description Fetch the article associated with a specific ID.
// @Param articleID path  string  true  "Id of a specific article."
// @Success 200 object Article "Article JSON"
// @Resource articles
// @Route /articles/{articleID} [get]
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

// @Title Add a new article.
// @Description Create a new article.
// @Param  article  body  Article  true  "Info for an article."
// @Success  201  object  Article  "Article JSON"
// @Resource articles
// @Route /article [post]
func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

// @Title Delete an article by ID.
// @Description Delete the article associated with a specific ID.
// @Param articleID path  string  true  "Id of a specific article."
// @Success 204 object Article "No response"
// @Resource articles
// @Route /articles/{articleID} [delete]
func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    })

    handler := c.Handler(myRouter)

    log.Fatal(http.ListenAndServe(":5000", handler))
}

func main() {
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }

    fmt.Println("Articles API")
    fmt.Println()
    fmt.Println("Serving requests")

    handleRequests()
}