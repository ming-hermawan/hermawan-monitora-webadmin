package uploadcsv

import (
    "log"
    "fmt"
    "io"
    "strconv"
    "strings"
    "encoding/csv"
    "encoding/json"
    "net/http"
    "github.com/gocarina/gocsv"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/initredis"
    "hermawan-monitora/webserver/module/stdlib"
)

const Lbl = "Upload from CSV"

type Record struct {
  IP      string `csv:"ip"`
  Name    string `csv:"name"`
  Group   string `csv:"group"`
  Service string `csv:"service"`
  Port    int    `csv:"port"`
}

type Server struct {
  IP string
  Name string
  Group string
}

type Port struct {
  IP string
  Port int
  Name string
}

func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      Lbl,
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.AdmMonPortUploadCsvHtmlFilepath,
    )
}

func formatInfo(src string, desc string, sts string) map[string]string {
    return map[string]string {
      "src": src,
      "desc": desc,
      "sts": sts}
}

func processSubmit(w http.ResponseWriter, r *http.Request) {
    // Maximum upload of 10 MB files
    r.ParseMultipartForm(10 << 20)

    // Get handler for filename, size and headers
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }

    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
        r := csv.NewReader(in)
        r.Comma = ';'
        return r
    })

    var records []Record
    if err := gocsv.UnmarshalMultipartFile(&file, &records); err != nil {
       panic(err)
    }

    serverGroupStates := []map[string]any {}
    serverStates := []map[string]any {}
    portStates := []map[string]any {}
    info := []map[string]string {}
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    for _, record := range records {
        isServerGroupFound := false;
        for _, x := range serverGroupStates {
            if record.Group == x["group"] {
                isServerGroupFound = true
                break
            }
        }
        if !isServerGroupFound {
            isServerGroupExists, _ := dbo.IsMPServerGroupExists(db, record.Group)
            serverGroupState := map[string]any {
              "group": record.Group,
              "exists": isServerGroupExists}
            serverGroupStates = append(serverGroupStates, serverGroupState)
            sts := "New Insert"
            if isServerGroupExists {
                sts = "Already Exist"
            }
            info = append(info, formatInfo("Server-Group", record.Group, sts))
        }
        isServerFound := false;
        for _, x := range serverStates {
            if record.IP == x["ip"] {
                isServerFound = true
                if record.Name != x["server"] {
                    err := fmt.Sprintf(
                      "IP %s has 2 different name, %s and %s\n",
                      record.IP,
                      record.Name,
                      x["server"])
                    http.Error(w, err, http.StatusInternalServerError)
                    return
                } else {
                    break
                }
            }
        }
        if !isServerFound {
            isServerExists, _ := dbo.IsServerExists(record.IP)
            serverState := map[string]any {
              "ip": record.IP,
              "server": record.Name,
              "group": record.Group,
              "exists": isServerExists}
            serverStates = append(serverStates, serverState)
            sts := "New Insert"
            if isServerExists {
                sts = "Update"
            }
            info = append(info, formatInfo("Server", record.Name, sts))
        }
        log.Println(fmt.Sprintf("Ports:\n%v\n", portStates))
        for _, x := range portStates {
            if (record.IP == x["ip"]) && (record.Port == x["port"]) {
                err := fmt.Sprintf(
                  "Ip %s Port %d duplicate!\n",
		  record.IP,
                  record.Port)
                http.Error(w, err, http.StatusInternalServerError)
                return
            }
        }
        isPortExists, _ := dbo.IsPortExists(record.IP, record.Port)
        portState := map[string]any {
          "ip": record.IP,
          "port": record.Port,
          "name": record.Service,
          "exists": isPortExists}
        sts := "New Insert"
        if isPortExists {
            sts = "Update"
        }
        info = append(info, formatInfo("Port", strconv.Itoa(record.Port), sts))
        portStates = append(portStates, portState)
    }
    result := map[string]any {
      "group": serverGroupStates,
      "server": serverStates,
      "port": portStates}
    var success bool
    redisKey := hmonredis.GetMonPortUploadCSVKey()
    success = initredis.Init(w)
    if !success {
        return
    }
    success = initredis.SendInitSignal(w)
    if !success {
        return
    }
    jsonInBytes, err := json.Marshal(&result)
    err = hmonredis.SetRawWithExpired(
      redisKey,
      jsonInBytes)
    if err != nil {
        httpresponse.ErrResponseSetRedis(
          w,
          redisKey,
          err.Error())
        return
    }
    jsonInBytes2, _ := json.Marshal(
      map[string]any {
        "id": redisKey,
        "info": info})
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes2)
}

func processConfirm(w http.ResponseWriter, r *http.Request) {
    var params map[string]interface{}
    json.NewDecoder(r.Body).Decode(&params);
    id, idFound := params["id"].(string)
    if idFound {
    }
    jsonInBytes2, _ := json.Marshal(map[string]any {"message": "SUCCESS", "id": id})
    var err error
    var jsonStr string
    jsonStr, err = hmonredis.Get(id)
    if err != nil {
        httpresponse.ErrResponseGetRedis(w, id, err.Error())
        return
    }
    var reading map[string]interface{}
    if err := json.Unmarshal([]byte(jsonStr), &reading); err != nil {
        httpresponse.ErrResponseWhenCnvToJson(w, err.Error())
        return
    }
    serverGroupStates := reading["group"].([]interface{})
    serverStates := reading["server"].([]interface{})
    portStates := reading["port"].([]interface{})

    var newMPServerGroupList []dbo.MPServerGroup
    var newMPServerList []dbo.MPServer
    var newMPPortList []dbo.MPPort
    var updServerList []Server
    var updPortList []Port
    for _, x := range serverGroupStates {
	group := x.(map[string]any)["group"].(string)
        if !x.(map[string]any)["exists"].(bool) {
            temp := dbo.MPServerGroup {ServerGroup: group, SortNumber: 99}
            newMPServerGroupList = append(newMPServerGroupList, temp)
        }
    }
    for _, x := range serverStates {
        temp := x.(map[string]any)
        exists := temp["exists"].(bool)
        ip := temp["ip"].(string)
        name := temp["server"].(string)
        group := temp["group"].(string)
        if exists {
            temp := Server{IP: ip, Name: name, Group: group}
            updServerList = append(updServerList, temp)
        } else {
            temp := dbo.MPServer {
              Ip: ip,
              Name: name,
              ServerGroup: group}
            newMPServerList = append(newMPServerList, temp)
        }
    }
    for _, x := range portStates {
        temp := x.(map[string]any)
        exists := temp["exists"].(bool)
        ip := temp["ip"].(string)
        port := int(temp["port"].(float64))
        name := temp["name"].(string)
        if exists {
            temp := Port{IP: ip, Port: port, Name: name}
            updPortList = append(updPortList, temp)
        } else {
            temp := dbo.MPPort {
              Ip: ip,
              Port: port,
              Name: name}
            newMPPortList = append(newMPPortList, temp)
        }
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    tx := db.Begin()
    dbo.BulkInsMPServerGroup(tx, newMPServerGroupList)
    dbo.BulkInsServer(tx, newMPServerList)
    dbo.BulkInsPort(tx, newMPPortList)
    for _, x := range updServerList {
        err = dbo.UpdServer(tx, x.IP, x.Name, x.Group)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    for _, x := range updPortList {
        err = dbo.UpdPort(tx, x.IP, x.Port, x.Name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    err = tx.Commit().Error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes2)
}

func post(w http.ResponseWriter, r *http.Request) {
    contentType := r.Header.Get("Content-type")
    if strings.Contains(contentType, "multipart/form-data;") {
        processSubmit(w, r)
    } else {
        processConfirm(w, r)
    }
}

func Process(username string,
             w http.ResponseWriter,
             r *http.Request) {
    switch r.Method {
      case "GET":
          get(username, w)
      case "POST":
          post(w, r)
    }
}
