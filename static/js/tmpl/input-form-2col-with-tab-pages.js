function ShowPage(page) {
    $("#divInputPageBody > div").css("display", "none");
    $("#divInput").css("display", "block");
    $("#" + page).css("display", "block");
}


$(document).ready(function() {
    $("#divInputButton button").click(function () {
        var page = $(this).data("page");
        ShowPage(page);
    });
    $("#buttonImgCloseInput").click(function() {
        $("#divInput").css("display", "none");
    });
    $("#btnNewRow").click(function() {
        $("#statDb").val("INS");
        $("#frmInp input:text").val("");
        $("#frmInp input:password").val("");
        //$("#frmInp input:email").val("");
        ShowPage("divPageMain");
    });
});


function InitTbl(idx, on_select) {
    const frmDatList = form_data_init();
    $.ajax({
        url: MAIN_URL,
        type: "POST",
        contentType: 'application/json',
        data: JSON.stringify(frmDatList),
        dataType: "json",
        processData: false,
        success: function (result, textStatus, jqXHR) {
            for (let x in result.headers) {
                AddHeaderIntoTable(result.headers[x]);
            }
            AddRowsIntoTable(result.rows, idx);
            initSelect("grpFilter", frmDatList.grpFilter,result.groupList, "");
            for (let x in result.groupList) {
                $("#grpFilter")
            }
            $("#tbodyRows > tr > .td-select").click(function () {
                const servergroup = $(this).data("id");
                const frmDatDetail = {key: servergroup};
                HighlightRow($(this).data("n"));
                $.ajax({
                  url: MAIN_URL,
                  type: "POST",
                  data: JSON.stringify(frmDatDetail),
                  dataType: "json",
                  success: function (result) {
                      on_select(result);
                      ShowPage("divPageMain");
                      $("#statDb").val("UPD");
                      DisableButton();
                  }
                });
            });
            params = {
              "txtFilter": frmDatList.txtFilter}
            InitPageButtons(frmDatList.page, result.pageCount, MAIN_URL, params);
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log("ERROR:" + textStatus + " " + errorThrown);
        }
    });
}

