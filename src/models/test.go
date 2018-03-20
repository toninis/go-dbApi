package models

import (
  "log"
  _"fmt"
)

type DataStruct struct {
    Id   int `json:"id" db:"id"`
    Project string `json:"project" db:"project"`
    Msisdn int `json:"msisdn" db:"msisdn"`
    Date_fetched string `json:"date_fetched" db:"date_fetched"`
    First_date_bulked string `json:"first_date_bulked" db:"first_date_bulked"`
    Last_date_bulked string `json:"last_date_bulked" db:"last_date_bulked"`
    Total_bulks int `json:"total_bulks" db:"total_bulks"`
  }

type Row struct {
	Results []DataStruct
}

func ResultData(data *Row) error {
  rows, err := db.Queryx("SELECT * FROM public.user_base")
  for rows.Next() {
      var row DataStruct
      err := rows.StructScan(&row)
      if err != nil {
          log.Fatalln(err)
      }
      // fmt.Printf("%#v\n", row)
      data.Results = append(data.Results,row)
  }
  return err
}
