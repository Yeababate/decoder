# art-decoder

A command line tool to encode and decode text-based art using a compact run-length bracket notation.

---

## Usage

```
go run . mode=<encode|decode> 'input string'
go run . mode=<encode|decode> -f filepath
```

`decode` is the default, so `mode` can be omitted when decoding.

Input can be provided either as an inline string or as a path to a file.

---

## Modes

### Decode

Expands bracket notation into plain text:

```
$ go run . "[5 #][5 -_]-[5 #]"
#####-_-_-_-_-_-#####

$ go run . "ABC[10 D]EFG"
ABCDDDDDDDDDDEFG
```

### Encode

Compresses repeated characters into bracket notation:

```
$ go run . mode=encode "#####-_-_-_-_-_-#####"
[5 #][5 -_]-[5 #]
```

The encoder checks both single-character and 2-character patterns, and picks whichever gives better compression. Sequences that don't repeat are left as-is.

### File input

Both modes accept a file path instead of an inline string:

```
$ go run . mode=encode filepath
$ go run . mode=decode filepath
```

The file is read and each line is processed independently, the same way as multi-line string input.

### Multi-line

Both modes handle multi-line input automatically. Embed newlines in your quoted string and each line is processed independently:

Input:

    go run . "[3 @]
    [5 -]
    [3 @]"

Output:

    @@@
    -----
    @@@

---
 
## Bracket format

A bracket expression has two arguments separated by a single space:

    [count unit]

- `count` — positive integer, how many times to repeat
- `unit` — one or more characters to repeat

| Encoded | Decoded |
|---|---|
| `[5 #]` | `#####` |
| `[4 -_]` | `-_-_-_-_` |
| `[10 D]` | `DDDDDDDDDD` |

Square brackets are reserved and cannot appear in output.

---

## Errors

| Cause | Message |
|---|---|
| Unbalanced `[` or `]` | `Close the bracket` |
| First token inside `[...]` is not a digit | `The first arg should be number` |
| No space between count and pattern | `There should be space between the arguments` |
| Count fails integer conversion | `Error converting to number` |
| Invalid `mode` value | `Invalid mode: use encode or decode` |
| No input provided | prints usage, exits |

---

## Project structure

```
.
├── main.go
├── decoder.go
├── encoder.go
└── README.md
```

Requires [Go](https://go.dev/dl/) to be installed.