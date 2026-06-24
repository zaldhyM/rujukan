# Sistem Rujukan Pasien

Aplikasi manajemen rujukan pasien antar fasilitas kesehatan. Terdiri dari dua komponen utama: REST API (Go) dan antarmuka web (SvelteKit).

---

## Struktur Proyek

```
rujukan/
├── api/              # Backend — REST API (Go + Gin)
├── rujukan-app/      # Frontend — Web App (SvelteKit)
├── db/               # SQL schema & migration scripts
│   ├── init.sql      # Script inisialisasi database
│   ├── aplikasi/     # Schema aplikasi (tabel user)
│   ├── master/       # Schema master (faskes, pasien, wilayah, referensi)
│   ├── rujukan/      # Schema rujukan (tabel rujukan)
│   └── patches/      # Patch SQL incremental
├── docker-compose.yml
└── .env.dist         # Template environment variables
```

---

## Prasyarat

Pastikan sudah terinstall:

| Tools | Versi minimum |
|---|---|
| [Go](https://go.dev/dl/) | 1.24+ |
| [Node.js](https://nodejs.org/) | 18+ |
| [Docker](https://www.docker.com/) + Docker Compose | — |

---

## Cara Menjalankan

### 1. Clone Repositori

```bash
git clone <url-repo>
cd rujukan
```

### 2. Konfigurasi Environment

Salin file template dan isi sesuai kebutuhan:

```bash
cp .env.dist docker-compose.yml
```

Buat file `.env` untuk API:

```bash
cp api/.env.example api/.env   # jika ada, atau buat manual
```

Isi `api/.env`:

```env
# Server
PORT=8081
GIN_MODE=debug

# Database
DB_HOST=127.0.0.1
DB_PORT=3307
DB_USER=admin
DB_PASS=password_anda
DB_NAME=mydb

# Redis
REDIS_ADDR=127.0.0.1:6379

# JWT — wajib diisi dengan string acak yang kuat
JWT_SECRET=isi_dengan_string_acak_minimal_32_karakter

# Session
TOKEN_EXPIRED_HOURS=24

# CORS — origin frontend yang diizinkan
CORS_ORIGIN=http://localhost:5173
```

### 3. Jalankan Database & Redis (Docker)

```bash
docker compose up -d
```

Verifikasi container berjalan:

```bash
docker compose ps
```

Dua container harus berstatus `running`:
- `mysql_server` — MySQL 8.0 di port `3307`
- `redis_server` — Redis 7 di port `6379`

### 4. Inisialisasi Database

Masuk ke MySQL dan jalankan script init:

```bash
docker exec -i mysql_server mysql -u root -p < db/init.sql
```

Script `init.sql` akan membuat semua schema dan tabel yang dibutuhkan.

### 5. Jalankan API

```bash
cd api
go mod download
go run cmd/app/main.go
```

API berjalan di `http://localhost:8081`.

### 6. Jalankan Frontend

Buka terminal baru:

```bash
cd rujukan-app
npm install
npm run dev
```

Frontend berjalan di `http://localhost:5173`.

---

## Endpoint API

Base URL: `http://localhost:8081/v1`

### Publik (tanpa token)

| Method | URL | Keterangan |
|---|---|---|
| `POST` | `/auth/login` | Login, mendapat JWT token |
| `POST` | `/auth/register` | Daftar user baru |

### Terproteksi (perlu `Authorization: Bearer <token>`)

**Auth & User**

| Method | URL | Keterangan |
|---|---|---|
| `GET` | `/auth/me` | Profil user yang sedang login |
| `POST` | `/auth/logout` | Logout, hapus sesi |
| `GET` | `/aplikasi/user` | List semua user |

**Rujukan**

| Method | URL | Keterangan |
|---|---|---|
| `GET` | `/rujukan` | List rujukan (`?search=`, `?status=`, `?limit=`, `?offset=`) |
| `GET` | `/rujukan/:id` | Detail rujukan |
| `POST` | `/rujukan` | Buat rujukan baru |
| `PUT` | `/rujukan/:id/status` | Update status rujukan |

**Data Master**

| Method | URL | Keterangan |
|---|---|---|
| `GET` | `/faskes` | List fasilitas kesehatan |
| `GET` | `/pasien` | List pasien |
| `GET` | `/wilayah` | List wilayah |
| `GET` | `/referensi` | List referensi |

---

## Format Response

Semua endpoint mengembalikan JSON dengan format konsisten:

```json
// Sukses data tunggal
{ "success": true, "data": { ... } }

// Sukses list
{ "success": true, "total": 100, "limit": 10, "offset": 0, "data": [ ... ] }

// Sukses pesan
{ "success": true, "message": "..." }

// Error
{ "success": false, "message": "..." }

// Error validasi
{ "success": false, "message": "Validasi input gagal", "errors": { "field": "pesan" } }
```

---

## Perintah Docker Berguna

```bash
# Jalankan semua container
docker compose up -d

# Hentikan tanpa hapus data
docker compose stop

# Hentikan dan hapus container (data tetap ada)
docker compose down

# Hentikan dan hapus container + volume (data hilang)
docker compose down -v

# Masuk ke MySQL
docker exec -it mysql_server mysql -u root -p

# Lihat log container
docker compose logs -f
```

---

## Dokumentasi Lebih Lanjut

- `api/ARCHITECTURE.md` — arsitektur lengkap backend, alur request, dan keamanan
- `api/NOTES.md` — catatan pengembangan, roadmap, dan backlog
