package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func toFils(s string) (int, error) {
	s = strings.TrimSpace(s)
	var err error
	fils := 0
	split := strings.Split(s, ".")

	if len(split) == 2 {
		fils, err = strconv.Atoi(split[1])
		if err != nil {
			return -1, err
		}
		if split[1] != "" {
			a, err := strconv.Atoi(string(split[0]))
			if err != nil {
				return -1, err
			}

			m := 100
			for a > 0 {
				n := a % 10
				fils += (n * m)
				a = a / 10
				m = m * 10
			}
		}
	} else {
		a, err := strconv.Atoi(string(split[0]))
		if err != nil {
			return -1, err
		}

		m := 100
		for a > 0 {
			n := a % 10
			fils += (n * m)
			a = a / 10
			m = m * 10
		}
	}

	return fils, nil
}

func diff(previous []string, current []string) map[string]bool {
	return nil
}

type Spend struct {
	Place  string
	Amount int
}

func main() {

	var path string
	flag.StringVar(&path, "path", "", "--path")
	flag.Parse()

	if path == "" {
		log.Panic("--path cannot be empty")
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	contents := string(bytes)

	tbl := tablewriter.NewWriter(os.Stdout)
	tbl.SetHeader([]string{"Spend", "Amount"})

	output := make([][]string, 0)

	spends := make(map[string]int)
	var str strings.Builder
	for _, v := range contents {
		if v == '\n' {
			sp := strings.TrimSpace(str.String())
			split := strings.Split(sp, " ")

			spend := strings.Join(split[0:len(split)-1], " ")
			amount, err := toFils(split[len(split)-1])
			if err != nil {
				log.Panicf("spend -> %s - %s", sp, err)
			}

			v, _ := spends[spend]
			spends[spend] = v + amount

			str.Reset()
			continue
		}
		str.WriteRune(v)
	}

	sp := make([]Spend, 0)
	for k, v := range spends {
		sp = append(sp, Spend{Place: k, Amount: v})
	}

	sort.Slice(sp, func(i, j int) bool {
		return sp[i].Amount > sp[j].Amount
	})

	total := 0
	for _, v := range sp {
		output = append(output, []string{v.Place, fmt.Sprintf("%d.%d", v.Amount/100, v.Amount%100)})
		total += v.Amount
	}
	tbl.AppendBulk(output)

	tbl.SetFooter([]string{"Total", fmt.Sprintf("%d.%d", total/100, total%100)})
	tbl.SetFooterAlignment(tablewriter.ALIGN_LEFT)
	tbl.Render()
}
