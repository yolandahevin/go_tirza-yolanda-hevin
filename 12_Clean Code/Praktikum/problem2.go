package main

type mobil struct {
	kendaraan
}

type kendaraan struct {
	kecepatanPerJam int
}

func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	cepat := mobil{}
	cepat.tambahKecepatan(10)
}
