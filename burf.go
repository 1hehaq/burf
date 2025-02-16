// haq (https://github.com/1hehaq)
//inspired from - https://x.com/h4x0r_dz/status/1573318682230530048

//This script generates a list of common paths for a given hostname. It can be used for fuzzing with ffuf or any tools you prefer.
//the script can also generate intermediate variants for hostnames with dots (when -both flag is set), like; with TLD or SUBDOMAINS or without both and so on...
//you can provide extra extensions using -ext flag or read from STDIN using 'pipe' value
//if no host is provided, the script will read from STDIN
//if no extensions are provided, the script will use extensions from defExts

// usage: burf -host github.com -ext pipe < common.txt | ffuf -w - -u https://github.com/FUZZ

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	hostFlag := flag.String("host", "", "hostname (e.g: github or sub.github.com) or pipe to read from STDIN\n\t(cat domains.txt | burf -ext .db,.sql,.bak,.old)")
	extFlag := flag.String("ext", "", "extra extensions (comma separated) or 'pipe' to read from STDIN \n\t(cat extensions.txt | burf -host github.com -ext pipe)")
	bothFlag := flag.Bool("both", false, "generate intermediate variants for hostnames with dots")
	helpFlag := flag.Bool("help", false, "show help")

	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.Usage = func() {}
	flag.CommandLine.SetOutput(io.Discard) // discard SetOutput

	err := flag.CommandLine.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			flag.Usage()
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	if *helpFlag {
		flag.CommandLine.SetOutput(os.Stderr)
		flag.Usage()
		os.Exit(0)
	}

	var hosts []string
	var pipedExts []string

	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	if fi.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		if *extFlag == "pipe" {
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" {
					pipedExts = append(pipedExts, line)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				os.Exit(1)
			}
		} else {
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" {
					hosts = append(hosts, line)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				os.Exit(1)
			}
		}
	}

	if len(hosts) == 0 {
		if *hostFlag == "" {
			fmt.Fprintln(os.Stderr, "no host provided. use -host flag or pipe input")
			os.Exit(1)
		}
		hosts = append(hosts, *hostFlag)
	}

	defExts := []string{
		".rar",
		".sql.tar",
		".tar.gz",
		".tar.bzip2",
		".tar",
		".sql.bz2",
		".7z",
		".sql.7z",
		".env",
		".config",
		".bak",
		".old",
		".backup",
		".zip",
		".zip.0",
		".zip1",
		".zipa",
		".zipA0",
		".zip.a0",
		".tgz",
		".sql",
		".db",
		".sqlite",
		".pgsql.txt",
		".mysql.txt",
		".gz",
		".log",
		".bkp",
		".crt",
		".dat",
		".eml",
		".java",
		".lst",
		".key",
		".passwd",
		".pl",
		".pwd",
		".mysql-connect",
		".jar",
		".cfg",
		".dir",
		".orig",
		".bz2",
		".vbs",
		".img",
		".inf",
		".sh",
		".py",
		".vbproj",
		".mysql-pconnect",
		".war",
		".go",
		".psql",
		".sql.gz",
		".vb",
		".webinfo",
		".jnlp",
		".cgi",
		".temp",
		".ini",
		".webproj",
		".xsql",
		".raw",
		".inc",
		".lck",
		".nz",
		".rc",
		".html.gz",
		".yml",
	}

	if *extFlag != "" && *extFlag != "pipe" {
		extra := strings.Split(*extFlag, ",")
		for i := range extra {
			extra[i] = strings.TrimSpace(extra[i])
		}
		defExts = append(defExts, extra...)
	}

	if *extFlag == "pipe" && len(pipedExts) > 0 {
		defExts = pipedExts
	}

	for _, host := range hosts {
		if *bothFlag {
			parts := strings.Split(host, ".")
			var variant string
			for i, p := range parts {
				if i == 0 {
					variant = p
				} else {
					variant = variant + "." + p
				}
				for _, ext := range defExts {
					fmt.Println("/" + variant + ext)
				}
			}
		} else {
			for _, ext := range defExts {
				fmt.Println("/" + host + ext)
			}
		}
	}
}
