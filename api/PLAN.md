# Rencana Implementasi & Dokumentasi Struktur API Rujukan

Dokumen ini berisi penjelasan struktur folder API saat ini serta rencana pengembangan (road map) untuk pengembangan sistem **Rujukan API**.

---

## 1. Struktur Folder API Saat Ini (`api/`)

Berikut adalah pohon struktur direktori dari folder `/api` beserta deskripsi singkat dari setiap komponen:

```text
api/
├── .env                  # Konfigurasi environment (DB host, port, credentials, port server, dll)
├── go.mod                # File deklarasi modul Go dan dependensi library
├── go.sum                # Berisi checksum dependensi untuk verifikasi keasamanan library
├── cmd/
│   └── app/
│       └── main.go       # Entry point utama aplikasi. Memuat .env, menginisialisasi DB, dan menjalankan server.
└── internal/
    ├── app/
    │   └── app.go        # Container utama inisialisasi aplikasi (Routing Gin, inisialisasi DB, dan registrasi modul).
    ├── infrastructure/
    │   └── database/
    │       └── mysql.go  # Driver koneksi database MySQL menggunakan GORM dengan kemampuan switch schema.
    └── modules/
        ├── user/         # Modul User (Struktur Modular Monolith)
        ├── faskes/       # Modul Faskes (Mengelola faskes dari db master)
        ├── pasien/       # Modul Pasien (Mengelola pasien dari db master)
        ├── wilayah/      # Modul Wilayah (Mengelola wilayah & jenis_wilayah dari db master)
        └── referensi/    # Modul Referensi (Mengelola referensi & jenis_referensi dari db master)
            # Setiap modul di atas memiliki struktur sub-folder standard:
            # ├── delivery/http/handler.go -> Handler HTTP
            # ├── domain/entity.go          -> Entitas Model Database
            # ├── repository/mysql.go       -> Query Database
            # └── module.go                 -> Inisialisasi & Registrasi Route
```

---

## 2. Penjelasan Detail Arsitektur Aplikasi

API ini dirancang dengan pendekatan **Modular Monolith** (Clean Architecture sederhana per modul). Pendekatan ini memastikan kode terorganisir dengan baik, mudah diuji, dan modular apabila di masa mendatang ingin dipisahkan menjadi microservices.

### A. Alur Pemrosesan Request (Request Flow)
1. **HTTP Client** mengirim request ke endpoint (contoh: `GET /v1/aplikasi/user`).
2. **Gin Router** (`internal/app/app.go`) mencocokkan route dan meneruskan request ke **Handler** (`internal/modules/user/delivery/http/handler.go`).
3. **Handler** membaca parameter input, memanggil fungsi pada **Repository** (`internal/modules/user/repository/mysql.go`) untuk mengambil data, lalu memformat response JSON.
4. **Repository** melakukan query ke database menggunakan instance **GORM** (`internal/infrastructure/database/mysql.go`) dan memetakan hasilnya ke **Entity/Domain** (`internal/modules/user/domain/entity.go`).

---

## 3. Rencana Pengembangan & Implementasi API (Plan Roadmap)

Untuk membangun API Rujukan yang lengkap dan aman, berikut adalah langkah-langkah implementasi selanjutnya:

### 🚀 Fase 1: Middleware & Security (Keamanan)
* **JWT Authentication Middleware**: Membuat token-based auth untuk melindungi endpoint sensitif.
* **CORS Middleware**: Mengatur Cross-Origin Resource Sharing agar frontend (`rujukan-app`) dapat mengakses API dengan aman.
* **Logger & Recovery Middleware**: Memastikan panic error tidak membuat server crash dan mencatat log request masuk.

### 📦 Fase 2: Peningkatan Fitur Modul User
* **Autentikasi (Login & Register)**:
  * Implementasi `POST /v1/auth/login` (verifikasi hash password bcrypt dan pembuatan token JWT).
  * Implementasi `POST /v1/auth/register` (pendaftaran user baru dengan enkripsi password).
* **CRUD User Lengkap**:
  * Tambahkan endpoint untuk:
    * `GET /v1/aplikasi/user/:id` (detail user berdasarkan ID)
    * `POST /v1/aplikasi/user` (tambah user baru)
    * `PUT /v1/aplikasi/user/:id` (update user)
    * `DELETE /v1/aplikasi/user/:id` (soft/hard delete user)

### 🏥 Fase 3: Pembuatan Modul Baru (Status: SEBAGIAN SELESAI ✅)
Mengikuti pola modular yang sama dengan `user`, kita telah mengimplementasikan modul data master pendukung rujukan dari database `master`:
* **Modul Faskes (Fasilitas Kesehatan) - SELESAI ✅**:
  * Mengelola data rumah sakit, puskesmas, klinik.
* **Modul Pasien - SELESAI ✅**:
  * Manajemen data medis dasar pasien yang dirujuk.
* **Modul Wilayah - SELESAI ✅**:
  * Menyediakan data provinsi, kota/kabupaten, kecamatan, kelurahan beserta jenis wilayah.
* **Modul Referensi - SELESAI ✅**:
  * Menyediakan data referensi kelompok faskes dan jenis kelamin.
* **Modul Rujukan - RENCANA SELANJUTNYA ⏳**:
  * Mengelola alur rujukan pasien antar faskes.
  * Endpoint `POST /v1/rujukan` (buat rujukan baru).
  * Endpoint `GET /v1/rujukan/:id` (cek status rujukan).

### 📝 Fase 4: Validasi & Standarisasi Error
* **Request Validation**: Integrasi `github.com/go-playground/validator` pada handler untuk memvalidasi input request body secara otomatis.
* **Standard Response Helper**: Membuat format response yang konsisten untuk sukses dan error di semua modul.

### 📖 Fase 5: Dokumentasi API & Deployment
* **Swagger/OpenAPI**: Integrasi `swaggo/swag` untuk generate dokumentasi API otomatis.
* **Database Migration & Patches**:
  * Pengelolaan perubahan skema database dilakukan menggunakan patch SQL terurut di dalam folder `db/patches/` (misalnya `0001_deskripsi_patch.sql`).
  * Integrasi library migrasi di Go (seperti `golang-migrate`) di masa mendatang untuk mengeksekusi patch secara otomatis saat server berjalan.

---

## 4. Struktur Database & Pengelolaan Patch (Database Patches)

Untuk melacak perubahan struktur database, proyek ini menyediakan direktori khusus untuk patch database di luar folder `api`:

```text
db/
├── init.sql              # Inisialisasi awal server MySQL Docker
├── aplikasi/             # Skema database 'aplikasi'
├── master/               # Skema database 'master'
└── patches/              # FOLDER BARU: Menyimpan perubahan schema incremental (Patches)
    ├── README.md         # Dokumentasi petunjuk penggunaan patch
    └── XXXX_deskripsi.sql# File patch SQL berurutan (misal: 0001_add_column_status.sql)
```

Setiap developer yang membuat perubahan skema tabel wajib mencatatnya dalam bentuk file patch baru di dalam folder `db/patches/` agar dapat dideploy ke environment staging maupun production secara konsisten.

