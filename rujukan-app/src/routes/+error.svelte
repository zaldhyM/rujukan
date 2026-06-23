<!-- src/routes/+error.svelte -->
<script lang="ts">
  import { page } from '$app/state';
  import { goto } from '$app/navigation';
  import kemkesLogo from '$lib/assets/kemkes.png';

  let status = $derived(page.status);
  let message = $derived(page.error?.message ?? '');

  let is404 = $derived(status === 404);
  let is500 = $derived(status === 500);
</script>

<svelte:head>
  <title>{status} | Rujukan Medis</title>
</svelte:head>

<div class="error-container">
  <div class="glow-orb orb-1"></div>
  <div class="glow-orb orb-2"></div>

  <div class="error-card animate-fade-in">
    <!-- Header -->
    <div class="card-header">
      <img src={kemkesLogo} alt="Logo Kemenkes" class="logo" />
      <span class="app-name">RUJUKAN<span>MEDIS</span></span>
    </div>

    <!-- Status Code -->
    <div class="status-code" class:is-500={is500}>
      {status}
    </div>

    <!-- Icon -->
    <div class="error-icon-wrap" class:is-500={is500}>
      {#if is404}
        <svg viewBox="0 0 24 24" fill="none" width="40" height="40" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"></circle>
          <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
          <line x1="11" y1="8" x2="11" y2="11"></line>
          <line x1="11" y1="14" x2="11.01" y2="14"></line>
        </svg>
      {:else}
        <svg viewBox="0 0 24 24" fill="none" width="40" height="40" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
          <line x1="12" y1="9" x2="12" y2="13"></line>
          <line x1="12" y1="17" x2="12.01" y2="17"></line>
        </svg>
      {/if}
    </div>

    <!-- Title & Description -->
    <div class="error-body">
      {#if is404}
        <h1>Halaman Tidak Ditemukan</h1>
        <p>Halaman yang Anda cari tidak tersedia atau telah dipindahkan. Periksa kembali URL yang dimasukkan.</p>
      {:else if is500}
        <h1>Kesalahan Server</h1>
        <p>Terjadi gangguan pada server. Tim teknis kami sedang menangani masalah ini. Silakan coba beberapa saat lagi.</p>
      {:else}
        <h1>Terjadi Kesalahan</h1>
        <p>{message || 'Terjadi kesalahan yang tidak terduga. Silakan coba kembali.'}</p>
      {/if}
    </div>

    {#if message && !is404}
      <div class="error-detail">
        <code>{message}</code>
      </div>
    {/if}

    <!-- Actions -->
    <div class="error-actions">
      <button class="btn-primary" onclick={() => goto('/')}>
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2.5">
          <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
          <polyline points="9 22 9 12 15 12 15 22"></polyline>
        </svg>
        <span>Ke Halaman Utama</span>
      </button>

      <button class="btn-outline" onclick={() => history.back()}>
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2.5">
          <line x1="19" y1="12" x2="5" y2="12"></line>
          <polyline points="12 19 5 12 12 5"></polyline>
        </svg>
        <span>Kembali</span>
      </button>
    </div>
  </div>

  <p class="footer-text">&copy; 2026 Rujukan Medis Terintegrasi &mdash; Kementerian Kesehatan RI</p>
</div>

<style>
  .error-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
    padding: 2rem 1.5rem;
    background: #ffffff;
    position: relative;
    overflow: hidden;
    gap: 2rem;
  }

  .glow-orb {
    position: absolute;
    border-radius: 50%;
    filter: blur(120px);
    opacity: 0.07;
    pointer-events: none;
  }
  .orb-1 {
    width: 400px;
    height: 400px;
    background: #2EC4B1;
    top: -80px;
    left: -80px;
  }
  .orb-2 {
    width: 350px;
    height: 350px;
    background: #AACC00;
    bottom: -60px;
    right: -60px;
  }

  .error-card {
    background: #ffffff;
    border: 1px solid rgba(0, 0, 0, 0.08);
    border-radius: 18px;
    box-shadow: 0 8px 40px rgba(0, 0, 0, 0.1);
    padding: 3rem 3.5rem;
    max-width: 520px;
    width: 100%;
    z-index: 1;
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
  }

  /* Header */
  .card-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }

  .logo {
    width: 36px;
    height: 36px;
    object-fit: contain;
  }

  .app-name {
    font-size: 1.1rem;
    font-weight: 800;
    letter-spacing: 0.05em;
    color: var(--text-primary);
    font-family: var(--font-title);
  }

  .app-name span {
    color: #2EC4B1;
  }

  /* Status Code */
  .status-code {
    font-size: 6rem;
    font-weight: 900;
    line-height: 1;
    font-family: var(--font-title);
    background: linear-gradient(135deg, #2EC4B1 0%, #AACC00 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
    letter-spacing: -0.04em;
  }

  .status-code.is-500 {
    background: linear-gradient(135deg, #EF4444 0%, #F59E0B 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  /* Icon */
  .error-icon-wrap {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    background: rgba(46, 196, 177, 0.1);
    border: 1px solid rgba(46, 196, 177, 0.25);
    color: #2EC4B1;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .error-icon-wrap.is-500 {
    background: rgba(239, 68, 68, 0.08);
    border-color: rgba(239, 68, 68, 0.2);
    color: #EF4444;
  }

  /* Body text */
  .error-body h1 {
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--text-primary);
    margin-bottom: 0.75rem;
    font-family: var(--font-title);
  }

  .error-body p {
    font-size: 0.95rem;
    color: var(--text-secondary);
    line-height: 1.7;
    max-width: 380px;
  }

  /* Error detail */
  .error-detail {
    background: rgba(0, 0, 0, 0.04);
    border: 1px solid rgba(0, 0, 0, 0.08);
    border-radius: 8px;
    padding: 0.75rem 1rem;
    width: 100%;
  }

  .error-detail code {
    font-size: 0.8rem;
    color: var(--text-muted);
    font-family: monospace;
    word-break: break-all;
  }

  /* Actions */
  .error-actions {
    display: flex;
    gap: 1rem;
    width: 100%;
    flex-wrap: wrap;
  }

  .btn-primary {
    flex: 1;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    padding: 0.85rem 1.5rem;
    border-radius: 8px;
    font-family: var(--font-title);
    font-weight: 700;
    font-size: 0.9rem;
    cursor: pointer;
    border: none;
    background: linear-gradient(135deg, #2EC4B1 0%, #239B90 100%);
    color: #ffffff;
    box-shadow: 0 4px 15px rgba(46, 196, 177, 0.35);
    transition: all 0.2s ease;
  }

  .btn-primary:hover {
    background: linear-gradient(135deg, #239B90 0%, #2EC4B1 100%);
    box-shadow: 0 6px 20px rgba(46, 196, 177, 0.5);
    transform: translateY(-2px);
  }

  .btn-outline {
    flex: 1;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    padding: 0.85rem 1.5rem;
    border-radius: 8px;
    font-family: var(--font-title);
    font-weight: 600;
    font-size: 0.9rem;
    cursor: pointer;
    background: transparent;
    border: 1.5px solid rgba(0, 0, 0, 0.15);
    color: var(--text-primary);
    transition: all 0.2s ease;
  }

  .btn-outline:hover {
    background: rgba(0, 0, 0, 0.04);
    border-color: #2EC4B1;
    color: #2EC4B1;
  }

  .footer-text {
    font-size: 0.8rem;
    color: var(--text-muted);
    z-index: 1;
  }

  @media (max-width: 540px) {
    .error-card {
      padding: 2rem 1.5rem;
    }

    .status-code {
      font-size: 4.5rem;
    }

    .error-actions {
      flex-direction: column;
    }
  }
</style>
