url = window.location.href+"/api"

generateText()

function generateText(){
    fetch("/api").then(function(r) {
        return r.text();
    }).then(function(data) {
        updateOutput(data)
        console.log("data: " + data);
    }).catch(function(err) {
        console.log('Fetch Error :-S', err);
    });
}

function updateOutput(text){
    document.getElementById("output").innerHTML = text

}

async function submitImage(){
    let image = document.getElementById("image").files[0];
    let r = await fetch("/api/image", {method : "POST", body : image})

    r.text().then(function(result){
        document.getElementById("output").innerHTML = result;
    });
}
