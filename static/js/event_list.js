
const newEventForm = document.getElementById('new-event-form')

const addEventButton = document.getElementById('add-button')
addEventButton.addEventListener('click', () => {
    newEventForm.classList.toggle('hidden')
})
