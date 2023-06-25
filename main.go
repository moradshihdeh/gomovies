package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomeData struct {
	Title   string
	Content string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	t, err := template.ParseFiles("static/layout.html", tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Home Page",
	}

	renderTemplate(w, "static/home.html", data)
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Movies []Movie
	}{
		Title:  "Movies List Page",
		Movies: nil,
	}
	renderTemplate(w, "static/movies.html", data)
}

func directorsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title     string
		Directors []struct {
			Name        string
			DateOfBirth string
			Nationality string
		}
	}{
		Title: "Director List",
		Directors: []struct {
			Name        string
			DateOfBirth string
			Nationality string
		}{
			{
				Name:        "Christopher Nolan",
				DateOfBirth: "July 30, 1970",
				Nationality: "British",
			},
			{
				Name:        "Quentin Tarantino",
				DateOfBirth: "March 27, 1963",
				Nationality: "American",
			},
			// Add more directors as needed
		},
	}
	renderTemplate(w, "static/directors.html", data)
}

func actorsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Actors []Actor
	}{
		Title: "Actor List",
		Actors: []Actor{
			{
				Name:        "Leonardo DiCaprio",
				DateOfBirth: "November 11, 1974",
				Nationality: "American",
			},
			{
				Name:        "Tom Hanks",
				DateOfBirth: "July 9, 1956",
				Nationality: "American",
			},
			// Add more actors as needed
		},
	}

	renderTemplate(w, "static/actors.html", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Content string
	}{
		Title:   "About",
		Content: "This is a web page about movies.",
	}

	renderTemplate(w, "static/about.html", data)
}

func main() {
	/*
		// Connect to the MySQL database
		db, err := sql.Open("mysql", "bkdr:n3t4cc3ss@tcp(192.168.0.102:3306)/movies")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Test the database connection
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to the database!")

		// Load all tables' results into arrays
		movies, err := getMovies(db)
		if err != nil {
			log.Fatal(err)
		}

		actors, err := getActors(db)
		if err != nil {
			log.Fatal(err)
		}

		reviews, err := getReviews(db)
		if err != nil {
			log.Fatal(err)
		}

		// Print the loaded data
		fmt.Println("Movies:")
		for _, movie := range movies {
			fmt.Printf("Title: %s, Director: %s\n", movie.Title, movie.Director)
		}

		fmt.Println("Actors:")
		for _, actor := range actors {
			fmt.Printf("Name: %s, Nationality: %s\n", actor.Name, actor.Nationality)
		}

		fmt.Println("Reviews:")
		for _, review := range reviews {
			fmt.Printf("Reviewer: %s, Rating: %.1f\n", review.ReviewerName, review.Rating)
		}
	*/
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/movies", moviesHandler)
	http.HandleFunc("/directors", directorsHandler)
	http.HandleFunc("/actors", actorsHandler)
	http.HandleFunc("/about", aboutHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}
