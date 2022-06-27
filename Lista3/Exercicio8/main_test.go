package main

import "testing"

func TestSomaEMediaDeCincoVariaveis(t *testing.T) {
	type args struct {
		primeiroNumero int
		segundoNumero  int
		terceiroNumero int
		quartoNumero   int
		quintoNumero   int
	}
	tests := []struct {
		name      string
		args      args
		wantSoma  int
		wantMedia int
	}{
		{
			name: "Quando todas as variáveis forem 0",
			args: args{
				primeiroNumero: 0,
				segundoNumero:  0,
				terceiroNumero: 0,
				quartoNumero:   0,
				quintoNumero:   0,
			},
			wantSoma:  0,
			wantMedia: 0,
		},
		{
			name: "Quanto a média for 30",
			args: args{
				primeiroNumero: 30,
				segundoNumero:  30,
				terceiroNumero: 30,
				quartoNumero:   30,
				quintoNumero:   30,
			},
			wantSoma:  150,
			wantMedia: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SomaEMediaDeCincoVariaveis(tt.args.primeiroNumero, tt.args.segundoNumero, tt.args.terceiroNumero, tt.args.quartoNumero, tt.args.quintoNumero)
			if got != tt.wantSoma {
				t.Errorf("SomaEMediaDeCincoVariaveis() got = %v, want %v", got, tt.wantSoma)
			}
			if got1 != tt.wantMedia {
				t.Errorf("SomaEMediaDeCincoVariaveis() got1 = %v, want %v", got1, tt.wantMedia)
			}
		})
	}
}
