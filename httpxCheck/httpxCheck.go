package httpxCheck

import (
	"log"
	"net"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"
)

func HttpxChecker(httpxSliceUrl []string) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)

	var urls []string

	for _, url := range httpxSliceUrl {
		// Resolve the domain name to an IP address
		ips, err := net.LookupIP(url)
		if err != nil {
			// fmt.Printf("Skipping '%s': %s\n", url, err)
			continue
		}

		// Add the URL to the slice if there are no errors
		if len(ips) > 0 {
			urls = append(urls, url)
		}
	}

	options := runner.Options{
		Methods:         "GET",
		InputTargetHost: urls,
		//InputFile: "./targetDomains.txt", // path to file containing the target domains list
	}
	if err := options.ValidateOptions(); err != nil {
		log.Fatal(err)
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {

		log.Fatal(err)

	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()
}
