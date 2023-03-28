# gotldwizard

Gotldwizard is a Go package that helps to change the top-level domain (TLD) of a given domain name. This package provides a command-line interface to change TLDs for one or multiple domain names in a single command.


# Using Gotldwizard

## Installation
To install Gotldwizard, use the following command:

```bash
go install github.com/90r1ll4/gotldwizard
```

## Usage

Gotldwizard provides a command-line interface to change TLDs for one or multiple domain names in a single command. The following are the available commands and their respective options:

```bash


        #   _______        ______       _  _  _ _____ ______ _______  ______ ______   #
        #      |    |      |     \      |  |  |   |    ____/ |_____| |_____/ |     \  #
        #      |    |_____ |_____/      |__|__| __|__ /_____ |     | |    \_ |_____/  #
        #                                                                             #

                                                                By:- Ashwin (@90r1ll4)
Swapping out the TLD, you can modify the domain name's extension to better suit your needs.

Usage:
  tldwizard [command]

Available Commands:
  domain      Domain Mode
  help        Help about any command

Flags:
  -h, --help      help for tldwizard
      --httpx     Use httpx for the output domains
  -v, --version   version for tldwizard

Use "tldwizard [command] --help" for more information about a command.
```
To see subcommands add `command -h`

```

        #   _______        ______       _  _  _ _____ ______ _______  ______ ______   #
        #      |    |      |     \      |  |  |   |    ____/ |_____| |_____/ |     \  #
        #      |    |_____ |_____/      |__|__| __|__ /_____ |     | |    \_ |_____/  #
        #                                                                             #

                                                                By:- Ashwin (@90r1ll4)
Domain Mode

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

Global Flags:
      --httpx   Use httpx for the output domains
```
To build use command `go build` to get genrate binary file.

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
example.com
example.nl
example.tech
example.com
example.nl
example.tech

```

To store the output in file use `-o  FILENAME` 

### Using `httpx` Flag

```bash
gotldwizard domain -d example.in -t com,org --httpx

        #   _______        ______       _  _  _ _____ ______ _______  ______ ______   #
        #      |    |      |     \      |  |  |   |    ____/ |_____| |_____/ |     \  #
        #      |    |_____ |_____/      |__|__| __|__ /_____ |     | |    \_ |_____/  #
        #                                                                             #

                                                                By:- Ashwin (@90r1ll4)
https://example.org
https://example.com
```

**NOTE:** Gotldwizard should be used responsibly and with permission from the target owner.
