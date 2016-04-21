# Colonoscopy

So it seems that semi-colons and braces (`;`, `{`, and `}`) are under threat. 
New languages are shunning them and their use seems to be dwindling. So what's 
happening to all those unused characters? We've put them to work, that's what. 
Introducing Colonoscopy, a 
[Rho&#8242;&#8242;](https://en.wikipedia.org/wiki/P%E2%80%B2%E2%80%B2), among
[others](https://en.wikipedia.org/wiki/Brainfuck) [NSFW],
influenced language designed purely to use up characters that are no longer seen
as _useful_ in trendy, modern languages.

## Syntax

Colonoscopy is a turing complete language with 8 operators. These operators 
translate to _BF_ and `C` as follows:

| Symbol | BF  | C                 | Description
| ------ | --- |------------------ | -----------
| `;}`   | `+` | `++ptr;`          | Move the pointer to the right
| `;{`   | `-` | `--ptr;`          | Move the pointer to the left
| `;;}`  | `>` | `++*ptr;`         | Increment the memory cell under the pointer
| `;;{`  | `<` | `--*ptr;`         | Decrement the memory cell under the pointer
| `;;;}` | `.` | `putchar(*ptr)`   | Output the character signified by the cell at the pointer
| `;;;{` | `,` | `*ptr=getchar();` | Input a character and store it in the cell at the pointer
| `{{`   | `[` | `while (*ptr) {`  | Jump past the matching `}}` if the cell under the pointer is 0
| `}}`   | `]` | `}`               | Jump back to the matching `{{`

In addition, each command **must** be terminated with a semi-colon. Other than 
`EOF`, no other characters are allowed in a colonoscopy program. Those are 
reserved for other languages that think that `;`, `{` and `}` are too good for
them.

## Interpreter 

An interpreter, written in Go (because irony) is available in this package to
allow you to play with Colonoscopy scripts. 

```
go get github.com/domdavis/colonoscopy
go install github.com/domdavis/colonoscopy
colonoscopy <sourcefile>
```

> Convention is that Colonoscopy files use a `.cl` suffix.

## Hello World

The following is an implementation of "Hello World!" written in Colonoscopy:

```
;;};;;};;;};;;};;;};;;};;;};;;};{{;;};;;};;;};;;};;;};{{;;};;;};;;};;};;;};;;};;;};;};;;};;;};;;};;};;;};;{;;{;;{;;{;;;{;}};;};;;};;};;;};;};;;{;;};;};;;};{{;;{;}};;{;;;{;}};;};;};;;;};;};;;{;;;{;;;{;;;;};;;};;;};;;};;;};;;};;;};;;};;;;};;;;};;;};;;};;;};;;;};;};;};;;;};;{;;;{;;;;};;{;;;;};;;};;;};;;};;;;};;;{;;;{;;;{;;;{;;;{;;;{;;;;};;;{;;;{;;;{;;;{;;;{;;;{;;;{;;;{;;;;};;};;};;;};;;;};;};;;};;;};;;;};
```

```
$ colonoscopy example.cl 
Tokenising...done.
Hello World!
Execution complete!
```

## FAQ

### Why?

Why not? Also, it helps with learning Go. Also 
[this](https://twitter.com/idomdavis/status/722050755616776192).

### Can I do `x` in Colonoscopy

It's turing complete, so yes.

### Do I want to do `x` in Colonoscopy?

Probably not.

### Should this be considered the reference implementation of Colonoscopy?

Doubtful. I'm not sure the code to read characters from `stdin` is cross
platform, the tokeniser is likely very naive and inefficient, plus I can't be 
sure it's bug free as I used SDD (Satire Driven Programming) and there is only
one manual test. ~~Hell, I don't even know if the install instructions work 
properly ;)~~ - Appears they do, at least on the Docker image I tried :D

## TODO

Converter from Ook!, and other such languages. Maybe a proper compiler. 
Who knows.

## License

Public Domain
