package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type trxMetrics struct {
	CpuUsage        prometheus.Gauge
	RamUsage        prometheus.Gauge
	FirmwareVersion prometheus.Gauge
	SerialNumber    prometheus.Gauge
}

func new() trxMetrics {
	trx = trxMetrics{
		CpuUsage: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_cpu_usage",
				Help: "CPU Usage of KLAS TRX Server",
			},
		),
		RamUsage: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_ram_usage",
				Help: "RAM Usage of KLAS TRX Server",
			},
		),
		FirmwareVersion: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_firmware_version",
				Help: "Firmware Version of KLAS TRX Server",
			},
		),
		SerialNumber: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "trx_serial_number",
				Help: "Serial Number of KLAS TRX Server",
			},
		),
	}
}

// Register metrics with prometheus
prometheus.MustRegister(trx.CpuUsage)
prometheus.MustRegister(trx.RamUsage)
prometheus.MustRegister(trx.FirmwareVersion)
prometheus.MustRegister(trx.SerialNumber)

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

	// Update values every 300s
	ticker := time.NewTicker(300 * time.Second)
	for range ticker.C {

		// SNMP GET

		// Set prometheus metrics
		serverMetrics.CpuUsage.Set()
		serverMetrics.RamUsage.Set()
		serverMetrics.FirmwareVersion.Set()
		serverMetrics.SerialNumber.Set()

	}

}

func convertToFloat(input [byte]) float64 {

	//Trim the quotation marks from the string and parse the result as a float
	floatValue, err := strconv.ParseFloat(strings.Trim(string(input), `"`), 64)

	if err != nil {
		     log.Fatal(err)
	}
	return floatValue
}