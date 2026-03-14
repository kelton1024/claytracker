
const hitBtn = document.getElementById("hit");
const missBtn = document.getElementById("miss");
const backBtn = document.getElementById("back");
const submitBtn = document.getElementById("submit");
const station_number = document.getElementById("station_number");
const prevStation = document.getElementById("prev_station");
const nextStation = document.getElementById("next_station");
const score = document.getElementById("results");
const scoreChildren = score.children;
let total = 0;
let column = 0;
let row = 0;
let stationNumber = 1;
scoreChildren[row].children[column].classList.add("bg-blue-700")

let currentSelected = null;
let prevValue = null;

hitBtn.addEventListener("click", hit);
missBtn.addEventListener("click", miss);
backBtn.addEventListener("click", back);
submitBtn.addEventListener("click", submit);
prevStation.addEventListener("click", () => {
    if(stationNumber == 1) return;
    stationNumber--;
    station_number.innerHTML = stationNumber
    // TODO: Go to backend and get scores
})
nextStation.addEventListener("click", () => {
    if(stationNumber == 20) return;
    stationNumber++;
    station_number.innerHTML = stationNumber
    // TODO: Go to backend and get scores
})
// TODO: Move this to an edit file
score.addEventListener("dblclick", edit)
document.addEventListener("click", unfocus)

function hit(score){
    recordScore(true)
}

function miss(){
    recordScore(false)
}

// TODO: This needs refactored so that the javascript has the array and not the html 
function recordScore(score){
    scoreChildren[row].children[column].innerHTML = score ? 1 : 0 // TODO: Remove ternary operator
    scoreChildren[row].children[column].classList.remove("bg-blue-700")

    total++;
    if(total >= 8){
        hitBtn.removeEventListener("click", hit)
        missBtn.removeEventListener("click", miss)
        return
    }
    column = total % 2;
    row = Math.floor((total) / 2)
    scoreChildren[row].children[column].classList.add("bg-blue-700")
}

function back(){
    if(total <= 0){
        return
    }
    
    if(total >= 8){
        hitBtn.addEventListener("click", hit);
        missBtn.addEventListener("click", miss);
    }

    scoreChildren[row].children[column].classList.remove("bg-blue-700")

    total--;
    column = total % 2;
    row = Math.floor((total) / 2)
    scoreChildren[row].children[column].classList.add("bg-blue-700")
}

// TODO: Move this to an edit file
function edit(event){
    if(event.target && event.target.nodeName == "TD"){
        currentSelected = event.target
        prevValue = event.target.innerHTML
        event.target.innerHTML = '<input id="focus" class="w-20">'
        const inputField = document.getElementById("focus")
        inputField.focus();
    }
}

function unfocus(){
    if(!currentSelected){
        return
    }
    const inputField = document.getElementById("focus")
    if(inputField == null || inputField.value == ""){
        currentSelected.innerHTML = prevValue
        return
    }
    currentSelected.innerHTML = inputField.value;
    currentSelected = null;

}

async function submit(){
    let scores = 0
    for(i = 0; i < 4; i++){
        for(j = 0; j < 2; j++){
            value = scoreChildren[i].children[j].innerHTML
            if(value == "NA"){
                continue
            }

            // OR and then bit shift to generate an integer
            scores = scores | parseInt(value)
            scores = scores << 1
        }
    }
    // Ideally we wouldn't have this, but it's hard to tell how many clays a station will have
    // so we'll just do this for now
    scores = scores >> 1

    const url = '/add_score'
    try {
        const response = await fetch(url, {
            method: "POST",
            headers: {"Content-type": "application/json"},
            body: JSON.stringify({course: stationNumber, score: scores})
        })
        if(!response.ok) {
            throw new Error(`response status ${response.status}`)
        }
        const result = await response.json()
        console.log(result)
    }
    catch(error){
        console.error(error)
    }
    console.log("here")
}
