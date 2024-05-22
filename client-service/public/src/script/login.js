document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('login__section_form');

  form.addEventListener('submit', (event) => {
    event.preventDefault();

    const email = document.getElementById('login__section_form_input_email').value;
    const password = document.getElementById('login__section_form_input_password').value;

    firebase.auth().signInWithEmailAndPassword(email, password)
      .then((userCredential) => {
        // Signed in
        const user = userCredential.user;
        console.log(userCredential);
        console.log('User logged in:', user);
      })
      .catch((error) => {
        const errorCode = error.code;
        const errorMessage = error.message;
        console.error('Login error:', errorMessage);
      });
  });
});
