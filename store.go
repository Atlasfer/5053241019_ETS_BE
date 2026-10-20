package main

import "sync"

type Store struct {
	mu sync.Mutex
	sumberDaya []SumberDaya
	peminjaman []Peminjaman
	nextSD int
	nextPem int
}

func NewStore() *Store {
	return &Store{
		sumberDaya: make([]SumberDaya, 0),
		peminjaman: make([]Peminjaman, 0),
		nextSD: 1, 
		nextPem: 1,
	}
}

func (st *Store) tambahSumberDaya(sd SumberDaya) SumberDaya {
	st.mu.Lock()
	defer st.mu.Unlock()

	sd.ID = st.nextSD
	st.nextSD++
	st.sumberDaya = append(st.sumberDaya, sd)
	return sd
}

func (st *Store) daftarSumberDaya() []SumberDaya {
	st.mu.Lock()
	defer st.mu.Unlock()

	hasil := make([]SumberDaya, len(st.sumberDaya)) // bukan nil -> JSON jadi []
	copy(hasil, st.sumberDaya)
	return hasil
}

func (st *Store) cariSumberDaya(id int) (SumberDaya, bool) {
	st.mu.Lock()
	defer st.mu.Unlock()

	for _, sd := range st.sumberDaya {
		if sd.ID == id {
			return sd, true
		}
	}
	return SumberDaya{}, false
}

func (st *Store) tambahPeminjaman(p Peminjaman) (Peminjaman, *apiError) {
	st.mu.Lock()
	defer st.mu.Unlock()

	sd, ada := st.cariSDLocked(p.SumberDayaID)
	if !ada {
		return Peminjaman{}, &apiError{400, "sumber_daya_id tidak ditemukan"}
	}
	if p.Jumlah > sd.Stok {
		return Peminjaman{}, &apiError{400, "jumlah melebihi stok sumber daya"}
	}

	p.ID = st.nextPem
	st.nextPem++
	p.Status = StatusAktif
	st.peminjaman = append(st.peminjaman, p)
	return p, nil
}

func (st *Store) cariSDLocked(id int) (SumberDaya, bool) {
	for _, sd := range st.sumberDaya {
		if sd.ID == id {
			return sd, true
		}
	}
	return SumberDaya{}, false
}

func (st *Store) daftarPeminjaman(f filterPeminjaman) []Peminjaman {
	st.mu.Lock()
	defer st.mu.Unlock()

	hasil := make([]Peminjaman, 0) // bukan nil -> JSON jadi []
	for _, p := range st.peminjaman {
		if f.SumberDayaID != nil && p.SumberDayaID != *f.SumberDayaID {
			continue
		}
		if f.Tanggal != nil && p.Tanggal != *f.Tanggal {
			continue
		}
		if f.Status != nil && p.Status != *f.Status {
			continue
		}
		hasil = append(hasil, p)
	}
	return hasil
}

func (st *Store) cariPeminjaman(id int) (Peminjaman, bool) {
	st.mu.Lock()
	defer st.mu.Unlock()

	for _, p := range st.peminjaman {
		if p.ID == id {
			return p, true
		}
	}
	return Peminjaman{}, false
}