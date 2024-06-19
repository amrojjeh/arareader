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
        submitBtn.dataset.type = "disabled"
        submitBtn.setAttribute("disabled", "")
      } else {
        submitBtn.dataset.type = "primary"
        submitBtn.removeAttribute("disabled")
      }
    }.bind(answer)
  }

  if (inputField) {
    inputField.addEventListener("input", function(e) {
      answer.set(e.target.value)
    })
  }

  substituteButtons.forEach(x => {
    x.addEventListener("click", function(e) {
      substituteButtons.forEach(b => {
        delete b.dataset.type
        delete b.dataset.selectedButton
      })
      const sub = e.target.dataset.substituteButton
      selected.innerText = sub
      e.target.dataset.type = "primary"
      answer.set(sub)
    })
  })

  selected.scrollIntoView({
    "behavior": "smooth",
    "block": "center",
    "inline": "center",
  })

  document.querySelector("#excerpt").addEventListener("htmx:oobBeforeSwap", function(evt) {
    evt.detail.shouldSwap = false
    const response = evt.detail.fragment
    const myExcerpt = document.body.querySelector("#excerpt")
    const responseExcerpt = response.querySelector("#excerpt")
    if (responseExcerpt) {
      const responseRefs = responseExcerpt.querySelectorAll("span")
      const refs = myExcerpt.querySelectorAll("span")
      for (let x = 0; x < refs.length; x++) {
        const ref = refs[x]
        ref.className = responseRefs[x].className
        delete ref.dataset.selectedSegment
        if (Object.hasOwn(responseRefs[x].dataset, "selectedSegment")) {
          ref.dataset.selectedSegment = ""
        }
        if (ref.children.length == 0) {
          ref.innerText = responseRefs[x].innerText
        }
      }
      const selected = document.querySelector("[data-selected-segment]")
      selected.scrollIntoView({
        "behavior": "smooth",
        "block": "center",
        "inline": "center",
      })
    }
  })
})

