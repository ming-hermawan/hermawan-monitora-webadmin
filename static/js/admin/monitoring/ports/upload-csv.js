function AddRowsIntoTable(info) {
    for(x1 in info) {
        markup = "<tr><td class=\"td-dat\"><div class=\"div-normal\">" + info[x1].src + "</div></td><td class=\"td-dat\"><div class=\"div-normal\">" + info[x1].desc + "<\div></td><td class=\"td-dat\"><div class=\"div-normal\">" + info[x1].sts + "<\div></td></tr>";
        $("#tbodyPorts").append(markup);
    }
}


$(document).ready(function() {
    RelocationUrl = "/dashboard";
    $("#myFile").on("change", function (e) {
        if (this.files.length > 0) {
            const fileSize = this.files.item(0).size;
            const fileMb = fileSize / 1024 ** 2;
            if (fileMb <= 2) {
            } else {
                OpenAlertModal("ERROR", "File not allowed more than 2 MB.");
            }
        }
    });

    $("#formUploadCSV").submit(function (e) {
        $.LoadingOverlay("show");
        e.preventDefault();
        var frmDat = new FormData(this);
        $.ajax({
            url: window.location.pathname,
            type: 'POST',
            data: frmDat,
            success: function (result) {
                $("#id").val(result.id);
                AddRowsIntoTable(result.info);
                $.LoadingOverlay("hide");
            },
            error: function (xhr, textStatus, errorThrown) {
              $.LoadingOverlay("hide");
              OpenAlertModal(textStatus, xhr.responseText);
            },
            cache: false,
            contentType: false,
            processData: false
        });
    })

    $("#btnConfirm").click(function (e) {
        id = $("#id").val();
        var frmDat = {
          "id": id};
        $.ajax({
            url: window.location.pathname,
            type: 'POST',
            data: JSON.stringify(frmDat),
            dataType: "json",
            success: function (result) {
                if (result.message == "SUCCESS") {
                    OpenAlertModal("SUCCESS", "Server & Ports Updated");
                }
            },
            cache: false,
            contentType: false,
            processData: false
        });
    })

});
