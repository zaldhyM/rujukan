# Catatan Pengembangan — rujukan-app (Frontend)

## Stack
- **Framework:** Svelte 5 (file-based routing)
- **Build tool:** Vite
- **Language:** TypeScript
- **Dev server:** `npm run dev` → default port `5174`

---

## Struktur Halaman (`src/routes/`)

| Route | File | Keterangan |
|---|---|---|
| `/` | `+page.svelte` | Landing page |
| `/login` | `login/+page.svelte` | Halaman login |
| `/register` | `register/+page.svelte` | Halaman daftar faskes |
| `/dashboard` | `dashboard/+page.svelte` | Ringkasan/KPI |
| `/dashboard/rujukan-masuk` | `dashboard/rujukan-masuk/+page.svelte` | Kelola rujukan masuk |
| `/dashboard/rujukan-keluar` | `dashboard/rujukan-keluar/+page.svelte` | Buat rujukan keluar |
| `/dashboard/pengaturan` | `dashboard/pengaturan/+page.svelte` | Pengaturan akun |

Layout dashboard (sidebar + topbar + footer): `dashboard/+layout.svelte`  
CSS global + variabel tema: `src/routes/global.css`

---

## Tema — Kemenkes Light

Warna utama diubah ke tema terang mengacu palet Kementerian Kesehatan RI.

```css
--bg-primary:          #F0F7F3   /* latar halaman */
--bg-secondary:        #FFFFFF   /* kartu/panel */
--color-primary:       #00873C   /* hijau Kemenkes */
--color-primary-hover: #006B2F
--color-secondary:     #8DC63F   /* lime/aksen */
--text-primary:        #111827
--text-secondary:      #374151
--text-muted:          #6B7280
```

**Aturan desain:**
- Sidebar dashboard: solid `#005C28` (hijau gelap), teks putih
- Nav item aktif: background `#8DC63F`, teks `#005C28`
- Topbar: putih `#FFFFFF` dengan shadow tipis
- Tombol CTA utama: `#8DC63F` (lime), teks `#005C28`
- Landing page — panel kiri: `#005C28`, panel kanan: `#FFFFFF`

---

## Status Halaman

| Halaman | Status | Catatan |
|---|---|---|
| Landing page | ✅ Selesai | Tema Kemenkes, dua panel |
| Login | ✅ Selesai | Tema Kemenkes, integrasi API belum |
| Register | ✅ Selesai | Tema Kemenkes, integrasi API belum |
| Dashboard | ✅ Selesai | Mock data, belum connect ke API |
| Rujukan Masuk | ⏳ Partial | Tampilan ada, belum connect ke API |
| Rujukan Keluar | ⏳ Partial | Tampilan ada, belum connect ke API |
| Pengaturan | ⏳ Partial | Tampilan ada, belum connect ke API |

---

## Yang Perlu Dilakukan Selanjutnya

- [ ] Integrasi API login ke `POST /v1/auth/login`
- [ ] Simpan JWT token ke localStorage/cookie setelah login
- [ ] Koneksi halaman Rujukan Masuk ke `GET /v1/rujukan?status=menunggu`
- [ ] Koneksi halaman Rujukan Keluar ke `POST /v1/rujukan`
- [ ] Update status rujukan via `PUT /v1/rujukan/:id/status`
- [ ] Halaman pengaturan: ganti password, update profil
