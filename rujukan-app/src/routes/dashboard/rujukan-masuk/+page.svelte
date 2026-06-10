<!-- src/routes/dashboard/rujukan-masuk/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';

  // Svelte 5 state
  let searchQuery = $state('');
  let activeTab = $state('All'); // All, Pending, Approved, Rejected
  let selectedReferral = $state<any>(null);
  let showDetailModal = $state(false);
  let showRejectReasonModal = $state(false);
  let rejectionReasonText = $state('');
  let referralToReject = $state<any>(null);

  let referrals = $state([
    { 
      id: 'REF-008', 
      patient: 'Andi Wijaya', 
      from: 'Puskesmas Kassi Kassi', 
      date: 'Hari ini, 09:15', 
      status: 'Pending', 
      type: 'Umum', 
      nik: '7371100203990001', 
      gender: 'Laki-laki', 
      age: '34 Tahun', 
      diagnosis: 'Appendicitis Acute (K35.8)', 
      senderDoctor: 'dr. Sarah Wijaya', 
      notes: 'Pasien mengalami nyeri perut kanan bawah sejak 2 hari yang lalu. Skala nyeri 7/10. Hasil lab leukositosis 15.000. Mohon tindakan operasi appendectomy segera.' 
    },
    { 
      id: 'REF-007', 
      patient: 'Siti Rahma', 
      from: 'RS Stella Maris', 
      date: 'Hari ini, 08:30', 
      status: 'Approved', 
      type: 'BPJS', 
      nik: '7371084205880003', 
      gender: 'Perempuan', 
      age: '52 Tahun', 
      diagnosis: 'Congestive Heart Failure (I50.9)', 
      senderDoctor: 'dr. Faisal Rahman, Sp.JP', 
      notes: 'Pasien sesak napas saat istirahat, edema pada tungkai bawah bilateral. Hasil rontgen kardiomegali dengan edema paru. Butuh ICU/ICCU segera.' 
    },
    { 
      id: 'REF-006', 
      patient: 'Budi Santoso', 
      from: 'Klinik Medika', 
      date: 'Kemarin, 16:45', 
      status: 'Approved', 
      type: 'BPJS', 
      nik: '7371121112750002', 
      gender: 'Laki-laki', 
      age: '48 Tahun', 
      diagnosis: 'Diabetes Mellitus Type II (E11)', 
      senderDoctor: 'dr. Herman Kurnia', 
      notes: 'GDS terakhir 450 mg/dL dengan ulkus diabetikum derajat III pada pedis dextra. Perlu debridement jaringan mati dan regulasi insulin.' 
    },
    { 
      id: 'REF-005', 
      patient: 'Dewi Lestari', 
      from: 'Puskesmas Mamajang', 
      date: 'Kemarin, 14:20', 
      status: 'Rejected', 
      type: 'Umum', 
      nik: '7371092809930005', 
      gender: 'Perempuan', 
      age: '28 Tahun', 
      diagnosis: 'Gastritis Chronis (K29.7)', 
      senderDoctor: 'dr. Indah Sari', 
      notes: 'Nyeri epigastrium berat disertai mual muntah berulang. Terapi standar rawat jalan kurang responsif. Rujuk ke poli penyakit dalam.', 
      rejectReason: 'Kapasitas poli penyakit dalam penuh untuk 3 hari ke depan. Disarankan rujuk ke RS tipe C terdekat.' 
    },
    { 
      id: 'REF-004', 
      patient: 'Rudi Hermawan', 
      from: 'Puskesmas Bara-Baraya', 
      date: '08 Juni 2026', 
      status: 'Pending', 
      type: 'BPJS', 
      nik: '7371021404670001', 
      gender: 'Laki-laki', 
      age: '59 Tahun', 
      diagnosis: 'Cataract Senilis (H25)', 
      senderDoctor: 'dr. Rian Hidayat', 
      notes: 'Visus VOD 1/60, VOS 6/60. Kekeruhan lensa grade IV. Rujuk untuk tindakan bedah phacoemulsification.' 
    }
  ]);

  // Derived filtered referrals
  let filteredReferrals = $derived(
    referrals.filter(ref => {
      const matchesSearch = 
        ref.patient.toLowerCase().includes(searchQuery.toLowerCase()) || 
        ref.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
        ref.from.toLowerCase().includes(searchQuery.toLowerCase()) ||
        ref.diagnosis.toLowerCase().includes(searchQuery.toLowerCase());
      
      const matchesTab = activeTab === 'All' || ref.status === activeTab;
      return matchesSearch && matchesTab;
    })
  );

  const viewDetails = (ref: any) => {
    selectedReferral = ref;
    showDetailModal = true;
  };

  const approveReferral = (id: string) => {
    referrals = referrals.map(ref => {
      if (ref.id === id) {
        return { ...ref, status: 'Approved' };
      }
      return ref;
    });
    
    // Update active selected object if modal is open
    if (selectedReferral && selectedReferral.id === id) {
      selectedReferral = { ...selectedReferral, status: 'Approved' };
    }
  };

  const openRejectModal = (ref: any) => {
    referralToReject = ref;
    rejectionReasonText = '';
    showRejectReasonModal = true;
  };

  const submitRejection = () => {
    if (!rejectionReasonText.trim()) return;

    referrals = referrals.map(ref => {
      if (ref.id === referralToReject.id) {
        return { ...ref, status: 'Rejected', rejectReason: rejectionReasonText };
      }
      return ref;
    });

    if (selectedReferral && selectedReferral.id === referralToReject.id) {
      selectedReferral = { ...selectedReferral, status: 'Rejected', rejectReason: rejectionReasonText };
    }

    showRejectReasonModal = false;
    referralToReject = null;
  };
</script>

<svelte:head>
  <title>Rujukan Masuk | RUJUKAN MEDIS</title>
</svelte:head>

<div class="rujukan-masuk-container animate-fade-in">
  <div class="page-header">
    <div class="header-left">
      <span class="breadcrumb">DASHBOARD / RUJUKAN MASUK</span>
      <h1>Rujukan Masuk</h1>
      <p class="subtitle">Kelola dan review rujukan pasien dari fasilitas kesehatan tingkat pertama (FKTP) atau RS lain</p>
    </div>
  </div>

  <!-- Search & Filters Panel -->
  <div class="glass-panel filter-panel">
    <div class="search-box-wrap">
      <svg class="search-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="11" cy="11" r="8"></circle>
        <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
      </svg>
      <input 
        type="text" 
        class="form-input search-input" 
        placeholder="Cari berdasarkan nama pasien, ID rujukan, diagnosis, atau faskes..." 
        bind:value={searchQuery}
      />
    </div>

    <div class="filter-tabs">
      <button 
        class="tab-btn" 
        class:active={activeTab === 'All'} 
        onclick={() => activeTab = 'All'}
      >
        Semua
      </button>
      <button 
        class="tab-btn" 
        class:active={activeTab === 'Pending'} 
        onclick={() => activeTab = 'Pending'}
      >
        Menunggu Review
        <span class="tab-badge bg-warning">{referrals.filter(r => r.status === 'Pending').length}</span>
      </button>
      <button 
        class="tab-btn" 
        class:active={activeTab === 'Approved'} 
        onclick={() => activeTab = 'Approved'}
      >
        Disetujui
      </button>
      <button 
        class="tab-btn" 
        class:active={activeTab === 'Rejected'} 
        onclick={() => activeTab = 'Rejected'}
      >
        Ditolak
      </button>
    </div>
  </div>

  <!-- Table List -->
  <div class="glass-panel list-panel">
    <div class="table-responsive">
      <table class="rujukan-table">
        <thead>
          <tr>
            <th>ID Rujukan</th>
            <th>Nama Pasien</th>
            <th>Diagnosa / Gejala</th>
            <th>Faskes Pengirim</th>
            <th>Tgl Masuk</th>
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
                <p>Tidak ditemukan data rujukan masuk.</p>
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
                <td>
                  <div class="diagnosis-cell">
                    <span class="diag-text">{ref.diagnosis}</span>
                    <span class="doctor-text">Dokter: {ref.senderDoctor}</span>
                  </div>
                </td>
                <td class="text-secondary">{ref.from}</td>
                <td class="text-secondary">{ref.date}</td>
                <td>
                  {#if ref.status === 'Pending'}
                    <span class="badge badge-warning">Review</span>
                  {:else if ref.status === 'Approved'}
                    <span class="badge badge-success">Disetujui</span>
                  {:else}
                    <span class="badge badge-danger">Ditolak</span>
                  {/if}
                </td>
                <td class="actions-cell">
                  <div class="btn-group">
                    <button class="btn btn-secondary btn-sm" onclick={() => viewDetails(ref)}>
                      Detail
                    </button>
                    {#if ref.status === 'Pending'}
                      <button class="btn btn-success btn-sm" onclick={() => approveReferral(ref.id)}>
                        Terima
                      </button>
                      <button class="btn btn-danger btn-sm" onclick={() => openRejectModal(ref)}>
                        Tolak
                      </button>
                    {/if}
                  </div>
                </td>
              </tr>
            {/each}
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- Referral Detail Modal -->
{#if showDetailModal && selectedReferral}
  <div class="modal-backdrop" onclick={() => showDetailModal = false}>
    <div class="modal-content glass-panel animate-fade-in" onclick={(e) => e.stopPropagation()}>
      <div class="modal-header">
        <div class="modal-title-area">
          <span class="modal-badge font-title">{selectedReferral.id}</span>
          <h2>Detail Rujukan Masuk</h2>
        </div>
        <button class="close-modal-btn" onclick={() => showDetailModal = false}>&times;</button>
      </div>

      <div class="modal-body">
        <div class="modal-grid">
          <!-- Patient Details Section -->
          <div class="info-section">
            <h3>Informasi Pasien</h3>
            <div class="info-grid">
              <div class="info-item">
                <span class="info-label">Nama Lengkap</span>
                <span class="info-val text-white font-title">{selectedReferral.patient}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Nomor NIK</span>
                <span class="info-val text-white">{selectedReferral.nik}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Jenis Kelamin / Usia</span>
                <span class="info-val">{selectedReferral.gender} / {selectedReferral.age}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Metode Pembayaran</span>
                <span class="info-val"><span class="badge badge-primary">{selectedReferral.type}</span></span>
              </div>
            </div>
          </div>

          <!-- Medical Referral Details Section -->
          <div class="info-section">
            <h3>Informasi Medis & Rujukan</h3>
            <div class="info-grid">
              <div class="info-item">
                <span class="info-label">Diagnosa Utama (ICD-10)</span>
                <span class="info-val text-white font-title">{selectedReferral.diagnosis}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Asal Fasilitas Kesehatan</span>
                <span class="info-val text-white">{selectedReferral.from}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Dokter Pengirim</span>
                <span class="info-val">{selectedReferral.senderDoctor}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Tanggal Dikirim</span>
                <span class="info-val">{selectedReferral.date}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="notes-section">
          <span class="info-label">Catatan Medis & Keterangan Klinis</span>
          <div class="notes-box">{selectedReferral.notes}</div>
        </div>

        {#if selectedReferral.status === 'Rejected' && selectedReferral.rejectReason}
          <div class="rejection-box">
            <span class="info-label text-danger">Alasan Penolakan Rujukan</span>
            <p>{selectedReferral.rejectReason}</p>
          </div>
        {/if}
      </div>

      <div class="modal-footer">
        <div class="footer-left">
          <span class="status-indicator">
            Status: 
            {#if selectedReferral.status === 'Pending'}
              <span class="badge badge-warning">Menunggu Keputusan</span>
            {:else if selectedReferral.status === 'Approved'}
              <span class="badge badge-success">Telah Disetujui</span>
            {:else}
              <span class="badge badge-danger">Telah Ditolak</span>
            {/if}
          </span>
        </div>
        <div class="footer-right">
          <button class="btn btn-secondary" onclick={() => showDetailModal = false}>Tutup</button>
          {#if selectedReferral.status === 'Pending'}
            <button class="btn btn-danger" onclick={() => { openRejectModal(selectedReferral); showDetailModal = false; }}>Tolak Rujukan</button>
            <button class="btn btn-success" onclick={() => { approveReferral(selectedReferral.id); showDetailModal = false; }}>Terima Rujukan</button>
          {/if}
        </div>
      </div>
    </div>
  </div>
{/if}

<!-- Rejection Reason Modal -->
{#if showRejectReasonModal && referralToReject}
  <div class="modal-backdrop" onclick={() => showRejectReasonModal = false}>
    <div class="modal-content sub-modal glass-panel animate-fade-in" onclick={(e) => e.stopPropagation()}>
      <div class="modal-header">
        <h2>Alasan Penolakan Rujukan</h2>
        <button class="close-modal-btn" onclick={() => showRejectReasonModal = false}>&times;</button>
      </div>
      <div class="modal-body">
        <p class="text-secondary mb-3">
          Tentukan alasan mengapa rujukan untuk pasien <strong>{referralToReject.patient}</strong> ditolak.
        </p>
        <div class="form-group">
          <label class="form-label" for="rejectReason">Alasan Penolakan</label>
          <textarea 
            id="rejectReason" 
            class="form-input textarea-input" 
            placeholder="Tulis alasan logis penolakan (misal: Ruang ICU Penuh, Alat rusak, dll.)" 
            bind:value={rejectionReasonText}
            rows="4"
          ></textarea>
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-secondary" onclick={() => showRejectReasonModal = false}>Batal</button>
        <button 
          class="btn btn-danger" 
          disabled={!rejectionReasonText.trim()} 
          onclick={submitRejection}
        >
          Konfirmasi Tolak
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  .rujukan-masuk-container {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .page-header h1 {
    font-size: 2rem;
    color: white;
  }

  /* Filter Panel styling */
  .filter-panel {
    padding: 1.5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1.5rem;
    flex-wrap: wrap;
  }

  .search-box-wrap {
    position: relative;
    flex-grow: 1;
    max-width: 500px;
    min-width: 280px;
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

  .filter-tabs {
    display: flex;
    gap: 0.5rem;
  }

  .tab-btn {
    background: none;
    border: 1px solid var(--border-glass);
    color: var(--text-secondary);
    padding: 0.5rem 1.25rem;
    border-radius: var(--border-radius-sm);
    cursor: pointer;
    font-family: var(--font-title);
    font-weight: 600;
    transition: all var(--transition-fast);
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .tab-btn:hover {
    color: white;
    background: rgba(255, 255, 255, 0.03);
  }

  .tab-btn.active {
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-primary-hover) 100%);
    color: #0b111e;
    border-color: var(--color-primary);
  }

  .tab-badge {
    padding: 0.1rem 0.4rem;
    border-radius: 4px;
    font-size: 0.7rem;
    font-weight: 700;
    color: #0b111e;
  }

  .bg-warning {
    background-color: var(--color-warning);
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

  .diagnosis-cell {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .diag-text {
    font-weight: 500;
    color: white;
  }

  .doctor-text {
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

  .btn-group {
    display: inline-flex;
    gap: 0.5rem;
  }

  .btn-sm {
    padding: 0.4rem 0.85rem;
    font-size: 0.8rem;
    border-radius: 6px;
  }

  /* Modal Backdrop & Content */
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

  .sub-modal {
    max-width: 500px;
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid var(--border-glass);
    padding-bottom: 1rem;
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

  .modal-header h2 {
    font-size: 1.5rem;
    color: white;
  }

  .close-modal-btn {
    background: none;
    border: none;
    color: var(--text-secondary);
    font-size: 1.75rem;
    cursor: pointer;
    line-height: 1;
    transition: color var(--transition-fast);
  }

  .close-modal-btn:hover {
    color: white;
  }

  .modal-body {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .modal-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1.5rem;
  }

  @media (max-width: 640px) {
    .modal-grid {
      grid-template-columns: 1fr;
    }
    .filter-panel {
      flex-direction: column;
      align-items: stretch;
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

  .modal-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-top: 1px solid var(--border-glass);
    padding-top: 1.25rem;
    margin-top: 1rem;
  }

  .footer-right {
    display: flex;
    gap: 0.75rem;
  }

  .status-indicator {
    font-size: 0.9rem;
    color: var(--text-secondary);
  }

  .textarea-input {
    resize: none;
    font-size: 0.95rem;
  }

  .mb-3 {
    margin-bottom: 0.75rem;
  }
</style>
