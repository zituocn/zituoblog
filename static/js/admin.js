var addnews = function () {
    var i = 1;
    var cid = $("#cid");
    var title = $("#title");
    var text = editor.html();

    if (cid.val() == '0') {
        alert("没有选择分类...");
        cid.focus();
        i = 2;
        return false;
    }
    if (title.val().length == 0 || title.val() == "") {
        alert("没有填写标题...");
        title.focus();
        i = 2;
        return false;
    }
    if (text.length == 0 || text == "") {
        alert("没有填写内容...");
        editor.focus();
        i = 2;
        return false;
    }

    if (i == 1)
        $("#enter").disabled = true;
}

/*
取cookies
*/
var getCookie = function (c_name) {
    if (document.cookie.length > 0) {
        c_start = document.cookie.indexOf(c_name + "=");
        if (c_start != -1) {
            c_start = c_start + c_name.length + 1;
            c_end = document.cookie.indexOf(";", c_start);
            if (c_end == -1) c_end = document.cookie.length;
            return decodeURIComponent(document.cookie.substring(c_start, c_end));
        }
    }
    return "";
}