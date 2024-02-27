package controllers
import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/msrpr20/pkg/project/utils"
	"github.com/msrpr20/pkg/project/models"

)

var NewBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks:= models.GetAllBooks()
	res, _ := 	json.Marshal(newBooks)
	w.Header().Set("Content-Type","publication/json")
	w.WriteHeader(https.Status.OK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter ,r *http.Request){
	var mux.Vars(r)
	bookId := vars["bookId"]
	Id ,err :=strconv.ParseInt(bookId.,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	res, _:= json.Marshall(bookDetails)
	w.Header().Set("Content-Type","publication/json")
	w.WriteHeader(https.Status.OK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter,r *http.Request){
    CreateBook:= &models.Book{}
	utils.Parsebody(r,CreateBook)
	b:= CreatBook.CreatBook()
	res, _:= json.Marshall(b)
	w.WriteHeader(https.Status.OK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter ,r *http.Request){
	var mux.Vars(r)
	bookId := vars["bookId"]
	Id ,err :=strconv.ParseInt(bookId.,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(Id)
	res, _:= json.Marshall(book)
	w.Header().Set("Content-Type","publication/json")
	w.WriteHeader(https.Status.OK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter,r *http.Request){
    CreateBook:= &models.Book{}
	utils.Parsebody(r,updatebook)
	var mux.Vars(r) 
	bookId : Vars["bookId"]
	Id,err :=strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
    bookDetail, db:= models.GetBookByID(Id)
	if updateBook.Name != ""{
		bookDetails.name = updateBook.name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	res, _:= json.Marshall(bookDetails)
	w.Header().Set("Content-Type","publication/json")
	w.WriteHeader(https.Status.OK)
	w.Write(res)
}