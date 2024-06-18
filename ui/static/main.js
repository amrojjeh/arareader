htmx.onLoad(function() {
  const selected = document.querySelector("[data-selected-segment]")
  const substituteButtons = document.querySelectorAll("[data-substitute-button]")
  const inputField = document.querySelector("[data-input]")
  const submitBtn = document.querySelector("#submit")
  const answer = document.querySelector("[data-form-answer]")

  if (answer) {
    answer.set = function(value) {
      this.value = value.trim().toLowerCase()
      if (this.value == "") {
        submitBtn.classList.remove("btn--primary")
        submitBtn.classList.add("btn--disabled")
        submitBtn.setAttribute("disabled", "")
      } else {
        submitBtn.classList.add("btn--primary")
        submitBtn.classList.remove("btn--disabled")
        submitBtn.removeAttribute("disabled")
      }
    }.bind(answer)
    answer.addEventListener("input", function(e) {
      console.log("ans changed")
    })
  }

  if (inputField) {
    inputField.addEventListener("input", function(e) {
      answer.set(e.target.value)
    })
  }

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
      answer.set(sub)
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
  const myExcerpt = document.body.querySelector("#excerpt")
  const responseExcerpt = response.querySelector("#excerpt")
  if (responseExcerpt) {
    const responseRefs = responseExcerpt.querySelectorAll("span")
    const refs = myExcerpt.querySelectorAll("span")
    for (let x = 0; x < refs.length; x++) {
      const ref = refs[x]
      ref.className = responseRefs[x].className
      ref.removeAttribute("data-selected-segment")
      if (responseRefs[x].hasAttribute("data-selected-segment")) {
        ref.setAttribute("data-selected-segment", "")
      }
      if (ref.children.length == 0) {
        ref.innerText = responseRefs[x].innerText
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
