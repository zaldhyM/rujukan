<script>
  import { onMount } from 'svelte';
  import Fa from 'svelte-fa';
  import {
    faSearch,
    faInfoCircle,
    faCheck,
    faTimes,
    faArrowLeft,
    faInbox,
    faCalendar,
    faHospital,
    faIdCard,
    faUser
  } from '@fortawesome/free-solid-svg-icons';
  import { getRujukanMasuk, saveRujukanMasuk } from '$lib/auth';

  /** @type {any[]} */
  let rujukanList = $state([]);
  let filterStatus = $state('Semua'); // Semua, Pending, Diterima, Ditolak
  let searchQuery = $state('');
  
  // Modal state
  /** @type {any} */
  let selectedRujukan = $state(null);
  let detailModalOpen = $state(false);
  let respondNote = $state('');

  onMount(() => {
    rujukanList = getRujukanMasuk();
  });

  // Filter & search logic
  let filteredList = $derived(() => {
    return rujukanList.filter(r => {
      const matchStatus = filterStatus === 'Semua' || r.status === filterStatus;
      const query = searchQuery.toLowerCase();
      const matchSearch = 
        r.noRujukan.toLowerCase().includes(query) ||
        r.pasien.nama.toLowerCase().includes(query) ||
        r.faskesAsal.toLowerCase().includes(query) ||
        r.diagnosa.toLowerCase().includes(query);
      return matchStatus && matchSearch;
    });
  });

  /**
   * @param {any} rujukan
   */
  function openDetail(rujukan) {
    selectedRujukan = { ...rujukan };
    respondNote = rujukan.catatanResponse || '';
    detailModalOpen = true;
  }

  function closeDetail() {
    detailModalOpen = false;
    selectedRujukan = null;
  }

  /**
   * @param {string} newStatus
   */
  function updateStatus(newStatus) {
    if (!selectedRujukan) return;
    
    rujukanList = rujukanList.map(r => {
      if (selectedRujukan && r.id === selectedRujukan.id) {
        return {
          ...r,
          status: newStatus,
          catatanResponse: respondNote
        };
      }
      return r;
    });

    saveRujukanMasuk(rujukanList);
    closeDetail();
  }
</script>

<svelte:head>
  <title>Rujukan Masuk - Aplikasi Rujukan Pasien</title>
</svelte:head>

<!-- Page Heading -->
<div class="d-sm-flex align-items-center justify-content-between mb-4">
  <div>
    <h1 class="h3 mb-1 text-gray-800 font-weight-bold">Rujukan Masuk</h1>
    <p class="text-muted mb-0 small">Kelola rujukan pasien yang dikirim oleh fasilitas kesehatan lain ke faskes Anda.</p>
  </div>
</div>

<!-- Filter and Search controls -->
<div class="card shadow-sm mb-4 border-0">
  <div class="card-body">
    <div class="row align-items-center">
      <!-- Filter Status -->
      <div class="col-lg-6 col-md-12 mb-3 mb-lg-0">
        <div class="btn-group" role="group" aria-label="Filter Status">
          {#each ['Semua', 'Pending', 'Diterima', 'Ditolak'] as status}
            <button
              type="button"
              class="btn btn-sm"
              class:btn-primary={filterStatus === status}
              class:btn-outline-primary={filterStatus !== status}
              onclick={() => filterStatus = status}
            >
              {status}
              {#if status === 'Semua'}
                <span class="badge bg-white text-primary ml-1">{rujukanList.length}</span>
              {:else}
                <span class="badge" class:bg-warning={status === 'Pending'} class:bg-success={status === 'Diterima'} class:bg-danger={status === 'Ditolak'}>
                  {rujukanList.filter(r => r.status === status).length}
                </span>
              {/if}
            </button>
          {/each}
        </div>
      </div>
      
      <!-- Search Input -->
      <div class="col-lg-6 col-md-12">
        <div class="input-group">
          <input
            type="text"
            class="form-control form-control-sm border-right-0"
            placeholder="Cari No. Rujukan, Nama Pasien, Diagnosa, atau Faskes..."
            bind:value={searchQuery}
          />
          <div class="input-group-append bg-white border border-left-0 rounded-right pr-2 d-flex align-items-center justify-content-center">
            <span class="text-muted"><Fa icon={faSearch} size="xs" /></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<!-- Table / Content -->
<div class="card shadow-sm border-0">
  <div class="card-body p-0">
    {#if filteredList().length === 0}
      <div class="text-center py-5">
        <Fa icon={faInbox} size="3x" class="text-gray-300 mb-3" />
        <h5 class="text-gray-500">Tidak ada data rujukan masuk</h5>
        <p class="text-muted small">Coba sesuaikan filter atau kata kunci pencarian Anda.</p>
      </div>
    {:else}
      <div class="table-responsive">
        <table class="table table-hover align-middle mb-0">
          <thead class="bg-light text-gray-700 font-weight-bold text-xs">
            <tr>
              <th class="px-4 py-3">No. Rujukan</th>
              <th>Tgl Rujukan</th>
              <th>Faskes Asal</th>
              <th>Pasien</th>
              <th>Diagnosa Utama</th>
              <th class="text-center">Status</th>
              <th class="text-center">Aksi</th>
            </tr>
          </thead>
          <tbody class="small">
            {#each filteredList() as r (r.id)}
              <tr>
                <td class="px-4 py-3 font-weight-bold text-primary">{r.noRujukan}</td>
                <td>
                  <div class="d-flex align-items-center">
                    <Fa icon={faCalendar} class="mr-1 text-muted" />
                    {r.tglRujukan}
                  </div>
                </td>
                <td class="font-weight-bold text-gray-800">
                  <div class="d-flex align-items-center">
                    <Fa icon={faHospital} class="mr-1 text-muted text-xs" />
                    {r.faskesAsal}
                  </div>
                </td>
                <td>
                  <div class="font-weight-bold text-gray-800">{r.pasien.nama}</div>
                  <div class="text-xs text-muted">No: {r.pasien.noKartu} • Usia: {r.pasien.usia} th</div>
                </td>
                <td>{r.diagnosa}</td>
                <td class="text-center">
                  <span class="badge py-2 px-3 rounded-pill text-white" 
                    class:bg-warning={r.status === 'Pending'} 
                    class:bg-success={r.status === 'Diterima'} 
                    class:bg-danger={r.status === 'Ditolak'}
                  >
                    {r.status}
                  </span>
                </td>
                <td class="text-center">
                  <button class="btn btn-sm btn-outline-primary px-3 rounded-pill" onclick={() => openDetail(r)}>
                    <Fa icon={faInfoCircle} class="mr-1" /> Detail
                  </button>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  </div>
</div>

<!-- Detail Modal -->
{#if detailModalOpen && selectedRujukan}
  <div class="modal-backdrop fade show" style="z-index: 1040;"></div>
  <div class="modal fade show d-block" tabindex="-1" role="dialog" style="z-index: 1050; overflow-y: auto;">
    <div class="modal-dialog modal-lg modal-dialog-centered" role="document">
      <div class="modal-content border-0 shadow-lg rounded-4">
        <!-- Header -->
        <div class="modal-header bg-gradient-primary text-white border-0 py-3">
          <h5 class="modal-title font-weight-bold">
            <Fa icon={faInbox} class="mr-2" /> Detail Rujukan Masuk
          </h5>
          <button type="button" class="btn bg-transparent border-0 text-white font-weight-bold" onclick={closeDetail}>
            <Fa icon={faTimes} />
          </button>
        </div>
        
        <!-- Body -->
        <div class="modal-body p-4">
          <!-- Rujukan Header Card -->
          <div class="card bg-light border-0 mb-4 rounded-3">
            <div class="card-body py-3 px-4">
              <div class="row">
                <div class="col-sm-6">
                  <span class="text-xs text-muted d-block uppercase font-weight-bold">Nomor Rujukan</span>
                  <span class="h6 font-weight-bold text-primary">{selectedRujukan.noRujukan}</span>
                </div>
                <div class="col-sm-6 text-sm-right mt-2 mt-sm-0">
                  <span class="text-xs text-muted d-block uppercase font-weight-bold">Tanggal Kirim</span>
                  <span class="h6 font-weight-bold text-gray-800">{selectedRujukan.tglRujukan}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="row">
            <!-- Left Column: Patient & Sender Info -->
            <div class="col-md-6 mb-4 mb-md-0">
              <h6 class="font-weight-bold text-primary mb-3"><Fa icon={faUser} class="mr-1" /> Data Pasien</h6>
              <ul class="list-group list-group-flush mb-4">
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">Nama Lengkap</span>
                  <span class="font-weight-bold text-gray-800">{selectedRujukan.pasien.nama}</span>
                </li>
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">NIK</span>
                  <span class="text-gray-800 font-monospace">{selectedRujukan.pasien.nik}</span>
                </li>
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">No. BPJS</span>
                  <span class="text-gray-800 font-monospace">{selectedRujukan.pasien.noKartu}</span>
                </li>
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">Jenis Kelamin</span>
                  <span class="text-gray-800">{selectedRujukan.pasien.gender}</span>
                </li>
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">Usia</span>
                  <span class="text-gray-800">{selectedRujukan.pasien.usia} Tahun</span>
                </li>
              </ul>

              <h6 class="font-weight-bold text-primary mb-3"><Fa icon={faHospital} class="mr-1" /> Informasi Pengirim</h6>
              <ul class="list-group list-group-flush">
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">Faskes Perujuk</span>
                  <span class="font-weight-bold text-gray-800">{selectedRujukan.faskesAsal}</span>
                </li>
                <li class="list-group-item bg-transparent px-0 py-2 d-flex justify-content-between">
                  <span class="text-muted">Poli Tujuan</span>
                  <span class="text-gray-800 font-weight-bold">{selectedRujukan.poliTujuan}</span>
                </li>
              </ul>
            </div>

            <!-- Right Column: Medical Info & Response Action -->
            <div class="col-md-6">
              <h6 class="font-weight-bold text-primary mb-3"><Fa icon={faIdCard} class="mr-1" /> Diagnosa & Alasan</h6>
              <div class="card border-left-primary bg-light mb-4 rounded-3">
                <div class="card-body py-3 px-4">
                  <span class="text-xs text-muted d-block uppercase font-weight-bold">Diagnosa Utama</span>
                  <span class="d-block font-weight-bold text-gray-800">{selectedRujukan.diagnosa}</span>
                  <hr class="my-2" />
                  <span class="text-xs text-muted d-block uppercase font-weight-bold">Catatan Klinis</span>
                  <span class="d-block text-gray-700 italic small mt-1">"{selectedRujukan.catatan}"</span>
                </div>
              </div>

              <!-- Response Form -->
              <h6 class="font-weight-bold text-primary mb-2">Tanggapan Penerima</h6>
              <div class="form-group mb-3">
                <textarea
                  class="form-control form-control-sm small"
                  rows="3"
                  placeholder="Masukkan catatan tanggapan medis / alasan (opsional)..."
                  bind:value={respondNote}
                ></textarea>
              </div>

              <!-- Status Badge -->
              <div class="d-flex align-items-center justify-content-between p-2 rounded border mb-3 bg-white">
                <span class="small text-muted font-weight-bold">Status Saat Ini:</span>
                <span class="badge py-2 px-3 rounded-pill text-white" 
                  class:bg-warning={selectedRujukan.status === 'Pending'} 
                  class:bg-success={selectedRujukan.status === 'Diterima'} 
                  class:bg-danger={selectedRujukan.status === 'Ditolak'}
                >
                  {selectedRujukan.status}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="modal-footer bg-light border-0 d-flex justify-content-between p-3">
          <button type="button" class="btn btn-sm btn-outline-secondary px-4 rounded-pill" onclick={closeDetail}>
            <Fa icon={faArrowLeft} class="mr-1" /> Kembali
          </button>
          
          <div class="d-flex gap-2">
            <button
              type="button"
              class="btn btn-sm btn-danger px-4 rounded-pill mr-2"
              onclick={() => updateStatus('Ditolak')}
            >
              <Fa icon={faTimes} class="mr-1" /> Tolak Rujukan
            </button>
            <button
              type="button"
              class="btn btn-sm btn-success px-4 rounded-pill"
              onclick={() => updateStatus('Diterima')}
            >
              <Fa icon={faCheck} class="mr-1" /> Terima Rujukan
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal-backdrop {
    background-color: rgba(0, 0, 0, 0.5);
  }
  .card-body {
    border-radius: 0.5rem;
  }
  .font-monospace {
    font-family: SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace !important;
  }
</style>
