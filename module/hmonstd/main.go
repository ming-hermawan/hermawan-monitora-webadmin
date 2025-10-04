package hmonstd

import (
    "net"
    "strconv"
    "time"
)

func ChkTcp(ip string, port int) error {
    strPort := strconv.Itoa(port)
    address := net.JoinHostPort(ip, strPort)

    timeout := 5 * time.Second
    conn, err := net.DialTimeout("tcp", address, timeout)

    defer conn.Close()

    return err
}
