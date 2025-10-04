function AddHeaderIntoTable(headerName) {
    let markup = "<th><div>" + headerName + "</div></th>";
    $("#trHeader").append(markup);
}


function AddRowsIntoTable(rows, idx) {
    var markup = "";
    for(x1 in rows) {
        markup += "<tr><td class=\"td-chkbox\"><div class=\"div-chkbox\"><label class=\"chkbox-tbl\"><input data-id=\"" + rows[x1][idx] + "\" type=\"checkbox\" value=0/><span class=\"checkmark\"><img src=\"/static/img/icon/icons8-close-24-black.png\"/></span></label></div><div class=\"div-chkbox-right-border\"></div></td>";
        for(x2 in rows[x1]) {
            markup += "<td class=\"td-select td-" + x1 + "\" data-n=\"" + x1 + "\" data-id=\"" + rows[x1][idx] + "\"><div class=\"div-normal\">" + rows[x1][x2] + "</div></td>";
        }
        markup += "</tr>"
    }
    $("#tbodyRows").append(markup);
}


function HighlightRow(n) {
    $("#tbodyRows > tr > .td-select > .div-highlight").addClass("div-normal").removeClass("div-highlight");
    $("#tbodyRows > tr > .td-" + n + " > div").removeClass("div-normal").addClass("div-highlight");
}


function InitPageButtons(page, pageCount, baseurl, params) {
    $("#labelPage").html(page);
    $("#labelPageCount").html(pageCount);
    var url = baseurl + "?";
    for (key in params) {
        if(params[key] !== "") {
            url += key + "=" + params[key];
        }
    }
    if(page == 1) {
        $("#linkPageFirst").addClass("btn-disabled");
        $("#linkPagePrev").addClass("btn-disabled");
    } else {
        $("#linkPageFirst").attr("href", url+"page=1")
        $("#linkPageFirst").addClass("btn");
        const pagePrev = page - 1;
        $("#linkPagePrev").attr("href", url+"page="+pagePrev);
        $("#linkPagePrev").addClass("btn");
    }
    if(pageCount == page) {
        $("#linkPageLast").addClass("btn-disabled");
        $("#linkPageNext").addClass("btn-disabled");
    } else {
        $("#linkPageLast").attr("href", url+"page="+pageCount);
        $("#linkPageLast").addClass("btn");
        const pageNext = page + 1;
        $("#linkPageNext").attr("href", url+"page="+pageNext);
        $("#linkPageNext").addClass("btn");
    }
}


function NewProcess(onNew) {
    $("#btnNewRow").click(function() {
        $("#statDb").val("INS");
        onNew();
    });
}


function SubmitProcess(submitValidation, getFrmDat) {
    $("#frmInp").submit(function(e) {
        $.LoadingOverlay("show");
        e.preventDefault();
        e.stopImmediatePropagation();
        if(!submitValidation()) {
            $.LoadingOverlay("hide");
            return
        }
        let frmDat = getFrmDat();
        frmDat["statDb"] = $("#statDb").val();
        $.ajax({
          url: MAIN_URL,
          type: "PUT",
          data: JSON.stringify(frmDat),
          dataType: "json",
          success: function(output, status, xhr) {
              $.LoadingOverlay("hide");
              if(output.Status === 1) {
                  window.location.reload();
              } else {
                  OpenAlertModal(output.Title, output.Message);
              }
          },
          error: function (xhr, textStatus, errorThrown) {
              $.LoadingOverlay("hide");
              OpenAlertModal(textStatus, xhr.responseText);
          }
        });
    });
}


$(document).ready(function() {
    $("#buttonImgCloseInput").click(function() {
        $("#divInput").css("display", "none");
    });
    $("#btnDelRow").click(function() {
        var message = "DELETE:<br>";
        $(".td-chkbox > div > label > input").map(function() {
            if($(this).prop("checked"))
                message += "- " + $(this).data("id") + "<br>";
        }).get();
        $("#inpModalAsk").val("DEL");
        OpenAskModal("CONFIRMATION", message);
    });
    $("#btnAskConfirm").click(function() {
        if ($("#inpModalAsk").val() == "DEL") {
            var keys = [];
            $(".chkbox-tbl > input").each(function(index) {
                if($(this).prop("checked")) {
                    keys.push($(this).data("id"));
                }
            });
            var frmDat = {
              "keys": keys};
            $.ajax({
              url: MAIN_URL,
              type: "DELETE",
              data: JSON.stringify(frmDat),
              dataType: "json",
              success: function(output, status, xhr) {
                  if(output.Status == 1) {
                      CloseAskModal(true);
                      window.location.reload();
                  }
                  else {
                      OpenAlertModal(output.Title, output.Message);
                  }
              },
              error: function (xhr, textStatus, errorThrown) {
                  OpenAlertModal(textStatus, xhr.responseText);
              }
            });
        }
    });
});
