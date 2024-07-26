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
	"time"
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
	Integer         bool
	IntervalSeconds float64
	Frequency       float64
	StartTime       time.Time
}

func (dg *Generator) Generate() map[string]any {
	var floatVal float64
	switch dg.FunctionType {
	case "random":
		floatVal = dg.generateRandomData()
	case "sin_wave":
		floatVal = dg.generateSinWaveData()
	case "sawtooth_wave":
		floatVal = dg.generateSawtoothWaveData()
	}

	var value any
	value = floatVal
	if dg.Integer {
		value = int64(math.Round(floatVal))
	}

	return map[string]any{
		dg.Attribute: value,
	}
}

func (dg *Generator) generateRandomData() float64 {
	rang := dg.MaxValue - dg.MinValue
	return rand.Float64()*rang + dg.MinValue
}

func (dg *Generator) generateSinWaveData() float64 {
	amplitude := (dg.MaxValue - dg.MinValue) / 2
	offset := dg.MinValue + amplitude
	return amplitude*math.Sin(2*math.Pi*dg.Frequency*dg.getElapsedTime()) + offset
}

func (dg *Generator) generateSawtoothWaveData() float64 {
	period := 1 / dg.Frequency
	fraction := dg.getElapsedTime() / period
	return dg.MinValue + (dg.MaxValue-dg.MinValue)*(fraction-math.Floor(fraction))
}

func (dg *Generator) getElapsedTime() float64 {
	return time.Since(dg.StartTime).Seconds()
}
