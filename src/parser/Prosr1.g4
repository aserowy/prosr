grammar Prosr1;

// token
SYNTAX : 'syntax';
MESSAGE : 'message';
HUB : 'hub';
ACTION : 'action';
CALLS : 'calls';
ON : 'on';
TYPE : 'string' | 'bool' | 'int32';
IDENT : [a-zA-Z]([a-zA-Z0-9] | '_')+;
NUMBER : [0-9]+;
WHITESPACE : [ \r\n\t]+ -> skip;

// parser rules
content : syntax definition* EOF;
syntax : SYNTAX '=' quote 'prosr1' quote';';
definition : (hub | message);

hub : HUB hubIdent '{' (sending | returning)+ '}';
sending : ACTION sendingIdent'('sendingMessageIdent')' (CALLS returningIdent'('returningMessageIdent')' ON sendingTarget)? ';';
sendingMessageIdent : (messageIdent)?;
sendingTarget : ('caller' | 'all');
returning : CALLS returningIdent'('returningMessageIdent')' ON returningTarget ';';
returningMessageIdent : (messageIdent)?;
returningTarget : 'all';

message : MESSAGE messageIdent '{' (field)* '}';
field : typeIdent fieldIdent '=' NUMBER';';
typeIdent : (messageIdent | TYPE);

// literals & identifier
quote : '\'' | '"';

hubIdent : IDENT;
fieldIdent : IDENT;
messageIdent : IDENT;
returningIdent : IDENT;
sendingIdent : IDENT;