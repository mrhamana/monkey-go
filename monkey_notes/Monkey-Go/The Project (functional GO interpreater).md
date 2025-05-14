
The book is divided into  3 parts. They are Lexical Analysis ,Parsing and Evaluation. There's 4'th extra part dedicated on how to apply optimization technique to make the interpreter faster .
Briefly describing each sections of the program. Any additions to the parts below is greatly entertained and encouraged.

## 1)Lexical Analysis
##### Code is mostly built around tokenization and segregation of string literals into different types .

This part of the code focuses on breaking down the raw string into chucks and categorizing it based on it's attribute. Lets' take the example of a sample C code. "int c =10;". Here the interpreter(compiler in this case) divides the raw string literal into several chunks and gives it separate category. "int" to KEYWORD, "c" is an IDENTIFIER , "=" as in EQUAL sign , 10 as an INTEGER and finally ";" which terminates the statement.


### Key things in Lexical Analysis
#### 1) The Lexer structure 
The structure is built around four main components . Let's give them names for us to understand it better. They are input<string> , position<int>, readPosition<int> and ch<byte>. Each of them are important and plays important roles to analyze the code as it was intended. the input field is literally a line of code. To me more concrete , it's the string starting from one statement breaker to the another. Another is position which is the index of the next character from the character being analyzed. This is important in some aspects as some of the functionality requires it . Common example is the double equal sign "==" which is the equality operator rather than the assignment operator. Some languages even analyzes two more characters from the one currently being analyzed (JavaScript for example) .Next is the readPosition which is the index of the character being read at the moment. Finally ch is the current character being read. 