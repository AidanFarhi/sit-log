
const newEventForm = document.getElementById('new-event-form')
const addEventButton = document.getElementById('add-button')
const addButtonIcon = document.getElementById('add-button-icon')
const minusButtonIcon = document.getElementById('minus-button-icon')
const createEventButton = document.getElementById('create-event-button')
const alertBox = document.getElementById('alertBox')

addEventButton.addEventListener('click', () => {
    newEventForm.classList.toggle('hidden')
    addButtonIcon.classList.toggle('hidden')
    minusButtonIcon.classList.toggle('hidden')
})

createEventButton.addEventListener('click', () => {
    newEventForm.classList.toggle('hidden')
    addButtonIcon.classList.toggle('hidden')
    minusButtonIcon.classList.toggle('hidden')
    showAlert()
})

function showAlert() {
    alertBox.classList.remove('hidden')
    alertBox.classList.add('visible')
    setTimeout(() => {
        alertBox.classList.remove('visible')
        alertBox.classList.add('hidden')
    }, 3000)
}