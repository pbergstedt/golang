package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
  "fmt"
  "log"
	"encoding/json"
	"os"
	"strconv"
)

type conditions struct {
  id int "json:id"
	cityname string "json:cityname"
	zipcode string "json:zipcode"
	tempk sql.NullFloat64
	descript string
	humidity sql.NullFloat64
	windspd sql.NullFloat64
	sunrise sql.NullFloat64
	sunset sql.NullFloat64
	ptime sql.NullFloat64
}

func main() {
  db, err := sql.Open("mysql", "root:byteme@tcp(127.0.0.1:3306)/weather")
  if err != nil {
    log.Fatal(err)
  }

  rows, err := db.Query("SELECT * FROM conditions")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  bks := make([]*conditions, 0)
  for rows.Next() {
    bk := new(conditions)
    err := rows.Scan(&bk.id, &bk.cityname, &bk.zipcode, &bk.tempk, &bk.descript, &bk.humidity, &bk.windspd, &bk.sunrise, &bk.sunset, &bk.ptime)
    if err != nil {
      log.Fatal(err)
    }
    bks = append(bks, bk)
	}
  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  for _, bk := range bks {
		var city string = bk.cityname
		var zip string = bk.zipcode
		var temp float64
		// temp = fmt.Sprintf("%.2f", bk.tempk.Float64)
		temp = bk.tempk.Float64
		temp = temp - 273.15
		temp = temp * 1.8
		temp = temp + 32
		// fmt.Printf("%s, %s, %0.2f\n", city, zip, temp)4
		fmt.Printf("%s, %s, %0.1f \n", city, zip, temp)
		temps := strconv.Itoa(temp)
		enc := json.NewEncoder(os.Stdout)
    d := map[string]string{"cityname": city, "zipcode": zip, "tempature": temps}
		// c := map[string]float64{"tempature": temp}
    enc.Encode(d)
	}
}
