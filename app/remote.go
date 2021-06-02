package app

import (
	"log"
	"plotcarrier/remote"
	"time"

	"github.com/spf13/viper"
)

type MachineCfg struct {
	IP  string `yaml:"ip"`
	Src string `yaml:"remote_src"`
	Dst string `yaml:"local_dst"`
}

func RemoteStart() error {
	machineCfgs := []*MachineCfg{}
	err := viper.UnmarshalKey("machines", &machineCfgs)
	if err != nil {
		return err
	}
	interval := viper.GetInt("interval")
	count := 0
	mtotal := len(machineCfgs)
	workerNum := viper.GetInt("worker")
	machine := make(chan *MachineCfg, workerNum)

	//Start worker
	for i := 0; i < workerNum; i++ {
		go worker(i, machine)
	}

	for {
		if count > mtotal-1 {
			log.Println("Already finish a round")
			time.Sleep(time.Duration(interval) * time.Second)
			count = 0
		}
		machine <- machineCfgs[count]
		count++
	}
}

func worker(id int, machine <-chan *MachineCfg) {
	log.Printf("Start worker_%d\n", id)
	for m := range machine {
		log.Printf("Worker_%d, start job. ip: %s, src: %s, dst: %s\n", id, m.IP, m.Src, m.Dst)
		//r := rand.Intn(10)
		//time.Sleep(time.Duration(r) * time.Second)
		err := remote.MV(m.IP, m.Src, m.Dst)
		if err != nil {
			log.Printf("Move file error: %s", err)
		}
		log.Printf("Worker_%d, finish job. ip: %s, src: %s, dst: %s\n", id, m.IP, m.Src, m.Dst)
	}
	log.Printf("End worker_%d\n", id)
}
