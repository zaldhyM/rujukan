<!-- src/routes/+page.svelte -->
<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import kemkesLogo from "$lib/assets/kemkes.png";

  let isLoggedIn = $state(false);

  onMount(() => {
    isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
  });
</script>

<svelte:head>
  <title>Rujukan Medis | Sehat Untuk Indonesia</title>
</svelte:head>

<div class="landing-page">
  <!-- ── Navbar ── -->
  <nav class="navbar">
    <div class="nav-inner">
      <div class="nav-logo">
        <img src={kemkesLogo} alt="Logo Kemenkes" />
        <span>Kemenkes</span>
      </div>
      <div class="nav-links">
        <a href="/login" class="nav-link">Tentang</a>
        <a href="/login" class="nav-link">Layanan</a>
        <a href="/login" class="nav-link">Kontak</a>
      </div>
      {#if isLoggedIn}
        <a href="/dashboard" class="nav-btn">Dashboard</a>
      {:else}
        <a href="/login" class="nav-btn">Masuk</a>
      {/if}
    </div>
  </nav>

  <!-- ── Hero Section ── -->
  <section class="hero">
    <!-- Left content -->
    <div class="hero-content">
      <div class="hero-badge">
        <span class="badge-dot"></span>
        Platform Rujukan Terintegrasi Kemenkes RI
      </div>

      <h1 class="hero-title">
        Sehat Untuk<br />
        <span class="title-accent">Indonesia</span>
      </h1>

      <p class="hero-desc">
        Kemenkes hadir untuk mewujudkan masyarakat sehat melalui layanan
        kesehatan yang mudah diakses, terintegrasi, dan terpercaya.
      </p>

      <div class="hero-actions">
        {#if isLoggedIn}
          <a href="/dashboard" class="btn-primary-hero">
            Dashboard <svg viewBox="0 0 24 24" fill="none" width="17" height="17" stroke="currentColor" stroke-width="2.5"><line x1="5" y1="12" x2="19" y2="12"/><polyline points="12 5 19 12 12 19"/></svg>
          </a>
        {:else}
          <a href="/login" class="btn-primary-hero">
            Mulai Sekarang <svg viewBox="0 0 24 24" fill="none" width="17" height="17" stroke="currentColor" stroke-width="2.5"><line x1="5" y1="12" x2="19" y2="12"/><polyline points="12 5 19 12 12 19"/></svg>
          </a>
        {/if}
        <a href="/register" class="btn-outline-hero">Pelajari Lebih Lanjut</a>
      </div>

      <!-- Stats row -->
      <div class="stats-row">
        <div class="stat-item">
          <span class="stat-num">2.400+</span>
          <span class="stat-label">Faskes Terdaftar</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat-item">
          <span class="stat-num">98%</span>
          <span class="stat-label">Tingkat Keberhasilan</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat-item">
          <span class="stat-num">24/7</span>
          <span class="stat-label">Layanan Aktif</span>
        </div>
      </div>
    </div>

    <!-- Right illustration -->
    <div class="hero-visual">
      <!-- Background blobs -->
      <div class="blob-green"></div>
      <div class="blob-teal"></div>

      <!-- Medical illustration SVG -->
      <div class="medical-illustration">
        <svg viewBox="0 0 420 480" fill="none" xmlns="http://www.w3.org/2000/svg" class="illus-svg">
          <!-- Background shapes -->
          <ellipse cx="210" cy="380" rx="160" ry="40" fill="rgba(20,184,166,0.08)"/>

          <!-- Doctor 1 (center, taller) -->
          <g class="doctor-center">
            <!-- Body -->
            <rect x="165" y="200" width="90" height="180" rx="8" fill="#FFFFFF" stroke="#E5E7EB" stroke-width="1.5"/>
            <!-- Coat detail -->
            <rect x="195" y="200" width="30" height="180" rx="2" fill="#F0FDF4"/>
            <!-- Head -->
            <circle cx="210" cy="175" r="32" fill="#FDDCB5"/>
            <!-- Hair -->
            <path d="M178 165 Q178 135 210 135 Q242 135 242 165" fill="#3D2B1F"/>
            <!-- Stethoscope -->
            <path d="M195 235 Q185 250 185 265 Q185 275 195 275 Q205 275 205 265" stroke="#9CA3AF" stroke-width="2.5" fill="none" stroke-linecap="round"/>
            <circle cx="185" cy="265" r="5" fill="#9CA3AF"/>
            <!-- Collar -->
            <path d="M195 200 L210 220 L225 200" fill="#E5E7EB" stroke="#D1D5DB" stroke-width="1"/>
            <!-- Face features -->
            <circle cx="201" cy="172" r="2.5" fill="#6B4226"/>
            <circle cx="219" cy="172" r="2.5" fill="#6B4226"/>
            <path d="M204 183 Q210 188 216 183" stroke="#6B4226" stroke-width="1.5" fill="none" stroke-linecap="round"/>
            <!-- Arms -->
            <rect x="145" y="210" width="22" height="80" rx="10" fill="#FDDCB5"/>
            <rect x="253" y="210" width="22" height="80" rx="10" fill="#FDDCB5"/>
            <!-- Tablet in right hand -->
            <rect x="252" y="265" width="32" height="24" rx="3" fill="#1F2937"/>
            <rect x="254" y="267" width="28" height="20" rx="2" fill="#14B8A6"/>
            <!-- Legs -->
            <rect x="175" y="375" width="30" height="60" rx="6" fill="#374151"/>
            <rect x="215" y="375" width="30" height="60" rx="6" fill="#374151"/>
            <!-- Shoes -->
            <ellipse cx="190" cy="435" rx="18" ry="8" fill="#1F2937"/>
            <ellipse cx="230" cy="435" rx="18" ry="8" fill="#1F2937"/>
          </g>

          <!-- Nurse left -->
          <g class="nurse-left">
            <rect x="68" y="230" width="76" height="155" rx="8" fill="#14B8A6"/>
            <rect x="93" y="230" width="26" height="155" rx="2" fill="rgba(255,255,255,0.2)"/>
            <circle cx="106" cy="208" r="27" fill="#FDDCB5"/>
            <path d="M79 200 Q79 174 106 174 Q133 174 133 200" fill="#1F2937"/>
            <!-- Nurse cap -->
            <rect x="94" y="174" width="24" height="10" rx="3" fill="#FFFFFF"/>
            <rect x="103" y="170" width="6" height="8" rx="1" fill="#DC2626"/>
            <!-- Face -->
            <circle cx="99" cy="206" r="2" fill="#6B4226"/>
            <circle cx="113" cy="206" r="2" fill="#6B4226"/>
            <path d="M101 216 Q106 220 111 216" stroke="#6B4226" stroke-width="1.5" fill="none" stroke-linecap="round"/>
            <!-- Arms -->
            <rect x="48" y="240" width="22" height="65" rx="10" fill="#FDDCB5"/>
            <rect x="142" y="240" width="20" height="65" rx="10" fill="#FDDCB5"/>
            <!-- Clipboard in left hand -->
            <rect x="32" y="285" width="22" height="28" rx="3" fill="#FFFFFF" stroke="#E5E7EB" stroke-width="1.5"/>
            <line x1="37" y1="294" x2="49" y2="294" stroke="#9CA3AF" stroke-width="1.2"/>
            <line x1="37" y1="300" x2="46" y2="300" stroke="#9CA3AF" stroke-width="1.2"/>
            <line x1="37" y1="306" x2="49" y2="306" stroke="#9CA3AF" stroke-width="1.2"/>
            <!-- Legs -->
            <rect x="78" y="380" width="22" height="50" rx="5" fill="#374151"/>
            <rect x="106" y="380" width="22" height="50" rx="5" fill="#374151"/>
            <ellipse cx="89" cy="430" rx="15" ry="7" fill="#1F2937"/>
            <ellipse cx="117" cy="430" rx="15" ry="7" fill="#1F2937"/>
          </g>

          <!-- Doctor right (hijab) -->
          <g class="doctor-right">
            <rect x="278" y="225" width="80" height="160" rx="8" fill="#FFFFFF" stroke="#E5E7EB" stroke-width="1.5"/>
            <rect x="303" y="225" width="30" height="160" rx="2" fill="#F0FDF4"/>
            <circle cx="318" cy="198" r="30" fill="#FDDCB5"/>
            <!-- Hijab -->
            <path d="M288 198 Q288 165 318 162 Q348 165 348 198 Q348 225 318 228 Q288 225 288 198Z" fill="#2563EB"/>
            <path d="M288 198 Q280 220 285 240" stroke="#2563EB" stroke-width="8" fill="none" stroke-linecap="round"/>
            <path d="M348 198 Q356 220 351 240" stroke="#2563EB" stroke-width="8" fill="none" stroke-linecap="round"/>
            <!-- Face showing -->
            <circle cx="310" cy="196" r="2" fill="#6B4226"/>
            <circle cx="326" cy="196" r="2" fill="#6B4226"/>
            <path d="M312 207 Q318 211 324 207" stroke="#6B4226" stroke-width="1.5" fill="none" stroke-linecap="round"/>
            <!-- Arms -->
            <rect x="258" y="234" width="22" height="68" rx="10" fill="#FDDCB5"/>
            <rect x="356" y="234" width="22" height="68" rx="10" fill="#FDDCB5"/>
            <!-- Stethoscope -->
            <path d="M303 250 Q293 265 293 278 Q293 288 303 288 Q313 288 313 278" stroke="#9CA3AF" stroke-width="2.5" fill="none" stroke-linecap="round"/>
            <circle cx="293" cy="278" r="5" fill="#9CA3AF"/>
            <!-- Legs -->
            <rect x="288" y="378" width="22" height="52" rx="5" fill="#1E3A8A"/>
            <rect x="316" y="378" width="22" height="52" rx="5" fill="#1E3A8A"/>
            <ellipse cx="299" cy="430" rx="16" ry="7" fill="#1F2937"/>
            <ellipse cx="327" cy="430" rx="16" ry="7" fill="#1F2937"/>
          </g>
        </svg>
      </div>

      <!-- Floating badge cards -->
      <div class="float-card card-top-left">
        <div class="float-icon teal">
          <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="currentColor" stroke-width="2">
            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
            <path d="M12 8v4M12 16h.01" stroke-width="2.5"/>
          </svg>
        </div>
        <div>
          <p class="float-label">Keamanan Data</p>
          <p class="float-sub">Terenkripsi SSL</p>
        </div>
      </div>

      <div class="float-card card-top-right">
        <div class="float-icon lime">
          <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="currentColor" stroke-width="2">
            <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/>
          </svg>
        </div>
        <div>
          <p class="float-label">Respon Cepat</p>
          <p class="float-sub">Rata-rata &lt; 2 menit</p>
        </div>
      </div>

      <div class="float-card card-bottom-left">
        <div class="float-icon teal">
          <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="currentColor" stroke-width="2">
            <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
            <circle cx="9" cy="7" r="4"/>
            <path d="M23 21v-2a4 4 0 0 0-3-3.87M16 3.13a4 4 0 0 1 0 7.75"/>
          </svg>
        </div>
        <div>
          <p class="float-label">2.400+ Faskes</p>
          <p class="float-sub">Seluruh Indonesia</p>
        </div>
      </div>
    </div>
  </section>

  <!-- ── Features strip ── -->
  <section class="features-strip">
    <div class="features-inner">
      <div class="feat-item">
        <div class="feat-icon">
          <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2">
            <polyline points="7 10 12 15 17 10"/><line x1="12" y1="3" x2="12" y2="15"/>
            <path d="M5 19h14a2 2 0 0 0 2-2V7"/>
          </svg>
        </div>
        <div>
          <h3>Rujukan Masuk</h3>
          <p>Review &amp; tindak lanjuti rujukan dari faskes pengirim</p>
        </div>
      </div>
      <div class="feat-item">
        <div class="feat-icon lime-icon">
          <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2">
            <polyline points="17 8 12 3 7 8"/><line x1="12" y1="3" x2="12" y2="15"/>
            <path d="M19 19H5a2 2 0 0 1-2-2V7"/>
          </svg>
        </div>
        <div>
          <h3>Rujukan Keluar</h3>
          <p>Kirim surat rujukan digital ke RS penerima spesialis</p>
        </div>
      </div>
      <div class="feat-item">
        <div class="feat-icon">
          <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
            <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
          </svg>
        </div>
        <div>
          <h3>Notifikasi Real-time</h3>
          <p>Alert instan via Telegram Bot untuk kasus Cito</p>
        </div>
      </div>
      <div class="feat-item">
        <div class="feat-icon lime-icon">
          <svg viewBox="0 0 24 24" fill="none" width="22" height="22" stroke="currentColor" stroke-width="2">
            <rect x="3" y="11" width="18" height="11" rx="2"/>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
          </svg>
        </div>
        <div>
          <h3>Keamanan Tinggi</h3>
          <p>Data terenkripsi end-to-end sesuai standar Kemenkes</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Footer -->
  <footer class="landing-footer">
    <p>&copy; 2026 Rujukan Medis Terintegrasi &mdash; Kementerian Kesehatan RI</p>
  </footer>
</div>

<style>
  .landing-page {
    min-height: 100vh;
    background: #FFFFFF;
    display: flex;
    flex-direction: column;
    overflow-x: hidden;
  }

  /* ── Navbar ── */
  .navbar {
    position: sticky;
    top: 0;
    z-index: 100;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(8px);
    border-bottom: 1px solid #F3F4F6;
  }

  .nav-inner {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2.5rem;
    height: 64px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 2rem;
  }

  .nav-logo {
    display: flex;
    align-items: center;
    gap: 0.625rem;
    text-decoration: none;
    flex-shrink: 0;
  }

  .nav-logo img {
    width: 32px;
    height: 32px;
    object-fit: contain;
  }

  .nav-logo span {
    font-family: var(--font-title);
    font-size: 1.1rem;
    font-weight: 700;
    color: var(--text-primary);
  }

  .nav-links {
    display: flex;
    align-items: center;
    gap: 2rem;
    flex: 1;
    justify-content: center;
  }

  .nav-link {
    font-size: 0.9rem;
    font-weight: 500;
    color: var(--text-muted);
    transition: color 0.15s;
    text-decoration: none;
  }

  .nav-link:hover { color: var(--color-primary); }

  .nav-btn {
    display: inline-flex;
    align-items: center;
    padding: 0.5rem 1.25rem;
    background: var(--color-primary);
    color: #FFFFFF;
    border-radius: 8px;
    font-family: var(--font-title);
    font-size: 0.875rem;
    font-weight: 600;
    text-decoration: none;
    transition: background 0.15s, box-shadow 0.15s;
    flex-shrink: 0;
  }

  .nav-btn:hover {
    background: var(--color-primary-hover);
    box-shadow: 0 2px 8px rgba(20, 184, 166, 0.3);
  }

  /* ── Hero Section ── */
  .hero {
    flex: 1;
    max-width: 1200px;
    margin: 0 auto;
    padding: 4rem 2.5rem 2rem;
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: center;
    gap: 3rem;
    width: 100%;
  }

  /* Left content */
  .hero-content {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .hero-badge {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    background: #F0FDF4;
    border: 1px solid #BBF7D0;
    color: #16A34A;
    border-radius: 999px;
    padding: 0.35rem 0.875rem;
    font-size: 0.78rem;
    font-weight: 600;
    width: fit-content;
  }

  .badge-dot {
    width: 7px;
    height: 7px;
    background: #16A34A;
    border-radius: 50%;
    animation: pulse-dot 2s infinite;
  }

  @keyframes pulse-dot {
    0%, 100% { opacity: 1; transform: scale(1); }
    50% { opacity: 0.6; transform: scale(1.3); }
  }

  .hero-title {
    font-size: 3.25rem;
    font-weight: 800;
    line-height: 1.15;
    color: #111827;
    letter-spacing: -0.02em;
    font-family: var(--font-title);
  }

  .title-accent { color: var(--color-primary); }

  .hero-desc {
    font-size: 1rem;
    color: var(--text-muted);
    line-height: 1.7;
    max-width: 420px;
  }

  .hero-actions {
    display: flex;
    align-items: center;
    gap: 1rem;
    flex-wrap: wrap;
  }

  .btn-primary-hero {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1.75rem;
    background: var(--color-primary);
    color: #FFFFFF;
    border-radius: 8px;
    font-family: var(--font-title);
    font-size: 0.95rem;
    font-weight: 700;
    text-decoration: none;
    transition: background 0.15s, box-shadow 0.15s;
  }

  .btn-primary-hero:hover {
    background: var(--color-primary-hover);
    box-shadow: 0 4px 16px rgba(20, 184, 166, 0.35);
  }

  .btn-outline-hero {
    display: inline-flex;
    align-items: center;
    padding: 0.75rem 1.5rem;
    background: transparent;
    border: 1.5px solid #D1D5DB;
    color: var(--text-secondary);
    border-radius: 8px;
    font-family: var(--font-title);
    font-size: 0.95rem;
    font-weight: 600;
    text-decoration: none;
    transition: border-color 0.15s, background 0.15s;
  }

  .btn-outline-hero:hover {
    border-color: var(--color-primary);
    color: var(--color-primary);
    background: rgba(20, 184, 166, 0.04);
  }

  /* Stats */
  .stats-row {
    display: flex;
    align-items: center;
    gap: 1.5rem;
    padding-top: 0.5rem;
  }

  .stat-item {
    display: flex;
    flex-direction: column;
    gap: 0.15rem;
  }

  .stat-num {
    font-size: 1.35rem;
    font-weight: 800;
    color: var(--text-primary);
    font-family: var(--font-title);
    line-height: 1;
  }

  .stat-label {
    font-size: 0.75rem;
    color: var(--text-muted);
  }

  .stat-divider {
    width: 1px;
    height: 32px;
    background: #E5E7EB;
  }

  /* Right visual */
  .hero-visual {
    position: relative;
    height: 500px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  /* Background blobs */
  .blob-green {
    position: absolute;
    width: 340px;
    height: 340px;
    background: #A3E635;
    border-radius: 60% 40% 55% 45% / 50% 60% 40% 50%;
    opacity: 0.15;
    top: 50%;
    left: 50%;
    transform: translate(-30%, -50%);
    z-index: 0;
  }

  .blob-teal {
    position: absolute;
    width: 200px;
    height: 200px;
    background: var(--color-primary);
    border-radius: 50%;
    opacity: 0.12;
    top: 15%;
    right: 5%;
    z-index: 0;
  }

  /* Medical illustration */
  .medical-illustration {
    position: relative;
    z-index: 1;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: flex-end;
    justify-content: center;
  }

  .illus-svg {
    width: 100%;
    height: 100%;
    max-height: 480px;
    object-fit: contain;
    filter: drop-shadow(0 8px 32px rgba(0, 0, 0, 0.08));
  }

  /* Floating cards */
  .float-card {
    position: absolute;
    z-index: 2;
    background: #FFFFFF;
    border: 1px solid #E5E7EB;
    border-radius: 12px;
    padding: 0.75rem 1rem;
    display: flex;
    align-items: center;
    gap: 0.625rem;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
    animation: float-card 4s ease-in-out infinite;
  }

  .card-top-left {
    top: 8%;
    left: -4%;
    animation-delay: 0s;
  }

  .card-top-right {
    top: 12%;
    right: -2%;
    animation-delay: 1.5s;
  }

  .card-bottom-left {
    bottom: 14%;
    left: -2%;
    animation-delay: 0.8s;
  }

  @keyframes float-card {
    0%, 100% { transform: translateY(0); }
    50% { transform: translateY(-6px); }
  }

  .float-icon {
    width: 36px;
    height: 36px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .float-icon.teal { background: #CCFBF1; color: #0D9488; }
  .float-icon.lime { background: #ECFCCB; color: #65A30D; }

  .float-label {
    font-size: 0.8rem;
    font-weight: 700;
    color: var(--text-primary);
    font-family: var(--font-title);
    white-space: nowrap;
  }

  .float-sub {
    font-size: 0.72rem;
    color: var(--text-muted);
    white-space: nowrap;
  }

  /* ── Features strip ── */
  .features-strip {
    background: #F9FAFB;
    border-top: 1px solid #F3F4F6;
    padding: 2.5rem 2.5rem;
  }

  .features-inner {
    max-width: 1200px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 2rem;
  }

  .feat-item {
    display: flex;
    align-items: flex-start;
    gap: 1rem;
  }

  .feat-icon {
    width: 44px;
    height: 44px;
    min-width: 44px;
    border-radius: 10px;
    background: #CCFBF1;
    color: #0D9488;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .feat-icon.lime-icon {
    background: #ECFCCB;
    color: #65A30D;
  }

  .feat-item h3 {
    font-size: 0.9rem;
    font-weight: 700;
    color: var(--text-primary);
    font-family: var(--font-title);
    margin-bottom: 0.25rem;
  }

  .feat-item p {
    font-size: 0.8rem;
    color: var(--text-muted);
    line-height: 1.5;
  }

  /* Footer */
  .landing-footer {
    padding: 1.5rem 2.5rem;
    text-align: center;
    color: var(--text-muted);
    font-size: 0.8rem;
    border-top: 1px solid #F3F4F6;
    background: #FFFFFF;
  }

  /* ── Responsive ── */
  @media (max-width: 1024px) {
    .features-inner { grid-template-columns: repeat(2, 1fr); }
  }

  @media (max-width: 768px) {
    .hero {
      grid-template-columns: 1fr;
      padding: 2.5rem 1.5rem 1.5rem;
      gap: 2rem;
      text-align: center;
    }

    .hero-content { align-items: center; }
    .hero-badge { margin: 0 auto; }
    .hero-title { font-size: 2.25rem; }
    .hero-desc { font-size: 0.95rem; }
    .hero-actions { justify-content: center; }
    .stats-row { justify-content: center; }

    .hero-visual {
      height: 320px;
      order: -1;
    }

    .float-card { display: none; }

    .nav-links { display: none; }
    .nav-inner { padding: 0 1.5rem; }

    .features-inner { grid-template-columns: 1fr 1fr; gap: 1.25rem; }
    .features-strip { padding: 2rem 1.5rem; }
  }

  @media (max-width: 480px) {
    .features-inner { grid-template-columns: 1fr; }
    .hero-title { font-size: 2rem; }
  }
</style>
