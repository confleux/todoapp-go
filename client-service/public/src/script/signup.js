document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('signup__section_form');

  form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const email = document.getElementById('signup__section_form_input_email').value;
    const password = document.getElementById('signup__section_form_input_password').value;

    const data = {
      email,
      password,
    };

    const response = await fetch('http://localhost:3000/api/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    const parsedResponse = await response.json();
    console.log(parsedResponse);
  });
});
