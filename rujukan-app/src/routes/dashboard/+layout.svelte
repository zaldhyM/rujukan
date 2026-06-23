<!-- src/routes/dashboard/+layout.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/state';

  let { children } = $props();

  let isSidebarOpen = $state(false);
  let username = $state('');
  let faskes = $state('');
  let role = $state('');
  let isLoaded = $state(false);

  // Active path detection
  let currentPath = $derived(page.url.pathname);
  let isRujukanDropdownOpen = $state(false);

  // Automatically open dropdown if on a rujukan subpage
  $effect(() => {
    if (currentPath.startsWith('/dashboard/rujukan-')) {
      isRujukanDropdownOpen = true;
    }
  });

  onMount(() => {
    const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
    if (!isLoggedIn) {
      goto('/login');
    } else {
      username = localStorage.getItem('username') || 'Tenaga Medis';
      faskes = localStorage.getItem('faskes') || 'Fasilitas Kesehatan';
      role = localStorage.getItem('role') || 'Staf Medis';
      isLoaded = true;
    }
  });

  const handleLogout = () => {
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('username');
    localStorage.removeItem('faskes');
    localStorage.removeItem('role');
    goto('/login');
  };

  const toggleSidebar = () => {
    isSidebarOpen = !isSidebarOpen;
  };
</script>

<svelte:head>
  <title>Dashboard | RUJUKAN MEDIS</title>
</svelte:head>

{#if isLoaded}
  <div class="dashboard-wrapper">
    <!-- Sidebar -->
    <aside class="sidebar glass-panel" class:active={isSidebarOpen}>
      <div class="sidebar-header">
        <svg class="sidebar-logo" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M19 10.5V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V10.5M12 3V17M12 3L7.5 7.5M12 3L16.5 7.5" stroke="url(#sideLogoGrad)" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
          <defs>
            <linearGradient id="sideLogoGrad" x1="5" y1="3" x2="19" y2="21" gradientUnits="userSpaceOnUse">
              <stop stop-color="#00C853" />
              <stop offset="1" stop-color="#8DC63F" />
            </linearGradient>
          </defs>
        </svg>
        <span class="logo-text">RUJUKAN<span>MEDIS</span></span>
      </div>

      <nav class="sidebar-nav">
        <a 
          href="/dashboard" 
          class="nav-item" 
          class:active={currentPath === '/dashboard'}
          onclick={() => isSidebarOpen = false}
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="7" height="9"></rect>
            <rect x="14" y="3" width="7" height="5"></rect>
            <rect x="14" y="12" width="7" height="9"></rect>
            <rect x="3" y="16" width="7" height="5"></rect>
          </svg>
          <span>Dashboard</span>
        </a>

        <!-- Rujukan Parent Menu Dropdown -->
        <div class="nav-dropdown-wrapper">
          <button 
            type="button"
            class="nav-item nav-dropdown-trigger" 
            class:active={currentPath.startsWith('/dashboard/rujukan-')}
            onclick={() => isRujukanDropdownOpen = !isRujukanDropdownOpen}
            aria-expanded={isRujukanDropdownOpen}
          >
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
              <polyline points="14 2 14 8 20 8"></polyline>
              <line x1="16" y1="13" x2="8" y2="13"></line>
              <line x1="16" y1="17" x2="8" y2="17"></line>
              <polyline points="10 9 9 9 8 9"></polyline>
            </svg>
            <span>Rujukan</span>
            <svg 
              class="chevron-icon" 
              class:rotate={isRujukanDropdownOpen}
              viewBox="0 0 24 24" 
              fill="none" 
              stroke="currentColor" 
              stroke-width="2.5"
            >
              <polyline points="6 9 12 15 18 9"></polyline>
            </svg>
          </button>

          {#if isRujukanDropdownOpen}
            <div class="submenu-list animate-slide-down">
              <a 
                href="/dashboard/rujukan-masuk" 
                class="submenu-item" 
                class:active={currentPath === '/dashboard/rujukan-masuk'}
                onclick={() => isSidebarOpen = false}
              >
                <span class="bullet"></span>
                <span>Rujukan Masuk</span>
              </a>

              <a 
                href="/dashboard/rujukan-keluar" 
                class="submenu-item" 
                class:active={currentPath === '/dashboard/rujukan-keluar'}
                onclick={() => isSidebarOpen = false}
              >
                <span class="bullet"></span>
                <span>Rujukan Keluar</span>
              </a>
            </div>
          {/if}
        </div>

        <a 
          href="/dashboard/pengaturan" 
          class="nav-item" 
          class:active={currentPath === '/dashboard/pengaturan'}
          onclick={() => isSidebarOpen = false}
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="3"></circle>
            <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
          </svg>
          <span>Pengaturan</span>
        </a>
      </nav>

      <div class="sidebar-footer">
        <button onclick={handleLogout} class="btn btn-secondary w-full logout-btn">
          <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
            <polyline points="16 17 21 12 16 7"></polyline>
            <line x1="21" y1="12" x2="9" y2="12"></line>
          </svg>
          <span>Keluar</span>
        </button>
      </div>
    </aside>

    <!-- Overlay when sidebar is open on mobile -->
    {#if isSidebarOpen}
      <button class="sidebar-overlay" onclick={toggleSidebar} aria-label="Close sidebar"></button>
    {/if}

    <!-- Main Container -->
    <div class="main-container">
      <!-- Topbar -->
      <header class="topbar glass-panel">
        <div class="topbar-left">
          <button class="hamburger-btn" onclick={toggleSidebar} aria-label="Toggle Navigation Menu">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="3" y1="12" x2="21" y2="12"></line>
              <line x1="3" y1="6" x2="21" y2="6"></line>
              <line x1="3" y1="18" x2="21" y2="18"></line>
            </svg>
          </button>
          <div class="topbar-welcome">
            <span class="welcome-label">Instansi Anda</span>
            <h3 class="faskes-name">{faskes}</h3>
          </div>
        </div>

        <div class="topbar-right">
          <div class="user-profile">
            <div class="avatar">
              <span>{username.split(' ').map(n => n[0]).join('').substring(0,2).toUpperCase()}</span>
            </div>
            <div class="user-info">
              <span class="user-name">{username}</span>
              <span class="user-role">{role}</span>
            </div>
          </div>
        </div>
      </header>

      <!-- Page Content -->
      <main class="page-content">
        {@render children()}
      </main>

      <!-- Footer -->
      <footer class="footer">
        <p>&copy; 2026 Rujukan Medis Terintegrasi. All Rights Reserved. Built with Svelte 5 and Premium Glassmorphism.</p>
      </footer>
    </div>
  </div>
{/if}

<style>
  .dashboard-wrapper {
    display: flex;
    min-height: 100vh;
    background: var(--bg-primary);
  }

  /* Sidebar Styling */
  .sidebar {
    width: 280px;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
    border-radius: 0;
    border: none;
    background: #005C28 !important;
    box-shadow: 2px 0 16px rgba(0, 0, 0, 0.12) !important;
    display: flex;
    flex-direction: column;
    padding: 2rem 1.5rem;
    transition: transform var(--transition-normal);
  }

  .sidebar-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    margin-bottom: 3rem;
  }

  .sidebar-logo {
    width: 32px;
    height: 32px;
    filter: drop-shadow(0 0 6px rgba(0, 200, 83, 0.4));
  }

  .logo-text {
    font-size: 1.25rem;
    font-weight: 800;
    letter-spacing: 0.05em;
    color: #ffffff;
    -webkit-text-fill-color: #ffffff;
  }

  .logo-text span {
    color: #8DC63F;
    -webkit-text-fill-color: #8DC63F;
    background: none;
  }

  .logo-text span {
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-primary-hover) 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
    flex-grow: 1;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 0.85rem 1.25rem;
    border-radius: var(--border-radius-sm);
    color: rgba(255, 255, 255, 0.7);
    font-family: var(--font-title);
    font-weight: 500;
    transition: all var(--transition-fast);
    border: 1px solid transparent;
  }

  .nav-item svg {
    width: 20px;
    height: 20px;
    transition: color var(--transition-fast);
  }

  .nav-item:hover {
    color: #ffffff;
    background: rgba(255, 255, 255, 0.1);
    border-color: rgba(255, 255, 255, 0.1);
    transform: translateX(4px);
  }

  .nav-item.active {
    color: #005C28;
    background: #8DC63F;
    border-color: #8DC63F;
    box-shadow: 0 4px 12px rgba(141, 198, 63, 0.3);
    font-weight: 700;
  }

  .nav-item.active:hover {
    transform: none;
  }

  /* Dropdown Wrapper & Trigger */
  .nav-dropdown-wrapper {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .nav-dropdown-trigger {
    background: none;
    border: 1px solid transparent;
    cursor: pointer;
    width: 100%;
    text-align: left;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .nav-dropdown-trigger span {
    flex-grow: 1;
  }

  .chevron-icon {
    width: 14px;
    height: 14px;
    margin-left: auto;
    transition: transform var(--transition-fast);
    color: var(--text-muted);
  }

  .nav-dropdown-trigger:hover .chevron-icon,
  .nav-dropdown-trigger.active .chevron-icon {
    color: inherit;
  }

  .chevron-icon.rotate {
    transform: rotate(180deg);
  }

  /* Submenu List */
  .submenu-list {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    padding-left: 1rem;
    margin-top: 0.25rem;
    border-left: 1px solid rgba(255, 255, 255, 0.15);
    margin-left: 1.5rem;
  }

  .submenu-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.65rem 1rem;
    font-size: 0.85rem;
    color: rgba(255, 255, 255, 0.65);
    border-radius: var(--border-radius-sm);
    transition: all var(--transition-fast);
  }

  .submenu-item:hover {
    color: white;
    background: rgba(255, 255, 255, 0.08);
  }

  .submenu-item.active {
    color: #8DC63F;
    font-weight: 600;
  }

  .submenu-item .bullet {
    width: 4px;
    height: 4px;
    background: rgba(255, 255, 255, 0.4);
    border-radius: 50%;
    transition: all var(--transition-fast);
  }

  .submenu-item:hover .bullet {
    background: white;
  }

  .submenu-item.active .bullet {
    background: #8DC63F;
    box-shadow: 0 0 6px rgba(141, 198, 63, 0.6);
  }

  @keyframes slideDown {
    from { opacity: 0; transform: translateY(-8px); }
    to { opacity: 1; transform: translateY(0); }
  }

  .animate-slide-down {
    animation: slideDown 0.2s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }

  .sidebar-footer {
    margin-top: auto;
  }

  .logout-btn {
    border-color: rgba(255, 255, 255, 0.2);
    color: rgba(255, 255, 255, 0.7);
    background: rgba(255, 255, 255, 0.05);
  }

  .logout-btn:hover {
    background: rgba(239, 68, 68, 0.15);
    color: #FCA5A5;
    border-color: rgba(239, 68, 68, 0.3);
  }

  /* Overlay */
  .sidebar-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(7, 26, 13, 0.85);
    backdrop-filter: blur(4px);
    z-index: 90;
    border: none;
    cursor: pointer;
  }

  /* Main Container */
  .main-container {
    flex-grow: 1;
    margin-left: 280px;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    min-width: 0; /* Prevents flex items from overflowing */
  }

  /* Topbar Styling */
  .topbar {
    height: 72px;
    border-radius: 0;
    border-top: none;
    border-left: none;
    border-right: none;
    border-bottom: 1px solid rgba(0, 0, 0, 0.08) !important;
    background: #FFFFFF !important;
    box-shadow: 0 1px 8px rgba(0, 0, 0, 0.06) !important;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 2.5rem;
    position: sticky;
    top: 0;
    z-index: 80;
  }

  .topbar-left {
    display: flex;
    align-items: center;
    gap: 1.5rem;
  }

  .hamburger-btn {
    display: none;
    background: none;
    border: none;
    color: var(--text-primary);
    cursor: pointer;
    padding: 0.25rem;
  }

  .hamburger-btn svg {
    width: 24px;
    height: 24px;
  }

  .welcome-label {
    font-size: 0.75rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    color: var(--text-muted);
  }

  .faskes-name {
    font-size: 1.1rem;
    color: var(--text-primary);
  }

  .topbar-right {
    display: flex;
    align-items: center;
    gap: 1.5rem;
  }

  .user-profile {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .avatar {
    width: 42px;
    height: 42px;
    border-radius: var(--border-radius-full);
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-secondary) 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: var(--font-title);
    font-weight: 700;
    color: #071a0d;
    box-shadow: 0 0 10px rgba(0, 200, 83, 0.25);
  }

  .user-info {
    display: flex;
    flex-direction: column;
  }

  .user-name {
    font-size: 0.95rem;
    font-weight: 600;
    color: var(--text-primary);
  }

  .user-role {
    font-size: 0.75rem;
    color: var(--text-secondary);
  }

  /* Page Content */
  .page-content {
    flex-grow: 1;
    padding: 2.5rem;
    max-width: 1600px;
    width: 100%;
    margin: 0 auto;
  }

  /* Footer */
  .footer {
    padding: 1.5rem 2.5rem;
    border-top: 1px solid var(--border-glass);
    background: #FFFFFF;
    text-align: center;
    color: var(--text-muted);
    font-size: 0.85rem;
  }

  /* Responsive Design */
  @media (max-width: 1024px) {
    .sidebar {
      transform: translateX(-100%);
    }

    .sidebar.active {
      transform: translateX(0);
    }

    .main-container {
      margin-left: 0;
    }

    .hamburger-btn {
      display: block;
    }

    .topbar {
      padding: 0 1.5rem;
    }

    .page-content {
      padding: 1.5rem;
    }

    .footer {
      padding: 1.5rem;
    }
  }
</style>
