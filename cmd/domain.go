/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gotldwizard/domainChanger"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var domainList string
var domainName string
var tld string
var tldList string

var output string

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:     "domain",
	Short:   "To choose domain mode",
	Aliases: []string{"dom"},
	Run: func(cmd *cobra.Command, args []string) {
		if domainName != "" && tld != "" {
			domainChanger.DomainCheckerCase1(domainName, tld, output)
		} else if domainList != "" && tld != "" {
			domainChanger.DomainCheckerCase2(domainList, tld, output)
		} else if domainList != "" && tldList != "" {
			domainChanger.DomainCheckerCase3(domainList, tldList, output)
		} else if domainName != "" && tldList != "" {
			domainChanger.DomainCheckerCase4(domainName, tldList, output)
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
