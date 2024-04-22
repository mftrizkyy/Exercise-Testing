package main

import (
	"errors"
	"fmt"
	"testing"
)
const (
	tax = 10
	app = 2000
)

func HitungHargaTotal(hargaItem, ongkir float64, qty int) (float64, error) {
	if hargaItem <= 0 {
		return 0, errors.New("harga barang tidak boleh nol")
	}

	if qty <= 0 {
		return 0, errors.New("jumlah barang tidak boleh nol")
	}

	hargaAkhirItem := hargaItem * float64(qty)

	if ongkir <= 0{
		return 0, errors.New("harga ongkir tidak boleh nol")
	}

	hargaSetelahOngkir := hargaAkhirItem + ongkir
	pajak := hargaAkhirItem * tax / 100
	total := hargaSetelahOngkir + pajak + app

    return total, nil
}

func TestHitungHargaTotal(t *testing.T) {
	type args struct {
		hargaItem float64
		ongkir    float64
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HitungHargaTotal(tt.args.hargaItem, tt.args.ongkir, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("HitungHargaTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HitungHargaTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
func PembayaranBarang(hargaTotal float64, metode string, dicicil bool) error {
	// Cek hargaTotal
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	// Cek metode pembayaran
	validMetode := map[string]bool{"cod": true, "transfer": true, "debit": true, "credit": true, "gerai": true}
	if !validMetode[metode] {
		return errors.New("metode tidak dikenali")
	}

	// Cek apakah pembayaran dicicil atau tidak
	if dicicil {
		// Jika dicicil, metode harus credit dan hargaTotal >= 500.000
		if metode != "credit" {
			return errors.New("credit harus dicicil")
		}
		if hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		// Jika tidak dicicil, metode tidak boleh credit
		if metode == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	return nil
}

func main() {
	fmt.Println(HitungHargaTotal(15000, 10000, 2))
	
}