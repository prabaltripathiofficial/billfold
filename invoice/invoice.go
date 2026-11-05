package invoice

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Invoice struct {
	ID       int       `json:"id"`
	Client   string    `json:"client"`
	Amount   int64     `json:"amount"`
	Currency string    `json:"currency"`
	Issued   time.Time `json:"issued"`
	Paid     bool      `json:"paid"`
	Note     string    `json:"note,omitempty"`
}

func (i Invoice) Total() string {
	return fmt.Sprintf("%d.%02d %s", i.Amount/100, i.Amount%100, i.Currency)
}

type Store struct {
	path     string
	Invoices []Invoice `json:"invoices"`
}

func DefaultPath() string {
	if p := os.Getenv("BILLFOLD_FILE"); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "billfold.json"
	}
	return filepath.Join(home, ".billfold", "invoices.json")
}

func Open(path string) (*Store, error) {
	s := &Store{path: path}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, s); err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return s, nil
}

func (s *Store) Add(inv Invoice) Invoice {
	next := 1
	for _, existing := range s.Invoices {
		if existing.ID >= next {
			next = existing.ID + 1
		}
	}
	inv.ID = next
	s.Invoices = append(s.Invoices, inv)
	return inv
}

func (s *Store) MarkPaid(id int) error {
	for i := range s.Invoices {
		if s.Invoices[i].ID == id {
			s.Invoices[i].Paid = true
			return nil
		}
	}
	return fmt.Errorf("no invoice with id %d", id)
}

func (s *Store) Sorted() []Invoice {
	out := append([]Invoice(nil), s.Invoices...)
	sort.Slice(out, func(a, b int) bool { return out[a].ID < out[b].ID })
	return out
}

func (s *Store) Save() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o644)
}
