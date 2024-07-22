package api

import (
	"database/sql"
	"example/web-service-gin/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	rows, err := database.Db.Query("SELECT * FROM public.albums ORDER BY id ASC LIMIT 100")
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Couldn't create the new album.")
	} else {
		columnTypes, err := rows.ColumnTypes()

		if err != nil {
			c.AbortWithStatusJSON(400, "Couldn't get col type")
		}

		count := len(columnTypes)
		finalRows := []interface{}{}

		for rows.Next() {

			scanArgs := make([]interface{}, count)

			for i, v := range columnTypes {
				switch v.DatabaseTypeName() {
				case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
					scanArgs[i] = new(sql.NullString)
				case "BOOL":
					scanArgs[i] = new(sql.NullBool)
				case "INT4":
					scanArgs[i] = new(sql.NullInt64)
				default:
					scanArgs[i] = new(sql.NullString)
				}
			}

			err := rows.Scan(scanArgs...)

			if err != nil {
				c.AbortWithStatusJSON(400, "Couldn't scan rows")

			}

			masterData := map[string]interface{}{}

			for i, v := range columnTypes {

				if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
					masterData[v.Name()] = z.Bool
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullString); ok {
					masterData[v.Name()] = z.String
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
					masterData[v.Name()] = z.Int64
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
					masterData[v.Name()] = z.Float64
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
					masterData[v.Name()] = z.Int32
					continue
				}

				masterData[v.Name()] = scanArgs[i]
			}

			finalRows = append(finalRows, masterData)
		}

		c.IndentedJSON(http.StatusOK, finalRows)
	}
}
