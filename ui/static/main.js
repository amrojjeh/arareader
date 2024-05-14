const selected = document.querySelector("[data-selected-segment]")
const substituteButtons = document.querySelectorAll("[data-substitute-button]")

substituteButtons.forEach(x => {
  x.addEventListener("click", function(e) {
    substituteButtons.forEach(b => {
      b.classList.remove("button--selected")
      b.removeAttribute("data-selected-button")
    })
    const sub = e.target.getAttribute("data-substitute-button")
    selected.innerText = sub
    e.target.classList.add("button--selected")
    e.target.setAttribute("data-selected-button", "")
  })
})
