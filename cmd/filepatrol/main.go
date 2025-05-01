package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"filepatrol/internal/patrol"
	"github.com/kardianos/service"
)

type program struct {
	configPath string
}

func (p *program) Start(s service.Service) error {
	go patrol.Start(p.configPath)
	return nil
}

func (p *program) Stop(s service.Service) error {
	patrol.Stop()
	return nil
}

func main() {
	fileInfo, err := os.Stat("../etc")
	if err == nil && fileInfo.IsDir() {
		_ = os.Chdir("../")
	}
	fileInfo, err = os.Stat("../../etc")
	if err == nil && fileInfo.IsDir() {
		_ = os.Chdir("../../")
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	log.Println("Starting service")
	log.Printf("Directory: %s", dir)

	configPath := flag.String("config", "etc/config.json", "Path to configuration file (default: etc/config.json)")
	fileInfo, err = os.Stat(*configPath)
	help := flag.Bool("help", false, "Show help")
	flag.Parse()
	if err != nil || !fileInfo.IsDir() {
		*help = true
	}

	if *help {
		showUsage()
		os.Exit(0)
	}

	log.Printf("Using config: %s", *configPath)

	svcConfig := &service.Config{
		Name:        "FilePatrol",
		DisplayName: "FilePatrol Service",
		Description: "Modular file monitoring and processing service.",
	}

	prg := &program{configPath: *configPath}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func showUsage() {
	fmt.Println("FilePatrol - Modular file watcher service")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  filepatrol [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -config <path>     Path to configuration file (default: etc/config.json)")
	fmt.Println("  -help              Show this help message")
}
