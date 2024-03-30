document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('main__section_form');
  const messages = document.getElementById('main__section_messages');

  form.addEventListener('submit', (event) => {
      event.preventDefault();

      const value = document.getElementById('main__section_form_input').value;
      if (!value.length) {
        return;
      }
      form.reset();

      const timestamp = Date.now();

      createMessageDiv(value, timestamp, messages);

      let storedMessages = localStorage.getItem('messages') ? localStorage.getItem('messages') : [];
      if (typeof storedMessages === 'string') {
        storedMessages = JSON.parse(storedMessages);
      }
      storedMessages.push({value, timestamp});
      localStorage.setItem('messages', JSON.stringify(storedMessages));
  });

  let storedMessages = localStorage.getItem('messages');
  if (storedMessages) {
    storedMessages = JSON.parse(storedMessages);
    for (const message of storedMessages) {
      createMessageDiv(message.value, message.timestamp, messages);
    }
  }
});

function createMessageDiv(value, timestamp, container) {
    const div = document.createElement('div');
    div.classList.add('main__section_messages_element');

    const valueText = document.createElement('p');
    valueText.textContent = value;

    const timeText = document.createElement('p');
    timeText.textContent = (new Date(timestamp)).toLocaleString('en-GB');
    timeText.classList.add('main__section_messages_element_timestamp');

    div.appendChild(valueText);
    div.appendChild(timeText);

    container.appendChild(div);
}