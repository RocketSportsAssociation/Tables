# Tables
Read in a csv, output HTML-formatted tables

## Installing

Just set `GOPATH` and run `go get github.com/RocketSportsAssociation/Tables`. This should pull and install a binary: `$GOPATH/bin/Tables`

## Running
```
$ ./Tables 
  -file string
    	A CSV file.
  -title string
    	The title of the table being generated. (default "Table")
    	```

`./bin/Tables -file my.csv -title FooBar`

```
<html lang="en">
<head>
    <meta charset="utf-8" />
    <title>Table</title>
    <meta name="viewport" content="initial-scale=1.0; maximum-scale=1.0; width=device-width;">
</head>

<body>
<table class="tablesorter">
<thead>
<th class="text-left">Team</th>
<th class="text-left">GP</th>
<th class="text-left">W</th>
<th class="text-left">L</th>
<th class="text-left">D</th>
<th class="text-left">Pts</th>
<th class="text-left">Win Pct</th>
<th class="text-left">GD</th>
<th class="text-left">GF</th>
<th class="text-left">GA</th>

</thead>
<tbody class="table-hover">
<tr>
<td class="text-left">Boosted Monkeys</td>
<td class="text-left">5</td>
<td class="text-left">4</td>
<td class="text-left">0</td>
<td class="text-left">1</td>
<td class="text-left">13</td>
<td class="text-left">0.800</td>
<td class="text-left">17</td>
<td class="text-left">33</td>
<td class="text-left">16</td>

</tr>
<tr>
<td class="text-left">Virtual Rockets</td>
<td class="text-left">4</td>
<td class="text-left">4</td>
<td class="text-left">0</td>
<td class="text-left">0</td>
<td class="text-left">12</td>
<td class="text-left">1.000</td>
<td class="text-left">43</td>
<td class="text-left">49</td>
<td class="text-left">6</td>

</tr>
<tr>
<td class="text-left">Squared Lions</td>
<td class="text-left">4</td>
<td class="text-left">3</td>
<td class="text-left">0</td>
<td class="text-left">1</td>
<td class="text-left">10</td>
<td class="text-left">0.750</td>
<td class="text-left">13</td>
<td class="text-left">23</td>
<td class="text-left">10</td>

</tr>
<tr>
<td class="text-left">Rykon Gaming </td>
<td class="text-left">4</td>
<td class="text-left">3</td>
<td class="text-left">0</td>
<td class="text-left">1</td>
<td class="text-left">10</td>
<td class="text-left">0.750</td>
<td class="text-left">10</td>
<td class="text-left">16</td>
<td class="text-left">6</td>

</tr>

</tbody>
</table>
  

  </body>
```
