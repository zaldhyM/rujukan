<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import Fa from "svelte-fa";
  import {
    faChartLine,
    faInbox,
    faPaperPlane,
    faGear,
    faBell,
    faBars,
    faRightFromBracket,
    faUser,
    faHospital,
    faTimes
  } from "@fortawesome/free-solid-svg-icons";
  import avatar from "$lib/assets/img/undraw_profile.svg";
  import { isLoggedIn, getCurrentUser, getFaskesSettings, logout, getRujukanMasuk } from '$lib/auth';

  let { children } = $props();

  let user: any = $state({ name: '', username: '', role: '' });
  let faskes: any = $state({ name: '', code: '' });
  let isMounted = $state(false);
  let pendingCount = $state(0);

  // Native Svelte UI states
  let showNotificationDropdown = $state(false);
  let showProfileDropdown = $state(false);
  let isSidebarOpen = $state(true);

  onMount(() => {
    if (!isLoggedIn()) {
      window.location.href = '/login';
      return;
    }
    user = getCurrentUser();
    faskes = getFaskesSettings();
    
    // Count pending incoming referrals for notification badge
    const incoming = getRujukanMasuk();
    pendingCount = incoming.filter((r: any) => r.status === 'Pending').length;

    // Responsive initial sidebar check
    if (window.innerWidth < 992) {
      isSidebarOpen = false;
    }

    isMounted = true;
  });

  function handleLogout() {
    logout();
    window.location.href = '/login';
  }

  // Active menu check helper
  function isActive(path: string) {
    if (!isMounted) return false;
    return $page.url.pathname === path;
  }

  function toggleSidebar() {
    isSidebarOpen = !isSidebarOpen;
  }

  // Close dropdowns on window click
  function handleDocumentClick(e: MouseEvent) {
    const target = e.target as HTMLElement;
    if (!target.closest('.dropdown-trigger')) {
      showNotificationDropdown = false;
      showProfileDropdown = false;
    }
  }
</script>

<svelte:head>
  <title>Dashboard SISRUTE - {faskes.name || 'Faskes'}</title>
</svelte:head>

<svelte:window onclick={handleDocumentClick} />

{#if isMounted}
<div class="dashboard-wrapper">
  <!-- Mobile Sidebar Backdrop -->
  {#if isSidebarOpen}
    <div class="sidebar-backdrop d-lg-none" onclick={toggleSidebar}></div>
  {/if}

  <!-- Sidebar -->
  <aside class="sidebar-modern" class:sidebar-hidden={!isSidebarOpen}>
    <div class="sidebar-header">
      <div class="brand-logo">
        <Fa icon={faHospital} size="lg" />
      </div>
      <div class="brand-info">
        <span class="brand-name">SISRUTE LOKAL</span>
        <span class="brand-tag">Integrasi Rujukan</span>
      </div>
      <button class="btn-close-sidebar d-lg-none" onclick={toggleSidebar}>
        <Fa icon={faTimes} />
      </button>
    </div>

    <hr class="sidebar-divider" />

    <nav class="sidebar-menu">
      <div class="menu-label">MENU UTAMA</div>
      
      <a class="menu-item" class:active={isActive('/admin')} href="/admin">
        <span class="menu-icon"><Fa icon={faChartLine} /></span>
        <span class="menu-text">Dashboard</span>
      </a>

      <div class="menu-label mt-4">LAYANAN MEDIS</div>

      <a class="menu-item" class:active={isActive('/admin/rujukan-masuk')} href="/admin/rujukan-masuk">
        <span class="menu-icon"><Fa icon={faInbox} /></span>
        <span class="menu-text">Rujukan Masuk</span>
        {#if pendingCount > 0}
          <span class="badge-count-custom">{pendingCount}</span>
        {/if}
      </a>

      <a class="menu-item" class:active={isActive('/admin/rujukan-keluar')} href="/admin/rujukan-keluar">
        <span class="menu-icon"><Fa icon={faPaperPlane} /></span>
        <span class="menu-text">Rujukan Keluar</span>
      </a>

      <div class="menu-label mt-4">ADMINISTRASI</div>

      <a class="menu-item" class:active={isActive('/admin/pengaturan')} href="/admin/pengaturan">
        <span class="menu-icon"><Fa icon={faGear} /></span>
        <span class="menu-text">Pengaturan</span>
      </a>
    </nav>

    <div class="sidebar-footer">
      <div class="faskes-badge">
        <i class="fas fa-circle-check text-success"></i>
        <span>{faskes.type || 'Puskesmas'} Aktif</span>
      </div>
    </div>
  </aside>

  <!-- Content Wrapper -->
  <div class="content-wrapper" class:content-expanded={!isSidebarOpen}>
    <!-- Topbar Header -->
    <header class="topbar-modern glass-panel">
      <div class="topbar-left">
        <button class="btn-toggle-sidebar" onclick={toggleSidebar}>
          <Fa icon={faBars} />
        </button>
        
        <div class="faskes-info d-none d-sm-flex align-items-center">
          <div class="faskes-icon">
            <Fa icon={faHospital} />
          </div>
          <div>
            <span class="faskes-name">{faskes.name}</span>
            <span class="faskes-code">Kode Instansi: {faskes.code}</span>
          </div>
        </div>
      </div>

      <div class="topbar-right">
        <!-- Notification Dropdown -->
        <div class="dropdown-trigger position-relative">
          <button class="btn-icon" onclick={() => { showNotificationDropdown = !showNotificationDropdown; showProfileDropdown = false; }}>
            <Fa icon={faBell} />
            {#if pendingCount > 0}
              <span class="badge-dot-custom"></span>
            {/if}
          </button>
          
          {#if showNotificationDropdown}
            <div class="dropdown-menu-custom alert-dropdown-custom animate-fade-in shadow-lg">
              <div class="dropdown-header-custom">
                <h6>Notifikasi Terbaru</h6>
                {#if pendingCount > 0}
                  <span class="badge bg-danger rounded-pill">{pendingCount} Baru</span>
                {/if}
              </div>
              <div class="dropdown-body-custom">
                {#if pendingCount > 0}
                  <a class="dropdown-item-custom" href="/admin/rujukan-masuk" onclick={() => showNotificationDropdown = false}>
                    <div class="item-icon bg-warning-light text-warning">
                      <Fa icon={faInbox} />
                    </div>
                    <div class="item-content">
                      <span class="item-title font-weight-bold">Rujukan Baru Masuk</span>
                      <span class="item-desc">Ada {pendingCount} rujukan baru masuk memerlukan peninjauan klinis.</span>
                    </div>
                  </a>
                {:else}
                  <div class="dropdown-empty">
                    <p class="text-muted small mb-0">Tidak ada notifikasi baru</p>
                  </div>
                {/if}
              </div>
            </div>
          {/if}
        </div>

        <div class="topbar-divider"></div>

        <!-- User Profile Dropdown -->
        <div class="dropdown-trigger position-relative">
          <button class="btn-profile" onclick={() => { showProfileDropdown = !showProfileDropdown; showNotificationDropdown = false; }}>
            <img class="user-avatar" src={avatar} alt="Profile" />
            <div class="user-info d-none d-md-flex">
              <span class="user-name">{user.name}</span>
              <span class="user-role">{user.role}</span>
            </div>
          </button>
          
          {#if showProfileDropdown}
            <div class="dropdown-menu-custom profile-dropdown-custom animate-fade-in shadow-lg">
              <div class="profile-header">
                <span class="profile-name">{user.name}</span>
                <span class="profile-username">NIP: {user.username}</span>
              </div>
              <hr class="my-1" style="border-color: var(--border-color);" />
              <div class="dropdown-body-custom py-1">
                <a class="dropdown-item-custom" href="/admin/pengaturan" onclick={() => showProfileDropdown = false}>
                  <Fa icon={faUser} class="mr-2 text-muted" /> Profil Saya
                </a>
                <a class="dropdown-item-custom" href="/admin/pengaturan" onclick={() => showProfileDropdown = false}>
                  <Fa icon={faGear} class="mr-2 text-muted" /> Pengaturan Faskes
                </a>
                <hr class="my-1" style="border-color: var(--border-color);" />
                <button class="dropdown-item-custom text-danger font-weight-bold w-100 bg-transparent border-0 text-left" onclick={handleLogout}>
                  <Fa icon={faRightFromBracket} class="mr-2 text-danger" /> Keluar Aplikasi
                </button>
              </div>
            </div>
          {/if}
        </div>
      </div>
    </header>

    <!-- Main Content Area -->
    <main class="main-content-body container-fluid animate-fade-in">
      {@render children?.()}
    </main>

    <!-- Footer -->
    <footer class="dashboard-footer">
      <div class="text-center py-4">
        <span class="text-muted small">Copyright &copy; SISRUTE LOKAL 2026 - {faskes.name}</span>
      </div>
    </footer>
  </div>
</div>
{/if}

<style>
  .dashboard-wrapper {
    display: flex;
    min-height: 100vh;
    background-color: var(--bg-slate);
    position: relative;
  }

  /* Mobile Sidebar Backdrop */
  .sidebar-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.4);
    backdrop-filter: blur(4px);
    z-index: 100;
  }

  /* Sidebar Design */
  .sidebar-modern {
    width: 260px;
    background-color: var(--sidebar-bg);
    color: #f8fafc;
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    z-index: 101;
    display: flex;
    flex-column: column;
    flex-direction: column;
    padding: 1.5rem 1rem;
    transition: var(--transition);
    border-right: 1px solid rgba(255, 255, 255, 0.05);
  }

  .sidebar-hidden {
    transform: translateX(-260px);
  }

  .sidebar-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    position: relative;
    padding: 0.5rem 0.5rem 1rem 0.5rem;
  }

  .brand-logo {
    width: 42px;
    height: 42px;
    background: linear-gradient(135deg, var(--primary) 0%, #6366f1 100%);
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 4px 12px rgba(79, 70, 229, 0.3);
  }

  .brand-info {
    display: flex;
    flex-direction: column;
  }

  .brand-name {
    font-weight: 800;
    font-size: 1.05rem;
    letter-spacing: 0.02em;
    color: #fff;
  }

  .brand-tag {
    font-size: 0.7rem;
    color: var(--text-light);
    font-weight: 600;
  }

  .btn-close-sidebar {
    position: absolute;
    right: 0;
    background: transparent;
    border: none;
    color: #cbd5e1;
    cursor: pointer;
  }

  .sidebar-divider {
    border-top: 1px solid rgba(255, 255, 255, 0.08);
    margin: 0.75rem 0;
  }

  .sidebar-menu {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 0.35rem;
    padding: 1rem 0;
  }

  .menu-label {
    font-size: 0.65rem;
    font-weight: 700;
    color: var(--text-light);
    letter-spacing: 0.1em;
    padding-left: 0.75rem;
    margin-bottom: 0.25rem;
  }

  .menu-item {
    display: flex;
    align-items: center;
    padding: 0.75rem 1rem;
    border-radius: var(--radius-md);
    color: #cbd5e1;
    text-decoration: none;
    font-size: 0.875rem;
    font-weight: 600;
    transition: var(--transition);
  }

  .menu-item:hover {
    background-color: rgba(255, 255, 255, 0.04);
    color: #fff;
  }

  .menu-item.active {
    background-color: var(--sidebar-active-bg);
    color: #fff;
    border-left: 3.5px solid var(--sidebar-active-border);
  }

  .menu-icon {
    width: 22px;
    display: flex;
    align-items: center;
    font-size: 1rem;
    margin-right: 0.75rem;
    color: #94a3b8;
    transition: var(--transition);
  }

  .menu-item.active .menu-icon {
    color: #818cf8;
  }

  .badge-count-custom {
    background-color: var(--danger);
    color: #fff;
    font-size: 0.7rem;
    font-weight: 800;
    padding: 0.15rem 0.5rem;
    border-radius: var(--radius-pill);
    margin-left: auto;
  }

  .sidebar-footer {
    padding-top: 1rem;
  }

  .faskes-badge {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    background-color: rgba(255, 255, 255, 0.03);
    padding: 0.5rem 0.75rem;
    border-radius: var(--radius-sm);
    font-size: 0.75rem;
    color: #cbd5e1;
    font-weight: 600;
  }

  /* Content wrapper layout */
  .content-wrapper {
    flex: 1;
    margin-left: 260px;
    display: flex;
    flex-direction: column;
    min-width: 0;
    transition: var(--transition);
  }

  .content-expanded {
    margin-left: 0;
  }

  /* Topbar styling */
  .topbar-modern {
    height: 70px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 2rem;
    position: sticky;
    top: 1rem;
    z-index: 99;
    margin: 1rem 1.5rem;
  }

  .btn-toggle-sidebar {
    background: transparent;
    border: none;
    color: var(--text-main);
    font-size: 1.2rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 38px;
    height: 38px;
    border-radius: 50%;
    transition: var(--transition);
  }

  .btn-toggle-sidebar:hover {
    background-color: #f1f5f9;
  }

  .topbar-left {
    display: flex;
    align-items: center;
    gap: 1.5rem;
  }

  .faskes-info {
    gap: 0.75rem;
  }

  .faskes-icon {
    width: 36px;
    height: 36px;
    background-color: var(--primary-light);
    color: var(--primary);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .faskes-name {
    display: block;
    font-weight: 700;
    font-size: 0.875rem;
    color: var(--text-main);
    line-height: 1.2;
  }

  .faskes-code {
    display: block;
    font-size: 0.7rem;
    color: var(--text-muted);
    font-weight: 500;
  }

  .topbar-right {
    display: flex;
    align-items: center;
    gap: 1.25rem;
  }

  .btn-icon {
    background: transparent;
    border: none;
    color: var(--text-muted);
    font-size: 1.15rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 38px;
    height: 38px;
    border-radius: 50%;
    position: relative;
    transition: var(--transition);
  }

  .btn-icon:hover {
    background-color: #f1f5f9;
    color: var(--text-main);
  }

  .badge-dot-custom {
    position: absolute;
    top: 8px;
    right: 8px;
    width: 8px;
    height: 8px;
    background-color: var(--danger);
    border-radius: 50%;
    border: 1.5px solid #fff;
  }

  .topbar-divider {
    width: 1px;
    height: 24px;
    background-color: var(--border-color);
  }

  .btn-profile {
    background: transparent;
    border: none;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    cursor: pointer;
    padding: 0.25rem;
    border-radius: var(--radius-pill);
    transition: var(--transition);
  }

  .btn-profile:hover {
    background-color: #f1f5f9;
  }

  .user-avatar {
    width: 34px;
    height: 34px;
    border-radius: 50%;
    object-fit: cover;
    border: 2px solid #fff;
    box-shadow: var(--shadow-sm);
  }

  .user-info {
    display: flex;
    flex-direction: column;
    text-align: left;
    line-height: 1.2;
  }

  .user-name {
    font-size: 0.85rem;
    font-weight: 700;
    color: var(--text-main);
  }

  .user-role {
    font-size: 0.7rem;
    color: var(--text-muted);
    font-weight: 500;
  }

  /* Dropdown Menu System */
  .dropdown-menu-custom {
    position: absolute;
    right: 0;
    top: 115%;
    background-color: #ffffff;
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    z-index: 1000;
    overflow: hidden;
  }

  .alert-dropdown-custom {
    width: 320px;
  }

  .profile-dropdown-custom {
    width: 220px;
  }

  .dropdown-header-custom {
    background-color: #f8fafc;
    padding: 0.85rem 1.25rem;
    border-bottom: 1px solid var(--border-color);
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .dropdown-header-custom h6 {
    margin: 0;
    font-size: 0.85rem;
    font-weight: 700;
  }

  .dropdown-empty {
    padding: 2rem 1.25rem;
    text-align: center;
  }

  .dropdown-item-custom {
    display: flex;
    align-items: center;
    padding: 0.75rem 1.25rem;
    color: var(--text-main);
    text-decoration: none;
    font-size: 0.85rem;
    font-weight: 500;
    transition: var(--transition);
    gap: 0.75rem;
  }

  .dropdown-item-custom:hover {
    background-color: #f8fafc;
    color: var(--primary);
  }

  .item-icon {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .bg-warning-light {
    background-color: var(--warning-light);
  }

  .item-content {
    display: flex;
    flex-direction: column;
    line-height: 1.3;
  }

  .item-title {
    font-size: 0.8rem;
    color: var(--text-main);
  }

  .item-desc {
    font-size: 0.7rem;
    color: var(--text-muted);
  }

  .profile-header {
    padding: 1rem 1.25rem;
    display: flex;
    flex-direction: column;
  }

  .profile-name {
    font-size: 0.85rem;
    font-weight: 700;
    color: var(--text-main);
  }

  .profile-username {
    font-size: 0.7rem;
    color: var(--text-muted);
    font-weight: 500;
  }

  /* Main content spacing */
  .main-content-body {
    flex: 1;
    padding: 1.5rem 1.5rem;
  }

  /* Footer minimal */
  .dashboard-footer {
    border-top: 1px solid var(--border-color);
    background-color: #fff;
    margin-top: auto;
  }

  /* Responsive Adjustments */
  @media (max-width: 991.98px) {
    .sidebar-modern {
      left: 0;
    }
    .content-wrapper {
      margin-left: 0;
    }
    .topbar-modern {
      padding: 0 1rem;
      margin: 0.5rem 1rem;
    }
  }
</style>
