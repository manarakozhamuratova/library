CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  "created_at" timestamp default now()
);

CREATE TABLE book_borrows (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  book_id INTEGER NOT NULL,
  borrow_date timestamp default now(),
  return_date timestamp
);

ALTER TABLE book_borrows
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id)
REFERENCES users(id);

ALTER TABLE book_borrows
ADD CONSTRAINT fk_book_id
FOREIGN KEY (book_id)
REFERENCES books(id);

