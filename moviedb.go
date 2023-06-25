package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	ID          int
	Title       string
	ReleaseDate string
	Duration    int
	Genre       string
	Director    string
	Rating      float64
	PlotSummary string
	PosterURL   string
}

type Actor struct {
	ID          int
	Name        string
	DateOfBirth string
	Nationality string
}

type Director struct {
	Name        string
	DateOfBirth string
	Nationality string
}

type Review struct {
	ID           int
	MovieID      int
	ReviewerName string
	ReviewText   string
	Rating       float64
	ReviewDate   string
}

func getMovies(db *sql.DB) ([]Movie, error) {
	rows, err := db.Query("SELECT * FROM Movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Duration,
			&movie.Genre,
			&movie.Director,
			&movie.Rating,
			&movie.PlotSummary,
			&movie.PosterURL,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func getActors(db *sql.DB) ([]Actor, error) {
	rows, err := db.Query("SELECT * FROM Actors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []Actor
	for rows.Next() {
		var actor Actor
		err := rows.Scan(
			&actor.ID,
			&actor.Name,
			&actor.DateOfBirth,
			&actor.Nationality,
		)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}

func getReviews(db *sql.DB) ([]Review, error) {
	rows, err := db.Query("SELECT * FROM Reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(
			&review.ID,
			&review.MovieID,
			&review.ReviewerName,
			&review.ReviewText,
			&review.Rating,
			&review.ReviewDate,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
