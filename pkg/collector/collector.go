package collector

import (
	"log"
	"os"
	"time"

	g "github.com/gosnmp/gosnmp"
	"github.com/prometheus/client_golang/prometheus"
)

type trxMetrics struct {
	CpuIdle         prometheus.Gauge
	RamTotal        prometheus.Gauge
	RamAvailable    prometheus.Gauge
	FirmwareVersion prometheus.Gauge
	SerialNumber    prometheus.Gauge
}

func new() trxMetrics {
	trx := trxMetrics{
		CpuIdle: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_cpu_idle",
				Help: "CPU IDLE of KLAS TRX Server",
			},
		),
		RamTotal: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_ram_total",
				Help: "Total RAM of KLAS TRX Server",
			},
		),
		RamAvailable: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_ram_available",
				Help: "Total RAM of KLAS TRX Server",
			},
		),
		// FirmwareVersion: prometheus.NewGauge(
		// 	prometheus.GaugeOpts{
		// 		Name: "trx_firmware_version",
		// 		Help: "Firmware Version of KLAS TRX Server",
		// 	},
		// ),
		// SerialNumber: prometheus.NewGauge(
		// 	prometheus.GaugeOpts{
		// 		Name: "trx_serial_number",
		// 		Help: "Serial Number of KLAS TRX Server",
		// 	},
		// ),
	}
	// Register metrics with prometheus
	prometheus.MustRegister(trx.CpuIdle)
	prometheus.MustRegister(trx.RamTotal)
	prometheus.MustRegister(trx.RamAvailable)
	// prometheus.MustRegister(trx.FirmwareVersion)
	// prometheus.MustRegister(trx.SerialNumber)
	return trx
}

func Collect() {

	// Get TRX IP from env variable $TRX_MGMT_IP
	ipAddress := os.Getenv("TRX_MGMT_IP")

	// Validate if IP is set
	if len(ipAddress) == 0 {
		log.Fatal("TRX_MGMT_IP env var not set")
		os.Exit(1)
	}

	// Get TRX IP from env variable $TRX_MGMT_IP
	snmpCommunity := os.Getenv("TRX_SNMP_COMMUNITY")

	// Validate if IP is set
	if len(snmpCommunity) == 0 {
		log.Fatal("TRX_SNMP_COMMUNITY env var not set")
		os.Exit(1)
	}

	serverMetrics := new()

	// Update values every 30s
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {

		// SNMP GET
		g.Default.Target = ipAddress
		g.Default.Community = snmpCommunity
		err := g.Default.Connect()
		if err != nil {
			log.Fatalf("Connect() err: %v", err)
		}
		defer g.Default.Conn.Close()

		oids := []string{".1.3.6.1.4.1.2021.11.53.0", "1.3.6.1.4.1.2021.4.5.0", "1.3.6.1.4.1.2021.4.6.0"}

		result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
		if err2 != nil {
			log.Fatalf("Get() err: %v", err2)
		}

		// Set prometheus metrics
		serverMetrics.CpuIdle.Set(float64(result.Variables[0].Value.(uint)))
		serverMetrics.RamTotal.Set(float64(result.Variables[1].Value.(uint)))
		serverMetrics.RamAvailable.Set(float64(result.Variables[2].Value.(uint)))
		// serverMetrics.FirmwareVersion.Set(3)
		// serverMetrics.SerialNumber.Set(4)

	}

}
