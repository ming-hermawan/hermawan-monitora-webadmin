package port

import (
    "net/http"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
)

const Lbl = "Monitoring-Port Report"

func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      "Monitoring-Port Report",
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.ReportMonitoringPortHtmlFilepath,
    )
}

func Process(w http.ResponseWriter, r *http.Request, username string) {
    switch r.Method {
       case "GET":
           get(username, w)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
