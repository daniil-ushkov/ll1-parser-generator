package main

import (
	"ll1-parser-generator/generator"
	"ll1-parser-generator/grammar"
)

const GrammarDesc = `
$NUM : "\\d+";
$PLUS: "\\+";
$SUB: "-";
$MUL: "\\*";
$DIV: "/";
$LP: "\\(";
$RP: "\\)";

skip: "[ \t\n]+";

start: $s;

$s -> e=$e {res["val"] = e["val"]};

$e -> t=$t 
	  e1=$e1!"acc" : t["val"]! 
	  {res["val"] = e1["val"]};

$e1 -> $PLUS 
       t=$t 
	   e1=$e1!"acc" : t["val"].(int) + attr["acc"].(int)! 
       {res["val"] = e1["val"]};

$e1 -> $SUB 
       t=$t 
	   e1=$e1!"acc" : attr["acc"].(int) - t["val"].(int)! 
       {res["val"] = e1["val"]};

$e1 -> {res["val"] = attr["acc"]};

$t -> f=$f 
      t1=$t1!"acc" : f["val"]! 
	  {res["val"] = t1["val"]};

$t1 -> $MUL 
	   f=$f 
       t1=$t1!"acc" : f["val"].(int) * attr["acc"].(int)!
       {res["val"] = t1["val"]};

$t1 -> $DIV 
	   f=$f 
       t1=$t1!"acc" : attr["acc"].(int) / f["val"].(int)!
       {res["val"] = t1["val"]};

$t1 -> {res["val"] = attr["acc"]};

$f -> $LP e=$e $RP {res["val"] = e["val"]};
$f -> n=$NUM {res["val"], _ = strconv.Atoi(n)};
`

func main() {
	desc, _ := grammar.ParseDesc(GrammarDesc)
	gr := grammar.New(desc)
	ctx := generator.Context{
		Package: "parser",
		Out:     "tests/calculator/parser",
		Grammar: gr,
	}

	_ = ctx.GenerateAstGo()
	_ = ctx.GenerateLexerGo()
	_ = ctx.GenerateAttrParserGo()
	_ = ctx.GenerateSymbolsGo()
}
