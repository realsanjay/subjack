package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./subjack"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	Project := "./"
	configFile := "fingerprints.json"
	defaultConfig := GOPATH + Project + configFile

	o := subjack.Options{}

	flag.StringVar(&o.Domain, "d", "", "Domain.")
	flag.StringVar(&o.Wordlist, "w", "", "Path to wordlist.")
	flag.StringVar(&o.DirPath, "f", "", "Path to folder contains list of subdomain files")
	flag.IntVar(&o.Threads, "t", 10, "Number of concurrent threads (Default: 10).")
	flag.IntVar(&o.Timeout, "timeout", 10, "Seconds to wait before connection timeout (Default: 10).")
	flag.BoolVar(&o.Ssl, "ssl", false, "Force HTTPS connections (May increase accuracy (Default: http://).")
	flag.BoolVar(&o.All, "a", false, "Find those hidden gems by sending requests to every URL. (Default: Requests are only sent to URLs with identified CNAMEs).")
	flag.BoolVar(&o.Verbose, "v", false, "Display more information per each request.")
	flag.StringVar(&o.Output, "o", "", "Output results to file (Subjack will write JSON if file ends with '.json').")
	flag.StringVar(&o.Config, "c", defaultConfig, "Path to configuration file.")
	flag.BoolVar(&o.Manual, "m", false, "Flag the presence of a dead record, but valid CNAME entry.")

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	o.Output = o.DirPath + "/subjack"
	os.MkdirAll(o.Output, os.ModePerm);
	o.Output += "/output.txt"

	var files []string

    filepath.Walk(o.DirPath, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        if info.IsDir() {
    		return nil
		}
        o.Wordlist = path
        subjack.Process(&o)
        return nil
    })
}
