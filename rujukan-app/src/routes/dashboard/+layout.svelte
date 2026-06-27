<!-- src/routes/dashboard/+layout.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/state';
  import kemkesLogo from '$lib/assets/kemkes.png';

  let { children } = $props();

  let isSidebarOpen  = $state(false);
  let isCollapsed    = $state(false);
  let username       = $state('');
  let faskes         = $state('');
  let role           = $state('');
  let isLoaded       = $state(false);

  let currentPath = $derived(page.url.pathname);
  let isRujukanDropdownOpen = $state(false);

  // Tooltip state
  type TooltipInfo = { label: string; y: number; isRujukan?: boolean };
  let tooltip = $state<TooltipInfo | null>(null);
  let SIDEBAR_EXPANDED = 280;
  let SIDEBAR_COLLAPSED = 72;

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
      username  = localStorage.getItem('username')  || 'Tenaga Medis';
      faskes    = localStorage.getItem('faskes')    || 'Fasilitas Kesehatan';
      role      = localStorage.getItem('role')      || 'Staf Medis';
      isCollapsed = localStorage.getItem('sidebarCollapsed') === 'true';
      isLoaded  = true;
    }
  });

  const handleLogout = () => {
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('username');
    localStorage.removeItem('faskes');
    localStorage.removeItem('role');
    goto('/login');
  };

  const toggleSidebar = () => { isSidebarOpen = !isSidebarOpen; };

  const toggleCollapse = () => {
    isCollapsed = !isCollapsed;
    if (isCollapsed) isRujukanDropdownOpen = false;
    tooltip = null;
    localStorage.setItem('sidebarCollapsed', String(isCollapsed));
  };

  const handleRujukanClick = () => {
    if (isCollapsed) {
      isCollapsed = false;
      localStorage.setItem('sidebarCollapsed', 'false');
      isRujukanDropdownOpen = true;
      tooltip = null;
    } else {
      isRujukanDropdownOpen = !isRujukanDropdownOpen;
    }
  };

  // Show tooltip anchored to the hovered element
  function showTooltip(e: MouseEvent, label: string, isRujukan = false) {
    if (!isCollapsed) return;
    const el = e.currentTarget as HTMLElement;
    const rect = el.getBoundingClientRect();
    tooltip = { label, y: rect.top + rect.height / 2, isRujukan };
  }

  function hideTooltip() { tooltip = null; }

  let tooltipLeft = $derived(isCollapsed ? SIDEBAR_COLLAPSED + 10 : SIDEBAR_EXPANDED + 10);
</script>

<svelte:head>
  <title>Dashboard | RUJUKAN MEDIS</title>
</svelte:head>

{#if isLoaded}
  <div class="dashboard-wrapper">

    <!-- ── Sidebar ── -->
    <aside class="sidebar" class:active={isSidebarOpen} class:collapsed={isCollapsed}>

      <div class="sidebar-header">
        <img src={kemkesLogo} alt="Logo Kemkes" class="sidebar-logo" />
        {#if !isCollapsed}
          <span class="logo-text">RUJUKAN<span>MEDIS</span></span>
        {/if}
      </div>

      <nav class="sidebar-nav">

        <!-- Dashboard -->
        <a
          href="/dashboard"
          class="nav-item"
          class:active={currentPath === '/dashboard'}
          onclick={() => isSidebarOpen = false}
          onmouseenter={(e) => showTooltip(e, 'Dashboard')}
          onmouseleave={hideTooltip}
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="7" height="9"></rect>
            <rect x="14" y="3" width="7" height="5"></rect>
            <rect x="14" y="12" width="7" height="9"></rect>
            <rect x="3" y="16" width="7" height="5"></rect>
          </svg>
          {#if !isCollapsed}<span>Dashboard</span>{/if}
        </a>

        <!-- Rujukan dropdown -->
        <div class="nav-dropdown-wrapper">
          <button
            type="button"
            class="nav-item nav-dropdown-trigger"
            class:active={currentPath.startsWith('/dashboard/rujukan-')}
            onclick={handleRujukanClick}
            aria-expanded={isRujukanDropdownOpen}
            onmouseenter={(e) => showTooltip(e, 'Rujukan', true)}
            onmouseleave={hideTooltip}
          >
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
              <polyline points="14 2 14 8 20 8"></polyline>
              <line x1="16" y1="13" x2="8" y2="13"></line>
              <line x1="16" y1="17" x2="8" y2="17"></line>
              <polyline points="10 9 9 9 8 9"></polyline>
            </svg>
            {#if !isCollapsed}
              <span>Rujukan</span>
              <svg class="chevron-icon" class:rotate={isRujukanDropdownOpen}
                viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            {/if}
          </button>

          {#if isRujukanDropdownOpen && !isCollapsed}
            <div class="submenu-list animate-slide-down">
              <a
                href="/dashboard/rujukan-masuk"
                class="submenu-item"
                class:active={currentPath === '/dashboard/rujukan-masuk'}
                onclick={() => isSidebarOpen = false}
                onmouseenter={(e) => showTooltip(e, 'Rujukan Masuk')}
                onmouseleave={hideTooltip}
              >
                <span class="bullet"></span>
                <span>Rujukan Masuk</span>
              </a>
              <a
                href="/dashboard/rujukan-keluar"
                class="submenu-item"
                class:active={currentPath === '/dashboard/rujukan-keluar'}
                onclick={() => isSidebarOpen = false}
                onmouseenter={(e) => showTooltip(e, 'Rujukan Keluar')}
                onmouseleave={hideTooltip}
              >
                <span class="bullet"></span>
                <span>Rujukan Keluar</span>
              </a>
            </div>
          {/if}
        </div>

        <!-- Pengaturan -->
        <a
          href="/dashboard/pengaturan"
          class="nav-item"
          class:active={currentPath === '/dashboard/pengaturan'}
          onclick={() => isSidebarOpen = false}
          onmouseenter={(e) => showTooltip(e, 'Pengaturan')}
          onmouseleave={hideTooltip}
        >
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="3"></circle>
            <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
          </svg>
          {#if !isCollapsed}<span>Pengaturan</span>{/if}
        </a>
      </nav>

      <div class="sidebar-footer">
        <button
          onclick={handleLogout}
          class="logout-btn"
          onmouseenter={(e) => showTooltip(e, 'Keluar')}
          onmouseleave={hideTooltip}
        >
          <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
            <polyline points="16 17 21 12 16 7"></polyline>
            <line x1="21" y1="12" x2="9" y2="12"></line>
          </svg>
          {#if !isCollapsed}<span>Keluar</span>{/if}
        </button>
      </div>
    </aside>

    <!-- ── Tooltip / Flyout Panel ── -->
    {#if tooltip}
      <div
        class="nav-tooltip"
        class:is-flyout={tooltip.isRujukan && isCollapsed}
        style="top: {tooltip.y}px; left: {tooltipLeft}px;"
        onmouseenter={() => { tooltip = tooltip; }}
        onmouseleave={hideTooltip}
      >
        <span class="tooltip-label">{tooltip.label}</span>

        {#if tooltip.isRujukan && isCollapsed}
          <div class="tooltip-divider"></div>
          <a
            href="/dashboard/rujukan-masuk"
            class="tooltip-link"
            class:active={currentPath === '/dashboard/rujukan-masuk'}
            onmouseenter={() => { tooltip = tooltip; }}
          >
            <svg viewBox="0 0 24 24" fill="none" width="14" height="14" stroke="currentColor" stroke-width="2">
              <polyline points="7 10 12 15 17 10"></polyline>
              <line x1="12" y1="3" x2="12" y2="15"></line>
              <line x1="3" y1="21" x2="21" y2="21"></line>
            </svg>
            Rujukan Masuk
          </a>
          <a
            href="/dashboard/rujukan-keluar"
            class="tooltip-link"
            class:active={currentPath === '/dashboard/rujukan-keluar'}
            onmouseenter={() => { tooltip = tooltip; }}
          >
            <svg viewBox="0 0 24 24" fill="none" width="14" height="14" stroke="currentColor" stroke-width="2">
              <polyline points="17 8 12 3 7 8"></polyline>
              <line x1="12" y1="3" x2="12" y2="15"></line>
              <line x1="3" y1="21" x2="21" y2="21"></line>
            </svg>
            Rujukan Keluar
          </a>
        {/if}
      </div>
    {/if}

    <!-- Overlay mobile -->
    {#if isSidebarOpen}
      <button class="sidebar-overlay" onclick={toggleSidebar} aria-label="Close sidebar"></button>
    {/if}

    <!-- ── Main Container ── -->
    <div class="main-container" class:collapsed={isCollapsed}>
      <header class="topbar">
        <div class="topbar-left">
          <button class="hamburger-btn" onclick={toggleSidebar} aria-label="Toggle Navigation Menu">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="3" y1="12" x2="21" y2="12"></line>
              <line x1="3" y1="6" x2="21" y2="6"></line>
              <line x1="3" y1="18" x2="21" y2="18"></line>
            </svg>
          </button>
          <button
            class="collapse-toggle-btn"
            onclick={toggleCollapse}
            title={isCollapsed ? 'Perluas sidebar' : 'Kecilkan sidebar'}
          >
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
              stroke-linecap="round" stroke-linejoin="round">
              <rect x="3" y="3" width="18" height="18" rx="2"></rect>
              <line x1="9" y1="3" x2="9" y2="21"></line>
            </svg>
          </button>
          <div class="topbar-search">
            <svg class="search-ic" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8"></circle>
              <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
            </svg>
            <input type="text" class="search-field" placeholder="Cari pasien, menu, laporan..." />
          </div>
        </div>
        <div class="topbar-right">
          <div class="faskes-chip">
            <svg viewBox="0 0 24 24" fill="none" width="14" height="14" stroke="currentColor" stroke-width="2">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
              <polyline points="9 22 9 12 15 12 15 22"></polyline>
            </svg>
            <span>{faskes}</span>
          </div>
          <div class="topbar-divider"></div>
          <div class="user-profile">
            <div class="avatar">
              <span>{username.split(' ').map((n: string) => n[0]).join('').substring(0,2).toUpperCase()}</span>
            </div>
            <div class="user-info">
              <span class="user-name">{username}</span>
              <span class="user-role">{role}</span>
            </div>
          </div>
        </div>
      </header>

      <main class="page-content">
        {@render children()}
      </main>

      <footer class="footer">
        <p>&copy; 2026 Rujukan Medis Terintegrasi. All Rights Reserved.</p>
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

  /* ─── Sidebar ─── */
  .sidebar {
    width: 280px;
    height: 100vh;
    position: fixed;
    top: 0; left: 0;
    z-index: 100;
    background: #FFFFFF;
    border-right: 1px solid var(--border-glass);
    box-shadow: none;
    display: flex;
    flex-direction: column;
    padding: 1.75rem 1rem 1.25rem;
    transition: width var(--transition-normal), padding var(--transition-normal);
    overflow: hidden;
  }

  .sidebar.collapsed {
    width: 72px;
    padding: 1.75rem 0.75rem 1.25rem;
  }

  /* ─── Header ─── */
  .sidebar-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    margin-bottom: 1.5rem;
    padding-bottom: 1.25rem;
    border-bottom: 1px solid var(--border-glass);
    min-height: 36px;
    overflow: hidden;
  }

  .sidebar-logo {
    width: 32px;
    height: 32px;
    flex-shrink: 0;
    object-fit: contain;
  }

  .logo-text {
    font-size: 1rem;
    font-weight: 800;
    letter-spacing: 0.04em;
    color: var(--text-primary);
    white-space: nowrap;
  }

  .logo-text span { color: var(--color-primary); }

  /* ─── Collapse Toggle (Topbar) ─── */
  .collapse-toggle-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    color: var(--text-muted);
    cursor: pointer;
    padding: 0.25rem;
    flex-shrink: 0;
    border-radius: 6px;
    transition: color var(--transition-fast), background var(--transition-fast);
  }

  .collapse-toggle-btn:hover { color: var(--text-primary); background: var(--bg-primary); }

  .collapse-toggle-btn svg { width: 22px; height: 22px; }

  /* ─── Nav Items ─── */
  .sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    flex-grow: 1;
    overflow: hidden;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 0.875rem;
    padding: 0.75rem 1rem;
    border-radius: var(--border-radius-sm);
    color: var(--text-muted);
    font-family: var(--font-title);
    font-weight: 500;
    font-size: 0.875rem;
    letter-spacing: 0.01em;
    transition: all var(--transition-fast);
    border: 1px solid transparent;
    white-space: nowrap;
    cursor: pointer;
  }

  .sidebar.collapsed .nav-item {
    justify-content: center;
    padding: 0.75rem;
    gap: 0;
  }

  .nav-item svg {
    width: 18px;
    height: 18px;
    flex-shrink: 0;
  }

  .nav-item:hover {
    color: var(--color-primary);
    background: rgba(20, 184, 166, 0.08);
    border-color: transparent;
  }

  .nav-item.active {
    color: #FFFFFF;
    background: var(--color-primary);
    border-color: var(--color-primary);
    font-weight: 600;
    box-shadow: 0 2px 8px rgba(20, 184, 166, 0.25);
  }

  .nav-item.active:hover { background: var(--color-primary); color: #FFFFFF; }

  /* ─── Dropdown ─── */
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
  }

  .nav-dropdown-trigger span:first-of-type { flex-grow: 1; }

  .chevron-icon {
    width: 14px;
    height: 14px;
    flex-shrink: 0;
    margin-left: auto;
    transition: transform var(--transition-fast);
    color: #9CA3AF;
  }

  .nav-dropdown-trigger:hover .chevron-icon { color: var(--color-primary); }
  .nav-dropdown-trigger.active .chevron-icon { color: #FFFFFF; }

  .chevron-icon.rotate { transform: rotate(180deg); }

  /* ─── Submenu ─── */
  .submenu-list {
    display: flex;
    flex-direction: column;
    gap: 0.1rem;
    padding-left: 0.875rem;
    margin-top: 0.15rem;
    border-left: 2px solid var(--border-glass);
    margin-left: 1.4rem;
  }

  .submenu-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.55rem 0.875rem;
    font-size: 0.82rem;
    font-family: var(--font-title);
    font-weight: 500;
    letter-spacing: 0.01em;
    color: var(--text-muted);
    border-radius: var(--border-radius-sm);
    transition: all var(--transition-fast);
    white-space: nowrap;
  }

  .submenu-item:hover { color: var(--color-primary); background: rgba(20, 184, 166, 0.06); }

  .submenu-item.active { color: var(--color-primary); font-weight: 600; }

  .submenu-item .bullet {
    width: 5px;
    height: 5px;
    flex-shrink: 0;
    background: #D1D5DB;
    border-radius: 50%;
    transition: all var(--transition-fast);
  }

  .submenu-item:hover .bullet { background: var(--color-primary); }
  .submenu-item.active .bullet { background: var(--color-primary); }

  @keyframes slideDown {
    from { opacity: 0; transform: translateY(-6px); }
    to   { opacity: 1; transform: translateY(0); }
  }

  .animate-slide-down {
    animation: slideDown 0.2s cubic-bezier(0.16,1,0.3,1) forwards;
  }

  /* ─── Footer / Logout ─── */
  .sidebar-footer {
    margin-top: auto;
    padding-top: 1rem;
    border-top: 1px solid var(--border-glass);
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .logout-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.75rem;
    width: 100%;
    padding: 0.625rem 1rem;
    border-radius: var(--border-radius-sm);
    border: 1px solid var(--border-glass);
    color: var(--text-muted);
    background: var(--bg-primary);
    cursor: pointer;
    font-family: var(--font-title);
    font-weight: 600;
    font-size: 0.78rem;
    letter-spacing: 0.05em;
    text-transform: uppercase;
    transition: all var(--transition-fast);
    white-space: nowrap;
    overflow: hidden;
  }

  .sidebar.collapsed .logout-btn {
    padding: 0.625rem;
    gap: 0;
  }

  .logout-btn:hover {
    background: #FEF2F2;
    color: var(--color-danger);
    border-color: #FECACA;
  }

  /* ─── Tooltip / Flyout ─── */
  .nav-tooltip {
    position: fixed;
    transform: translateY(-50%);
    z-index: 200;
    background: #1F2937;
    border: 1px solid #374151;
    border-radius: 8px;
    padding: 0.5rem 0.85rem;
    pointer-events: none;
    animation: tooltipIn 0.15s ease forwards;
    min-width: max-content;
    box-shadow: 0 4px 12px rgba(0,0,0,0.15);
  }

  .nav-tooltip.is-flyout {
    padding: 0.75rem 0;
    pointer-events: auto;
  }

  @keyframes tooltipIn {
    from { opacity: 0; transform: translateY(-50%) translateX(-6px); }
    to   { opacity: 1; transform: translateY(-50%) translateX(0); }
  }

  .nav-tooltip::before {
    content: '';
    position: absolute;
    left: -6px;
    top: 50%;
    transform: translateY(-50%);
    border-width: 5px 6px 5px 0;
    border-style: solid;
    border-color: transparent #1F2937 transparent transparent;
  }

  .tooltip-label {
    display: block;
    font-family: var(--font-title);
    font-size: 0.78rem;
    font-weight: 700;
    letter-spacing: 0.05em;
    text-transform: uppercase;
    color: #fff;
    padding: 0 0.15rem;
  }

  .is-flyout .tooltip-label {
    padding: 0 1rem 0.5rem;
    color: rgba(255,255,255,0.5);
    font-size: 0.7rem;
  }

  .tooltip-divider {
    height: 1px;
    background: #374151;
    margin: 0 0 0.35rem;
  }

  .tooltip-link {
    display: flex;
    align-items: center;
    gap: 0.6rem;
    padding: 0.6rem 1rem;
    font-family: var(--font-title);
    font-size: 0.8rem;
    font-weight: 600;
    letter-spacing: 0.03em;
    text-transform: uppercase;
    color: rgba(255,255,255,0.8);
    transition: all var(--transition-fast);
    white-space: nowrap;
  }

  .tooltip-link:hover {
    background: rgba(255,255,255,0.08);
    color: #fff;
  }

  .tooltip-link.active { color: var(--color-primary); }

  /* ─── Overlay mobile ─── */
  .sidebar-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.35);
    backdrop-filter: blur(2px);
    z-index: 90;
    border: none;
    cursor: pointer;
  }

  /* ─── Main Container ─── */
  .main-container {
    flex-grow: 1;
    margin-left: 280px;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    min-width: 0;
    transition: margin-left var(--transition-normal);
  }

  .main-container.collapsed { margin-left: 72px; }

  /* ─── Topbar ─── */
  .topbar {
    height: 64px;
    background: #FFFFFF;
    border-bottom: 1px solid var(--border-glass);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 2rem;
    position: sticky;
    top: 0;
    z-index: 80;
    gap: 1rem;
  }

  .topbar-left {
    display: flex;
    align-items: center;
    gap: 1rem;
    flex: 1;
  }

  .hamburger-btn {
    display: none;
    background: none;
    border: none;
    color: var(--text-muted);
    cursor: pointer;
    padding: 0.25rem;
    flex-shrink: 0;
    border-radius: 6px;
    transition: color var(--transition-fast), background var(--transition-fast);
  }

  .hamburger-btn:hover { color: var(--text-primary); background: var(--bg-primary); }

  .hamburger-btn svg { width: 22px; height: 22px; }

  /* Search bar in topbar */
  .topbar-search {
    position: relative;
    flex: 1;
    max-width: 420px;
  }

  .search-ic {
    position: absolute;
    left: 0.875rem;
    top: 50%;
    transform: translateY(-50%);
    width: 16px;
    height: 16px;
    color: #9CA3AF;
    pointer-events: none;
  }

  .search-field {
    width: 100%;
    height: 38px;
    background: var(--bg-primary);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-full);
    padding: 0 1rem 0 2.5rem;
    font-family: var(--font-body);
    font-size: 0.875rem;
    color: var(--text-primary);
    transition: all var(--transition-fast);
  }

  .search-field::placeholder { color: #9CA3AF; }

  .search-field:focus {
    outline: none;
    border-color: var(--color-primary);
    background: #FFFFFF;
    box-shadow: 0 0 0 3px rgba(20, 184, 166, 0.1);
  }

  .topbar-right {
    display: flex;
    align-items: center;
    gap: 1rem;
    flex-shrink: 0;
  }

  .faskes-chip {
    display: flex;
    align-items: center;
    gap: 0.4rem;
    font-size: 0.8rem;
    font-weight: 500;
    color: var(--text-muted);
    background: var(--bg-primary);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-full);
    padding: 0.3rem 0.875rem;
    white-space: nowrap;
    max-width: 200px;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .faskes-chip span {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .topbar-divider {
    width: 1px;
    height: 28px;
    background: var(--border-glass);
    flex-shrink: 0;
  }

  .user-profile { display: flex; align-items: center; gap: 0.75rem; }

  .avatar {
    width: 38px;
    height: 38px;
    border-radius: var(--border-radius-full);
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-secondary) 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: var(--font-title);
    font-weight: 700;
    font-size: 0.85rem;
    color: #fff;
    flex-shrink: 0;
  }

  .user-info { display: flex; flex-direction: column; }

  .user-name { font-size: 0.875rem; font-weight: 600; color: var(--text-primary); }
  .user-role { font-size: 0.72rem; color: var(--text-muted); }

  /* ─── Content ─── */
  .page-content {
    flex-grow: 1;
    padding: 2rem 2.5rem;
    max-width: 1600px;
    width: 100%;
    margin: 0 auto;
  }

  .footer {
    padding: 1.25rem 2.5rem;
    border-top: 1px solid var(--border-glass);
    background: #FFFFFF;
    text-align: center;
    color: var(--text-muted);
    font-size: 0.82rem;
  }

  /* ─── Responsive ─── */
  @media (max-width: 1024px) {
    .sidebar {
      transform: translateX(-100%);
      width: 280px !important;
      padding: 1.75rem 1rem 1.25rem !important;
    }

    .sidebar.active { transform: translateX(0); }

    .collapse-toggle-btn { display: none; }

    .main-container,
    .main-container.collapsed { margin-left: 0; }

    .hamburger-btn { display: flex; }

    .topbar { padding: 0 1.25rem; }
    .page-content { padding: 1.5rem; }
    .footer { padding: 1.25rem 1.5rem; }

    .nav-tooltip { display: none; }

    .faskes-chip { display: none; }
    .topbar-divider { display: none; }
  }

  @media (max-width: 640px) {
    .topbar-search { max-width: 200px; }
  }
</style>
