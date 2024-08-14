grammar Expression;

// 解析入口
expr:   expr op=('*'|'/') expr   # MulDivExpr
    |   expr op=('+'|'-') expr   # AddSubExpr
    |   '-' expr                 # NegateExpr
    |   expr '!'                 # FactorialExpr
    |   INT                      # IntExpr
    |   '(' expr ')'             # ParenExpr
    ;

// 词法规则
INT:    [0-9]+ ;
WS:     [ \t\r\n]+ -> skip ;
