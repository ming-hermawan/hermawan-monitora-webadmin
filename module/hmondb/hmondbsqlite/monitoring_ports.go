package hmondbsqlite

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "log"
    "strings"
    "hermawan-monitora/module/hmonenv"
)

func GetListOfMPServer(grpFilter string, txtFilter string, pageNumber int, pageLimit int) []MPServer {
    offset := (pageNumber - 1) * pageLimit
    db, err := GetDb()
    if err != nil {
        panic("failed to connect database")
    }
    if (grpFilter != "") && (txtFilter != "") {
        sql := "(server_group = ?) AND (ip LIKE '%?%') OR (name LIKE '%?%')"
        db.Where(sql, grpFilter, txtFilter)
    } else if (grpFilter != "") {
        sql := "server_group = ?"
        db.Where(sql, grpFilter)
    } else if (txtFilter != "") {
        sql := "(ip LIKE '%?%') OR (name LIKE '%?%')"
        db.Where(sql, txtFilter)
    }
    var servers []MPServer
    db.Limit(pageLimit).Offset(offset).Find(&servers)
    return servers
}

// func GetMPServerGroup() []string {
//     db, err := gorm.Open(sqlite.Open(hmonenv.GetSQLiteDbFilePath()), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database")
//     }
//     var serverGroups []MPServerGroup
//     db.Find(&serverGroups)
//     out := []string{}
//     for k := range serverGroups {
//         out = append(out, serverGroups[k].ServerGroup)
//     }
//     return out
// }

func GetServerPorts(serverGroupParam string, serverKeywordParam string, serviceKeywordParam string) map[string]interface{} {
    serverGroupParamHasValue := false
    if serverGroupParam != "" {
        serverGroupParamHasValue = true
    }
    serverKeywordParamHasValue := false
    if serverKeywordParam != "" {
        serverKeywordParamHasValue = true
    }
    serviceKeywordParamHasValue := false
    if serviceKeywordParam != "" {
        serviceKeywordParamHasValue = true
    }
    leftJoinQuery := "LEFT JOIN mp_servers ON mp_servers.ip = mp_ports.ip"
    if serverKeywordParamHasValue {
        leftJoinQuery = fmt.Sprintf("%s AND mp_servers.name LIKE '%%%s%%'", leftJoinQuery, serverKeywordParam)
    }
    whereQuery := ""
    if serverGroupParamHasValue {
        whereQuery = fmt.Sprintf("mp_servers.server_group = '%s'", serverGroupParam)
    }
    if serverKeywordParamHasValue {
        if serverGroupParamHasValue {
            whereQuery = fmt.Sprintf("%s AND ", whereQuery)
        }
        whereQuery = fmt.Sprintf("%s mp_servers.name LIKE '%%%s%%'", whereQuery, serverKeywordParam)
    }
    if serviceKeywordParamHasValue {
      if serverGroupParamHasValue || serverKeywordParamHasValue {
          whereQuery = fmt.Sprintf("%s AND ", whereQuery)
      }
        whereQuery = fmt.Sprintf("%s mp_ports.name LIKE '%%%s%%'", whereQuery, serviceKeywordParam)
    }
    db, err := gorm.Open(sqlite.Open(hmonenv.GetSQLiteDbFilePath()), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.Model(&MPPort{}).Select("mp_servers.server_group", "mp_ports.ip", "mp_servers.name", "mp_ports.port", "mp_ports.name").Joins(leftJoinQuery).Where(whereQuery).Order("mp_servers.server_group, mp_ports.ip, mp_ports.port").Scan(&result{})
    row, err := db.Table("mp_ports").Select("mp_servers.server_group", "mp_ports.ip", "mp_servers.name", "mp_ports.port", "mp_ports.name").Joins(leftJoinQuery).Where(whereQuery).Order("mp_servers.server_group, mp_ports.ip, mp_ports.port").Rows()
    if err != nil {
        panic(err)
    }

    serverPorts := map[string]interface{}{}
    portCount := 0
    for row.Next() {
        var serverGroup string
        var serverIp string
        var serverName string
        var portNumber int
        var portName string
        row.Scan(
          &serverGroup,
          &serverIp,
          &serverName,
          &portNumber,
          &portName)
        UniqueId := fmt.Sprintf(
          "%s-%d",
          strings.Replace(serverIp, ".", "-", 4),
          portNumber)
        Port := map[string]string{
          "uniqueId": UniqueId,
          "name": portName,
          "status": ""}
        log.Println(fmt.Sprintf("Unique-ID in Port=%s\n", Port["uniqueId"]))
        entry1, ok := serverPorts[serverIp]
        if ok {
            portCount++
            entry2 := entry1.(map[string]interface{})["ports"].(map[int]map[string]string)
            entry2[portNumber] = Port
            entry1.(map[string]interface{})["portCount"] = portCount
        } else {
            portCount = 1
            serverPorts[serverIp] = map[string]interface{} {
              "name": serverName,
              "serverGroup": serverGroup,
              "portCount": portCount,
              "ports": map[int]map[string]string{portNumber: Port}}
              log.Println(fmt.Sprintf("Unique-ID in map=%s\n", serverPorts[serverIp].(map[string]interface{})["ports"].(map[int]map[string]string)[portNumber]["uniqueId"]))
        }
    }
    return serverPorts
}

func GetServerPorts0(db *gorm.DB, id string) ServerPorts {
    var server MPServer
    db.Where("ip = ?", id).First(&server)
    var ports []MPPort
    db.Where("ip = ?", id).Find(&ports)
    return ServerPorts {
        Server: server,
        Ports: ports}
}

func UpdateServerPorts(ip string, name string, server_group string, ports map[int]string) {
    db, err := gorm.Open(sqlite.Open(hmonenv.GetSQLiteDbFilePath()), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    var server MPServer
    updateStatus := true
    if err := db.Where("ip = ?", ip).First(&server).Error; err != nil {
        updateStatus = false
    }
    server.Name = name
    server.ServerGroup = server_group
    if updateStatus {
        db.Where("ip = ?", ip).Save(&server)
        var ports []MPPort
        db.Where("ip = ?", ip).Find(&ports)
        db.Where("ip = ?", ip).Unscoped().Delete(ports)
    } else {
        server.Ip = ip
        result := db.Create(server)
        log.Println(fmt.Sprintf("server Create result = %+v\n", result))
    }
    for k, v := range ports {
        var ports MPPort
        ports.Ip = ip
        ports.Port = k
        ports.Name = v
        result := db.Create(ports)
        log.Println(fmt.Sprintf("serverPort Create result = %+v\n", result))
    }
}
