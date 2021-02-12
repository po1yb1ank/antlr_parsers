grammar lightc;

Space : ' ';
For : 'for';
If : 'if';
Bool : 'bool';
Int : 'int';
Return : 'return';
Void : 'void';
Main : 'main';
LeftParen : '(';
RightParen : ')';
LeftBrace : '{';
RightBrace : '}';
Less : '<';
Greater : '>';
Question : '?';
Colon : ':';
Semi : ';';
Comma : ',';
Assign : '=';
Equal : '==';
NotEqual : '!=';
WS  :   (' '|'\r'|'\n'|'\t') -> channel(HIDDEN);

program : typedef ' ' 'main' '(' ')' '{' statement '}' ;
typedef
	: 'int'
	| 'bool'
	| 'void'
	;
statement
	: declaration ';' statement
	| '{' statement '}' statement
	| forloop statement
	| ifstat statement
	| returnval
	;
declaration : typedef ' ' Identifier assign ;
assign	: '=' assign_end ;
assign_end
	: Identifier
	| Number
	;
forloop : 'for' '{' declaration ';' bool_expression ';' '}' ;
bool_expression 
	: Identifier relop Identifier
	| Number relop Identifier
	;
relop 
	: '<'
	| '>' 
	| '==' 
	| '!='
	;
ifstat : 'if' '(' bool_expression ')' ;
returnval : 'return' ' ' Number ';' ;

fragment NONDIGIT : [a-zA-Z_];
fragment DIGIT : [0-9];
Number : DIGIT+ ;
Identifier : NONDIGIT+ ;
