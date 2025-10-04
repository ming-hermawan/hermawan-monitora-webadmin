function initSelect(id, val, vals, initialMarkup) {
    var markup = initialMarkup;
    for(let n in vals) {
        const optionVal = vals[n];
        const selected = (optionVal == val)? " selected": "";
        markup += "<option value=\"" + optionVal + "\"" + selected + ">" + optionVal + "</option>";
    }
    $("#"+id)
    .find('option')
    .remove()
    .end()
    .append(markup)
    .val(val)
    ;
}
