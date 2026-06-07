<script>
  import { register } from '$lib/auth';

  let name = $state('');
  let username = $state('');
  let password = $state('');
  let confirmPassword = $state('');
  let errorMessage = $state('');
  let successMessage = $state('');

  /**
   * @param {SubmitEvent} e
   */
  function handleRegister(e) {
    e.preventDefault();
    errorMessage = '';
    successMessage = '';

    if (!name || !username || !password || !confirmPassword) {
      errorMessage = 'Semua kolom wajib diisi!';
      return;
    }

    if (password !== confirmPassword) {
      errorMessage = 'Konfirmasi password tidak cocok!';
      return;
    }

    if (password.length < 5) {
      errorMessage = 'Password minimal terdiri dari 5 karakter!';
      return;
    }

    const res = register(name, username, password);
    if (res.success) {
      successMessage = 'Registrasi berhasil! Mengalihkan ke halaman login...';
      setTimeout(() => {
        window.location.href = '/login';
      }, 1500);
    } else {
      errorMessage = res.message || 'Registrasi gagal';
    }
  }
</script>

<svelte:head>
  <title>Registrasi - Aplikasi Rujukan Pasien</title>
</svelte:head>

<div class="auth-container min-vh-100 d-flex align-items-center justify-content-center position-relative overflow-hidden py-5">
  <!-- Glowing Background Blobs -->
  <div class="ambient-glow" style="top: -10%; left: -5%; width: 500px; height: 500px; background: radial-gradient(circle, rgba(79, 70, 229, 0.25) 0%, rgba(79, 70, 229, 0) 75%);"></div>
  <div class="ambient-glow" style="bottom: -10%; right: -5%; width: 500px; height: 500px; background: radial-gradient(circle, rgba(14, 165, 233, 0.2) 0%, rgba(14, 165, 233, 0) 75%);"></div>

  <div class="row justify-content-center w-100 z-1 px-3">
    <div class="col-xl-5 col-lg-6 col-md-8 col-sm-10">
      <div class="glass-panel p-4 p-sm-5 animate-fade-in">
        <div class="text-center mb-4">
          <div class="icon-badge d-inline-flex align-items-center justify-content-center mb-3">
            <i class="fas fa-user-plus text-primary" style="font-size: 24px;"></i>
          </div>
          <h2 class="h3 font-weight-bold mb-1">Daftar Akun Baru</h2>
          <p class="text-muted small">Buat akun petugas untuk mengakses SISRUTE</p>
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

        <form class="mt-4" onsubmit={handleRegister}>
          <div class="mb-3">
            <label for="name" class="form-label">Nama Lengkap & Gelar</label>
            <div class="input-icon-wrapper">
              <i class="fas fa-user-md input-icon"></i>
              <input
                type="text"
                class="form-control ps-5"
                id="name"
                placeholder="Contoh: dr. John Doe, Sp.PD"
                bind:value={name}
              />
            </div>
          </div>
          
          <div class="mb-3">
            <label for="username" class="form-label">Username / NIP</label>
            <div class="input-icon-wrapper">
              <i class="fas fa-id-badge input-icon"></i>
              <input
                type="text"
                class="form-control ps-5"
                id="username"
                placeholder="Masukkan Username baru"
                bind:value={username}
              />
            </div>
          </div>
          
          <div class="row">
            <div class="col-sm-6 mb-3">
              <label for="password" class="form-label">Password</label>
              <div class="input-icon-wrapper">
                <i class="fas fa-lock input-icon"></i>
                <input
                  type="password"
                  class="form-control ps-5"
                  id="password"
                  placeholder="Password"
                  bind:value={password}
                />
              </div>
            </div>
            
            <div class="col-sm-6 mb-4">
              <label for="confirmPassword" class="form-label">Ulangi Password</label>
              <div class="input-icon-wrapper">
                <i class="fas fa-lock input-icon"></i>
                <input
                  type="password"
                  class="form-control ps-5"
                  id="confirmPassword"
                  placeholder="Ulangi Password"
                  bind:value={confirmPassword}
                />
              </div>
            </div>
          </div>
          
          <button
            type="submit"
            class="btn-modern-primary w-100"
          >
            Daftar Akun Baru
          </button>
        </form>
        
        <hr class="my-4" style="border-color: rgba(255,255,255,0.08);" />
        
        <div class="text-center">
          <span class="small text-muted">Sudah memiliki akun? </span>
          <a class="small font-weight-bold text-decoration-none text-primary" href="/login">Masuk di sini</a>
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
</style>
