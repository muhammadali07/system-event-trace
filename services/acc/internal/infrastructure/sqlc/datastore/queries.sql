-- queries.sql

-- Buat akun baru
INSERT INTO akun (nama, nik, no_hp, pin, saldo)
VALUES ($1, $2, $3, $4, $5)
RETURNING no_rekening, nama, nik, no_hp, pin, saldo;

-- Dapatkan akun berdasarkan no_rekening
SELECT no_rekening, nama, nik, no_hp, pin, saldo
FROM akun
WHERE no_rekening = $1;

-- Perbarui saldo akun
UPDATE akun
SET saldo = $1
WHERE no_rekening = $2;
