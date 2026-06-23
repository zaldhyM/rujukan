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
          class="collapse-btn"
          onclick={toggleCollapse}
          title={isCollapsed ? 'Perluas sidebar' : 'Kecilkan sidebar'}
        >
          <svg viewBox="0 0 24 24" fill="none" width="16" height="16"
            stroke="currentColor" stroke-width="2.5"
            stroke-linecap="round" stroke-linejoin="round"
            class:flip={isCollapsed}>
            <polyline points="15 18 9 12 15 6"></polyline>
          </svg>
          {#if !isCollapsed}<span>Kecilkan Menu</span>{/if}
        </button>

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
          <div class="topbar-welcome">
            <span class="welcome-label">Instansi Anda</span>
            <h3 class="faskes-name">{faskes}</h3>
          </div>
        </div>
        <div class="topbar-right">
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
    background: #0D7A72;
    box-shadow: 2px 0 16px rgba(0,0,0,0.12);
    display: flex;
    flex-direction: column;
    padding: 2rem 1.25rem 1.5rem;
    transition: width var(--transition-normal), padding var(--transition-normal);
    overflow: hidden;
  }

  .sidebar.collapsed {
    width: 72px;
    padding: 2rem 0.75rem 1.5rem;
  }

  /* ─── Header ─── */
  .sidebar-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    margin-bottom: 1rem;
    min-height: 36px;
    overflow: hidden;
  }

  .sidebar-logo {
    width: 36px;
    height: 36px;
    flex-shrink: 0;
    object-fit: contain;
    filter: drop-shadow(0 0 6px rgba(46,196,177,0.4));
  }

  .logo-text {
    font-size: 1.1rem;
    font-weight: 800;
    letter-spacing: 0.05em;
    color: #fff;
    white-space: nowrap;
  }

  .logo-text span { color: #AACC00; }

  /* ─── Collapse Button ─── */
  .collapse-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.75rem;
    width: 100%;
    padding: 0.75rem 1rem;
    background: rgba(255,255,255,0.06);
    border: 1px solid rgba(255,255,255,0.12);
    border-radius: var(--border-radius-sm);
    color: rgba(255,255,255,0.6);
    cursor: pointer;
    font-family: var(--font-title);
    font-weight: 600;
    font-size: 0.82rem;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    white-space: nowrap;
    overflow: hidden;
    transition: all var(--transition-fast);
  }

  .sidebar.collapsed .collapse-btn {
    padding: 0.75rem;
    gap: 0;
  }

  .collapse-btn:hover {
    background: rgba(255,255,255,0.12);
    color: #fff;
    border-color: rgba(255,255,255,0.2);
  }

  .collapse-btn svg {
    transition: transform var(--transition-normal);
    flex-shrink: 0;
  }

  .collapse-btn svg.flip { transform: rotate(180deg); }

  /* ─── Nav Items ─── */
  .sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    flex-grow: 1;
    overflow: hidden;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 0.85rem 1rem;
    border-radius: var(--border-radius-sm);
    color: rgba(255,255,255,0.7);
    font-family: var(--font-title);
    font-weight: 500;
    font-size: 0.82rem;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    transition: all var(--transition-fast);
    border: 1px solid transparent;
    white-space: nowrap;
    cursor: pointer;
  }

  .sidebar.collapsed .nav-item {
    justify-content: center;
    padding: 0.85rem;
    gap: 0;
  }

  .nav-item svg {
    width: 20px;
    height: 20px;
    flex-shrink: 0;
  }

  .nav-item:hover {
    color: #fff;
    background: rgba(255,255,255,0.12);
    border-color: rgba(255,255,255,0.12);
  }

  .sidebar:not(.collapsed) .nav-item:hover {
    transform: translateX(4px);
  }

  .nav-item.active {
    color: #0D7A72;
    background: #AACC00;
    border-color: #AACC00;
    box-shadow: 0 4px 12px rgba(170,204,0,0.3);
    font-weight: 700;
  }

  .nav-item.active:hover { transform: none; }

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
    color: rgba(255,255,255,0.5);
  }

  .nav-dropdown-trigger:hover .chevron-icon,
  .nav-dropdown-trigger.active .chevron-icon { color: inherit; }

  .chevron-icon.rotate { transform: rotate(180deg); }

  /* ─── Submenu ─── */
  .submenu-list {
    display: flex;
    flex-direction: column;
    gap: 0.15rem;
    padding-left: 1rem;
    margin-top: 0.25rem;
    border-left: 1px solid rgba(255,255,255,0.15);
    margin-left: 1.5rem;
  }

  .submenu-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.6rem 1rem;
    font-size: 0.75rem;
    font-family: var(--font-title);
    font-weight: 500;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: rgba(255,255,255,0.65);
    border-radius: var(--border-radius-sm);
    transition: all var(--transition-fast);
    white-space: nowrap;
  }

  .submenu-item:hover { color: #fff; background: rgba(255,255,255,0.08); }

  .submenu-item.active { color: #AACC00; font-weight: 600; }

  .submenu-item .bullet {
    width: 4px;
    height: 4px;
    flex-shrink: 0;
    background: rgba(255,255,255,0.4);
    border-radius: 50%;
    transition: all var(--transition-fast);
  }

  .submenu-item:hover .bullet { background: #fff; }

  .submenu-item.active .bullet {
    background: #AACC00;
    box-shadow: 0 0 6px rgba(170,204,0,0.6);
  }

  @keyframes slideDown {
    from { opacity: 0; transform: translateY(-8px); }
    to   { opacity: 1; transform: translateY(0); }
  }

  .animate-slide-down {
    animation: slideDown 0.2s cubic-bezier(0.16,1,0.3,1) forwards;
  }

  /* ─── Footer / Logout ─── */
  .sidebar-footer {
    margin-top: auto;
    padding-top: 1rem;
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
    padding: 0.75rem 1rem;
    border-radius: var(--border-radius-sm);
    border: 1px solid rgba(255,255,255,0.2);
    color: rgba(255,255,255,0.7);
    background: rgba(255,255,255,0.05);
    cursor: pointer;
    font-family: var(--font-title);
    font-weight: 600;
    font-size: 0.82rem;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    transition: all var(--transition-fast);
    white-space: nowrap;
    overflow: hidden;
  }

  .sidebar.collapsed .logout-btn {
    padding: 0.75rem;
    gap: 0;
  }

  .logout-btn:hover {
    background: rgba(239,68,68,0.15);
    color: #FCA5A5;
    border-color: rgba(239,68,68,0.3);
  }

  /* ─── Tooltip / Flyout ─── */
  .nav-tooltip {
    position: fixed;
    transform: translateY(-50%);
    z-index: 200;
    background: #0D4F4B;
    border: 1px solid rgba(255,255,255,0.12);
    border-radius: 8px;
    padding: 0.5rem 0.85rem;
    pointer-events: none;
    animation: tooltipIn 0.15s ease forwards;
    min-width: max-content;
    box-shadow: 0 6px 20px rgba(0,0,0,0.2);
  }

  .nav-tooltip.is-flyout {
    padding: 0.75rem 0;
    pointer-events: auto;
  }

  @keyframes tooltipIn {
    from { opacity: 0; transform: translateY(-50%) translateX(-6px); }
    to   { opacity: 1; transform: translateY(-50%) translateX(0); }
  }

  /* Arrow kiri */
  .nav-tooltip::before {
    content: '';
    position: absolute;
    left: -6px;
    top: 50%;
    transform: translateY(-50%);
    border-width: 5px 6px 5px 0;
    border-style: solid;
    border-color: transparent #0D4F4B transparent transparent;
  }

  .tooltip-label {
    display: block;
    font-family: var(--font-title);
    font-size: 0.78rem;
    font-weight: 700;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: #fff;
    padding: 0 0.15rem;
  }

  .is-flyout .tooltip-label {
    padding: 0 1rem 0.5rem;
    color: rgba(255,255,255,0.6);
    font-size: 0.7rem;
  }

  .tooltip-divider {
    height: 1px;
    background: rgba(255,255,255,0.1);
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
    letter-spacing: 0.05em;
    text-transform: uppercase;
    color: rgba(255,255,255,0.8);
    transition: all var(--transition-fast);
    white-space: nowrap;
  }

  .tooltip-link:hover {
    background: rgba(255,255,255,0.1);
    color: #fff;
  }

  .tooltip-link.active {
    color: #AACC00;
  }

  /* ─── Overlay mobile ─── */
  .sidebar-overlay {
    position: fixed;
    inset: 0;
    background: rgba(13,122,114,0.6);
    backdrop-filter: blur(4px);
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
    height: 72px;
    background: #fff;
    border-bottom: 1px solid rgba(0,0,0,0.08);
    box-shadow: 0 1px 8px rgba(0,0,0,0.06);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 2.5rem;
    position: sticky;
    top: 0;
    z-index: 80;
  }

  .topbar-left { display: flex; align-items: center; gap: 1.5rem; }

  .hamburger-btn {
    display: none;
    background: none;
    border: none;
    color: var(--text-primary);
    cursor: pointer;
    padding: 0.25rem;
  }

  .hamburger-btn svg { width: 24px; height: 24px; }

  .welcome-label {
    font-size: 0.75rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    color: var(--text-muted);
  }

  .faskes-name { font-size: 1.1rem; color: var(--text-primary); }

  .topbar-right { display: flex; align-items: center; gap: 1.5rem; }

  .user-profile { display: flex; align-items: center; gap: 1rem; }

  .avatar {
    width: 42px;
    height: 42px;
    border-radius: var(--border-radius-full);
    background: linear-gradient(135deg, #2EC4B1 0%, #AACC00 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: var(--font-title);
    font-weight: 700;
    color: #fff;
    box-shadow: 0 0 10px rgba(46,196,177,0.25);
    flex-shrink: 0;
  }

  .user-info { display: flex; flex-direction: column; }

  .user-name { font-size: 0.95rem; font-weight: 600; color: var(--text-primary); }
  .user-role { font-size: 0.75rem; color: var(--text-secondary); }

  /* ─── Content ─── */
  .page-content {
    flex-grow: 1;
    padding: 2.5rem;
    max-width: 1600px;
    width: 100%;
    margin: 0 auto;
  }

  .footer {
    padding: 1.5rem 2.5rem;
    border-top: 1px solid var(--border-glass);
    background: #fff;
    text-align: center;
    color: var(--text-muted);
    font-size: 0.85rem;
  }

  /* ─── Responsive ─── */
  @media (max-width: 1024px) {
    .sidebar {
      transform: translateX(-100%);
      width: 280px !important;
      padding: 2rem 1.25rem 1.5rem !important;
    }

    .sidebar.active { transform: translateX(0); }

    .collapse-btn { display: none; }

    .main-container,
    .main-container.collapsed { margin-left: 0; }

    .hamburger-btn { display: block; }

    .topbar { padding: 0 1.5rem; }
    .page-content { padding: 1.5rem; }
    .footer { padding: 1.5rem; }

    .nav-tooltip { display: none; }
  }
</style>
