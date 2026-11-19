// Package conventions holds tests that assert repository-wide conventions
// rather than the behaviour of any one package.
package conventions

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestQueryStructFieldsHaveURITags asserts that every field of every *Query
// struct under pkg/ carries a uri tag.
//
// Query structs reach the server two ways: the newclient functions expand them
// through uritemplates, which reads the uri tag, and the older services encode
// them with go-querystring, which reads url. A field with only a url tag is
// silently dropped by the newclient path — the request succeeds and the filter
// does not apply. That is how actiontemplates.Get shipped broken (#437), and how
// DashboardQuery.IncludeLatest sat unusable for four years.
func TestQueryStructFieldsHaveURITags(t *testing.T) {
	var files []string
	err := filepath.WalkDir(filepath.Join("..", "..", "pkg"), func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			files = append(files, path)
		}
		return nil
	})
	require.NoError(t, err)
	require.NotEmpty(t, files, "found no Go files under pkg/")

	fset := token.NewFileSet()
	structsChecked := 0

	for _, path := range files {
		parsed, err := parser.ParseFile(fset, path, nil, 0)
		require.NoErrorf(t, err, "parsing %s", path)

		ast.Inspect(parsed, func(node ast.Node) bool {
			typeSpec, ok := node.(*ast.TypeSpec)
			if !ok || !strings.HasSuffix(typeSpec.Name.Name, "Query") {
				return true
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				return true
			}

			structsChecked++
			for _, field := range structType.Fields.List {
				if len(field.Names) == 0 {
					continue // embedded
				}

				var tag reflect.StructTag
				if field.Tag != nil {
					unquoted, err := strconv.Unquote(field.Tag.Value)
					require.NoError(t, err)
					tag = reflect.StructTag(unquoted)
				}

				if _, hasURI := tag.Lookup("uri"); hasURI {
					continue
				}

				reason := "has no struct tag"
				if _, hasURL := tag.Lookup("url"); hasURL {
					reason = "has a url tag but no uri tag, so the newclient path drops it"
				}

				for _, name := range field.Names {
					t.Errorf("%s: %s.%s %s", path, typeSpec.Name.Name, name.Name, reason)
				}
			}

			return true
		})
	}

	require.NotZero(t, structsChecked, "found no *Query structs to check")
	t.Logf("checked %d *Query structs across %d files", structsChecked, len(files))
}
