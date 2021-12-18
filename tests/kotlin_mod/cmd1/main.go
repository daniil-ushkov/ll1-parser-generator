package main

import "fmt"

const GrammarDesc = `
$FUN : "fun";
$ID  : "[a-zA-Z][a-zA-Z0-9]*";
$COL : ":";
$COM : ",";
$LP  : "\\(";
$RP  : "\\)";

skip: "[ \t\n]+";

start: $sig;

$sig -> $FUN $ID $LP $params $RP $ret;

$params -> $param $params1;
$params -> ;

$param -> $ID $COL $ID;

$params1 -> $COM $params2;
$params1 -> ;

$params2 -> $param $params1;
$params2 -> ;

$ret -> $COL $ID;
$ret -> ;
`

func log(val1, val2 int) int {
	res := 0
	for val1 != 1 {
		val1 /= val2
		res++
	}
	return res
}

func main() {
	//desc, _ := grammar.ParseDesc(GrammarDesc)
	//gr := grammar.New(desc)
	//ctx := generator.Context{
	//	Package: "parser",
	//	Out:     "tests/kotlin_mod/parser",
	//	Grammar: gr,
	//}
	//
	//_ = ctx.GenerateAstGo()
	//_ = ctx.GenerateLexerGo()
	//_ = ctx.GenerateAstParserGo()
	//_ = ctx.GenerateSymbolsGo()

	fmt.Println(log(81, log(8, 2)))
}
