function form_data_init() {
    var page = parseInt(QUERY_STRING_PARAMS.get("page"));
    if(!page) {
        page = 1;
    }
    var txtFilter = QUERY_STRING_PARAMS.get("txtFilter");
    if(txtFilter) {
        console.log("txtFilter="+txtFilter);
        $("#txtFilter").val(txtFilter);
    } else {
        txtFilter = "";
    }
    return {
      page: page,
      txtFilter: txtFilter}
}


$(document).ready(function(){
    $("#btnFilter").click(function() {
        var queryString = "?";
        var page = parseInt(QUERY_STRING_PARAMS.get('page'));
        if(page) {
            queryString += "page="+page;
        } else {
            queryString += "page=1";
        }
        var txtFilter = $('#txtFilter').val();
        if((typeof txtFilter !== "undefined") && (txtFilter.trim() !== "")) {
            queryString += "&txtFilter="+txtFilter;
        }
        if(queryString === "?") {
            window.location.href = new URL(MAIN_URL);
        } else {
            window.location.href = new URL(`${MAIN_URL}${queryString}`);
        }
    });
});
