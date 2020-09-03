grammar Prosr1;

// token
REPEATED : 'repeated';
TYPE : 'string' | 'bool' | 'int32';
IDENT : [a-zA-Z]([a-zA-Z0-9_])+;
NUMBER : [0-9]+;
WHITESPACE : [ \r\n\t]+ -> skip;

// parser rules
content : syntax definition* EOF;
syntax : 'syntax' '=' quote 'prosr1' quote';';
definition : (packageDefinition | hubDefinition | messageDefinition);

packageDefinition : 'package' fullIdent ';';

hubDefinition : 'hub' IDENT '{' (actionDefinition | callsDefinition)+ '}';
actionDefinition : 'action' IDENT'('(fullIdent)?')' (callsDefinition)? ';';
callsDefinition : 'calls' IDENT'('fullIdent')' 'on' target=('caller' | 'all') ';';

messageDefinition : 'message' IDENT '{' (fieldDefinition)* '}';
fieldDefinition : (REPEATED)? typeIdent fullIdent '=' NUMBER';';

// literals & identifier
quote : '\'' | '"';

fullIdent : IDENT ('.' IDENT)*;
typeIdent : fullIdent | TYPE;