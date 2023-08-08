package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
	"encoding/csv"
	"strings"
	"strconv"
)

func main() {
	data_filename := flag.String("data", "data.csv", "mail address list separated by CR/LF")
	template_filename := flag.String("template", "template.txt", "message template with place holder $N")
	out_dir := flag.String("out", "out/", "directory path that output files located")
	help := flag.Bool("help", false, "display help message")
	flag.Parse()

	//
	if *help {
		displayGuideMessage()
		return
	}

	//
	data_fp, err := os.Open(*data_filename)
	if err != nil {
		fmt.Println("ERROR: data file cannot be opened")
		displayGuideMessage()
		return
	}
	defer data_fp.Close()

	reader := csv.NewReader(data_fp)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("ERROR: csv format is invalid")
		return
	}

	//
	template_fp, err := os.Open(*template_filename)
	if err != nil {
		fmt.Println("ERROR: template file cannot be opened")
		displayGuideMessage()
		return
	}
	defer template_fp.Close()

	template_scanner := bufio.NewScanner(template_fp)
	serialized_template := ""
	for template_scanner.Scan() {
		serialized_template += template_scanner.Text() + "\n"
	}

	//
	if (*out_dir)[len(*out_dir)-1:] != "/" {
		*out_dir = *out_dir + "/"
	}

	if _, err := os.Stat(*out_dir); os.IsNotExist(err) {
		if err := os.Mkdir(*out_dir, 0777); err != nil {
			fmt.Println(err)
			return
		}
	}

	//
	for _, row := range rows {
		out_fp, err := os.Create(*out_dir + row[0] + ".txt")
		if err != nil {
			panic(err)
		}
		defer out_fp.Close()

		template_double := serialized_template
		for i, e := range row {
			if i == 0 {
				continue
			}
			placeholder := "$" + strconv.Itoa(i)
			template_double = strings.Replace(template_double, placeholder, e, 1)
		}
		out_fp.WriteString(template_double)
	}

	fmt.Println("generated", len(rows), "files")
}

func displayGuideMessage() {
	fmt.Println(" -data     : default: [data.csv]     address and more (column 0 must be mail address)")
	fmt.Println(" -template : default: [template.txt] template message")
	fmt.Println(" -out      : default: [out/]         output dir")
}
