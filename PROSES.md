# PROSES

## 1. Persona dan masalahnya
(2-3 kalimat dengan kata-katamu sendiri: siapa pengguna persona kamu, masalah apa yang diselesaikan fitur khasnya.)

## 2. Rancangan struktur data
Data disimpan di memori dalam satu struct `Store`:

- `sumberDaya []SumberDaya` berisi id, nama, jenis ("ruang"/"alat"), dan stok.
- `peminjaman []Peminjaman` berisi id, sumber_daya_id, peminjam, tanggal,
  slot_mulai, slot_selesai, jumlah, dan status.
- `nextSD` dan `nextPem` adalah counter id untuk masing-masing jenis data.
- `mu sync.Mutex` melindungi semua data di atas.

## 3. Satu keputusan desain
Saya memisahkan validasi menjadi dua tahap:

1. `validasiPeminjaman` adalah fungsi murni yang memeriksa bentuk data (format
   tanggal, rentang slot, field wajib) tanpa akses ke penyimpanan.
2. Pemeriksaan yang butuh data tersimpan (sumber daya ada, jumlah tidak melebihi
   stok, dan nanti deteksi bentrok) dilakukan di dalam `store.tambahPeminjaman`,
   di bawah satu `Lock()` bersama proses penyimpanan.
Pilihan lain yang saya pertimbangkan adalah memeriksa semuanya di handler dan
mengunci hanya saat `append`. Saya tidak memilihnya karena pengecekan dan
penyimpanan yang dikunci terpisah membuka celah race: dua request bisa sama-sama
lolos pengecekan lalu sama-sama menyimpan. Dengan satu kunci, pengecekan dan
penyimpanan menjadi satu langkah yang tidak bisa disela request lain. Fungsi
murni juga lebih mudah diuji.

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala: Kendala stup pertama kali
- Cara menemukan:
- Penyebab dan solusi:

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
