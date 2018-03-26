package api

import (
	"encoding/json"
	"net/http"
)

// Book type with Name, Author and ISBN
type Book struct {
	//define the book
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// BooksHandleFun to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}

