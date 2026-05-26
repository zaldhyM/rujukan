<script lang="ts">
  import { onMount } from 'svelte';
  import Fa from 'svelte-fa';
  import {
    faInbox,
    faPaperPlane,
    faClock,
    faCheckCircle,
    faExchangeAlt,
    faUserCheck,
    faChevronRight
  } from '@fortawesome/free-solid-svg-icons';
  import { getRujukanMasuk, getRujukanKeluar, getCurrentUser } from '$lib/auth';

  let user: any = $state({ name: '' });
  let incoming: any[] = $state([]);
  let outgoing: any[] = $state([]);

  let totalIncoming = $derived(incoming.length);
  let totalOutgoing = $derived(outgoing.length);
  let totalRujukan = $derived(totalIncoming + totalOutgoing);
  
  let pendingIncoming = $derived(incoming.filter((r: any) => r.status === 'Pending').length);
  let acceptedIncoming = $derived(incoming.filter((r: any) => r.status === 'Diterima').length);
  let processedOutgoing = $derived(outgoing.filter((r: any) => r.status === 'Diproses').length);

  // Combine and sort recent actions
  let recentActivities = $derived(() => {
    const combined = [
      ...incoming.map((r: any) => ({ ...r, type: 'Masuk' })),
      ...outgoing.map((r: any) => ({ ...r, type: 'Keluar' }))
    ];
    return combined.sort((a: any, b: any) => new Date(b.tglRujukan).getTime() - new Date(a.tglRujukan).getTime()).slice(0, 5);
  });

  onMount(() => {
    user = getCurrentUser();
    incoming = getRujukanMasuk();
    outgoing = getRujukanKeluar();
  });
</script>

<!-- Page Heading -->
<div class="d-sm-flex align-items-center justify-content-between mb-4">
  <div>
    <h1 class="h3 mb-1 text-gray-800 font-weight-bold">Dashboard Sisrute</h1>
    <p class="text-muted mb-0 small">Selamat datang kembali, <strong>{user.name}</strong>. Berikut ringkasan aktivitas rujukan medis hari ini.</p>
  </div>
</div>

<!-- Info Cards Row -->
<div class="row mb-4">
  <!-- Total Rujukan -->
  <div class="col-xl-3 col-md-6 mb-4 mb-xl-0">
    <div class="kpi-card kpi-primary h-100 py-3 px-4">
      <div class="d-flex align-items-center justify-content-between">
        <div>
          <span class="kpi-title">TOTAL RUJUKAN</span>
          <h2 class="kpi-value mb-0">{totalRujukan}</h2>
        </div>
        <div class="kpi-icon-wrapper">
          <Fa icon={faExchangeAlt} />
        </div>
      </div>
      <div class="kpi-footer mt-3 text-xs">
        <span class="text-primary font-weight-bold">{totalIncoming} Masuk</span> • <span class="text-info font-weight-bold">{totalOutgoing} Keluar</span>
      </div>
    </div>
  </div>

  <!-- Rujukan Masuk Pending -->
  <div class="col-xl-3 col-md-6 mb-4 mb-xl-0">
    <div class="kpi-card kpi-warning h-100 py-3 px-4">
      <div class="d-flex align-items-center justify-content-between">
        <div>
          <span class="kpi-title">PENDING MASUK</span>
          <h2 class="kpi-value mb-0">{pendingIncoming}</h2>
        </div>
        <div class="kpi-icon-wrapper">
          <Fa icon={faClock} />
        </div>
      </div>
      <div class="kpi-footer mt-3 text-xs">
        <span class="badge bg-warning-light text-warning font-weight-bold">Perlu Tindakan Segera</span>
      </div>
    </div>
  </div>

  <!-- Rujukan Masuk Diterima -->
  <div class="col-xl-3 col-md-6 mb-4 mb-sm-0">
    <div class="kpi-card kpi-success h-100 py-3 px-4">
      <div class="d-flex align-items-center justify-content-between">
        <div>
          <span class="kpi-title">DITERIMA MASUK</span>
          <h2 class="kpi-value mb-0">{acceptedIncoming}</h2>
        </div>
        <div class="kpi-icon-wrapper">
          <Fa icon={faCheckCircle} />
        </div>
      </div>
      <div class="kpi-footer mt-3 text-xs">
        <span class="badge bg-success-light text-success font-weight-bold">Telah Disetujui</span>
      </div>
    </div>
  </div>

  <!-- Rujukan Keluar Diproses -->
  <div class="col-xl-3 col-md-6">
    <div class="kpi-card kpi-info h-100 py-3 px-4">
      <div class="d-flex align-items-center justify-content-between">
        <div>
          <span class="kpi-title">DIPROSES KELUAR</span>
          <h2 class="kpi-value mb-0">{processedOutgoing}</h2>
        </div>
        <div class="kpi-icon-wrapper">
          <Fa icon={faPaperPlane} />
        </div>
      </div>
      <div class="kpi-footer mt-3 text-xs">
        <span class="badge bg-info-light text-info font-weight-bold">Menunggu Hasil Rumah Sakit</span>
      </div>
    </div>
  </div>
</div>

<div class="row">
  <!-- Recent Referrals -->
  <div class="col-lg-8 mb-4 mb-lg-0">
    <div class="card-modern">
      <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between bg-white border-bottom">
        <h6 class="m-0 font-weight-bold text-primary">Aktivitas Rujukan Terkini</h6>
        <span class="badge bg-light text-muted font-weight-bold">{recentActivities().length} Rujukan Terbaru</span>
      </div>
      <div class="card-body p-0">
        {#if recentActivities().length === 0}
          <div class="text-center py-5">
            <p class="text-muted mb-0">Belum ada aktivitas rujukan baru hari ini.</p>
          </div>
        {:else}
          <div class="table-responsive">
            <table class="table-modern">
              <thead>
                <tr>
                  <th>No. Rujukan</th>
                  <th>Tanggal</th>
                  <th>Jenis</th>
                  <th>Pasien</th>
                  <th>Faskes Relasi</th>
                  <th class="text-center">Status</th>
                </tr>
              </thead>
              <tbody>
                {#each recentActivities() as r (r.id)}
                  <tr>
                    <td class="font-weight-bold text-primary">{r.noRujukan}</td>
                    <td>{r.tglRujukan}</td>
                    <td>
                      <span class="badge-modern" class:badge-modern-info={r.type === 'Masuk'} class:badge-modern-success={r.type === 'Keluar'}>
                        {r.type === 'Masuk' ? 'Rujukan Masuk' : 'Rujukan Keluar'}
                      </span>
                    </td>
                    <td>
                      <div class="font-weight-bold">{r.pasien.nama}</div>
                      <div class="text-xs text-muted">BPJS: {r.pasien.noKartu}</div>
                    </td>
                    <td>{r.type === 'Masuk' ? r.faskesAsal : r.faskesTujuan}</td>
                    <td class="text-center">
                      <span class="badge-modern" 
                        class:badge-modern-warning={r.status === 'Pending' || r.status === 'Diproses'}
                        class:badge-modern-success={r.status === 'Diterima' || r.status === 'Selesai' || r.status === 'Dikirim'}
                        class:badge-modern-danger={r.status === 'Ditolak'}
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
  </div>

  <!-- Quick Actions & Stats -->
  <div class="col-lg-4">
    <!-- Quick Access Card -->
    <div class="card-modern mb-4">
      <div class="card-header py-3 bg-white border-bottom">
        <h6 class="m-0 font-weight-bold text-primary">Akses Cepat Fitur</h6>
      </div>
      <div class="card-body py-3">
        <a href="/admin/rujukan-masuk" class="quick-link-btn mb-3 p-3">
          <div class="d-flex align-items-center gap-3">
            <div class="quick-link-icon ql-warning">
              <Fa icon={faInbox} />
            </div>
            <div class="quick-link-info">
              <strong class="d-block text-main">Kelola Rujukan Masuk</strong>
              <span class="text-xs text-muted">Review dan proses pengajuan rujukan dari faskes lain.</span>
            </div>
          </div>
          <Fa icon={faChevronRight} class="quick-link-arrow" />
        </a>
        
        <a href="/admin/rujukan-keluar" class="quick-link-btn p-3">
          <div class="d-flex align-items-center gap-3">
            <div class="quick-link-icon ql-primary">
              <Fa icon={faPaperPlane} />
            </div>
            <div class="quick-link-info">
              <strong class="d-block text-main">Buat Rujukan Keluar</strong>
              <span class="text-xs text-muted">Kirim data rujukan pasien ke rumah sakit tujuan.</span>
            </div>
          </div>
          <Fa icon={faChevronRight} class="quick-link-arrow" />
        </a>
      </div>
    </div>

    <!-- Info Box -->
    <div class="profile-banner-card p-4 text-white shadow-sm position-relative overflow-hidden">
      <div class="ambient-glow" style="top: -20%; right: -20%; background: radial-gradient(circle, rgba(255,255,255,0.15) 0%, rgba(255,255,255,0) 70%);"></div>
      <div class="position-relative z-1">
        <div class="profile-banner-icon mb-3">
          <Fa icon={faUserCheck} />
        </div>
        <h5 class="font-weight-bold mb-2">Informasi Akun</h5>
        <p class="small opacity-75 mb-0 leading-relaxed">
          Anda masuk sebagai <strong>{user.name}</strong> dengan hak akses penuh ke sistem rujukan elektronik rumah sakit/puskesmas.
        </p>
      </div>
    </div>
  </div>
</div>

<style>
  /* KPI Custom Cards */
  .kpi-card {
    background-color: var(--bg-card);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    box-shadow: var(--shadow-sm);
    transition: var(--transition);
    position: relative;
    overflow: hidden;
  }
  .kpi-card::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 4px;
    height: 100%;
  }
  .kpi-card:hover {
    transform: translateY(-3px);
    box-shadow: var(--shadow-md);
  }
  
  .kpi-primary::after { background-color: var(--primary); }
  .kpi-warning::after { background-color: var(--warning); }
  .kpi-success::after { background-color: var(--success); }
  .kpi-info::after { background-color: var(--info); }
  
  .kpi-title {
    font-size: 0.65rem;
    font-weight: 700;
    color: var(--text-muted);
    letter-spacing: 0.1em;
    display: block;
    margin-bottom: 0.25rem;
  }
  
  .kpi-value {
    font-weight: 800;
    color: var(--text-main);
    font-size: 1.85rem;
  }
  
  .kpi-icon-wrapper {
    width: 46px;
    height: 46px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 1.25rem;
  }
  .kpi-primary .kpi-icon-wrapper { background-color: var(--primary-light); color: var(--primary); }
  .kpi-warning .kpi-icon-wrapper { background-color: var(--warning-light); color: var(--warning); }
  .kpi-success .kpi-icon-wrapper { background-color: var(--success-light); color: var(--success); }
  .kpi-info .kpi-icon-wrapper { background-color: var(--info-light); color: var(--info); }
  
  .bg-warning-light { background-color: var(--warning-light) !important; }
  .bg-success-light { background-color: var(--success-light) !important; }
  .bg-info-light { background-color: var(--info-light) !important; }
  
  .kpi-footer {
    border-top: 1px solid var(--border-color);
    padding-top: 0.75rem;
    color: var(--text-muted);
  }

  /* Quick Actions Styling */
  .quick-link-btn {
    display: flex;
    align-items: center;
    justify-content: space-between;
    text-decoration: none;
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    transition: var(--transition);
  }
  .quick-link-btn:hover {
    transform: translateY(-2px);
    border-color: var(--primary);
    box-shadow: var(--shadow-sm);
  }
  .quick-link-icon {
    width: 44px;
    height: 44px;
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 1.15rem;
    flex-shrink: 0;
  }
  .ql-primary { background-color: var(--primary-light); color: var(--primary); }
  .ql-warning { background-color: var(--warning-light); color: var(--warning); }
  
  .quick-link-info {
    line-height: 1.3;
  }
  .quick-link-arrow {
    color: var(--text-light);
    transition: var(--transition);
  }
  .quick-link-btn:hover .quick-link-arrow {
    color: var(--primary);
    transform: translateX(3px);
  }
  
  /* Profile Banner Card */
  .profile-banner-card {
    background: linear-gradient(135deg, #1e1b4b 0%, #312e81 100%);
    border-radius: var(--radius-md);
  }
  .profile-banner-icon {
    width: 46px;
    height: 46px;
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 1.25rem;
  }
</style>
