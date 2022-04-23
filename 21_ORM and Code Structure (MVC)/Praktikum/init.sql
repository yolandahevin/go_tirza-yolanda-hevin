CREATE TABLE IF NOT EXISTS books(
   id INT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(40),
   author VARCHAR(40)
);

INSERT IGNORE INTO books
    (id, title, author)
VALUES
    (1, 'Tapak Jejak', 'Fujiko F. Fujio'),
    (2, 'The Hobbit', 'J.R.R. Tolkien'),
    (3, 'The Da Vinci Code', 'Dan Brown');
