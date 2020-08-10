grammar CalcEngin;
// Rules
start : expression EOF;

expression
	:expression op=(AND|OR) expression	 #BoolExp
	|expression op=IN expression				# InStat
	|expression op=(GT|GE|EQ|NE|LT|LE) expression	#CompareStat
	|expression op=(MUL|DIV) expression # MulDiv
	|expression op=(ADD|SUB) expression # AddSub
	|NUMBER                     # Number
	|STRING						# StrStat
	|BOOL						# BoolStat
	|IP  						# IpStat
	|CIDR						# CidrStat
	|LIST						# ListStat
	;


AND :'&&';
OR:'||';
GT:'>';
GE:'>=';
LT:'<';
LE:'<=';

EQ:'=';
NE:'!=';

IN:'in';

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
BOOL:TRUE|FALSE;
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

ID:[a-zA-Z][a-zA-Z0-9]*;


STRING :  '"' (ESC | ~["\\])* '"' ;
fragment ESC :   '\\' (["\\/bfnrt] | UNICODE) ;
fragment HEX : [0-9a-fA-F] ;
fragment UNICODE : 'u' HEX HEX HEX HEX ;


TRUE:'true';
FALSE:'false';

CIDR:SET DOT SET DOT SET DOT SET SP ZNUM 
	;
IP
   : SET DOT SET DOT SET DOT SET
   ;
SET
   :'2''5'[0-5]|'2'[0-4][0-9]|[0-1]?[0-9]?[0-9]
   ;
fragment DOT
   : '.'
   ;

fragment SP
   : '/'
   ;
 
fragment ZNUM
   : [0-2][0-9]|'3'[0-1]|[0-9]
   ;
   
LIST:
	LK PARAM RK
	;
 
LK:
	'['
	;
 
PARAM:
	P1|P2
	;
 
P1	:
	'\'' [a-zA-Z][0-9a-zA-Z]* '\'' (',' '\''[a-zA-Z][0-9a-zA-Z]* '\'')+
	;
 
P2:
	[0-9]+(','[0-9]+)+ 
	;
	
 
RK:
	']'
	;








