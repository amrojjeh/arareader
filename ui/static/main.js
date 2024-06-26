var dialog = undefined

htmx.onLoad(function() {
  const selectedRef = document.querySelector("[data-selected-segment]")
  const selectedQuestion = document.querySelector("[data-selected-question]")
  const substituteButtons = document.querySelectorAll("[data-substitute-button]")
  const inputField = document.querySelector("[data-input]")
  const submitBtn = document.querySelector("#submit")
  const answer = document.querySelector("[data-form-answer]")
  const excerpt = document.querySelector("#excerpt")
  dialog = document.querySelector("dialog")

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
      selectedRef.innerText = sub
      e.target.dataset.type = "primary"
      answer.set(sub)
    })
  })

  if (selectedRef) {
    selectedRef.scrollIntoView({
      "behavior": "smooth",
      "block": "center",
      "inline": "center",
    })
  }

  if (selectedQuestion) {
    selectedQuestion.scrollIntoView({
      "behavior": "smooth",
      "block": "center",
      "inline": "center",
    })
  }

  if (excerpt) {
    excerpt.addEventListener("htmx:oobBeforeSwap", updateExcerpt)
  }

  document.body.addEventListener("keydown", shortcut)
  document.querySelectorAll("[data-sidebar-toggle]").forEach(x => x.addEventListener("click", sidebarToggle))
  if (dialog) {
    // to fix this bug:
    // 1) Open dialog
    // 2) Select a question
    // 3) Press back
    // The dialog will be in a "open" state where the question is interactable and the dialog is open at the same time
    // Not sure if this is a browser bug
    dialog.close()

    dialog.addEventListener("click", clickoutSidebar)
    dialog.addEventListener("htmx:oobBeforeSwap", updateSidebar)
  }
})

function clickoutSidebar(e) {
  const dialog = document.querySelector("dialog")
  const dialogDimensions = dialog.getBoundingClientRect()
  if (
    e.clientX < dialogDimensions.left ||
    e.clientX > dialogDimensions.right ||
    e.clientY < dialogDimensions.top ||
    e.clientY > dialogDimensions.bottom
  ) {
    dialog.close()
  }
}

function sidebarToggle() {
  const sidebar = document.querySelector("[data-sidebar]")
  if (sidebar.open) {
    sidebar.close()
  } else {
    sidebar.showModal()
  }
}

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

function updateSidebar(evt) {
  evt.detail.shouldSwap = false
  const responseDialog = evt.detail.fragment.firstChild
  const mydialog = evt.currentTarget
  mydialog.replaceChild(responseDialog.firstChild, mydialog.firstChild)
  const selected = mydialog.querySelector("[data-selected-question]")
  console.log(selected)
  selected.scrollIntoView({
    "behavior": "smooth",
    "block": "center",
    "inline": "center",
  })
  htmx.process(mydialog)
}

function updateExcerpt(evt) {
  evt.detail.shouldSwap = false
  const response = evt.detail.fragment
  const myExcerpt = evt.currentTarget
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

window.addEventListener("resize", function(evnt) {
  if (window.innerWidth >= 1024 && dialog) { // equivalent to tailwind's "lg"
    dialog.close()
  }
})
