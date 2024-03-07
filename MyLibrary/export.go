package calculate

import (
	"bufio"
	"fmt"
	article "goappcards/Controllers/Article"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Document represent the document's data.
type Document struct {
	Title string
	Body  []byte
}

// Save dumps document as txt file on disc.
func (p *Document) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage loads a document from disc.
func loadPage(title string) (*Document, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
			return nil, err
	}
	return &Document{Title: title, Body: body}, nil
}

func Export(){
	reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter Title: ")
        title, err := reader.ReadString('\n')
        if err != nil {
                log.Fatal(err)
        }
        title = strings.TrimSpace(title)

        fmt.Print("Enter Body: ")
        body, err := reader.ReadString('\n')
        if err != nil {
                log.Fatal(err)
        }
        body = strings.TrimSpace(body)

        //Save document and display on command line
        p1 := &Document{Title: title, Body: []byte(body)}
        if err := p1.save(); err != nil {
                log.Fatal(err)
        }
        p2, err := loadPage(title)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(string(p2.Body))
}

func ExportListArticle(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Content-Disposition", "attachment; filename=List.txt")
        listArticle := &Document{Title: "List",Body: (article.ArrListArticle())}
        if err := listArticle.save(); err != nil {
                log.Fatal(err)
        }             
        files, err := ioutil.ReadFile("List.txt")
        if err != nil {
                log.Fatalf("unable to read file: %v", err)
        }
        w.Write(files)
}