<script>
  import { onMount } from 'svelte';
  import Fa from 'svelte-fa';
  import {
    faSearch,
    faPlus,
    faPaperPlane,
    faCalendar,
    faHospital,
    faTimes,
    faUserPlus
  } from '@fortawesome/free-solid-svg-icons';
  import { getRujukanKeluar, saveRujukanKeluar } from '$lib/auth';

  /** @type {any[]} */
  let rujukanList = $state([]);
  let searchQuery = $state('');
  let filterStatus = $state('Semua'); // Semua, Dikirim, Diproses, Selesai

  // Modal State
  let showCreateModal = $state(false);
  
  // New Rujukan Form Inputs
  let inputNama = $state('');
  let inputNik = $state('');
  let inputNoKartu = $state('');
  let inputGender = $state('Laki-laki');
  let inputUsia = $state('');
  let inputFaskesTujuan = $state('RSUD Pasar Minggu');
  let inputPoliTujuan = $state('Poli Penyakit Dalam');
  let inputDiagnosa = $state('');
  let inputCatatan = $state('');
  let formErrorMessage = $state('');

  onMount(() => {
    rujukanList = getRujukanKeluar();
  });

  // Filter & Search Outgoing Referrals
  let filteredList = $derived(() => {
    return rujukanList.filter(r => {
      const matchStatus = filterStatus === 'Semua' || r.status === filterStatus;
      const query = searchQuery.toLowerCase();
      const matchSearch = 
        r.noRujukan.toLowerCase().includes(query) ||
        r.pasien.nama.toLowerCase().includes(query) ||
        r.faskesTujuan.toLowerCase().includes(query) ||
        r.diagnosa.toLowerCase().includes(query);
      return matchStatus && matchSearch;
    });
  });

  function openCreateModal() {
    inputNama = '';
    inputNik = '';
    inputNoKartu = '';
    inputGender = 'Laki-laki';
    inputUsia = '';
    inputFaskesTujuan = 'RSUD Pasar Minggu';
    inputPoliTujuan = 'Poli Penyakit Dalam';
    inputDiagnosa = '';
    inputCatatan = '';
    formErrorMessage = '';
    showCreateModal = true;
  }

  function closeCreateModal() {
    showCreateModal = false;
  }

  /**
   * @param {SubmitEvent} e
   */
  function handleSubmit(e) {
    e.preventDefault();
    formErrorMessage = '';

    if (!inputNama || !inputNik || !inputNoKartu || !inputUsia || !inputDiagnosa) {
      formErrorMessage = 'Semua kolom bertanda bintang (*) wajib diisi!';
      return;
    }

    // Generate random Referral number
    const randomSuffix = Math.floor(1000 + Math.random() * 9000);
    const dateCode = new Date().toISOString().slice(0, 10).replace(/-/g, '').slice(2, 6);
    const generatedNoRujukan = `3174R002${dateCode}K${randomSuffix}`;

    const newRujukan = {
      id: `RK-2026-${Date.now()}`,
      noRujukan: generatedNoRujukan,
      tglRujukan: new Date().toISOString().split('T')[0],
      pasien: {
        nama: inputNama,
        nik: inputNik,
        noKartu: inputNoKartu,
        gender: inputGender,
        usia: parseInt(inputUsia)
      },
      faskesTujuan: inputFaskesTujuan,
      poliTujuan: inputPoliTujuan,
      diagnosa: inputDiagnosa,
      status: 'Dikirim', // Default status for fresh outgoing
      catatan: inputCatatan
    };

    rujukanList = [newRujukan, ...rujukanList];
    saveRujukanKeluar(rujukanList);
    closeCreateModal();
  }
</script>

<svelte:head>
  <title>Rujukan Keluar - Aplikasi Rujukan Pasien</title>
</svelte:head>

<!-- Page Heading -->
<div class="d-sm-flex align-items-center justify-content-between mb-4">
  <div>
    <h1 class="h3 mb-1 text-gray-800 font-weight-bold">Rujukan Keluar</h1>
    <p class="text-muted mb-0 small">Kirim dan pantau rujukan pasien ke rumah sakit rujukan tingkat lanjut.</p>
  </div>
  <button class="btn-modern-primary py-2 px-4 shadow-sm mt-3 mt-sm-0" onclick={openCreateModal}>
    <Fa icon={faPlus} class="mr-1" /> Buat Rujukan Keluar
  </button>
</div>

<!-- Filter and Search controls -->
<div class="card-modern mb-4">
  <div class="card-body p-3">
    <div class="row align-items-center">
      <!-- Filter Status -->
      <div class="col-lg-7 mb-3 mb-lg-0">
        <div class="filter-pills-wrapper d-flex flex-wrap gap-2">
          {#each ['Semua', 'Dikirim', 'Diproses', 'Selesai'] as status}
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
            placeholder="Cari No. Rujukan, Pasien, Rumah Sakit..."
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
          <Fa icon={faPaperPlane} size="3x" />
        </div>
        <h5 class="text-gray-500 font-weight-bold">Tidak ada data rujukan keluar</h5>
        <p class="text-muted small">Kirim rujukan baru dengan menekan tombol buat rujukan di pojok kanan atas.</p>
      </div>
    {:else}
      <div class="table-responsive">
        <table class="table-modern">
          <thead>
            <tr>
              <th>No. Rujukan</th>
              <th>Tanggal Kirim</th>
              <th>Faskes Penerima</th>
              <th>Pasien</th>
              <th>Diagnosa Utama</th>
              <th class="text-center">Status</th>
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
                    <span>{r.faskesTujuan}</span>
                  </div>
                </td>
                <td>
                  <div class="font-weight-bold">{r.pasien.nama}</div>
                  <div class="text-xs text-muted">No: {r.pasien.noKartu} • {r.pasien.usia} th</div>
                </td>
                <td>{r.diagnosa}</td>
                <td class="text-center">
                  <span class="badge-modern" 
                    class:badge-modern-success={r.status === 'Dikirim'} 
                    class:badge-modern-warning={r.status === 'Diproses'} 
                    class:badge-modern-info={r.status === 'Selesai'}
                  >
                    <i class="fas fa-circle text-xs mr-1" style="font-size: 6px;"></i>
                    {r.status}
                  </span>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  </div>
</div>

<!-- Create Outgoing Referral Modal -->
{#if showCreateModal}
  <div class="modal-backdrop-custom animate-fade-in"></div>
  <div class="modal d-block animate-fade-in" tabindex="-1" role="dialog" style="overflow-y: auto;">
    <div class="modal-dialog modal-lg modal-dialog-centered" role="document">
      <div class="modal-content modal-modern-content">
        <form onsubmit={handleSubmit}>
          <!-- Header -->
          <div class="modal-header modal-modern-header d-flex justify-content-between align-items-center text-white">
            <h5 class="modal-title font-weight-bold mb-0">
              <Fa icon={faUserPlus} class="mr-2" /> Buat Rujukan Keluar Baru
            </h5>
            <button type="button" class="btn bg-transparent border-0 text-white font-weight-bold p-0" onclick={closeCreateModal}>
              <Fa icon={faTimes} />
            </button>
          </div>
          
          <!-- Body -->
          <div class="modal-body modal-modern-body">
            {#if formErrorMessage}
              <div class="alert alert-danger-modern py-2 px-3 mb-3 text-center small" role="alert">
                <i class="fas fa-exclamation-circle mr-1"></i> {formErrorMessage}
              </div>
            {/if}

            <div class="row">
              <!-- Left Column: Patient Bio -->
              <div class="col-md-6 mb-3 mb-md-0 border-end border-color-light">
                <h6 class="font-weight-bold text-primary mb-3">Identitas Pasien</h6>
                
                <div class="mb-3">
                  <label for="namaPasien" class="form-label">Nama Lengkap Pasien *</label>
                  <input
                    type="text"
                    id="namaPasien"
                    class="form-control"
                    placeholder="Masukkan nama lengkap pasien"
                    bind:value={inputNama}
                    required
                  />
                </div>

                <div class="mb-3">
                  <label for="nikPasien" class="form-label">NIK Pasien (16 Digit) *</label>
                  <input
                    type="text"
                    id="nikPasien"
                    class="form-control"
                    placeholder="Masukkan 16 digit NIK"
                    bind:value={inputNik}
                    required
                  />
                </div>

                <div class="mb-3">
                  <label for="noKartu" class="form-label">No. Kartu BPJS *</label>
                  <input
                    type="text"
                    id="noKartu"
                    class="form-control"
                    placeholder="Masukkan 13 digit nomor BPJS"
                    bind:value={inputNoKartu}
                    required
                  />
                </div>

                <div class="row">
                  <div class="col-sm-6 mb-3 mb-sm-0">
                    <label for="gender" class="form-label">Jenis Kelamin</label>
                    <select id="gender" class="form-select" bind:value={inputGender}>
                      <option value="Laki-laki">Laki-laki</option>
                      <option value="Perempuan">Perempuan</option>
                    </select>
                  </div>
                  <div class="col-sm-6">
                    <label for="usia" class="form-label">Usia (Tahun) *</label>
                    <input
                      type="number"
                      id="usia"
                      class="form-control"
                      placeholder="Usia"
                      bind:value={inputUsia}
                      required
                    />
                  </div>
                </div>
              </div>

              <!-- Right Column: Destination Faskes & Med Info -->
              <div class="col-md-6 ps-md-4">
                <h6 class="font-weight-bold text-primary mb-3">Instansi Tujuan & Medis</h6>
                
                <div class="mb-3">
                  <label for="faskesTujuan" class="form-label">Fasilitas Kesehatan Tujuan *</label>
                  <select id="faskesTujuan" class="form-select" bind:value={inputFaskesTujuan}>
                    <option value="RSCM (Rumah Sakit Cipto Mangunkusumo)">RSCM Jakarta</option>
                    <option value="RS Kanker Dharmais">RS Kanker Dharmais</option>
                    <option value="RS Jantung Harapan Kita">RS Jantung Harapan Kita</option>
                    <option value="RSUD Fatmawati">RSUD Fatmawati</option>
                    <option value="RSUP Persahabatan">RSUP Persahabatan</option>
                  </select>
                </div>

                <div class="mb-3">
                  <label for="poliTujuan" class="form-label">Spesialisasi / Poli Tujuan *</label>
                  <select id="poliTujuan" class="form-select" bind:value={inputPoliTujuan}>
                    <option value="Poli Penyakit Dalam">Poli Penyakit Dalam (Interna)</option>
                    <option value="Poli Jantung & Pembuluh Darah">Poli Jantung & Pembuluh Darah</option>
                    <option value="Poli Bedah Onkologi">Poli Bedah Onkologi</option>
                    <option value="Poli Bedah Umum">Poli Bedah Umum</option>
                    <option value="Poli Mata">Poli Mata</option>
                    <option value="Poli Kebidanan & Kandungan">Poli Kebidanan & Kandungan (Obgyn)</option>
                  </select>
                </div>

                <div class="mb-3">
                  <label for="diagnosa" class="form-label">Diagnosa Utama Medis *</label>
                  <input
                    type="text"
                    id="diagnosa"
                    class="form-control"
                    placeholder="Contoh: CKD Stage V / Heart Failure"
                    bind:value={inputDiagnosa}
                    required
                  />
                </div>

                <div class="mb-3">
                  <label for="catatan" class="form-label">Catatan Klinis / Keterangan Medis</label>
                  <textarea
                    id="catatan"
                    class="form-control"
                    rows="3"
                    placeholder="Masukkan indikasi klinis rujukan medis..."
                    bind:value={inputCatatan}
                  ></textarea>
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="modal-footer modal-modern-footer d-flex justify-content-between">
            <button type="button" class="btn btn-modern-secondary" onclick={closeCreateModal}>
              Batal
            </button>
            <button type="submit" class="btn btn-modern-primary py-2 px-4">
              <Fa icon={faPaperPlane} class="mr-1" /> Kirim Rujukan
            </button>
          </div>
        </form>
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

  .badge-color-dikirim { background-color: var(--success-light); color: var(--success-dark); }
  .badge-color-diproses { background-color: var(--warning-light); color: var(--warning-dark); }
  .badge-color-selesai { background-color: var(--info-light); color: var(--info-dark); }

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

  .alert-danger-modern {
    background-color: var(--danger-light);
    color: var(--danger-dark);
    border: 1px solid rgba(244, 63, 94, 0.2);
    border-radius: var(--radius-md);
    font-weight: 600;
  }
</style>
