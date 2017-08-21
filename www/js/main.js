var scrollingEnabled = true;
var lineLimit = {{.Limit}};
var numberOfLines = 0;
var expr = "";

function isNumber(n) {
	return !isNaN(parseFloat(n)) && isFinite(n);
}

window.onload = function () {
	var conn;

	var log = document.getElementById("log");

	function appendLog(msg) {

		if (!msg.Line.match(expr)) {
			return;
		}
		var item = document.createElement("div");
		if ( msg.Priority === 0 ) {
			item.setAttribute("class", "normal")
		}
		if ( msg.Priority > 0 && msg.Priority <= 2) {
			item.setAttribute("class", "raised")
		}
		if ( msg.Priority > 2 && msg.Priority <= 5) {
			item.setAttribute("class", "medium")
		}
		if ( msg.Priority > 5 && msg.Priority <= 8) {
			item.setAttribute("class", "warning")
		}
		if ( msg.Priority > 8 ) {
			item.setAttribute("class", "critical")
		}
		if ( msg.Line === "") {
			item.innerText = " ";
		} else {
			item.innerText = msg.Line;
		}
		log.appendChild(item);
		if (scrollingEnabled) {
			window.scrollTo(0,document.body.scrollHeight);
		}
		if (lineLimit > 0) {
			numberOfLines++;
			while (numberOfLines > lineLimit) {
				log.removeChild(log.firstChild);
				numberOfLines--;
			}
		}
	}

	if (window["WebSocket"]) {
		conn = new WebSocket("{{.Scheme}}://{{.Host}}/api/go-logsink/ws");
		conn.onclose = function (evt) {
			log.innerHTML = "<b>Connection closed.</b>";
		}
		conn.onmessage = function (evt) {
			var json = eval( "(" + evt.data + ")" );
			appendLog(json);
		};
	} else {
		log.innerHTML = "<b>Your browser does not support WebSockets.</b>";
	}

	document.getElementById("limit").value = lineLimit;
	document.getElementById("limit").onblur = function (evt) {
		if (isNumber(document.getElementById("limit").value)) {
			lineLimit = document.getElementById("limit").value;
		}
	};
	document.getElementById("expr").onblur = function (evt) {
		expr = document.getElementById("expr").value;
	};
};

togglScrolling = function () {
	scrollingEnabled=!scrollingEnabled;
	if (scrollingEnabled) {
		document.getElementById("scrollToggler").innerText = "scrolling";
	} else {
		document.getElementById("scrollToggler").innerText = "paused";
	}
}
