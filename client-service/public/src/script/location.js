document.addEventListener("DOMContentLoaded", () => {
  const navMenu = document.getElementById("header__nav");
  const menuItems = navMenu.getElementsByTagName("a");

  for (let i = 0; i < menuItems.length; ++i) {
    if (document.location.href === menuItems[i].href) {
      menuItems[i].className = "header__nav_link_enabled";
    }
  }
});
