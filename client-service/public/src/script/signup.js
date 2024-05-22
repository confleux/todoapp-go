document.addEventListener('DOMContentLoaded', () => {
  const url = window.location.href.includes('localhost') ? 'http://localhost:3000' : 'https://web-confleux.onrender.com';
  const form = document.getElementById('signup__section_form');

  form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const email = document.getElementById('signup__section_form_input_email').value;
    const password = document.getElementById('signup__section_form_input_password').value;

    const data = {
      email,
      password,
    };

    try {
      const response = await fetch(`${url}/api/signup`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        console.log(await response.json());
        window.alert('Signed up successfully');
        window.location.href = '/login';
      } else {
        window.alert(`Failed to sign up: ${await response.text()}`);
      }

    } catch (error) {
      console.log(`Failed to sign up: ${error}`);
    }
  });
});
