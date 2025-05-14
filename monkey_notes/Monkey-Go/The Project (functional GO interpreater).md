
The book is divided into  3 parts. They are Lexical Analysis ,Parsing and Evaluation. There's 4'th extra part dedicated on how to apply optimization technique to make the interpreter faster .
Briefly describing each sections of the program. Any additions to the parts below is greatly entertained and encouraged.

### 1)Lexical Analysis
##### Code is mostly built around tokenization and segregation of string literals into different types .

This part of the code focuses on breaking down the raw string into chucks and categorizing it based on it's attribute. Lets' take the example of a sample C code. "int c =10;". Here the interpreter(compiler in this case) divides the raw string literal into several chunks and gives it separate category. "int" to KEYWORD, "c" is an IDENTIFIER , "=" as in EQUAL sign , 10 as an INTEGER and finally ";" which terminates the statement.


