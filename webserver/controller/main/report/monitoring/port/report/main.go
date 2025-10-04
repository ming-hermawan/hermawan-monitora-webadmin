package report

import (
    "fmt"
    "log"
    "encoding/csv"
    "net/http"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/validation"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
)

const Lbl = "Monitoring-Port Report"

// var userData = [][]string{
//   {"id", "name", "email"},
//   {"1", "John Doe", "john.doe@example.com"},
//   {"2", "Jane Smith", "jane.smith@example.com"},
//   {"3", "Peter Jones", "peter.jones@example.com"},
// }

func get(w http.ResponseWriter, r *http.Request, username string) {
    params := r.URL.Query()

    from := params.Get("from")
    to := params.Get("to")

    dateFrom, dateFromValid := validation.DateParamValidation(w, from, "Date-From", true)
    if !dateFromValid {
        return
    }
    dateTo, dateToValid := validation.DateParamValidation(w, to, "Date-To", true)
    if !dateToValid {
        return
    }
    unixTimestampFrom := dateFrom.UnixMicro()
    unixTimestampTo := dateTo.AddDate(0, 0, 1).UnixMicro()

    db := stdlib.GetDb(w)
    histories, err := dbo.GetMPHistory(db, unixTimestampFrom, unixTimestampTo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error read DB: %v", err)
        return
    }

    fileName := fmt.Sprintf("%s_%s.csv", from, to)

    w.Header().Set("Content-Type", "text/csv")
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
    wr := csv.NewWriter(w)

    if err := wr.WriteAll(histories); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error writing CSV to ResponseWriter: %v", err)
        return
    }

    wr.Flush()
    if err := wr.Error(); err != nil {
        log.Printf("Error flushing CSV writer: %v", err)
    }
}

func Process(w http.ResponseWriter, r *http.Request, username string) {
    switch r.Method {
      case "GET":
          get(w, r, username)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
