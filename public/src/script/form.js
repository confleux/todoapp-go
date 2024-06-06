document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('main__section_form');
  const messages = document.getElementById('main__section_messages');

  const url = window.location.href.includes('localhost') ? 'ws://localhost:3000' : 'wss://web-confleux.onrender.com';
  const socket = new WebSocket(`${url}/ws`);
  socket.addEventListener("message", (event) => {
    const data = JSON.parse(event.data);
    createMessageDiv(data.text, data.email, messages);
  });

  form.addEventListener('submit', (event) => {
      event.preventDefault();

      const value = document.getElementById('main__section_form_input').value;
      const email = localStorage.getItem("email");
      if (!value.length) {
        return;
      }
      if (email === null) {
        window.alert("You have to be signed in!");
        return;
      }
      form.reset();

      socket.send(JSON.stringify({
        text: value,
        email,
      }));
  });

  // let storedMessages = localStorage.getItem('messages');
  // if (storedMessages) {
  //   storedMessages = JSON.parse(storedMessages);
  //   for (const message of storedMessages) {
  //     createMessageDiv(message.value, message.timestamp, messages);
  //   }
  // }
});

function createMessageDiv(value, email, container) {
    const div = document.createElement('div');
    div.classList.add('main__section_messages_element');

    const valueText = document.createElement('p');
    valueText.textContent = value;

    const timeText = document.createElement('p');
    // timeText.textContent = (new Date(timestamp)).toLocaleString('en-GB');
    timeText.textContent = email;
    timeText.classList.add('main__section_messages_element_timestamp');

    div.appendChild(valueText);
    div.appendChild(timeText);

    container.appendChild(div);
}
