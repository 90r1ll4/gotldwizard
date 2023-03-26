# gotldwizard

Gotldwizard is a Go package that helps to change the top-level domain (TLD) of a given domain name. This package provides a command-line interface to change TLDs for one or multiple domain names in a single command.


# Using Gotldwizard

## Installation
To install Gotldwizard, use the following command:

```bash
go get -u github.com/username/gotldwizard
```

## Usage

Gotldwizard provides a command-line interface to change TLDs for one or multiple domain names in a single command. The following are the available commands and their respective options:

```bash

__/\\\\\\\\\\\\\\\__/\\\______________/\\\\\\\\\\\\_______________/\\\______________/\\\__/\\\\\\\\\\\__/\\\\\\\\\\\\\\\_____/\\\\\\\\\_______/\\\\\\\\\______/\\\\\\\\\\\\____        
 _\///////\\\/////__\/\\\_____________\/\\\////////\\\____________\/\\\_____________\/\\\_\/////\\\///__\////////////\\\____/\\\\\\\\\\\\\___/\\\///////\\\___\/\\\////////\\\__       
  _______\/\\\_______\/\\\_____________\/\\\______\//\\\___________\/\\\_____________\/\\\_____\/\\\_______________/\\\/____/\\\/////////\\\_\/\\\_____\/\\\___\/\\\______\//\\\_      
   _______\/\\\_______\/\\\_____________\/\\\_______\/\\\___________\//\\\____/\\\____/\\\______\/\\\_____________/\\\/_____\/\\\_______\/\\\_\/\\\\\\\\\\\/____\/\\\_______\/\\\_     
    _______\/\\\_______\/\\\_____________\/\\\_______\/\\\____________\//\\\__/\\\\\__/\\\_______\/\\\___________/\\\/_______\/\\\\\\\\\\\\\\\_\/\\\//////\\\____\/\\\_______\/\\\_    
     _______\/\\\_______\/\\\_____________\/\\\_______\/\\\_____________\//\\\/\\\/\\\/\\\________\/\\\_________/\\\/_________\/\\\/////////\\\_\/\\\____\//\\\___\/\\\_______\/\\\_   
      _______\/\\\_______\/\\\_____________\/\\\_______/\\\_______________\//\\\\\\//\\\\\_________\/\\\_______/\\\/___________\/\\\_______\/\\\_\/\\\_____\//\\\__\/\\\_______/\\\__  
       _______\/\\\_______\/\\\\\\\\\\\\\\\_\/\\\\\\\\\\\\/_________________\//\\\__\//\\\_______/\\\\\\\\\\\__/\\\\\\\\\\\\\\\_\/\\\_______\/\\\_\/\\\______\//\\\_\/\\\\\\\\\\\\/___ 
        _______\///________\///////////////__\////////////____________________\///____\///_______\///////////__\///////////////__\///________\///__\///________\///__\////////////_____

                                                                                                                                             By:- Ashwin (@90r1ll4)
Swapping out the TLD, you can modify the domain name's extension to better suit your needs.

Usage:  tldwizard [command]

Available Commands:
  domain      To choose domain mode
  help        Help about any command

Flags:
  -h, --help      help for tldwizard
  -v, --version   version for tldwizard

Use "tldwizard [command] --help" for more information about a command.
```
To see subcommands add `command -h`

```
Usage:
  tldwizard domain [flags]

Aliases:
  domain, dom

Flags:
  -l, --domainList string   domain name list
  -d, --domainName string   domain name
  -h, --help                help for domain
  -o, --output string       output file name or path
  -t, --tld string          TLD(Top Level Domain) to add[com,in,tech]
  -T, --tldList string      TLD (Top Level Domain) list to add
```

For an overview of all commands use the following command:

```bash
gotldwizard domain [-d DOMAIN_NAME | -l DOMAIN_LIST] [-t TLD | -T TLD_LIST] [-o OUTPUT_FILE]

```

- `domain`: This command is used to choose the domain mode.
- `-d DOMAIN_NAME`: This flag is used to specify a single domain name to change its TLD.
- `-l DOMAIN_LIST`: This flag is used to specify a list of domain names to change their TLDs.
- `-t TLD`: This flag is used to specify a TLD to add to the domain name(s).
- `-T TLD_LIST`: This flag is used to specify a list of TLDs to add to the domain name(s).
- `-o OUTPUT_FILE`: This flag is used to specify an output file name or path.

## Examples
### To change the TLD of a single domain name, use the following command:



```bash
gotldwizard domain -d example.com -t in,nl

```
This will change the TLD of `example.com` to `in` and `n;`.
```
example.com
example.in
example.nl
```
### To change the TLD of multiple domain names, use the following command:

```bash
gotldwizard domain -l domains.txt -t tech
```
This will change the TLD of all the domain names listed in `domains.txt` to `tech`.
```bash
example.comexample.nl
example.tech
example.comexample.nl
example.tech

```

To store the output in file use `-o  FILENAME` 

**NOTE:** Gotldwizard should be used responsibly and with permission from the target owner.