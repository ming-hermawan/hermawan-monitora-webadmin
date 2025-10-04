function tbl_header() {
    return "<tr>\
<th><div>Server IP</div></th>\
<th><div>Server Name</div></th>\
<th class=\"th-icon\"><div></div></th>\
<th class=\"th-icon\"><div></div></th>\
<th class=\"th-icon\"><div></div></th>\
</tr>";
}

function get_server_ports() {
    $.post("/get-server-ports", {}, function(data, status){
      tableBody = $("table tbody");
      $.each(data, function(m, item1) {
          $.each(item1.ports, function(n, item2) {
              console.log(item2)

              markup = "<tr><td>"+lineNo+"</td></tr>";
              tableBody.append(markup);

          });
      });
    });
}

$(document).ready(function(){
    $.post("/get-server-groups", {}, function(data, status){
        $.each(data, function(index, value) {
            option = "<option value="+value+">"+value+"</option>";
            $("#groupServer").append(option);
        });
    });
    $("#filterFrm").submit(function(e){
        console.log("SUBMIT");
        e.preventDefault();
        var frmDat = {
          groupServer: $("#groupServer").val(),
          serverKeyword: $("#serverKeyword").val(),
          serviceKeyword: $("#serviceKeyword").val(),
        };
        $.post("/get-server-ports0", JSON.stringify(frmDat), function(data, status) {
            tableBody = $("#tbodyContent");
            tableBody.html(tbl_header());
            $.each(data, function(index1, value1){
                markup = "<tr>\
<td><div>" + index1 + "</div></td>\
<td><div>" + value1["name"] + "</div></td>\
<td class=\"td-icon\"><div><button><img src=\"/static/img/icon/icons8-edit-24.png\"></img></button></div></td>\
<td class=\"td-icon\"><div><button><img src=\"/static/img/icon/icons8-delete-24.png\"></img></button></div></td>\
<td class=\"td-icon\"><div><button><img src=\"/static/img/icon/icons8-drop-down-24.png\"></img></button></div></td>\
</tr>";
                tableBody.append(markup);
            });
        });
    });
});
