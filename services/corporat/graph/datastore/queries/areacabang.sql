-- name: GetListWilayah :many
select * from ibent.areacabang a order by tanggal_input asc ;

-- name: GetDetilWilayahByKodeWilayah :many
select * from ibent.areacabang a where a.kode_wilayah = $1 ;

-- name: UpdateWilayahByKodeWilayah :many
update ibent.areacabang set kode_wilayah_induk = $1 where kode_wilayah = $2 
RETURNING *;

-- name: DeleteWilayahByKodeWilayah :many
delete from ibent.areacabang where kode_wilayah = $1
RETURNING *;

-- name: CreateNewDataWilayah :one
INSERT INTO ibent.areacabang (kode_wilayah, keterangan, tanggal_input, tanggal_perbarui, terminal_input, terminal_perbarui, user_input, user_perbarui, kode_wilayah_induk)
VALUES(@kode_wilayah::varchar, @keterangan::varchar, @tanggal_input::datetime, @tanggal_perbarui::datetime, @terminal_input::varchar, @terminal_perbarui::varchar, @user_input::varchar, @user_perbarui::varchar, @kode_wilayah_induk::varchar)
RETURNING *;


