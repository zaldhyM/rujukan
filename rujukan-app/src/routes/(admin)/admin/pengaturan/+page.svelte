<script lang="ts">
  import { onMount } from 'svelte';
  import Fa from 'svelte-fa';
  import {
    faGear,
    faUser,
    faHospital,
    faSave,
    faCheckCircle,
    faLock
  } from '@fortawesome/free-solid-svg-icons';
  import { getCurrentUser, updateCurrentUser, getFaskesSettings, updateFaskesSettings } from '$lib/auth';

  // State Profile
  let userName = $state('');
  let userRole = $state('');
  let userUsername = $state('');
  
  // State Faskes
  let faskesName = $state('');
  let faskesCode = $state('');
  let faskesType = $state('');
  let faskesAddress = $state('');
  let faskesPhone = $state('');
  let faskesEmail = $state('');

  // Alerts State
  let showSuccessUser = $state(false);
  let showSuccessFaskes = $state(false);

  onMount(() => {
    // Load current values
    const user = getCurrentUser();
    userName = user.name;
    userRole = user.role;
    userUsername = user.username;

    const faskes = getFaskesSettings();
    faskesName = faskes.name;
    faskesCode = faskes.code;
    faskesType = faskes.type;
    faskesAddress = faskes.address;
    faskesPhone = faskes.phone;
    faskesEmail = faskes.email;
  });

  function handleSaveProfile(e: SubmitEvent) {
    e.preventDefault();
    updateCurrentUser({
      name: userName,
      role: userRole,
      username: userUsername
    });
    
    showSuccessUser = true;
    setTimeout(() => {
      showSuccessUser = false;
      // Reload page to reflect name changes in Topbar
      window.location.reload();
    }, 1000);
  }

  function handleSaveFaskes(e: SubmitEvent) {
    e.preventDefault();
    updateFaskesSettings({
      name: faskesName,
      code: faskesCode,
      type: faskesType,
      address: faskesAddress,
      phone: faskesPhone,
      email: faskesEmail
    });

    showSuccessFaskes = true;
    setTimeout(() => {
      showSuccessFaskes = false;
      // Reload page to reflect faskes changes in Topbar
      window.location.reload();
    }, 1000);
  }
</script>

<svelte:head>
  <title>Pengaturan - Aplikasi Rujukan Pasien</title>
</svelte:head>

<!-- Page Heading -->
<div class="d-sm-flex align-items-center justify-content-between mb-4">
  <div>
    <h1 class="h3 mb-1 text-gray-800 font-weight-bold">Pengaturan</h1>
    <p class="text-muted mb-0 small">Kelola profil pengguna dan konfigurasi instansi fasilitas kesehatan Anda.</p>
  </div>
</div>

<div class="row">
  <!-- Left Side: User Profile Settings -->
  <div class="col-lg-6 mb-4">
    <div class="card shadow-sm border-0 h-100">
      <div class="card-header bg-white py-3 border-bottom d-flex align-items-center">
        <Fa icon={faUser} class="text-primary mr-2" />
        <h6 class="m-0 font-weight-bold text-primary">Profil Pengguna</h6>
      </div>
      <div class="card-body">
        {#if showSuccessUser}
          <div class="alert alert-success d-flex align-items-center py-2 mb-3 small" role="alert">
            <Fa icon={faCheckCircle} class="mr-2" /> Profil berhasil diperbarui!
          </div>
        {/if}

        <form onsubmit={handleSaveProfile}>
          <div class="form-group mb-3">
            <label for="profName" class="form-label small font-weight-bold text-gray-700">Nama Lengkap & Gelar</label>
            <input
              type="text"
              id="profName"
              class="form-control form-control-sm"
              bind:value={userName}
              required
            />
          </div>

          <div class="form-group mb-3">
            <label for="profRole" class="form-label small font-weight-bold text-gray-700">Jabatan / Peran Medis</label>
            <input
              type="text"
              id="profRole"
              class="form-control form-control-sm"
              bind:value={userRole}
              required
            />
          </div>

          <div class="form-group mb-4">
            <label for="profUser" class="form-label small font-weight-bold text-gray-700">Username / NIP <span class="text-xs text-muted">(ID Pengguna)</span></label>
            <div class="input-group">
              <input
                type="text"
                id="profUser"
                class="form-control form-control-sm bg-light"
                bind:value={userUsername}
                disabled
              />
              <div class="input-group-append bg-light border border-left-0 rounded-right pr-2 d-flex align-items-center justify-content-center">
                <span class="text-muted"><Fa icon={faLock} size="xs" /></span>
              </div>
            </div>
            <span class="text-xs text-muted mt-1 d-block">Username tidak dapat diubah setelah terdaftar.</span>
          </div>

          <button type="submit" class="btn btn-sm btn-primary rounded-pill px-4">
            <Fa icon={faSave} class="mr-1" /> Simpan Profil
          </button>
        </form>
      </div>
    </div>
  </div>

  <!-- Right Side: Faskes Facility Settings -->
  <div class="col-lg-6 mb-4">
    <div class="card shadow-sm border-0 h-100">
      <div class="card-header bg-white py-3 border-bottom d-flex align-items-center">
        <Fa icon={faHospital} class="text-primary mr-2" />
        <h6 class="m-0 font-weight-bold text-primary">Informasi Instansi (Faskes)</h6>
      </div>
      <div class="card-body">
        {#if showSuccessFaskes}
          <div class="alert alert-success d-flex align-items-center py-2 mb-3 small" role="alert">
            <Fa icon={faCheckCircle} class="mr-2" /> Data Faskes berhasil diperbarui!
          </div>
        {/if}

        <form onsubmit={handleSaveFaskes}>
          <div class="row">
            <div class="col-md-8 mb-3">
              <label for="fasName" class="form-label small font-weight-bold text-gray-700">Nama Fasilitas Kesehatan</label>
              <input
                type="text"
                id="fasName"
                class="form-control form-control-sm"
                bind:value={faskesName}
                required
              />
            </div>
            <div class="col-md-4 mb-3">
              <label for="fasCode" class="form-label small font-weight-bold text-gray-700">Kode Faskes</label>
              <input
                type="text"
                id="fasCode"
                class="form-control form-control-sm font-monospace"
                bind:value={faskesCode}
                required
              />
            </div>
          </div>

          <div class="form-group mb-3">
            <label for="fasType" class="form-label small font-weight-bold text-gray-700">Tipe / Kategori Faskes</label>
            <select id="fasType" class="form-control form-control-sm" bind:value={faskesType}>
              <option value="Puskesmas">Puskesmas (Faskes Tingkat I)</option>
              <option value="Klinik Pratama">Klinik Pratama</option>
              <option value="Rumah Sakit Tipe C">Rumah Sakit Tipe C (Rujukan)</option>
              <option value="Rumah Sakit Tipe B">Rumah Sakit Tipe B (Rujukan Utama)</option>
              <option value="Rumah Sakit Tipe A">Rumah Sakit Tipe A (Rujukan Nasional)</option>
            </select>
          </div>

          <div class="form-group mb-3">
            <label for="fasAddress" class="form-label small font-weight-bold text-gray-700">Alamat Lengkap</label>
            <textarea
              id="fasAddress"
              class="form-control form-control-sm"
              rows="2"
              bind:value={faskesAddress}
              required
            ></textarea>
          </div>

          <div class="row mb-4">
            <div class="col-sm-6 mb-3 mb-sm-0">
              <label for="fasPhone" class="form-label small font-weight-bold text-gray-700">No. Telepon / Fax</label>
              <input
                type="text"
                id="fasPhone"
                class="form-control form-control-sm"
                bind:value={faskesPhone}
              />
            </div>
            <div class="col-sm-6">
              <label for="fasEmail" class="form-label small font-weight-bold text-gray-700">Email Instansi</label>
              <input
                type="email"
                id="fasEmail"
                class="form-control form-control-sm"
                bind:value={faskesEmail}
              />
            </div>
          </div>

          <button type="submit" class="btn btn-sm btn-primary rounded-pill px-4">
            <Fa icon={faSave} class="mr-1" /> Simpan Data Faskes
          </button>
        </form>
      </div>
    </div>
  </div>
</div>

<style>
  .font-monospace {
    font-family: SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace !important;
  }
</style>
