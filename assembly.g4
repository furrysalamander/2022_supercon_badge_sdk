grammar assembly;
options {}

start
: (label? instruction)* EOF;

// Is there an easy way to enforce size limits on the immediates here?  Or will we need to do that when walking the tree?

instruction
//: unary_mnemonic operand
//| binary_mnemonic operand operand
: ADD register COMMA register
| ADC register COMMA register
| SBB register COMMA register
| OR  register COMMA register
| AND register COMMA register
| XOR register COMMA register
| MOV register COMMA register
| MOV register COMMA nibble
| MOV register_combo COMMA R0
| MOV R0 COMMA register_combo
| MOV R0 COMMA '[' byte ']' // Maybe this should be a special case, like the register combo.
//| MOV PC COMMA // At one point in the docs, this instruction is called LPC.  Personally, I think that's cleaner than all these MOV variants, but whatever.
| LPC byte
| JR byte
| CP R0 COMMA nibble
| ADD R0 COMMA nibble
| INC register
| DEC register
| DSZ register
| OR R0 COMMA nibble
| AND R0 COMMA nibble
| XOR R0 COMMA nibble
| EXR nibble
| BIT  RG COMMA quarter
| BSET RG COMMA quarter
| BCLR RG COMMA quarter
| BTG  RG COMMA quarter
| RRC register
| RET R0 COMMA nibble
| mn_SKIP flag COMMA quarter
;

register
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

flag // TODO: not yet finished.
: 'z'
| 'nz'
| 'c'
| 'nc'
;


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

RG
: R0
| R1
| R2
| RS // FIXME: What the heck is RS?
;

// TODO
byte: '';
nibble: '';
quarter: ''; // Half a nibble

register_combo: '[' register ':' register ']';

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

LPC: 'LPC'; // This is listed in the instruction sheet for MOV PC... maybe it was a typo, but I like it.

// operand
// : 'R1';
// binary_mnemonic
// : 'ADD'
// | 'ADC'
// | 'SUB'
// | 'SBB'
// | 'OR'
// | 'AND'
// | 'XOR'
// | 'MOV' // Should MOV be a special case?  How do we want to handle these kinds of things?
// ;

// immediate_mnemonic
// : 'JR'
// ;
// unary_mnemonic
// : 'INC';

COMMA: ',';

WHITESPACE : ' ' -> skip;