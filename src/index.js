// this is just a method created for testing the read file
async function fetchJsonFile(){
    // this res variable corresponds to the response of the try and catch handling exceptions
    try{
        // tries to fetch the JSON data that is stored in the localhost server
        res = await fetch('http://127.0.0.1:5500/src/json/algorithms.json');
        if (!res.ok){
            throw new Error("HTTPS error! Status: ${response.status}");
        }
        return res.json();
    } catch(error){
        throw new Error("HTTP request couldn't fetch from algorithms.json file AND/OR file doesn't exist anymore.");
    }
}
function inputDropButtonData(){
    const data = fetchJsonFile();
    const btnDropDown = document.querySelector("#btnDropDown");
    setDropButtonList(data,btnDropDown);
}
function setDropButtonList(data,btnDropDown){
    var i = data[0]["algorithms"].length;
    var j = 0;
    while(j <= i){
        createObjectOption(btnDropDown, data[0]["algorithms"][j]);
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





