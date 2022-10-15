grammar Assembly;
options {}

start
: (directive (instruction | synthetic_instruction))* EOF;

// Is there an easy way to enforce size limits on the immediates here?  Or will we need to do that when walking the tree?

instruction
: ADD register_symbol COMMA register_symbol
| ADC register_symbol COMMA register_symbol
| SBB register_symbol COMMA register_symbol
| OR  register_symbol COMMA register_symbol
| AND register_symbol COMMA register_symbol
| XOR register_symbol COMMA register_symbol
| MOV register_symbol COMMA register_symbol
| MOV register_symbol COMMA nibble
| MOV register_combo COMMA R0
| MOV R0 COMMA register_combo
| MOV R0 COMMA L_BRACKET data_byte R_BRACKET // Maybe this should be a special case, like the register combo.
//| MOV PC COMMA // At one point in the docs, this instruction is called LPC.  Personally, I think that's cleaner than all these MOV variants, but whatever.
| LPC data_byte
| JR data_byte
| CP R0 COMMA nibble
| ADD R0 COMMA nibble
| INC register_symbol
| DEC register_symbol
| DSZ register_symbol
| OR R0 COMMA nibble
| AND R0 COMMA nibble
| XOR R0 COMMA nibble
| EXR nibble
| BIT  rg COMMA quarter
| BSET rg COMMA quarter
| BCLR rg COMMA quarter
| BTG  rg COMMA quarter
| RRC register_symbol
| RET R0 COMMA nibble
| mn_SKIP flag COMMA quarter
;

synthetic_instruction
: RLC register_symbol COMMA register_symbol
| SL register_symbol COMMA register_symbol
| LSR register_symbol
| CPL R0
| CPL register_symbol COMMA register_symbol
| NEG register_symbol COMMA register_symbol
| NOP
;

macro:
DOT
( STRING
| 
);

register_symbol
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


data_byte: literal;
nibble: literal;
quarter: literal; // Half a nibble
register_combo: L_BRACKET register_symbol COLON register_symbol R_BRACKET;

literal: (NUMBER | CHARACTER); // TODO support for non base 10 values.

directive
: DOT (
    STRING
|   IDENTIFIER
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
mn_SKIP: 'SKIP';

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

COMMENT
: '/*' .*? '*/' -> skip
;
LINE_COMMENT
: '//' ~[\r\n]* -> skip
;

fragment LETTER      : [A-Za-z_];
IDENTIFIER           : LETTER(LETTER|NUMBER)*;

NUMBER      : '-'? [0-9]+ ;
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