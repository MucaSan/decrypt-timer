// this is just a method created for testing the read file
async function fetchJsonFile(){
    // this res variable corresponds to the response of the try and catch handling exceptions
    try{
        // tries to fetch the JSON data that is stored in the localhost server
        res = await fetch('http://127.0.0.1:5500/src/json/algorithms.json');
        if (!res.ok){
            throw new Error("HTTPS error! Status: ${res.status}");
        }
        return res.json();
    } catch(error){
        throw new Error("HTTP request couldn't fetch from algorithms.json file AND/OR file doesn't exist anymore.");
    }
}
async function inputDropButtonData(){
    // fetch the data, and use the await keyword to automatically make the promise interactable
    const data = await fetchJsonFile();
    // create DOM object for interacting with select button
    const btnDropDown = document.querySelector("#btnDropDown");
    console.log(data["algorithms"]);
    setDropButtonList(data,btnDropDown);
}
function setDropButtonList(data,btnDropDown){
    var i = data["algorithms"].length;
    var j = 0;
    while(j <= i){
        createObjectOption(btnDropDown, data["algorithms"][j]);
        j++;
    }
}
function createObjectOption(btnDropButton, optionName){
    var option = document.createElement('option');
    option.value = optionName;
    option.innerHTML = optionName; 
    btnDropButton.appendChild(option);
}
document.addEventListener("DOMContentLoaded", inputDropButtonData());





