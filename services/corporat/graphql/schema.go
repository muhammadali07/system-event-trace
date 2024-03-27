package graphql

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// var rootQuery = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "RootQuery",
// 		Fields: graphql.Fields{
// 			"hello": &graphql.Field{
// 				Type: graphql.String,
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					return "world", nil
// 				},
// 			},
// 		},
// 	},
// )

// var schema, _ = graphql.NewSchema(
// 	graphql.SchemaConfig{
// 		Query: rootQuery,
// 	},
// )

// func ExecuteQuery(query string) *graphql.Result {
// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 	})

// 	if len(result.Errors) > 0 {
// 		return &graphql.Result{
// 			Errors: result.Errors,
// 		}
// 	}

// 	return result
// }

type Cabang struct {
	KodeCabang   string `db:"kode_cabang"`
	NamaCabang   string `db:"nama_cabang"`
	TipeCabang   string `db:"tipe_cabang"`
	KantorAlamat string `db:"kantor_alamat"`
	// Tambahkan field lain sesuai kebutuhan
}

func GetAllCabang(db *sqlx.DB) ([]Cabang, error) {
	var cabangs []Cabang
	err := db.Select(&cabangs, "SELECT kode_cabang, nama_cabang, tipe_cabang, kantor_alamat FROM ibent.cabang")
	if err != nil {
		return nil, err
	}
	return cabangs, nil
}

func UpdateCabangByCode(db *sqlx.DB, kode_cabang string, nama_cabang string, tipe_cabang string, kantor_alamat string) error {
	query := "UPDATE ibent.cabang SET nama_cabang=$1, tipe_cabang=$2, kantor_alamat=$3 WHERE kode_cabang=$4"
	_, err := db.Exec(query, nama_cabang, tipe_cabang, kantor_alamat, kode_cabang)
	return err
}

func GetCabangByCondition(db *sqlx.DB, condition string) ([]Cabang, error) {
	var cabangs []Cabang
	query := "SELECT * FROM ibent.cabang WHERE " + condition
	fmt.Println(query)
	err := db.Select(&cabangs, query)
	if err != nil {
		return nil, err
	}
	return cabangs, nil
}

func GraphqlHandler(db *sqlx.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Query string `json:"query"`
			// Tambahkan field lain jika ada
		}

		if err := c.BodyParser(&requestBody); err != nil {
			return err
		}

		var result interface{}
		switch {
		case strings.Contains(requestBody.Query, "cabang"):
			cabangs, err := GetAllCabang(db)
			if err != nil {
				return err
			}
			result = cabangs
		case strings.Contains(requestBody.Query, "GetCabangByCondition"):
			var params struct {
				NamaCabang string `json:"nama_cabang"`
			}

			if err := json.Unmarshal([]byte(requestBody.Query), &params); err != nil {
				return err
			}
			condition := fmt.Sprintf("nama_cabang = '%v' ", params.NamaCabang)
			cabangs, err := GetCabangByCondition(db, condition)
			if err != nil {
				return err
			}
			result = cabangs
		case strings.Contains(requestBody.Query, "UpdateCabangByCode"):
			// Parsing parameter yang diperlukan dari requestBody
			var params struct {
				KodeCabang   string `json:"kode_cabang"`
				NamaCabang   string `json:"nama_cabang"`
				TipeCabang   string `json:"tipe_cabang"`
				AlamatKantor string `json:"alamat_kantor"`
			}
			if err := json.Unmarshal([]byte(requestBody.Query), &params); err != nil {
				return err
			}
			err := UpdateCabangByCode(db, params.KodeCabang, params.NamaCabang, params.TipeCabang, params.AlamatKantor)
			if err != nil {
				return err
			}
			result = "Cabang berhasil diupdate"
		default:
			result = "Operasi tidak valid"
		}

		return c.JSON(result)
	}
}
