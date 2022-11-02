# compiler60 service

This is the source code for the "compiler60" service from [FAUST CTF 2022](https://2022.faustctf.net).

**As it was written for a CTF service, the code is deliberately insecure and contains exploitable bugs. It
is provided for educational purposes only, do not even think about (re-) using it for anything productive!**

The code is released under the ISC License, see LICENSE.txt for details.

----

Remember the days when you could invent new features for programming languages like recursion or BNF? To relive those
distant times we took ALGOL60 and spiced it up a bit. ALGOL60v2 now has safe strings, 64-bit, UTF-8...  
Make sure to read the specification ([./compiler/ALGOL60v2-spec.md](./compiler/ALGOL60v2-spec.md)) to find out all the
differences to the original language.

## Getting started

Open [gui.html](./gui.html) or http://[fd66:666:<team-num>::2]:6061 in your favourite browser.

## Structure of the service

- compiler (port 6061): compiles source code to elf files and signs them, written in Java
    - contains multiple vulnerabilities so if you found one keep looking
- executor (port 6062): given an elf file verifies the signature and executes it in a sandbox,
    - no intended vulnerabilities in there, if you find something tell us, reversing this will be a waste of time
    - If you want to debug the execution of a binary in the sandbox see the [docker-compose.yml](docker-compose.yml).

## Flag IDs

The gameserver will publish flag IDs for stored flag. The flag ID will be the filename of the file which contains the
flag.

## Contracts

- The signature for the ELF files MUST NOT include the `.rodata` section. You can therefore modify the `.rodata` section
  and still execute it.
- The compiler MUST be deterministic. Compiling the same source code even with changed int constants (except -1, 0, 1)
  must result in the same ELF file (except the `.rodata` section).
