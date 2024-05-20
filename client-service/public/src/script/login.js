// import { initializeApp } from "https://www.gstatic.com/firebasejs/10.12.0/firebase-app.js";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyDsGv_1JXQD-cuwPSLu5aqGmLowaZtNK3s",
  authDomain: "web-confleux.firebaseapp.com",
  projectId: "web-confleux",
  storageBucket: "web-confleux.appspot.com",
  messagingSenderId: "62743178164",
  appId: "1:62743178164:web:40bde3b14704c9b7c6a0e1"
};

const app = firebase.initializeApp(firebaseConfig);

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
        console.log(userCredential)
        console.log('User logged in:', user);
      })
      .catch((error) => {
        const errorCode = error.code;
        const errorMessage = error.message;
        console.error('Login error:', errorMessage);
      });

  });

});
