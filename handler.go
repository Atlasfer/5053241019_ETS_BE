package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type server struct {
	store *Store
}

func NewHandler() http.Handler {
	s := &server{store: NewStore()}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /sumber-daya", s.buatSumberDaya)
	mux.HandleFunc("GET /sumber-daya", s.daftarSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", s.detailSumberDaya)
	mux.HandleFunc("POST /peminjaman", s.buatPeminjaman)
	mux.HandleFunc("GET /peminjaman", s.daftarPeminjaman)
	mux.HandleFunc("GET /peminjaman/{id}", s.detailPeminjaman)
	return mux
}

func (s *server) buatSumberDaya(w http.ResponseWriter, r *http.Request) {
	var req sumberDayaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tulisError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}

	sd, pesan := validasiSumberDaya(req)
	if pesan != "" {
		tulisError(w, http.StatusBadRequest, pesan)
		return
	}

	sd = s.store.tambahSumberDaya(sd)
	tulisJSON(w, http.StatusCreated, sd) 
}

func (s *server) daftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	tulisJSON(w, http.StatusOK, s.store.daftarSumberDaya()) 
}

func (s *server) detailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}

	sd, ada := s.store.cariSumberDaya(id)
	if !ada {
		tulisError(w, http.StatusNotFound, "sumber daya tidak ditemukan")
		return
	}

	tulisJSON(w, http.StatusOK, sd) 
}

func (s *server) buatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var req peminjamanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tulisError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}

	p, pesan := validasiPeminjaman(req)
	if pesan != "" {
		tulisError(w, http.StatusBadRequest, pesan)
		return
	}

	hasil, apiErr := s.store.tambahPeminjaman(p)
	if apiErr != nil {
		tulisApiError(w, apiErr)
		return
	}

	tulisJSON(w, http.StatusCreated, hasil)
}

func (s *server) daftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var f filterPeminjaman

	if q.Has("sumber_daya_id") {
		id, err := strconv.Atoi(q.Get("sumber_daya_id"))
		if err != nil {
			tulisError(w, http.StatusBadRequest, "sumber_daya_id harus berupa angka")
			return
		}
		f.SumberDayaID = &id
	}
	if q.Has("tanggal") {
		t := q.Get("tanggal")
		f.Tanggal = &t
	}
	if q.Has("status") {
		st := q.Get("status")
		f.Status = &st
	}

	tulisJSON(w, http.StatusOK, s.store.daftarPeminjaman(f))
}

func (s *server) detailPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}

	p, ada := s.store.cariPeminjaman(id)
	if !ada {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan") // 404
		return
	}

	tulisJSON(w, http.StatusOK, p)
}