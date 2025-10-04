package usergroup

import (
    "gorm.io/gorm"
    "net/http"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/decorator"
)


// CONST

const Lbl = "user_groups"


// PRIVATE

func get(db *gorm.DB) ([]string, error) {
    return dbo.GetUserGroupsForDDL(db)
}


// PUBLIC

func Get(db *gorm.DB,
         w http.ResponseWriter) []string {
    return decorator.GetDDL(
      db,
      w,
      Lbl,
      get)
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
