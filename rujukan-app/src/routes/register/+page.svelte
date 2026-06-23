<!-- src/routes/register/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  let nama = $state('');
  let username = $state('');
  let nik = $state('');
  let faskes = $state('');
  let password = $state('');
  let confirmPassword = $state('');

  let errorMessage = $state('');
  let successMessage = $state('');
  let isLoading = $state(false);

  onMount(() => {
    if (localStorage.getItem('isLoggedIn') === 'true') {
      goto('/dashboard');
    }
  });

  const handleRegister = async (e: SubmitEvent) => {
    e.preventDefault();
    errorMessage = '';
    successMessage = '';

    if (!nama || !username || !nik || !faskes || !password || !confirmPassword) {
      errorMessage = 'Semua field wajib diisi.';
      return;
    }

    if (password !== confirmPassword) {
      errorMessage = 'Konfirmasi password tidak cocok.';
      return;
    }

    isLoading = true;

    // Simulate database write
    await new Promise(resolve => setTimeout(resolve, 1000));

    // For simulation, save to localstorage and redirect
    successMessage = 'Registrasi Faskes berhasil! Mengalihkan ke halaman login...';
    
    setTimeout(() => {
      goto('/login');
    }, 1500);
  };
</script>

<svelte:head>
  <title>Daftar Faskes Baru | Rujukan Medis Modern</title>
</svelte:head>

<div class="auth-container">
  <div class="logo-area">
    <div class="logo-glow"></div>
    <svg class="logo-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M19 10.5V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V10.5M12 3V17M12 3L7.5 7.5M12 3L16.5 7.5" stroke="url(#logoGrad)" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
      <defs>
        <linearGradient id="logoGrad" x1="5" y1="3" x2="19" y2="21" gradientUnits="userSpaceOnUse">
          <stop stop-color="#00C853" />
          <stop offset="1" stop-color="#8DC63F" />
        </linearGradient>
      </defs>
    </svg>
    <h1>RUJUKAN<span>MEDIS</span></h1>
    <p class="subtitle">Pendaftaran Fasilitas Kesehatan Baru & Tenaga Medis</p>
  </div>

  <div class="glass-panel auth-card animate-fade-in">
    <div class="card-header">
      <h2>Pendaftaran Faskes</h2>
      <p>Lengkapi berkas untuk mendapat akses sistem rujukan</p>
    </div>

    {#if errorMessage}
      <div class="alert-error">
        <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="12" y1="8" x2="12" y2="12"></line>
          <line x1="12" y1="16" x2="12.01" y2="16"></line>
        </svg>
        <span>{errorMessage}</span>
      </div>
    {/if}

    {#if successMessage}
      <div class="alert-success">
        <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="currentColor" stroke-width="2">
          <polyline points="20 6 9 17 4 12"></polyline>
        </svg>
        <span>{successMessage}</span>
      </div>
    {/if}

    <form onsubmit={handleRegister} class="auth-form">
      <div class="form-row">
        <div class="form-group flex-1">
          <label class="form-label" for="nama">Nama Lengkap & Gelar</label>
          <input 
            type="text" 
            id="nama" 
            class="form-input" 
            placeholder="Contoh: dr. Ahmad" 
            bind:value={nama}
            disabled={isLoading}
          />
        </div>
        <div class="form-group flex-1">
          <label class="form-label" for="nik">NIK / No. KTP</label>
          <input 
            type="text" 
            id="nik" 
            class="form-input" 
            placeholder="Nomor NIK 16 digit" 
            bind:value={nik}
            disabled={isLoading}
          />
        </div>
      </div>

      <div class="form-row">
        <div class="form-group flex-1">
          <label class="form-label" for="faskes">Instansi / Faskes</label>
          <input 
            type="text" 
            id="faskes" 
            class="form-input" 
            placeholder="Contoh: Puskesmas Kassi Kassi" 
            bind:value={faskes}
            disabled={isLoading}
          />
        </div>
        <div class="form-group flex-1">
          <label class="form-label" for="username">Username</label>
          <input 
            type="text" 
            id="username" 
            class="form-input" 
            placeholder="Buat username baru" 
            bind:value={username}
            disabled={isLoading}
          />
        </div>
      </div>

      <div class="form-row">
        <div class="form-group flex-1">
          <label class="form-label" for="password">Kata Sandi</label>
          <input 
            type="password" 
            id="password" 
            class="form-input" 
            placeholder="Minimal 6 karakter" 
            bind:value={password}
            disabled={isLoading}
          />
        </div>
        <div class="form-group flex-1">
          <label class="form-label" for="confirmPassword">Konfirmasi Kata Sandi</label>
          <input 
            type="password" 
            id="confirmPassword" 
            class="form-input" 
            placeholder="Ketik ulang kata sandi" 
            bind:value={confirmPassword}
            disabled={isLoading}
          />
        </div>
      </div>

      <button type="submit" class="btn btn-primary w-full" disabled={isLoading}>
        {#if isLoading}
          <div class="btn-spinner"></div>
          <span>Mendaftarkan Faskes...</span>
        {:else}
          <span>Daftar Akun</span>
          <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2.5">
            <path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2" />
            <circle cx="8.5" cy="7" r="4" />
            <line x1="20" y1="8" x2="20" y2="14" />
            <line x1="17" y1="11" x2="23" y2="11" />
          </svg>
        {/if}
      </button>
    </form>

    <div class="card-footer">
      <p>Sudah memiliki akun? <a href="/login">Masuk ke Sistem</a></p>
    </div>
  </div>
</div>

<style>
  .auth-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
    padding: 2rem 1.5rem;
    position: relative;
    z-index: 1;
  }

  .logo-area {
    text-align: center;
    margin-bottom: 2rem;
    position: relative;
  }

  .logo-glow {
    position: absolute;
    width: 120px;
    height: 120px;
    background: radial-gradient(circle, rgba(0, 200, 83, 0.4) 0%, transparent 70%);
    top: -30px;
    left: 50%;
    transform: translateX(-50%);
    filter: blur(10px);
    z-index: -1;
  }

  .logo-icon {
    width: 60px;
    height: 60px;
    margin-bottom: 0.5rem;
    filter: drop-shadow(0 0 8px rgba(0, 200, 83, 0.4));
  }

  .logo-area h1 {
    font-size: 1.8rem;
    font-weight: 800;
    letter-spacing: 0.1em;
    background: none;
    -webkit-text-fill-color: var(--text-primary);
    color: var(--text-primary);
  }

  .logo-area h1 span {
    color: var(--color-primary);
    -webkit-text-fill-color: var(--color-primary);
    font-weight: 800;
    background: none;
  }

  .subtitle {
    color: var(--text-secondary);
    font-size: 0.875rem;
    margin-top: 0.5rem;
    max-width: 400px;
  }

  .auth-card {
    width: 100%;
    max-width: 680px;
    padding: 2.5rem;
    border-radius: var(--border-radius-lg);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    animation: fadeIn 0.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }

  .card-header {
    text-align: center;
    margin-bottom: 2rem;
  }

  .card-header h2 {
    font-size: 1.5rem;
    color: var(--text-primary);
    margin-bottom: 0.5rem;
  }

  .card-header p {
    color: var(--text-secondary);
    font-size: 0.9rem;
  }

  .alert-error {
    background: rgba(255, 117, 140, 0.1);
    border: 1px solid rgba(255, 117, 140, 0.25);
    border-radius: var(--border-radius-sm);
    color: #ff758c;
    padding: 0.85rem 1rem;
    margin-bottom: 1.5rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    font-size: 0.875rem;
    animation: fadeIn 0.3s ease;
  }

  .alert-success {
    background: rgba(0, 245, 160, 0.1);
    border: 1px solid rgba(0, 245, 160, 0.25);
    border-radius: var(--border-radius-sm);
    color: var(--color-success);
    padding: 0.85rem 1rem;
    margin-bottom: 1.5rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    font-size: 0.875rem;
    animation: fadeIn 0.3s ease;
  }

  .auth-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .form-row {
    display: flex;
    gap: 1.25rem;
  }

  .flex-1 {
    flex: 1;
  }

  @media (max-width: 640px) {
    .form-row {
      flex-direction: column;
      gap: 1.25rem;
    }
  }

  .btn-primary {
    background: linear-gradient(135deg, #00C853 0%, #43A047 100%);
    color: #071a0d;
    font-weight: 700;
    box-shadow: 0 4px 15px rgba(0, 200, 83, 0.3);
    border: none;
    transition: all var(--transition-normal);
  }

  .btn-primary:hover {
    background: linear-gradient(135deg, #43A047 0%, #00C853 100%);
    box-shadow: 0 6px 20px rgba(0, 200, 83, 0.5);
    transform: translateY(-2px);
  }

  .w-full {
    width: 100%;
  }

  .btn-spinner {
    width: 18px;
    height: 18px;
    border: 2px solid rgba(7, 26, 13, 0.15);
    border-top-color: #071a0d;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
  }

  .card-footer {
    text-align: center;
    margin-top: 2rem;
    font-size: 0.9rem;
    color: var(--text-secondary);
  }

  .card-footer a {
    color: var(--color-primary);
    font-weight: 600;
    transition: color var(--transition-fast);
  }

  .card-footer a:hover {
    color: var(--color-primary-hover);
    text-decoration: underline;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }
</style>
