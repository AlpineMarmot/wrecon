package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"io"
	"os"
	"sync"
	"time"
	"wrecon"
)

var (
	errors        = []string{}
	pingResponses = []*wrecon.PingResponse{}
	cliTable      = newCliTable()
)

func pingSite(wg *sync.WaitGroup, lock *sync.Mutex, site wrecon.Site) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	resp, err := wrecon.Request(site.Address)

	if err != nil {
		pingResponses = append(pingResponses, wrecon.FailedPingResponse(site, err))
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	pingResponses = append(pingResponses, wrecon.SucceedPingResponse(site, resp))
	return
}

func newCliTable() table.Table {
	tbl := table.New("Site", "Status", "Code")
	headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgHiWhite).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func main() {
	configFile := flag.String("config", "", "config json file to use")
	interval := flag.Duration("interval", 0, "check every X seconds")
	flag.Parse()

	if configFile == nil || *configFile == "" {
		fmt.Println("a config file is needed")
		os.Exit(1)
	}

	configBytes, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config wrecon.Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		fmt.Printf("Error while decoding '%s' file:\n", *configFile)
		fmt.Println(err)
		os.Exit(1)
	}

loop:
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(len(config.Sites))

	for _, site := range config.Sites {
		go pingSite(&wg, &lock, site)
	}

	wg.Wait()

	wrecon.ClearTerminal()
	fmt.Println("Wrecon v0.1")

	for _, resp := range pingResponses {
		resultName := "Failed"
		var status int
		if resp.Succeed {
			resultName = "Ok"
			status = resp.Response.StatusCode
		}
		cliTable.AddRow(resp.Site.GetName(), resultName, status)
		if !resp.Succeed {
			errors = append(errors, resp.ErrMsg)
		}
	}
	cliTable.Print()

	if len(errors) > 0 {
		fmt.Printf("\nGot %d error(s):\n", len(errors))
		for _, err := range errors {
			fmt.Println(err)
		}
	}

	if *interval > 0 {
		cliTable = newCliTable()
		sleepForNSec := interval.Seconds()
		pingResponses = []*wrecon.PingResponse{}
		errors = []string{}
		fmt.Printf("\nRetry in %.0f sec(s) ...\n\n", sleepForNSec)
		time.Sleep(*interval)
		wrecon.ClearTerminal()
		fmt.Printf("\nPinging sites ... \n\n")
		goto loop
	}

	fmt.Println("End")
}
