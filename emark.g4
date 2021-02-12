grammar emark;

Block : '<block ' ;
BlockEnd : '</block>' ;
Rows : 'rows=' ;
Columns : 'columns=' ;
Row : '<row>' ;
RowEnd : '</row>' ;
Column : '<column' ;
ColumnEnd : '</column>' ;
Walign : 'walign' ;
Halign : 'halign' ;
BgColor : 'bgcolor' ;
Width : 'width' ;
fragment DIGIT : [0-9];
fragment TXT : [a-zA-Z];
Ctx : TXT ;
Number : DIGIT ;
TextColor : 'textcolor' ;
Spacebar : ' ' ;
CloseTag : '>' ;
OpenTag : '<' ;

WS  :   (' '|'\r'|'\n'|'\t') -> channel(HIDDEN);

program : '<block ' Rows Number+ ' ' Columns Number+ '>' tag* '</block>' EOF;
tag
	: ctx* rowop
	| ctx* blop
	| ctx* colop
	| ctx+
	;
rowop : Row tag* RowEnd;
blop : Block Rows Number+ ' ' Columns Number+ '>' tag* BlockEnd;
colop : Column tag* ColumnEnd;
ctx 
	: ' '
	| Number
	| Ctx
	;
