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
    <p class="text-muted mb-0 small">Kirim rujukan pasien dari fasilitas kesehatan Anda ke rumah sakit rujukan tingkat lanjut.</p>
  </div>
  <button class="btn btn-primary btn-sm rounded-pill px-4 py-2 font-weight-bold shadow-sm mt-3 mt-sm-0" onclick={openCreateModal}>
    <Fa icon={faPlus} class="mr-1" /> Buat Rujukan Keluar
  </button>
</div>

<!-- Filter and Search controls -->
<div class="card shadow-sm mb-4 border-0">
  <div class="card-body">
    <div class="row align-items-center">
      <!-- Filter Status -->
      <div class="col-lg-6 col-md-12 mb-3 mb-lg-0">
        <div class="btn-group" role="group" aria-label="Filter Status">
          {#each ['Semua', 'Dikirim', 'Diproses', 'Selesai'] as status}
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
                <span class="badge" class:bg-success={status === 'Dikirim'} class:bg-warning={status === 'Diproses'} class:bg-primary={status === 'Selesai'}>
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
        <Fa icon={faPaperPlane} size="3x" class="text-gray-300 mb-3" />
        <h5 class="text-gray-500">Tidak ada data rujukan keluar</h5>
        <p class="text-muted small">Coba buat rujukan keluar baru dengan menekan tombol di pojok kanan atas.</p>
      </div>
    {:else}
      <div class="table-responsive">
        <table class="table table-hover align-middle mb-0">
          <thead class="bg-light text-gray-700 font-weight-bold text-xs">
            <tr>
              <th class="px-4 py-3">No. Rujukan</th>
              <th>Tgl Rujukan</th>
              <th>Faskes Tujuan</th>
              <th>Pasien</th>
              <th>Diagnosa Utama</th>
              <th class="text-center">Status</th>
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
                    {r.faskesTujuan}
                  </div>
                </td>
                <td>
                  <div class="font-weight-bold text-gray-800">{r.pasien.nama}</div>
                  <div class="text-xs text-muted">No: {r.pasien.noKartu} • Usia: {r.pasien.usia} th</div>
                </td>
                <td>{r.diagnosa}</td>
                <td class="text-center">
                  <span class="badge py-2 px-3 rounded-pill text-white" 
                    class:bg-success={r.status === 'Dikirim'} 
                    class:bg-warning={r.status === 'Diproses'} 
                    class:bg-primary={r.status === 'Selesai'}
                  >
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
  <div class="modal-backdrop fade show" style="z-index: 1040;"></div>
  <div class="modal fade show d-block" tabindex="-1" role="dialog" style="z-index: 1050; overflow-y: auto;">
    <div class="modal-dialog modal-lg modal-dialog-centered" role="document">
      <div class="modal-content border-0 shadow-lg rounded-4">
        <form onsubmit={handleSubmit}>
          <!-- Header -->
          <div class="modal-header bg-gradient-primary text-white border-0 py-3">
            <h5 class="modal-title font-weight-bold">
              <Fa icon={faUserPlus} class="mr-2" /> Buat Rujukan Keluar Baru
            </h5>
            <button type="button" class="btn bg-transparent border-0 text-white font-weight-bold" onclick={closeCreateModal}>
              <Fa icon={faTimes} />
            </button>
          </div>
          
          <!-- Body -->
          <div class="modal-body p-4">
            {#if formErrorMessage}
              <div class="alert alert-danger py-2 mb-3 text-center small" role="alert">
                {formErrorMessage}
              </div>
            {/if}

            <div class="row">
              <!-- Left Column: Patient Bio -->
              <div class="col-md-6 mb-3 mb-md-0">
                <h6 class="font-weight-bold text-primary mb-3">Identitas Pasien</h6>
                
                <div class="form-group mb-3">
                  <label for="namaPasien" class="form-label small font-weight-bold text-gray-700">Nama Pasien *</label>
                  <input
                    type="text"
                    id="namaPasien"
                    class="form-control form-control-sm"
                    placeholder="Masukkan nama lengkap pasien"
                    bind:value={inputNama}
                    required
                  />
                </div>

                <div class="form-group mb-3">
                  <label for="nikPasien" class="form-label small font-weight-bold text-gray-700">NIK Pasien *</label>
                  <input
                    type="text"
                    id="nikPasien"
                    class="form-control form-control-sm"
                    placeholder="Masukkan 16 digit NIK"
                    bind:value={inputNik}
                    required
                  />
                </div>

                <div class="form-group mb-3">
                  <label for="noKartu" class="form-label small font-weight-bold text-gray-700">No. Kartu BPJS *</label>
                  <input
                    type="text"
                    id="noKartu"
                    class="form-control form-control-sm"
                    placeholder="Masukkan 13 digit nomor BPJS"
                    bind:value={inputNoKartu}
                    required
                  />
                </div>

                <div class="row">
                  <div class="col-sm-6 mb-3 mb-sm-0">
                    <label for="gender" class="form-label small font-weight-bold text-gray-700">Jenis Kelamin</label>
                    <select id="gender" class="form-control form-control-sm" bind:value={inputGender}>
                      <option value="Laki-laki">Laki-laki</option>
                      <option value="Perempuan">Perempuan</option>
                    </select>
                  </div>
                  <div class="col-sm-6">
                    <label for="usia" class="form-label small font-weight-bold text-gray-700">Usia *</label>
                    <input
                      type="number"
                      id="usia"
                      class="form-control form-control-sm"
                      placeholder="Usia"
                      bind:value={inputUsia}
                      required
                    />
                  </div>
                </div>
              </div>

              <!-- Right Column: Destination Faskes & Med Info -->
              <div class="col-md-6">
                <h6 class="font-weight-bold text-primary mb-3">Tujuan & Detail Medis</h6>
                
                <div class="form-group mb-3">
                  <label for="faskesTujuan" class="form-label small font-weight-bold text-gray-700">Faskes Tujuan Rujukan *</label>
                  <select id="faskesTujuan" class="form-control form-control-sm" bind:value={inputFaskesTujuan}>
                    <option value="RSCM (Rumah Sakit Cipto Mangunkusumo)">RSCM Jakarta</option>
                    <option value="RS Kanker Dharmais">RS Kanker Dharmais</option>
                    <option value="RS Jantung Harapan Kita">RS Jantung Harapan Kita</option>
                    <option value="RSUD Fatmawati">RSUD Fatmawati</option>
                    <option value="RSUP Persahabatan">RSUP Persahabatan</option>
                  </select>
                </div>

                <div class="form-group mb-3">
                  <label for="poliTujuan" class="form-label small font-weight-bold text-gray-700">Spesialisasi / Poli Tujuan *</label>
                  <select id="poliTujuan" class="form-control form-control-sm" bind:value={inputPoliTujuan}>
                    <option value="Poli Penyakit Dalam">Poli Penyakit Dalam (Interna)</option>
                    <option value="Poli Jantung & Pembuluh Darah">Poli Jantung & Pembuluh Darah</option>
                    <option value="Poli Bedah Onkologi">Poli Bedah Onkologi</option>
                    <option value="Poli Bedah Umum">Poli Bedah Umum</option>
                    <option value="Poli Mata">Poli Mata</option>
                    <option value="Poli Kebidanan & Kandungan">Poli Kebidanan & Kandungan (Obgyn)</option>
                  </select>
                </div>

                <div class="form-group mb-3">
                  <label for="diagnosa" class="form-label small font-weight-bold text-gray-700">Diagnosa Utama *</label>
                  <input
                    type="text"
                    id="diagnosa"
                    class="form-control form-control-sm"
                    placeholder="Contoh: Chronic Kidney Disease Stage V"
                    bind:value={inputDiagnosa}
                    required
                  />
                </div>

                <div class="form-group mb-3">
                  <label for="catatan" class="form-label small font-weight-bold text-gray-700">Catatan Klinis / Keterangan Tambahan</label>
                  <textarea
                    id="catatan"
                    class="form-control form-control-sm"
                    rows="3"
                    placeholder="Masukkan indikasi rujukan medis..."
                    bind:value={inputCatatan}
                  ></textarea>
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="modal-footer bg-light border-0 d-flex justify-content-between p-3">
            <button type="button" class="btn btn-sm btn-outline-secondary px-4 rounded-pill" onclick={closeCreateModal}>
              Batal
            </button>
            <button type="submit" class="btn btn-sm btn-primary px-4 rounded-pill">
              <Fa icon={faPaperPlane} class="mr-1" /> Kirim Rujukan
            </button>
          </div>
        </form>
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
</style>
