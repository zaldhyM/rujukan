# Database Patches / Migrations

Direktori ini digunakan untuk menyimpan file patch SQL (migrasi database manual) secara berurutan.

## Cara Penamaan File Patch
Untuk menjaga urutan eksekusi database patch, gunakan format penamaan berikut:
`XXXX_deskripsi_patch.sql`

Di mana `XXXX` adalah nomor urut 4 digit (dimulai dari `0001`).

Contoh:
* `0001_add_status_telegram_to_user.sql`
* `0002_create_table_rujukan.sql`

## Cara Menjalankan Patch
Jika Anda melakukan perubahan database pada database development/production secara manual:
1. Tulis query SQL perubahan tersebut ke dalam file patch baru di folder ini.
2. Jalankan file SQL tersebut di server database MySQL Anda menggunakan command tool SQL client (seperti DBeaver, TablePlus, Navicat) atau terminal:
   ```bash
   docker exec -i mysql_server mysql -u root -p[password] aplikasi < db/patches/XXXX_deskripsi_patch.sql
   ```
