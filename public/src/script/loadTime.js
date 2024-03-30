(function () {
  const startTime = performance.now();

  window.addEventListener("load", () => {
      const endTime = performance.now();

      const delta = endTime - startTime;

      document.getElementById("footer__loadTime_value").textContent = delta;
  });
})();