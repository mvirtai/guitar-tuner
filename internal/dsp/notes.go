package dsp

import (
	"fmt"
	"math"
)

// A4ConcertPitch defines the standard reference pitch for A4 in Hz.
const A4ConcertPitch = 440.0

// DefaultInTuneTolerance defines the acceptable deviation in cents for in-tune idication.
const DefaultInTuneTolerance = 3.0

// GuitarString represents a standard guitar string.
type GuitarString struct {
	Number int     // 1 to 6 (1 = 'high e', 6 = 'low E')
	Name   string  //e.g. "E2", "A2"
	Target float64 // frequency in Hz (hertz)
}

// TunerMode holds the result of frequency analysis against musical tuning.
type TunerNote struct {
	Frequency     float64       // Measured frequency in Hz
	TargetFreq    float64       // Ideal target frequency in Hz
	NoteName      string        // Note name without octave (e.g. "A", "C#")
	Octave        int           // Octave number (e.g. 2, 4)
	DisplayName   string        // Full name (e.g. "A2", "C#4")
	Cents         float64       // Deviation in centers (-50.0 to +50.0)
	InTune        bool          // True if with tolerance (e.g. +/- 3 cents from target)
	ClosestString *GuitarString // Matched standard guitar string (or nil if non-guitar pitch)
}

// StandardTuning lists standard guitar strings ordered from lowest to highest.
var StandardTuning = []GuitarString{
	{Number: 6, Name: "E2", Target: 82.4069},
	{Number: 5, Name: "A2", Target: 110.0000},
	{Number: 4, Name: "D3", Target: 146.8324},
	{Number: 3, Name: "G3", Target: 195.9977},
	{Number: 2, Name: "B3", Target: 246.9417},
	{Number: 1, Name: "E4", Target: 329.6276},
}

// CentsDeviation calculates the difference in cents between freq and targetFreq.
func CentsDeviation(freq, targetFreq float64) float64 {
	if freq <= 0 || targetFreq <= 0 || math.IsNaN(freq) || math.IsNaN(targetFreq) {
		return 0.0
	}
	return 1200.0 * math.Log2(freq/targetFreq)
}

// ClosestGuitarString finds the nearest standard guitar string for a given frequency.
// Returns nil if the frequency is outside reasonable guitar range (60 Hz - 450 Hz)
// or further than 150 cents from any string.
func ClosestGuitarString(freq float64) *GuitarString {
	if freq < 60.0 || freq > 450.0 || math.IsNaN(freq) || math.IsInf(freq, 0) {
		return nil
	}

	var bestString *GuitarString
	minCentsAbs := math.MaxFloat64

	for i := range StandardTuning {
		str := &StandardTuning[i]
		centsDiff := math.Abs(CentsDeviation(freq, str.Target))
		if centsDiff < minCentsAbs {
			minCentsAbs = centsDiff
			bestString = str
		}
	}

	// Match only if within 150 cents of the string
	if minCentsAbs <= 150.0 {
		return bestString
	}
	return nil
}

// FrequencyToNote converts a measured frequency to its closest musical note,
// including octave, cents deviation, and guitar string matching.
func FrequencyToNote(freq float64) TunerNote {
	// Handle edge cases
	if freq <= 0 || math.IsNaN(freq) || math.IsInf(freq, 0) {
		return TunerNote{
			Frequency:     freq,
			NoteName:      "",
			Octave:        0,
			DisplayName:   "",
			Cents:         0.0,
			InTune:        false,
			ClosestString: nil,
		}
	}

	// Calculate exact MIDI value
	exactMIDI := 69.0 + 12.0*math.Log2(freq/A4ConcertPitch)
	roundMIDI := int(math.Round(exactMIDI))

	// Note names in 12-tone equal temperament
	noteNames := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	// Calculate note index and octave
	noteIndex := ((roundMIDI % 12) + 12) % 12
	octave := (roundMIDI / 12) - 1
	noteName := noteNames[noteIndex]
	displayName := fmt.Sprintf("%s%d", noteName, octave)

	// Calculate target frequency and cents deviation
	targetFreq := A4ConcertPitch * math.Pow(2.0, float64(roundMIDI-69)/12.0)
	cents := CentsDeviation(freq, targetFreq)
	inTune := math.Abs(cents) <= DefaultInTuneTolerance

	return TunerNote{
		Frequency:     freq,
		TargetFreq:    targetFreq,
		NoteName:      noteName,
		Octave:        octave,
		DisplayName:   displayName,
		Cents:         cents,
		InTune:        inTune,
		ClosestString: ClosestGuitarString(freq),
	}
}
