document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('header__auth_auth_logout');

  form.addEventListener('click', (event) => {
    event.preventDefault();
    console.log(event);

    console.log(firebase.auth().signOut());
    window.location.href = '/login';
  });
});
