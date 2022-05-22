// the grammar for tdat RuleEngine
grammar Tdat;

// the first parser rule, also the first rule of RuleEngine grammar
// prog is a sequence of rule lines.

prog
    : ruleLine+
    ;

ruleLine
    : ruleID ':' enumerableFunc '{' windowsRange conditionExpr '}' '=>' result ';'
    ;

ruleID
    : ID
    ; 

enumerableFunc
    : 'Each'
    | 'None'
    | 'Any'
    ; 

windowsRange
    : '|' '|'               #WindowsWithZeroIndex
    | '|' INT? ',' INT? '|' #WindowsWithLowAndHighIndex
    ;

conditionExpr
    : conditionExpr logicalOp conditionExpr
    | '(' conditionExpr ')'
    | primaryExpr comparisonOp primaryExpr
    ;

primaryExpr
    : '(' primaryExpr ')'                  #BracketExprInPrimaryExpr
    | primaryExpr arithmeticOp primaryExpr #ArithmeticExprInPrimaryExpr
    | METRIC                               #MetricInPrimaryExpr
    | builtin '(' primaryExpr ')'          #BuildinExprInPrimaryExpr
    | literal                              #RightLiteralInPrimaryExpr
    ;

arithmeticOp
    : '+'
    | '-'
    | '*'
    | '/'
    | '%'
    ;

builtin
    : 'roundUp'
    | 'roundDown'
    | 'abs'
    ;
    

logicalOp
    : 'or'
    | 'and'
    ;

comparisonOp
    : '<'
    | '>'
    | '<='
    | '>='
    | '=='
    | '!='
    ;

literal
    : INT
    | FLOAT
    | STRING
    ;

result
    : '(' STRING (',' STRING)* ')' # ResultWithElements
    | '(' ')'                      # ResultWithoutElements
    ;

// the first char of ID must be a letter
ID
    : ID_LETTER (ID_LETTER | DIGIT)*
    ;

METRIC
    : '$' ID // match $speed
    ;

INT
    : DIGIT+
    ;

FLOAT
    : DIGIT+ '.' DIGIT* // match 1. 39. 3.14159 etc...
    | '.' DIGIT+        // match .1 .14159
    ;

STRING
    : '"' (ESC|.)*? '"'
    ;

fragment
ID_LETTER
    : 'a'..'z'|'A'..'Z'|'_'  // [a-zA-Z_]
    ;

fragment
ESC
    : '\\"' | '\\\\'
    ;

fragment
DIGIT
    : [0-9]  // match single digit
    ;

CHAR
    : [a-z|A-Z]
    ;

LINE_COMMENT
    : '//' .*? '\r'? '\n' -> skip
    ;

COMMENT
    : '/*' .*? '*/' -> skip
    ;

WS 
    : [ \t\r\n]+ -> skip
    ;
