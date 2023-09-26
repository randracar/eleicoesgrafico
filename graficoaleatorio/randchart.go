package main

import (
	"time"
	"math/rand"
	"os"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateBarItems(xaxis []string) ([]opts.BarData, []int) {
	items := make([]opts.BarData, 0)
	values := make([]int, 0)
	for i := 0; i < len(xaxis); i++ {
		rn := rand.Intn(100)
		items = append(items, opts.BarData{Value: rn})
		values = append(values, rn)
	}
	return items, values
}

func generateBarItems2(xaxis []string, values []int) ([]opts.BarData) {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(xaxis); i++ {
		num := 100 - values[i]
		items = append(items, opts.BarData{Value: num})
	}
	return items
}

func returnEstados(regiao string) []string {
	if regiao == "Norte" {
		items := []string{"AM", "PA", "RO", "AC", "RR", "AP", "TO"}
		return items
	} else if regiao == "Nordeste" {
		items := []string{"BA", "PI", "MA", "CE", "RN", "PB", "PE", "AL", "SE"}
		return items
	} else if regiao == "Centro Oeste" {
		items := []string{"DF", "GO", "MT", "MS"}
		return items
	} else if regiao == "Sudeste" {
		items := []string{"SP", "RJ", "ES", "MG"}
		return items
	} else if regiao == "Sul" {
		items := []string{"PR", "SC", "RS"}
		return items
	}
	return []string{}
}

func createBar(xaxis []string) *charts.Bar {
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title:    "Gráfico eleições interativo"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80px"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show: true,
			Right: "50%", 
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Download image",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					Lang: []string{"data view", "turn off", "refresh"},
				},
			}},
		),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "inside",
			Start: 10,
			End:   50,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Estado",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Porcentagem de votos",
		}),
	)

	cand1items, values := generateBarItems(xaxis)
	cand2items := generateBarItems2(xaxis, values)
	
	bar.SetXAxis(xaxis).
		AddSeries("Candidato 1", cand1items).
		AddSeries("Candidato 2", cand2items)

	return bar
}

func main() {
	start := time.Now()

	axis := returnEstados("Nordeste")
	bar := createBar(axis)

	f, _ := os.Create("resultado.html")
	bar.Render(f)
	fmt.Println(time.Since(start))
}