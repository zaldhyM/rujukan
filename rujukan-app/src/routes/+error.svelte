<div id="wrapper">
  <div class="text-center">
    <div class="error mx-auto" data-text="404">404</div>
    <p class="lead mb-5">Page Not Found</p>
  </div>
</div>

<style>
  html,
  body {
    margin: 0;
    height: 100%;
  }

  #wrapper {
    display: flex;
    justify-content: center; /* horizontal center */
    align-items: center; /* vertical center */
    height: 100vh; /* full layar */
    width: 100vw;
    background: linear-gradient(
      to bottom,
      #4f86f7,
      #1e40af
    ); /* contoh gradasi */
  }

  .error {
    font-size: 8rem;
    font-weight: bold;
    color: white;
  }

  .lead {
    font-size: 1.5rem;
    color: white;
  }
</style>
