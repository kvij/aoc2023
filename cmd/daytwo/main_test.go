package main

import "testing"

func Test_parseGameLogLine(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name   string
		args   args
		wantId int
		wantR  int
		wantG  int
		wantB  int
	}{
		{
			args:   struct{ l string }{l: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 30 green`},
			wantId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotR, gotG, gotB := parseGameLogLine(tt.args.l)
			if gotId != tt.wantId {
				t.Errorf("parseGameLogLine() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotR != tt.wantR {
				t.Errorf("parseGameLogLine() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("parseGameLogLine() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("parseGameLogLine() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
