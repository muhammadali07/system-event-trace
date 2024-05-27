create database set_dev;

create table journal (
	id serial primary key,
	transaction_date timestamp,
	account_number_credit varchar(30),
	account_number_debit varchar(30),
	amount_credit decimal,
	amount_debit decimal,
	type_transaction varchar(1),
	created_at timestamp
);

create table accounts (
	id serial primary key,
	name varchar(200),
	nik varchar(36),
	phone_number varchar(20),
	pin text,
	account_number varchar(20),
	balance numeric(28,2),
	status varchar(1),
	created_at timestamp
);
