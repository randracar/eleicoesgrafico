package main

import (
	"time"
	"os"
	"fmt"
	"bufio"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"strconv"
)

func app() *charts.Bar{
	textScanner := bufio.NewScanner(os.Stdin)

	fmt.Println("----- GRAFICO INTERATIVO ELEICOES -----")
	fmt.Print("Insira o nome do candidato 1: ")
	textScanner.Scan()

	cand1nome := textScanner.Text()

	fmt.Println("")
	fmt.Print("Insira o nome do candidato 2: ")
	textScanner.Scan()

	cand2nome := textScanner.Text()

	fmt.Println("Qual região o gráfico deve ser sobre? ")
	fmt.Println("Norte")
	fmt.Println("Nordeste")
	fmt.Println("Centro Oeste")
	fmt.Println("Sudeste")
	fmt.Println("Sul")
	fmt.Print("Insira aqui a região: ")
	
	textScanner.Scan()

	reg := textScanner.Text()

	values := make([]int, 0)
	bar1data := make([]opts.BarData, 0)

	axis := returnEstados(reg)
	
	for a := 0; a < len(axis); a++ {
		fmt.Print("Quantos votos o candidato " + cand1nome + " teve em " + axis[a] + "? (em % dos votos): ")
		textScanner.Scan()
		convval, _ := strconv.Atoi(textScanner.Text())
		values = append(values, convval)
		bar1data = append(bar1data, opts.BarData{Value: convval})
	}

	bar2data := generateBarItems2(axis, values)
	
	bar := createBar(axis, cand1nome, cand2nome, bar1data, bar2data)
	return bar
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

func createBar(xaxis []string, cand1nome, cand2nome string, cand1items, cand2items []opts.BarData) *charts.Bar {
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
	
	bar.SetXAxis(xaxis).
		AddSeries(cand1nome, cand1items).
		AddSeries(cand2nome, cand2items)

	return bar
}

func main() {
	start := time.Now()

	bar := app()

	f, _ := os.Create("resultado.html")
	bar.Render(f)
	fmt.Println(time.Since(start))
}