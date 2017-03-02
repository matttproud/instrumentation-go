package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/instrumentation-go/stats"
	"github.com/google/instrumentation-go/stats/tagging"
)

func main() {
	// ---------------------------------------MEASURES---------------------------
	// Create/register measure descriptions of type string
	muString := &stats.MeasurementUnit{}
	mDescString := stats.NewMeasureDescString("SystemLoad", "At any time its value can be: HIGH, MEDIUM, LOW", muString)
	if err := stats.RegisterMeasureDesc(mDescString); err != nil {
		fmt.Printf("Error measure registration mDescString: %v\n", err)
	}

	// Create/register measure descriptions of type bool
	muBool := &stats.MeasurementUnit{}
	mDescBool := stats.NewMeasureDescBool("SystemHot", "At any time its value can be: TRUE, FALSE", muBool)
	if err := stats.RegisterMeasureDesc(mDescBool); err != nil {
		fmt.Printf("Error measure registration mDescmuBool: %v\n", err)
	}

	// Create/register measure descriptions of type float64
	muFloat64 := &stats.MeasurementUnit{
		Power10: 6,
		Numerators: []stats.BasicUnit{
			stats.BytesUnit,
		},
	}
	mDescFloat64 := stats.NewMeasureDescFloat64("DiskRead", "Read MBs", muFloat64)
	if err := stats.RegisterMeasureDesc(mDescFloat64); err != nil {
		fmt.Printf("Error measure registration mDescFloat64: %v\n", err)
	}

	// Create/register measure descriptions of type Int64
	muInt64 := &stats.MeasurementUnit{
		Power10: 6,
		Numerators: []stats.BasicUnit{
			stats.BytesUnit,
		},
	}
	mDescInt64 := stats.NewMeasureDescInt64("ConnectedUsers", "Connected users. It is for a snapshot", muInt64)
	if err := stats.RegisterMeasureDesc(mDescInt64); err != nil {
		fmt.Printf("Error measure registration mDescInt64: %v\n", err)
	}

	// ---------------------------------------VIEWS---------------------------
	// Creates keys/views
	mgr := tagging.DefaultKeyManager()
	k1, err := mgr.CreateKeyInt64("key1")
	if err != nil {
		panic(fmt.Sprintf("Key k1 not created %v", err))
	}
	k2, err := mgr.CreateKeyInt64("key2")
	if err != nil {
		panic(fmt.Sprintf("Key k2 not created %v", err))
	}
	k3, err := mgr.CreateKeyStringUTF8("key3")
	if err != nil {
		panic(fmt.Sprintf("Key k3 not created %v", err))
	}

	// Create/register view description of type gauge float
	c := make(chan *stats.View, 256)

	vwGaugeS := &stats.GaugeStringViewDesc{
		Vdc: &stats.ViewDescCommon{
			Name:            "View1",
			Description:     "",
			MeasureDescName: "SystemLoad",
			TagKeys:         []tagging.Key{k1, k2, k3},
		},
	}
	if err := stats.RegisterViewDesc(vwGaugeS, c); err != nil {
		fmt.Printf("Error view registration: %v\n", err)
	}

	vwGaugeB := &stats.GaugeBoolViewDesc{
		Vdc: &stats.ViewDescCommon{
			Name:            "View2",
			Description:     "",
			MeasureDescName: "SystemHot",
			TagKeys:         []tagging.Key{k1, k2, k3},
		},
	}
	if err := stats.RegisterViewDesc(vwGaugeB, c); err != nil {
		fmt.Printf("Error view registration: %v\n", err)
	}

	vwGaugeF := &stats.GaugeFloat64ViewDesc{
		Vdc: &stats.ViewDescCommon{
			Name:            "View3",
			Description:     "",
			MeasureDescName: "DiskRead",
			TagKeys:         []tagging.Key{k1, k2, k3},
		},
	}
	if err := stats.RegisterViewDesc(vwGaugeF, c); err != nil {
		fmt.Printf("Error view registration: %v\n", err)
	}

	vwGaugeI := &stats.GaugeInt64ViewDesc{
		Vdc: &stats.ViewDescCommon{
			Name:            "View4",
			Description:     "",
			MeasureDescName: "ConnectedUsers",
			TagKeys:         []tagging.Key{k1, k2, k3},
		},
	}
	if err := stats.RegisterViewDesc(vwGaugeI, c); err != nil {
		fmt.Printf("Error view registration: %v\n", err)
	}

	vwDist := &stats.DistributionViewDesc{
		Vdc: &stats.ViewDescCommon{
			Name:            "View5",
			Description:     "",
			MeasureDescName: "DiskRead",
			TagKeys:         []tagging.Key{k1, k2, k3},
		},
		Bounds: []float64{0, 10, 100},
	}
	if err := stats.RegisterViewDesc(vwDist, c); err != nil {
		fmt.Printf("Error view registration: %v\n", err)
	}

	vwInterval := &stats.IntervalViewDesc{
		Vdc: &stats.ViewDescCommon{
			Name:            "View6",
			Description:     "",
			MeasureDescName: "DiskRead",
			TagKeys:         []tagging.Key{k1, k2, k3},
		},
		SubIntervals: 5,
		Intervals:    []time.Duration{time.Second * 5, time.Second * 30, time.Second * 60},
	}
	if err := stats.RegisterViewDesc(vwInterval, c); err != nil {
		fmt.Printf("Error view registration: %v\n", err)
	}
	// ---------------------------------------CREATES TAGS---------------------------
	// Set tags values in mutations
	mut1 := k1.CreateMutation(1, tagging.BehaviorAdd)
	mut2 := k2.CreateMutation(1, tagging.BehaviorAdd)
	mut3 := k3.CreateMutation("s1", tagging.BehaviorAdd)
	// ---------------------------------------RECORD USAGE---------------------------
	ctx := tagging.NewContextWithMutations(context.Background(), mut1, mut2, mut3)

	//...
	// DoStuff()
	//...
	// Creates few measurements
	m1 := mDescString.CreateMeasurement("HIGH")
	m2 := mDescBool.CreateMeasurement(true)
	m3 := mDescFloat64.CreateMeasurement(209)
	m4 := mDescFloat64.CreateMeasurement(9)
	m5 := mDescInt64.CreateMeasurement(20)
	m6 := mDescInt64.CreateMeasurement(40)
	stats.RecordMeasurements(ctx, m1, m2, m3, m4, m5, m6)

	// ---------------------------------------CREATES TAGS---------------------------
	// Set tags values in mutations
	mut1 = k1.CreateMutation(2, tagging.BehaviorAdd)
	mut2 = k2.CreateMutation(2, tagging.BehaviorAdd)
	mut3 = k3.CreateMutation("s2", tagging.BehaviorAdd)
	// ---------------------------------------RECORD USAGE---------------------------
	ctx = tagging.NewContextWithMutations(context.Background(), mut1, mut2, mut3)

	//...
	// DoStuff()
	//...
	// Create context with mutations
	// Creates few measurements
	m1 = mDescString.CreateMeasurement("LOW")
	m2 = mDescBool.CreateMeasurement(false)
	m3 = mDescFloat64.CreateMeasurement(218)
	m4 = mDescFloat64.CreateMeasurement(18)
	m5 = mDescInt64.CreateMeasurement(5)
	m6 = mDescInt64.CreateMeasurement(10)
	stats.RecordMeasurements(ctx, m1, m2, m3, m4, m5, m6)

	// ---------------------------------------RETRIEVE USAGE---------------------------
	/*for i, name := range []string{"View1", "View2"} {
		view, err := stats.RetrieveViewByName(name)
		if err != nil {
			fmt.Printf("%v -->\nError view retrieve: %v\n", i+1, err)
			continue
		}
		fmt.Printf("%v -->\n%v\n", i+1, view)
	}*/

	// Retrieve Views
	done := make(chan bool)
	go func(c chan *stats.View) {
		i := 0
		for {
			i++
			v := <-c
			fmt.Printf("%v -->\n%v\n", i, v)
		}
		done <- true
	}(c)
	<-done

	/*y := mDesc1.Meta()
	fmt.Printf("Hello, world:\n%v\n%v\n", m1_1, y)
	*/
}