let ws = new WebSocket("ws://" + window.location.host + "/connect");

ws.onmessage = function(e) {
    document.getElementById("output").innerHTML = escapeHTML(e.data);
};

ws.onopen = function() {
    sendOptions();
};

function escapeHTML(unsafe) {
    return unsafe.replaceAll('&', '&amp;').replaceAll('<', '&lt;').replaceAll('>', '&gt;').replaceAll('"', '&quot;').replaceAll("'", '&#039;');
}

function updateImage() {
    let image = document.getElementById("image").files[0];
    ws.send(image);

    sendOptions()
}

function sendOptions() {
    let form = document.getElementById("options");

    let options = {
        CharSetName: form.charset.value,
        Invert: form.invert.checked,
        ScaleWidth: Number(form.scale.value),
    };

    console.log(options);
    ws.send(JSON.stringify(options));
}
