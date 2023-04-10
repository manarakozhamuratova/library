alter table book_borrows drop CONSTRAINT fk_user_id;
alter table book_borrows drop CONSTRAINT fk_book_id;
drop table books;
drop table book_borrows;