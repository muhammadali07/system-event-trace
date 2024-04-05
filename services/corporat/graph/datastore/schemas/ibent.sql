-- Create the "ibacc" schema
CREATE SCHEMA IF NOT EXISTS "ibent";

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."aplikasi" (
    "kode_aplikasi" varchar(30) NOT NULL,
    "nama_aplikasi" varchar(100),
    "deskripsi" varchar(100),
    "nama_server" varchar(30),
    "protokol_akses" varchar(15),
    "lokasi_source" varchar(100),
    "lokasi_session" varchar(100),
    "nama_session" varchar(50),
    PRIMARY KEY ("kode_aplikasi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."appsubsystem" (
    "subsystem_code" varchar(5) NOT NULL,
    "subsystem_name" varchar(30),
    "is_local" varchar(1),
    "session_name" varchar(20),
    "remote_host" varchar(23),
    "remote_port" int4,
    "is_ssl" varchar(1),
    "ssl_profile_name" varchar(30),
    "remote_appid" varchar(50),
    "kode_aplikasi" varchar(30),
    PRIMARY KEY ("subsystem_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."areacabang" (
    "kode_wilayah" varchar(3) NOT NULL,
    "keterangan" varchar(50),
    "tanggal_input" timestamp,
    "tanggal_perbarui" timestamp,
    "terminal_input" varchar(15),
    "terminal_perbarui" varchar(15),
    "user_input" varchar(20),
    "user_perbarui" varchar(20),
    "kode_wilayah_induk" varchar(10),
    PRIMARY KEY ("kode_wilayah")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."blobdata" (
    "id_data" int4 NOT NULL,
    "data" text,
    PRIMARY KEY ("id_data")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpparallelstep" (
    "parallel_step_id" int4 NOT NULL,
    "parameter_name" varchar(100),
    "parameter_order" int4,
    "parameter_type" int4,
    "step_id" int4,
    PRIMARY KEY ("parallel_step_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpparallelstepdetail" (
    "pstep_detail_id" int4 NOT NULL,
    "svalue" varchar(500),
    "ivalue" int4,
    "parameter_type" int4,
    "unrestrict_next_step" varchar(1),
    "parallel_step_id" int4,
    "disabled" varchar(1),
    PRIMARY KEY ("pstep_detail_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpparallelstepinstance" (
    "pstepinst_id" int4 NOT NULL,
    "stepinst_id" int4,
    "pstep_detail_id" int4,
    PRIMARY KEY ("pstepinst_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpprocess" (
    "bpprocess_id" int4 NOT NULL,
    "pid" int4,
    "status" int4,
    "main_pid" int4,
    "phost" varchar(24),
    "error_info" varchar(255),
    "step_id" int4,
    "ctlcommand" int4,
    "processed" int4,
    PRIMARY KEY ("bpprocess_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpscenario" (
    "scenario_code" varchar(10) NOT NULL,
    "scenario_name" varchar(30),
    PRIMARY KEY ("scenario_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpscenariocollection" (
    "collection_code" varchar(10) NOT NULL,
    "collection_name" varchar(100),
    PRIMARY KEY ("collection_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpscenarioinstance" (
    "instance_id" int4 NOT NULL,
    "start_time" timestamp,
    "stop_time" timestamp,
    "user_id" varchar(20),
    "terminal_id" varchar(30),
    "status" int4,
    "main_pid" int4,
    "scenario_code" varchar(10),
    "process_date" timestamp,
    PRIMARY KEY ("instance_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpscenariostatus" (
    "id_pod" int4,
    "last_scenario" varchar(10),
    "last_proctime" timestamp,
    "last_status" int4,
    "last_errinfo" varchar(100)
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpscollectionmember" (
    "member_id" int4 NOT NULL,
    "member_order" numeric(1000,53),
    "scenario_code" varchar(10),
    "collection_code" varchar(10),
    PRIMARY KEY ("member_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpscript" (
    "script_id" int4 NOT NULL,
    "description" varchar(100),
    "script_path" varchar(100),
    "kode_aplikasi" varchar(30),
    PRIMARY KEY ("script_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpstep" (
    "step_id" int4 NOT NULL,
    "description" varchar(100),
    "step_order" int4,
    "disabled" varchar(1),
    "scenario_code" varchar(10),
    "script_id" int4,
    PRIMARY KEY ("step_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."bpstepinstance" (
    "stepinst_id" int4 NOT NULL,
    "start_time" timestamp,
    "stop_time" timestamp,
    "status" int4,
    "step_error_message" varchar(255),
    "main_pid" int4,
    "pid" int4,
    "step_id" int4,
    "instance_id" int4,
    PRIMARY KEY ("stepinst_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."cabang" (
    "kode_cabang" varchar(20) NOT NULL,
    "nama_cabang" varchar(100),
    "tipe_cabang" varchar(2),
    "status_cabang" varchar(1),
    "sandi_bi" varchar(7),
    "zona_waktu" varchar(2),
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_otorisasi" varchar(20),
    "tanggal_otorisasi" timestamp,
    "terminal_otorisasi" varchar(15),
    "kode_cabang_induk" varchar(20),
    "status_aktif" varchar(1),
    "status_otorisasi" varchar(1),
    "kode_wilayah" varchar(3),
    "kode_kantor" varchar(5),
    "kode_lokasi" varchar(10),
    "status_akses" varchar(1),
    "kode_wilayah_kliring" varchar(10),
    "level_caban" numeric(1000,53),
    "level_cabang" numeric(1000,53),
    "id_sandilokasi" int4,
    "id_sandibi" int4,
    "kantor_alamat" varchar(200),
    "kantor_kode_pos" varchar(5),
    "kantor_nama" varchar(100),
    "kantor_npwp" varchar(30),
    "kantor_telepon1" varchar(15),
    "kantor_telepon2" varchar(15),
    "kantor_telepon3" varchar(15),
    "kantor_tipe" varchar(2),
    "is_kordinator_kliring" varchar(1),
    "pejabat_cabang" varchar(150),
    "jabatan_pejabat" varchar(150),
    "kantor_kota" varchar(100),
    "status_bds" varchar(1),
    PRIMARY KEY ("kode_cabang")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."cfg_refdata" (
    "ref_type" varchar(50) NOT NULL,
    "ref_code" varchar(20) NOT NULL,
    "ref_desc" varchar(250),
    "status_data" varchar(1),
    "is_leaf" varchar(1),
    "code_slik" varchar(10),
    "desc_slik" varchar(100),
    "code_lsmk" varchar(10),
    "desc_lsmk" varchar(100),
    "code_lbus" varchar(10),
    "desc_lbus" varchar(100),
    "code_anta" varchar(10),
    "desc_anta" varchar(100),
    PRIMARY KEY ("ref_type","ref_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."cfg_reftype" (
    "ref_type" varchar(50) NOT NULL,
    "description" varchar(100) NOT NULL,
    "def_code" varchar(1),
    PRIMARY KEY ("ref_type")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."crewpayroll" (
    "nomor_rekening" varchar(10) NOT NULL,
    "nilai_mutasi" float8,
    "keterangan" varchar(100),
    PRIMARY KEY ("nomor_rekening")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."departemen" (
    "kode_departemen" varchar(10) NOT NULL,
    "nama_departemen" varchar(50),
    "deskripsi" varchar(50),
    "status_aktif" varchar(1),
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_perbarui" varchar(20),
    "tanggal_perbarui" timestamp,
    "terminal_perbarui" varchar(15),
    "tipe_akses" varchar(1),
    PRIMARY KEY ("kode_departemen")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."employee" (
    "nomor_karyawan" varchar(20) NOT NULL,
    "nama_lengkap" varchar(50),
    "alamat" varchar(200),
    "kode_pos" varchar(5),
    "nomor_telepon" varchar(15),
    "nomor_seluler" varchar(15),
    "nama_supervisor" varchar(20),
    "is_supervisor" varchar(1),
    "is_personalia" varchar(1),
    "is_lemburberlaku" varchar(1),
    "nomor_kartu" varchar(18),
    "kode_cabang" varchar(20),
    "kode_departemen" varchar(10),
    "kode_jabatan" varchar(10),
    "status_aktif" varchar(1),
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_perbarui" varchar(20),
    "tanggal_perbarui" timestamp,
    "terminal_perbarui" varchar(15),
    "fingerstation_terminal" varchar(15),
    "kode_kantor" varchar(20),
    PRIMARY KEY ("nomor_karyawan")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."enum_int" (
    "enum_name" varchar(32) NOT NULL,
    "enum_value" int4 NOT NULL,
    "enum_description" varchar(50),
    PRIMARY KEY ("enum_name","enum_value")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."enum_varchar" (
    "enum_name" varchar(32) NOT NULL,
    "enum_value" varchar(2) NOT NULL,
    "enum_description" varchar(50),
    PRIMARY KEY ("enum_name","enum_value")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."executablefungsi" (
    "is_dual_control" varchar(1),
    "level_dual_control" int4,
    "level_override" int4,
    "level_akses" int4,
    "hari_diperbolehkan_akses" varchar(7),
    "waktu_boleh_diakses" timestamp,
    "waktu_selesai_diakses" timestamp,
    "id_fungsi" varchar(10) NOT NULL,
    "restriksi_hari_akses" varchar(1),
    "restriksi_waktu_akses" varchar(1),
    "limit_transaksi_user" float8,
    PRIMARY KEY ("id_fungsi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."fa_ttran" (
    "mstype" varchar(4),
    "nrekgl" varchar(20),
    "nocabg" varchar(3),
    "kdvalt" varchar(3),
    "respct" varchar(4),
    "kdtrex" varchar(2),
    "kdmnem" varchar(1),
    "tgltra" timestamp,
    "norefs" varchar(50),
    "jumtra" numeric(20,2),
    "userid" varchar(12),
    "overid" varchar(12),
    "otorid" varchar(12),
    "kettra" varchar(200),
    "kettr2" varchar(200),
    "nbatch" varchar(30),
    "messid" varchar(100),
    "kodapl" bpchar(2),
    "salawl" numeric(20,2)
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."fingerstation" (
    "fingerstation_id" int4 NOT NULL,
    "kode_cabang" varchar(20),
    "id_terminal" varchar(15),
    PRIMARY KEY ("fingerstation_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."fungsi" (
    "id_fungsi" varchar(10) NOT NULL,
    "nama_fungsi" varchar(60),
    "tipe_fungsi" varchar(1),
    "id_parent_fungsi" varchar(10),
    "kode_aplikasi" varchar(30),
    PRIMARY KEY ("id_fungsi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."glinterface" (
    "kode_interface" varchar(20),
    "tipe_interface" varchar(1),
    "kode_account" varchar(20),
    "kode_rc" varchar(20),
    "deskripsi" varchar(50),
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_perbarui" varchar(20),
    "tanggal_perbarui" timestamp,
    "terminal_perbarui" varchar(15),
    "id_interface" int4 NOT NULL,
    "kode_valuta" varchar(20),
    "kode_cabang" varchar(20),
    PRIMARY KEY ("id_interface")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."historikehadiran" (
    "tanggal" timestamp,
    "id_histori_kehadiran" int4 NOT NULL,
    "waktu_datang" timestamp,
    "waktu_pulang" timestamp,
    "nama_cabang" varchar(30),
    "nama_departemen" varchar(30),
    "nomor_karyawan" varchar(20),
    "ip_datang" varchar(20),
    "ip_pulang" varchar(20),
    PRIMARY KEY ("id_histori_kehadiran")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."id_gen" (
    "id_code" varchar(50) NOT NULL,
    "last_id" int4,
    "locked" varchar(1),
    PRIMARY KEY ("id_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."inbox" (
    "cabang" varchar(20),
    "deskripsi" varchar(100),
    "grup_user" varchar(10),
    "id_inbox" int4 NOT NULL,
    "inputer" varchar(20),
    "is_terproses" varchar(1),
    "level_user" int4,
    "limit" float8,
    "nama_file" varchar(200),
    "status" varchar(1),
    "parameter_code" varchar(10),
    PRIMARY KEY ("id_inbox")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."inboxparameter" (
    "parameter_code" varchar(10) NOT NULL,
    "description" varchar(60),
    "parameter_type" varchar(1),
    "form_script_name" varchar(100),
    "id_fungsi" varchar(10),
    PRIMARY KEY ("parameter_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."jabatan" (
    "kode_jabatan" varchar(10) NOT NULL,
    "nama_jabatan" varchar(30),
    "deskripsi" varchar(50),
    PRIMARY KEY ("kode_jabatan")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."kantor" (
    "kode_kantor" varchar(5) NOT NULL,
    "nama_kantor" varchar(100),
    "tipe_kantor" varchar(2),
    "alamat" varchar(200),
    "nomor_npwp" varchar(30),
    "kode_pos" varchar(5),
    "nomor_telepon1" varchar(15),
    "nomor_telepon2" varchar(15),
    "nomor_telepon3" varchar(15),
    "kode_cabang" varchar(20),
    PRIMARY KEY ("kode_kantor")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."lastobjauditdata" (
    "lastaudit_id" int4 NOT NULL,
    "param_id" int4,
    "datakey_int" int4,
    "datakey_str" varchar(25),
    "audit_id" int4,
    PRIMARY KEY ("lastaudit_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."limitoperasi" (
    "id_limit_operasi" int4 NOT NULL,
    "jenis_limit" varchar(1),
    "nilai_limit" float8,
    "nilai_limit_akumulasi" float8,
    "jumlah_item_limit" int4,
    "status_otorisasi" varchar(1),
    "kode_cabang" varchar(20),
    PRIMARY KEY ("id_limit_operasi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."limittransaksi" (
    "id_limit_transaksi" int4 NOT NULL,
    "jenis_limit" varchar(10),
    "nilai_limit" float8,
    "nilai_limit_akumulasi" float8,
    "jumlah_item_limit" int4,
    "id_user" varchar(20),
    PRIMARY KEY ("id_limit_transaksi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listaplikasiperan" (
    "kode_aplikasi" varchar(30) NOT NULL,
    "id_peran" varchar(30) NOT NULL,
    PRIMARY KEY ("kode_aplikasi","id_peran")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listaplikasiuser" (
    "id_listaplikasiuser" int4 NOT NULL,
    "id_user" varchar(20),
    "kode_aplikasi" varchar(30),
    PRIMARY KEY ("id_listaplikasiuser")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listareadiizinkan" (
    "id_user" varchar(20) NOT NULL,
    "kode_wilayah" varchar(3) NOT NULL,
    PRIMARY KEY ("id_user","kode_wilayah")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listcabangdiizinkan" (
    "id_user" varchar(20) NOT NULL,
    "kode_cabang" varchar(20) NOT NULL,
    PRIMARY KEY ("id_user","kode_cabang")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listdepartemendiizinkan" (
    "nomor_karyawan" varchar(20) NOT NULL,
    "kode_departemen" varchar(10) NOT NULL,
    PRIMARY KEY ("nomor_karyawan","kode_departemen")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listfungsiaplikasiperan" (
    "kode_aplikasi" varchar(30) NOT NULL,
    "id_peran" varchar(30) NOT NULL,
    "id_fungsi" varchar(10) NOT NULL,
    PRIMARY KEY ("kode_aplikasi","id_peran","id_fungsi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listfungsiaplikasiuser" (
    "id_listfungsiaplikasiuser" int4 NOT NULL,
    "id_fungsi" varchar(10),
    "id_listaplikasiuser" int4,
    PRIMARY KEY ("id_listfungsiaplikasiuser")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listperanuser" (
    "id_peran" varchar(30) NOT NULL,
    "id_user" varchar(20) NOT NULL,
    "level_peranuser" int4,
    PRIMARY KEY ("id_peran","id_user")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."listterminaldiizinkan" (
    "id_user" varchar(20) NOT NULL,
    "id_terminal" varchar(15) NOT NULL,
    PRIMARY KEY ("id_user","id_terminal")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."loginsession" (
    "login_id" int4 NOT NULL,
    "id_user" varchar(20),
    "login_time" timestamp,
    "terminal_ip" varchar(25),
    "app_id" varchar(25),
    "session_id" varchar(25),
    "last_access_time" timestamp,
    "branch_code" varchar(20),
    "user_id" varchar(20),
    PRIMARY KEY ("login_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."lokasi" (
    "kode_lokasi" varchar(10) NOT NULL,
    "nama_lokasi" varchar(50),
    "status_data" varchar(1)
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."menu" (
    "menu_id" int4 NOT NULL,
    "menu_name" varchar(100),
    "caption" varchar(100),
    "menu_level" numeric(1000,53),
    "parent_id" int4,
    "is_main_menu" varchar(1),
    "menu_idx" varchar(21),
    "filename_id" varchar(100),
    "is_popup_menu" varchar(1),
    "kode_aplikasi" varchar(30),
    "menu_type" varchar(1),
    PRIMARY KEY ("menu_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."menuhierarchy" (
    "hierarchy_id" int4 NOT NULL,
    "parent_id" int4,
    "child_id" int4,
    PRIMARY KEY ("hierarchy_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."objevent" (
    "event_id" int4 NOT NULL,
    "evtclass_id" int4 NOT NULL,
    "event_code" varchar(20),
    "description" varchar(100),
    PRIMARY KEY ("event_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."objeventclass" (
    "evtclass_id" int4 NOT NULL,
    "classname" varchar(32),
    "description" varchar(100),
    PRIMARY KEY ("evtclass_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."objeventproc" (
    "instance_id" int4 NOT NULL,
    "event_id" int4 NOT NULL,
    "script_id" varchar(20),
    "description" varchar(100),
    "session_id" varchar(32),
    PRIMARY KEY ("instance_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otordata" (
    "id_data" int4 NOT NULL,
    "actualfilename" varchar(50),
    "is_data_packet" varchar(1),
    "clob_data" text,
    PRIMARY KEY ("id_data")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otorentri" (
    "id_otorisasi" int4 NOT NULL,
    "tgl_input" timestamp,
    "kode_entri" varchar(10),
    "user_input" varchar(20),
    "terminal_input" varchar(25),
    "tipe_target" int4,
    "keterangan" varchar(100),
    "info1" varchar(30),
    "info2" varchar(30),
    "id_user" varchar(20),
    "id_peran" varchar(30),
    "kode_cabang" varchar(20),
    "id_data" int4,
    "id_data1" int4,
    "nama_tabel" varchar(50),
    "tipe_key" int4,
    "key_int" int4,
    "key_string" varchar(20),
    "tgl_valuta" timestamp,
    "transaction_amount" numeric(38,16),
    "override_state" varchar(1),
    "user_override" varchar(20),
    "add_info" varchar(100),
    PRIMARY KEY ("id_otorisasi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otorentrihistori" (
    "id_otorisasi" int4 NOT NULL,
    "tgl_input" timestamp,
    "kode_entri" varchar(10),
    "user_input" varchar(20),
    "terminal_input" varchar(25),
    "tipe_target" int4,
    "tgl_otorisasi" timestamp,
    "user_otorisasi" varchar(20),
    "terminal_otorisasi" varchar(20),
    "status_otorisasi" varchar(1),
    "keterangan" varchar(100),
    "info1" varchar(30),
    "info2" varchar(30),
    "id_data" int4,
    "id_user" varchar(20),
    "id_peran" varchar(30),
    "kode_cabang" varchar(20),
    "id_data1" int4,
    "nama_entri" varchar(50),
    "nama_tabel" varchar(50),
    "tipe_key" int4,
    "key_int" int4,
    "key_string" varchar(20),
    "tgl_valuta" timestamp,
    "user_override" varchar(20),
    "add_info" varchar(100),
    PRIMARY KEY ("id_otorisasi")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otorobject" (
    "kode_object" varchar(20) NOT NULL,
    "keterangan" varchar(50),
    "nama_tabel_cari" varchar(50),
    PRIMARY KEY ("kode_object")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otorparameter" (
    "kode_entri" varchar(10) NOT NULL,
    "remote_form_id" varchar(100),
    "method_id_save" varchar(50),
    "subsystem_code" varchar(5),
    "method_id_reject" varchar(50),
    "nama_entri" varchar(50),
    "classic_form_call" varchar(1),
    "kode_object" varchar(20),
    PRIMARY KEY ("kode_entri")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otorparameterperan" (
    "id_otorparameterperan" int4 NOT NULL,
    "id_peran" varchar(30),
    "kode_entri" varchar(10),
    "status" varchar(1),
    "terminal_input" varchar(15),
    "terminal_perbarui" varchar(15),
    "user_input" varchar(20),
    "user_perbarui" varchar(20),
    "tanggal_input" timestamp,
    "tanggal_perbarui" timestamp,
    PRIMARY KEY ("id_otorparameterperan")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."otpkey" (
    "userid" varchar(20) NOT NULL,
    "token_base" varchar(16),
    "t_challenge" timestamp,
    "key1" int4,
    "key2" int4,
    "key3" int4,
    "key4" int4,
    "key5" int4,
    PRIMARY KEY ("userid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."parameterglobal" (
    "kode_parameter" varchar(20) NOT NULL,
    "tipe_parameter" varchar(1),
    "nilai_parameter" float8,
    "deskripsi" varchar(60),
    "is_parameter_system" varchar(1),
    "nilai_parameter_tanggal" timestamp,
    "nilai_parameter_string" varchar(30),
    "status" varchar(1),
    PRIMARY KEY ("kode_parameter")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."peran" (
    "id_peran" varchar(30) NOT NULL,
    "nama_peran" varchar(50),
    "deskripsi" varchar(50),
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_perbarui" varchar(20),
    "tanggal_perbarui" timestamp,
    "terminal_perbarui" varchar(15),
    "status_peran" varchar(1),
    "kategori_peran" varchar(1),
    PRIMARY KEY ("id_peran")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."publication" (
    "publication_code" varchar(5) NOT NULL,
    "description" varchar(100),
    PRIMARY KEY ("publication_code")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ibent.reference_limit_id_seq;

-- Table Definition
CREATE TABLE "ibent"."reference_limit" (
    "id" int4 NOT NULL DEFAULT nextval('ibent.reference_limit_id_seq'::regclass),
    "group_limit" varchar(20),
    "jenis_limit" varchar(20),
    "nilai_limit" float8,
    "nilai_limit_akumulasi" float8,
    "jumlah_item_limit" int4,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ibent.reference_map_limit_id_seq;

-- Table Definition
CREATE TABLE "ibent"."reference_map_limit" (
    "id" int4 NOT NULL DEFAULT nextval('ibent.reference_map_limit_id_seq'::regclass),
    "rolecode_hrmis" varchar(20),
    "group_limit" varchar(20),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ibent.reference_map_role_id_seq;

-- Table Definition
CREATE TABLE "ibent"."reference_map_role" (
    "id" int4 NOT NULL DEFAULT nextval('ibent.reference_map_role_id_seq'::regclass),
    "rolecode_hrmis" varchar(20),
    "rolecode_ihsan" varchar(20),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."referencedata" (
    "refdata_id" int4 NOT NULL,
    "reftype_id" int4 NOT NULL,
    "reference_code" varchar(50) NOT NULL,
    "reference_desc" varchar(250),
    "parentrefdata_id" int4,
    "is_leaf" varchar(1),
    "reference_level" int4,
    "status_data" varchar(1),
    PRIMARY KEY ("refdata_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."referencehierarchy" (
    "refhierarchy_id" int4 NOT NULL,
    "parentrefdata_id" int4 NOT NULL,
    "childrefdata_id" int4 NOT NULL,
    PRIMARY KEY ("refhierarchy_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."referencemap" (
    "refmap_id" int4 NOT NULL,
    "refmap_type" varchar(1),
    "refmapdef_id" int4 NOT NULL,
    "destrefdata_id" int4 NOT NULL,
    "srcrefdata_id" int4,
    "intkey_src" int4,
    "strkey_src" varchar(25),
    PRIMARY KEY ("refmap_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."referencemapdef" (
    "refmapdef_id" int4 NOT NULL,
    "refmaptype" varchar(1) NOT NULL,
    "description" varchar(50),
    "destreftype_id" int4 NOT NULL,
    "srcreftype_id" int4,
    "src_classname" varchar(25),
    "src_keyname" varchar(25),
    "src_keytype" int4,
    PRIMARY KEY ("refmapdef_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."referencetype" (
    "reftype_id" int4 NOT NULL,
    "reference_name" varchar(50),
    "description" varchar(100),
    "kode_map_kategori" varchar(3),
    PRIMARY KEY ("reftype_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."refmapitem" (
    "kode_kategori" varchar(3) NOT NULL,
    "sandi_lbus" varchar(10) NOT NULL,
    "sandi_lsmk" varchar(10) NOT NULL,
    "sandi_sid" varchar(10) NOT NULL,
    "status_data" varchar(1),
    PRIMARY KEY ("kode_kategori","sandi_lbus","sandi_lsmk","sandi_sid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."refmapkategori" (
    "kode_kategori" varchar(3) NOT NULL,
    "nama_kategori" varchar(50),
    "kolom_detil" varchar(10),
    PRIMARY KEY ("kode_kategori")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."refmapreport" (
    "kode_map_kategori" varchar(3),
    "reftype_id" int4 NOT NULL,
    "report_type" varchar(1),
    PRIMARY KEY ("reftype_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."rolemenulist" (
    "list_id" int4 NOT NULL,
    "id_peran" varchar(30),
    "menu_id" int4,
    PRIMARY KEY ("list_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."salcrewpayroll" (
    "nomor_rekening" varchar(10) NOT NULL,
    "salsek" float8,
    "salseh" float8,
    PRIMARY KEY ("nomor_rekening")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."subscribeitem" (
    "id_sitem" int4 NOT NULL,
    "script_path" varchar(200),
    "kode_aplikasi" varchar(30),
    "publication_code" varchar(5),
    PRIMARY KEY ("id_sitem")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."teller" (
    "kode_rekening_teller" varchar(7),
    "id_user" varchar(20) NOT NULL,
    PRIMARY KEY ("id_user")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."terminal" (
    "id_terminal" varchar(15) NOT NULL,
    "nama_terminal" varchar(30),
    "lokasi_terminal" varchar(15),
    "deskripsi" varchar(50),
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_perbarui" varchar(20),
    "tanggal_perbarui" timestamp,
    "terminal_perbarui" varchar(15),
    "kode_kantor" varchar(5),
    "kode_cabang" varchar(20),
    PRIMARY KEY ("id_terminal")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."userapp" (
    "id_user" varchar(20) NOT NULL,
    "password" varchar(20),
    "status_profil" varchar(1),
    "level_user" int4,
    "tanggal_didaftarkan" timestamp,
    "tanggal_kadaluarsa" timestamp,
    "tanggal_password_kadaluarsa" timestamp,
    "tanggal_ubah_password_terakhir" timestamp,
    "tanggal_aktif_terakhir" timestamp,
    "tipe_akses_cabang" varchar(1),
    "tipe_akses_terminal" varchar(1),
    "hari_diperbolehkan_akses" varchar(7),
    "status_restriksi_waktu" varchar(1),
    "waktu_boleh_mulai" timestamp,
    "waktu_harus_selesai" timestamp,
    "maksimum_jumlah_login" int4,
    "user_input" varchar(20),
    "tanggal_input" timestamp,
    "terminal_input" varchar(15),
    "user_perbarui" varchar(20),
    "tanggal_perbarui" timestamp,
    "terminal_perbarui" varchar(15),
    "tipe_user" varchar(1),
    "jumlah_login_gagal" int4,
    "jumlah_login_aktif" int4,
    "kunci_akses" varchar(1),
    "nama_user" varchar(50),
    "jumlah_validasi_gagal" int4,
    "status_password" int4,
    "maksimum_tidak_aktif" int4,
    "maksimum_ganti_password" int4,
    "no_urut_teller" varchar(3),
    "id_terminal" varchar(15),
    "kode_cabang" varchar(20),
    "nomor_karyawan" varchar(20),
    "maksimum_gagal_login" numeric(1000,53),
    "status_akses" varchar(1),
    "tipe_login" varchar(1),
    "fingerstation_terminal" varchar(15),
    "kode_kantor" varchar(10),
    PRIMARY KEY ("id_user")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."useraudit" (
    "audit_id" int4 NOT NULL,
    "user_id" varchar(20) NOT NULL,
    "app_id" varchar(20),
    "terminal_ip" varchar(15),
    "operation_code" varchar(32),
    "event_time" timestamp,
    "event_description" varchar(250),
    "info1" varchar(100),
    "info2" varchar(100),
    PRIMARY KEY ("audit_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."useraudithistori" (
    "audit_id" int4 NOT NULL,
    "user_id" varchar(20) NOT NULL,
    "app_id" varchar(20),
    "terminal_ip" varchar(15),
    "operation_code" varchar(32),
    "event_time" timestamp,
    "event_description" varchar(250),
    "info1" varchar(100),
    "info2" varchar(100),
    PRIMARY KEY ("audit_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."userbiodata" (
    "id_biodata" int4 NOT NULL,
    "id_data" int4 NOT NULL,
    "id_user" varchar(20),
    "data_tag" varchar(20) NOT NULL,
    "description" varchar(100),
    "nomor_karyawan" varchar(20),
    PRIMARY KEY ("id_biodata")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."usermenulist" (
    "list_id" int4 NOT NULL,
    "id_user" varchar(30),
    "menu_id" int4,
    PRIMARY KEY ("list_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."usersession" (
    "id_session" int4 NOT NULL,
    "session_stamp" varchar(32),
    "id_user" varchar(20),
    "login_time" timestamp,
    "logout_time" timestamp,
    "last_access_time" timestamp,
    "status" varchar(1),
    "ip_user" varchar(25),
    "app_id" varchar(32),
    PRIMARY KEY ("id_session")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."usertbl" (
    "userid" varchar(20) NOT NULL,
    "enc_password1" varchar(16),
    "enc_password2" varchar(16),
    PRIMARY KEY ("userid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "ibent"."wilayahkliring" (
    "kode_wilayah_kliring" varchar(5) NOT NULL,
    "nama_wilayah_kliring" varchar(50),
    "keterangan" varchar(100),
    "hari_efektif_kliring" numeric(1000,53),
    PRIMARY KEY ("kode_wilayah_kliring")
);