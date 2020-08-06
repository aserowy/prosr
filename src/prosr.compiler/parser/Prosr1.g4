grammar Prosr1;

// token
SYNTAX : 'syntax';
MESSAGE : 'message';
HUB : 'hub';
ACTION : 'action';
RETURNS : 'returns';
TO : 'to';
TYPE : 'string' | 'bool' | 'int32';
IDENT : [a-zA-Z]([a-zA-Z0-9] | '_')+;
NUMBER : [0-9]+;
WHITESPACE : [ \r\n\t]+ -> skip;

// parser rules
content : syntax definition* EOF;
syntax : SYNTAX '=' quote 'prosr1' quote';';
definition : (hub | message);

hub : HUB hubIdent '{' (sending | returning)+ '}';
sending : ACTION sendingIdent'('inputType=messageIdent')' (RETURNS '('outputType=messageIdent')' TO sendingTarget) ';';
sendingTarget : ('caller' | 'all');
returning : RETURNS '('messageIdent')' TO returningTarget ';';
returningTarget : 'all';

message : MESSAGE messageIdent '{' (field)* '}';
field : typeIdent fieldIdent '=' NUMBER';';
typeIdent : (messageIdent | TYPE);

// literals & identifier
quote : '\'' | '"';

hubIdent : IDENT;
sendingIdent : IDENT;
messageIdent : IDENT;
fieldIdent : IDENT;