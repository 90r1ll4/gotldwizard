package domainChanger

import (
	"os"
	"sort"
	"strings"

	"github.com/90r1ll4/gotldwizard/httpxCheck"
	"github.com/fatih/color"
	tld "github.com/jpillora/go-tld"
)

func DomainCheckerCase1(domainName string, tldName string, output string, httpxS bool) {
	var parsedUrl *tld.URL
	if strings.HasPrefix(domainName, "http") {
		parsedUrl, _ = tld.Parse(domainName)
	} else {
		parsedUrl, _ = tld.Parse("http://" + domainName)
	}
	tldSlice := strings.Split(tldName, ",")
	tldSlice = append(tldSlice, parsedUrl.TLD)
	uniqueTld := FindUniqueTld(tldSlice)

	var content string

	for _, value := range uniqueTld {
		content += parsedUrl.Domain + "." + value + "\n"
	}
	if httpxS {
		httpSlice := strings.Split(content, "\n")
		httpxCheck.HttpxChecker(httpSlice)
		os.Exit(0)
	}

	if output != "" {
		outputInFile(output, content)
	} else {
		color.Green(content)
	}
}

func DomainCheckerCase2(domainList string, tldName string, output string, httpxS bool) {
	fileDomain := string(readFile(domainList))
	fileDomainSlice := strings.Split(string(fileDomain), "\n")
	var tldSlice []string = strings.Split(tldName, ",")
	var content []string

	for _, value := range fileDomainSlice {
		if value != "" {
			var parsedUrl *tld.URL
			value = strings.TrimRight(value, "\r")
			if strings.HasPrefix(value, "http") {
				parsedUrl, _ = tld.Parse(value)
			} else {
				parsedUrl, _ = tld.Parse("http://" + value)
			}
			if parsedUrl.TLD != "" {
				tldSlice = append(tldSlice, parsedUrl.TLD)
			}
		}
	}
	uniqueTld := FindUniqueTld(tldSlice)
	for _, value2 := range fileDomainSlice {
		for _, value := range uniqueTld {
			if value2 != "" {
				value = strings.TrimRight(value, "\r")
				value2 = strings.TrimRight(value2, "\r")
				var parsedUrl *tld.URL
				if strings.HasPrefix(value2, "http") {
					parsedUrl, _ = tld.Parse(value2)
				} else {
					parsedUrl, _ = tld.Parse("http://" + value2)
				}
				content = append(content, parsedUrl.Domain+"."+value)
			}
		}
	}
	if httpxS {
		httpxCheck.HttpxChecker(content)
		os.Exit(0)
	}

	if output != "" {
		outputInFile(output, strings.Join(content, "\n"))
	} else {
		color.Green(strings.Join(content, "\n"))
	}

}

func DomainCheckerCase3(domainList string, tldList string, output string, httpxS bool) {
	fileDomain := string(readFile(domainList))
	fileDomainSlice := strings.Split(string(fileDomain), "\n")
	tldListSlice := string(readFile(tldList))
	var tldSlice []string = strings.Split(tldListSlice, "\n")
	var content []string

	for _, value := range fileDomainSlice {
		if value != "" {
			var parsedUrl *tld.URL
			value = strings.TrimRight(value, "\r")
			if strings.HasPrefix(value, "http") {
				parsedUrl, _ = tld.Parse(value)
			} else {
				parsedUrl, _ = tld.Parse("http://" + value)
			}
			if parsedUrl.TLD != "" {
				tldSlice = append(tldSlice, parsedUrl.TLD)
			}
		}
	}
	uniqueTld := FindUniqueTld(tldSlice)
	for _, value2 := range fileDomainSlice {
		for _, value := range uniqueTld {
			if value2 != "" {
				value = strings.TrimRight(value, "\r")
				value2 = strings.TrimRight(value2, "\r")
				var parsedUrl *tld.URL
				if strings.HasPrefix(value2, "http") {
					parsedUrl, _ = tld.Parse(value2)
				} else {
					parsedUrl, _ = tld.Parse("http://" + value2)
				}
				content = append(content, parsedUrl.Domain+"."+value)
			}
		}
	}

	if httpxS {

		httpxCheck.HttpxChecker(content)
		os.Exit(0)
	}
	if output != "" {
		outputInFile(output, strings.Join(content, "\n"))
	} else {
		color.Green(strings.Join(content, "\n"))
	}

}
func DomainCheckerCase4(domainName string, tldList string, output string, httpxS bool) {
	var parsedUrl *tld.URL
	if strings.HasPrefix(domainName, "http") {
		parsedUrl, _ = tld.Parse(domainName)
	} else {
		parsedUrl, _ = tld.Parse("http://" + domainName)
	}
	tldListSlice := string(readFile(tldList))
	var tldSlice []string = strings.Split(tldListSlice, "\n")
	tldSlice = append(tldSlice, parsedUrl.TLD)
	uniqueTld := FindUniqueTld(tldSlice)
	var content string

	for _, value := range uniqueTld {
		value = strings.TrimRight(value, "\r")
		content += parsedUrl.Domain + "." + value + "\n"
	}

	if httpxS {
		httpSlice := strings.Split(content, "\n")
		httpxCheck.HttpxChecker(httpSlice)
		os.Exit(0)
	}

	if output != "" {
		outputInFile(output, content)
	} else {
		color.Green(content)
	}
}

func FindUniqueTld(tldSlice []string) []string {
	uniqueStrings := make(map[string]bool)
	for _, s := range tldSlice {
		uniqueStrings[s] = true
	}

	uniqueTlds := make([]string, len(uniqueStrings))
	i := 0
	for s := range uniqueStrings {
		uniqueTlds[i] = s
		i++
	}
	sort.Strings(uniqueTlds)
	return uniqueTlds

}

func outputInFile(filePath string, content string) {
	// Check if file exists
	if _, err := os.Stat(filePath); err == nil {
		// If file exists, overwrite it
		file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
		if err != nil {
			color.Red("Error opening file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			color.Red("Error writing to file:", err)
			return
		}
	} else {
		// If file does not exist, create it and write to it
		file, err := os.Create(filePath)
		if err != nil {
			color.Red("Error creating file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			color.Red("Error writing to file:", err)
			return
		}
	}
}

func readFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		color.Red("Error reading file:", err)
	}
	return string(content)
}
