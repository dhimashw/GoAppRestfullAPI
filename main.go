package main

import (
	"fmt"
	article "goappcards/Controllers/Article"
	calculate "goappcards/MyLibrary"
	"net/http"

	"github.com/gorilla/mux"
)

// type ReturnMessage struct{
// 	code string `json:code`
// 	message string `json:message`
// 	data <T>
// }

func main() {

	// var cards string = "loha loha"
	// dec := 10.33
	// dec = 10.44
	// fmt.Println(cards +  fmt.Sprintf(" %f",dec))	
	// fmt.Println("total : ", calculate.MyCalculate(11,11))
	// fmt.Println("total : ", calculate.MyMultiply(11,11))

	// calculate.LoopByValue(5)
	// calculate.FuncFloat()
	// fmt.Println(calculate.Percent)

	// arrayItems := []string{"array1","array2","array3"}
	// arrayItems2 := []string{"array4", "array5"}

	// arrayItems = append(arrayItems, arrayItems2...)

	// calculate.LoopArray(arrayItems)

	// deck := arrayItems;
	// calculate.LoopArrayByType(deck)

	// calc1 := mystruct.CalcVY{3,4}

	// calc1.X = 9
	// calc1.Scale(10)
	// fmt.Println(calc1.CalcPrint())
    
	// fmt.Println(mystruct.DecimalToString(1.3))
	// fmt.Println(mystruct.IntToString(1))

	// A,X := mystruct.MultipleReturn()
	// fmt.Println(A,X)

	// calculate.Export()

	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome to the Homepage!")
	fmt.Println("Enpoint Hit : homepage")
}

func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/api/article/listarticle", article.GetListArticle)
	myRouter.HandleFunc("/api/article/exportarticle", calculate.ExportListArticle)
	myRouter.HandleFunc("/api/article/getarticlebyid/{id}", article.GetListArticleById)
	myRouter.HandleFunc("/api/article/deletearticlebyid", article.DeleteArticleById).Methods("POST")	
	myRouter.HandleFunc("/api/article/createarticle", article.CreateArticle).Methods("POST")	
	myRouter.HandleFunc("/api/article/updatearticle", article.UpdateArticleById).Methods("POST")	
	fmt.Println(http.ListenAndServe(":8081",myRouter))	
}
