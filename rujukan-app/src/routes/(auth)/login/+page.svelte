<script>
  import { onMount } from 'svelte';
  import { login, isLoggedIn } from '$lib/auth';

  let username = $state('');
  let password = $state('');
  let errorMessage = $state('');
  let successMessage = $state('');

  onMount(() => {
    if (isLoggedIn()) {
      window.location.href = '/admin';
    }
  });

  /**
   * @param {SubmitEvent} e
   */
  function handleLogin(e) {
    e.preventDefault();
    errorMessage = '';
    
    if (!username || !password) {
      errorMessage = 'Username dan password wajib diisi!';
      return;
    }

    const res = login(username, password);
    if (res.success) {
      successMessage = 'Login Berhasil! Mengalihkan...';
      setTimeout(() => {
        window.location.href = '/admin';
      }, 1000);
    } else {
      errorMessage = res.message || 'Login gagal';
    }
  }
</script>

<svelte:head>
  <title>Login - Aplikasi Rujukan Pasien</title>
</svelte:head>

<div class="container d-flex align-items-center justify-content-center min-vh-100">
  <div class="row justify-content-center w-100">
    <div class="col-xl-5 col-lg-6 col-md-8">
      <div class="card o-hidden border-0 shadow-lg my-5 rounded-4 overflow-hidden">
        <div class="card-body p-0">
          <div class="p-5">
            <div class="text-center mb-4">
              <div class="d-inline-flex align-items-center justify-content-center bg-primary text-white rounded-circle mb-3" style="width: 60px; height: 60px; font-size: 24px;">
                <i class="fas fa-hospital"></i>
              </div>
              <h1 class="h3 text-gray-900 font-weight-bold mb-1">Aplikasi Rujukan</h1>
              <p class="text-muted small">Silakan masuk untuk mengakses dashboard rujukan</p>
            </div>
            
            {#if errorMessage}
              <div class="alert alert-danger alert-dismissible fade show text-center py-2 mb-3 small" role="alert">
                {errorMessage}
              </div>
            {/if}

            {#if successMessage}
              <div class="alert alert-success text-center py-2 mb-3 small" role="alert">
                {successMessage}
              </div>
            {/if}

            <form class="user" onsubmit={handleLogin}>
              <div class="form-group mb-3">
                <label for="username" class="form-label small text-gray-700 font-weight-bold">Username / NIP</label>
                <input
                  type="text"
                  class="form-control form-control-user rounded-pill py-2 px-3"
                  id="username"
                  placeholder="Masukkan Username (contoh: 12345)"
                  bind:value={username}
                />
              </div>
              <div class="form-group mb-4">
                <label for="password" class="form-label small text-gray-700 font-weight-bold">Password</label>
                <input
                  type="password"
                  class="form-control form-control-user rounded-pill py-2 px-3"
                  id="password"
                  placeholder="Masukkan Password (contoh: 12345)"
                  bind:value={password}
                />
              </div>
              
              <button
                type="submit"
                class="btn btn-primary btn-user btn-block w-100 rounded-pill py-2 font-weight-bold shadow-sm"
              >
                Masuk
              </button>
            </form>
            
            <hr class="my-4" />
            
            <div class="text-center">
              <span class="small text-muted">Belum memiliki akun? </span>
              <a class="small font-weight-bold" href="/register">Buat Akun Baru</a>
            </div>
            
            <div class="text-center mt-3">
              <span class="text-xs text-muted">Akun demo: <strong>12345</strong> / <strong>12345</strong></span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  /* Custom modern overrides */
  :global(body) {
    background: linear-gradient(135deg, #4e73df 0%, #224abe 100%) !important;
  }
  .card {
    border-radius: 1rem !important;
  }
  .form-control-user {
    font-size: 0.9rem !important;
    padding: 0.75rem 1.25rem !important;
  }
  .btn-user {
    font-size: 0.9rem !important;
    padding: 0.75rem 1.25rem !important;
  }
</style>
