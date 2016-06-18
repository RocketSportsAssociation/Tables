package main

import (
    matcher "github.com/RocketSportsAssociation/matcher"
    "github.com/gocarina/gocsv"
    "flag"
    "fmt"
    "os"
    "container/heap"
)

var rankFileName = flag.String("ranks", "", "A CSV file with team ranks.")
var week = flag.Int("week", 0, "The week number")

func main() {
    flag.Parse()

    if *rankFileName == "" {
        flag.PrintDefaults()
        return
    }

    rankFile, err := os.Open(*rankFileName)
    if err != nil {
        panic(err)
    }
    defer rankFile.Close()

    unsortedRanks := []*matcher.TeamRank{}
    if err := gocsv.UnmarshalFile(rankFile, &unsortedRanks); err != nil {
        panic(err)
    }

    ranks := &matcher.RankedList{}
    heap.Init(ranks)

    for _, r := range unsortedRanks {
        heap.Push(ranks, *r)
    }

    tableFmt := `<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /><link href="http://dynasties.operationsports.com/css/osdyn.css" rel="stylesheet" type="text/css"></head><body><table cellpadding="3" cellspacing="1" class="osdyn" width="900"><tr class="masthead_alt"><td bgcolor="#025C9A" colspan="7">

Week %d Standings</td></tr>

<tr class="stathead"><td width="31%%">TEAM</td>
<td align="right" width="10%%">GP</td>
<td align="right" width="7%%">W</td>
<td align="right" width="7%%">L</td>
<td align="right" width="7%%">D</td>
<td align="right" width="7%%">PTS</td>
<td align="right" width="10%%">WIN %%</td>
<td align="right" width="7%%">GD</td>
<td align="right" width="7%%">GF</td>
<td align="right" width="7%%">GA</td></tr>\n\n
%s
</table></body></html>
`
/*
<tr class="evenrow"><td>[TEAMNAME]</td><td align="right">[GAMESPLAYED]</td><td align="right">[WINS]</td><td align="right">[LOSSES]</td><td align="right">[DRAWS]</td><td align="right">[POINTS]</td><td align="right">[WIN%]</td><td align="right">[GD]</td><td align="right">[GF]</td><td align="right">[GA]</td></tr>
*/
    teamStr := ""
    teamFmt := `<tr class="%s"><td>%s</td><td align="right">%d</td><td align="right">%d</td><td align="right">%d</td><td align="right">%d</td><td align="right">%d</td><td align="right">%s</td><td align="right">%d</td><td align="right">%d</td><td align="right">%d</td></tr>

`
    teamNum := 1
    for ranks.Len() > 0 {
        row := "oddrow"
        if teamNum % 2 == 0 {
            row = "evenrow"
        }
        teamNum++
        team := heap.Pop(ranks).(matcher.TeamRank)
        teamStr += fmt.Sprintf(teamFmt, row, team.Name, team.Played, team.Wins, team.Losses, team.Draws, team.Points, team.WinPercentage, team.GoalDiff, team.GoalsFor, team.GoalsAgainst)
    }
    fmt.Printf(tableFmt, *week, teamStr)
}
