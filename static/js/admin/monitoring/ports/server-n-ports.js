const tbodyPortsHeaders = "<tr><th class=\"th-chk-box\"><div></div></th><th id=\"thPort\"><div>Port</div></th><th id=\"thServerName\"><div>Server Name</div></th></tr>"
const tbodyEmailsHeaders = "<tr><th class=\"th-chk-box\"><div></div></th><th id=\"thEmail\"><div>Email</div></th></tr>"


function getMarkupTrTdPort(port, service) {
    return "<tr data-id=\"" + port + "\"><td class=\"td-chkbox-in-tab td-chkbox-port\"><div><label class=\"chkbox-tbl\">"+
"<input data-id=\"" + port + "\" type=\"checkbox\" value=0/>"+
"<span class=\"checkmark\"><img src=\"/static/img/icon/icons8-close-24-black.png\"/></span></label><div></td><td><div>" + port + "</div></td><td><div>" + service +"</div></td></tr>";
}
function getMarkupTrTdEmail(email) {
    return "<tr data-id=\"" + email + "\"><td class=\"td-chkbox-in-tab td-chkbox-email\"><div><label class=\"chkbox-tbl\">"+
"<input data-id=\"" + email + "\" type=\"checkbox\" value=0/>"+
"<span class=\"checkmark\"><img src=\"/static/img/icon/icons8-close-24-black.png\"/></span></label><div></td><td><div>" + email + "</div></td></tr>";
}



function onSelect(result) {
    $("#ip").val(result.ip);
    $("#ip0").val(result.ip);
    $("#name").val(result.name);
    $("#name0").val(result.ip);
    initSelect("group", result.servergroup, result.servergroups, "");
    $("#group0").val(result.servergroup);
    jsonPorts = JSON.stringify(result.ports)
    jsonEmails = JSON.stringify(result.emails)
    $("#inpPorts").val(jsonPorts);
    $("#inpEmails").val(jsonEmails);
    $("#ports0").val(jsonPorts);
    $("#emails0").val(jsonEmails);
    $("#tbodyPorts").html(tbodyPortsHeaders);
    for (const key in result.ports) {
        $("#tbodyPorts").append(getMarkupTrTdPort(key, result.ports[key]));
    }
    $("#tbodyEmails").html(tbodyEmailsHeaders);
    for (const key in result.emails) {
        $("#tbodyEmails").append(getMarkupTrTdEmail(result.emails[key]));
    }
    $("#ip").prop('disabled', true);
}


function onNew() {
    $.ajax({
      url: "/ddl/admin/monitoring/port/server-group",
      type: "POST",
      dataType: "json",
      success: function (result) {
        initSelect("group", null, result, "");
      }
    });
    $("#frmInp input:text").val("");
    $("#frmInp input:email").val("");
    $("#ip").prop('disabled', false);
}


function onCancel() {
    $("#ip").val($("#ip0").val());
    $("#name").val($("#name0").val());
    $("#group").val($("#group0").val());
    ports = JSON.parse($("#ports0").val())
    emails = JSON.parse($("#emails0").val())
    $("#tbodyPorts").html(tbodyPortsHeaders);
    for (const key in ports) {
        $("#tbodyPorts").append(getMarkupTrTdPort(key, ports[key]));
    }
    $("#tbodyEmails").html(tbodyEmailsHeaders);
    for (const key in emails) {
        $("#tbodyEmails").append(getMarkupTrTdEmail(emails[key]));
    }
}


function submitValidation() {
    let ip = $("#ip").val().trim();
    if(ip === "") {
        $("#divInputMessage").html("IP can't be empty!");
        return false;
    }
    let group = $("#group").find(":selected").text();
    if(group === "") {
        $("#divInputMessage").html("Group can't be empty!");
        return false;
    }
    return true;
}


function getFrmDat() {
    let ip = $("#ip").val();
    let name = $("#name").val();
    let group = $("#group").find(":selected").text();
    var ports;
    try {
        ports = JSON.parse($("#inpPorts").val());
    }
    catch(err) {
        ports = {};
    }
    var emails;
    try {
        emails = JSON.parse($("#inpEmails").val());
    }
    catch(err) {
        emails = [];
    }
    out = {
      "ip": ip,
      "name": name,
      "group": group,
      "ports": ports,
      "emails": emails};
    return out;
}


$(document).ready(function(){
    InitTbl(0, onSelect);
    NewProcess(onNew);

    $("#btnAddPort").click(function() {
        const valPort = $("#port").val();
        const valService = $("#service").val();
        if( (isNaN(valPort)) || (valPort.trim() == "")) {
            $("#port").val("");
            $("#service").val("");
        } else {
            $("#tbodyPorts").append(getMarkupTrTdPort(valPort, valService));
            var obj;
            try {
                obj = JSON.parse($("#inpPorts").val());
            }
            catch(err) {
                obj = {};
            }
            obj[valPort] = valService;
            $("#inpPorts").val(JSON.stringify(obj));
            EnableButton();
        }
    });
    $("#btnDelPort").click(function() {
        $(".td-chkbox-port > div > label > input").map(function() {
            if($(this).prop("checked")) {
                var obj;
                try {
                    obj = JSON.parse($("#inpPorts").val());
                }
                catch(err) {
                    obj = {};
                }
                delete obj[$(this).data("id")];
                $("#inpPorts").val(JSON.stringify(obj));
                $("tr[data-id=\""+$(this).data("id")+"\"]").remove();
                EnableButton();
            }
        }).get();
    });

    $("#btnAddEmail").click(function() {
        const email = $("#email").val();
        console.log("email = "+email);
        if(email.trim() == "") {

        } else {
            console.log(getMarkupTrTdEmail(email));
            $("#tbodyEmails").append(getMarkupTrTdEmail(email));
            var obj;
            try {
                obj = JSON.parse($("#inpEmails").val());
            }
            catch(err) {
                obj = [];
            }
            obj.push(email);
            $("#inpEmails").val(JSON.stringify(obj));
            EnableButton();
        }
    });
    $("#btnDelEmail").click(function() {
        $(".td-chkbox-email > div > label > input").map(function() {
            if($(this).prop("checked")) {
                var obj;
                try {
                    obj = JSON.parse($("#inpEmails").val());
                }
                catch(err) {
                    obj = [];
                }
                const index = obj.indexOf($(this).data("id"));
                if (index > -1) { // only splice array when item is found
                    obj.splice(index, 1); // 2nd parameter means remove one item only
                }
                $("#inpEmails").val(JSON.stringify(obj));
                $("tr[data-id=\""+$(this).data("id")+"\"]").remove();
                EnableButton();
            }
        }).get();
    });

    CancelProcess(onCancel);
    SubmitProcess(submitValidation, getFrmDat);
});
