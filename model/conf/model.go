//  This file is part of the Eliona project.
//  Copyright © 2024 IoTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package confmodel

import (
	"math"
	"math/rand"
)

type Generator struct {
	Id              int32
	AssetId         int32
	Attribute       string
	Subtype         string
	AssetType       string
	FunctionType    string
	MinValue        float64
	MaxValue        float64
	IntervalSeconds int32
	Frequency       float64
	ElapsedTime     float64
}

func (dg *Generator) Generate() map[string]any {
	var value any
	switch dg.FunctionType {
	case "boolean":
		value = dg.generateBooleanData()
	case "sin_wave":
		value = dg.generateSinWaveData()
	case "sawtooth_wave":
		value = dg.generateSawtoothWaveData()
	}

	dg.ElapsedTime += float64(dg.IntervalSeconds)
	return map[string]any{
		dg.Attribute: value,
	}
}

func (dg *Generator) generateBooleanData() bool {
	return rand.Intn(2) == 0
}

func (dg *Generator) generateSinWaveData() float64 {
	amplitude := (dg.MaxValue - dg.MinValue) / 2
	offset := dg.MinValue + amplitude
	return amplitude*math.Sin(2*math.Pi*dg.Frequency*dg.ElapsedTime) + offset
}

func (dg *Generator) generateSawtoothWaveData() float64 {
	period := 1 / dg.Frequency
	fraction := dg.ElapsedTime / period
	return dg.MinValue + (dg.MaxValue-dg.MinValue)*(fraction-math.Floor(fraction))
}
