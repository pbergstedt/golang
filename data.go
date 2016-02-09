package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
  "log"
	"encoding/json"
	"os"
	"strconv"
)

type conditions struct {
  id int
	cityname string
	zipcode string
	tempk sql.NullFloat64
	descript string
	humidity sql.NullFloat64
	windspd sql.NullFloat64
	sunrise string
	sunset string
	ptime string
}

func main() {
	var zip string = os.Getenv("ZIP")
	var usrpwd string = os.Getenv("DB_USER_PWD")
  db, err := sql.Open("mysql", usrpwd + "@tcp(127.0.0.1:3306)/weather")
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
	  if zip == bk.zipcode {
		  var temp float64 = bk.tempk.Float64
      //convert temp from Kelvin to Fahrenheit
			// temp = ((temp - 273.15) * 1.8) + 32
	    // fmt.Printf("%s, %s, %0.1f \n", city, zip, temp)
		  var stemp string = strconv.FormatFloat(float64(temp), 'f', 2, 32)
      var hum float64 = bk.humidity.Float64
		  var shumidity string = strconv.FormatFloat(float64(hum), 'f', 1, 32)
      var ws float64 = bk.windspd.Float64
		  var swindspd string = strconv.FormatFloat(float64(ws), 'f', 1, 32)
		  enc := json.NewEncoder(os.Stdout)
      d := map[string]string{"data": "weather", "city": bk.cityname,
			                       "zipcode": bk.zipcode, "tempature": stemp,
			  										 "condition": bk.descript, "updated": bk.ptime,
			  									   "humidity": shumidity, "windspeed": swindspd,
			  									   "sunrise": bk.sunrise, "sunset": bk.sunset}
      enc.Encode(d)
		}
	}
}
