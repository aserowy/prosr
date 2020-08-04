grammar Prosr1;

// token
SYNTAX : 'syntax';
MESSAGE : 'message';
HUB : 'hub';
ACTION : 'action';
ACTIONTARGET: ('caller' | 'all');
RETURNS : 'returns';
RETURNSTARGET: 'all';
TO : 'to';
TYPE : 'string' | 'bool' | 'int32';
IDENT : [a-zA-Z]([a-zA-Z0-9] | '_')+;
NUMBER : [0-9]+;
WHITESPACE : [ \r\n\t]+ -> skip;

// parser rules
main : syntax definition EOF;
syntax : SYNTAX '=' quote 'prosr1' quote';';
definition : (hub | message)*;

hub : HUB hubIdent '{' (sending | returning)+ '}';
sending : ACTION actionIdent'('messageIdent')' (RETURNS '('messageIdent')' TO ACTIONTARGET';')';';
returning : RETURNS '('messageIdent')' TO RETURNSTARGET';';

message : MESSAGE messageIdent '{' (field)* '}';
field : (messageIdent | TYPE) fieldIdent '=' NUMBER';';

// literals & identifier
quote : '\'' | '"';

hubIdent : IDENT;
actionIdent : IDENT;
messageIdent : IDENT;
fieldIdent : IDENT;