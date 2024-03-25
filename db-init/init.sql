CREATE TABLE albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    artist VARCHAR(100),
    price FLOAT
);

INSERT INTO albums (title, artist, price) VALUES
    ("Blue Train", "John Coltrane", 56.99),
	("Jeru", "Gerry Mulligan", 17.99),
	("Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99);