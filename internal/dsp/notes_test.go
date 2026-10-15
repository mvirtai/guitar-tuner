package dsp_test

import (
	"math"
	"testing"

	"guitar_tuner/internal/dsp"
)

func TestFrequencyToNote_StandardPitches(t *testing.T) {
	tests := []struct {
		name        string
		freq        float64
		wantName    string
		wantOctave  int
		wantDisplay string
		wantTarget  float64
		wantInTune  bool
	}{
		{
			name:        "Concert A4",
			freq:        440.0,
			wantName:    "A",
			wantOctave:  4,
			wantDisplay: "A4",
			wantTarget:  440.0,
			wantInTune:  true,
		},
		{
			name:        "Guitar Low E2",
			freq:        82.4069,
			wantName:    "E",
			wantOctave:  2,
			wantDisplay: "E2",
			wantTarget:  82.4069,
			wantInTune:  true,
		},
		{
			name:        "Guitar A2",
			freq:        110.0,
			wantName:    "A",
			wantOctave:  2,
			wantDisplay: "A2",
			wantTarget:  110.0,
			wantInTune:  true,
		},
		{
			name:        "Guitar D3",
			freq:        146.8324,
			wantName:    "D",
			wantOctave:  3,
			wantDisplay: "D3",
			wantTarget:  146.8324,
			wantInTune:  true,
		},
		{
			name:        "Guitar G3",
			freq:        195.9977,
			wantName:    "G",
			wantOctave:  3,
			wantDisplay: "G3",
			wantTarget:  195.9977,
			wantInTune:  true,
		},
		{
			name:        "Guitar B3",
			freq:        246.9417,
			wantName:    "B",
			wantOctave:  3,
			wantDisplay: "B3",
			wantTarget:  246.9417,
			wantInTune:  true,
		},
		{
			name:        "Guitar High E4",
			freq:        329.6276,
			wantName:    "E",
			wantOctave:  4,
			wantDisplay: "E4",
			wantTarget:  329.6276,
			wantInTune:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := dsp.FrequencyToNote(tc.freq)

			if res.NoteName != tc.wantName {
				t.Errorf("NoteName: got %q, want %q", res.NoteName, tc.wantName)
			}

			if res.Octave != tc.wantOctave {
				t.Errorf("Octave: got %d, want %d", res.Octave, tc.wantOctave)
			}
			if res.DisplayName != tc.wantDisplay {
				t.Errorf("DisplayName: got %q, want %q", res.DisplayName, tc.wantDisplay)
			}
			if math.Abs(res.TargetFreq-tc.wantTarget) > 0.01 {
				t.Errorf("TargetFreq: got %.4f, want %.4f", res.TargetFreq, tc.wantTarget)
			}
			if math.Abs(res.Cents) > 0.01 {
				t.Errorf("Cents: got %.4f, want 0.0", res.Cents)
			}
			if res.InTune != tc.wantInTune {
				t.Errorf("InTune: got %v, want %v", res.InTune, tc.wantInTune)
			}
		})
	}
}

func TestFrequencyToNote_Detuned(t *testing.T) {
	tests := []struct {
		name       string
		freq       float64
		wantCents  float64
		wantInTune bool
	}{
		{
			name:       "A4 slightly sharp (+9.8 cents)",
			freq:       442.5,
			wantCents:  9.80,
			wantInTune: false,
		},
		{
			name:       "A4 slightly flat (-9.9 cents)",
			freq:       437.5,
			wantCents:  -9.88,
			wantInTune: false,
		},
		{
			name:       "A4 within tolerance (+2.0 cents)",
			freq:       440.51,
			wantCents:  2.00,
			wantInTune: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := dsp.FrequencyToNote(tc.freq)
			if math.Abs(res.Cents-tc.wantCents) > 0.1 {
				t.Errorf("Cents: got %.2f, want %.2f", res.Cents, tc.wantCents)
			}
			if res.InTune != tc.wantInTune {
				t.Errorf("InTune: got %v, want %v", res.InTune, tc.wantInTune)
			}
		})
	}
}

func TestFrequencyToNote_EdgeCases(t *testing.T) {
	edgeCases := []struct {
		name string
		freq float64
	}{
		{"Zero frequency", 0.0},
		{"Negative frequency", -50.0},
		{"NaN", math.NaN()},
		{"Positive Infinity", math.Inf(1)},
		{"Sub-audible tiny frequency", 0.001},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			res := dsp.FrequencyToNote(tc.freq)
			// Special handling for NaN comparison
			if math.IsNaN(tc.freq) {
				if !math.IsNaN(res.Frequency) {
					t.Errorf("Frequency not preserved for NaN: got %v, want NaN", res.Frequency)
				}
			} else {
				if res.Frequency != tc.freq {
					t.Errorf("Frequency not preserved: got %v, want %v", res.Frequency, tc.freq)
				}
			}
			if res.InTune {
				t.Errorf("InTune should be false for edge case frequency %v", tc.freq)
			}
			if res.ClosestString != nil {
				t.Errorf("ClosestString should be nil for invalid frequency %v", tc.freq)
			}
		})
	}
}

func TestClosestGuitarString(t *testing.T) {
	tests := []struct {
		name       string
		freq       float64
		wantString int // 1-6 or 0 for nil
	}{
		{"Exact Low E2 (82.41 Hz)", 82.4069, 6},
		{"Flat Low E2 (80.0 Hz)", 80.0, 6},
		{"Sharp High E4 (335.0 Hz)", 335.0, 1},
		{"Exact A2 (110.0 Hz)", 110.0, 5},
		{"D3 near 146 Hz", 146.0, 4},
		{"Way too low (20 Hz)", 20.0, 0},
		{"Way too high (1000 Hz)", 1000.0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matched := dsp.ClosestGuitarString(tc.freq)
			if tc.wantString == 0 {
				if matched != nil {
					t.Errorf("Expected nil string, got %v", matched)
				}
			} else {
				if matched == nil {
					t.Fatalf("Expected string %d, got nil", tc.wantString)
				}
				if matched.Number != tc.wantString {
					t.Errorf("String number: got %d, want %d", matched.Number, tc.wantString)
				}
			}
		})
	}
}
