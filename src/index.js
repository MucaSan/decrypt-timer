function inputDropButtonData(){

}
function setDropButtonList(){

}
// this is just a method created for testing the read file
async function displayJsonFile(){
    data = await fetch(
        'http://127.0.0.1:5500/src/json/algorithms.json'
    );
    alert(data);
}