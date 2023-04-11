alter table transactions drop CONSTRAINT foreignkey_user_id;
alter table transactions drop CONSTRAINT foreignkey_book_id;
DROP TABLE transactions;