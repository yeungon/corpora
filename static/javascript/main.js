// const search_button = document.querySelector("#search_btm");
// const result_grid = document.querySelector("#result_grid");
// const searching_loading = document.querySelector("#searching_loading");


// search_button.onclick = function () {  
//   console.log("click");
//   console.log(result_grid);
//   console.log(searching_loading)
//   result_grid.style.display = "none"
//   searching_loading.style.dispay = "none"
// };


const search_input = document.querySelector("#input_search");
const search_loading = document.querySelector("#searching_loading");
const search_btm = document.querySelector("#search_btm");
const result_grid = document.querySelector("#result_grid");


window.onload = function () {
    console.log(search_input)    
}
search_input.oninput = function () {
    const value = search_input.value
     if (value.length >= 1) {
        console.log("ready")
        search_btm.disabled = false
    }
}

search_btm.onclick = function () {
    console.log("onclicked button")
    const value = search_input.value
    if(value.length <=0){
        search_btm.disabled = true
        search_input.placeholder = "Hi, bạn chưa nhập ngữ đoạn!"        
        return
    }

    search_btm.disabled = false
    console.log(search_loading)
    console.log(result_grid)
    // result_grid.style.display = "none"
}