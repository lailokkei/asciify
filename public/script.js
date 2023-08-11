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

let form = document.getElementById("options");

function sendOptions() {

    let options = {
        CharSetName: form.charset.value,
        Invert: form.invert.checked,
        ScaleWidth: Number(form.scale.value),
    };

    console.log(options);
    ws.send(JSON.stringify(options));
}

let scaleRange = document.getElementById("scale-range")
let scaleInput = form.scale

scaleRange.addEventListener("input", updateScaleInput);
scaleInput.addEventListener("change", updateScaleRange);

function updateScaleRange() {
    scaleRange.value = scaleInput.value;
}

function updateScaleInput() {
    scaleInput.value = scaleRange.value;
}

function copyText() {
    let text = document.getElementById("output");
    navigator.clipboard.writeText(text.innerText);
}

function updateFont() {
    let selector = document.getElementsByName("font");
    let text = document.getElementById("output");

    text.style.fontFamily = selector[0].value;
}
