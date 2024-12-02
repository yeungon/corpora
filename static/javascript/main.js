const search_input = document.querySelector("#input_search");
const search_loading = document.querySelector("#searching_loading");
const search_btm = document.querySelector("#search_btm");
const result_grid = document.querySelector("#result_grid");
const corpus_query_form = document.querySelector("#corpus_query");
const result_pagination = document.querySelector("#result_pagination");

document.addEventListener("DOMContentLoaded", function () {
  console.log(search_input);
});

window.onload = function () {
  console.log(search_input);
  console.log(search_loading);
};
search_input.oninput = function () {
  const value = search_input.value;
  if (value.length >= 1) {
    console.log("ready");
    search_btm.disabled = false;
  }
};

search_btm.addEventListener("click", function (event) {
  event.preventDefault();
  console.log("onclicked button");
  const value = search_input.value;
  if (value.length <= 0) {
    search_input.placeholder = "Hi, bạn chưa nhập ngữ đoạn!";
    search_btm.disabled = true;
    return;
  }
  search_btm.disabled = false;

  if (search_loading) {
    search_loading.style.display = "block";
  }
  if (result_pagination) {
    result_pagination.style.display = "none";
  }
  if (result_grid) {
    result_grid.style.display = "none";
  }
  corpus_query_form.submit();
});
