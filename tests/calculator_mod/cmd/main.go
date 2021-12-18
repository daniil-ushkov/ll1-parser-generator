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
$LOG: "//";
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


$f -> l=$l l1=$l1!"acc" : l["val"]! {res["val"] = l1["val"]};
$l1 -> $LOG f=$f {res["val"] = log(attr["acc"].(int), f["val"].(int))}; 
$l1 -> {res["val"] = attr["acc"]};

$l -> $LP e=$e $RP {res["val"] = e["val"]};
$l -> n=$NUM {res["val"], _ = strconv.Atoi(n)};
$l -> $SUB n=$NUM {
x, _ := strconv.Atoi(n)
res["val"] = -x
};
`

func main() {
	desc, _ := grammar.ParseDesc(GrammarDesc)
	gr := grammar.New(desc)
	ctx := generator.Context{
		Package: "parser",
		Out:     "tests/calculator_mod/parser",
		Grammar: gr,
	}

	_ = ctx.GenerateAstGo()
	_ = ctx.GenerateLexerGo()
	_ = ctx.GenerateAttrParserGo()
	_ = ctx.GenerateSymbolsGo()
}
