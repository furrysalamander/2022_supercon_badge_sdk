# 2022_supercon_badge_sdk
An assembler for the 2022 Hackaday Superconference Badge

The code can be compiled by running

`go build`

If you need to update the lexer/parser code, you can run

`go generate ./...`

Run the resulting binary with `./assembler your_assembly_file.asm` and your program will be saved as `out.bin`