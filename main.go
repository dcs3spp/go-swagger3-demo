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

type ArticlesResponse struct {
    Data []Article `json:"data" example:"[{\"id\":123, \"title\":\"Article title\", \"desc\":\"Article Description\", \"content\":\"Article content\"}]"`
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


// @Title Get all articles.
// @Description Fetch all articles.
// @Success 200 object ArticlesResponse "A list of all articles"
// @Resource articles
// @Route /articles [get]
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

// @Title Get an article by ID.
// @Description Fetch the article associated with a specific ID.
// @Param id path  string  true  "Id of a specific article."
// @Success 200 object Article "The  article referenced by the id parameter"
// @Resource articles
// @Route /articles/{id} [get]
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
// @Success  201  object  Article  "The created article."
// @Resource articles
// @Route /articles [post]
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
// @Param  id  path  string  true  "Id of a specific article."
// @success 204 "No Response"
// @Resource articles
// @Route /articles/{id} [delete]
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
    myRouter.HandleFunc("/", homePage).Methods("GET")
    myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
    myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
        AllowCredentials: true,
    })

    handler := c.Handler(myRouter)

    log.Fatal(http.ListenAndServe(":5000", handler))
}


// @Version 1.0.0
// @Title Backend API
// @Description API TutorialEdge/create-rest-api-in-go-tutorial with swagger!
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
