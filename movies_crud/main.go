package main
import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)

type Movie struct{
 ID string `json:"id"`
 Isbn string `json:"isbn"`
 Tile string `json:"title"`
 Director *Director `json:"director"`

}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	
}

var moives []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type","Application/json")
	json.NewEncoder(w).Encode(moives)
}


func deleteMovie(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content Type","Application/json")
	params:= mux.Vars(r)
	for index, item:=range moives{
		if item.ID==params["id"]{
			moives = append(moives[:index],moives[index+1:]...)
			break
		}

	}

	json.NewEncoder(w).Encode(moives)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type","Application/json")
	params:=mux.Vars(r)
	for _, item:=range moives{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type","Application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(10000))
	moives =append(moives, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	// set json content type
	w.Header().Set("Content Type","Application/json")
	params:=mux.Vars(r)
	// loop over the 
	for index, item:=range moives{
		if item.ID==params["id"]{
			moives = append(moives[:index],moives[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			moives=append(moives, movie)
			json.NewEncoder(w).Encode(movie)

			return
		}

	}


}





func main(){
	r:=mux.NewRouter()
	moives=append(moives, Movie{ID:"1",Isbn: "342323",Tile: "Movie One",Director: &Director{Firstname: "Quan",Lastname: "Pham"}})
	moives=append(moives, Movie{ID:"2",Isbn: "342432",Tile: "Movie Two",Director: &Director{Firstname: "Anh",Lastname: "Pham"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000",r))
	
}