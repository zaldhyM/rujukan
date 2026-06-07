<script>
  import { onMount } from 'svelte';
  import { isLoggedIn } from '$lib/auth';

  let authenticated = $state(false);

  onMount(() => {
    authenticated = isLoggedIn();
  });

  function masuk() {
    if (authenticated) {
      window.location.href = "/admin";
    } else {
      window.location.href = "/login";
    }
  }

  function daftar() {
    window.location.href = "/register";
  }
</script>

<svelte:head>
  <title>Selamat Datang - SISRUTE LOKAL</title>
</svelte:head>

<div class="welcome-container min-vh-100 d-flex flex-column justify-content-between position-relative overflow-hidden">
  <!-- Glowing Background Blobs -->
  <div class="ambient-glow" style="top: -15%; left: -10%; width: 600px; height: 600px; background: radial-gradient(circle, rgba(92, 103, 242, 0.2) 0%, rgba(92, 103, 242, 0) 70%);"></div>
  <div class="ambient-glow" style="bottom: -15%; right: -10%; width: 600px; height: 600px; background: radial-gradient(circle, rgba(56, 189, 248, 0.15) 0%, rgba(56, 189, 248, 0) 70%);"></div>

  <!-- Header Bar -->
  <header class="welcome-header px-4 py-3 z-1">
    <div class="container d-flex align-items-center justify-content-between">
      <div class="brand d-flex align-items-center gap-2">
        <div class="brand-logo-icon">
          <i class="fas fa-hospital text-white"></i>
        </div>
        <div class="brand-text">
          <span class="brand-name">SISRUTE LOKAL</span>
          <span class="brand-tag">Rujukan Terintegrasi</span>
        </div>
      </div>
      <div>
        {#if authenticated}
          <button class="btn-modern-secondary px-3 py-2" onclick={masuk}>
            <i class="fas fa-chart-line me-2"></i> Ke Dasbor
          </button>
        {:else}
          <button class="btn-modern-secondary px-3 py-2" onclick={masuk}>
            <i class="fas fa-sign-in-alt me-2"></i> Masuk
          </button>
        {/if}
      </div>
    </div>
  </header>

  <!-- Hero Content -->
  <main class="welcome-main container d-flex align-items-center justify-content-center flex-grow-1 z-1 py-5">
    <div class="row align-items-center justify-content-center w-100">
      <div class="col-lg-6 mb-5 mb-lg-0 text-center text-lg-start pe-lg-5">
        <span class="badge bg-primary-light text-primary px-3 py-2 rounded-pill font-weight-bold mb-3 animate-fade-in">
          <i class="fas fa-shield-halved me-1"></i> Sistem Rujukan Medis Digital
        </span>
        <h1 class="welcome-title mb-3 animate-fade-in" style="font-size: 3rem; line-height: 1.2;">
          Proses Rujukan Medis <span class="text-gradient">Cepat & Aman</span>
        </h1>
        <p class="welcome-desc text-muted mb-4 animate-fade-in" style="font-size: 1.1rem; line-height: 1.6;">
          SISRUTE LOKAL memfasilitasi komunikasi dan transfer data rujukan antar instansi kesehatan secara real-time. Membantu tenaga medis memberikan pelayanan rujukan yang cepat, transparan, dan efisien demi keselamatan pasien.
        </p>
        <div class="d-flex flex-wrap gap-3 justify-content-center justify-content-lg-start animate-fade-in">
          <button class="btn-modern-primary py-3 px-4" onclick={masuk}>
            <i class="fas fa-arrow-right me-2"></i> 
            {authenticated ? 'Masuk ke Dasbor' : 'Masuk Aplikasi'}
          </button>
          {#if !authenticated}
            <button class="btn-modern-secondary py-3 px-4" onclick={daftar}>
              <i class="fas fa-user-plus me-2"></i> Daftar Akun
            </button>
          {/if}
        </div>
      </div>

      <div class="col-lg-5 col-md-8">
        <!-- Glassmorphic Visual Card -->
        <div class="glass-panel p-4 p-sm-5 text-center animate-fade-in position-relative">
          <div class="welcome-illustration mb-4">
            <div class="illustration-circle">
              <i class="fas fa-user-md text-primary" style="font-size: 3.5rem;"></i>
            </div>
            <div class="pulse-ring"></div>
            <div class="pulse-ring delay-1"></div>
          </div>
          
          <h4 class="font-weight-bold mb-2">Fasilitas Kesehatan Terhubung</h4>
          <p class="text-muted small mb-4">
            Puskesmas, Klinik Pratama, hingga Rumah Sakit Rujukan saling berintegrasi dalam satu jaringan rujukan elektronik yang aman.
          </p>

          <div class="row g-3 text-start">
            <div class="col-6">
              <div class="feature-item p-3">
                <i class="fas fa-bolt text-warning mb-2" style="font-size: 1.25rem;"></i>
                <h6 class="font-weight-bold mb-1">Real-time</h6>
                <span class="text-xs text-muted">Pengiriman instan ke RS rujukan</span>
              </div>
            </div>
            <div class="col-6">
              <div class="feature-item p-3">
                <i class="fas fa-file-prescription text-success mb-2" style="font-size: 1.25rem;"></i>
                <h6 class="font-weight-bold mb-1">Terintegrasi</h6>
                <span class="text-xs text-muted">Data klinis terkirim lengkap</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>

  <!-- Footer -->
  <footer class="welcome-footer py-4 z-1">
    <div class="container text-center">
      <span class="text-muted small">Copyright &copy; SISRUTE LOKAL 2026. Dikembangkan untuk efisiensi rujukan kesehatan.</span>
    </div>
  </footer>
</div>

<style>
  .welcome-container {
    background-color: #0f172a;
    color: #f8fafc;
  }

  .brand-logo-icon {
    width: 38px;
    height: 38px;
    background: linear-gradient(135deg, var(--primary) 0%, #818cf8 100%);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: var(--shadow-glow);
  }

  .brand-text {
    display: flex;
    flex-direction: column;
    line-height: 1.2;
  }

  .brand-name {
    font-weight: 800;
    font-size: 0.95rem;
    letter-spacing: 0.02em;
    color: #fff;
  }

  .brand-tag {
    font-size: 0.65rem;
    color: var(--text-light);
    font-weight: 600;
  }

  .text-gradient {
    background: linear-gradient(135deg, #a5b4fc 0%, var(--secondary) 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .welcome-title {
    font-weight: 800;
    color: #ffffff;
  }

  .welcome-desc {
    color: #cbd5e1 !important;
  }

  /* Illustration Pulsing Effect */
  .welcome-illustration {
    position: relative;
    width: 120px;
    height: 120px;
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .illustration-circle {
    width: 90px;
    height: 90px;
    background: rgba(92, 103, 242, 0.1);
    border: 1px solid rgba(92, 103, 242, 0.25);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 2;
    box-shadow: 0 8px 20px rgba(92, 103, 242, 0.15);
  }

  .pulse-ring {
    position: absolute;
    width: 90px;
    height: 90px;
    border: 1px solid rgba(92, 103, 242, 0.2);
    border-radius: 50%;
    animation: pulse 3s infinite linear;
    z-index: 1;
  }

  .pulse-ring.delay-1 {
    animation-delay: 1.5s;
  }

  @keyframes pulse {
    0% {
      transform: scale(1);
      opacity: 1;
    }
    100% {
      transform: scale(1.6);
      opacity: 0;
    }
  }

  .feature-item {
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid rgba(255, 255, 255, 0.05);
    border-radius: var(--radius-md);
    transition: var(--transition);
  }

  .feature-item:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(92, 103, 242, 0.2);
    transform: translateY(-2px);
  }

  .welcome-footer {
    border-top: 1px solid rgba(255, 255, 255, 0.05);
  }
</style>
