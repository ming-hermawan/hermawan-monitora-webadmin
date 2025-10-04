package hmondbsqlite

import (
    "strconv"
    "time"
    "gorm.io/gorm"
)


// DB Struc

type MPHistory struct {
    Ip string `gorm:"primaryKey"`
    Port int `gorm:"primaryKey"`
    EpochTimestamp int64 `gorm:"primaryKey"`
    Status string
}

func GetMPHistory(db *gorm.DB, from int64, to int64) ([][]string, error) {
    leftJoinQuery := "LEFT JOIN mp_servers ON mp_servers.ip = mp_histories.ip LEFT JOIN mp_ports ON mp_ports.ip = mp_histories.ip AND mp_ports.port = mp_histories.port"
    whereQuery := "mp_histories.epoch_timestamp >= ? AND mp_histories.epoch_timestamp < ?"
    row, err := db.Table("mp_histories").Select("mp_histories.ip", "mp_servers.name", "mp_histories.port", "mp_ports.name", "mp_histories.epoch_timestamp", "mp_histories.status").Joins(leftJoinQuery).Where(whereQuery, from, to).Order("mp_histories.ip, mp_histories.port, mp_histories.epoch_timestamp").Rows()
    if err != nil {
        return nil, err
    }
    out := [][]string{}
    for row.Next() {
        var ip string
        var serverName string
        var port int
        var serviceName string
        var epochTimestamp int64
        var status string
        row.Scan(
          &ip,
          &serverName,
          &port,
          &serviceName,
          &epochTimestamp,
          &status)
        seconds := epochTimestamp / 1_000_000
        nanoseconds := (epochTimestamp % 1_000_000) * 1_000
        t := time.Unix(seconds, nanoseconds)
	strPort := strconv.Itoa(port)
        temp := []string{ip, serverName, strPort, serviceName, t.Format("2006-01-02 15:04:05"), status}
        out = append(out, temp)
    }
    return out, nil
}
