package datastore

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/log"
)

type JournalDatabase struct {
	log         *log.Logger
	driver      string
	db          *pgxpool.Pool
	core_schema string
}

// func (f *JournalDatabase) Begin() (tx *pgx.Conn, err error) {
// 	// Gunakan `pool.Begin` untuk memulai transaksi
// 	tx, err = f.db.pool.Begin(context.Background())
// 	if err != nil {
// 		remark := "ds: gagal memulai transaksi"
// 		f.log.Error(logrus.Fields{
// 			"error": err.Error(), // Gunakan pesan error original
// 		}, nil, remark)
// 		return // Kembalikan error original
// 	}
// 	return tx, nil
// }

func InitDatastore(driver, host, user, password, database string, port int, map_schema map[string]string, log *log.Logger) *JournalDatabase {
	// Buat konfigurasi koneksi
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	config := fmt.Sprintf("%v://%v:%v@%v:%v/%v", driver, user, password, host, port, database)

	// Buat koneksi ke database
	conn, err := pgx.Connect(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		panic(err)
	}
	defer conn.Close(context.Background())

	return &JournalDatabase{
		log:         log,
		driver:      driver,
		db:          conn,
		core_schema: map_schema["public"],
	}
}
