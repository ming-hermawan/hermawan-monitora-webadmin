function adjustHeight() {
    let documenth = $(document).height();
    let divMainContentH = $("#divMainContent").height();
    if ( documenth > 250 ) {
        var padding_top = (documenth - divMainContentH) / 2;
        $("#divInnerBody").css("padding-top", padding_top+"px")
    }
}


$(document).ready(function() {
    adjustHeight();
    $(window).resize(function(){
        adjustHeight();
    });
});
