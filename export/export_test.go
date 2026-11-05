package export

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"testing"
	"time"

	"github.com/prabaltripathiofficial/billfold/invoice"
)

func sample() []invoice.Invoice {
	return []invoice.Invoice{
		{ID: 1, Client: "Globex", Amount: 125000, Currency: "USD", Issued: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), Paid: true},
		{ID: 2, Client: "Initech", Amount: 9950, Currency: "EUR", Issued: time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC), Note: "design review"},
	}
}

func TestCSVHeaderAndRows(t *testing.T) {
	var buf bytes.Buffer
	if err := CSV(&buf, sample()); err != nil {
		t.Fatal(err)
	}
	rows, err := csv.NewReader(&buf).ReadAll()
	if err != nil {
		t.Fatalf("output is not valid CSV: %v", err)
	}
	if len(rows) != 3 {
		t.Fatalf("want 3 rows, got %d", len(rows))
	}
	if got := rows[1][2]; got != "1250.00" {
		t.Errorf("amount: want 1250.00, got %s", got)
	}
	if got := rows[2][6]; got != "design review" {
		t.Errorf("note: want %q, got %q", "design review", got)
	}
}

func TestJSONRoundTrip(t *testing.T) {
	var buf bytes.Buffer
	if err := JSON(&buf, sample()); err != nil {
		t.Fatal(err)
	}
	var back []invoice.Invoice
	if err := json.Unmarshal(buf.Bytes(), &back); err != nil {
		t.Fatal(err)
	}
	if len(back) != 2 || back[1].Client != "Initech" {
		t.Fatalf("unexpected round trip: %+v", back)
	}
}

func TestUnknownFormat(t *testing.T) {
	if err := Write(&bytes.Buffer{}, "xml", nil); err == nil {
		t.Fatal("want error for unknown format")
	}
}
