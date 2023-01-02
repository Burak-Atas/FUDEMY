/**
 *  cookie kontrol
ü * 
 */
function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}
window.onload = function() {
    if (getCookie("token") == "") {
        console.log("token boş")
    } else {
        var name=document.getElementsByClassName('name')
        var elems = document.getElementsByClassName('signup');
        for (var
                i = 0; i < elems.length; i += 1) {
            elems[i].style.display = 'none';
        }
    }
};