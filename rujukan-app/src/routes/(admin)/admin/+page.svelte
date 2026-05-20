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
    <h1 class="h3 mb-1 text-gray-800 font-weight-bold">Dashboard</h1>
    <p class="text-muted mb-0 small">Selamat datang kembali, <strong>{user.name}</strong>. Berikut adalah ikhtisar data rujukan hari ini.</p>
  </div>
</div>

<!-- Info Cards Row -->
<div class="row">
  <!-- Total Rujukan -->
  <div class="col-xl-3 col-md-6 mb-4">
    <div class="card border-left-primary shadow-sm h-100 py-2">
      <div class="card-body">
        <div class="row no-gutters align-items-center">
          <div class="col mr-2">
            <div class="text-xs font-weight-bold text-primary text-uppercase mb-1">
              Total Rujukan (Aktif)
            </div>
            <div class="h3 mb-0 font-weight-bold text-gray-800">{totalRujukan}</div>
          </div>
          <div class="col-auto text-primary">
            <Fa icon={faExchangeAlt} size="2x" class="text-gray-300" />
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Rujukan Masuk Pending -->
  <div class="col-xl-3 col-md-6 mb-4">
    <div class="card border-left-warning shadow-sm h-100 py-2">
      <div class="card-body">
        <div class="row no-gutters align-items-center">
          <div class="col mr-2">
            <div class="text-xs font-weight-bold text-warning text-uppercase mb-1">
              Rujukan Masuk Pending
            </div>
            <div class="h3 mb-0 font-weight-bold text-gray-800">{pendingIncoming}</div>
          </div>
          <div class="col-auto text-warning">
            <Fa icon={faClock} size="2x" class="text-gray-300" />
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Rujukan Masuk Diterima -->
  <div class="col-xl-3 col-md-6 mb-4">
    <div class="card border-left-success shadow-sm h-100 py-2">
      <div class="card-body">
        <div class="row no-gutters align-items-center">
          <div class="col mr-2">
            <div class="text-xs font-weight-bold text-success text-uppercase mb-1">
              Rujukan Masuk Diterima
            </div>
            <div class="h3 mb-0 font-weight-bold text-gray-800">{acceptedIncoming}</div>
          </div>
          <div class="col-auto text-success">
            <Fa icon={faCheckCircle} size="2x" class="text-gray-300" />
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Rujukan Keluar Diproses -->
  <div class="col-xl-3 col-md-6 mb-4">
    <div class="card border-left-info shadow-sm h-100 py-2">
      <div class="card-body">
        <div class="row no-gutters align-items-center">
          <div class="col mr-2">
            <div class="text-xs font-weight-bold text-info text-uppercase mb-1">
              Rujukan Keluar Diproses
            </div>
            <div class="h3 mb-0 font-weight-bold text-gray-800">{processedOutgoing}</div>
          </div>
          <div class="col-auto text-info">
            <Fa icon={faPaperPlane} size="2x" class="text-gray-300" />
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<div class="row">
  <!-- Recent Referrals -->
  <div class="col-lg-8">
    <div class="card shadow-sm mb-4">
      <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between bg-white border-bottom-0">
        <h6 class="m-0 font-weight-bold text-primary">Rujukan Terkini</h6>
        <span class="badge badge-light text-muted font-weight-normal">{recentActivities().length} Teratas</span>
      </div>
      <div class="card-body p-0">
        {#if recentActivities().length === 0}
          <div class="text-center py-5">
            <p class="text-muted mb-0">Belum ada aktivitas rujukan</p>
          </div>
        {:else}
          <div class="table-responsive">
            <table class="table table-hover align-middle mb-0">
              <thead class="bg-light text-gray-600 font-weight-bold text-xs">
                <tr>
                  <th class="border-0 px-4">No. Rujukan</th>
                  <th class="border-0">Tgl Rujukan</th>
                  <th class="border-0">Jenis</th>
                  <th class="border-0">Pasien</th>
                  <th class="border-0">Faskes Asal/Tujuan</th>
                  <th class="border-0 text-center">Status</th>
                </tr>
              </thead>
              <tbody class="small">
                {#each recentActivities() as r (r.id)}
                  <tr>
                    <td class="font-weight-bold px-4 py-3 text-primary">{r.noRujukan}</td>
                    <td>{r.tglRujukan}</td>
                    <td>
                      <span class="badge" class:bg-primary={r.type === 'Masuk'} class:bg-info={r.type === 'Keluar'}>
                        {r.type}
                      </span>
                    </td>
                    <td>
                      <div class="font-weight-bold text-gray-800">{r.pasien.nama}</div>
                      <div class="text-xs text-muted">No: {r.pasien.noKartu}</div>
                    </td>
                    <td>{r.type === 'Masuk' ? r.faskesAsal : r.faskesTujuan}</td>
                    <td class="text-center">
                      <span class="badge" 
                        class:bg-warning={r.status === 'Pending' || r.status === 'Diproses'}
                        class:bg-success={r.status === 'Diterima' || r.status === 'Selesai' || r.status === 'Dikirim'}
                        class:bg-danger={r.status === 'Ditolak'}
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
  </div>

  <!-- Quick Actions & Stats -->
  <div class="col-lg-4">
    <!-- Quick Access Card -->
    <div class="card shadow-sm mb-4">
      <div class="card-header py-3 bg-white border-bottom-0">
        <h6 class="m-0 font-weight-bold text-primary">Akses Cepat Menu</h6>
      </div>
      <div class="card-body">
        <a href="/admin/rujukan-masuk" class="btn btn-outline-primary btn-block text-left mb-2 d-flex align-items-center justify-content-between p-3 rounded-3">
          <div>
            <Fa icon={faInbox} class="mr-2" />
            <strong>Rujukan Masuk</strong>
            <div class="text-xs text-muted mt-1">Lihat dan verifikasi rujukan masuk</div>
          </div>
          <Fa icon={faChevronRight} class="text-gray-400" />
        </a>
        <a href="/admin/rujukan-keluar" class="btn btn-outline-primary btn-block text-left mb-2 d-flex align-items-center justify-content-between p-3 rounded-3">
          <div>
            <Fa icon={faPaperPlane} class="mr-2" />
            <strong>Rujukan Keluar</strong>
            <div class="text-xs text-muted mt-1">Buat dan monitoring rujukan keluar</div>
          </div>
          <Fa icon={faChevronRight} class="text-gray-400" />
        </a>
      </div>
    </div>

    <!-- Info Box -->
    <div class="card bg-primary text-white shadow-sm">
      <div class="card-body p-4">
        <div class="text-center mb-3">
          <Fa icon={faUserCheck} size="3x" />
        </div>
        <h5 class="font-weight-bold text-center mb-2">Informasi Akun</h5>
        <p class="small text-center mb-0 opacity-75">
          Anda login sebagai <strong>{user.name}</strong> dengan hak akses penuh sistem rujukan rumah sakit.
        </p>
      </div>
    </div>
  </div>
</div>
