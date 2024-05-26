package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDiagnose(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     []Diagnostic
	}{
		{
			name: "no diagnostics",
			template: `
package main

t1 template () {
	<p>Hello, World!</p>
}`,
			want: nil,
		},

		// useOfLegacyCallSyntaxDiagnoser

		{
			name: "useOfLegacyCallSyntaxDiagnoser: template root",
			template: `
package main

t1 template () {
	{! t1.Raw("foo") }
}`,
			want: []Diagnostic{{
				Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
				Range:   Range{Position{39, 4, 4}, Position{55, 4, 20}},
			}},
		},
		{
			name: "useOfLegacyCallSyntaxDiagnoser: in div",
			template: `
package main

t1 template () {
	<div>
		{! t1.Raw("foo") }
	</div>
}`,
			want: []Diagnostic{{
				Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
				Range:   Range{Position{47, 5, 5}, Position{63, 5, 21}},
			}},
		},
		{
			name: "useOfLegacyCallSyntaxDiagnoser: in if",
			template: `
package main

t1 template () {
	if true {
		{! t1.Raw("foo") }
	}
}`,
			want: []Diagnostic{{
				Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
				Range:   Range{Position{51, 5, 5}, Position{67, 5, 21}},
			}},
		},
		{
			name: "useOfLegacyCallSyntaxDiagnoser: in for",
			template: `
package main

t1 template () {
	for i := range x {
		{! t1.Raw("foo") }
	}
}`,
			want: []Diagnostic{{
				Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
				Range:   Range{Position{60, 5, 5}, Position{76, 5, 21}},
			}},
		},
		{
			name: "useOfLegacyCallSyntaxDiagnoser: in switch",
			template: `
package main

t1 template () {
	switch x {
	case 1:
		{! t1.Raw("foo") }
	default:
		{! x }
	}
}`,
			want: []Diagnostic{
				{
					Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
					Range:   Range{Position{61, 6, 5}, Position{77, 6, 21}},
				},
				{
					Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
					Range:   Range{Position{95, 8, 5}, Position{96, 8, 6}},
				},
			},
		},
		{
			name: "useOfLegacyCallSyntaxDiagnoser: in block",
			template: `
package main

t1 template () {
	@layout("Home") {
		{! t1.Raw("foo") }
	}
}`,
			want: []Diagnostic{{
				Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `t1 fmt .` to fix all instances.",
				Range:   Range{Position{59, 5, 5}, Position{75, 5, 21}},
			}},
		},
		{
			name: "voidElementWithChildrenDiagnoser: no diagnostics",
			template: `
package main

t1 template () {
	<div>
		<input/>
	</div>
}`,
			want: nil,
		},
		{
			name: "voidElementWithChildrenDiagnoser: with diagnostics",
			template: `
package main

t1 template () {
	<div>
	  <input>Child content</input>
	</div>
}`,
			want: []Diagnostic{{
				Message: "void element <input> should not have child content",
				Range:   Range{Position{46, 5, 4}, Position{51, 5, 9}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf, err := ParseString(tt.template)
			if err != nil {
				t.Fatalf("ParseTemplateFile() error = %v", err)
			}
			got, err := Diagnose(tf)
			if err != nil {
				t.Fatalf("Diagnose() error = %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Diagnose() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
