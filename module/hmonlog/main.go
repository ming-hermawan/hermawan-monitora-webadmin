package hmonlog

import (
    "fmt"
    "os"
    "time"
    "path/filepath"
    "github.com/google/uuid"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/module/hmonenv"
)


// PUBLIC

// 1. const

const (
    DbLog = iota
    RedisLog = iota
    WebserverLog = iota
)

// 2. func

func WriteLog(logType int,
              title string,
              msg string) (string, error) {
    currentTime := time.Now()
    var id string
    var logFilepath string
    var err error
    var sysFileLocation string
    var file *os.File
    switch logType {
      case DbLog:
          id = fmt.Sprintf("Db-%s", uuid.New())
          logFilepath = hmonglobal.DbLogFilepath
      case WebserverLog:
          id = fmt.Sprintf("Ws-%s", uuid.New())
          logFilepath = hmonglobal.WebserverLogFilepath
      default:
          id = fmt.Sprintf("%s", uuid.New())
          logFilepath = hmonglobal.DefaultLogFilepath
    }
    sysFileLocation = filepath.Join(
      hmonenv.GetLogDirPath(),
      logFilepath)
    file, err = os.OpenFile(
      sysFileLocation,
      os.O_APPEND|os.O_CREATE|os.O_WRONLY,
      0644)
    if err != nil {
        return "", err
    }
    defer file.Close()
    outMsg := fmt.Sprintf(
      "%s;%s;%s;%s\n",
      currentTime.Format("2006-01-02 15:04:05"),
      id,
      title,
      msg)
    _, err = file.WriteString(outMsg)
    if err != nil {
        return "", err
    }
    err = file.Sync()
    if err != nil {
        return "", err
    }
    return id, nil
}

func WriteLogDb(title string,
                msg string) (string, error) {
    return WriteLog(DbLog, title, msg)
}

func WriteLogRedis(title string,
                   msg string) (string, error) {
    return WriteLog(RedisLog, title, msg)
}

func WriteLogWebserver(title string,
                       msg string) (string, error) {
    return WriteLog(WebserverLog, title, msg)
}
