grammar GrammarDesc;

rules: (lexer_rule|skip_rule|start_rule|parser_rule)*;

lexer_rule: TERM_ID ':' REGEXP ';';

skip_rule: 'skip' ':' REGEXP ';';

start_rule: 'start' ':' NON_TERM_ID ';';

parser_rule: NON_TERM_ID '->' (term_parser_rule_item|non_term_parser_rule_item|code_item)* ';';

term_parser_rule_item: (ID '=')? TERM_ID;

non_term_parser_rule_item: (ID '=')? NON_TERM_ID ATTR?;

code_item: CODE;

ID : [a-zA-Z] [a-zA-Z0-9]*;
TERM_ID : '$' [A-Z] [A-Z0-9]*;
NON_TERM_ID : '$' [a-z] [a-z0-9]*;

REGEXP: '"' .*? '"';
ATTR: '!' .*? '!';
CODE: '{' .*? '}';

WS : [ \n\t\r] -> skip;
