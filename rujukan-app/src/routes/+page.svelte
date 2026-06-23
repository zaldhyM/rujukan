<!-- src/routes/+page.svelte -->
<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import kemkesLogo from '$lib/assets/kemkes.png';

  let isLoggedIn = $state(false);
  let copied = $state(false);

  onMount(() => {
    isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
  });

  const handleCTA = () => {
    if (isLoggedIn) {
      goto("/dashboard");
    } else {
      goto("/login");
    }
  };

  const copyCredentials = () => {
    navigator.clipboard.writeText("12345");
    copied = true;
    setTimeout(() => { copied = false; }, 2000);
  };
</script>

<svelte:head>
  <title>Rujukan Medis | Portal Utama</title>
</svelte:head>

<div class="landing-container animate-fade-in">
  <div class="glow-orb orb-1"></div>
  <div class="glow-orb orb-2"></div>

  <div class="unified-portal-card">
    <!-- Panel Kiri: Brand & Akses -->
    <div class="access-panel">
      <div class="logo-area">
        <div class="logo-wrapper">
          <img src={kemkesLogo} alt="Logo Kemenkes" class="logo-icon" />
        </div>
        <h1>RUJUKAN<span>MEDIS</span></h1>
        <span class="portal-badge">
          <span class="pulse-dot"></span>
          PORTAL UTAMA INTEGRASI FASKES
        </span>
      </div>

      <div class="action-buttons">
        <button onclick={handleCTA} class="btn-cta w-full">
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
          <a href="/register" class="btn-outline w-full">Daftar Faskes Baru</a>
        {/if}
      </div>

      <button onclick={copyCredentials} class="test-tip-btn" aria-label="Salin akun uji coba">
        <div class="test-tip-content">
          <svg viewBox="0 0 24 24" fill="none" width="16" height="16" stroke="rgba(255,255,255,0.7)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
          <span>Uji Coba: <code class="cred-code">12345</code> / <code class="cred-code">12345</code></span>
        </div>
        <div class="copy-action-text">
          {#if copied}
            <span class="text-copied">Tersalin!</span>
          {:else}
            <svg viewBox="0 0 24 24" fill="none" width="14" height="14" stroke="rgba(255,255,255,0.5)" stroke-width="2">
              <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
              <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
            </svg>
          {/if}
        </div>
      </button>
    </div>

    <!-- Panel Kanan: Deskripsi & Fitur -->
    <div class="description-panel">
      <div class="description-header">
        <h2>Tentang Platform</h2>
        <div class="decor-line">
          <svg viewBox="0 0 100 10" width="100" height="15" preserveAspectRatio="none">
            <path d="M0,5 L35,5 L40,1 L45,9 L50,5 L100,5" fill="none" stroke="url(#lineGrad2)" stroke-width="2" />
            <defs>
              <linearGradient id="lineGrad2" x1="0" y1="0" x2="1" y2="0">
                <stop offset="0%" stop-color="#2EC4B1" />
                <stop offset="100%" stop-color="#AACC00" />
              </linearGradient>
            </defs>
          </svg>
        </div>
      </div>
      <p class="description">
        Platform rujukan medis digital terintegrasi yang memfasilitasi
        pembuatan, peninjauan, dan pengelolaan dokumen rujukan pasien antar
        Fasilitas Kesehatan (Faskes) secara instan, aman, dan transparan.
      </p>

      <div class="divider"></div>

      <div class="features-list">
        <div class="feature-card">
          <div class="feature-icon-wrapper icon-primary">
            <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="7 10 12 15 17 10"></polyline>
              <line x1="12" y1="15" x2="12" y2="3"></line>
            </svg>
          </div>
          <div class="feature-info">
            <h3>Rujukan Masuk</h3>
            <p>Review diagnosis klinis ICD-10 pasien dan berikan tindakan persetujuan/penolakan dengan cepat.</p>
          </div>
        </div>

        <div class="feature-card">
          <div class="feature-icon-wrapper icon-success">
            <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="17 8 12 3 7 8"></polyline>
              <line x1="12" y1="3" x2="12" y2="15"></line>
            </svg>
          </div>
          <div class="feature-info">
            <h3>Rujukan Keluar</h3>
            <p>Buat rujukan transfer pasien ke rumah sakit dengan spesialisasi lebih lengkap secara paperless.</p>
          </div>
        </div>

        <div class="feature-card">
          <div class="feature-icon-wrapper icon-warning">
            <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
              <line x1="12" y1="9" x2="12" y2="13"></line>
              <line x1="12" y1="17" x2="12.01" y2="17"></line>
            </svg>
          </div>
          <div class="feature-info">
            <h3>Notifikasi Alert</h3>
            <p>Integrasi peringatan kasus darurat (Cito) melalui bot Telegram agar segera mendapat perhatian.</p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <footer class="landing-footer">
    <p>&copy; 2026 Rujukan Medis Terintegrasi &mdash; Kementerian Kesehatan RI</p>
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
    position: relative;
    overflow: hidden;
  }

  .glow-orb {
    position: absolute;
    border-radius: 50%;
    filter: blur(120px);
    z-index: 0;
    opacity: 0.06;
    pointer-events: none;
    animation: orb-bounce 15s ease-in-out infinite alternate;
  }
  .orb-1 {
    width: 400px;
    height: 400px;
    background: #2EC4B1;
    top: 0;
    left: 5%;
  }
  .orb-2 {
    width: 400px;
    height: 400px;
    background: #AACC00;
    bottom: 0;
    right: 5%;
    animation-delay: -5s;
  }
  @keyframes orb-bounce {
    0%   { transform: translate(0, 0) scale(1); }
    100% { transform: translate(40px, 30px) scale(1.1); }
  }

  /* Card utama */
  .unified-portal-card {
    display: grid;
    grid-template-columns: 1fr 1.2fr;
    width: 100%;
    max-width: 1000px;
    z-index: 1;
    overflow: hidden;
    border-radius: var(--border-radius-lg);
    box-shadow: 0 8px 40px rgba(0, 0, 0, 0.12);
    border: 1px solid rgba(0, 0, 0, 0.06);
  }

  /* Panel kiri — teal Kemenkes */
  .access-panel {
    background: #0D7A72;
    padding: 3.5rem 3rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 2rem;
    text-align: center;
  }

  .logo-area {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
    width: 100%;
  }

  .logo-wrapper {
    width: 68px;
    height: 68px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 0.5rem;
  }

  .logo-icon {
    width: 72px;
    height: 72px;
    object-fit: contain;
    filter: drop-shadow(0 0 10px rgba(170, 204, 0, 0.5));
    animation: float 4s ease-in-out infinite;
  }

  @keyframes float {
    0%, 100% { transform: translateY(0); }
    50%       { transform: translateY(-6px); }
  }

  .logo-area h1 {
    font-size: 2.2rem;
    font-weight: 800;
    letter-spacing: 0.05em;
    color: #ffffff;
    line-height: 1.1;
  }

  .logo-area h1 span {
    color: #AACC00;
  }

  .portal-badge {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    background: rgba(255, 255, 255, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.2);
    padding: 0.35rem 0.85rem;
    border-radius: var(--border-radius-full);
    font-size: 0.7rem;
    font-weight: 700;
    color: rgba(255, 255, 255, 0.85);
    letter-spacing: 0.08em;
    margin-top: 0.25rem;
  }

  .pulse-dot {
    width: 6px;
    height: 6px;
    background: #AACC00;
    border-radius: 50%;
    animation: pulse 1.8s infinite;
  }
  @keyframes pulse {
    0%   { transform: scale(0.9); opacity: 0.6; }
    50%  { transform: scale(1.4); opacity: 1; box-shadow: 0 0 8px #AACC00; }
    100% { transform: scale(0.9); opacity: 0.6; }
  }

  .action-buttons {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 0.85rem;
  }

  .w-full { width: 100%; }

  .btn-cta {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    width: 100%;
    padding: 0.85rem 1.5rem;
    border-radius: var(--border-radius-sm);
    font-family: var(--font-title);
    font-weight: 700;
    font-size: 0.95rem;
    cursor: pointer;
    border: none;
    background: #AACC00;
    color: #0D7A72;
    box-shadow: 0 4px 14px rgba(170, 204, 0, 0.35);
    transition: all var(--transition-normal);
  }

  .btn-cta:hover {
    background: #BBDD00;
    box-shadow: 0 6px 20px rgba(170, 204, 0, 0.5);
    transform: translateY(-2px);
  }

  .btn-outline {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 0.85rem 1.5rem;
    border-radius: var(--border-radius-sm);
    font-family: var(--font-title);
    font-weight: 600;
    font-size: 0.95rem;
    cursor: pointer;
    background: transparent;
    border: 1.5px solid rgba(255, 255, 255, 0.35);
    color: rgba(255, 255, 255, 0.85);
    transition: all var(--transition-normal);
    text-align: center;
  }

  .btn-outline:hover {
    background: rgba(255, 255, 255, 0.1);
    border-color: rgba(255, 255, 255, 0.6);
    color: #ffffff;
  }

  .test-tip-btn {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 0.5rem;
    font-size: 0.8rem;
    color: rgba(255, 255, 255, 0.6);
    background: rgba(255, 255, 255, 0.06);
    padding: 0.65rem 1rem;
    border-radius: var(--border-radius-sm);
    border: 1px solid rgba(255, 255, 255, 0.12);
    width: 100%;
    cursor: pointer;
    text-align: left;
    transition: all var(--transition-fast);
    outline: none;
  }

  .test-tip-btn:hover {
    background: rgba(255, 255, 255, 0.1);
    border-color: rgba(255, 255, 255, 0.2);
  }

  .test-tip-content {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .cred-code {
    background: rgba(170, 204, 0, 0.2);
    color: #AACC00;
    padding: 0.1rem 0.35rem;
    border-radius: 4px;
    font-weight: 700;
    font-family: monospace;
  }

  .copy-action-text {
    display: flex;
    align-items: center;
    font-size: 0.75rem;
  }

  .text-copied {
    color: #AACC00;
    font-weight: 600;
  }

  /* Panel kanan — putih */
  .description-panel {
    background: #FFFFFF;
    padding: 3.5rem 3rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 1.5rem;
  }

  .description-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 1rem;
  }

  .description-panel h2 {
    font-size: 1.6rem;
    color: var(--text-primary);
    font-family: var(--font-title);
  }

  .decor-line {
    width: 100px;
    opacity: 0.9;
  }

  .description {
    color: var(--text-secondary);
    font-size: 0.95rem;
    line-height: 1.7;
  }

  .divider {
    height: 1px;
    background: rgba(0, 0, 0, 0.07);
    width: 100%;
  }

  .features-list {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .feature-card {
    display: flex;
    align-items: flex-start;
    gap: 1.25rem;
    padding: 1.1rem 1.25rem;
    border-radius: var(--border-radius-md);
    background: var(--bg-primary);
    border: 1px solid rgba(0, 0, 0, 0.06);
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .feature-card:hover {
    transform: translateX(4px);
    background: #FFFFFF;
    border-color: rgba(46, 196, 177, 0.25);
    box-shadow: 0 4px 16px rgba(46, 196, 177, 0.1);
  }

  .feature-icon-wrapper {
    width: 44px;
    height: 44px;
    border-radius: var(--border-radius-sm);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    transition: all var(--transition-fast);
  }

  .icon-primary {
    background: rgba(46, 196, 177, 0.1);
    border: 1px solid rgba(46, 196, 177, 0.25);
    color: #2EC4B1;
  }
  .feature-card:hover .icon-primary {
    background: #2EC4B1;
    color: #ffffff;
    box-shadow: 0 4px 12px rgba(46, 196, 177, 0.35);
  }

  .icon-success {
    background: rgba(170, 204, 0, 0.1);
    border: 1px solid rgba(170, 204, 0, 0.25);
    color: #7A9400;
  }
  .feature-card:hover .icon-success {
    background: #AACC00;
    color: #ffffff;
    box-shadow: 0 4px 12px rgba(170, 204, 0, 0.35);
  }

  .icon-warning {
    background: rgba(245, 158, 11, 0.08);
    border: 1px solid rgba(245, 158, 11, 0.2);
    color: #D97706;
  }
  .feature-card:hover .icon-warning {
    background: #F59E0B;
    color: #ffffff;
    box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
  }

  .feature-info h3 {
    font-size: 1rem;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 0.25rem;
    font-family: var(--font-title);
  }

  .feature-info p {
    font-size: 0.85rem;
    color: var(--text-secondary);
    line-height: 1.5;
  }

  .landing-footer {
    color: var(--text-muted);
    font-size: 0.8rem;
    z-index: 1;
  }

  @media (max-width: 850px) {
    .unified-portal-card {
      grid-template-columns: 1fr;
    }
    .access-panel {
      padding: 3rem 2rem;
    }
    .description-panel {
      padding: 3rem 2rem;
    }
  }
</style>
