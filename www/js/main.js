var scrollingEnabled = true;

window.onload = function () {
    var conn;
    
    var log = document.getElementById("log");
    
    function appendLog(line) {
        var item = document.createElement("div");
        item.innerText = line
        log.appendChild(item);
        if (scrollingEnabled) {
            window.scrollTo(0,document.body.scrollHeight);
        }
    }
    
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://{{$}}/api/go-logsink/ws");
        conn.onclose = function (evt) {
            log.innerHTML = "<b>Connection closed.</b>";
        }
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            appendLog(evt.data);
        };
    } else {
        log.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    }
};

togglScrolling = function () {
    scrollingEnabled=!scrollingEnabled;
    if (scrollingEnabled) {
        document.getElementById("scrollToggler").innerText = "scrolling"
    } else {
        document.getElementById("scrollToggler").innerText = "paused"
    }
}