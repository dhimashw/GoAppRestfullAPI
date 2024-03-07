package article

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	Article{ID : 1 , Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	Article{ID : 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func GetListArticle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func ArrListArticle() []byte {
	urljson,_ := json.Marshal(Articles)
	return urljson;
}

func GetListArticleById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]	
    
	for _, item := range Articles {
		itemId := strconv.Itoa(item.ID);
		if itemId == key {	
			json.NewEncoder(w).Encode(item)
		}
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request){
	reqBody,_ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody,&article)
	
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(Articles)
}

func UpdateArticleById(w http.ResponseWriter, r *http.Request){
	reqBody,_ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody,&article)	
		
	for index, item := range Articles {		
		if item.ID == article.ID {
			Articles[index].Title = article.Title
			Articles[index].Desc = article.Desc			
			Articles[index].Content = article.Content			
		}
	}
	json.NewEncoder(w).Encode(Articles)

}
func DeleteArticleById(w http.ResponseWriter, r *http.Request){
	reqBody,_ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody,&article)	
	for index, item := range Articles {		
		if item.ID == article.ID {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(Articles)
}