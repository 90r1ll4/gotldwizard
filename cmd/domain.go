/*
Copyright © 2023 Ashwin Singh <ashwinsinghsingh666672@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/90r1ll4/gotldwizard/domainChanger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var domainList string
var domainName string
var tld string
var tldList string

var output string
var httpxS bool

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:     "domain [flags]",
	Short:   "Domain Mode",
	Aliases: []string{"dom"},
	Run: func(cmd *cobra.Command, args []string) {
		// if httpxS {
		// 	fmt.Println("Verbose mode active")
		// } else {
		// 	fmt.Println("Verbose mode not active")
		// }

		if domainName != "" && tld != "" {
			domainChanger.DomainCheckerCase1(domainName, tld, output, httpxS)
		} else if domainList != "" && tld != "" {
			domainChanger.DomainCheckerCase2(domainList, tld, output, httpxS)
		} else if domainList != "" && tldList != "" {
			domainChanger.DomainCheckerCase3(domainList, tldList, output, httpxS)
		} else if domainName != "" && tldList != "" {
			domainChanger.DomainCheckerCase4(domainName, tldList, output, httpxS)
		}

	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
	domainCmd.PersistentFlags().StringVarP(&domainName, "domainName", "d", "", "domain name")
	domainCmd.PersistentFlags().StringVarP(&domainList, "domainList", "l", "", "domain name list")
	domainCmd.PersistentFlags().StringVarP(&tld, "tld", "t", "", "TLD(Top Level Domain) to add[com,in,tech]")
	domainCmd.PersistentFlags().StringVarP(&tldList, "tldList", "T", "", "TLD (Top Level Domain) list to add")
	domainCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output file name or path")
	rootCmd.PersistentFlags().BoolVar(&httpxS, "httpx", false, "Use httpx for the output domains")

	domainCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {

		if domainName == "" && domainList == "" {
			color.Red("Either Domain or Domain List Argument must be provided use Flag[-d][-l]")
			os.Exit(1)
		}
		if domainName != "" && domainList != "" {
			color.Red("Do not use both flag same time Flag[[-d] or[-l]]")
			os.Exit(1)
		}

		if tld == "" && tldList == "" {
			color.Red("Either TLD or TLD List Argument must be provided use Flag[-t][-T]")
			os.Exit(1)
		}
		return nil
	}

}
