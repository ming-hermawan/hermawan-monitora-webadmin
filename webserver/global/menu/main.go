package menu

const AdmUsrMgtUsrGrp = "A-UM-usergroup"
const AdmUsrMgtUsr = "A-UM-user"
const AdmMonPortSettings = "A-M-ports-settings"
const AdmMonPortServerGrp = "A-M-ports-servergroup"
const AdmMonPortServerNPorts = "A-M-ports-serversNports"
const AdmMonPortAlertEmail = "A-M-P-email"
const AdmMonPortUploadCsv = "A-M-P-csv"
const MonPort = "M-ports"
const ReportMonPort = "R-M-ports"
const MenuCount = 9

var Menus = [MenuCount]string{
  AdmUsrMgtUsrGrp,
  AdmUsrMgtUsr,
  AdmMonPortSettings,
  AdmMonPortServerGrp,
  AdmMonPortServerNPorts,
  AdmMonPortAlertEmail,
  AdmMonPortUploadCsv,
  ReportMonPort,
  MonPort}
