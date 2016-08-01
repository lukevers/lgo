# lgo - Literate Go

A CLI to build Literate Go (lgo) files.

## What's Literate Go?

Literate Go is the [Go](https://golang.org/) programming language, but with a [literate programming](https://en.wikipedia.org/wiki/Literate_programming) twist.

This specific tool assumes that your lgo files are basically [Markdown](https://en.wikipedia.org/wiki/Markdown) files with code in triple back-tick blocks. You may include the language name in the triple back-tick blocks if you wish, but it's not required at this time.

## Why?

Because why not.

I've been interested in literate programming for some time now, but I've never had to chance to do any of it. I've also never written a lexer/parser before, so this is also a fun way to do that.

## Credits

Thanks to [benbjohnson/sql-parser](https://github.com/benbjohnson/sql-parser) for giving me a base parser/scanner to work with.
