package main

import (
	"strings"
	"time"
	"unicode/utf8"
)

func validasiSumberDaya(req sumberDayaRequest) (SumberDaya, string) {
	if req.Nama == nil || req.Jenis == nil || req.Stok == nil {
		return SumberDaya{}, "nama, jenis, dan stok wajib diisi"
	}

	nama := strings.TrimSpace(*req.Nama)
	jenis := strings.TrimSpace(*req.Jenis)
	stok := *req.Stok

	if nama == "" {
		return SumberDaya{}, "nama tidak boleh kosong"
	}
	if utf8.RuneCountInString(nama) > 50 { // per karakter, bukan per byte
		return SumberDaya{}, "nama maksimal 50 karakter"
	}
	if jenis != "ruang" && jenis != "alat" {
		return SumberDaya{}, `jenis harus "ruang" atau "alat"`
	}
	if stok < 1 {
		return SumberDaya{}, "stok minimal 1"
	}

	return SumberDaya{Nama: nama, Jenis: jenis, Stok: stok}, ""
}

func validasiPeminjaman(req peminjamanRequest) (Peminjaman, string) {
	if req.SumberDayaID == nil || req.Peminjam == nil || req.Tanggal == nil ||
		req.SlotMulai == nil || req.SlotSelesai == nil || req.Jumlah == nil {
		return Peminjaman{}, "sumber_daya_id, peminjam, tanggal, slot_mulai, slot_selesai, dan jumlah wajib diisi"
	}

	peminjam := strings.TrimSpace(*req.Peminjam)
	tanggal := strings.TrimSpace(*req.Tanggal)

	if peminjam == "" {
		return Peminjaman{}, "peminjam tidak boleh kosong"
	}
	if utf8.RuneCountInString(peminjam) > 50 {
		return Peminjaman{}, "peminjam maksimal 50 karakter"
	}
	if !tanggalValid(tanggal) {
		return Peminjaman{}, "tanggal harus berformat YYYY-MM-DD dan merupakan tanggal yang ada"
	}
	if *req.SlotMulai < 1 || *req.SlotMulai > 8 || *req.SlotSelesai < 1 || *req.SlotSelesai > 8 {
		return Peminjaman{}, "slot harus antara 1 sampai 8"
	}
	if *req.SlotMulai > *req.SlotSelesai {
		return Peminjaman{}, "slot_mulai tidak boleh lebih besar dari slot_selesai"
	}
	if *req.Jumlah < 1 {
		return Peminjaman{}, "jumlah minimal 1"
	}

	return Peminjaman{
		SumberDayaID: *req.SumberDayaID,
		Peminjam:     peminjam,
		Tanggal:      tanggal,
		SlotMulai:    *req.SlotMulai,
		SlotSelesai:  *req.SlotSelesai,
		Jumlah:       *req.Jumlah,
	}, ""
}

func tanggalValid(s string) bool {
	if len(s) != 10 {
		return false
	}
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}