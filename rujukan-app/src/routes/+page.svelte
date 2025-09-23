<script>
  function masuk() {
    window.location.href = "/admin";
  }
</script>

<div class="hero">
  <div class="hero-box">
    <h1>Selamat Datang di Aplikasi Rujukan</h1>
    <p>
      Aplikasi ini membantu proses rujukan pasien menjadi lebih cepat, efisien,
      dan transparan. Dengan sistem digital, tenaga medis dapat mengirim dan
      menerima rujukan antar fasilitas kesehatan secara real-time.
    </p>
    <button class="btn" on:click={masuk}>Masuk</button>
  </div>
</div>

<style>
  .hero {
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #4e73df;
    color: white;
    text-align: center;
    padding: 2rem;
  }

  .hero-box {
    background: rgba(255, 255, 255, 0.1);
    padding: 3rem;
    border-radius: 20px;
    backdrop-filter: blur(10px);
    max-width: 700px;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
  }

  h1 {
    font-size: 2.8rem;
    margin-bottom: 1rem;
  }

  p {
    font-size: 1.2rem;
    margin-bottom: 2rem;
    line-height: 1.6;
  }

  .btn {
    padding: 0.8rem 2rem;
    border: none;
    border-radius: 50px;
    font-size: 1.1rem;
    font-weight: bold;
    background: #fff;
    color: #4e73df;
    cursor: pointer;
    transition: 0.3s;
  }

  .btn:hover {
    background: #f8f9fc;
    transform: translateY(-3px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  }
</style>
