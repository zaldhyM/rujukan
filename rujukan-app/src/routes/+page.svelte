<!-- src/routes/+page.svelte -->
<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";

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
    setTimeout(() => {
      copied = false;
    }, 2000);
  };
</script>

<svelte:head>
  <title>Rujukan Medis | Portal Utama</title>
</svelte:head>

<div class="landing-container animate-fade-in">
  <!-- Decorative background orbs -->
  <div class="glow-orb orb-1"></div>
  <div class="glow-orb orb-2"></div>

  <div class="unified-portal-card">
    <!-- Layout 1: Brand & Access Portal -->
    <div class="access-panel">
      <div class="logo-area">
        <div class="logo-wrapper">
          <svg
            class="logo-icon"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M19 10.5V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V10.5M12 3V17M12 3L7.5 7.5M12 3L16.5 7.5"
              stroke="url(#logoGrad)"
              stroke-width="2.5"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <defs>
              <linearGradient
                id="logoGrad"
                x1="5"
                y1="3"
                x2="19"
                y2="21"
                gradientUnits="userSpaceOnUse"
              >
                <stop stop-color="#00f2fe" />
                <stop offset="1" stop-color="#8b5cf6" />
              </linearGradient>
            </defs>
          </svg>
        </div>
        <h1>RUJUKAN<span>MEDIS</span></h1>
        <span class="portal-badge">
          <span class="pulse-dot"></span>
          PORTAL UTAMA INTEGRASI FASKES
        </span>
      </div>

      <div class="action-buttons">
        <button onclick={handleCTA} class="btn btn-primary w-full">
          {#if isLoggedIn}
            <span>Masuk Dashboard</span>
          {:else}
            <span>Masuk ke Sistem</span>
          {/if}
          <svg
            viewBox="0 0 24 24"
            fill="none"
            width="18"
            height="18"
            stroke="currentColor"
            stroke-width="2.5"
          >
            <line x1="5" y1="12" x2="19" y2="12"></line>
            <polyline points="12 5 19 12 12 19"></polyline>
          </svg>
        </button>

        {#if !isLoggedIn}
          <a href="/register" class="btn btn-secondary w-full"
            >Daftar Faskes Baru</a
          >
        {/if}
      </div>

      <button
        onclick={copyCredentials}
        class="test-tip-btn"
        aria-label="Salin akun uji coba"
      >
        <div class="test-tip-content">
          <svg
            viewBox="0 0 24 24"
            fill="none"
            width="16"
            height="16"
            stroke="var(--color-primary)"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
          <span
            >Uji Coba: <code class="cred-code">12345</code> /
            <code class="cred-code">12345</code></span
          >
        </div>
        <div class="copy-action-text">
          {#if copied}
            <span class="text-success">Tersalin!</span>
          {:else}
            <svg
              viewBox="0 0 24 24"
              fill="none"
              width="14"
              height="14"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
              <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"
              ></path>
            </svg>
          {/if}
        </div>
      </button>
    </div>

    <!-- Layout 2: Platform Description & Information -->
    <div class="description-panel">
      <div class="description-header">
        <h2>Tentang Platform</h2>
        <div class="decor-line">
          <svg
            viewBox="0 0 100 10"
            width="100"
            height="15"
            preserveAspectRatio="none"
          >
            <path
              d="M0,5 L35,5 L40,1 L45,9 L50,5 L100,5"
              fill="none"
              stroke="url(#lineGrad2)"
              stroke-width="2"
            />
            <defs>
              <linearGradient id="lineGrad2" x1="0" y1="0" x2="1" y2="0">
                <stop offset="0%" stop-color="#00f2fe" />
                <stop offset="100%" stop-color="#8b5cf6" />
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
          <div class="feature-icon-wrapper primary-glow">
            <svg
              viewBox="0 0 24 24"
              fill="none"
              width="22"
              height="22"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="7 10 12 15 17 10"></polyline>
              <line x1="12" y1="15" x2="12" y2="3"></line>
            </svg>
          </div>
          <div class="feature-info">
            <h3>Rujukan Masuk</h3>
            <p>
              Review diagnosis klinis ICD-10 pasien dan berikan tindakan
              persetujuan/penolakan dengan cepat.
            </p>
          </div>
        </div>

        <div class="feature-card">
          <div class="feature-icon-wrapper success-glow">
            <svg
              viewBox="0 0 24 24"
              fill="none"
              width="22"
              height="22"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="17 8 12 3 7 8"></polyline>
              <line x1="12" y1="3" x2="12" y2="15"></line>
            </svg>
          </div>
          <div class="feature-info">
            <h3>Rujukan Keluar</h3>
            <p>
              Buat rujukan transfer pasien ke rumah sakit dengan spesialisasi
              lebih lengkap secara paperless.
            </p>
          </div>
        </div>

        <div class="feature-card">
          <div class="feature-icon-wrapper warning-glow">
            <svg
              viewBox="0 0 24 24"
              fill="none"
              width="22"
              height="22"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path
                d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"
              ></path>
              <line x1="12" y1="9" x2="12" y2="13"></line>
              <line x1="12" y1="17" x2="12.01" y2="17"></line>
            </svg>
          </div>
          <div class="feature-info">
            <h3>Notifikasi Alert</h3>
            <p>
              Integrasi peringatan kasus darurat (Cito) melalui bot Telegram
              agar segera mendapat perhatian.
            </p>
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
    position: relative;
    overflow: hidden;
  }

  /* Decorative glow orbs */
  .glow-orb {
    position: absolute;
    border-radius: 50%;
    filter: blur(140px);
    z-index: -1;
    opacity: 0.12;
    pointer-events: none;
    animation: orb-bounce 15s ease-in-out infinite alternate;
  }
  .orb-1 {
    width: 350px;
    height: 350px;
    background: var(--color-primary);
    top: 5%;
    left: 10%;
  }
  .orb-2 {
    width: 450px;
    height: 450px;
    background: var(--color-secondary);
    bottom: 5%;
    right: 10%;
    animation-delay: -5s;
  }
  @keyframes orb-bounce {
    0% {
      transform: translate(0, 0) scale(1);
    }
    100% {
      transform: translate(60px, 40px) scale(1.15);
    }
  }

  /* Single Unified Portal Card */
  .unified-portal-card {
    display: grid;
    grid-template-columns: 1fr 1.2fr;
    width: 100%;
    max-width: 1000px;
    z-index: 1;
    overflow: hidden;
    background: var(--bg-glass);
    backdrop-filter: blur(20px);
    -webkit-backdrop-filter: blur(20px);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-lg);
    box-shadow: var(--shadow-card);
    transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .unified-portal-card:hover {
    border-color: rgba(255, 255, 255, 0.15);
    box-shadow: 0 12px 40px 0 rgba(0, 0, 0, 0.4);
  }

  /* Layout 1: Access Portal */
  .access-panel {
    align-items: center;
    text-align: center;
    gap: 2rem;
    position: relative;
    padding: 3.5rem 3rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background: rgba(
      9,
      15,
      26,
      0.55
    ); /* Bedah warna: darker side panel restored */
    border-right: 1px solid var(--border-glass); /* Garis pembatas panel */
  }

  .logo-area {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
    width: 100%;
  }

  .logo-wrapper {
    position: relative;
    width: 68px;
    height: 68px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 0.5rem;
  }

  .logo-icon {
    width: 56px;
    height: 56px;
    filter: drop-shadow(0 0 12px rgba(0, 242, 254, 0.4));
    z-index: 2;
    animation: float 4s ease-in-out infinite;
  }

  @keyframes float {
    0%,
    100% {
      transform: translateY(0);
    }
    50% {
      transform: translateY(-6px);
    }
  }

  .logo-area h1 {
    font-size: 2.2rem;
    font-weight: 800;
    letter-spacing: 0.05em;
    background: linear-gradient(135deg, #ffffff 0%, #a5b4fc 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    line-height: 1.1;
  }

  .logo-area h1 span {
    background: linear-gradient(
      135deg,
      var(--color-primary) 0%,
      var(--color-secondary) 100%
    );
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    text-shadow: 0 0 20px rgba(0, 242, 254, 0.15);
  }

  .portal-badge {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    background: rgba(0, 242, 254, 0.06);
    border: 1px solid rgba(0, 242, 254, 0.15);
    padding: 0.35rem 0.85rem;
    border-radius: var(--border-radius-full);
    font-size: 0.7rem;
    font-weight: 700;
    color: var(--color-primary);
    letter-spacing: 0.08em;
    margin-top: 0.25rem;
  }

  .pulse-dot {
    width: 6px;
    height: 6px;
    background-color: var(--color-success);
    border-radius: 50%;
    box-shadow: 0 0 8px var(--color-success);
    animation: pulse 1.8s infinite;
  }
  @keyframes pulse {
    0% {
      transform: scale(0.9);
      opacity: 0.6;
    }
    50% {
      transform: scale(1.3);
      opacity: 1;
      box-shadow: 0 0 12px var(--color-success);
    }
    100% {
      transform: scale(0.9);
      opacity: 0.6;
    }
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

  /* Tombol Utama menjadi putih dengan teks gelap */
  .btn-primary {
    position: relative;
    overflow: hidden;
    background: #00f2fe;
    color: #0b111e;
    font-weight: 700;
    box-shadow: 0 4px 15px rgba(255, 255, 255, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.4);
  }

  .btn-primary:hover {
    background: rgba(255, 255, 255, 0.9);
    box-shadow: 0 6px 20px rgba(255, 255, 255, 0.45);
    transform: translateY(-2px);
  }

  .btn-primary::after {
    content: "";
    position: absolute;
    top: 0;
    left: -50%;
    width: 20%;
    height: 100%;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0) 0%,
      rgba(0, 0, 0, 0.1) 50%,
      rgba(0, 0, 0, 0) 100%
    );
    transform: skewX(-25deg);
    transition: 0.75s;
    opacity: 0;
  }

  .btn-primary:hover::after {
    left: 125%;
    opacity: 1;
    transition: 0.75s;
  }

  .btn-secondary {
    border: 1px solid var(--border-glass);
    background: rgba(255, 255, 255, 0.03);
    color: var(--text-primary);
    transition: all 0.3s ease;
  }

  .btn-secondary:hover {
    background: rgba(255, 255, 255, 0.08);
    border-color: var(--color-primary);
    box-shadow: 0 0 15px rgba(0, 242, 254, 0.15);
  }

  /* Test account copy card style */
  .test-tip-btn {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 0.5rem;
    font-size: 0.8rem;
    color: var(--text-secondary);
    background: rgba(255, 255, 255, 0.02);
    padding: 0.65rem 1rem;
    border-radius: var(--border-radius-sm);
    border: 1px solid var(--border-glass);
    width: 100%;
    cursor: pointer;
    text-align: left;
    transition: all var(--transition-fast);
    outline: none;
  }

  .test-tip-btn:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.15);
  }

  .test-tip-content {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .cred-code {
    background: rgba(0, 242, 254, 0.1);
    color: var(--color-primary);
    padding: 0.1rem 0.35rem;
    border-radius: 4px;
    font-weight: 600;
    font-family: monospace;
  }

  .copy-action-text {
    display: flex;
    align-items: center;
    font-size: 0.75rem;
    color: var(--text-muted);
  }

  .test-tip-btn:hover .copy-action-text {
    color: var(--color-primary);
  }

  .text-success {
    color: var(--color-success);
    font-weight: 600;
  }

  /* Layout 2: Description Panel */
  .description-panel {
    gap: 1.5rem;
    padding: 3.5rem 3rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background: rgba(255, 255, 255, 0.01); /* Bedah warna: lighter side */
  }

  .description-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    flex-wrap: wrap;
    gap: 1rem;
  }

  .description-panel h2 {
    font-size: 1.6rem;
    color: white;
    font-family: var(--font-title);
  }

  .decor-line {
    flex-grow: 0;
    width: 100px;
    opacity: 0.8;
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

  .features-list {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
    width: 100%;
  }

  .feature-card {
    display: flex;
    align-items: flex-start;
    gap: 1.25rem;
    padding: 1.25rem;
    border-radius: var(--border-radius-md);
    background: rgba(255, 255, 255, 0.02);
    border: 1px solid rgba(255, 255, 255, 0.04);
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .feature-card:hover {
    transform: translateY(-2px) translateX(4px);
    background: rgba(255, 255, 255, 0.04);
    border-color: rgba(255, 255, 255, 0.1);
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
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

  .primary-glow {
    background: rgba(0, 242, 254, 0.08);
    border: 1px solid rgba(0, 242, 254, 0.2);
    color: var(--color-primary);
  }
  .feature-card:hover .primary-glow {
    background: var(--color-primary);
    color: #0b111e;
    box-shadow: 0 0 15px var(--color-primary);
  }

  .success-glow {
    background: rgba(0, 245, 160, 0.08);
    border: 1px solid rgba(0, 245, 160, 0.2);
    color: var(--color-success);
  }
  .feature-card:hover .success-glow {
    background: var(--color-success);
    color: #0b111e;
    box-shadow: 0 0 15px var(--color-success);
  }

  .warning-glow {
    background: rgba(255, 174, 52, 0.08);
    border: 1px solid rgba(255, 174, 52, 0.2);
    color: var(--color-warning);
  }
  .feature-card:hover .warning-glow {
    background: var(--color-warning);
    color: #0b111e;
    box-shadow: 0 0 15px var(--color-warning);
  }

  .feature-info h3 {
    font-size: 1rem;
    font-weight: 600;
    color: white;
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

  /* Responsive Design adjustments */
  @media (max-width: 850px) {
    .unified-portal-card {
      grid-template-columns: 1fr;
    }
    .access-panel {
      border-right: none;
      border-bottom: 1px solid var(--border-glass);
      padding: 3rem 2rem;
    }
    .description-panel {
      padding: 3rem 2rem;
    }
  }
</style>
