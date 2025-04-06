🧭 High-Level System Role
User (Penumpang) – Request, bayar, kasih rating

Driver – Terima order, antar user, update status

Admin – Monitoring, verifikasi driver, kontrol sistem

🔄 Flow Kompleks Ojek Online (User ↔ Driver ↔ Admin)
🎯 1. Registrasi & Verifikasi
User:
POST /auth/register → daftar user

POST /auth/login → login

Driver:
POST /auth/register-driver

Admin review → POST /admin/approve-driver

Admin:
Verifikasi dokumen driver

Approve/reject driver via dashboard

🧭 2. Order Ride (User Request → Driver Matching)
User:
Pilih titik jemput & tujuan → POST /trips/request

Sistem mencari driver terdekat (radius 3 km) via geo-indexing

Backend:
Cari driver aktif dengan lokasi paling dekat (bisa Redis atau PostGIS)

Kirim notifikasi (misalnya pakai WebSocket atau Firebase) ke driver: POST /driver/notify

🏁 3. Accept & Jalankan Order
Driver:
POST /trips/:id/accept

Update posisi driver secara berkala → POST /driver/location

User:
Lihat posisi driver → GET /trips/:id/track

Admin:
Monitor live trip & status driver (real-time tracking dashboard)

💸 4. Pembayaran (Integrasi Payment SDK)
Sebelum jalan:
Sistem hitung biaya → GET /trips/:id/estimate

User bayar via SDK (Xendit/Midtrans/DANA/etc)

POST /payments/initiate

Redirect / SDK callback → update trip status jadi paid

Trip tidak bisa dimulai kalau belum dibayar

Setelah selesai:
POST /trips/:id/complete

Sistem bagi hasil:

Misal: 80% ke driver, 20% ke platform

Tambah saldo driver (virtual wallet)

⭐ 5. Rating & Feedback
User:
POST /trips/:id/rate → kasih rating dan komentar ke driver

Driver:
Bisa kasih rating balik ke user (optional)

Admin:
Lihat daftar rating, banned user/driver toxic

⚙️ 6. Admin Panel Features (Opsional tapi powerful)
Verifikasi driver & user

Blokir akun nakal

Lihat statistik:

Order terbanyak, rating rendah, driver dengan penghasilan tertinggi

Export laporan harian