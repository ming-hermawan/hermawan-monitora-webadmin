package hmonenv

import (
    "fmt"
    "os"
    "time"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
)

func CnvToTimeDurationSeconds(val int64) time.Duration {
    return time.Duration(val) * time.Second
}

func CnvToTimeDurationMilliSeconds(val int64) time.Duration {
    return time.Duration(val) * time.Millisecond
}

func getClientCert(crtFilepath string, keyFilepath string) (tls.Certificate, string) {
    var err error
    _, err = os.Stat(crtFilepath)
    if err != nil {
        return tls.Certificate{}, fmt.Sprintf("REDIS_CLIENT_CRT_PATH file path error! %s", err.Error())
    }
    _, err = os.Stat(keyFilepath)
    if err != nil {
        return tls.Certificate{}, fmt.Sprintf("REDIS_CLIENT_CRT_PATH file path error! %s", err.Error())
    }
    obj, err := tls.LoadX509KeyPair(crtFilepath, keyFilepath)
    if err != nil {
        return tls.Certificate{}, fmt.Sprintf("Error loading client certificate! %s", err.Error())
    }
    return obj, ""
}

func getCaCertPool(caFilePath string) (*x509.CertPool, string) {
    var err error
    _, err = os.Stat(caFilePath)
    if err != nil {
        return nil, fmt.Sprintf("REDIS_CLIENT_CA_PATH file path error! %s", err.Error())
    }
    caCert, err := ioutil.ReadFile(caFilePath)
    if err != nil {
        return nil, fmt.Sprintf("Error loading CA certificate! %s", err.Error())
    }
    obj := x509.NewCertPool()
    obj.AppendCertsFromPEM(caCert)
    return obj, ""
}

func getTlsConfig(crtFilepath string, keyFilepath string, caFilePath string) (*tls.Config, string) {
    var errMsg string
    var certs []tls.Certificate
    var caCertPool *x509.CertPool
    if crtFilepath != "" && keyFilepath != "" {
        var clientCert tls.Certificate
        clientCert, errMsg = getClientCert(crtFilepath, keyFilepath)
        if errMsg != "" {
            return nil, errMsg
        }
        certs = []tls.Certificate{clientCert}
    }
    if caFilePath != "" {
        caCertPool, errMsg = getCaCertPool(caFilePath)
        if errMsg != "" {
            return nil, errMsg
        }
    }
    tlsConfig := &tls.Config{
      Certificates: certs,
      RootCAs:caCertPool,
    }
    return tlsConfig, ""
}
