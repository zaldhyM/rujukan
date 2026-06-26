<!-- src/routes/register/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import kemkesLogo from '$lib/assets/kemkes.png';

  let nama = $state('');
  let username = $state('');
  let nik = $state('');
  let faskes = $state('');
  let password = $state('');
  let confirmPassword = $state('');
  let showPassword = $state(false);
  let showConfirm = $state(false);

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
    await new Promise(resolve => setTimeout(resolve, 1000));

    successMessage = 'Registrasi Faskes berhasil! Mengalihkan ke halaman login...';
    setTimeout(() => { goto('/login'); }, 1500);
  };
</script>

<svelte:head>
  <title>Daftar Faskes | Rujukan Medis</title>
</svelte:head>

<div class="page-bg">
  <!-- Decorative waves -->
  <svg class="wave-tl" viewBox="0 0 300 300" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
    <path d="M-50 200 Q50 100 150 150 T350 100" stroke="#14B8A6" stroke-width="1.5" fill="none" opacity="0.35"/>
    <path d="M-80 240 Q30 130 140 180 T380 130" stroke="#14B8A6" stroke-width="1" fill="none" opacity="0.25"/>
    <path d="M-30 160 Q80 60 180 110 T400 60" stroke="#A3E635" stroke-width="1" fill="none" opacity="0.2"/>
    <path d="M-60 280 Q60 160 170 220 T420 170" stroke="#14B8A6" stroke-width="0.8" fill="none" opacity="0.2"/>
  </svg>
  <svg class="wave-br" viewBox="0 0 300 300" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
    <path d="M350 100 Q250 200 150 150 T-50 200" stroke="#A3E635" stroke-width="1.5" fill="none" opacity="0.3"/>
    <path d="M380 130 Q270 240 160 180 T-80 240" stroke="#14B8A6" stroke-width="1" fill="none" opacity="0.2"/>
    <path d="M400 60 Q280 160 180 110 T-30 160" stroke="#A3E635" stroke-width="1" fill="none" opacity="0.18"/>
  </svg>

  <div class="reg-center">
    <!-- Logo -->
    <div class="logo-block">
      <img src={kemkesLogo} alt="Logo Kemenkes" class="kemkes-logo" />
      <span class="kemkes-label">Kemenkes</span>
    </div>

    <!-- Heading -->
    <div class="heading-block">
      <h1>Daftar Faskes Baru</h1>
      <p>Lengkapi data berikut untuk mendapatkan akses sistem rujukan</p>
    </div>

    <!-- Card -->
    <div class="reg-card">
      {#if errorMessage}
        <div class="alert-error">
          <svg viewBox="0 0 24 24" fill="none" width="16" height="16" stroke="currentColor" stroke-width="2" style="flex-shrink:0">
            <circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/>
          </svg>
          <span>{errorMessage}</span>
        </div>
      {/if}

      {#if successMessage}
        <div class="alert-success">
          <svg viewBox="0 0 24 24" fill="none" width="16" height="16" stroke="currentColor" stroke-width="2" style="flex-shrink:0">
            <polyline points="20 6 9 17 4 12"/>
          </svg>
          <span>{successMessage}</span>
        </div>
      {/if}

      <form onsubmit={handleRegister}>
        <div class="form-row">
          <div class="field">
            <label for="nama">Nama Lengkap &amp; Gelar</label>
            <div class="input-wrap">
              <svg class="field-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/>
              </svg>
              <input type="text" id="nama" placeholder="Contoh: dr. Ahmad Fauzi" bind:value={nama} disabled={isLoading} />
            </div>
          </div>

          <div class="field">
            <label for="nik">NIK / No. KTP</label>
            <div class="input-wrap">
              <svg class="field-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <rect x="2" y="5" width="20" height="14" rx="2"/><line x1="2" y1="10" x2="22" y2="10"/>
              </svg>
              <input type="text" id="nik" placeholder="16 digit NIK" bind:value={nik} disabled={isLoading} maxlength="16" />
            </div>
          </div>
        </div>

        <div class="form-row">
          <div class="field">
            <label for="faskes">Instansi / Faskes</label>
            <div class="input-wrap">
              <svg class="field-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/>
              </svg>
              <input type="text" id="faskes" placeholder="Puskesmas / RS / Klinik" bind:value={faskes} disabled={isLoading} />
            </div>
          </div>

          <div class="field">
            <label for="username">Username</label>
            <div class="input-wrap">
              <svg class="field-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-3.92 7.94"/>
              </svg>
              <input type="text" id="username" placeholder="Buat username unik" bind:value={username} disabled={isLoading} />
            </div>
          </div>
        </div>

        <div class="form-row">
          <div class="field">
            <label for="password">Kata Sandi</label>
            <div class="input-wrap">
              <svg class="field-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/>
              </svg>
              <input
                type={showPassword ? 'text' : 'password'}
                id="password"
                placeholder="Min. 6 karakter"
                bind:value={password}
                disabled={isLoading}
              />
              <button type="button" class="eye-btn" onclick={() => showPassword = !showPassword}>
                {#if showPassword}
                  <svg viewBox="0 0 24 24" fill="none" width="17" height="17" stroke="currentColor" stroke-width="1.8">
                    <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/>
                    <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/>
                    <line x1="1" y1="1" x2="23" y2="23"/>
                  </svg>
                {:else}
                  <svg viewBox="0 0 24 24" fill="none" width="17" height="17" stroke="currentColor" stroke-width="1.8">
                    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/>
                  </svg>
                {/if}
              </button>
            </div>
          </div>

          <div class="field">
            <label for="confirmPassword">Konfirmasi Kata Sandi</label>
            <div class="input-wrap">
              <svg class="field-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                <rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/>
              </svg>
              <input
                type={showConfirm ? 'text' : 'password'}
                id="confirmPassword"
                placeholder="Ulangi kata sandi"
                bind:value={confirmPassword}
                disabled={isLoading}
              />
              <button type="button" class="eye-btn" onclick={() => showConfirm = !showConfirm}>
                {#if showConfirm}
                  <svg viewBox="0 0 24 24" fill="none" width="17" height="17" stroke="currentColor" stroke-width="1.8">
                    <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/>
                    <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/>
                    <line x1="1" y1="1" x2="23" y2="23"/>
                  </svg>
                {:else}
                  <svg viewBox="0 0 24 24" fill="none" width="17" height="17" stroke="currentColor" stroke-width="1.8">
                    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/>
                  </svg>
                {/if}
              </button>
            </div>
          </div>
        </div>

        <button type="submit" class="submit-btn" disabled={isLoading}>
          {#if isLoading}
            <div class="spinner"></div>
            <span>Mendaftarkan...</span>
          {:else}
            Daftar Akun
          {/if}
        </button>
      </form>

      <div class="card-footer">
        Sudah memiliki akun? <a href="/login">Masuk ke Sistem</a>
      </div>
    </div>

    <!-- Security badge -->
    <div class="security-badge">
      <div class="shield-wrap">
        <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2">
          <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
        </svg>
      </div>
      <div>
        <p class="badge-title">Sistem ini aman dan terlindungi</p>
        <p class="badge-desc">Data Anda dienkripsi dan kami menjaga kerahasiaan informasi pribadi Anda.</p>
      </div>
    </div>
  </div>
</div>

<style>
  .page-bg {
    min-height: 100vh;
    background: #F0F4F8;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem 1rem;
    position: relative;
    overflow: hidden;
  }

  .wave-tl {
    position: absolute;
    top: 0; left: 0;
    width: 380px; height: 380px;
    pointer-events: none;
  }

  .wave-br {
    position: absolute;
    bottom: 0; right: 0;
    width: 380px; height: 380px;
    pointer-events: none;
  }

  .reg-center {
    position: relative;
    z-index: 1;
    width: 100%;
    max-width: 640px;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .logo-block {
    display: flex;
    align-items: center;
    gap: 0.6rem;
    margin-bottom: 1.25rem;
  }

  .kemkes-logo { width: 36px; height: 36px; object-fit: contain; }

  .kemkes-label {
    font-family: var(--font-title);
    font-size: 1.2rem;
    font-weight: 700;
    color: var(--text-primary);
  }

  .heading-block {
    text-align: center;
    margin-bottom: 1.5rem;
  }

  .heading-block h1 {
    font-size: 1.45rem;
    font-weight: 700;
    color: var(--text-primary);
    margin-bottom: 0.3rem;
  }

  .heading-block p {
    font-size: 0.875rem;
    color: var(--text-muted);
  }

  .reg-card {
    width: 100%;
    background: #FFFFFF;
    border: 1px solid #E5E7EB;
    border-radius: 16px;
    padding: 2rem 1.75rem;
    box-shadow: 0 2px 16px rgba(0, 0, 0, 0.07);
  }

  .alert-error {
    background: #FEF2F2;
    border: 1px solid #FECACA;
    border-radius: 8px;
    color: #DC2626;
    padding: 0.7rem 0.875rem;
    margin-bottom: 1.25rem;
    display: flex;
    align-items: flex-start;
    gap: 0.5rem;
    font-size: 0.825rem;
    line-height: 1.45;
  }

  .alert-success {
    background: #F0FDF4;
    border: 1px solid #BBF7D0;
    border-radius: 8px;
    color: #16A34A;
    padding: 0.7rem 0.875rem;
    margin-bottom: 1.25rem;
    display: flex;
    align-items: flex-start;
    gap: 0.5rem;
    font-size: 0.825rem;
  }

  .form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
  }

  .field {
    display: flex;
    flex-direction: column;
    gap: 0.4rem;
    margin-bottom: 0.875rem;
  }

  .field label {
    font-size: 0.82rem;
    font-weight: 600;
    color: var(--text-secondary);
    font-family: var(--font-title);
  }

  .input-wrap {
    position: relative;
    display: flex;
    align-items: center;
  }

  .field-ic {
    position: absolute;
    left: 0.875rem;
    width: 16px; height: 16px;
    color: #9CA3AF;
    pointer-events: none;
    transition: color 0.15s;
  }

  .input-wrap:focus-within .field-ic { color: #14B8A6; }

  .input-wrap input {
    width: 100%;
    height: 44px;
    padding: 0 1rem 0 2.6rem;
    background: #FFFFFF;
    border: 1px solid #E5E7EB;
    border-radius: 8px;
    font-family: var(--font-body);
    font-size: 0.875rem;
    color: #374151;
    transition: border-color 0.15s, box-shadow 0.15s;
  }

  .input-wrap input::placeholder { color: #9CA3AF; }

  .input-wrap input:focus {
    outline: none;
    border-color: #14B8A6;
    box-shadow: 0 0 0 3px rgba(20, 184, 166, 0.1);
  }

  .input-wrap input:disabled { opacity: 0.6; cursor: not-allowed; }

  .eye-btn {
    position: absolute;
    right: 0.75rem;
    background: none;
    border: none;
    color: #9CA3AF;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
    transition: color 0.15s;
  }
  .eye-btn:hover { color: #6B7280; }

  .submit-btn {
    width: 100%;
    height: 46px;
    background: #14B8A6;
    color: #FFFFFF;
    border: none;
    border-radius: 8px;
    font-family: var(--font-title);
    font-weight: 700;
    font-size: 0.975rem;
    cursor: pointer;
    transition: background 0.15s, box-shadow 0.15s;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    margin-top: 0.25rem;
  }

  .submit-btn:hover:not(:disabled) {
    background: #0F9488;
    box-shadow: 0 2px 10px rgba(20, 184, 166, 0.3);
  }

  .submit-btn:disabled { opacity: 0.7; cursor: not-allowed; }

  .spinner {
    width: 16px; height: 16px;
    border: 2px solid rgba(255,255,255,0.35);
    border-top-color: #fff;
    border-radius: 50%;
    animation: spin 0.75s linear infinite;
  }

  @keyframes spin { to { transform: rotate(360deg); } }

  .card-footer {
    text-align: center;
    margin-top: 1.25rem;
    font-size: 0.875rem;
    color: #6B7280;
    border-top: 1px solid #F3F4F6;
    padding-top: 1.25rem;
  }

  .card-footer a {
    color: #14B8A6;
    font-weight: 600;
    transition: color 0.15s;
  }

  .card-footer a:hover { color: #0F9488; }

  .security-badge {
    width: 100%;
    background: #FFFFFF;
    border: 1px solid #E5E7EB;
    border-radius: 12px;
    padding: 1rem 1.25rem;
    margin-top: 0.875rem;
    display: flex;
    align-items: flex-start;
    gap: 1rem;
    box-shadow: 0 1px 4px rgba(0,0,0,0.05);
  }

  .shield-wrap {
    width: 40px; height: 40px; min-width: 40px;
    border-radius: 10px;
    background: #CCFBF1;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #14B8A6;
  }

  .badge-title { font-size: 0.85rem; font-weight: 600; color: #374151; margin-bottom: 0.2rem; }
  .badge-desc { font-size: 0.775rem; color: #6B7280; line-height: 1.5; }

  @media (max-width: 600px) {
    .form-row { grid-template-columns: 1fr; }
    .reg-card { padding: 1.5rem 1.25rem; }
  }
</style>
