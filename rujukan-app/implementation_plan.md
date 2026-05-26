# Rencana Modernisasi UI/UX Rujukan App

Rencana ini bertujuan untuk memperbarui tampilan aplikasi **rujukan-app** dari template SB Admin 2 klasik (berbasis Bootstrap lama) menjadi antarmuka dasbor medis yang modern, premium, responsif, dan dinamis dengan vanilla CSS dan Svelte 5.

## Masalah Tampilan Saat Ini
1. **Desain Ketinggalan Zaman**: Menggunakan template SB Admin 2 berbasis Bootstrap 4/5 lama dengan font Nunito dan sudut-sudut kaku.
2. **Kurang Dinamis**: Efek hover, transisi, dan animasi sangat minimal.
3. **Palet Warna Standar**: Warna biru-gelap default yang generic dan kurang profesional untuk aplikasi teknologi medis modern.
4. **Layout Kaku**: Sidebar dan navbar standard Bootstrap yang memakan banyak ruang dan kurang elegan.

## Desain Baru (Modern & Premium)
- **Tipografi**: Menggunakan font **Plus Jakarta Sans** (bersih, modern, dan sangat profesional untuk aplikasi SaaS/Medis).
- **Tema Warna**:
  - **Primary**: Deep Indigo (`#4F46E5` / `#6366F1`)
  - **Success**: Emerald Green (`#10B981`)
  - **Warning**: Amber Amber (`#F59E0B`)
  - **Danger**: Rose Red (`#F43F5E`)
  - **Background**: Soft Medical Grey/Slate (`#F8FAFC`)
- **Visual Styles**:
  - **Glassmorphism**: Efek blur backdrop lembut pada kartu, modal, dan sidebar.
  - **Hover Micro-Animations**: Tombol, menu sidebar, dan kartu akan bergeser ke atas (`translateY(-3px)`) dan memiliki bayangan lembut (`glowing shadows`) saat di-hover.
  - **Elegant Cards & Tables**: Baris tabel dengan sudut tumpul (`border-radius`), spasi luas, dan status tag dengan kombinasi warna soft (contoh: teks merah dengan latar merah transparan untuk status "Ditolak").
  - **Responsive Floating Layout**: Sidebar gelap modern yang sleek dengan tombol active berbentuk kapsul bercahaya.

---

## Rencana Perubahan

### 1. Global Setup & Variables

#### [MODIFY] [app.html](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/app.html)
- Ganti link font Nunito dengan **Plus Jakarta Sans**.
- Hapus link stylesheet SB Admin 2 (`sb-admin-2.css`) dan skrip jquery/sb-admin-2 yang lama.
- Definisikan layout body tanpa kelas `bg-gradient-primary` default.
- Impor file stylesheet modern global baru (`/src/app.css`).

#### [NEW] [app.css](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/app.css)
- Buat file CSS global baru yang menampung seluruh variabel CSS (`:root`), pengaturan dasar tipografi, styling tombol modern, inputs, transisi, efek glassmorphic, dan layout container dasar.

---

### 2. Autentikasi (Login & Register)

#### [MODIFY] [login/+page.svelte](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(auth)/login/+page.svelte)
- Ubah tampilan login card menjadi glassmorphic, di tengah background gradien indigo-violet yang dinamis.
- Gunakan rounded inputs, button premium dengan hover effect, dan ikon modern.

#### [MODIFY] [register/+page.svelte](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(auth)/register/+page.svelte)
- Modernisasi registrasi card dengan struktur serupa dengan login page.

---

### 3. Layout Admin

#### [MODIFY] [+layout.svelte (Admin)](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(admin)/admin/+layout.svelte)
- Desain ulang Sidebar menjadi menu mengambang (floating layout) berwarna Indigo Gelap (`#0F172A`/`#1E1B4B`) dengan logo yang modern.
- Desain ulang Topbar dengan background semi-transparan blur (glassmorphism), serta dropdown info profil yang clean.
- Perbaiki struktur footer agar lebih minimalis.

---

### 4. Halaman Dashboard & Fitur

#### [MODIFY] [+page.svelte (Dashboard)](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(admin)/admin/+page.svelte)
- Perbarui empat kartu informasi (Total Rujukan, Pending, Diterima, Diproses) menggunakan gradien warna modern, ikon bercahaya, dan hover micro-interaction.
- Perbarui tabel Rujukan Terkini dengan rounded-row styling dan status badges yang lebih estetik.
- Percantik menu akses cepat dengan tombol interaktif.

#### [MODIFY] [rujukan-masuk/+page.svelte](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(admin)/admin/rujukan-masuk/+page.svelte)
- Desain ulang filter rujukan (kapsul tab).
- Modernisasi search input dan tabel data rujukan masuk.
- Rancang ulang modal detail rujukan masuk: bagi menjadi bagian visual yang rapi (identitas pasien di kartu khusus, diagnosa dengan border berwarna, form tanggapan yang bersih, tombol aksi ber-kapsul).

#### [MODIFY] [rujukan-keluar/+page.svelte](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(admin)/admin/rujukan-keluar/+page.svelte)
- Modernisasi tabel rujukan keluar dan filter status.
- Desain ulang formulir pembuatan rujukan baru (modal popup) dengan layout 2-kolom yang seimbang, floating effect, select box yang premium, dan error message yang elegan.

#### [MODIFY] [pengaturan/+page.svelte](file:///Users/zaldimasud/Documents/PROJECT/GIT/rujukan/rujukan-app/src/routes/(admin)/admin/pengaturan/+page.svelte)
- Modernisasi layout pengaturan profil dan faskes dengan card-based design, input fields premium, dan animasi sukses/menyimpan.

---

## Rencana Verifikasi

### Manual Verification
1. Jalankan aplikasi secara lokal dengan perintah `npm run dev` di folder `rujukan-app/`.
2. Verifikasi tampilan halaman login & register apakah sudah modern, glassmorphic, dan responsif.
3. Masuk ke dashboard (`12345` / `12345`) dan cek layout sidebar, topbar, dan footer.
4. Cek interaksi hover pada menu, kartu KPI, dan baris tabel.
5. Verifikasi fungsionalitas dan estetika halaman **Rujukan Masuk** (termasuk modal detail, aksi terima/tolak).
6. Verifikasi fungsionalitas dan estetika halaman **Rujukan Keluar** (termasuk modal tambah rujukan baru).
7. Verifikasi fungsionalitas dan estetika halaman **Pengaturan**.
