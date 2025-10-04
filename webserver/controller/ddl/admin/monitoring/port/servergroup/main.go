package servergroup

import (
    "gorm.io/gorm"
    "net/http"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/decorator"
)


// CONST

const Lbl = "server_groups"


// PRIVATE

func get(db *gorm.DB) ([]string, error) {
    return dbo.GetMPServerGroupsForDDL(db)
}

func PostProcess(w http.ResponseWriter,
                 r *http.Request) {
    if (r.Method == "POST") {
        decorator.HttpPostToGetDDL(
          w,
          Lbl,
          get)
    }
}
