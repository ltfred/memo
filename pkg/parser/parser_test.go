package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	par := Parser{}

	t.Run("add", func(t *testing.T) {
		err := par.Add(Memo{
			Name:     "test_add",
			Date:     "2023",
			Content:  "test_add",
			Priority: "high",
			CreateAt: 1684480661,
		})
		if err != nil {
			t.Fatalf("add err: %v", err)
		}
	})

	t.Run("show", func(t *testing.T) {
		show, err := par.Show()
		if err != nil {
			t.Fatalf("show err: %v", err)
		} else {
			t.Logf("records: %v", show)
		}
	})

	t.Run("modify", func(t *testing.T) {
		err := par.Modify("1684480661", Memo{
			Date: "2024",
		})
		if err != nil {
			t.Fatalf("modify err: %v", err)
		}
	})

	t.Run("delete", func(t *testing.T) {
		err := par.Delete("1684480661")
		if err != nil {
			t.Fatalf("delete err: %v", err)
		}
	})
}
