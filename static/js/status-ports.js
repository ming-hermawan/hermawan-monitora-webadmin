function tbl_header() {
    return "<tr>\
<th class=\"thd-server-left\"><div>Server</div></th>\
<th class=\"thd-port\"><div>Port</div></th>\
<th class=\"thd-status\"><div>Status</div></th>\
</tr>";
}

function svg_init() {
    return "<svg height=\"30\" width=\"30\">\
<defs>\
<linearGradient id=\"gradInit\" x1=\"0%\" x2=\"100%\" y1=\"0%\" y2=\"0%\">\
<stop offset=\"0%\" stop-color=\"#777777\" />\
<stop offset=\"100%\" stop-color=\"#ffffff\" />\
</linearGradient>\
<linearGradient id=\"gradOn\" x1=\"0%\" x2=\"0%\" y1=\"0%\" y2=\"100%\">\
<stop offset=\"0%\" stop-color=\"#ffffff\" />\
<stop offset=\"100%\" stop-color=\"#00ff00\" />\
</linearGradient>\
<linearGradient id=\"gradOff\" x1=\"0%\" x2=\"0%\" y1=\"0%\" y2=\"100%\">\
<stop offset=\"0%\" stop-color=\"#ffffff\" />\
<stop offset=\"100%\" stop-color=\"#ff0000\" />\
</linearGradient>\
</defs>\
<circle r=\"12\" cx=\"15\" cy=\"15\" fill=\"url(#gradInit)\"/>\
<circle r=\"10\" cx=\"15\" cy=\"15\" fill=\"#777777\"/>\
</svg>";
}

function svg_on() {
    return "<svg class=\"light-on\" height=\"30\" width=\"30\">\
<defs>\
<linearGradient id=\"gradInit\" x1=\"0%\" x2=\"100%\" y1=\"0%\" y2=\"0%\">\
<stop offset=\"0%\" stop-color=\"#777777\" />\
<stop offset=\"100%\" stop-color=\"#ffffff\" />\
</linearGradient>\
<linearGradient id=\"gradOn\" x1=\"0%\" x2=\"0%\" y1=\"0%\" y2=\"100%\">\
<stop offset=\"0%\" stop-color=\"#ffffff\" />\
<stop offset=\"100%\" stop-color=\"#00ff00\" />\
</linearGradient>\
<linearGradient id=\"gradOff\" x1=\"0%\" x2=\"0%\" y1=\"0%\" y2=\"100%\">\
<stop offset=\"0%\" stop-color=\"#ffffff\" />\
<stop offset=\"100%\" stop-color=\"#ff0000\" />\
</linearGradient>\
</defs>\
<circle r=\"12\" cx=\"15\" cy=\"15\" fill=\"url(#gradOn)\"/>\
</svg>";
}

function svg_off() {
    return "<svg class=\"light-off\" height=\"30\" width=\"30\">\
      <defs>\
        <linearGradient id=\"gradInit\" x1=\"0%\" x2=\"100%\" y1=\"0%\" y2=\"0%\">\
          <stop offset=\"0%\" stop-color=\"#777777\" />\
          <stop offset=\"100%\" stop-color=\"#ffffff\" />\
        </linearGradient>\
        <linearGradient id=\"gradOn\" x1=\"0%\" x2=\"0%\" y1=\"0%\" y2=\"100%\">\
          <stop offset=\"0%\" stop-color=\"#ffffff\" />\
          <stop offset=\"100%\" stop-color=\"#00ff00\" />\
        </linearGradient>\
        <linearGradient id=\"gradOff\" x1=\"0%\" x2=\"0%\" y1=\"0%\" y2=\"100%\">\
          <stop offset=\"0%\" stop-color=\"#ffffff\" />\
          <stop offset=\"100%\" stop-color=\"#ff0000\" />\
        </linearGradient>\
      </defs>\
      <circle r=\"12\" cx=\"15\" cy=\"15\" fill=\"url(#gradOff)\"/>\
    </svg>"
}

var socket = io();
socket.on('reply', function(msg) {
    try {
        const obj = JSON.parse(msg);
        if(obj.status == "RESET") {
            console.log("RESET");
            window.location.reload();
        }
        if(obj.status == "UP") {
            console.log("UP");
            $("#status-"+obj.uniqueId).html(svg_on());
        }
        else if(obj.status == "DOWN") {
            console.log("DOWN");
            $("#status-"+obj.uniqueId).html(svg_off());
        }
    } catch(err) {
        console.log(err);
        console.log("MSG:");
        console.log(msg);
    }
});

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

function run_filter() {
  console.log("groupServer = "+$("#groupServer").val());
  var frmDat = {
    groupServer: $("#groupServer").val(),
    serverKeyword: $("#serverKeyword").val(),
    serviceKeyword: $("#serviceKeyword").val(),
  };
  $.post("/get-server-ports", JSON.stringify(frmDat), function(data, status){
      console.log(data);
      tableBody = $("#tbodyContent");
      tableBody.html(tbl_header());
      $.each(data, function(index1, value1){
          portCount = parseInt(value1["portCount"]);
          nPort = 0;
          first = true;
          $.each(value1["ports"], function(index2, value2) {
              nPort++;
              if(value2.status == "INIT") {
                  svgMarkup = svg_init();
              } else if (value2.status == "UP") {
                  svgMarkup = svg_on();
              } else if (value2.status == "DOWN") {
                  svgMarkup = svg_off();
              } else {
                  console.log("STATUS = "+value2.status);
                  svgMarkup = svg_init();
              }
              if(portCount > 1) {
                  if(first) {
                      markup = "<tr class=\"tr-server-top\"><td colspan=\"2\"><div>" + index1 + " (" + value1.name + ")</div></td></tr>\
<tr>\
<td class=\"td-first td-last thd-server-left td-server-left\" rowspan=\""+value1["portCount"]+"\"><div>" + index1 + " (" + value1.name + ")</div></td>\
<td class=\"td-first thd-port td-port\"><div>" + index2 + " (" + value2.name +")</div></td>\
<td class=\"td-first thd-status td-status\"><div id=\"status-" + value2.uniqueId + "\">" + svgMarkup + "</div></td>\
</tr>";
                      first = false;
                  }
                  else {
                      if(nPort == portCount) {
                          markup = "<tr>\
<td class=\"td-last thd-port td-port\"><div>" + index2 + " (" + value2.name +")</div></td>\
<td class=\"td-last thd-status td-status\"><div id=\"status-" + value2.uniqueId + "\">" + svgMarkup + "</div></td>\
</tr>";
                      } else {
                        markup = "<tr>\
<td class=\"thd-port td-port\"><div>" + index2 + " (" + value2.name +")</div></td>\
<td class=\"thd-status td-status\"><div id=\"status-" + value2.uniqueId + "\">" + svgMarkup + "</div></td>\
</tr>";
                      }
                  }
              } else {
                  markup = "<tr class=\"tr-server-top\"><td colspan=\"2\"><div>" + index1 + " (" + value1.name + ")</div></td></tr>\
<tr>\
<td class=\"td-first td-last thd-server-left td-server-left\"><div>" + index1 + " (" + value1.name + ")</div></td>\
<td class=\"td-first td-last thd-port td-port\"><div>" + index2 + " (" + value2.name +")</div></td>\
<td class=\"td-first td-last thd-status td-status\"><div id=\"status-" + value2.uniqueId + "\">" + svgMarkup + "</div></td>\
</tr>";
              }
              tableBody.append(markup);
          });
      });
  });
}

$(document).ready(function(){
    const urlParams = new URLSearchParams(window.location.search);
    const filterParam = urlParams.get('filter');
    const serverKeyword = urlParams.get('server-keyword');
    if(serverKeyword) {
        $("#serverKeyword").val(serverKeyword);
    }
    const serviceKeyword = urlParams.get('service-keyword');
    if(serviceKeyword) {
        $("#serviceKeyword").val(serviceKeyword);
    }
    if(filterParam == "1") {
        run_filter();
    }
    $.ajax({
      url: "/ddl/admin/monitoring/port/server-group",
      type: "POST",
      dataType: "json",
      success: function (result) {
        initSelect("groupServer", groupServer, result, "");
      }
    });
    $("#filterFrm").submit(function(e){
        console.log("SUBMIT");
        e.preventDefault();
        run_filter();
    });
});
