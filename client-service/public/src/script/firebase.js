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
firebase.auth().onAuthStateChanged((user) => {
  const authHeaderApp = document.getElementById('header__auth_auth_app');
  const authHeaderLogout = document.getElementById('header__auth_auth_logout');
  const unauthHeaderSignin = document.getElementById('header__auth_unauth_login');
  const unauthHeaderSignup = document.getElementById('header__auth_unauth_signup');

  if (user) {
    unauthHeaderSignin.style.display = 'none';
    unauthHeaderSignup.style.display = 'none';
    authHeaderApp.style.display = 'block';
    authHeaderLogout.style.display = 'block';

    localStorage.setItem('accessToken', user.multiFactor.user.accessToken);
    localStorage.setItem('email', user.multiFactor.user.email);
    localStorage.setItem('uid', user.multiFactor.user.uid);
  } else {
    authHeaderApp.style.display = 'none';
    authHeaderLogout.style.display = 'none';
    unauthHeaderSignin.style.display = 'block';
    unauthHeaderSignup.style.display = 'block';

    localStorage.removeItem('accessToken');
    localStorage.removeItem('email');
    localStorage.removeItem('uid');
    console.log(window.location.href);
    // window.location.href = '/login';
  }
});
