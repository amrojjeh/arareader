htmx.onLoad(function() {
  const selected = document.querySelector("[data-selected-segment]")
  const substituteButtons = document.querySelectorAll("[data-substitute-button]")

  substituteButtons.forEach(x => {
    x.addEventListener("click", function(e) {
      const submitBtn = document.querySelector("#submit")
      submitBtn.classList.add("btn--primary")
      submitBtn.classList.remove("btn--disabled")
      submitBtn.removeAttribute("disabled")
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

  const highlighted = document.querySelector(".highlight")
  highlighted.scrollIntoView({
    "behavior": "smooth",
    "block": "center",
    "inline": "center",
  })
})

document.querySelector("#target").addEventListener("htmx:beforeSwap", function(evt) {
  const response = new DOMParser().parseFromString(evt.detail.xhr.response, "text/html")
  const responseExcerpt = response.querySelector("#excerpt")
  if (responseExcerpt) {
    const responseRefs = responseExcerpt.querySelectorAll("span")
    const myExcerpt = document.body.querySelector("#excerpt")
    const refs = myExcerpt.querySelectorAll("span")
    for (let x = 0; x < refs.length; x++) {
      const ref = refs[x]
      ref.className = responseRefs[x].className
      ref.removeAttribute("data-selected-segment")
      if (responseRefs[x].hasAttribute("data-selected-segment")) {
        ref.setAttribute("data-selected-segment", "")
      }
    }
    const highlighted = document.querySelector(".highlight")
    highlighted.scrollIntoView({
      "behavior": "smooth",
      "block": "center",
      "inline": "center",
    })
  }
})
