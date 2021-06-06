package app

import (
	"fmt"
	"log"
	"plotcarrier/remote"
	"time"

	"github.com/spf13/viper"
)

type MachineCfg struct {
	IP          string `yaml:"ip"`
	BindAddress string `yaml:"bindAddress"`
	Src         string `yaml:"remote_src"`
	Dst         string `yaml:"local_dst"`
}

func RemoteStart() error {
	hostName := viper.GetString("host.username")
	keyPath := viper.GetString("host.keypath")
	if hostName == "" || keyPath == "" {
		return fmt.Errorf("HostName or ssh priavate key path is empty")
	}
	machineCfgs := []*MachineCfg{}
	err := viper.UnmarshalKey("machines", &machineCfgs)
	if err != nil {
		return err
	}
	interval := viper.GetInt("interval")
	count := 0
	mtotal := len(machineCfgs)
	bindAddress := viper.GetStringSlice("bindAddress")
	if len(bindAddress) == 0 {
		bindAddress = []string{""}
	}
	workerNum := len(bindAddress)
	machine := make(chan *MachineCfg, workerNum)

	carrierWorker := viper.GetInt("worker")
	if carrierWorker <= 0 {
		log.Println("Woker doesn't set, default is 0")
		carrierWorker = 1
	}
	if carrierWorker > 8 {
		log.Println("Woker larger than 8, default is 8")
		carrierWorker = 8
	}

	//Start worker
	for i := 0; i < workerNum; i++ {
		go worker(i, hostName, keyPath, interval, bindAddress, machine, carrierWorker)
	}

	for {
		if count > mtotal-1 {
			count = 0
		}
		machine <- machineCfgs[count]
		count++
	}
}

func worker(id int, hostname, keypath string, interval int, bindAddress []string, machine <-chan *MachineCfg, carrierWorker int) {
	log.Printf("Start from network interface: %d\n", id)
	for m := range machine {
		currentBindAddress := bindAddress[id]
		log.Printf("Network interface: %d, start job. ip: %s, bindAddress: %s, src: %s, dst: %s\n", id, m.IP, currentBindAddress, m.Src, m.Dst)
		// r := rand.Intn(20)
		// time.Sleep(time.Duration(r) * time.Second)
		err := remote.StartSCPSimple(m.IP, currentBindAddress, m.Src, m.Dst, hostname, keypath, carrierWorker)
		if err != nil {
			log.Printf("Worker_%d, Move file error: %s", id, err)
		}
		log.Printf("Network interface: %d, finish job. ip: %s, bindAddress: %s, src: %s, dst: %s, sleep %d second\n", id, m.IP, currentBindAddress, m.Src, m.Dst, interval)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
