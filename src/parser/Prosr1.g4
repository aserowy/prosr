grammar Prosr1;

// token
REPEATED : 'repeated';
KEYTYPE : 'int32' | 'string';
TYPE : 'bool';
IDENT : [a-zA-Z]([a-zA-Z0-9_])+;
NUMBER : [0-9]+;
WHITESPACE : [ \r\n\t]+ -> skip;

// parser rules
content : syntaxDefinition definitions+=bodyDefinition* EOF;
syntaxDefinition : 'syntax' '=' '"' 'prosr1' '"' ';';
bodyDefinition : packageDefinition | hubDefinition | messageDefinition;

packageDefinition : 'package' fullIdent ';';

hubDefinition : 'hub' IDENT '{' (actionDefinition | callsDefinition)+ '}';
actionDefinition : 'action' IDENT'('fullIdent?')' (callsDefinition | ';');
callsDefinition : 'calls' IDENT'('fullIdent?')' 'on' target=('caller' | 'all') ';';

messageDefinition : 'message' IDENT '{' (fieldDefinition | mapDefinition)* '}';
fieldDefinition : REPEATED? typeIdent IDENT '=' NUMBER';';
mapDefinition : 'map' '<' KEYTYPE ',' typeIdent '>' IDENT '=' NUMBER';';

// literals & identifier
fullIdent : IDENT ('.' IDENT)*;
typeIdent : fullIdent | KEYTYPE | TYPE;