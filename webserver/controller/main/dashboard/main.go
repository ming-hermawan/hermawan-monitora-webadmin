package dashboard

import (
    "net/http"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/webserver/module/html"
)


func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      "Dashboard",
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.DashboardHtmlFilepath)
}

func Process(username string, w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
          get(username, w)
    }
}
