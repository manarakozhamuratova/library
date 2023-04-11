CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    sum INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE transactions
ADD CONSTRAINT foreignkey_user_id
FOREIGN KEY (user_id)
REFERENCES users(id);

ALTER TABLE transactions
ADD CONSTRAINT foreignkey_book_id
FOREIGN KEY (book_id)
REFERENCES books(id);