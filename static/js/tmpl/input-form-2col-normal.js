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
                      $("#divInput").css("display","block");
                      DisableButton();
                      $("#statDb").val("UPD");
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
