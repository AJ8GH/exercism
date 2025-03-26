package ledger

import (
	"errors"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type ChanEntry struct {
	i int
	s string
	e error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry = make([]Entry, len(entries))
	copy(entriesCopy, entries)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{
			Date:        "2014-01-01",
			Description: "",
			Change:      0,
		}}); err != nil {
			return "", err
		}
	}

	doStuffWithEntries(entriesCopy)
	out, err := buildOut(locale)
	if err != nil {
		return "", err
	}

	// Parallelism, always a great idea
	channel := make(chan ChanEntry)

	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			switch {
			case len(entry.Date) != 10:
				channel <- ChanEntry{e: errors.New("")}
			case entry.Date[4] != '-':
				channel <- ChanEntry{e: errors.New("")}
			case entry.Date[7] != '-':
				channel <- ChanEntry{e: errors.New("")}
			}

			date := getDate(locale, entry)

			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}

			a, al := buildA(locale, currency, channel, cents, negative)
			desc := entry.Description
			if len(desc) > 25 {
				desc = desc[:22] + "..."
			} else {
				desc = desc + strings.Repeat(" ", 25-len(desc))
			}

			channel <- ChanEntry{
				i: i,
				s: date + strings.Repeat(" ", 10-len(date)) + " | " + desc + " | " + strings.Repeat(" ", 13-al) + a + "\n",
			}
		}(i, et)
	}

	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-channel
		if v.e != nil {
			return "", v.e
		}

		ss[v.i] = v.s
	}

	for i := 0; i < len(entriesCopy); i++ {
		out += ss[i]
	}

	return out, nil
}

func getCentsStr(cents int) string {
	centsStr := strconv.Itoa(cents)
	switch len(centsStr) {
	case 1:
		centsStr = "00" + centsStr
	case 2:
		centsStr = "0" + centsStr
	}
	return centsStr
}

func buildOut(locale string) (string, error) {
	var out string
	switch locale {
	case "nl-NL":
		out = "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	case "en-US":
		out = "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	default:
		return "", errors.New("")
	}
	return out, nil
}

func doStuffWithEntries(entriesCopy []Entry) {
	map1 := map[bool]int{true: 0, false: 1}
	map2 := map[bool]int{true: -1, false: 1}
	entriesCopy2 := entriesCopy
	for len(entriesCopy2) > 1 {
		first, rest := entriesCopy2[0], entriesCopy2[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (map1[e.Date == first.Date]*map2[e.Date < first.Date]*4 +
					map1[e.Description == first.Description]*map2[e.Description < first.Description]*2 +
					map1[e.Change == first.Change]*map2[e.Change < first.Change]*1) < 0 {
					entriesCopy2[0], entriesCopy2[i+1] = entriesCopy2[i+1], entriesCopy2[0]
					success = false
				}
			}
		}
		entriesCopy2 = entriesCopy2[1:]
	}
}

func buildCentsStrAndParts(
	cents int,
	a string,
	negative bool,
	delim1 string,
	delim2 string,
	delim3 string,
) string {
	centsStr := getCentsStr(cents)
	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}

	if len(rest) > 0 {
		parts = append(parts, rest)
	}

	for i := len(parts) - 1; i >= 0; i-- {
		a += parts[i] + delim1
	}

	a = a[:len(a)-1]
	a += delim2
	a += centsStr[len(centsStr)-2:]
	if negative {
		a += delim3
	} else {
		a += " "
	}
	return a
}

func getDate(locale string, entry Entry) string {
	var date string
	if locale == "nl-NL" {
		date = entry.Date[8:10] + "-" + entry.Date[5:7] + "-" + entry.Date[0:4]
	} else if locale == "en-US" {
		date = entry.Date[5:7] + "/" + entry.Date[8:10] + "/" + entry.Date[0:4]
	}
	return date
}

func handleCurrency(a string, currency string, channel chan ChanEntry) string {
	switch currency {
	case "EUR":
		a += "€"
	case "USD":
		a += "$"
	default:
		channel <- ChanEntry{e: errors.New("")}
	}
	return a
}

func buildA(
	locale, currency string,
	channel chan ChanEntry,
	cents int,
	negative bool,
) (string, int) {
	var a string
	switch locale {
	case "nl-NL":
		a = handleCurrency(a, currency, channel)
		a += " "
		a = buildCentsStrAndParts(cents, a, negative, ".", ",", "-")
	case "en-US":
		if negative {
			a += "("
		}
		a = handleCurrency(a, currency, channel)
		a = buildCentsStrAndParts(cents, a, negative, ",", ".", ")")
	default:
		channel <- ChanEntry{e: errors.New("")}
	}

	var al int
	for range a {
		al++
	}
	return a, al
}
