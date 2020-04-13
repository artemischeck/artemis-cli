package services

import (
	"fmt"
	"log"

	"github.com/opalmer/check-go-version/api"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

//Usage xxx
func Usage() {

	o, err := host.Info()
	if err != nil {
		fmt.Println("Could not retrieve OS details.")
		log.Fatal(err)
	}
	fmt.Printf("OS: %v Version %v \nOS Active Processes: %v \nOS Uptime: %v\n", o.OS, o.PlatformVersion, o.Procs, o.Uptime)

	m, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Could not retrieve RAM details.")
		log.Fatal(err)
	}
	fmt.Printf("RAM Total: %v \nRAM Available: %v\nRAM Used: %v\nRAM Used Percent:%f%%\n", m.Total, m.Available, m.Used, m.UsedPercent)

	d, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Could not retrieve disk details.")
		log.Fatal(err)
	}
	fmt.Printf("Disk Total: %v \nDisk Available: %v\nDisk Used: %v\nDisk Used Percent:%f%%\n", d.Total, d.Free, d.Used, d.UsedPercent)

	c, err := cpu.Info()
	if err != nil {
		fmt.Println("Could not retrieve CPU details.")
		log.Fatal(err)
	}
	fmt.Printf("CPU Model: %v \nCPU Cores: %v \n", c[0].ModelName, c[0].Cores)

	r, err := api.GetRunningVersion()
	if err != nil {
		fmt.Println("Could not retrieve Go installation details.")
		log.Fatal(err)
	}
	fmt.Printf("Go Version: %s\nGo Platform: %s\nGo Architecture: %s\n\n", r.Version, r.Platform, r.Architecture)

}
