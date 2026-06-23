<!-- src/routes/login/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import kemkesLogo from '$lib/assets/kemkes.png';

  let username = $state('');
  let password = $state('');
  let errorMessage = $state('');
  let isLoading = $state(false);

  onMount(() => {
    // If already logged in, skip login page
    if (localStorage.getItem('isLoggedIn') === 'true') {
      goto('/dashboard');
    }
  });

  const handleLogin = async (e: SubmitEvent) => {
    e.preventDefault();
    errorMessage = '';
    
    if (!username || !password) {
      errorMessage = 'Username dan Password wajib diisi.';
      return;
    }

    isLoading = true;

    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 800));

    if (username === '12345' && password === '12345') {
      localStorage.setItem('isLoggedIn', 'true');
      localStorage.setItem('username', 'dr. Eka Pratama');
      localStorage.setItem('faskes', 'RSUD Labuang Baji');
      localStorage.setItem('role', 'Dokter Spesialis');
      goto('/dashboard');
    } else {
      errorMessage = 'Kredensial salah. Gunakan username "12345" dan password "12345".';
      isLoading = false;
    }
  };
</script>

<svelte:head>
  <title>Login | Rujukan Medis Modern</title>
</svelte:head>

<div class="auth-container">
  <div class="logo-area">
    <div class="logo-glow"></div>
    <img src={kemkesLogo} alt="Logo Kemenkes" class="logo-icon" />
    <h1>RUJUKAN<span>MEDIS</span></h1>
    <p class="subtitle">Sistem Rujukan Antar Fasilitas Kesehatan Terintegrasi</p>
  </div>

  <div class="glass-panel auth-card animate-fade-in">
    <div class="card-header">
      <h2>Masuk ke Akun</h2>
      <p>Gunakan kredensial yang terdaftar pada sistem faskes</p>
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

    <form onsubmit={handleLogin} class="auth-form">
      <div class="form-group">
        <label class="form-label" for="username">Username / NIP / NIK</label>
        <div class="input-wrapper">
          <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
          <input 
            type="text" 
            id="username" 
            class="form-input" 
            placeholder="Masukkan username (12345)" 
            bind:value={username}
            disabled={isLoading}
          />
        </div>
      </div>

      <div class="form-group">
        <label class="form-label" for="password">Kata Sandi</label>
        <div class="input-wrapper">
          <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
            <path d="M7 11V7a5 5 0 0 1 10 0v4" />
          </svg>
          <input 
            type="password" 
            id="password" 
            class="form-input" 
            placeholder="Masukkan kata sandi (12345)" 
            bind:value={password}
            disabled={isLoading}
          />
        </div>
      </div>

      <button type="submit" class="btn btn-primary w-full" disabled={isLoading}>
        {#if isLoading}
          <div class="btn-spinner"></div>
          <span>Menghubungkan...</span>
        {:else}
          <span>Masuk Sistem</span>
          <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2.5">
            <line x1="5" y1="12" x2="19" y2="12"></line>
            <polyline points="12 5 19 12 12 19"></polyline>
          </svg>
        {/if}
      </button>
    </form>

    <div class="card-footer">
      <p>Belum memiliki akun? <a href="/register">Daftar Faskes Baru</a></p>
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
    background: radial-gradient(circle, rgba(46, 196, 177, 0.3) 0%, transparent 70%);
    top: -30px;
    left: 50%;
    transform: translateX(-50%);
    filter: blur(10px);
    z-index: -1;
  }

  .logo-icon {
    width: 72px;
    height: 72px;
    object-fit: contain;
    margin-bottom: 0.75rem;
    filter: drop-shadow(0 0 8px rgba(46, 196, 177, 0.35));
    animation: pulse-glow 3s infinite ease-in-out;
  }

  .logo-area h1 {
    font-size: 1.8rem;
    font-weight: 800;
    letter-spacing: 0.1em;
    background: linear-gradient(135deg, #ffffff 0%, #8DC63F 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .logo-area h1 {
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
    max-width: 480px;
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
    line-height: 1.4;
    animation: fadeIn 0.3s ease;
  }

  .alert-error svg {
    flex-shrink: 0;
  }

  .auth-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .input-wrapper {
    position: relative;
    display: flex;
    align-items: center;
  }

  .input-icon {
    position: absolute;
    left: 1rem;
    width: 20px;
    height: 20px;
    color: var(--text-muted);
    pointer-events: none;
    transition: color var(--transition-fast);
  }

  .input-wrapper .form-input {
    width: 100%;
    padding-left: 3rem;
  }

  .input-wrapper:focus-within .input-icon {
    color: var(--color-primary);
  }

  .btn-primary {
    background: linear-gradient(135deg, #2EC4B1 0%, #239B90 100%);
    color: #ffffff;
    font-weight: 700;
    box-shadow: 0 4px 15px rgba(46, 196, 177, 0.35);
    border: none;
    transition: all var(--transition-normal);
  }

  .btn-primary:hover {
    background: linear-gradient(135deg, #239B90 0%, #2EC4B1 100%);
    box-shadow: 0 6px 20px rgba(46, 196, 177, 0.5);
    transform: translateY(-2px);
  }

  .w-full {
    width: 100%;
  }

  .btn-spinner {
    width: 18px;
    height: 18px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: #ffffff;
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
