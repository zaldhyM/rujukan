<!-- src/routes/dashboard/rujukan-keluar/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';

  // Svelte 5 state
  let searchQuery = $state('');
  let showCreateModal = $state(false);
  let showDetailModal = $state(false);
  let selectedReferral = $state<any>(null);
  let showToast = $state(false);
  let toastMessage = $state('');

  // Form State
  let patientName = $state('');
  let patientNik = $state('');
  let patientAge = $state('');
  let patientGender = $state('Laki-laki');
  let destinationFaskes = $state('');
  let diagnosis = $state('');
  let rujukanType = $state('BPJS');
  let clinicalNotes = $state('');

  let referrals = $state([
    { 
      id: 'REF-K01', 
      patient: 'Rahmat Hidayat', 
      to: 'RSUP Dr. Wahidin Sudirohusodo', 
      date: 'Hari ini, 10:00', 
      status: 'Pending', 
      type: 'BPJS', 
      nik: '7371021404670002', 
      gender: 'Laki-laki', 
      age: '45 Tahun', 
      diagnosis: 'Hernia Inguinalis Lateralis (K40.9)', 
      notes: 'Benjolan di lipat paha kanan yang bisa keluar masuk. Nyeri saat mengangkat beban berat. Rencana operasi hernioplasty.' 
    },
    { 
      
      id: 'REF-K02', 
      patient: 'Siti Aminah', 
      to: 'RS Siloam Makassar', 
      date: 'Kemarin, 11:20', 
      status: 'Approved', 
      type: 'Umum', 
      nik: '7371084205880004', 
      gender: 'Perempuan', 
      age: '38 Tahun', 
      diagnosis: 'Cholelithiasis (K80.2)', 
      notes: 'Nyeri kolik kuadran kanan atas menjalar ke punggung belakang. Hasil USG abdomen: Cholelithiasis dengan diameter batu 1.2 cm.' 
    }
  ]);

  let filteredReferrals = $derived(
    referrals.filter(ref => 
      ref.patient.toLowerCase().includes(searchQuery.toLowerCase()) || 
      ref.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
      ref.to.toLowerCase().includes(searchQuery.toLowerCase()) ||
      ref.diagnosis.toLowerCase().includes(searchQuery.toLowerCase())
    )
  );

  const viewDetails = (ref: any) => {
    selectedReferral = ref;
    showDetailModal = true;
  };

  const handleCreateReferral = (e: SubmitEvent) => {
    e.preventDefault();

    if (!patientName || !patientNik || !patientAge || !destinationFaskes || !diagnosis || !clinicalNotes) {
      alert('Mohon isi semua field formulir.');
      return;
    }

    const newId = `REF-K0${referrals.length + 1}`;
    const now = new Date();
    const formattedDate = `Hari ini, ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}`;

    const newReferral = {
      id: newId,
      patient: patientName,
      to: destinationFaskes,
      date: formattedDate,
      status: 'Pending',
      type: rujukanType,
      nik: patientNik,
      gender: patientGender,
      age: `${patientAge} Tahun`,
      diagnosis: diagnosis,
      notes: clinicalNotes
    };

    referrals = [newReferral, ...referrals];

    // Reset Form
    patientName = '';
    patientNik = '';
    patientAge = '';
    patientGender = 'Laki-laki';
    destinationFaskes = '';
    diagnosis = '';
    rujukanType = 'BPJS';
    clinicalNotes = '';

    showCreateModal = false;

    // Toast feedback
    toastMessage = `Sukses membuat rujukan baru ${newId} untuk ${newReferral.patient}`;
    showToast = true;
    setTimeout(() => {
      showToast = false;
    }, 3000);
  };
</script>

<svelte:head>
  <title>Rujukan Keluar | RUJUKAN MEDIS</title>
</svelte:head>

<div class="rujukan-keluar-container animate-fade-in">
  <!-- Toast Notification -->
  {#if showToast}
    <div class="toast-notification glass-panel">
      <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="var(--color-success)" stroke-width="2.5">
        <polyline points="20 6 9 17 4 12"></polyline>
      </svg>
      <span>{toastMessage}</span>
    </div>
  {/if}

  <div class="page-header">
    <div class="header-left">
      <span class="breadcrumb">DASHBOARD / RUJUKAN KELUAR</span>
      <h1>Rujukan Keluar</h1>
      <p class="subtitle">Buat surat rujukan baru ke rumah sakit penerima spesialis / rujukan tingkat lanjut</p>
    </div>
    <div class="header-right">
      <button class="btn btn-primary" onclick={() => showCreateModal = true}>
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2.5">
          <line x1="12" y1="5" x2="12" y2="19"></line>
          <line x1="5" y1="12" x2="19" y2="12"></line>
        </svg>
        <span>Buat Rujukan Baru</span>
      </button>
    </div>
  </div>

  <!-- Search & Filter Panel -->
  <div class="glass-panel filter-panel">
    <div class="search-box-wrap">
      <svg class="search-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="11" cy="11" r="8"></circle>
        <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
      </svg>
      <input 
        type="text" 
        class="form-input search-input" 
        placeholder="Cari rujukan keluar berdasarkan nama pasien, ID rujukan, diagnosa..." 
        bind:value={searchQuery}
      />
    </div>
  </div>

  <!-- Rujukan Keluar Table -->
  <div class="glass-panel list-panel">
    <div class="table-responsive">
      <table class="rujukan-table">
        <thead>
          <tr>
            <th>ID Rujukan</th>
            <th>Nama Pasien</th>
            <th>Diagnosa</th>
            <th>Faskes Penerima (Tujuan)</th>
            <th>Tgl Keluar</th>
            <th>Status</th>
            <th class="actions-header">Aksi</th>
          </tr>
        </thead>
        <tbody>
          {#if filteredReferrals.length === 0}
            <tr>
              <td colspan="7" class="empty-cell">
                <svg viewBox="0 0 24 24" fill="none" width="48" height="48" stroke="var(--text-muted)" stroke-width="1.5">
                  <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path>
                </svg>
                <p>Tidak ditemukan data rujukan keluar.</p>
              </td>
            </tr>
          {:else}
            {#each filteredReferrals as ref}
              <tr class="table-row">
                <td class="ref-id font-title">{ref.id}</td>
                <td>
                  <div class="patient-info-cell">
                    <span class="patient-name">{ref.patient}</span>
                    <div class="patient-meta">
                      <span class="badge badge-primary">{ref.type}</span>
                      <span>{ref.age} / {ref.gender}</span>
                    </div>
                  </div>
                </td>
                <td>{ref.diagnosis}</td>
                <td class="text-secondary font-title">{ref.to}</td>
                <td class="text-secondary">{ref.date}</td>
                <td>
                  {#if ref.status === 'Pending'}
                    <span class="badge badge-warning">Menunggu</span>
                  {:else if ref.status === 'Approved'}
                    <span class="badge badge-success">Diterima</span>
                  {:else}
                    <span class="badge badge-danger">Ditolak</span>
                  {/if}
                </td>
                <td class="actions-cell">
                  <button class="btn btn-secondary btn-sm" onclick={() => viewDetails(ref)}>
                    Detail
                  </button>
                </td>
              </tr>
            {/each}
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- Modal Form Buat Rujukan Baru -->
{#if showCreateModal}
  <div class="modal-backdrop" onclick={() => showCreateModal = false}>
    <div class="modal-content glass-panel animate-fade-in" onclick={(e) => e.stopPropagation()}>
      <div class="modal-header">
        <h2>Buat Rujukan Keluar Baru</h2>
        <button class="close-modal-btn" onclick={() => showCreateModal = false}>&times;</button>
      </div>

      <form onsubmit={handleCreateReferral}>
        <div class="modal-body">
          <div class="form-row">
            <div class="form-group flex-1">
              <label class="form-label" for="ptName">Nama Pasien</label>
              <input 
                type="text" 
                id="ptName" 
                class="form-input" 
                placeholder="Masukkan nama lengkap pasien" 
                bind:value={patientName}
                required
              />
            </div>
            <div class="form-group flex-1">
              <label class="form-label" for="ptNik">NIK Pasien</label>
              <input 
                type="text" 
                id="ptNik" 
                class="form-input" 
                placeholder="Nomor NIK 16 digit" 
                bind:value={patientNik}
                required
              />
            </div>
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label class="form-label" for="ptAge">Usia Pasien (Tahun)</label>
              <input 
                type="number" 
                id="ptAge" 
                class="form-input" 
                placeholder="Contoh: 35" 
                bind:value={patientAge}
                required
              />
            </div>
            <div class="form-group flex-1">
              <label class="form-label" for="ptGender">Jenis Kelamin</label>
              <select id="ptGender" class="form-input select-input" bind:value={patientGender}>
                <option value="Laki-laki">Laki-laki</option>
                <option value="Perempuan">Perempuan</option>
              </select>
            </div>
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label class="form-label" for="destFaskes">Faskes Tujuan Penerima</label>
              <input 
                type="text" 
                id="destFaskes" 
                class="form-input" 
                placeholder="Contoh: RSUP Dr. Wahidin Sudirohusodo" 
                bind:value={destinationFaskes}
                required
              />
            </div>
            <div class="form-group flex-1">
              <label class="form-label" for="diag">Diagnosis Utama (ICD-10)</label>
              <input 
                type="text" 
                id="diag" 
                class="form-input" 
                placeholder="Contoh: Hernia (K40.9)" 
                bind:value={diagnosis}
                required
              />
            </div>
          </div>

          <div class="form-group">
            <label class="form-label" for="type">Metode Pembayaran</label>
            <div class="radio-group">
              <label class="radio-label">
                <input type="radio" name="rujType" value="BPJS" bind:group={rujukanType}>
                <span>BPJS Kesehatan</span>
              </label>
              <label class="radio-label">
                <input type="radio" name="rujType" value="Umum" bind:group={rujukanType}>
                <span>Umum / Mandiri</span>
              </label>
            </div>
          </div>

          <div class="form-group">
            <label class="form-label" for="notes">Indikasi Medis & Catatan Klinis</label>
            <textarea 
              id="notes" 
              class="form-input textarea-input" 
              placeholder="Deskripsikan kondisi klinis pasien secara rinci..." 
              bind:value={clinicalNotes}
              rows="4"
              required
            ></textarea>
          </div>
        </div>

        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" onclick={() => showCreateModal = false}>Batal</button>
          <button type="submit" class="btn btn-primary">Kirim Surat Rujukan</button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- Detail Rujukan Keluar Modal -->
{#if showDetailModal && selectedReferral}
  <div class="modal-backdrop" onclick={() => showDetailModal = false}>
    <div class="modal-content glass-panel animate-fade-in" onclick={(e) => e.stopPropagation()}>
      <div class="modal-header">
        <div class="modal-title-area">
          <span class="modal-badge font-title">{selectedReferral.id}</span>
          <h2>Detail Rujukan Keluar</h2>
        </div>
        <button class="close-modal-btn" onclick={() => showDetailModal = false}>&times;</button>
      </div>

      <div class="modal-body">
        <div class="modal-grid">
          <!-- Patient Details -->
          <div class="info-section">
            <h3>Informasi Pasien</h3>
            <div class="info-grid">
              <div class="info-item">
                <span class="info-label">Nama Lengkap</span>
                <span class="info-val text-white font-title">{selectedReferral.patient}</span>
              </div>
              <div class="info-item">
                <span class="info-label">NIK Pasien</span>
                <span class="info-val text-white">{selectedReferral.nik}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Usia / Jenis Kelamin</span>
                <span class="info-val">{selectedReferral.age} / {selectedReferral.gender}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Jenis Rujukan</span>
                <span class="info-val"><span class="badge badge-primary">{selectedReferral.type}</span></span>
              </div>
            </div>
          </div>

          <!-- Referral Info -->
          <div class="info-section">
            <h3>Tujuan & Medis</h3>
            <div class="info-grid">
              <div class="info-item">
                <span class="info-label">Diagnosa</span>
                <span class="info-val text-white font-title">{selectedReferral.diagnosis}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Faskes Tujuan</span>
                <span class="info-val text-white">{selectedReferral.to}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Tanggal Pengajuan</span>
                <span class="info-val">{selectedReferral.date}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="notes-section">
          <span class="info-label">Catatan Medis</span>
          <div class="notes-box">{selectedReferral.notes}</div>
        </div>

        {#if selectedReferral.status === 'Rejected' && selectedReferral.rejectReason}
          <div class="rejection-box">
            <span class="info-label text-danger">Catatan Penolakan Faskes Penerima</span>
            <p>{selectedReferral.rejectReason}</p>
          </div>
        {/if}
      </div>

      <div class="modal-footer">
        <span class="status-indicator">
          Status Rujukan: 
          {#if selectedReferral.status === 'Pending'}
            <span class="badge badge-warning">Menunggu Konfirmasi</span>
          {:else if selectedReferral.status === 'Approved'}
            <span class="badge badge-success">Diterima / Disetujui</span>
          {:else}
            <span class="badge badge-danger">Ditolak</span>
          {/if}
        </span>
        <button class="btn btn-secondary" onclick={() => showDetailModal = false}>Tutup</button>
      </div>
    </div>
  </div>
{/if}

<style>
  .rujukan-keluar-container {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    position: relative;
  }

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1.5rem;
  }

  .page-header h1 {
    font-size: 2rem;
    color: white;
  }

  /* Toast Notification */
  .toast-notification {
    position: fixed;
    top: 24px;
    right: 24px;
    background: rgba(18, 27, 45, 0.9);
    border: 1px solid var(--color-success);
    padding: 1rem 1.5rem;
    border-radius: var(--border-radius-md);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    gap: 0.75rem;
    z-index: 2000;
    color: white;
    font-weight: 500;
    animation: slideInRight 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }

  @keyframes slideInRight {
    from { transform: translateX(120%); opacity: 0; }
    to { transform: translateX(0); opacity: 1; }
  }

  /* Filters styling */
  .filter-panel {
    padding: 1.5rem;
  }

  .search-box-wrap {
    position: relative;
    max-width: 500px;
    width: 100%;
  }

  .search-icon {
    position: absolute;
    left: 1rem;
    top: 50%;
    transform: translateY(-50%);
    width: 20px;
    height: 20px;
    color: var(--text-muted);
  }

  .search-input {
    width: 100%;
    padding-left: 3rem !important;
  }

  /* List Table Styling */
  .list-panel {
    padding: 1rem;
  }

  .rujukan-table {
    width: 100%;
    border-collapse: collapse;
    text-align: left;
  }

  .rujukan-table th {
    font-family: var(--font-title);
    font-size: 0.75rem;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: var(--text-muted);
    padding: 0.75rem 1rem;
    border-bottom: 1px solid var(--border-glass);
  }

  .rujukan-table td {
    padding: 1rem;
    font-size: 0.9rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.03);
    vertical-align: middle;
  }

  .table-row {
    transition: background-color var(--transition-fast);
  }

  .table-row:hover {
    background-color: rgba(255, 255, 255, 0.02);
  }

  .ref-id {
    color: var(--color-primary);
    font-weight: 700;
    font-size: 0.85rem;
  }

  .patient-info-cell {
    display: flex;
    flex-direction: column;
    gap: 0.35rem;
  }

  .patient-name {
    font-weight: 600;
    color: white;
  }

  .patient-meta {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-size: 0.75rem;
    color: var(--text-secondary);
  }

  .empty-cell {
    text-align: center;
    padding: 3rem 0;
    color: var(--text-muted);
  }

  .empty-cell p {
    margin-top: 0.75rem;
    font-size: 0.95rem;
  }

  .actions-header {
    text-align: right;
  }

  .actions-cell {
    text-align: right;
  }

  .btn-sm {
    padding: 0.4rem 0.85rem;
    font-size: 0.8rem;
    border-radius: 6px;
  }

  /* Modal Styles */
  .modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(11, 17, 30, 0.85);
    backdrop-filter: blur(8px);
    z-index: 1000;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1.5rem;
  }

  .modal-content {
    width: 100%;
    max-width: 800px;
    padding: 2.5rem;
    border-radius: var(--border-radius-lg);
    box-shadow: 0 25px 60px rgba(0, 0, 0, 0.5);
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    max-height: 90vh;
    overflow-y: auto;
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid var(--border-glass);
    padding-bottom: 1rem;
  }

  .modal-header h2 {
    font-size: 1.5rem;
    color: white;
  }

  .modal-title-area {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .modal-badge {
    background: rgba(0, 242, 254, 0.1);
    border: 1px solid rgba(0, 242, 254, 0.3);
    padding: 0.35rem 0.75rem;
    border-radius: 6px;
    color: var(--color-primary);
    font-size: 0.85rem;
    font-weight: 700;
  }

  .close-modal-btn {
    background: none;
    border: none;
    color: var(--text-secondary);
    font-size: 1.75rem;
    cursor: pointer;
    line-height: 1;
  }

  .close-modal-btn:hover {
    color: white;
  }

  .modal-body {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
  }

  .form-row {
    display: flex;
    gap: 1.25rem;
  }

  .flex-1 {
    flex: 1;
  }

  @media (max-width: 640px) {
    .form-row {
      flex-direction: column;
      gap: 1.25rem;
    }
    
    .page-header {
      flex-direction: column;
      align-items: flex-start;
      gap: 1rem;
    }
  }

  .select-input {
    background-color: var(--bg-secondary);
    cursor: pointer;
  }

  .select-input option {
    background-color: var(--bg-secondary);
    color: white;
  }

  /* Radio Group */
  .radio-group {
    display: flex;
    gap: 2rem;
    padding: 0.25rem 0;
  }

  .radio-label {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    cursor: pointer;
    font-size: 0.95rem;
    color: var(--text-secondary);
  }

  .radio-label input[type="radio"] {
    width: 18px;
    height: 18px;
    accent-color: var(--color-primary);
  }

  .textarea-input {
    resize: none;
    font-size: 0.95rem;
  }

  .modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 0.75rem;
    border-top: 1px solid var(--border-glass);
    padding-top: 1.25rem;
    margin-top: 1rem;
  }

  /* Detail Section inside modal */
  .modal-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1.5rem;
  }

  @media (max-width: 640px) {
    .modal-grid {
      grid-template-columns: 1fr;
    }
  }

  .info-section {
    background: rgba(255, 255, 255, 0.01);
    border: 1px solid rgba(255, 255, 255, 0.03);
    border-radius: var(--border-radius-md);
    padding: 1.25rem;
  }

  .info-section h3 {
    font-size: 0.95rem;
    color: var(--color-primary);
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
    padding-bottom: 0.5rem;
    margin-bottom: 1rem;
  }

  .info-grid {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .info-item {
    display: flex;
    flex-direction: column;
    gap: 0.15rem;
  }

  .info-label {
    font-size: 0.75rem;
    font-weight: 600;
    color: var(--text-muted);
    text-transform: uppercase;
    letter-spacing: 0.05em;
  }

  .info-val {
    font-size: 0.9rem;
    color: var(--text-secondary);
  }

  .text-white {
    color: white;
  }

  .notes-section {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .notes-box {
    background: rgba(0, 0, 0, 0.2);
    border: 1px solid var(--border-glass);
    border-radius: var(--border-radius-sm);
    padding: 1rem;
    font-size: 0.9rem;
    line-height: 1.6;
    color: var(--text-secondary);
  }

  .rejection-box {
    background: rgba(255, 117, 140, 0.05);
    border: 1px solid rgba(255, 117, 140, 0.15);
    border-radius: var(--border-radius-sm);
    padding: 1rem;
  }

  .rejection-box p {
    font-size: 0.9rem;
    color: #ff758c;
    margin-top: 0.35rem;
  }

  .status-indicator {
    align-self: center;
    margin-right: auto;
    font-size: 0.9rem;
    color: var(--text-secondary);
  }
</style>
