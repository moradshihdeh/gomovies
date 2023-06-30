-- Create Movies table
CREATE TABLE Movies (
  movie_id INT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  release_date DATE,
  duration INT,
  genre VARCHAR(255),
  director VARCHAR(255),
  rating DECIMAL(3, 1),
  plot_summary TEXT,
  poster_url VARCHAR(255)
);

-- Create Actors table
CREATE TABLE Actors (
  actor_id INT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  date_of_birth DATE,
  nationality VARCHAR(255)
);

-- Create Movies_Actors table
CREATE TABLE Movies_Actors (
  movie_id INT,
  actor_id INT,
  PRIMARY KEY (movie_id, actor_id),
  FOREIGN KEY (movie_id) REFERENCES Movies(movie_id),
  FOREIGN KEY (actor_id) REFERENCES Actors(actor_id)
);

-- Create Reviews table
CREATE TABLE Reviews (
  review_id INT PRIMARY KEY,
  movie_id INT,
  reviewer_name VARCHAR(255),
  review_text TEXT,
  rating DECIMAL(2, 1),
  review_date DATE,
  FOREIGN KEY (movie_id) REFERENCES Movies(movie_id)
);

CREATE TABLE directors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    nationality VARCHAR(255) NOT NULL
);

-- Drop Movies_Actors table
DROP TABLE IF EXISTS Movies_Actors;

-- Drop Reviews table
DROP TABLE IF EXISTS Reviews;

-- Drop Actors table
DROP TABLE IF EXISTS Actors;

-- Drop Movies table
DROP TABLE IF EXISTS Movies;


INSERT INTO directors (name, date_of_birth, nationality)
VALUES
    ('Christopher Nolan', '1970-07-30', 'British'),
    ('Quentin Tarantino', '1963-03-27', 'American'),
    ('Steven Spielberg', '1946-12-18', 'American'),
    ('Martin Scorsese', '1942-11-17', 'American'),
    ('Stanley Kubrick', '1928-07-26', 'American'),
    ('Alfred Hitchcock', '1899-08-13', 'British'),
    ('Hayao Miyazaki', '1941-01-05', 'Japanese'),
    ('Spike Lee', '1957-03-20', 'American'),
    ('Francis Ford Coppola', '1939-04-07', 'American'),
    ('Wong Kar-wai', '1958-07-17', 'Chinese'),
    ('Denis Villeneuve', '1967-10-03', 'Canadian'),
    ('David Fincher', '1962-08-28', 'American');
