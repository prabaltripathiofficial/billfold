package export

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/prabaltripathiofficial/billfold/invoice"
)

var header = []string{"id", "client", "amount", "currency", "issued", "paid", "note"}

func Write(w io.Writer, format string, invoices []invoice.Invoice) error {
	switch format {
	case "csv":
		return CSV(w, invoices)
	case "json":
		return JSON(w, invoices)
	default:
		return fmt.Errorf("unknown format %q (use csv or json)", format)
	}
}

func CSV(w io.Writer, invoices []invoice.Invoice) error {
	if _, err := fmt.Fprintln(w, strings.Join(header, ",")); err != nil {
		return err
	}
	for _, inv := range invoices {
		row := []string{
			fmt.Sprint(inv.ID),
			inv.Client,
			fmt.Sprintf("%d.%02d", inv.Amount/100, inv.Amount%100),
			inv.Currency,
			inv.Issued.Format("2006-01-02"),
			fmt.Sprint(inv.Paid),
			inv.Note,
		}
		if _, err := fmt.Fprintln(w, strings.Join(row, ",")); err != nil {
			return err
		}
	}
	return nil
}

func JSON(w io.Writer, invoices []invoice.Invoice) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if invoices == nil {
		invoices = []invoice.Invoice{}
	}
	return enc.Encode(invoices)
}
