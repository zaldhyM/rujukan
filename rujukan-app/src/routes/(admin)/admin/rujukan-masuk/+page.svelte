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
    <p class="text-muted mb-0 small">Kelola dan tinjau pengajuan rujukan dari fasilitas kesehatan lain.</p>
  </div>
</div>

<!-- Filter and Search controls -->
<div class="card-modern mb-4">
  <div class="card-body p-3">
    <div class="row align-items-center">
      <!-- Filter Status (Pills layout) -->
      <div class="col-lg-7 mb-3 mb-lg-0">
        <div class="filter-pills-wrapper d-flex flex-wrap gap-2">
          {#each ['Semua', 'Pending', 'Diterima', 'Ditolak'] as status}
            <button
              type="button"
              class="btn-filter-pill"
              class:active={filterStatus === status}
              onclick={() => filterStatus = status}
            >
              <span class="pill-text">{status}</span>
              {#if status === 'Semua'}
                <span class="pill-badge">{rujukanList.length}</span>
              {:else}
                <span class="pill-badge badge-color-{status.toLowerCase()}">
                  {rujukanList.filter(r => r.status === status).length}
                </span>
              {/if}
            </button>
          {/each}
        </div>
      </div>
      
      <!-- Search Input -->
      <div class="col-lg-5">
        <div class="input-icon-wrapper">
          <i class="fas fa-search input-icon" style="color: var(--text-light);"></i>
          <input
            type="text"
            class="form-control ps-5"
            placeholder="Cari No. Rujukan, Nama Pasien, Diagnosa..."
            bind:value={searchQuery}
          />
        </div>
      </div>
    </div>
  </div>
</div>

<!-- Table / Content -->
<div class="card-modern">
  <div class="card-body p-0">
    {#if filteredList().length === 0}
      <div class="text-center py-5">
        <div class="empty-icon mb-3">
          <Fa icon={faInbox} size="3x" />
        </div>
        <h5 class="text-gray-500 font-weight-bold">Tidak ada data rujukan masuk</h5>
        <p class="text-muted small">Coba sesuaikan filter atau kata kunci pencarian Anda.</p>
      </div>
    {:else}
      <div class="table-responsive">
        <table class="table-modern">
          <thead>
            <tr>
              <th>No. Rujukan</th>
              <th>Tanggal Kirim</th>
              <th>Faskes Asal</th>
              <th>Pasien</th>
              <th>Diagnosa Utama</th>
              <th class="text-center">Status</th>
              <th class="text-center">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {#each filteredList() as r (r.id)}
              <tr>
                <td class="font-weight-bold text-primary">{r.noRujukan}</td>
                <td>
                  <div class="d-flex align-items-center gap-1">
                    <Fa icon={faCalendar} class="text-muted" />
                    <span>{r.tglRujukan}</span>
                  </div>
                </td>
                <td class="font-weight-bold">
                  <div class="d-flex align-items-center gap-1">
                    <Fa icon={faHospital} class="text-muted" />
                    <span>{r.faskesAsal}</span>
                  </div>
                </td>
                <td>
                  <div class="font-weight-bold">{r.pasien.nama}</div>
                  <div class="text-xs text-muted">No: {r.pasien.noKartu} • {r.pasien.usia} th</div>
                </td>
                <td>{r.diagnosa}</td>
                <td class="text-center">
                  <span class="badge-modern" 
                    class:badge-modern-warning={r.status === 'Pending'} 
                    class:badge-modern-success={r.status === 'Diterima'} 
                    class:badge-modern-danger={r.status === 'Ditolak'}
                  >
                    <i class="fas fa-circle text-xs mr-1" style="font-size: 6px;"></i>
                    {r.status}
                  </span>
                </td>
                <td class="text-center">
                  <button class="btn btn-sm btn-outline-primary px-3 rounded-pill font-weight-bold" onclick={() => openDetail(r)}>
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
  <div class="modal-backdrop-custom animate-fade-in"></div>
  <div class="modal d-block animate-fade-in" tabindex="-1" role="dialog" style="overflow-y: auto;">
    <div class="modal-dialog modal-lg modal-dialog-centered" role="document">
      <div class="modal-content modal-modern-content">
        <!-- Header -->
        <div class="modal-header modal-modern-header d-flex justify-content-between align-items-center text-white">
          <h5 class="modal-title font-weight-bold mb-0">
            <Fa icon={faInbox} class="mr-2" /> Detail Rujukan Masuk
          </h5>
          <button type="button" class="btn bg-transparent border-0 text-white font-weight-bold p-0" onclick={closeDetail}>
            <Fa icon={faTimes} />
          </button>
        </div>
        
        <!-- Body -->
        <div class="modal-body modal-modern-body">
          <!-- Rujukan Header Card -->
          <div class="kpi-card p-3 mb-4 bg-light border-0">
            <div class="row align-items-center">
              <div class="col-sm-6">
                <span class="text-xs text-muted d-block uppercase font-weight-bold mb-1">Nomor Rujukan SISRUTE</span>
                <span class="h6 font-weight-bold text-primary font-monospace">{selectedRujukan.noRujukan}</span>
              </div>
              <div class="col-sm-6 text-sm-end mt-2 mt-sm-0">
                <span class="text-xs text-muted d-block uppercase font-weight-bold mb-1">Tanggal Pengiriman</span>
                <span class="h6 font-weight-bold text-gray-800">{selectedRujukan.tglRujukan}</span>
              </div>
            </div>
          </div>

          <div class="row">
            <!-- Left Column: Patient & Sender Info -->
            <div class="col-md-6 mb-4 mb-md-0 border-end border-color-light">
              <h6 class="font-weight-bold text-primary mb-3"><Fa icon={faUser} class="mr-1" /> Data Identitas Pasien</h6>
              
              <div class="patient-info-list mb-4">
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">Nama Lengkap</span>
                  <span class="font-weight-bold text-main">{selectedRujukan.pasien.nama}</span>
                </div>
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">NIK Pasien</span>
                  <span class="text-main font-monospace">{selectedRujukan.pasien.nik}</span>
                </div>
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">No. BPJS Kesehatan</span>
                  <span class="text-main font-monospace">{selectedRujukan.pasien.noKartu}</span>
                </div>
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">Jenis Kelamin</span>
                  <span class="text-main">{selectedRujukan.pasien.gender}</span>
                </div>
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">Usia Pasien</span>
                  <span class="text-main font-weight-bold">{selectedRujukan.pasien.usia} Tahun</span>
                </div>
              </div>

              <h6 class="font-weight-bold text-primary mb-3"><Fa icon={faHospital} class="mr-1" /> Asal Instansi Pengirim</h6>
              <div class="patient-info-list">
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">Faskes Perujuk</span>
                  <span class="font-weight-bold text-main">{selectedRujukan.faskesAsal}</span>
                </div>
                <div class="info-row py-2 d-flex justify-content-between">
                  <span class="text-muted small">Poli / Spesialisasi</span>
                  <span class="text-main font-weight-bold text-primary">{selectedRujukan.poliTujuan}</span>
                </div>
              </div>
            </div>

            <!-- Right Column: Medical Info & Response Action -->
            <div class="col-md-6 ps-md-4">
              <h6 class="font-weight-bold text-primary mb-3"><Fa icon={faIdCard} class="mr-1" /> Indikasi Klinis & Diagnosa</h6>
              
              <div class="medical-condition-card p-3 mb-4 rounded-3">
                <span class="text-xs text-muted d-block uppercase font-weight-bold mb-1">Diagnosa Utama</span>
                <span class="d-block font-weight-bold text-danger">{selectedRujukan.diagnosa}</span>
                
                <hr class="my-2 border-color-light" />
                
                <span class="text-xs text-muted d-block uppercase font-weight-bold mb-1">Keterangan Medis</span>
                <span class="d-block text-gray-700 italic small mt-1">"{selectedRujukan.catatan || '-'}"</span>
              </div>

              <!-- Response Form -->
              <h6 class="font-weight-bold text-primary mb-2">Tanggapan Peninjauan Medis</h6>
              <div class="form-group mb-3">
                <textarea
                  class="form-control form-control-sm small"
                  rows="3"
                  placeholder="Berikan catatan persetujuan / alasan penolakan rujukan..."
                  bind:value={respondNote}
                ></textarea>
              </div>

              <!-- Status Badge -->
              <div class="d-flex align-items-center justify-content-between p-3 rounded-3 border mb-3 bg-light">
                <span class="small text-muted font-weight-bold">Status Saat Ini:</span>
                <span class="badge-modern" 
                  class:badge-modern-warning={selectedRujukan.status === 'Pending'} 
                  class:badge-modern-success={selectedRujukan.status === 'Diterima'} 
                  class:badge-modern-danger={selectedRujukan.status === 'Ditolak'}
                >
                  {selectedRujukan.status}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="modal-footer modal-modern-footer d-flex justify-content-between">
          <button type="button" class="btn btn-modern-secondary" onclick={closeDetail}>
            <Fa icon={faArrowLeft} class="mr-1" /> Kembali
          </button>
          
          <div class="d-flex gap-2">
            <button
              type="button"
              class="btn btn-modern-danger py-2 px-3 mr-2"
              onclick={() => updateStatus('Ditolak')}
            >
              <Fa icon={faTimes} class="mr-1" /> Tolak Rujukan
            </button>
            <button
              type="button"
              class="btn btn-modern-primary py-2 px-3"
              onclick={() => updateStatus('Diterima')}
            >
              <Fa icon={faCheck} class="mr-1" /> Setujui Rujukan
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal-backdrop-custom {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(11, 15, 25, 0.6);
    backdrop-filter: blur(5px);
    z-index: 1040;
  }
  .modal {
    z-index: 1050;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }
  
  .empty-icon {
    width: 60px;
    height: 60px;
    margin: 0 auto;
    background-color: #f1f5f9;
    color: var(--text-light);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .border-color-light {
    border-color: var(--border-color) !important;
  }

  /* Filter Pills Custom Styles */
  .btn-filter-pill {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.4rem 1rem;
    border-radius: var(--radius-pill);
    border: 1px solid var(--border-color);
    background-color: #ffffff;
    font-weight: 700;
    font-size: 0.8rem;
    color: var(--text-muted);
    cursor: pointer;
    transition: var(--transition);
  }

  .btn-filter-pill:hover {
    border-color: var(--primary);
    color: var(--primary);
  }

  .btn-filter-pill.active {
    background-color: var(--primary);
    border-color: var(--primary);
    color: #ffffff;
  }

  .pill-badge {
    padding: 0.15rem 0.45rem;
    font-size: 0.7rem;
    border-radius: var(--radius-pill);
    font-weight: 800;
    background-color: #f1f5f9;
    color: var(--text-muted);
  }

  .btn-filter-pill.active .pill-badge {
    background-color: rgba(255, 255, 255, 0.2);
    color: #ffffff;
  }

  .badge-color-pending { background-color: var(--warning-light); color: var(--warning-dark); }
  .badge-color-diterima { background-color: var(--success-light); color: var(--success-dark); }
  .badge-color-ditolak { background-color: var(--danger-light); color: var(--danger-dark); }

  /* Input icon wrapper */
  .input-icon-wrapper {
    position: relative;
  }
  .input-icon {
    position: absolute;
    left: 1rem;
    top: 50%;
    transform: translateY(-50%);
  }
  .ps-5 {
    padding-left: 2.5rem !important;
  }

  /* Patient Details List */
  .info-row {
    border-bottom: 1px dashed var(--border-color);
  }
  .info-row:last-child {
    border-bottom: none;
  }

  /* Medical Indication Card */
  .medical-condition-card {
    background-color: var(--danger-light);
    border-left: 4px solid var(--danger);
  }
  
  .font-monospace {
    font-family: SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace !important;
  }
</style>
