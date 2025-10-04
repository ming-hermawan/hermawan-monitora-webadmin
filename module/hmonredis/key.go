package hmonredis

import (
    "fmt"
    "github.com/google/uuid"
)


const LogoutAfter1Hour = "logout-after-1-hour"
const PortScanStatus = "ports-scan-status"

func GetPubSubServerNPort(ip string, port int) string {
    return fmt.Sprintf("pubsub:%s:%d", ip, port)
}

func GetLastServerPortStatus(ip string, port int) string {
    return fmt.Sprintf("last:%s:%d", ip, port)
}

func GetServerNPortKey(ip string, port int) string {
    return fmt.Sprintf("serverport:%s:%d", ip, port)
}

func GetMonPortUploadCSVKey() string {
    return fmt.Sprintf("mp-csv-%s", uuid.New())
}

func GetFailedLoginKey(username string) string {
    return fmt.Sprintf("failed-login-%s", username)
}

func GetUsrMenu(username string, menu string) string {
    return fmt.Sprintf("menu-%s-%s", username, menu)
}

func GetServerMails(ip string) string {
    return fmt.Sprintf("mail-%s", ip)
}
