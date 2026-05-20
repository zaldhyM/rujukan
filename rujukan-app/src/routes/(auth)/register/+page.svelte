<script>
  import { register } from '$lib/auth';

  let name = $state('');
  let username = $state('');
  let password = $state('');
  let confirmPassword = $state('');
  let errorMessage = $state('');
  let successMessage = $state('');

  /**
   * @param {SubmitEvent} e
   */
  function handleRegister(e) {
    e.preventDefault();
    errorMessage = '';
    successMessage = '';

    if (!name || !username || !password || !confirmPassword) {
      errorMessage = 'Semua kolom wajib diisi!';
      return;
    }

    if (password !== confirmPassword) {
      errorMessage = 'Konfirmasi password tidak cocok!';
      return;
    }

    if (password.length < 5) {
      errorMessage = 'Password minimal terdiri dari 5 karakter!';
      return;
    }

    const res = register(name, username, password);
    if (res.success) {
      successMessage = 'Registrasi berhasil! Mengalihkan ke halaman login...';
      setTimeout(() => {
        window.location.href = '/login';
      }, 1500);
    } else {
      errorMessage = res.message || 'Registrasi gagal';
    }
  }
</script>

<svelte:head>
  <title>Registrasi Akun Baru - Aplikasi Rujukan Pasien</title>
</svelte:head>

<div class="container d-flex align-items-center justify-content-center min-vh-100">
  <div class="row justify-content-center w-100">
    <div class="col-xl-6 col-lg-7 col-md-9">
      <div class="card o-hidden border-0 shadow-lg my-5 rounded-4 overflow-hidden">
        <div class="card-body p-0">
          <div class="p-5">
            <div class="text-center mb-4">
              <div class="d-inline-flex align-items-center justify-content-center bg-primary text-white rounded-circle mb-3" style="width: 60px; height: 60px; font-size: 24px;">
                <i class="fas fa-user-plus"></i>
              </div>
              <h1 class="h3 text-gray-900 font-weight-bold mb-1">Daftar Akun Baru</h1>
              <p class="text-muted small">Buat akun untuk mengelola rujukan pasien</p>
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

            <form class="user" onsubmit={handleRegister}>
              <div class="form-group mb-3">
                <label for="name" class="form-label small text-gray-700 font-weight-bold">Nama Lengkap & Gelar</label>
                <input
                  type="text"
                  class="form-control form-control-user rounded-pill py-2 px-3"
                  id="name"
                  placeholder="Contoh: Dr. John Doe, Sp.PD"
                  bind:value={name}
                />
              </div>
              
              <div class="form-group mb-3">
                <label for="username" class="form-label small text-gray-700 font-weight-bold">Username / NIP</label>
                <input
                  type="text"
                  class="form-control form-control-user rounded-pill py-2 px-3"
                  id="username"
                  placeholder="Masukkan Username baru"
                  bind:value={username}
                />
              </div>
              
              <div class="row mb-4">
                <div class="col-sm-6 mb-3 mb-sm-0">
                  <label for="password" class="form-label small text-gray-700 font-weight-bold">Password</label>
                  <input
                    type="password"
                    class="form-control form-control-user rounded-pill py-2 px-3"
                    id="password"
                    placeholder="Password"
                    bind:value={password}
                  />
                </div>
                <div class="col-sm-6">
                  <label for="confirmPassword" class="form-label small text-gray-700 font-weight-bold">Konfirmasi Password</label>
                  <input
                    type="password"
                    class="form-control form-control-user rounded-pill py-2 px-3"
                    id="confirmPassword"
                    placeholder="Ulangi Password"
                    bind:value={confirmPassword}
                  />
                </div>
              </div>
              
              <button
                type="submit"
                class="btn btn-primary btn-user btn-block w-100 rounded-pill py-2 font-weight-bold shadow-sm"
              >
                Daftar Akun
              </button>
            </form>
            
            <hr class="my-4" />
            
            <div class="text-center">
              <span class="small text-muted">Sudah memiliki akun? </span>
              <a class="small font-weight-bold" href="/login">Masuk (Login) di sini</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
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
