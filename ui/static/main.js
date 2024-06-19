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

  document.querySelector("#excerpt").addEventListener("htmx:oobBeforeSwap", updateExcerpt)

  document.body.addEventListener("keydown", shortcut)
})

function shortcut(ev) {
  const shortcuts = document.querySelectorAll("[data-shortcut]")
  shortcuts.forEach(function(elt) {
    if (ev.repeat) {
      return
    }
    const modifiers = ["ctrl", "alt", "shift"]
    var [modifier, key] = elt.dataset.shortcut.split(" ")
    if (modifiers.indexOf(modifier) == -1) {
      key = modifier
      modifier = true
    } else if (modifier == "ctrl") {
      modifier = ev.ctrlKey
    } else if (modifier == "alt") {
      modifier = ev.altKey
    } else if (modifier == "shift") {
      modifier = ev.shiftKey
    }

    if (modifier && ev.code === key) {
      elt.click()
      ev.preventDefault()
    }
  })
}

function updateExcerpt(evt) {
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
}
