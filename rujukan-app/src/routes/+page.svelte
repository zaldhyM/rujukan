<!-- src/routes/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  let isLoggedIn = $state(false);

  onMount(() => {
    isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
  });

  const handleCTA = () => {
    if (isLoggedIn) {
      goto('/dashboard');
    } else {
      goto('/login');
    }
  };
</script>

<svelte:head>
  <title>Rujukan Medis | Portal Utama</title>
</svelte:head>

<div class="landing-container animate-fade-in">
  <div class="split-layout">
    <!-- Layout 1: Brand & Access Portal -->
    <div class="glass-panel access-panel">
      <div class="logo-area">
        <svg class="logo-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M19 10.5V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V10.5M12 3V17M12 3L7.5 7.5M12 3L16.5 7.5" stroke="url(#logoGrad)" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
          <defs>
            <linearGradient id="logoGrad" x1="5" y1="3" x2="19" y2="21" gradientUnits="userSpaceOnUse">
              <stop stop-color="#00f2fe" />
              <stop offset="1" stop-color="#4facfe" />
            </linearGradient>
          </defs>
        </svg>
        <h1>RUJUKAN<span>MEDIS</span></h1>
        <span class="portal-badge">PINTU GERBANG UTAMA</span>
      </div>

      <div class="action-buttons">
        <button onclick={handleCTA} class="btn btn-primary w-full">
          {#if isLoggedIn}
            <span>Masuk Dashboard</span>
          {:else}
            <span>Masuk ke Sistem</span>
          {/if}
          <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2.5">
            <line x1="5" y1="12" x2="19" y2="12"></line>
            <polyline points="12 5 19 12 12 19"></polyline>
          </svg>
        </button>

        {#if !isLoggedIn}
          <a href="/register" class="btn btn-secondary w-full">Daftar Faskes Baru</a>
        {/if}
      </div>

      <div class="test-tip">
        <svg viewBox="0 0 24 24" fill="none" width="16" height="16" stroke="var(--color-primary)" stroke-width="2">
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="12" y1="16" x2="12" y2="12"></line>
          <line x1="12" y1="8" x2="12.01" y2="8"></line>
        </svg>
        <span>Akun Uji Coba: <strong>12345</strong> / <strong>12345</strong></span>
      </div>
    </div>

    <!-- Layout 2: Platform Description & Information -->
    <div class="glass-panel description-panel">
      <h2>Tentang Platform</h2>
      <p class="description">
        Platform rujukan medis digital terintegrasi yang memfasilitasi pembuatan, peninjauan, dan pengelolaan dokumen rujukan pasien antar Fasilitas Kesehatan (Faskes) secara instan, aman, dan transparan.
      </p>

      <div class="divider"></div>

      <div class="features-summary">
        <div class="summary-item">
          <span class="dot primary-dot"></span>
          <div class="item-text">
            <h4>Rujukan Masuk</h4>
            <p>Review diagnosis klinis ICD-10 pasien dan berikan tindakan persetujuan/penolakan dengan cepat.</p>
          </div>
        </div>
        <div class="summary-item">
          <span class="dot success-dot"></span>
          <div class="item-text">
            <h4>Rujukan Keluar</h4>
            <p>Buat rujukan transfer pasien ke rumah sakit dengan spesialisasi lebih lengkap secara paperless.</p>
          </div>
        </div>
        <div class="summary-item">
          <span class="dot warning-dot"></span>
          <div class="item-text">
            <h4>Notifikasi Alert</h4>
            <p>Integrasi peringatan kasus darurat (Cito) melalui bot Telegram agar segera mendapat perhatian.</p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <footer class="landing-footer">
    <p>&copy; 2026 Rujukan Medis Terintegrasi. All rights reserved.</p>
  </footer>
</div>

<style>
  .landing-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
    padding: 2rem 1.5rem;
    gap: 2rem;
  }

  .split-layout {
    display: grid;
    grid-template-columns: 1fr 1.2fr;
    gap: 2rem;
    width: 100%;
    max-width: 960px;
  }

  @media (max-width: 800px) {
    .split-layout {
      grid-template-columns: 1fr;
    }
  }

  .glass-panel {
    background: var(--bg-glass);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-lg);
    box-shadow: var(--shadow-card);
    padding: 3rem 2.5rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  /* Layout 1: Access Portal */
  .access-panel {
    align-items: center;
    text-align: center;
    gap: 2rem;
  }

  .logo-area {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.5rem;
  }

  .logo-icon {
    width: 54px;
    height: 54px;
    filter: drop-shadow(0 0 8px rgba(0, 242, 254, 0.3));
    margin-bottom: 0.5rem;
  }

  .logo-area h1 {
    font-size: 2.2rem;
    font-weight: 800;
    letter-spacing: 0.1em;
    background: linear-gradient(135deg, #ffffff 0%, #a5b4fc 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    line-height: 1;
  }

  .logo-area h1 span {
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-primary-hover) 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .portal-badge {
    background: rgba(0, 242, 254, 0.08);
    border: 1px solid rgba(0, 242, 254, 0.2);
    padding: 0.25rem 0.75rem;
    border-radius: var(--border-radius-full);
    font-size: 0.7rem;
    font-weight: 700;
    color: var(--color-primary);
    letter-spacing: 0.1em;
    margin-top: 0.5rem;
  }

  .action-buttons {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .w-full {
    width: 100%;
  }

  .test-tip {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    font-size: 0.8rem;
    color: var(--text-muted);
    background: rgba(255, 255, 255, 0.02);
    padding: 0.5rem 1rem;
    border-radius: var(--border-radius-sm);
    border: 1px solid var(--border-glass);
    width: 100%;
  }

  .test-tip strong {
    color: white;
  }

  /* Layout 2: Description Panel */
  .description-panel {
    gap: 1.5rem;
  }

  .description-panel h2 {
    font-size: 1.6rem;
    color: white;
    font-family: var(--font-title);
  }

  .description {
    color: var(--text-secondary);
    font-size: 0.95rem;
    line-height: 1.6;
  }

  .divider {
    height: 1px;
    background-color: var(--border-glass);
    width: 100%;
  }

  .features-summary {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .summary-item {
    display: flex;
    align-items: flex-start;
    gap: 1rem;
  }

  .dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    margin-top: 0.4rem;
    flex-shrink: 0;
  }

  .primary-dot {
    background-color: var(--color-primary);
    box-shadow: 0 0 8px var(--color-primary);
  }

  .success-dot {
    background-color: var(--color-success);
    box-shadow: 0 0 8px var(--color-success);
  }

  .warning-dot {
    background-color: var(--color-warning);
    box-shadow: 0 0 8px var(--color-warning);
  }

  .item-text h4 {
    font-size: 0.95rem;
    color: white;
    font-family: var(--font-title);
    margin-bottom: 0.15rem;
  }

  .item-text p {
    font-size: 0.85rem;
    color: var(--text-secondary);
    line-height: 1.4;
  }

  .landing-footer {
    color: var(--text-muted);
    font-size: 0.8rem;
  }
</style>
