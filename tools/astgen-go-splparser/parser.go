// This file was generated from goparser.go.tmpl, DO NOT MODIFY.

package main
import "fmt"
import "github.com/zarevucky/astgen"
func emitOptionType(t *astgen.OptionType) {
fmt.Print("func ParseAST")
fmt.Printf("%v",  t.Name )
fmt.Print("(p *spl.SeqParser) (ret AST")
fmt.Printf("%v",  t.Name )
fmt.Print(") {\n")
fmt.Print("\tif !p.IsList() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tp.Down()\n")
fmt.Print("\t\n")
fmt.Print("\tif !p.IsString() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\ttag := p.String()\n")
fmt.Print("\t\n")
fmt.Print("\tp.Next()\n")
fmt.Print("\t\n")
fmt.Print("\tif p.IsEnd() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tswitch tag {\n")
for _, tn := range t.ConcreteTypes() {
fmt.Print("\tcase \"")
fmt.Printf("%v",  tn )
fmt.Print("\":\n")
fmt.Print("\t\tret = ParseAST")
fmt.Printf("%v",  tn )
fmt.Print("(p)\n")
}
fmt.Print("\tdefault:\n")
fmt.Print("\t\tpanic(\"missing case for \" + tag)\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tp.Next()\n")
fmt.Print("\tif !p.IsEnd() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\tp.Up()\n")
fmt.Print("\t\n")
fmt.Print("\treturn ret\n")
fmt.Print("}\n")
}
func emitEnumType(t *astgen.EnumType) {
fmt.Print("func ParseAST")
fmt.Printf("%v",  t.Name )
fmt.Print("(p *spl.SeqParser) (ret AST")
fmt.Printf("%v",  t.Name )
fmt.Print(") {\n")
fmt.Print("\tif !p.IsString() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\ttag := p.String()\n")
fmt.Print("\tp.Next()\n")
fmt.Print("\t\n")
fmt.Print("\tswitch tag {\n")
for _, tok := range t.EnumTokens {
fmt.Print("\tcase \"")
fmt.Printf("%v",  tok.Name )
fmt.Print("\":\n")
fmt.Print("\t\treturn AST")
fmt.Printf("%v",  t.Name )
fmt.Print("_")
fmt.Printf("%v",  tok.Name )
fmt.Print("\n")
}
fmt.Print("\tdefault:\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("}\n")
}
func goTypeName(t astgen.Type) string {
	switch t.(type) {
	case *astgen.StructType:
		return "*AST" + t.Common().Name
	case *astgen.LexicalType:
		return "string"
	case *astgen.BoolType:
		return "bool"
	default:
		return "AST" + t.Common().Name
	}
}
func emitStructType(t *astgen.StructType) {
fmt.Print("func ParseAST")
fmt.Printf("%v",  t.Name )
fmt.Print("(p *spl.SeqParser) (ret *AST")
fmt.Printf("%v",  t.Name )
fmt.Print(") {\n")
fmt.Print("\tif !p.IsList() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tp.Down()\n")
fmt.Print("\t\n")
fmt.Print("\tif p.IsEnd() {\n")
fmt.Print("\t\tp.Up()\n")
fmt.Print("\t\treturn nil\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tret = new(AST")
fmt.Printf("%v",  t.Name )
fmt.Print(")\n")
for _, m := range t.Members {
fmt.Print("\t\n")
fmt.Print("\tif p.IsEnd() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
	if m.Array {
fmt.Print("\tret._")
fmt.Printf("%v",  m.Name )
fmt.Print(" = ParseAST")
fmt.Printf("%v",  t.Name )
fmt.Print("_")
fmt.Printf("%v",  m.Name )
fmt.Print("(p)\n")
	} else if m.Type.Kind() == "Bool" {
fmt.Print("\tret._")
fmt.Printf("%v",  m.Name )
fmt.Print(" = p.IsString()\n")
	} else if m.Type.Kind() == "Lexical" {
		if m.Nullable {
fmt.Print("\tif p.IsString() {\n")
fmt.Print("\t\ts := p.String()\n")
fmt.Print("\t\tret._")
fmt.Printf("%v",  m.Name )
fmt.Print(" = &s\n")
fmt.Print("\t}\n")
		} else {
fmt.Print("\tret._")
fmt.Printf("%v",  m.Name )
fmt.Print(" = p.String()\n")
		}
	} else {
fmt.Print("\tret._")
fmt.Printf("%v",  m.Name )
fmt.Print(" = ParseAST")
fmt.Printf("%v",  m.Type.Common().Name )
fmt.Print("(p)\n")
	}
fmt.Print("\t\n")
fmt.Print("\tp.Next()\n")
}
fmt.Print("\t\n")
fmt.Print("\tif !p.IsEnd() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\tp.Up()\n")
fmt.Print("\t\n")
fmt.Print("\treturn\n")
fmt.Print("}\n")
for _, m := range t.Members {
	if !m.Array {
		continue
	}

	tn := goTypeName(m.Type)

fmt.Print("func ParseAST")
fmt.Printf("%v",  t.Name )
fmt.Print("_")
fmt.Printf("%v",  m.Name )
fmt.Print("(p *spl.SeqParser) (ret []")
fmt.Printf("%v",  tn )
fmt.Print(") {\n")
fmt.Print("\tif !p.IsList() {\n")
fmt.Print("\t\tpanic(\"bad file\")\n")
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tret = make([]")
fmt.Printf("%v",  tn )
fmt.Print(")\n")
fmt.Print("\t\n")
fmt.Print("\tfor p.Down(); !p.IsEnd(); p.Next() {\n")
	if m.Type.Kind() == "Lexical" {
fmt.Print("\t\tret = append(ret, p.String())\n")
	} else {
fmt.Print("\t\tret = append(ret, ParseAST")
fmt.Printf("%v",  m.Type.Common().Name )
fmt.Print("(p))\n")
	}
fmt.Print("\t}\n")
fmt.Print("\t\n")
fmt.Print("\tp.Up()\n")
fmt.Print("\t\n")
fmt.Print("\treturn ret\n")
fmt.Print("}\n")
}
}