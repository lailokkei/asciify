let ws = new WebSocket("ws://" + window.location.host + "/connect");
ws.onmessage = function(e) {
    document.getElementById("output").innerHTML = e.data;
};

ws.onopen = function() {
    sendOptions();
};

let options = {
    CharSetName: "simple",
    Invert: false,
    ScaleWidth: 20,
};

// function sendImage() {
//     ws.send
// }
//
function sendOptions() {
    ws.send(JSON.stringify(getOptions()));
}

function getOptions() {
    let form = document.getElementById("options");

    let options = {
        CharSetName: form.charset.value,
        Invert: form.invert.checked,
        ScaleWidth: Number(form.scale.value),
    };

    console.log(options);
    return options
}

// url = window.location.href + "/api";
//
// let value = document.querySelector("#scale-output");
// let range = document.querySelector("#scale-range");
//
// range.addEventListener("input", console.log("input"));
//
// function updateValue() {
//     console.log("jdasflksajf");
//     value.value = range.value;
// }
//
// async function submitImage() {
//     let image = document.getElementById("image").files[0];
//     let r = await fetch("/api/image", { method: "POST", body: "sakdljf", image });
//
//     r.text().then(function(result) {
//         document.getElementById("output").innerHTML = result;
//     });
// }
