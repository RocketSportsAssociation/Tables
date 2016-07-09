package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	tableFmt = `<html lang="en">
<head>
    <meta charset="utf-8" />
    <title>%s</title>
    <meta name="viewport" content="initial-scale=1.0; maximum-scale=1.0; width=device-width;">
</head>

<body>
<table class="tablesorter">
<thead>
%s
</thead>
<tbody class="table-hover">
%s
</tbody>
</table>
  

  </body>
`

	tableHeader = `<th class="text-left">%s</th>
`

	tableRow = `<tr>
%s
</tr>
`

	tableRowElem = `<td class="text-left">%s</td>
`
)

var title = flag.String("title", "Table", "The title of the table being generated.")
var csvFile = flag.String("file", "", "A CSV file.")

func main() {
	flag.Parse()

	if *csvFile == "" {
		flag.PrintDefaults()
		return
	}

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Println("The file is messed up in some way, dummy")
		panic(err)
	}
	defer file.Close()

	lineScan := bufio.NewScanner(file)

	_ = lineScan.Scan()
	headers := lineScan.Text()
	headerList := strings.Split(headers, ",")
	headerStr := ""
	for _, h := range headerList {
		headerStr += fmt.Sprintf(tableHeader, h)
	}
	//fmt.Println(headerStr)

	tableRowStr := ""
	for lineScan.Scan() {
		s := lineScan.Text()
		ss := strings.Split(s, ",")
		rowElemStr := ""
		for _, elem := range ss {
			rowElemStr += fmt.Sprintf(tableRowElem, elem)
		}
		tableRowStr += fmt.Sprintf(tableRow, rowElemStr)
	}

	fmt.Printf(tableFmt, *title, headerStr, tableRowStr)
}
