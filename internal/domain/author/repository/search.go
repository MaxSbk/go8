package repository

import (
	"context"
	"fmt"

	"github.com/gmhafiz/go8/ent/gen"
	entAuthor "github.com/gmhafiz/go8/ent/gen/author"
	"github.com/gmhafiz/go8/ent/gen/predicate"
	"github.com/gmhafiz/go8/internal/domain/author"
)

func NewSearch(db *gen.Client) *repository {
	return &repository{ent: db}
}

// Search using the same store. May use other store e.g. elasticsearch/bleve as
// the search repository.
func (r *repository) Search(ctx context.Context, f *author.Filter) ([]*gen.Author, int, error) {
	var predicateUser []predicate.Author
	if f.FirstName != "" {
		predicateUser = append(predicateUser, entAuthor.FirstNameContainsFold(f.FirstName))
	}
	if f.MiddleName != "" {
		predicateUser = append(predicateUser, entAuthor.MiddleNameContainsFold(f.MiddleName))
	}
	if f.LastName != "" {
		predicateUser = append(predicateUser, entAuthor.LastNameContainsFold(f.LastName))
	}

	// sort by field
	var orderFunc []gen.OrderFunc
	for col, ord := range f.Base.Sort {
		if ord == "ASC" {
			orderFunc = append(orderFunc, gen.Asc(col))
		} else {
			orderFunc = append(orderFunc, gen.Desc(col))
		}
	}

	total, err := r.ent.Author.Query().
		Where(entAuthor.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("error retrieving Author list: %w", err)
	}

	// The section where the search is done
	//
	// ILIKE is to search for case-insensitive in postgres
	//
	// Speed optimization is possible by creating an index on lowered names
	// Reference: https://www.postgresql.org/docs/current/indexes-expressional.html
	// CREATE INDEX on authors (LOWER(first_name));
	// CREATE INDEX on authors (LOWER(middle_name));
	// CREATE INDEX on authors (LOWER(last_name));
	//
	// Second alternative is to use ~*
	// mods = append(mods, qm.Or(models.AuthorColumns.MiddleName+" ~* ?", f.Name))
	//
	// postgres has a builtin full-text search: https://www.postgresql.org/docs/current/textsearch.html
	//
	// Also, may use term frequency-inverted index search (tf-idf) like
	// elasticsearch or bleve.
	authors, err := r.ent.Author.Query().
		WithBooks().
		Where(predicateUser...).
		Where(entAuthor.DeletedAtIsNil()).
		Limit(f.Base.Limit).
		Offset(f.Base.Offset).
		Order(orderFunc...).
		All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("error retrieving Author list: %w", err)
	}

	return authors, total, nil
}
