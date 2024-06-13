htmx.onLoad(function() {
  const selected = document.querySelector("[data-selected-segment]")
  const substituteButtons = document.querySelectorAll("[data-substitute-button]")

  substituteButtons.forEach(x => {
    x.addEventListener("click", function(e) {
      substituteButtons.forEach(b => {
        b.classList.remove("btn--primary")
        b.removeAttribute("data-selected-button")
      })
      const sub = e.target.getAttribute("data-substitute-button")
      selected.innerText = sub
      e.target.classList.add("btn--primary")
      e.target.setAttribute("data-selected-button", "")
      const answer = document.querySelector("[data-form-answer]")
      answer.value = sub
    })
  })
})
