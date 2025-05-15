# The Golang Interpreter

The book is divided into 3 parts. They are Lexical Analysis ,Parsing and Evaluation. There's 4'th extra part dedicated on how to apply optimization technique to make the interpreter faster .
Briefly describing each sections of the program. Any additions to the parts below is greatly entertained and encouraged.

## 1)Lexical Analysis

##### Code is mostly built around tokenization and segregation of string literals into different types .

This part of the code focuses on breaking down the raw string into chucks and categorizing it based on it's attribute. Lets' take the example of a sample C code. "int c =10;". Here the interpreter(compiler in this case) divides the raw string literal into several chunks and gives it separate category. "int" to KEYWORD, "c" is an IDENTIFIER , "=" as in EQUAL sign , 10 as an INTEGER and finally ";" which terminates the statement.

## Key things in Lexical Analysis

### The Lexical Structure

The structure is built around four main components . Let's give them names for us to understand it better. They are input<string> , position<int>, readPosition<int> and ch<byte>.

a) The string is sum of characters between two line breakers. Semicolon in some cases and "\n" in some.
b)position is the index of string being analyzed
c) Position+1
d) ch is the character being analyzed


