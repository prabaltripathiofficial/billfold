package invoice

import (
	"path/filepath"
	"testing"
	"time"
)

func TestParseAmount(t *testing.T) {
	cases := map[string]int64{"12": 1200, "12.5": 1250, "12.05": 1205, "0.99": 99, ".5": 50}
	for in, want := range cases {
		got, err := ParseAmount(in)
		if err != nil || got != want {
			t.Errorf("ParseAmount(%q) = %d, %v; want %d", in, got, err, want)
		}
	}
	for _, bad := range []string{"", "abc", "-3", "1.234", "1."} {
		if _, err := ParseAmount(bad); err == nil {
			t.Errorf("ParseAmount(%q): want error", bad)
		}
	}
}

func TestStoreRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "invoices.json")
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	a := s.Add(Invoice{Client: "Globex", Amount: 1000, Currency: "USD", Issued: time.Now()})
	b := s.Add(Invoice{Client: "Initech", Amount: 2000, Currency: "USD", Issued: time.Now()})
	if a.ID != 1 || b.ID != 2 {
		t.Fatalf("ids: %d %d", a.ID, b.ID)
	}
	if err := s.MarkPaid(2); err != nil {
		t.Fatal(err)
	}
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}
	again, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(again.Invoices) != 2 || !again.Invoices[1].Paid {
		t.Fatalf("reloaded: %+v", again.Invoices)
	}
}

func TestTotalsForMonth(t *testing.T) {
	invs := []Invoice{
		{Amount: 1000, Currency: "USD", Issued: time.Date(2026, 10, 3, 12, 0, 0, 0, time.UTC)},
		{Amount: 2500, Currency: "USD", Issued: time.Date(2026, 10, 20, 12, 0, 0, 0, time.UTC)},
		{Amount: 700, Currency: "EUR", Issued: time.Date(2026, 10, 9, 12, 0, 0, 0, time.UTC)},
		{Amount: 9900, Currency: "USD", Issued: time.Date(2026, 11, 2, 12, 0, 0, 0, time.UTC)},
	}
	got := TotalsForMonth(invs, "2026-10")
	if len(got) != 2 || got[0].Currency != "EUR" || got[1].Amount != 3500 || got[1].Count != 2 {
		t.Fatalf("unexpected totals: %+v", got)
	}
}
