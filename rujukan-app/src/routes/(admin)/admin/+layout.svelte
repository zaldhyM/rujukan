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
    faHospital
  } from "@fortawesome/free-solid-svg-icons";
  import avatar from "$lib/assets/img/undraw_profile.svg";
  import { isLoggedIn, getCurrentUser, getFaskesSettings, logout, getRujukanMasuk } from '$lib/auth';

  let { children } = $props();

  let user: any = $state({ name: '', username: '', role: '' });
  let faskes: any = $state({ name: '', code: '' });
  let isMounted = $state(false);
  let pendingCount = $state(0);

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
</script>

<svelte:head>
  <title>Dashboard Rujukan - {faskes.name || 'Faskes'}</title>
</svelte:head>

{#if isMounted}
<div id="wrapper">
  <!-- Sidebar -->
  <ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">
    <!-- Sidebar - Brand -->
    <a class="sidebar-brand d-flex align-items-center justify-content-center py-4" href="/admin">
      <div class="sidebar-brand-icon">
        <Fa icon={faHospital} size="lg" />
      </div>
      <div class="sidebar-brand-text mx-2 text-capitalize small font-weight-bold">SISRUTE LOKAL</div>
    </a>

    <!-- Divider -->
    <hr class="sidebar-divider my-0" />

    <!-- Nav Item - Dashboard -->
    <li class="nav-item" class:active={isActive('/admin')}>
      <a class="nav-link" href="/admin">
        <Fa icon={faChartLine} class="mr-2" />
        <span>Dashboard</span>
      </a>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider" />

    <!-- Heading -->
    <div class="sidebar-heading">Layanan Rujukan</div>

    <!-- Nav Item - Rujukan Masuk -->
    <li class="nav-item" class:active={isActive('/admin/rujukan-masuk')}>
      <a class="nav-link" href="/admin/rujukan-masuk">
        <Fa icon={faInbox} class="mr-2" />
        <span>Rujukan Masuk</span>
        {#if pendingCount > 0}
          <span class="badge badge-danger badge-counter ml-2">{pendingCount}</span>
        {/if}
      </a>
    </li>

    <!-- Nav Item - Rujukan Keluar -->
    <li class="nav-item" class:active={isActive('/admin/rujukan-keluar')}>
      <a class="nav-link" href="/admin/rujukan-keluar">
        <Fa icon={faPaperPlane} class="mr-2" />
        <span>Rujukan Keluar</span>
      </a>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider" />

    <!-- Heading -->
    <div class="sidebar-heading">Konfigurasi</div>

    <!-- Nav Item - Pengaturan -->
    <li class="nav-item" class:active={isActive('/admin/pengaturan')}>
      <a class="nav-link" href="/admin/pengaturan">
        <Fa icon={faGear} class="mr-2" />
        <span>Pengaturan</span>
      </a>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider d-none d-md-block" />
  </ul>
  <!-- End of Sidebar -->

  <!-- Content Wrapper -->
  <div id="content-wrapper" class="d-flex flex-column">
    <!-- Main Content -->
    <div id="content">
      <!-- Topbar -->
      <nav class="navbar navbar-expand navbar-light bg-white topbar mb-4 static-top shadow-sm">
        <!-- Sidebar Toggle (Topbar) -->
        <button id="sidebarToggleTop" class="btn btn-link d-md-none rounded-circle mr-3">
          <Fa icon={faBars} />
        </button>

        <!-- Faskes Name Display -->
        <div class="d-none d-sm-inline-block ml-md-3 my-2 my-md-0 mw-100 text-gray-800 font-weight-bold">
          <Fa icon={faHospital} class="text-primary mr-1" />
          {faskes.name} ({faskes.code})
        </div>

        <!-- Topbar Navbar -->
        <ul class="navbar-nav ml-auto">
          <!-- Nav Item - Alerts -->
          <li class="nav-item dropdown no-arrow mx-1">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              id="alertsDropdown"
              role="button"
              data-toggle="dropdown"
              aria-haspopup="true"
              aria-expanded="false"
            >
              <Fa icon={faBell} style="color:gray" />
              <!-- Counter - Alerts -->
              {#if pendingCount > 0}
                <span class="badge badge-danger badge-counter">{pendingCount}</span>
              {/if}
            </a>
            <!-- Dropdown - Alerts -->
            <div
              class="dropdown-list dropdown-menu dropdown-menu-right shadow animated--grow-in"
              aria-labelledby="alertsDropdown"
            >
              <h6 class="dropdown-header">Notifikasi Masuk</h6>
              {#if pendingCount > 0}
                <a class="dropdown-item d-flex align-items-center" href="/admin/rujukan-masuk">
                  <div class="mr-3">
                    <div class="icon-circle bg-warning text-white p-2 rounded-circle">
                      <Fa icon={faInbox} />
                    </div>
                  </div>
                  <div>
                    <div class="small text-gray-500">Hari ini</div>
                    <span class="font-weight-bold">Ada {pendingCount} rujukan baru yang perlu ditinjau!</span>
                  </div>
                </a>
              {:else}
                <div class="p-3 text-center small text-gray-500">Tidak ada notifikasi baru</div>
              {/if}
            </div>
          </li>

          <div class="topbar-divider d-none d-sm-block"></div>

          <!-- Nav Item - User Information -->
          <li class="nav-item dropdown no-arrow">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              id="userDropdown"
              role="button"
              data-toggle="dropdown"
              aria-haspopup="true"
              aria-expanded="false"
            >
              <div class="d-flex flex-column text-right mr-2">
                <span class="d-none d-lg-inline text-gray-700 font-weight-bold small">{user.name}</span>
                <span class="d-none d-lg-inline text-gray-500 text-xs">{user.role}</span>
              </div>
              <img class="img-profile rounded-circle" src={avatar} alt="Profile" style="width: 32px; height: 32px;" />
            </a>
            <!-- Dropdown - User Information -->
            <div
              class="dropdown-menu dropdown-menu-right shadow animated--grow-in"
              aria-labelledby="userDropdown"
            >
              <a class="dropdown-item" href="/admin/pengaturan">
                <Fa icon={faUser} class="mr-2 text-gray-400" />
                Profil Saya
              </a>
              <a class="dropdown-item" href="/admin/pengaturan">
                <Fa icon={faGear} class="mr-2 text-gray-400" />
                Pengaturan Faskes
              </a>
              <div class="dropdown-divider"></div>
              <button class="dropdown-item text-danger font-weight-bold btn-link bg-transparent border-0 w-100 text-left" onclick={handleLogout}>
                <Fa icon={faRightFromBracket} class="mr-2 text-danger" />
                Keluar Aplikasi
              </button>
            </div>
          </li>
        </ul>
      </nav>

      <!-- Main Content -->
      <div class="container-fluid py-2">
        {@render children?.()}
      </div>
    </div>
    <!-- Footer -->
    <footer class="sticky-footer bg-white shadow-sm mt-5">
      <div class="container my-auto">
        <div class="copyright text-center my-auto">
          <span>Copyright &copy; Aplikasi Rujukan Pasien Svelte 2026 - {faskes.name}</span>
        </div>
      </div>
    </footer>
    <!-- End of Footer -->
  </div>
  <!-- End of Content Wrapper -->
</div>
{/if}

<style>
  /* Premium Sidebar modifications */
  .sidebar {
    background: linear-gradient(180deg, #4e73df 10deg, #224abe 100%) !important;
  }
  .nav-item.active .nav-link {
    background-color: rgba(255, 255, 255, 0.15) !important;
    font-weight: bold;
  }
  .nav-link {
    display: flex;
    align-items: center;
    padding: 0.75rem 1rem !important;
  }
  .navbar-nav.sidebar .nav-item .nav-link span {
    font-size: 0.85rem !important;
  }
  .topbar-divider {
    border-right: 1px solid #e3e6f0;
    height: calc(4.375rem - 2rem);
    margin: auto 1rem;
  }
</style>
