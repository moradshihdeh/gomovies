package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Movie struct {
	Title       string
	Description string
	Year        int
}

type PageData struct {
	Title  string
	Movies []Movie
}

var movies []Movie

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("layout.html", "static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error 1", http.StatusInternalServerError)
		fmt.Printf("%v", err)
		return
	}

	data := struct {
		Title  string
		Movies []Movie
		Years  []int
	}{
		Title:  "Movies",
		Movies: nil,
		Years:  generateYears(2010, 2023),
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error 2", http.StatusInternalServerError)
		fmt.Printf("%v", err)
		return
	}
}

func generateYears(start, end int) []int {
	years := make([]int, end-start+1)
	for i := 0; i < len(years); i++ {
		years[i] = start + i
	}
	return years
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "About",
	}

	tmpl, err := template.ParseFiles("layout.html", "static/about.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Contact",
	}

	tmpl, err := template.ParseFiles("layout.html", "static/contact.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func filterMoviesByYear(movies []Movie, year int) []Movie {
	var filteredMovies []Movie

	for _, movie := range movies {
		if movie.Year == year {
			filteredMovies = append(filteredMovies, movie)
		}
	}

	return filteredMovies
}

func moviesByYearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	currentYear := time.Now().Year()
	if year < 2010 || year > currentYear {
		http.Error(w, "Year out of range", http.StatusBadRequest)
		return
	}

	filteredMovies := filterMoviesByYear(movies, year)

	tmpl, err := template.ParseFiles("layout.html", "static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error after parsing", http.StatusInternalServerError)
		fmt.Printf("%v", err)
		return
	}

	data := struct {
		Years  []int
		Year   int
		Movies []Movie
		Title  string
	}{
		Years:  generateYears(2010, time.Now().Year()),
		Year:   year,
		Movies: filteredMovies,
		Title:  "Movies",
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("%v", err)
		return
	}
}

func addMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		description := r.FormValue("description")
		yearStr := r.FormValue("year")

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "Invalid year", http.StatusBadRequest)
			return
		}

		newMovie := Movie{
			Title:       title,
			Description: description,
			Year:        year,
		}

		movies = append(movies, newMovie)

		http.Redirect(w, r, "/?success=true", http.StatusFound)
		return
	}

	tmpl, err := template.ParseFiles("layout.html", "static/add_movie.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Add Movie",
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {

	// TODO: Fetch movies for the selected year from your database or data source
	// For this example, we'll use dummy movie data
	movies = []Movie{
		{Title: "Movie 1", Description: "Description 1", Year: 2012},
		{Title: "Movie 2", Description: "Description 2", Year: 2020},
		{Title: "Movie 3", Description: "Description 3", Year: 2023},
		{Title: "Movie 34", Description: "Description 34", Year: 2023},
		{Title: "Movie 33", Description: "Description 33", Year: 2013},
		{Title: "Movie 23", Description: "Description 23", Year: 2013},
		{Title: "Movie 13", Description: "Description 13", Year: 2015},
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/movies-by-year", moviesByYearHandler)
	http.HandleFunc("/add-movie", addMovieHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
