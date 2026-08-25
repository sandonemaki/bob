package orm

import (
	"context"
	"io"

	"github.com/stephenafamo/bob"
)

// CountPreloadColumn is the SELECT-list entry a count Preloader (see the
// generated per-relation CountPreloader, gen/templates/counts) appends to
// a query's preload columns via bob.Query's AppendPreloadSelect. Exposing
// Name and Alias in typed, exported fields lets a caller that walks
// clause.SelectList.PreloadColumns (a public field, clause/select.go)
// recognize a declared count preload without parsing SQL text.
type CountPreloadColumn struct {
	// Name is the relationship name PreloadCount(name, count) matches
	// against in the generated model's switch.
	Name string
	// Alias is the SQL column alias assigned to Expr: always
	// "__count_" + Name.
	Alias string
	// Expr is the correlated subquery expression itself.
	Expr bob.Expression
}

// WriteSQL renders "<Expr> AS "<Alias>"" — byte-identical to what the
// previous unexported aliasedExpr produced.
func (c CountPreloadColumn) WriteSQL(ctx context.Context, w io.StringWriter, d bob.Dialect, start int) ([]any, error) {
	args, err := c.Expr.WriteSQL(ctx, w, d, start)
	if err != nil {
		return nil, err
	}
	w.WriteString(" AS ")
	d.WriteQuoted(w, c.Alias)
	return args, nil
}
