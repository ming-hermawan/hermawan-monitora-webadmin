var ports = {}

function tbl_ports_header() {
    return "<tr>\
<th><div>Port</div></th>\
<th><div>Service Name</div></th>\
<th><div>Del</div></th>\
</tr>";
}

function input_edit() {
    $(".input-default").prop('disabled', false);
    $("#divButtonControl").hide();
    $("#divButtonEdit").show();
}

function init() {
    var urlParams = new URLSearchParams(window.location.search);
    if(urlParams.has('ip')) {
        ip = urlParams.get('ip');
        if(ip == "new") {
            input_edit();
        } else {
            var frmDat = {ip: ip};
            $.post("/admin/monitoring/ports",  JSON.stringify(frmDat), function(data, status) {
                $("#serverIP").val(data.Ip);
                $("#serverName").val(data.Name);
                $("#serverGroup").val(data.ServerGroup);
                tableBodyPorts = $("#tbodyPorts");
                tableBodyPorts.html(tbl_ports_header());
                serverGroupsDataList = $("#serverGroupsDataList");
                $.each(data.Ports, function(n, item) {
                    markup = "<tr class=\"tr-" + n + "\">\
<td><div>" + n + "</div></td>\
<td><div>" + item + "</div></td>\
<td><div><input class=\"del-port-checkbox\" type=\"checkbox\" data-port=\"" + n + "\" data-servicename=\"" + item + "\" /></div></td>\
</tr>";
                    tableBodyPorts.append(markup);
                    ports[n] = item;
                });
                $.each(data.MPServerGroup, function(n, item) {
                    markup = "<option value=\"" + item + "\">";
                    serverGroupsDataList.append(markup);
                });
            });
        }
    }

    $("#divInputPorts").css('visibility', 'hidden');
    $("#divButtonEdit").hide();
    $("#divPortsButtonControl").show();
    $("#divButtonControl").show();
    $(".input-default").prop('disabled', true);
}

$(document).ready(function() {
    init();
    // var urlParams = new Proxy(new URLSearchParams(window.location.search), {
    //     get: (searchParams, prop) => searchParams.get(prop),
    // });
    // console.log(urlParams.ip)

    $("#btnEdit").click(function() {
        input_edit();
    });
    $("#btnSave").click(function() {
        ip = $("#serverIP").val();
        name = $("#serverName").val();
        server_group = $("#serverGroup").val();
        ports = {};

        var elements = $(".del-port-checkbox");
        $.each(elements, function(n, item) {
            ports[item.dataset.port] = item.dataset.servicename;
        });

        var frmDat = {
          "ip": ip,
          "name": name,
          "server_group": server_group,
          "ports": ports};
        $.ajax({
          url: "/setting/monitoring/ports",
          type: "PUT",
          data: JSON.stringify(frmDat),
          dataType: "json",
          success: function (result) {
            console.log(result);
          }
        });
    });
    $("#btnDelPorts").click(function() {
        console.log("CLICK");
        var elements = $(".del-port-checkbox");
        $.each(elements, function(n, item) {
            if(item.checked) {
                $(".tr-"+item.dataset.port).remove();
            }
        });
    });
    $("#btnCancel").click(function() {
        init();
    });
    $("#idButtonNewPort").click(function() {
        $("#divInputPorts").css('visibility', 'visible');
        $("#divPortsButtonControl").hide();
        $("#servicePort").focus();
    });
    $("#idButtonSavePort").click(function() {
      port = $("#servicePort").val();
      serviceName = $("#serviceName").val();

      tableBodyPorts = $("#tbodyPorts");
      markup = "<tr class=\"tr-" + port + "\">\
<td><div>" + port + "</div></td>\
<td><div>" + serviceName + "</div></td>\
<td><div><input class=\"del-port-checkbox\" type=\"checkbox\" data-port=\"" + port + "\" data-servicename=\"" + serviceName + "\" /></div></td>\
</tr>";
      tableBodyPorts.append(markup);
    });
});
