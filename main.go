package main

import (
	"fmt"
	"time"

	"os/exec"

	"github.com/robfig/cron"

)

func main() {
	fmt.Println("suresh")

	cron := cron.New()

	//err := cron.AddJob("@every 00h00m05s", njobs{cmd: "mkdir", agr: "test/bro"})
	//if err != nil{
	//	fmt.Printf(err.Error())
	//}

	now := time.Now()
	fmt.Println(now)
	//maintain := now.Add(time.Second * 10)
	runat := now.Add(time.Second * 5)
	
	cron.Schedule(scheduler{runat: runat}, njobs{cmd: "mkdir", agr: "hello"})

	cron.Start()

	select{}
}

type scheduler struct {
	runat time.Time
}

func (s scheduler) Next(t time.Time) time.Time {
	if t.After(s.runat) {
		return time.Time{}
	}
	return s.runat
}

type njobs struct {
	cmd string
	agr string
}

func (j njobs) Run () {

	x := exec.Command(j.cmd, j.agr)
	op, err := x.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(op)
	}

	fmt.Println("done", op)
}