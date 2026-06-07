<script>
  import { onMount } from 'svelte';
  import { login, isLoggedIn } from '$lib/auth';

  let username = $state('');
  let password = $state('');
  let errorMessage = $state('');
  let successMessage = $state('');

  onMount(() => {
    if (isLoggedIn()) {
      window.location.href = '/admin';
    }
  });

  /**
   * @param {SubmitEvent} e
   */
  function handleLogin(e) {
    e.preventDefault();
    errorMessage = '';
    
    if (!username || !password) {
      errorMessage = 'Username dan password wajib diisi!';
      return;
    }

    const res = login(username, password);
    if (res.success) {
      successMessage = 'Login Berhasil! Mengalihkan...';
      setTimeout(() => {
        window.location.href = '/admin';
      }, 1000);
    } else {
      errorMessage = res.message || 'Login gagal';
    }
  }
</script>

<svelte:head>
  <title>Masuk - Aplikasi Rujukan Pasien</title>
</svelte:head>

<div class="auth-container min-vh-100 d-flex align-items-center justify-content-center position-relative overflow-hidden">
  <!-- Glowing Background Blobs -->
  <div class="ambient-glow" style="top: -10%; left: -5%; width: 500px; height: 500px; background: radial-gradient(circle, rgba(79, 70, 229, 0.25) 0%, rgba(79, 70, 229, 0) 75%);"></div>
  <div class="ambient-glow" style="bottom: -10%; right: -5%; width: 500px; height: 500px; background: radial-gradient(circle, rgba(14, 165, 233, 0.2) 0%, rgba(14, 165, 233, 0) 75%);"></div>

  <div class="row justify-content-center w-100 z-1 px-3">
    <div class="col-xl-4 col-lg-5 col-md-7 col-sm-9">
      <div class="glass-panel p-4 p-sm-5 animate-fade-in">
        <div class="text-center mb-4">
          <div class="icon-badge d-inline-flex align-items-center justify-content-center mb-3">
            <i class="fas fa-hospital-user text-primary" style="font-size: 26px;"></i>
          </div>
          <h2 class="h3 font-weight-bold mb-1">SISRUTE LOKAL</h2>
          <p class="text-muted small">Sistem Informasi Rujukan Terintegrasi</p>
        </div>
        
        {#if errorMessage}
          <div class="alert alert-danger-modern text-center py-2 px-3 mb-3 small animate-fade-in" role="alert">
            <i class="fas fa-exclamation-circle mr-1"></i> {errorMessage}
          </div>
        {/if}

        {#if successMessage}
          <div class="alert alert-success-modern text-center py-2 px-3 mb-3 small animate-fade-in" role="alert">
            <i class="fas fa-check-circle mr-1"></i> {successMessage}
          </div>
        {/if}

        <form class="mt-4" onsubmit={handleLogin}>
          <div class="mb-3">
            <label for="username" class="form-label">Username / NIP</label>
            <div class="input-icon-wrapper">
              <i class="fas fa-user input-icon"></i>
              <input
                type="text"
                class="form-control ps-5"
                id="username"
                placeholder="Masukkan username (12345)"
                bind:value={username}
              />
            </div>
          </div>
          
          <div class="mb-4">
            <label for="password" class="form-label">Password</label>
            <div class="input-icon-wrapper">
              <i class="fas fa-lock input-icon"></i>
              <input
                type="password"
                class="form-control ps-5"
                id="password"
                placeholder="Masukkan password (12345)"
                bind:value={password}
              />
            </div>
          </div>
          
          <button
            type="submit"
            class="btn-modern-primary w-100"
          >
            Masuk ke Dasbor
          </button>
        </form>
        
        <div class="divider-text my-4">
          <span>ATAU</span>
        </div>
        
        <div class="text-center mb-4">
          <span class="small text-muted">Belum memiliki akun petugas? </span>
          <a class="small font-weight-bold text-decoration-none text-primary" href="/register">Buat Akun Baru</a>
        </div>
        
        <div class="demo-badge p-3 text-center">
          <span class="text-xs text-muted d-block">Gunakan Akun Demo untuk Uji Coba:</span>
          <code class="text-primary font-weight-bold small mt-1 d-inline-block">User: 12345 &nbsp;|&nbsp; Pass: 12345</code>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  .auth-container {
    background: radial-gradient(circle at 50% 50%, #1e1b4b 0%, #0f172a 100%);
  }
  
  .icon-badge {
    width: 60px;
    height: 60px;
    border-radius: 16px;
    background: rgba(79, 70, 229, 0.1);
    border: 1px solid rgba(79, 70, 229, 0.2);
    box-shadow: 0 8px 16px rgba(79, 70, 229, 0.1);
  }

  .input-icon-wrapper {
    position: relative;
  }

  .input-icon {
    position: absolute;
    left: 1rem;
    top: 50%;
    transform: translateY(-50%);
    color: var(--text-light);
    font-size: 0.9rem;
    transition: var(--transition);
  }

  .ps-5 {
    padding-left: 2.75rem !important;
  }

  .form-control:focus + .input-icon,
  .form-control:focus ~ .input-icon {
    color: var(--primary);
  }

  .alert-danger-modern {
    background-color: var(--danger-light);
    color: var(--danger-dark);
    border: 1px solid rgba(244, 63, 94, 0.2);
    border-radius: var(--radius-md);
    font-weight: 600;
  }

  .alert-success-modern {
    background-color: var(--success-light);
    color: var(--success-dark);
    border: 1px solid rgba(16, 185, 129, 0.2);
    border-radius: var(--radius-md);
    font-weight: 600;
  }

  .divider-text {
    display: flex;
    align-items: center;
    text-align: center;
    color: var(--text-light);
    font-size: 0.75rem;
    font-weight: 700;
  }

  .divider-text::before,
  .divider-text::after {
    content: '';
    flex: 1;
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  }

  .divider-text:not(:empty)::before {
    margin-right: .75em;
  }

  .divider-text:not(:empty)::after {
    margin-left: .75em;
  }

  .demo-badge {
    background: rgba(255, 255, 255, 0.03);
    border: 1px dashed rgba(255, 255, 255, 0.1);
    border-radius: var(--radius-md);
  }
</style>
