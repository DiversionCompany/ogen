package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ogen-go/ogen"
	ast "github.com/ogen-go/ogen/internal/ast2"
)

func TestSchemaSimple(t *testing.T) {
	gen := &schemaGen{
		localRefs: make(map[string]*ast.Schema),
	}

	out, err := gen.Generate(ogen.Schema{
		Type: "object",
		Properties: map[string]ogen.Schema{
			"id":   {Type: "integer"},
			"name": {Type: "string"},
		},
		Required: []string{"id", "name"},
	})
	require.NoError(t, err)

	expect := &ast.Schema{
		Type: ast.Object,
		Fields: []ast.SchemaField{
			{
				Name: "id",
				Type: &ast.Schema{Type: ast.Integer},
			},
			{
				Name: "name",
				Type: &ast.Schema{Type: ast.String},
			},
		},
	}

	require.Equal(t, expect, out)
}

func TestSchemaRecursive(t *testing.T) {
	spec := &ogen.Spec{
		Components: &ogen.Components{
			Schemas: map[string]ogen.Schema{
				"Pet": {
					Type: "object",
					Properties: map[string]ogen.Schema{
						"id":   {Type: "integer"},
						"name": {Type: "string"},
						"friends": {
							Type: "array",
							Items: &ogen.Schema{
								Ref: "#/components/schemas/Pet",
							},
						},
					},
					Required: []string{"id", "name", "friends"},
				},
			},
		},
	}

	pet := &ast.Schema{
		Type: ast.Object,
		Ref:  "#/components/schemas/Pet",
	}
	pet.Fields = []ast.SchemaField{
		{
			Name: "friends",
			Type: &ast.Schema{
				Type: ast.Array,
				Item: pet,
			},
		},
		{
			Name: "id",
			Type: &ast.Schema{Type: ast.Integer},
		},
		{
			Name: "name",
			Type: &ast.Schema{Type: ast.String},
		},
	}

	expectLocalRefs := map[string]*ast.Schema{
		"#/components/schemas/Pet": {
			Type: ast.Object,
			Ref:  "#/components/schemas/Pet",
			Fields: []ast.SchemaField{
				{
					Name: "friends",
					Type: &ast.Schema{
						Type: ast.Array,
						Item: pet,
					},
				},
				{
					Name: "id",
					Type: &ast.Schema{Type: ast.Integer},
				},
				{
					Name: "name",
					Type: &ast.Schema{Type: ast.String},
				},
			},
		},
	}

	gen := &schemaGen{
		spec:      spec,
		localRefs: make(map[string]*ast.Schema),
	}

	out, err := gen.Generate(ogen.Schema{
		Ref: "#/components/schemas/Pet",
	})
	require.NoError(t, err)
	require.Equal(t, expectLocalRefs, gen.localRefs)
	require.Equal(t, pet, out)
}

func TestSchemaSideEffects(t *testing.T) {
	expectSide := []*ast.Schema{
		{
			Type: ast.Object,
			Fields: []ast.SchemaField{
				{
					Name: "age",
					Type: &ast.Schema{Type: ast.Integer},
				},
				{
					Name: "id",
					Type: &ast.Schema{Type: ast.Integer},
				},
				{
					Name: "name",
					Type: &ast.Schema{Type: ast.String},
				},
			},
		},
	}

	expect := &ast.Schema{
		Type: ast.Object,
		Fields: []ast.SchemaField{
			{
				Name: "name",
				Type: &ast.Schema{Type: ast.String},
			},
			{
				Name: "owner",
				Type: expectSide[0],
			},
		},
	}

	gen := &schemaGen{
		localRefs: make(map[string]*ast.Schema),
	}

	out, err := gen.Generate(ogen.Schema{
		Type: "object",
		Properties: map[string]ogen.Schema{
			"name": {Type: "string"},
			"owner": {
				Type: "object",
				Properties: map[string]ogen.Schema{
					"name": {Type: "string"},
					"id":   {Type: "integer"},
					"age":  {Type: "integer"},
				},
				Required: []string{"name", "id", "age"},
			},
		},
		Required: []string{"id", "name", "owner"},
	})

	require.NoError(t, err)
	require.Equal(t, expect, out)
}

func TestSchemaReferencedArray(t *testing.T) {
	spec := &ogen.Spec{
		Components: &ogen.Components{
			Schemas: map[string]ogen.Schema{
				"Pets": {
					Type: "array",
					Items: &ogen.Schema{
						Type: "string",
					},
				},
			},
		},
	}

	pets := &ast.Schema{
		Type: ast.Array,
		Ref:  "#/components/schemas/Pets",
		Item: &ast.Schema{Type: ast.String},
	}

	expectLocalRefs := map[string]*ast.Schema{
		"#/components/schemas/Pets": pets,
	}

	expect := &ast.Schema{
		Type: ast.Object,
		Fields: []ast.SchemaField{
			{
				Name: "pets",
				Type: pets,
			},
		},
	}

	gen := &schemaGen{
		spec:      spec,
		localRefs: make(map[string]*ast.Schema),
	}

	out, err := gen.Generate(ogen.Schema{
		Type: "object",
		Properties: map[string]ogen.Schema{
			"pets": {
				Ref: "#/components/schemas/Pets",
			},
		},
		Required: []string{"pets"},
	})

	require.NoError(t, err)
	require.Equal(t, expectLocalRefs, gen.localRefs)
	require.Equal(t, expect, out)
}
