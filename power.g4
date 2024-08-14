grammar expr;

// "1 + 2 * 3" 会被识别为 expression = "1", term = "2 * 3"
// 这里其实用到了一些 BNF 的小技巧
// BNF的递归是深度优先的递归，所以将优先级最高的放在层级最深的文法中可以保证优先级；
expression : term | expression '+' term | expression '-' term;
// term 只能是 factor，或者通过自身递归的乘/除 factor，这是为了区分不同运算符的优先级
// 通过将乘/除放到 term 中，而加法和减法放在 expression 中可以保证运算符的优先级的正确性
term : factor | term '*' factor | term '/' factor;

// 注意 factor，这里表示的是一个因子
// a + b + -c
// 对应的是三个因子 ['a', 'b', '-c']
// 而 a + (b + -c)
// 对应的是两个因子 ['a', 'b + -c']，其中 b + -c 可以基于 expression 递归的去解析
factor : power | factor '!' | factor '^' power;

digit : '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9';
integer : digit | digit integer;
primary : integer | '(' expression ')';
power : primary | primary '^' power;
