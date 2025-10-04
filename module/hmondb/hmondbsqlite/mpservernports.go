package hmondbsqlite

import (
  "errors"
  "fmt"
  "gorm.io/gorm"
)

type MPServer struct {
    Ip string  `gorm:"primaryKey"`
    Name string
    ServerGroup string
}

type MPPort struct {
    Ip string `gorm:"primaryKey"`
    Port int `gorm:"primaryKey"`
    Name string
}

type MPServerAlertEmail struct {
    Ip string `gorm:"primaryKey"`
    Email string `gorm:"primaryKey"`
}

type ServerNPort struct {
    Ip string
    Port int
    ServerName string
    ServiceName string
}

func GetCntOfMPServer(db *gorm.DB, grpFilter string, txtFilter string) (int64, error) {
    var out int64
    if (grpFilter != "") && (grpFilter != "null") && (txtFilter != "") && (txtFilter != "null") {
        sql := "(server_group = ?) AND (ip LIKE ?) OR (name LIKE ?)"
        db.Model(&MPServer{}).Where(sql, grpFilter, txtFilter).Count(&out)
    } else if (grpFilter != "") && (grpFilter != "null") {
        sql := "server_group = ?"
        db.Model(&MPServer{}).Where(sql, grpFilter).Count(&out)
    } else if (txtFilter != "") && (txtFilter != "null") {
        sql := "(ip LIKE ?) OR (name LIKE ?)"
        db.Model(&MPServer{}).Where(sql, txtFilter).Count(&out)
    } else {
        db.Model(&MPServer{}).Count(&out)
    }
    return out, nil
}

func GetMPServerForTbl(db *gorm.DB, grpFilter string, txtFilter string, pageNumber int, pageLimit int) ([][]string, error) {
    var dbResult *gorm.DB
    var servers []MPServer
    offset := (pageNumber - 1) * pageLimit
    if (grpFilter != "") && (grpFilter != "null") && (txtFilter != "") && (txtFilter != "null") {
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        sql := "(server_group = ?) AND (ip LIKE ?) OR (name LIKE ?)"
        dbResult = db.Where(
          sql,
          grpFilter,
          txtLikeFilter,
          txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&servers)
    } else if (grpFilter != "") && (grpFilter != "null") {
        sql := "server_group = ?"
        dbResult = db.Where(
          sql,
          grpFilter).Limit(pageLimit).Offset(offset).Find(&servers)
    } else if (txtFilter != "") && (txtFilter != "null") {
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        sql := "(ip LIKE ?) OR (name LIKE ?)"
        dbResult = db.Where(
          sql,
          txtLikeFilter,
          txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&servers)
    } else {
        dbResult = db.Limit(pageLimit).Offset(offset).Find(&servers)
    }
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := [][]string{}
    for k := range servers {
        temp := []string{servers[k].Ip, servers[k].Name, servers[k].ServerGroup}
        out = append(out, temp)
    }
    return out, nil
}

func GetServerEmails(db *gorm.DB, mpAlertEmails []MPAlertEmail) (map[string]string, error) {
    defaultEmails := ""
    if mpAlertEmails != nil {
        lenMPAlertEmails := len(mpAlertEmails)
        n := 0
        for _, val := range mpAlertEmails {
            defaultEmails += val.Email
            if n < (lenMPAlertEmails - 1) {
                defaultEmails += ","
            }
            n += 1
        }
    }
    var dbResult *gorm.DB
    var servers []MPServer
    dbResult = db.Find(&servers)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    var serverAlertEmails []MPServerAlertEmail
    dbResult = db.Find(&serverAlertEmails)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    emails := map[string]string{}
    for _, val := range serverAlertEmails {
        _, exists := emails[val.Ip]
        if exists {
            emails[val.Ip] = fmt.Sprintf("%s,%s", emails[val.Ip], val.Email)
        } else {
            if defaultEmails == "" {
                emails[val.Ip] = val.Email
            } else {
                emails[val.Ip] = fmt.Sprintf("%s,%s", defaultEmails, val.Email)
            }
        }
    }
    out := map[string]string{}
    for _, val := range servers {
        email, exists := emails[val.Ip]
        if exists {
            out[val.Ip] = email
        } else {
            out[val.Ip] = defaultEmails
        }
    }
    return out, nil
}


func GetMPServerAlertEmail(db *gorm.DB, ip string) ([]string, error) {
    var serverAlertEmails []MPServerAlertEmail
    dbResult := db.Where(
      "ip = ?",
      ip).Find(&serverAlertEmails)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := []string{}
    for k := range serverAlertEmails {
        out = append(out, serverAlertEmails[k].Email)
    }
    return out, nil

}

func IsServerExists(ip string) (bool, error) {
    db, err := GetDb()
    if err != nil {
        return false, err
    }
    var count int64
    sql := "ip = ?"
    db.Model(&MPServer{}).Where(sql, ip).Count(&count)
    if (count == 0) {
        return false, nil
    } else {
        return true, nil
    }
}

func IsPortExists(ip string, port int) (bool, error) {
    db, err := GetDb()
    if err != nil {
        return false, err
    }
    var count int64
    sql := "ip = ? AND port = ?"
    db.Model(&MPPort{}).Where(sql, ip, port).Count(&count)
    if (count == 0) {
        return false, nil
    } else {
        return true, nil
    }
}

func IsServerAlertEmailExists(ip string, email string) (bool, error) {
    db, err := GetDb()
    if err != nil {
        return false, err
    }
    var count int64
    sql := "ip = ? AND email = ?"
    db.Model(&MPPort{}).Where(sql, ip, email).Count(&count)
    if (count == 0) {
        return false, nil
    } else {
        return true, nil
    }
}

func IsServerGroupUsedInServers(servergroups []string) (bool, error) {
    db, err := GetDb()
    if err != nil {
        return false, err
    }
    var count int64
    sql := "server_group IN (?)"
    db.Model(&MPServer{}).Where(sql, servergroups).Count(&count)
    if (count == 0) {
        return false, nil
    } else {
        return true, nil
    }
}

func InsServer(db *gorm.DB, ip string, name string, group string) error {
    server := MPServer {
      Ip: ip,
      Name: name,
      ServerGroup: group}
    result := db.Create(server)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 1 {
        return nil
    } else {
        return errors.New("No Rows Affected")
    }
}

func BulkInsServer(db *gorm.DB, servers []MPServer) error {
    result := db.Create(servers)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected >= 0 {
        return nil
    } else {
        return errors.New("No Rows Affected")
    }
}

func UpdServer(db *gorm.DB, ip string, name string, group string) error {
    var serverRow MPServer
    if err := db.Where("ip = ?", ip).First(&serverRow).Error; err != nil {
        return err
    }
    serverRow.Name = name
    serverRow.ServerGroup = group
    result := db.Where("ip = ?", ip).Save(&serverRow)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func DelMPServer(db *gorm.DB, ips []string) error {
    var mpServers []MPServer
    if err := db.Where(
      "ip in (?)", ips).Find(&mpServers).Error; err != nil {
        return err
    }
    result := db.Where(
      "ip in (?)",
      ips).Delete(&mpServers)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func BulkInsPort(db *gorm.DB, ports []MPPort) error {
    result := db.Create(ports)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected >= 0 {
        return nil
    } else {
        return errors.New("No Rows Affected")
    }
}

func UpdPort(db *gorm.DB, ip string, port int, name string) error {
    var portRow MPPort
    if err := db.Where("ip = ? AND port = ?", ip, port).First(&portRow).Error; err != nil {
        return err
    }
    portRow.Name = name
    result := db.Where("ip = ? AND port = ?", ip, port).Save(&portRow)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func DelMPPorts(db *gorm.DB, ips []string) error {
    var mpPorts []MPPort
    if err := db.Where(
      "ip in (?)",
      ips).Find(&mpPorts).Error; err != nil {
        return err
    }
    dbResult := db.Where(
      "ip in (?)",
      ips).Delete(&mpPorts)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    return nil
}

func UpdMPPorts(db *gorm.DB, ip string, ports map[int]string) error {
    dbResult := db.Where("ip = ?", ip).Delete(&MPPort{})
    if dbResult.Error != nil {
        return dbResult.Error
    }
    for key, val := range ports {
        mpPort := &MPPort {Ip: ip, Port: key, Name: val}
        dbResult := db.Create(mpPort)
        if dbResult.Error != nil {
            return dbResult.Error
        }
    }
    return nil
}


func BulkInsMPServerAlertEmail(db *gorm.DB, serverAlertEmails []MPServerAlertEmail) error {
    result := db.Create(serverAlertEmails)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected >= 0 {
        return nil
    } else {
        return errors.New("No Rows Affected")
    }
}
func DelMPServerAlertEmail(db *gorm.DB, ips []string) error {
    var mpServerAlertEmail []MPServerAlertEmail
    if err := db.Where(
      "ip in (?)",
      ips).Find(&mpServerAlertEmail).Error; err != nil {
        return err
    }
    dbResult := db.Where(
      "ip in (?)",
      ips).Delete(&mpServerAlertEmail)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    return nil
}

func UpdMPServerAlertEmail(db *gorm.DB, ip string, emails []string) error {
    dbResult := db.Where("ip = ?", ip).Delete(&MPServerAlertEmail{})
    if dbResult.Error != nil {
        return dbResult.Error
    }
    for _, val := range emails {
        mpServerAlertEmail := &MPServerAlertEmail {Ip: ip, Email: val}
        dbResult := db.Create(mpServerAlertEmail)
        if dbResult.Error != nil {
            return dbResult.Error
        }
    }
    return nil
}


func GetMPServerNPortsForRedis(db *gorm.DB) ([]ServerNPort, error) {
    leftJoinQuery := "LEFT JOIN mp_servers ON mp_servers.ip = mp_ports.ip"
    row, err := db.Table("mp_ports").Select("mp_ports.ip", "mp_ports.port", "mp_servers.name", "mp_ports.name").Joins(leftJoinQuery).Order("mp_ports.ip, mp_ports.port").Rows()
    if err != nil {
        return nil, err
    }
    out := []ServerNPort{}
    for row.Next() {
        var ip string
        var port int
        var serverName string
        var serviceName string
        row.Scan(
          &ip,
          &port,
          &serverName,
          &serviceName)
        temp := ServerNPort{Ip: ip, Port: port, ServerName: serverName, ServiceName: serviceName}
        out = append(out, temp)
    }
    return out, nil
}
