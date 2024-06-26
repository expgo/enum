package example

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type makeTest struct {
	M Make `json:"make"`
}

func TestMake(t *testing.T) {
	assert.Equal(t, "Ford", MakeFord.String())
	assert.Equal(t, "Make(99).Name", Make(99).String())
	ford := MakeFord
	assert.Implements(t, (*flag.Value)(nil), &ford)
	assert.Implements(t, (*flag.Getter)(nil), &ford)

	values := MakeValues()
	assert.Len(t, values, 11)
}

var makeTests = []struct {
	name          string
	input         string
	output        *makeTest
	errorExpected bool
	err           error
	caseChanged   bool
}{
	{
		name:          "toyota",
		input:         `{"make":"Toyota"}`,
		output:        &makeTest{M: MakeToyota},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "chevy",
		input:         `{"make":"Chevy"}`,
		output:        &makeTest{M: MakeChevy},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "ford",
		input:         `{"make":"Ford"}`,
		output:        &makeTest{M: MakeFord},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "tesla",
		input:         `{"make":"Tesla"}`,
		output:        &makeTest{M: MakeTesla},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "hyundai",
		input:         `{"make":"Hyundai"}`,
		output:        &makeTest{M: MakeHyundai},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "nissan",
		input:         `{"make":"Nissan"}`,
		output:        &makeTest{M: MakeNissan},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "jaguar",
		input:         `{"make":"Jaguar"}`,
		output:        &makeTest{M: MakeJaguar},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "audi",
		input:         `{"make":"Audi"}`,
		output:        &makeTest{M: MakeAudi},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "AUDI",
		input:         `{"make":"AUDI"}`,
		output:        &makeTest{M: MakeAudi},
		errorExpected: false,
		err:           nil,
		caseChanged:   true,
	},
	{
		name:          "bmw",
		input:         `{"make":"BMW"}`,
		output:        &makeTest{M: MakeBmw}, // TODO fix me
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "mercedes",
		input:         `{"make":"Mercedes_Benz"}`,
		output:        &makeTest{M: MakeMercedesBenz},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "volkswagon",
		input:         `{"make":"Volkswagon"}`,
		output:        &makeTest{M: MakeVolkswagon},
		errorExpected: false,
		err:           nil,
	},
	{
		name:          "porsche",
		input:         `{"make":"Porsche"}`,
		output:        &makeTest{M: MakeVolkswagon},
		errorExpected: true,
		err:           errors.New("Porsche is not a valid Make, try [Toyota, Chevy, Ford, Tesla, Hyundai, Nissan, Jaguar, Audi, BMW, Mercedes_Benz, Volkswagon]"),
	},
}

func TestMakeUnmarshal(t *testing.T) {
	for _, test := range makeTests {
		t.Run(test.name, func(tt *testing.T) {
			x := &makeTest{}
			err := json.Unmarshal([]byte(test.input), x)
			if !test.errorExpected {
				require.NoError(tt, err, "failed unmarshalling the json.")
				assert.Equal(tt, test.output.M, x.M)

				// Values won't be exactly the same, so we just validate that it was unmarshalled correctly.
				if !test.caseChanged {
					// Marshal back
					raw, err := json.Marshal(test.output)
					require.NoError(tt, err, "failed marshalling back to json")
					require.JSONEq(tt, test.input, string(raw), "json didn't match")
				}
			} else {
				require.Error(tt, err)
				assert.EqualError(tt, err, test.err.Error())
			}
		})
	}
}

func TestFlagInterface(t *testing.T) {
	for _, test := range makeTests {
		t.Run(test.name, func(tt *testing.T) {
			if test.output != nil {
				var tmp Make
				m := &tmp
				require.Equal(tt, Make(0), m.Get(), "Unset value should be default")
				require.NoError(tt, m.Set(test.output.M.String()), "failed setting flag value")
			}
		})
	}
}

func TestNoZeroValues(t *testing.T) {
	t.Run("base", func(tt *testing.T) {
		assert.Equal(tt, 20, int(NoZerosStart))
		assert.Equal(tt, 21, int(NoZerosMiddle))
		assert.Equal(tt, 22, int(NoZerosEnd))
		assert.Equal(tt, 23, int(NoZerosPs))
		assert.Equal(tt, 24, int(NoZerosPps))
		assert.Equal(tt, 25, int(NoZerosPpps))
		assert.Equal(tt, "ppps", NoZerosPpps.String())
		assert.Equal(tt, "NoZeros(4).Name", NoZeros(4).String())

		assert.Len(tt, NoZerosValues(), 6)

		_, err := ParseNoZeros("pppps")
		assert.EqualError(tt, err, "pppps is not a valid NoZeros, try [start, middle, end, ps, pps, ppps]")

		tmp, _ := ParseNoZeros("ppps")
		assert.Equal(tt, NoZerosPpps, tmp)

		tmp, _ = ParseNoZeros("PppS")
		assert.Equal(tt, NoZerosPpps, tmp)

		val := map[string]*NoZeros{}

		err = json.Unmarshal([]byte(`{"nz":"pppps"}`), &val)
		assert.EqualError(tt, err, "pppps is not a valid NoZeros, try [start, middle, end, ps, pps, ppps]")
	})

	for _, value := range NoZerosValues() {
		t.Run(value.Name(), func(tt *testing.T) {
			val := map[string]*NoZeros{}

			rawJSON := fmt.Sprintf(`{"val":"%s"}`, value.Name())

			require.NoError(tt, json.Unmarshal([]byte(rawJSON), &val), "Failed unmarshalling no zero")
			marshalled, err := json.Marshal(val)
			require.NoError(tt, err, "failed marshalling back to json")
			require.JSONEq(tt, rawJSON, string(marshalled), "marshalled json did not match")

			// Flag
			var tmp NoZeros
			nz := &tmp
			require.Equal(tt, NoZeros(0), nz.Get(), "Unset value should be default")
			require.NoError(tt, nz.Set(value.Name()), "failed setting flag value")
		})
	}
}

func BenchmarkMakeParse(b *testing.B) {
	knownItems := map[string]struct {
		input  string
		output Make
		err    error
	}{
		"cased lookup": {
			input:  MakeAudi.String(),
			output: MakeAudi,
		},
		"lowercase lookup": {
			input:  strings.ToLower(MakeVolkswagon.String()),
			output: MakeVolkswagon,
		},
		"hyphenated upper": {
			input:  strings.ToUpper(MakeMercedesBenz.String()),
			output: MakeMercedesBenz,
		},
		// Leave this in to add an int as string parsing option in future.
		// "numeric": {
		// 	input:  "2",
		// 	output: MakeChevy,
		// },
		// "last numeric": {
		// 	input:  "20",
		// 	output: MakeVolkswagon,
		// },
		"failure": {
			input:  "xyz",
			output: MakeToyota,
			err:    errors.New("xyz is not a valid Make, try [Toyota, Chevy, Ford, Tesla, Hyundai, Nissan, Jaguar, Audi, BMW, Mercedes_Benz, Volkswagon]"),
		},
	}

	for name, tc := range knownItems {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				out, err := ParseMake(tc.input)
				assert.Equal(b, tc.output, out)
				if tc.err != nil {
					assert.Error(b, err)
				} else {
					assert.NoError(b, err)
				}
			}
		})
	}
}
