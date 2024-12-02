const search_button = document.querySelector("#search_btm");
const result_grid = document.querySelector("#result_grid");
const searching_loading = document.querySelector("#searching_loading");


search_button.onclick = function () {  
  console.log("click");
  console.log(result_grid);
  console.log(searching_loading)
  result_grid.style.display = "none"
  searching_loading.style.dispay = "none"
};
