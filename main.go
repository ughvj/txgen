package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
)

func main() {
	addresses_filename := flag.String("addresses", "addresses.txt", "mail address list separated by CR/LF")
	contents_filename := flag.String("contents", "contents.txt", "contents of mail body")
	out_dir := flag.String("out", "out/", "directory path that output files located")
	help := flag.Bool("help", false, "display help message")
	flag.Parse()

	//
	if *help {
		fmt.Println("-addresses: address list file path")
		fmt.Println("-contents : contents file path")
		fmt.Println("-out      : output dir (must attach separated character of dir at end of string)")
		return
	}

	//
	addresses_fp, err := os.Open(*addresses_filename)
	if err != nil {
		fmt.Println("addresses file cannot be opened")
		return
	}
	defer addresses_fp.Close()

	addresses_scanner := bufio.NewScanner(addresses_fp)
	var addresses []string
	for addresses_scanner.Scan() {
		addresses = append (addresses, addresses_scanner.Text())
	}

	//
	contents_fp, err := os.Open(*contents_filename)
	if err != nil {
		fmt.Println("contents file cannot be opened")
		return
	}
	defer contents_fp.Close()

	contents_scanner := bufio.NewScanner(contents_fp)
	var contents []string
	for contents_scanner.Scan() {
		contents = append (contents, contents_scanner.Text())
	}

	//
	if _, err := os.Stat(*out_dir); os.IsNotExist(err) {
		if err := os.Mkdir(*out_dir, 0644); err != nil {
        	fmt.Println(err)
			return
    	}
	}

	//
	for _, address := range addresses {
		out_fp, err := os.Create(*out_dir + address + ".txt")
		if err != nil {
			panic(err)
		}
		defer out_fp.Close()
		for _, content := range contents {
			out_fp.WriteString(content + "\n")
		}
	}

	fmt.Println("generated", len(addresses), "files")
}
