package main

import (
	"github.com/zarevucky/astgen"
	"io/ioutil"
	"os"
	"sort"
	"fmt"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	langdef, err := astgen.Load(file)
	if err != nil {
		panic(err)
	}

	var sortedTypes sort.StringSlice

	for s, _ := range langdef.Types {
		sortedTypes = append(sortedTypes, s)
	}

	// TODO: This sorting can be done by the library.

	sort.Sort(sortedTypes)

	emitVisitorInterface(langdef, sortedTypes)

	for _, s := range sortedTypes {
		emitZeframVisitor(langdef, langdef.Types[s])
	}
}



func emitVisitorInterface(l *astgen.LangDef, sortedTypes sort.StringSlice) {
	fmt.Printf("type Visitor = interface {\n")
	
	for _, s := range sortedTypes {
		t := l.Types[s]
		
		switch tt := t.(type) {
		case *astgen.LexicalType:
		case *astgen.EnumType:
		case *astgen.OptionType:
			// nothing
			
		case *astgen.StructType:
			fmt.Printf("\tpreprocess_%s(node: *AST%s) (enter: bool)\n", strings.ToLower(tt.Name), tt.Name)
			fmt.Printf("\tpostprocess_%s(node: *AST%s)\n", strings.ToLower(tt.Name), tt.Name)
		}
	}
	
	fmt.Printf("}\n")
}

func emitZeframVisitor(l *astgen.LangDef, t astgen.Type) {
	switch tt := t.(type) {
	case *astgen.LexicalType:
	case *astgen.EnumType:
		// nothing
	case *astgen.OptionType:
		emitOptionVisitor(l, tt)
	case *astgen.StructType:
		emitStructVisitor(tt)
	}
}
