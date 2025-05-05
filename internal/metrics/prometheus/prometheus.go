package prometheus

import (
	"fmt"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var counterMetrics map[string]prometheus.Counter
var gaugeMetrics map[string]prometheus.Gauge
var histogramMetrics map[string]prometheus.Histogram
var summaryMetrics map[string]prometheus.Summary

func Init() {

	// Initialize the maps
	counterMetrics = make(map[string]prometheus.Counter)
	gaugeMetrics = make(map[string]prometheus.Gauge)
	histogramMetrics = make(map[string]prometheus.Histogram)
	summaryMetrics = make(map[string]prometheus.Summary)

	for i := range 5 {
		metricName := fmt.Sprintf("example_counter_%d", i)
		counterMetrics[metricName] = promauto.NewCounter(
			prometheus.CounterOpts{
				Name: metricName,
				Help: fmt.Sprintf("An example counter metric #%d", i),
			},
		)
	}

	for i := range 5 {
		metricName := fmt.Sprintf("example_gauge_%d", i)
		gaugeMetrics[metricName] = promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: metricName,
				Help: fmt.Sprintf("An example gauge metric #%d", i),
			},
		)
	}

	for i := range 5 {
		metricName := fmt.Sprintf("example_histogram_%d", i)
		histogramMetrics[metricName] = promauto.NewHistogram(
			prometheus.HistogramOpts{
				Name:    metricName,
				Help:    fmt.Sprintf("An example histogram metric #%d", i),
				Buckets: []float64{0.1, 0.2, 0.5, 1.0, 2.0},
			},
		)
	}

	for i := range 5 {
		metricName := fmt.Sprintf("example_summary_%d", i)
		summaryMetrics[metricName] = promauto.NewSummary(
			prometheus.SummaryOpts{
				Name:       metricName,
				Help:       fmt.Sprintf("An example summary metric #%d", i),
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, // Quantiles
			},
		)
	}

}

func ChangeValues() {
	for _, counter := range counterMetrics {

		// Generate a random number between 1 and 50
		randomNumber := rand.Intn(50) + 1

		counter.Add(float64(randomNumber))
	}

	for _, gauge := range gaugeMetrics {

		// Generate a random number between 1 and 100
		randomNumber := rand.Intn(100) + 1

		gauge.Set(float64(randomNumber))
	}

	for _, histogram := range histogramMetrics {
		// Generate a random number between 0.0 and 2.0
		randomNumber := rand.Float64()*1.9 + 0.0

		fmt.Println(randomNumber)

		histogram.Observe(randomNumber)
	}

	for _, summary := range summaryMetrics {
		// Generate a random number between 0.0 and 2.0
		randomNumber := rand.Float64()*1.9 + 0.0

		summary.Observe(randomNumber)
	}
}
