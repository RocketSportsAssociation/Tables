# Tables
Read in a csv or google spreadsheet of team ranks, output HTML-formatted tables

## Installing

Just set `GOPATH` and run `go get github.com/RocketSportsAssociation/Tables`. This should pull and install a binary: `$GOPATH/bin/Tables`

## Running
  -ranks string
        A CSV file with team ranks.
  -week int
        The week number

`./bin/Tables -ranks “path/to/sheets/ORSAQualifiersWeek1.xlsx - PC-PS4 1v1.csv” -week 2`

```
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /><link href="http://dynasties.operationsports.com/css/osdyn.css" rel="stylesheet" type="text/css"></head><body><table cellpadding="3" cellspacing="1" class="osdyn" width="900"><tr class="masthead_alt"><td bgcolor="#025C9A" colspan="7">

Week 2 Standings</td></tr>

<tr class="stathead"><td width="10%">TEAM</td>
<td align="right" width="10%">GP</td>
<td align="right" width="10%">W</td>
<td align="right" width="10%">L</td>
<td align="right" width="10%">D</td>
<td align="right" width="10%">PTS</td>
<td align="right" width="10%">WIN %</td>
<td align="right" width="10%">GD</td>
<td align="right" width="10%">GF</td>
<td align="right" width="10%">GA</td></tr>
<tr class="oddrow"><td>Recycer</td><td align="right">5</td><td align="right">4</td><td align="right">0</td><td align="right">1</td><td align="right">13</td><td align="right">0.800</td><td align="right">-6</td><td align="right">18</td><td align="right">24</td></tr>

<tr class="evenrow"><td>IceyCocaine</td><td align="right">5</td><td align="right">4</td><td align="right">1</td><td align="right">0</td><td align="right">12</td><td align="right">0.800</td><td align="right">46</td><td align="right">65</td><td align="right">19</td></tr>

<tr class="oddrow"><td>[sNb] Gemini</td><td align="right">4</td><td align="right">4</td><td align="right">0</td><td align="right">0</td><td align="right">12</td><td align="right">1.000</td><td align="right">22</td><td align="right">48</td><td align="right">26</td></tr>

<tr class="evenrow"><td>adamschaub</td><td align="right">5</td><td align="right">4</td><td align="right">1</td><td align="right">0</td><td align="right">12</td><td align="right">0.800</td><td align="right">19</td><td align="right">49</td><td align="right">30</td></tr>

<tr class="oddrow"><td>nicckc9</td><td align="right">5</td><td align="right">3</td><td align="right">2</td><td align="right">0</td><td align="right">9</td><td align="right">0.600</td><td align="right">-10</td><td align="right">18</td><td align="right">28</td></tr>

<tr class="evenrow"><td>DoctorMcBoop</td><td align="right">4</td><td align="right">3</td><td align="right">1</td><td align="right">0</td><td align="right">9</td><td align="right">0.750</td><td align="right">-11</td><td align="right">3</td><td align="right">14</td></tr>

<tr class="oddrow"><td>ICEWS IDontThink</td><td align="right">4</td><td align="right">2</td><td align="right">1</td><td align="right">1</td><td align="right">7</td><td align="right">0.500</td><td align="right">-3</td><td align="right">8</td><td align="right">11</td></tr>

<tr class="evenrow"><td>McMonkey</td><td align="right">4</td><td align="right">2</td><td align="right">1</td><td align="right">1</td><td align="right">7</td><td align="right">0.500</td><td align="right">-23</td><td align="right">2</td><td align="right">25</td></tr>

<tr class="oddrow"><td>Loldongs95</td><td align="right">5</td><td align="right">2</td><td align="right">3</td><td align="right">0</td><td align="right">6</td><td align="right">0.400</td><td align="right">0</td><td align="right">41</td><td align="right">41</td></tr>

<tr class="evenrow"><td>RezHades</td><td align="right">1</td><td align="right">1</td><td align="right">0</td><td align="right">0</td><td align="right">3</td><td align="right">1.000</td><td align="right">29</td><td align="right">30</td><td align="right">1</td></tr>

<tr class="oddrow"><td>MacGyver</td><td align="right">0</td><td align="right">0</td><td align="right">0</td><td align="right">0</td><td align="right">0</td><td align="right">#DIV/0!</td><td align="right">0</td><td align="right">0</td><td align="right">0</td></tr>

<tr class="evenrow"><td>Ohflawless</td><td align="right">0</td><td align="right">0</td><td align="right">0</td><td align="right">0</td><td align="right">0</td><td align="right">#DIV/0!</td><td align="right">0</td><td align="right">0</td><td align="right">0</td></tr>

<tr class="oddrow"><td>SnipaEagleEye</td><td align="right">1</td><td align="right">0</td><td align="right">1</td><td align="right">0</td><td align="right">0</td><td align="right">0.000</td><td align="right">-13</td><td align="right">4</td><td align="right">17</td></tr>


</table></body></html>
```
