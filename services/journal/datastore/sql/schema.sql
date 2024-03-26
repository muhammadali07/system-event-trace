create table journal (
	id serial primary key,
	transcation_date timestamp,
	account_number_credit varchar(30),
	account_number_debit varchar(30),
	amount_credit decimal,
	amount_debit decimal,
	type_transaction varchar(1),
	created_at timestamp
);