grammar Assembly;
options {}

start
: (directive? (instruction | syntheticInstruction))* EOF;

// Is there an easy way to enforce size limits on the immediates here?  Or will we need to do that when walking the tree?

instruction
: ADD registerSymbol COMMA registerSymbol
| ADC registerSymbol COMMA registerSymbol
| SUB registerSymbol COMMA registerSymbol
| SBB registerSymbol COMMA registerSymbol
| OR  registerSymbol COMMA registerSymbol
| AND registerSymbol COMMA registerSymbol
| XOR registerSymbol COMMA registerSymbol
| MOV registerSymbol COMMA registerSymbol
| MOV registerSymbol COMMA nibble
| MOV registerCombo COMMA R0
| MOV R0 COMMA registerCombo
| MOV R0 COMMA L_BRACKET dataByte R_BRACKET // Maybe this should be a special case, like the register combo.
//| MOV PC COMMA // At one point in the docs, this instruction is called LPC.  Personally, I think that's cleaner than all these MOV variants, but whatever.
| LPC dataByte
| JR dataByte
| CP R0 COMMA nibble
| ADD R0 COMMA nibble
| INC registerSymbol
| DEC registerSymbol
| DSZ registerSymbol
| OR R0 COMMA nibble
| AND R0 COMMA nibble
| XOR R0 COMMA nibble
| EXR nibble
| BIT  rg COMMA quarter
| BSET rg COMMA quarter
| BCLR rg COMMA quarter
| BTG  rg COMMA quarter
| RRC registerSymbol
| RET R0 COMMA nibble
| SKIPI flag COMMA quarter
;

syntheticInstruction
: RLC registerSymbol COMMA registerSymbol
| SL registerSymbol COMMA registerSymbol
| LSR registerSymbol
| CPL R0
| CPL registerSymbol COMMA registerSymbol
| NEG registerSymbol COMMA registerSymbol
| NOP
;

macro:
DOT
( STRING
| 
);

registerSymbol
: R0
| R1
| R2
| R3
| R4
| R5
| R6
| R7
| R8
| R9
| OUT
| IN
| JSR
| PCL
| PCM
| PCH
;

rg
: R0
| R1
| R2
// | RS // FIXME: What the heck is RS?
;

flag // TODO: not yet finished.
: 'z'
| 'nz'
| 'c'
| 'nc'
;

// Registers
R0: 'R0';
R1: 'R1';
R2: 'R2';
R3: 'R3';
R4: 'R4';
R5: 'R5';
R6: 'R6';
R7: 'R7';
R8: 'R8';
R9: 'R9';
OUT: 'OUT';
IN: 'IN';
JSR: 'JSR';
PCL: 'PCL';
PCM: 'PCM';
PCH: 'PCH';


dataByte: literal;
nibble: literal;
quarter: literal; // Half a nibble
registerCombo: L_BRACKET registerSymbol COLON registerSymbol R_BRACKET;

literal: (NUMBER | CHARACTER); // TODO support for non base 10 values.

directive
: DOT (
    STRING
|   LABEL
);

// Instructions
ADD: 'ADD';
ADC: 'ADC';
SUB: 'SUB';
SBB: 'SBB';
OR: 'OR';
AND: 'AND';
XOR: 'XOR';
MOV: 'MOV';
JR: 'JR';
CP: 'CP';
INC: 'INC';
DEC: 'DEC';
DSZ: 'DSZ';
EXR: 'EXR';
BIT: 'BIT';
BSET: 'BSET';
BCLR: 'BCLR';
BTG: 'BTG';
RRC: 'RRC';
RET: 'RET';
SKIPI: 'SKIP';

// Synthetic Instructions
RLC: 'RLC';
SL: 'SL';
LSR: 'LSR';
CPL: 'CPL';
NEG: 'NEG';
NOP: 'NOP';

LPC: 'LPC'; // This is listed in the instruction sheet for MOV PC... maybe it was a typo, but I like it.

L_BRACKET: '[';
R_BRACKET: ']';

// Directives
STRING: 'string';

DOT: '.';
COMMA: ',';
COLON: ':';

// STRING_LITERAL: '"' . '"' // TODO

CHARACTER
: '\'' . '\''
| '\'\\' . '\'';


// BLOCK_COMMENT
// : '/*' .*? '*/' -> skip
// ;
LINE_COMMENT
: ';' ~[\r\n]* -> skip
;

LABEL                : LETTER(LETTER|NUMBER)*;
fragment LETTER      : [A-Za-z_];
NUMBER               : '-'? [0-9]+ ;
WHITESPACE           : ' ' -> skip ;
NEWLINE              : '\n' -> skip ;
CARRIAGE             : '\r' -> skip ;
TAB                  : '\t' -> skip ;

UNKNOWN : .;

/* 

Various TODOs

Macro support
Data section and program section...?
Labels
Literals
Assembler directives (stringz, char, int, word, etc.)
A list of potential directives to add:

https://docs.oracle.com/cd/E26502_01/html/E28388/eoiyg.html

*/