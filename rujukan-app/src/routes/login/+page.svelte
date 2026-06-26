<!-- src/routes/login/+page.svelte -->
<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import kemkesLogo from "$lib/assets/kemkes.png";

  let username = $state("");
  let password = $state("");
  let showPassword = $state(false);
  let rememberMe = $state(false);
  let errorMessage = $state("");
  let isLoading = $state(false);

  onMount(() => {
    if (localStorage.getItem("isLoggedIn") === "true") {
      goto("/dashboard");
    }
  });

  const handleLogin = async (e: SubmitEvent) => {
    e.preventDefault();
    errorMessage = "";

    if (!username || !password) {
      errorMessage = "Username dan Password wajib diisi.";
      return;
    }

    isLoading = true;
    await new Promise((resolve) => setTimeout(resolve, 800));

    if (username === "12345" && password === "12345") {
      localStorage.setItem("isLoggedIn", "true");
      localStorage.setItem("username", "dr. Eka Pratama");
      localStorage.setItem("faskes", "RSUD Labuang Baji");
      localStorage.setItem("role", "Dokter Spesialis");
      goto("/dashboard");
    } else {
      errorMessage =
        'Kredensial tidak valid. Gunakan username dan password "12345".';
      isLoading = false;
    }
  };
</script>

<svelte:head>
  <title>Masuk | Rujukan Medis</title>
</svelte:head>

<div class="page-bg">
  <!-- Decorative wave background -->
  <svg
    class="wave-tl"
    viewBox="0 0 300 300"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    aria-hidden="true"
  >
    <path
      d="M-50 200 Q50 100 150 150 T350 100"
      stroke="#14B8A6"
      stroke-width="1.5"
      fill="none"
      opacity="0.35"
    />
    <path
      d="M-80 240 Q30 130 140 180 T380 130"
      stroke="#14B8A6"
      stroke-width="1"
      fill="none"
      opacity="0.25"
    />
    <path
      d="M-30 160 Q80 60 180 110 T400 60"
      stroke="#A3E635"
      stroke-width="1"
      fill="none"
      opacity="0.2"
    />
    <path
      d="M-60 280 Q60 160 170 220 T420 170"
      stroke="#14B8A6"
      stroke-width="0.8"
      fill="none"
      opacity="0.2"
    />
  </svg>
  <svg
    class="wave-br"
    viewBox="0 0 300 300"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    aria-hidden="true"
  >
    <path
      d="M350 100 Q250 200 150 150 T-50 200"
      stroke="#A3E635"
      stroke-width="1.5"
      fill="none"
      opacity="0.3"
    />
    <path
      d="M380 130 Q270 240 160 180 T-80 240"
      stroke="#14B8A6"
      stroke-width="1"
      fill="none"
      opacity="0.2"
    />
    <path
      d="M400 60 Q280 160 180 110 T-30 160"
      stroke="#A3E635"
      stroke-width="1"
      fill="none"
      opacity="0.18"
    />
    <path
      d="M420 170 Q290 280 170 220 T-60 280"
      stroke="#14B8A6"
      stroke-width="0.8"
      fill="none"
      opacity="0.18"
    />
  </svg>

  <div class="login-center">
    <!-- Logo -->
    <div class="logo-block">
      <img src={kemkesLogo} alt="Logo Kemenkes" class="kemkes-logo" />
      <span class="kemkes-label">Kemenkes</span>
    </div>

    <!-- Heading -->
    <div class="heading-block">
      <h1>Selamat Datang Kembali</h1>
      <p>Masuk untuk melanjutkan ke dashboard Anda</p>
    </div>

    <!-- Card -->
    <div class="login-card">
      {#if errorMessage}
        <div class="alert-error">
          <svg
            viewBox="0 0 24 24"
            fill="none"
            width="16"
            height="16"
            stroke="currentColor"
            stroke-width="2"
            style="flex-shrink:0"
          >
            <circle cx="12" cy="12" r="10" /><line
              x1="12"
              y1="8"
              x2="12"
              y2="12"
            /><line x1="12" y1="16" x2="12.01" y2="16" />
          </svg>
          <span>{errorMessage}</span>
        </div>
      {/if}

      <form onsubmit={handleLogin}>
        <!-- Email / NIK field -->
        <div class="field">
          <label for="username">Email atau NIK</label>
          <div class="input-wrap">
            <svg
              class="field-ic"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.8"
            >
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" /><circle
                cx="12"
                cy="7"
                r="4"
              />
            </svg>
            <input
              type="text"
              id="username"
              placeholder="Masukkan email atau NIK Anda"
              bind:value={username}
              disabled={isLoading}
              autocomplete="username"
            />
          </div>
        </div>

        <!-- Password field -->
        <div class="field">
          <div class="field-label-row">
            <label for="password">Kata Sandi</label>
            <a
              href="/login"
              class="forgot-link"
              tabindex="-1"
              onclick={(e) => e.preventDefault()}>Lupa kata sandi?</a
            >
          </div>
          <div class="input-wrap">
            <svg
              class="field-ic"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.8"
            >
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2" /><path
                d="M7 11V7a5 5 0 0 1 10 0v4"
              />
            </svg>
            <input
              type={showPassword ? "text" : "password"}
              id="password"
              placeholder="Masukkan kata sandi Anda"
              bind:value={password}
              disabled={isLoading}
              autocomplete="current-password"
            />
            <button
              type="button"
              class="eye-btn"
              onclick={() => (showPassword = !showPassword)}
              aria-label={showPassword
                ? "Sembunyikan kata sandi"
                : "Tampilkan kata sandi"}
            >
              {#if showPassword}
                <svg
                  viewBox="0 0 24 24"
                  fill="none"
                  width="18"
                  height="18"
                  stroke="currentColor"
                  stroke-width="1.8"
                >
                  <path
                    d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"
                  />
                  <path
                    d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"
                  />
                  <line x1="1" y1="1" x2="23" y2="23" />
                </svg>
              {:else}
                <svg
                  viewBox="0 0 24 24"
                  fill="none"
                  width="18"
                  height="18"
                  stroke="currentColor"
                  stroke-width="1.8"
                >
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
                  <circle cx="12" cy="12" r="3" />
                </svg>
              {/if}
            </button>
          </div>
        </div>

        <!-- Remember me -->
        <label class="remember-row">
          <input type="checkbox" bind:checked={rememberMe} />
          <span>Ingat saya</span>
        </label>

        <!-- Submit -->
        <button type="submit" class="submit-btn" disabled={isLoading}>
          {#if isLoading}
            <div class="spinner"></div>
            <span>Memproses...</span>
          {:else}
            Masuk
          {/if}
        </button>
      </form>

      <div class="card-footer">
        Belum punya akun? <a href="/register">Daftar di sini</a>
      </div>
    </div>

    <!-- Security badge -->
    <div class="security-badge">
      <div class="shield-wrap">
        <svg
          viewBox="0 0 24 24"
          fill="none"
          width="22"
          height="22"
          stroke="currentColor"
          stroke-width="2"
        >
          <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
        </svg>
      </div>
      <div>
        <p class="badge-title">Sistem ini aman dan terlindungi</p>
        <p class="badge-desc">
          Data Anda dienkripsi dan kami menjaga kerahasiaan informasi pribadi
          Anda.
        </p>
      </div>
    </div>
  </div>
</div>

<style>
  /* ── Page background ── */
  .page-bg {
    min-height: 100vh;
    background: #f0f4f8;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem 1rem;
    position: relative;
    overflow: hidden;
  }

  /* Decorative waves */
  .wave-tl {
    position: absolute;
    top: 0;
    left: 0;
    width: 380px;
    height: 380px;
    pointer-events: none;
  }

  .wave-br {
    position: absolute;
    bottom: 0;
    right: 0;
    width: 380px;
    height: 380px;
    pointer-events: none;
  }

  /* ── Center column ── */
  .login-center {
    position: relative;
    z-index: 1;
    width: 100%;
    max-width: 420px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0;
  }

  /* ── Logo ── */
  .logo-block {
    display: flex;
    align-items: center;
    gap: 0.6rem;
    margin-bottom: 1.5rem;
  }

  .kemkes-logo {
    width: 38px;
    height: 38px;
    object-fit: contain;
  }

  .kemkes-label {
    font-family: var(--font-title);
    font-size: 1.25rem;
    font-weight: 700;
    color: var(--text-primary);
  }

  /* ── Heading ── */
  .heading-block {
    text-align: center;
    margin-bottom: 1.75rem;
  }

  .heading-block h1 {
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--text-primary);
    margin-bottom: 0.35rem;
  }

  .heading-block p {
    font-size: 0.875rem;
    color: var(--text-muted);
  }

  /* ── Card ── */
  .login-card {
    width: 100%;
    background: #ffffff;
    border: 1px solid #e5e7eb;
    border-radius: 16px;
    padding: 2rem 1.75rem;
    box-shadow: 0 2px 16px rgba(0, 0, 0, 0.07);
  }

  /* Alert */
  .alert-error {
    background: #fef2f2;
    border: 1px solid #fecaca;
    border-radius: 8px;
    color: #dc2626;
    padding: 0.7rem 0.875rem;
    margin-bottom: 1.25rem;
    display: flex;
    align-items: flex-start;
    gap: 0.5rem;
    font-size: 0.825rem;
    line-height: 1.45;
    animation: fadeIn 0.2s ease;
  }

  /* ── Fields ── */
  .field {
    margin-bottom: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.4rem;
  }

  .field label {
    font-size: 0.82rem;
    font-weight: 600;
    color: var(--text-secondary);
    font-family: var(--font-title);
  }

  .field-label-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .forgot-link {
    font-size: 0.8rem;
    color: var(--color-primary);
    font-weight: 600;
    transition: color var(--transition-fast);
  }

  .forgot-link:hover {
    color: var(--color-primary-hover);
  }

  .input-wrap {
    position: relative;
    display: flex;
    align-items: center;
  }

  .field-ic {
    position: absolute;
    left: 0.875rem;
    width: 17px;
    height: 17px;
    color: #9ca3af;
    pointer-events: none;
    flex-shrink: 0;
    transition: color var(--transition-fast);
  }

  .input-wrap:focus-within .field-ic {
    color: var(--color-primary);
  }

  .input-wrap input {
    width: 100%;
    height: 46px;
    padding: 0 1rem 0 2.75rem;
    background: #ffffff;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    font-family: var(--font-body);
    font-size: 0.9rem;
    color: #374151;
    transition:
      border-color 0.15s,
      box-shadow 0.15s;
  }

  .input-wrap input::placeholder {
    color: #9ca3af;
  }

  .input-wrap input:focus {
    outline: none;
    border-color: #14b8a6;
    box-shadow: 0 0 0 3px rgba(20, 184, 166, 0.1);
  }

  .input-wrap input:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  /* Eye toggle */
  .eye-btn {
    position: absolute;
    right: 0.875rem;
    background: none;
    border: none;
    color: #9ca3af;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
    transition: color 0.15s;
  }
  .eye-btn:hover {
    color: #6b7280;
  }

  /* Remember me */
  .remember-row {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-size: 0.875rem;
    color: var(--text-secondary);
    cursor: pointer;
    margin-bottom: 1.25rem;
    user-select: none;
  }

  .remember-row input[type="checkbox"] {
    width: 15px;
    height: 15px;
    accent-color: var(--color-primary);
    cursor: pointer;
  }

  /* Submit */
  .submit-btn {
    width: 100%;
    height: 48px;
    background: #14b8a6;
    color: #ffffff;
    border: none;
    border-radius: 8px;
    font-family: var(--font-title);
    font-weight: 700;
    font-size: 1rem;
    cursor: pointer;
    transition:
      background 0.15s,
      box-shadow 0.15s;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
  }

  .submit-btn:hover:not(:disabled) {
    background: #0f9488;
    box-shadow: 0 2px 10px rgba(20, 184, 166, 0.3);
  }

  .submit-btn:disabled {
    opacity: 0.7;
    cursor: not-allowed;
  }

  .spinner {
    width: 16px;
    height: 16px;
    border: 2px solid rgba(255, 255, 255, 0.35);
    border-top-color: #fff;
    border-radius: 50%;
    animation: spin 0.75s linear infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  /* Divider */
  .divider {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    margin: 1.25rem 0;
    color: #9ca3af;
    font-size: 0.8rem;
  }

  .divider::before,
  .divider::after {
    content: "";
    flex: 1;
    height: 1px;
    background: #e5e7eb;
  }

  /* SSO row */
  .sso-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 0.75rem;
    margin-bottom: 0.25rem;
  }

  .sso-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    height: 42px;
    background: #ffffff;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    font-family: var(--font-title);
    font-size: 0.85rem;
    font-weight: 600;
    color: #374151;
    cursor: pointer;
    transition:
      border-color 0.15s,
      background 0.15s;
  }

  .sso-btn:hover {
    background: #f9fafb;
    border-color: #d1d5db;
  }

  /* Card footer */
  .card-footer {
    text-align: center;
    margin-top: 1.25rem;
    font-size: 0.875rem;
    color: #6b7280;
    border-top: 1px solid #f3f4f6;
    padding-top: 1.25rem;
  }

  .card-footer a {
    color: #14b8a6;
    font-weight: 600;
    transition: color 0.15s;
  }

  .card-footer a:hover {
    color: #0f9488;
  }

  /* ── Security badge ── */
  .security-badge {
    width: 100%;
    background: #ffffff;
    border: 1px solid #e5e7eb;
    border-radius: 12px;
    padding: 1rem 1.25rem;
    margin-top: 1rem;
    display: flex;
    align-items: flex-start;
    gap: 1rem;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
  }

  .shield-wrap {
    width: 40px;
    height: 40px;
    min-width: 40px;
    border-radius: 10px;
    background: #ccfbf1;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #14b8a6;
  }

  .badge-title {
    font-size: 0.85rem;
    font-weight: 600;
    color: #374151;
    margin-bottom: 0.2rem;
  }

  .badge-desc {
    font-size: 0.775rem;
    color: #6b7280;
    line-height: 1.5;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(6px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .animate-fade-in {
    animation: fadeIn 0.35s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }

  @media (max-width: 480px) {
    .login-card {
      padding: 1.5rem 1.25rem;
    }
    .sso-row {
      grid-template-columns: 1fr;
    }
  }
</style>
