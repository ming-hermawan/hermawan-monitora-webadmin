package cookieswarning

import (
    "net/http"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
)


// CONST

const Lbl = "Cookies Warning"


// PUBLIC

func Get(w http.ResponseWriter) {
    html.GetTmpl0(
      Lbl,
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.CookiesWarningHtmlFilepath,
    )
}

func Process(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
          Get(w)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
