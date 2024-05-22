if (window.location.href.includes('todo-app') && !localStorage.getItem('accessToken')) {
  window.location.href = '/login';
}

document.addEventListener('DOMContentLoaded', async () => {
  const form = document.getElementById('main__section_form');
  const messages = document.getElementById('main__section_messages');

  try {
    const response = await fetch('http://localhost:3000/api/todos', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
      },
    });

    if (response.ok) {
      const list = await response.json();
      for (const todo of list.todos) {
        createMessageDiv(todo.description, new Date(todo.createdAt).getTime(), todo.id, messages);
      }
    }
  } catch (error) {
    window.alert(`Error loading todos: ${error}`)
  }

  form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const value = document.getElementById('main__section_form_input').value;
    if (!value.length) {
      return;
    }
    form.reset();

    const data = {
      description: value,
    };

    try {
      const response = await fetch('http://localhost:3000/api/todos', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('accessToken')}`,
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        const list = await response.json();
        const todo = list;
        createMessageDiv(todo.description, new Date(todo.createdAt).getTime(), todo.id, messages);
      }
    } catch (error) {
      window.alert(`Error loading todos: ${error}`)
    }

  });
});

function createMessageDiv(value, timestamp, id, container) {
  const div = document.createElement('div');
  div.classList.add('main__section_messages_element');
  div.setAttribute("id", id);

  const valueText = document.createElement('p');
  valueText.textContent = value;

  const timeText = document.createElement('p');
  timeText.textContent = (new Date(timestamp)).toLocaleString('en-GB');
  timeText.classList.add('main__section_messages_element_timestamp');

  const deleteForm = document.createElement('form');
  deleteForm.classList.add('main__section_form');
  deleteForm.addEventListener('submit', async (event) => {
    event.preventDefault();

    try {
      const todoId = div.getAttribute('id');
      const response = await fetch(`http://localhost:3000/api/todos/${todoId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('accessToken')}`,
        },
      });

      if (response.ok) {
        const parsedResponse = await response.json();
        console.log(parsedResponse);
        div.remove();
      }
    } catch (error) {
      window.alert(`Error deleting todo: ${error}`)
    }
  });

  const deleteButton = document.createElement('button');
  deleteButton.innerHTML = 'Delete';
  deleteButton.classList.add('main__section_form_button');

  deleteForm.appendChild(deleteButton);

  div.appendChild(valueText);
  div.appendChild(timeText);
  div.appendChild(deleteForm);

  container.appendChild(div);
}
