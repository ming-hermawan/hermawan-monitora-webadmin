package port

import (
    "log"
    "encoding/json"
    "context"
    "fmt"
    "os"
    "path"
    "net/http"
    socketio "github.com/googollee/go-socket.io"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)

const Lbl = "Ports Monitoring"

type Stat struct {
    status string `json:"status"`
    time string `json:"time"`
}

type ServerPortParam struct {
    groupServer string `json:"groupServer"`
    serverKeyword string `json:"serverKeyword"`
    serviceKeyword string `json:"serviceKeyword"`
}

var stat Stat

var ctx = context.Background()

//dipakai di temp
var Menuhtmlfilepath = path.Join("view", "header.html")


func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      Lbl,
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.MonPortHtmlFilepath,
    )
}


func writeLog(uniqueId string, txt string) {
    f, err := os.Create(fmt.Sprintf("/home/workdir/out/%s.log", uniqueId))
    if err != nil {
        log.Println(fmt.Println(err))
        return
    }
    n, err := f.WriteString(txt)
    if err != nil {
        log.Println(fmt.Println(err))
        f.Close()
        return
    }
    log.Println(fmt.Println(n))
    err = f.Close()
    if err != nil {
        log.Println(fmt.Println(err))
        return
    }
}

func redisGetLastStatus(ip string, port int) string {
    val, err := hmonredis.Get(hmonredis.GetLastServerPortStatus(ip, port))
    if err != nil {
        fmt.Println(err.Error())
        return "INIT"
    }
    return val
}

func RedisSubscribe(s socketio.Conn, ip string, port int, uniqueId string) {
    key := hmonredis.GetPubSubServerNPort(ip, port)
    var reading map[string]interface{}
    log.Println(fmt.Printf("LISTEN TO %s\n", key))
    subscriber := hmonredis.Subscribe(ctx, key)
    for {
        msg, err := hmonredis.SubscriberReceiveMessage(subscriber)
        if err != nil {
            txt := fmt.Sprintf("ERROR Redis subscribe\nMESSAGE = \"%s\"\nERROR = \"%s\"\n", msg, err)
            log.Println(fmt.Printf("%s\n", txt))
            writeLog(uniqueId, txt)
            return
        }
        log.Println(fmt.Printf("CHANNEL = %s\nPAYLOAD = %s\n", msg.Channel, msg.Payload))
        if err := json.Unmarshal([]byte(msg.Payload), &reading); err != nil {
            txt := fmt.Sprintf("ERROR json\nCHANNEL = %s\nPAYLOAD = \"%s\"\nERROR = \"%s\"\n", msg.Channel, msg.Payload, err)
            log.Println(fmt.Printf("%s\n", txt))
            return
        }
        status := reading["status"]
        time := reading["time"]
        log.Println(fmt.Printf("--- map[string]interface{} BEGIN ---\n"))
        log.Println(fmt.Printf("%+v\n", reading))
        log.Println(fmt.Printf("STATUS = %s, TIME = %f\n", status, time))
        log.Println(fmt.Printf("---- map[string]interface{} END ----\n"))
        message1 := fmt.Sprintf(
          "{\"uniqueId\":%q,\"status\":%q}",
	  uniqueId,
	  status)
        s.Emit("reply", message1)
    }
}

func getServerPort0(serverGroupParam string, serverKeywordParam string, serviceKeywordParam string)  map[string]interface{} {
    temp := dbo.GetServerPorts(serverGroupParam, serverKeywordParam, serviceKeywordParam)
    return temp
}

func GetServerPort(serverGroupParam string, serverKeywordParam string, serviceKeywordParam string)  map[string]interface{} {
    temp := dbo.GetServerPorts(serverGroupParam, serverKeywordParam, serviceKeywordParam)
    for kServer, rowServerPort := range temp {
        temp1 := rowServerPort.(map[string]interface{})["ports"].(map[int]map[string]string)
        for kPort, rowPort := range temp1 {
            Status := redisGetLastStatus(kServer, kPort)
            rowPort["status"] = Status
        }
    }
    return temp
}

func GetPortsMonitoringPage(username string, w http.ResponseWriter, r *http.Request) {
    // filter := r.URL.Query().Get("filter")
    // groupServer := r.URL.Query().Get("group-server")
    // serverKeyword := r.URL.Query().Get("server-keyword")
    // serviceKeyword := r.URL.Query().Get("service-keyword")
    // if filter == "1" {
    //     ServerPort := GetServerPort(groupServer, serverKeyword, serviceKeyword)
    //     log.Println(fmt.Printf("GET ServerPort:\n%+v\n---\n", ServerPort))
    // }
    get(username, w)
}


// func RedisSetScanStatusInit(w http.ResponseWriter) {
//     redisKey := "ports-scan-status"
//     hmonredis.SetStr(redisKey, "INIT")
//     if err != nil {
//         httpresponse.ErrResponseSetRedis(
//           w,
//           redisKey,
//           err)
//         return
//     }
// }


func ProcessGetServerPort(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    log.Println(fmt.Printf("params:\n%+v\n---\n", params))
    servergroup, _ := validation.StrParamValidation(
      w,
      params,
      "groupServer",
      "Server-Group",
      false,
      hmonglobal.RegexId)
    serverKeyword, _ := validation.StrParamValidation(
      w,
      params,
      "serverKeyword",
      "Server Keyword",
      false,
      "")
    serviceKeyword, _ := validation.StrParamValidation(
      w,
      params,
      "serviceKeyword",
      "Service Keyword",
      false,
      "")

    ServerPort := GetServerPort(servergroup, serverKeyword, serviceKeyword)
    log.Println(fmt.Printf("ServerPort:\n%+v\n---\n", ServerPort))

    jsonInBytes, err := json.Marshal(&ServerPort)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}
