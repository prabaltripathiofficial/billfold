package invoice

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type MonthTotal struct {
	Currency string
	Amount   int64
	Count    int
}

func TotalsForMonth(invoices []Invoice, month string) []MonthTotal {
	byCurrency := map[string]*MonthTotal{}
	for _, inv := range invoices {
		if inv.Issued.UTC().Format("2006-01") != month {
			continue
		}
		t, ok := byCurrency[inv.Currency]
		if !ok {
			t = &MonthTotal{Currency: inv.Currency}
			byCurrency[inv.Currency] = t
		}
		t.Amount += inv.Amount
		t.Count++
	}
	out := make([]MonthTotal, 0, len(byCurrency))
	for _, t := range byCurrency {
		out = append(out, *t)
	}
	sort.Slice(out, func(a, b int) bool { return out[a].Currency < out[b].Currency })
	return out
}

func ParseAmount(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("invalid amount: empty")
	}
	whole, frac, hasFrac := strings.Cut(s, ".")
	if whole == "" {
		whole = "0"
	}
	w, err := strconv.ParseInt(whole, 10, 64)
	if err != nil || w < 0 {
		return 0, fmt.Errorf("invalid amount %q", s)
	}
	cents := int64(0)
	if hasFrac {
		if len(frac) == 0 || len(frac) > 2 {
			return 0, fmt.Errorf("invalid amount %q: use at most two decimals", s)
		}
		if len(frac) == 1 {
			frac += "0"
		}
		c, err := strconv.ParseInt(frac, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid amount %q", s)
		}
		cents = c
	}
	return w*100 + cents, nil
}
