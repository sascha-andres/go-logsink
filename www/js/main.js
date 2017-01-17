window.onload = function () {
    var conn;
    
    var log = document.getElementById("log");
    
    function appendLog(line) {
        var item = document.createElement("div");
        item.innerText = line
        var doScroll = log.scrollTop === log.scrollHeight - log.clientHeight;
        log.insertBefore(item, log.firstChild);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }
    
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://{{$}}/ws");
        conn.onclose = function (evt) {
            appendLog("<b>Connection closed.</b>");
        }
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            appendLog(evt.data);
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};