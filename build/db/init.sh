CREATE DATABASE set_dev;

create table journal (
	id serial primary key,
	transcation_date timestamp,
	account_number_credit varchar(30),
	account_number_debit varchar(30),
	amount_credit decimal,
	amount_debit decimal,
	type_transaction varchar(1)
);

create table accounts (
	id serial primary key,
	nama varchar(200),
	nik varchar(36),
	no_hp varchar(20),
	pin text,
	nomor_rekening varchar(20),
	created_at timestamp
);