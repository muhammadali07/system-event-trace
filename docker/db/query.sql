
create table journal (
	id serial primary key,
	tanggal_transaksi timestamp,
	no_rekening_kredit varchar(30),
	no_rekening_debit varchar(30),
	nominal_kredit decimal,
	nominal_debit decimal
);

create table accounts (
	id serial primary key,
	nama varchar(200),
	nik varchar(36),
	np_hp varchar(20),
	pin text,
	create_at timestamp
);