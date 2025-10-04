function adjustHeight() {
    let documenth = $(document).height();
    let divMainContentH = $("#divMainContent").height();
    if ( documenth > (divMainContentH + 36) ) {
        var padding_top = (documenth - divMainContentH - 68) / 2;
        $("#divInnerBody").css("padding-top", padding_top+"px")
    }
}


$(document).ready(function() {
    adjustHeight();
    $(window).resize(function(){
        adjustHeight();
    });
});
