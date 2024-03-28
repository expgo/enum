package enum

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type animalData struct {
	AnimalX Animal `json:"animal"`
}

func TestAnimalString(t *testing.T) {
	x := Animal(109)
	assert.Equal(t, "Animal(109).Name", x.String())
	x = Animal(1)
	assert.Equal(t, "Dog", x.String())

	y, err := ParseAnimal("Cat")
	require.NoError(t, err, "Failed parsing cat")
	assert.Equal(t, AnimalCat, y)

	z, err := ParseAnimal("Snake")
	require.Error(t, err, "Shouldn't parse a snake")
	assert.Equal(t, Animal(0), z)
}

func TestAnimalUnmarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		output        *animalData
		errorExpected bool
		err           error
	}{
		{
			name:          "cat",
			input:         `{"animal":0}`,
			output:        &animalData{AnimalX: AnimalCat},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "dog",
			input:         `{"animal":1}`,
			output:        &animalData{AnimalX: AnimalDog},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "fish",
			input:         `{"animal":2}`,
			output:        &animalData{AnimalX: AnimalFish},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "notananimal",
			input:         `{"animal":22}`,
			output:        &animalData{AnimalX: Animal(22)},
			errorExpected: false,
			err:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			x := &animalData{}
			err := json.Unmarshal([]byte(test.input), x)
			if !test.errorExpected {
				require.NoError(tt, err, "failed unmarshalling the json.")
				assert.Equal(tt, test.output.AnimalX, x.AnimalX)
			} else {
				require.Error(tt, err)
				assert.EqualError(tt, err, test.err.Error())
			}
		})
	}
}

func TestAnimalMarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         *animalData
		output        string
		errorExpected bool
		err           error
	}{
		{
			name:          "cat",
			output:        `{"animal":0}`,
			input:         &animalData{AnimalX: AnimalCat},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "dog",
			output:        `{"animal":1}`,
			input:         &animalData{AnimalX: AnimalDog},
			errorExpected: false,
			err:           nil,
		},
		{
			name:          "fish",
			output:        `{"animal":2}`,
			input:         &animalData{AnimalX: AnimalFish},
			errorExpected: false,
			err:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			raw, err := json.Marshal(test.input)
			require.NoError(tt, err, "failed marshalling to json")
			assert.JSONEq(tt, test.output, string(raw))
		})
	}
}

func TestAnimal_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		animal Animal
		want   bool
	}{
		{
			name:   "Valid Animal Enum",
			animal: Animal(1), // Assuming 1 is valid
			want:   true,
		},
		{
			name:   "Invalid Animal Enum",
			animal: Animal(100), // Assuming 100 is not valid
			want:   false,
		},
		{
			name:   "Edge case Animal Enum",
			animal: Animal(-1), // Assuming negative numbers are not valid
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.animal.IsValid(); got != tt.want {
				t.Errorf("Animal.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnimalName(t *testing.T) {
	testCases := []struct {
		desc string
		in   Animal
		out  string
	}{
		{
			desc: "Valid Animal",
			in:   AnimalDog,
			out:  "Dog",
		},
		{
			desc: "Invalid Animal",
			in:   Animal(100),
			out:  "Animal(100).Name",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && tC.out != "" {
					t.Errorf("The code panicked %v", r)
				}
			}()
			got := tC.in.Name()
			if got != tC.out {
				t.Errorf("Expected Animal name %v, got %v", tC.out, got)
			}
		})
	}
}
