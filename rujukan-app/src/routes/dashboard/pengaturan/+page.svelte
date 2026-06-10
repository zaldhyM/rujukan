<!-- src/routes/dashboard/pengaturan/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';

  // Svelte 5 state
  let activeTab = $state('profile'); // profile, notification, password
  let showToast = $state(false);
  let toastMessage = $state('');

  // Profile Form States
  let username = $state('');
  let nama = $state('');
  let faskes = $state('');
  let role = $state('');
  let nik = $state('7371100203990001');

  // Telegram Integration States
  let telegramStatus = $state('Disconnected'); // Disconnected, Connected
  let telegramId = $state('');

  // Password States
  let currentPassword = $state('');
  let newPassword = $state('');
  let confirmNewPassword = $state('');

  // System Preference States
  let emailAlerts = $state(true);
  let audioAlerts = $state(true);

  onMount(() => {
    username = localStorage.getItem('username') || 'dr. Eka Pratama';
    nama = username;
    faskes = localStorage.getItem('faskes') || 'RSUD Labuang Baji';
    role = localStorage.getItem('role') || 'Dokter Spesialis';
    
    const savedTelegramId = localStorage.getItem('telegramId');
    if (savedTelegramId) {
      telegramId = savedTelegramId;
      telegramStatus = 'Connected';
    }
  });

  const saveProfile = (e: SubmitEvent) => {
    e.preventDefault();
    localStorage.setItem('username', nama);
    localStorage.setItem('faskes', faskes);
    localStorage.setItem('role', role);

    triggerToast('Profil Fasilitas Kesehatan berhasil diperbarui!');
  };

  const savePassword = (e: SubmitEvent) => {
    e.preventDefault();
    if (!currentPassword || !newPassword || !confirmNewPassword) {
      alert('Semua kolom kata sandi wajib diisi.');
      return;
    }
    if (newPassword !== confirmNewPassword) {
      alert('Konfirmasi kata sandi baru tidak sesuai.');
      return;
    }

    currentPassword = '';
    newPassword = '';
    confirmNewPassword = '';
    triggerToast('Kata sandi berhasil diperbarui!');
  };

  const handleConnectTelegram = (e: SubmitEvent) => {
    e.preventDefault();
    if (!telegramId.trim()) return;

    localStorage.setItem('telegramId', telegramId);
    telegramStatus = 'Connected';
    triggerToast('Notifikasi Telegram Bot berhasil diaktifkan!');
  };

  const handleDisconnectTelegram = () => {
    localStorage.removeItem('telegramId');
    telegramId = '';
    telegramStatus = 'Disconnected';
    triggerToast('Notifikasi Telegram Bot dinonaktifkan.');
  };

  const triggerToast = (msg: string) => {
    toastMessage = msg;
    showToast = true;
    setTimeout(() => {
      showToast = false;
    }, 3000);
  };
</script>

<svelte:head>
  <title>Pengaturan | RUJUKAN MEDIS</title>
</svelte:head>

<div class="pengaturan-container animate-fade-in">
  <!-- Toast Notification -->
  {#if showToast}
    <div class="toast-notification glass-panel">
      <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="var(--color-success)" stroke-width="2.5">
        <polyline points="20 6 9 17 4 12"></polyline>
      </svg>
      <span>{toastMessage}</span>
    </div>
  {/if}

  <div class="page-header">
    <div class="header-left">
      <span class="breadcrumb">DASHBOARD / PENGATURAN</span>
      <h1>Pengaturan Sistem</h1>
      <p class="subtitle">Sesuaikan profil medis Anda, konfigurasi telegram bot rujukan, dan ubah kata sandi keamanan</p>
    </div>
  </div>

  <div class="settings-layout">
    <!-- Sidebar Settings Navigation -->
    <div class="glass-panel settings-sidebar">
      <button 
        class="settings-nav-btn" 
        class:active={activeTab === 'profile'} 
        onclick={() => activeTab = 'profile'}
      >
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
          <circle cx="12" cy="7" r="4"></circle>
        </svg>
        <span>Profil Medis & Faskes</span>
      </button>

      <button 
        class="settings-nav-btn" 
        class:active={activeTab === 'notification'} 
        onclick={() => activeTab = 'notification'}
      >
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2">
          <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
          <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
        </svg>
        <span>Integrasi Notifikasi</span>
      </button>

      <button 
        class="settings-nav-btn" 
        class:active={activeTab === 'password'} 
        onclick={() => activeTab = 'password'}
      >
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2">
          <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
          <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
        </svg>
        <span>Keamanan & Password</span>
      </button>
    </div>

    <!-- Settings Contents Panel -->
    <div class="glass-panel settings-content">
      {#if activeTab === 'profile'}
        <div class="content-tab animate-fade-in">
          <h2>Profil Medis & Faskes</h2>
          <p class="tab-desc text-secondary">Perbarui data tenaga medis dan instansi layanan rujukan Anda</p>
          
          <form onsubmit={saveProfile} class="settings-form">
            <div class="form-row">
              <div class="form-group flex-1">
                <label class="form-label" for="profName">Nama Lengkap & Gelar</label>
                <input type="text" id="profName" class="form-input" bind:value={nama} required />
              </div>
              <div class="form-group flex-1">
                <label class="form-label" for="profNik">NIK Tenaga Medis</label>
                <input type="text" id="profNik" class="form-input" bind:value={nik} readonly />
              </div>
            </div>

            <div class="form-row">
              <div class="form-group flex-1">
                <label class="form-label" for="profFaskes">Instansi Fasilitas Kesehatan</label>
                <input type="text" id="profFaskes" class="form-input" bind:value={faskes} required />
              </div>
              <div class="form-group flex-1">
                <label class="form-label" for="profRole">Jabatan / Role Medis</label>
                <input type="text" id="profRole" class="form-input" bind:value={role} required />
              </div>
            </div>

            <button type="submit" class="btn btn-primary self-start">Simpan Perubahan</button>
          </form>
        </div>
      {/if}

      {#if activeTab === 'notification'}
        <div class="content-tab animate-fade-in">
          <h2>Integrasi Notifikasi Real-time</h2>
          <p class="tab-desc text-secondary">Konfigurasi jalur pemberitahuan instan saat ada rujukan masuk atau keluar</p>

          <!-- Telegram Bot section -->
          <div class="integration-card">
            <div class="integration-header">
              <div class="brand-info">
                <div class="telegram-icon-wrap">
                  <svg viewBox="0 0 24 24" width="24" height="24" fill="currentColor">
                    <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm4.64 6.8c-.15 1.58-.8 5.42-1.13 7.19-.14.75-.42 1-.68 1.03-.58.05-1.02-.38-1.58-.75-.88-.58-1.38-.94-2.23-1.5-.99-.65-.35-1.01.22-1.59.15-.15 2.71-2.48 2.76-2.69.01-.03.01-.14-.07-.2-.08-.06-.19-.04-.27-.02-.12.02-1.96 1.25-5.54 3.69-.52.36-1 .53-1.42.52-.47-.01-1.37-.26-2.03-.48-.82-.27-1.47-.42-1.42-.88.03-.24.35-.49.97-.74 3.79-1.65 6.31-2.74 7.57-3.27 3.61-1.51 4.36-1.77 4.85-1.78.11 0 .35.03.5.16.13.12.17.28.19.39.02.12.02.26.01.4z"/>
                  </svg>
                </div>
                <div>
                  <h3 class="text-white">Telegram Bot Alert</h3>
                  <p class="text-secondary font-size-sm">Dapatkan notifikasi rujukan baru secara otomatis di Telegram Anda</p>
                </div>
              </div>

              {#if telegramStatus === 'Connected'}
                <span class="badge badge-success">Terhubung</span>
              {:else}
                <span class="badge badge-danger">Belum Aktif</span>
              {/if}
            </div>

            <div class="integration-body">
              {#if telegramStatus === 'Disconnected'}
                <form onsubmit={handleConnectTelegram} class="tg-connect-form">
                  <p class="info-text text-secondary">
                    Untuk mulai menerima notifikasi, ikuti langkah berikut:<br>
                    1. Cari bot <strong>@RujukanMedisAlertBot</strong> di aplikasi Telegram.<br>
                    2. Kirim pesan <code>/start</code> untuk mendapatkan <strong>ID Telegram Chat</strong> Anda.<br>
                    3. Masukkan ID Telegram tersebut pada formulir di bawah ini.
                  </p>
                  <div class="form-group mb-0">
                    <label class="form-label" for="tgId">ID Telegram Chat</label>
                    <div class="input-inline">
                      <input 
                        type="text" 
                        id="tgId" 
                        class="form-input flex-1" 
                        placeholder="Contoh: 184729482" 
                        bind:value={telegramId}
                        required 
                      />
                      <button type="submit" class="btn btn-primary">Hubungkan Bot</button>
                    </div>
                  </div>
                </form>
              {:else}
                <div class="connected-info">
                  <p>Bot terhubung ke ID Telegram Chat: <strong>{telegramId}</strong></p>
                  <button class="btn btn-danger btn-sm mt-3" onclick={handleDisconnectTelegram}>Putuskan Sambungan</button>
                </div>
              {/if}
            </div>
          </div>

          <!-- Alert settings -->
          <div class="preferences-box mt-4">
            <h3 class="text-white mb-3">Preferensi Sistem</h3>
            
            <div class="preference-item">
              <div class="pref-desc">
                <span class="pref-title text-white">Notifikasi Email</span>
                <span class="pref-sub text-muted">Kirim rekap rujukan mingguan ke email terdaftar</span>
              </div>
              <label class="switch">
                <input type="checkbox" bind:checked={emailAlerts}>
                <span class="slider"></span>
              </label>
            </div>

            <div class="preference-item">
              <div class="pref-desc">
                <span class="pref-title text-white">Alert Suara</span>
                <span class="pref-sub text-muted">Putar suara notifikasi saat ada rujukan masuk yang mendesak (Cito)</span>
              </div>
              <label class="switch">
                <input type="checkbox" bind:checked={audioAlerts}>
                <span class="slider"></span>
              </label>
            </div>
          </div>
        </div>
      {/if}

      {#if activeTab === 'password'}
        <div class="content-tab animate-fade-in">
          <h2>Keamanan & Password</h2>
          <p class="tab-desc text-secondary">Ubah kata sandi secara berkala demi keamanan data rekam medis pasien</p>

          <form onsubmit={savePassword} class="settings-form">
            <div class="form-group">
              <label class="form-label" for="currPass">Kata Sandi Saat Ini</label>
              <input 
                type="password" 
                id="currPass" 
                class="form-input max-w-400" 
                placeholder="Masukkan kata sandi lama" 
                bind:value={currentPassword}
                required 
              />
            </div>

            <div class="form-group">
              <label class="form-label" for="newPass">Kata Sandi Baru</label>
              <input 
                type="password" 
                id="newPass" 
                class="form-input max-w-400" 
                placeholder="Minimal 6 karakter" 
                bind:value={newPassword}
                required 
              />
            </div>

            <div class="form-group">
              <label class="form-label" for="confirmNewPass">Konfirmasi Kata Sandi Baru</label>
              <input 
                type="password" 
                id="confirmNewPass" 
                class="form-input max-w-400" 
                placeholder="Ulangi kata sandi baru" 
                bind:value={confirmNewPassword}
                required 
              />
            </div>

            <button type="submit" class="btn btn-primary self-start">Perbarui Password</button>
          </form>
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  .pengaturan-container {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .page-header h1 {
    font-size: 2rem;
    color: white;
  }

  /* Toast Notification */
  .toast-notification {
    position: fixed;
    top: 24px;
    right: 24px;
    background: rgba(18, 27, 45, 0.9);
    border: 1px solid var(--color-success);
    padding: 1rem 1.5rem;
    border-radius: var(--border-radius-md);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    gap: 0.75rem;
    z-index: 2000;
    color: white;
    font-weight: 500;
    animation: slideInRight 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  @keyframes slideInRight {
    from { transform: translateX(120%); opacity: 0; }
    to { transform: translateX(0); opacity: 1; }
  }

  /* Layout Settings */
  .settings-layout {
    display: grid;
    grid-template-columns: 280px 1fr;
    gap: 1.5rem;
  }

  @media (max-width: 850px) {
    .settings-layout {
      grid-template-columns: 1fr;
    }
  }

  .settings-sidebar {
    padding: 1.25rem;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    height: fit-content;
  }

  .settings-nav-btn {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    background: none;
    border: 1px solid transparent;
    color: var(--text-secondary);
    padding: 0.85rem 1.25rem;
    border-radius: var(--border-radius-sm);
    cursor: pointer;
    font-family: var(--font-title);
    font-weight: 600;
    transition: all var(--transition-fast);
    text-align: left;
  }

  .settings-nav-btn:hover {
    color: white;
    background: rgba(255, 255, 255, 0.02);
  }

  .settings-nav-btn.active {
    background: rgba(0, 242, 254, 0.05);
    border-color: var(--border-glass-focus);
    color: var(--color-primary);
  }

  .settings-content {
    padding: 2.5rem;
  }

  .content-tab h2 {
    font-size: 1.5rem;
    color: white;
    margin-bottom: 0.25rem;
  }

  .tab-desc {
    font-size: 0.9rem;
    margin-bottom: 2rem;
  }

  .settings-form {
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

  .max-w-400 {
    max-width: 400px;
  }

  .self-start {
    align-self: flex-start;
  }

  @media (max-width: 640px) {
    .form-row {
      flex-direction: column;
      gap: 1.25rem;
    }
    
    .settings-content {
      padding: 1.5rem;
    }
  }

  /* Telegram bot integration box */
  .integration-card {
    background: rgba(255, 255, 255, 0.01);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-md);
    overflow: hidden;
  }

  .integration-header {
    background: rgba(255, 255, 255, 0.02);
    border-bottom: 1px solid var(--border-glass);
    padding: 1.25rem 1.5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
    flex-wrap: wrap;
  }

  .brand-info {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .telegram-icon-wrap {
    width: 42px;
    height: 42px;
    background: rgba(0, 136, 204, 0.15);
    color: #0088cc;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .font-size-sm {
    font-size: 0.8rem;
  }

  .integration-body {
    padding: 1.5rem;
  }

  .tg-connect-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .info-text {
    font-size: 0.875rem;
    line-height: 1.6;
  }

  .info-text code {
    background: rgba(255, 255, 255, 0.1);
    padding: 0.1rem 0.4rem;
    border-radius: 4px;
    color: var(--color-primary);
    font-family: monospace;
    font-size: 0.9rem;
  }

  .input-inline {
    display: flex;
    gap: 0.75rem;
  }

  @media (max-width: 480px) {
    .input-inline {
      flex-direction: column;
    }
  }

  .mb-0 {
    margin-bottom: 0;
  }

  .mt-3 {
    margin-top: 0.75rem;
  }

  .mt-4 {
    margin-top: 1.5rem;
  }

  .connected-info {
    font-size: 0.95rem;
    color: white;
  }

  /* Preference Switch / Sliders */
  .preferences-box {
    background: rgba(255, 255, 255, 0.01);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-md);
    padding: 1.5rem;
  }

  .preference-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 0;
    border-bottom: 1px solid rgba(255, 255, 255, 0.03);
    gap: 1.5rem;
  }

  .preference-item:last-child {
    border-bottom: none;
    padding-bottom: 0;
  }

  .pref-desc {
    display: flex;
    flex-direction: column;
    gap: 0.15rem;
  }

  .pref-title {
    font-size: 0.95rem;
    font-weight: 600;
  }

  .pref-sub {
    font-size: 0.8rem;
  }

  /* Switch styling */
  .switch {
    position: relative;
    display: inline-block;
    width: 46px;
    height: 24px;
    flex-shrink: 0;
  }

  .switch input {
    opacity: 0;
    width: 0;
    height: 0;
  }

  .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(255, 255, 255, 0.1);
    transition: .3s;
    border-radius: 24px;
    border: 1px solid var(--border-glass);
  }

  .slider:before {
    position: absolute;
    content: "";
    height: 16px;
    width: 16px;
    left: 3px;
    bottom: 3px;
    background-color: white;
    transition: .3s;
    border-radius: 50%;
  }

  input:checked + .slider {
    background-color: var(--color-primary-hover);
    border-color: var(--color-primary);
  }

  input:checked + .slider:before {
    transform: translateX(22px);
    background-color: #0b111e;
  }
</style>
