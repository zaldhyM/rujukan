<!-- src/routes/dashboard/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';

  let faskes = $state('');

  onMount(() => {
    faskes = localStorage.getItem('faskes') || 'Fasilitas Kesehatan';
  });

  // Mock Data
  const stats = [
    { title: 'Total Rujukan Masuk', value: '24', change: '+12% minggu ini', type: 'primary', icon: 'M19 10.5V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V10.5M12 3V17M12 3L7.5 7.5M12 3L16.5 7.5' },
    { title: 'Rujukan Disetujui', value: '18', change: '80% Rate Penerimaan', type: 'success', icon: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z' },
    { title: 'Menunggu Review', value: '4', change: 'Butuh tindakan segera', type: 'warning', icon: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z' },
    { title: 'Rujukan Ditolak', value: '2', change: 'Penolakan dengan catatan', type: 'danger', icon: 'M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z' }
  ];

  const recentReferrals = [
    { id: 'REF-008', patient: 'Andi Wijaya', from: 'Puskesmas Kassi Kassi', date: 'Hari ini, 09:15', status: 'Pending', type: 'Umum' },
    { id: 'REF-007', patient: 'Siti Rahma', from: 'RS Stella Maris', date: 'Hari ini, 08:30', status: 'Approved', type: 'BPJS' },
    { id: 'REF-006', patient: 'Budi Santoso', from: 'Klinik Medika', date: 'Kemarin, 16:45', status: 'Approved', type: 'BPJS' },
    { id: 'REF-005', patient: 'Dewi Lestari', from: 'Puskesmas Mamajang', date: 'Kemarin, 14:20', status: 'Rejected', type: 'Umum' }
  ];
</script>

<div class="overview-container animate-fade-in">
  <div class="page-header">
    <div class="header-left">
      <span class="breadcrumb">BERANDA / DASHBOARD</span>
      <h1>Ringkasan Rujukan</h1>
      <p class="subtitle">Monitor status rujukan pasien secara real-time untuk {faskes}</p>
    </div>
    <div class="header-right">
      <div class="date-badge glass-panel">
        <svg viewBox="0 0 24 24" fill="none" width="18" height="18" stroke="currentColor" stroke-width="2">
          <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
          <line x1="16" y1="2" x2="16" y2="6"></line>
          <line x1="8" y1="2" x2="8" y2="6"></line>
          <line x1="3" y1="10" x2="21" y2="10"></line>
        </svg>
        <span>10 Juni 2026</span>
      </div>
    </div>
  </div>

  <!-- KPI Grid -->
  <div class="stats-grid">
    {#each stats as stat}
      <div class="glass-panel stat-card card-glow-{stat.type}">
        <div class="stat-header">
          <span class="stat-title">{stat.title}</span>
          <div class="stat-icon-wrap icon-{stat.type}">
            <svg viewBox="0 0 24 24" fill="none" width="20" height="20" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <path d={stat.icon} />
            </svg>
          </div>
        </div>
        <div class="stat-value">{stat.value}</div>
        <div class="stat-footer">
          <span class="stat-change text-{stat.type}">{stat.change}</span>
        </div>
      </div>
    {/each}
  </div>

  <!-- Visual Charts & Recent Table -->
  <div class="dashboard-grid">
    <!-- Chart Panel -->
    <div class="glass-panel chart-panel">
      <div class="panel-header">
        <h3>Tren Rujukan Medis</h3>
        <span class="badge badge-primary">Mingguan</span>
      </div>
      <div class="chart-container">
        <!-- Premium Custom SVG Bar/Line Chart -->
        <svg class="custom-chart" viewBox="0 0 500 220" width="100%" height="100%">
          <defs>
            <linearGradient id="chartGrad" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stop-color="rgba(0, 200, 83, 0.4)"/>
              <stop offset="100%" stop-color="rgba(0, 200, 83, 0)"/>
            </linearGradient>
            <linearGradient id="lineGrad" x1="0" y1="0" x2="1" y2="0">
              <stop offset="0%" stop-color="#00C853"/>
              <stop offset="100%" stop-color="#8DC63F"/>
            </linearGradient>
          </defs>
          
          <!-- Grid Lines -->
          <line x1="30" y1="20" x2="480" y2="20" stroke="rgba(255, 255, 255, 0.05)" stroke-width="1" />
          <line x1="30" y1="70" x2="480" y2="70" stroke="rgba(255, 255, 255, 0.05)" stroke-width="1" />
          <line x1="30" y1="120" x2="480" y2="120" stroke="rgba(255, 255, 255, 0.05)" stroke-width="1" />
          <line x1="30" y1="170" x2="480" y2="170" stroke="rgba(255, 255, 255, 0.05)" stroke-width="1" />
          
          <!-- Chart Paths -->
          <path d="M 40 160 Q 110 90 180 130 T 320 60 T 460 30 L 460 170 L 40 170 Z" fill="url(#chartGrad)" />
          <path d="M 40 160 Q 110 90 180 130 T 320 60 T 460 30" fill="none" stroke="url(#lineGrad)" stroke-width="3" stroke-linecap="round" />
          
          <!-- Chart Data Points -->
          <circle cx="40" cy="160" r="4" fill="#00C853" stroke="#071a0d" stroke-width="2" />
          <circle cx="180" cy="130" r="4" fill="#00C853" stroke="#071a0d" stroke-width="2" />
          <circle cx="320" cy="60" r="4" fill="#8DC63F" stroke="#071a0d" stroke-width="2" />
          <circle cx="460" cy="30" r="4" fill="#8DC63F" stroke="#071a0d" stroke-width="2" />
          
          <!-- Labels -->
          <text x="40" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Sen</text>
          <text x="110" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Sel</text>
          <text x="180" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Rab</text>
          <text x="250" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Kam</text>
          <text x="320" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Jum</text>
          <text x="390" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Sab</text>
          <text x="460" y="195" fill="var(--text-muted)" font-size="10" text-anchor="middle">Min</text>

          <text x="22" y="174" fill="var(--text-muted)" font-size="10" text-anchor="end">0</text>
          <text x="22" y="124" fill="var(--text-muted)" font-size="10" text-anchor="end">10</text>
          <text x="22" y="74" fill="var(--text-muted)" font-size="10" text-anchor="end">20</text>
          <text x="22" y="24" fill="var(--text-muted)" font-size="10" text-anchor="end">30</text>
        </svg>
      </div>
    </div>

    <!-- Table Panel -->
    <div class="glass-panel table-panel">
      <div class="panel-header">
        <h3>Aktivitas Rujukan Terkini</h3>
        <a href="/dashboard/rujukan-masuk" class="view-all-link">Lihat Semua</a>
      </div>
      <div class="table-responsive">
        <table class="dashboard-table">
          <thead>
            <tr>
              <th>ID Rujukan</th>
              <th>Nama Pasien</th>
              <th>Pengirim</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {#each recentReferrals as ref}
              <tr class="table-row">
                <td class="ref-id font-title">{ref.id}</td>
                <td>
                  <div class="patient-cell">
                    <span class="patient-name">{ref.patient}</span>
                    <span class="patient-type badge badge-primary">{ref.type}</span>
                  </div>
                </td>
                <td class="text-secondary">{ref.from}</td>
                <td>
                  {#if ref.status === 'Pending'}
                    <span class="badge badge-warning">Review</span>
                  {:else if ref.status === 'Approved'}
                    <span class="badge badge-success">Disetujui</span>
                  {:else}
                    <span class="badge badge-danger">Ditolak</span>
                  {/if}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>

<style>
  .overview-container {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1.5rem;
  }

  .breadcrumb {
    font-size: 0.75rem;
    font-weight: 600;
    color: var(--color-primary);
    letter-spacing: 0.1em;
  }

  .page-header h1 {
    font-size: 2rem;
    color: var(--text-primary);
    margin-top: 0.25rem;
  }

  .date-badge {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.75rem 1.25rem;
    border-radius: var(--border-radius-sm);
    color: var(--text-primary);
    font-family: var(--font-title);
    font-weight: 600;
  }

  /* Stats Cards Grid */
  .stats-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1.5rem;
  }

  @media (max-width: 1200px) {
    .stats-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 640px) {
    .stats-grid {
      grid-template-columns: 1fr;
    }
    
    .page-header {
      flex-direction: column;
      align-items: flex-start;
    }
  }

  .stat-card {
    padding: 1.5rem;
    position: relative;
    overflow: hidden;
  }

  /* Glow effects per card status type */
  .card-glow-primary:hover {
    border-color: var(--color-primary);
    box-shadow: 0 8px 30px rgba(0, 200, 83, 0.15);
  }
  .card-glow-success:hover {
    border-color: var(--color-success);
    box-shadow: 0 8px 30px rgba(0, 200, 83, 0.15);
  }
  .card-glow-warning:hover {
    border-color: var(--color-warning);
    box-shadow: 0 8px 30px rgba(255, 174, 52, 0.15);
  }
  .card-glow-danger:hover {
    border-color: var(--color-danger);
    box-shadow: 0 8px 30px rgba(255, 117, 140, 0.15);
  }

  .stat-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
  }

  .stat-title {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--text-secondary);
    font-family: var(--font-title);
    text-transform: uppercase;
    letter-spacing: 0.05em;
  }

  .stat-icon-wrap {
    width: 36px;
    height: 36px;
    border-radius: var(--border-radius-sm);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .icon-primary { background: rgba(0, 200, 83, 0.1); color: var(--color-primary); }
  .icon-success { background: rgba(0, 200, 83, 0.1); color: var(--color-success); }
  .icon-warning { background: rgba(255, 174, 52, 0.1); color: var(--color-warning); }
  .icon-danger { background: rgba(255, 117, 140, 0.1); color: var(--color-danger); }

  .stat-value {
    font-size: 2.25rem;
    font-weight: 800;
    font-family: var(--font-title);
    color: var(--text-primary);
    line-height: 1;
    margin-bottom: 0.5rem;
  }

  .stat-change {
    font-size: 0.8rem;
    font-weight: 500;
  }
  
  .text-primary { color: var(--color-primary); }
  .text-success { color: var(--color-success); }
  .text-warning { color: var(--color-warning); }
  .text-danger { color: var(--color-danger); }

  /* Dashboard Panels Grid */
  .dashboard-grid {
    display: grid;
    grid-template-columns: 3fr 2fr;
    gap: 1.5rem;
  }

  @media (max-width: 1024px) {
    .dashboard-grid {
      grid-template-columns: 1fr;
    }
  }

  .chart-panel, .table-panel {
    padding: 1.75rem;
    display: flex;
    flex-direction: column;
  }

  .panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
  }

  .panel-header h3 {
    font-size: 1.2rem;
    color: var(--text-primary);
  }

  .view-all-link {
    font-size: 0.85rem;
    color: var(--color-primary);
    font-weight: 600;
    transition: color var(--transition-fast);
  }

  .view-all-link:hover {
    color: var(--color-primary-hover);
    text-decoration: underline;
  }

  .chart-container {
    height: 220px;
    position: relative;
  }

  .custom-chart {
    width: 100%;
    height: 100%;
  }

  /* Table styling inside dashboard overview */
  .table-responsive {
    overflow-x: auto;
    width: 100%;
  }

  .dashboard-table {
    width: 100%;
    border-collapse: collapse;
    text-align: left;
  }

  .dashboard-table th {
    font-family: var(--font-title);
    font-size: 0.75rem;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: var(--text-muted);
    padding: 0.75rem 1rem;
    border-bottom: 1px solid var(--border-glass);
  }

  .dashboard-table td {
    padding: 1rem;
    font-size: 0.9rem;
    border-bottom: 1px solid rgba(0, 0, 0, 0.04);
    vertical-align: middle;
    color: var(--text-secondary);
  }

  .table-row {
    transition: background-color var(--transition-fast);
  }

  .table-row:hover {
    background-color: rgba(0, 135, 60, 0.03);
  }

  .ref-id {
    color: var(--color-primary);
    font-weight: 700;
    font-size: 0.85rem;
  }

  .patient-cell {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .patient-name {
    font-weight: 600;
    color: var(--text-primary);
  }

  .patient-type {
    align-self: flex-start;
    padding: 0.1rem 0.4rem;
    font-size: 0.65rem;
  }
</style>
