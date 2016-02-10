package main

import ("database/sql"
        _ "github.com/go-sql-driver/mysql"
       "log"
       "os"
       "encoding/json"
       "net/http"
       "fmt")
// for database
type Conditions struct {
  id string
  cityname string
  zipcode string
  tempk string
  descript string
  humidity string
  windspd string
  sunrise string
  sunset string
  ptime string
}
//for web json
type Profile struct {
  Cityname string `json:"city"`
  Zipcode string `json:"zipcode"`
  Tempk string `json:"tempature"`
  Descript string `json:"condition"`
  Humidity string `json:"humidity"`
  Windspd string `json:"windspeed"`
  Sunrise string `json:"sunrise"`
  Sunset string `json:"sunset"`
  Ptime string `json:"updated"`
}

func main() {
  http.HandleFunc("/", WeatherData)
  http.ListenAndServe(":3000", nil)
}

func WeatherData(w http.ResponseWriter, r *http.Request) {
  zip := r.URL.Path[1:]
  switch {
  case zip == "45402":
  case zip == "45042":
  case zip == "45036":
  case zip == "45241":
  case zip == "45202":
  case zip == "29901":
  case zip == "89101":
  default:
      fmt.Fprintf(w, "Not a valid zipcode \n")
      fmt.Fprintf(w, "Valid zipcodes are: 45402, 45042, 45036, 45241, 45202, 29901, 89101")
      return
  }
  var usrpwd string = os.Getenv("DB_USER_PWD")
	var dbhost string = os.Getenv("DB_HOST")
  db, err := sql.Open("mysql", usrpwd + "@tcp(" + dbhost + ":3306)/weather")
  if err != nil {
    log.Fatal(err)
  }
  rows, err := db.Query("SELECT * FROM conditions")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  cndts := make([]*Conditions, 0)
  for rows.Next() {
    cndt := new(Conditions)
    err := rows.Scan(&cndt.id, &cndt.cityname, &cndt.zipcode, &cndt.tempk, &cndt.descript, &cndt.humidity, &cndt.windspd, &cndt.sunrise, &cndt.sunset, &cndt.ptime)
    if err != nil {
      log.Fatal(err)
    }
    cndts = append(cndts, cndt)
  }
  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  for _, cndt := range cndts {
    if zip == cndt.zipcode {
      profile := Profile{cndt.cityname, cndt.zipcode, cndt.tempk, cndt.descript, cndt.humidity, cndt.windspd, cndt.sunrise, cndt.sunset, cndt.ptime}
      js, err := json.Marshal(profile)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.Write(js)
    }
  }
}
