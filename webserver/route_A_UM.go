package webserver

import (
    "strconv"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/admin/usermanagement/usergroup"
    "hermawan-monitora/webserver/controller/main/admin/usermanagement/user"
    ddlUserGroup "hermawan-monitora/webserver/controller/ddl/admin/usermanagement/usergroup"
)


func adminUserGroupProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admUsrMgtUsrGrpRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmUsrMgtUsrGrp))
    admUsrMgtUsrGrp, _ := strconv.Atoi(admUsrMgtUsrGrpRaw)
    if admUsrMgtUsrGrp != 1 {
        login.Get(w)
        return
    }
    usergroup.Process(username, w, r)
}

func adminUserProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admUsrMgtUsrRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmUsrMgtUsr))
    admUsrMgtUsr, _ := strconv.Atoi(admUsrMgtUsrRaw)
    if admUsrMgtUsr != 1 {
        login.Get(w)
        return
    }
    user.Process(username, w, r)
}

func ddlUserGroupProcess(w http.ResponseWriter, r *http.Request) {
    _, err := authProcess(r)
    if err != nil {
        httpresponse.ErrResponseForInvalidToken(w, err.Error())
        return
    }
    ddlUserGroup.PostProcess(w, r)
}
