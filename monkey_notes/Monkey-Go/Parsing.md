# Parsing the Code

Parser is a program used to process tokens made by the tokenizer and is used to make syntax tree (AST which stands for abstract syntax tree).
There are mainly two ways to parse code . Them being top-down approach and bottom-up approach.  And they have sub types of their own. The one we're implementing is called "Recursive Descent parser". 
Some details regarding this parser is mentioned below.
1) Founded by Vaughan Pratt that's why sometimes it is referred as "Pratt parser".
2) It's a top-bottom approach meaning we're analyzing the code from top and ending it at the bottom until we encounter EOF(End of file).
3) The top bottom parsers starts with constructing root node of the AST and then descends while the latter does it the other way around.
4) We'll be implementing Recursive Descent Parser

### The infamous Let
The let statement in out code looks like this . "let <IDENTIFIER>=<EXPRESSION>"Before we move on , there's an important difference between statements and expressions. Expressions outputs something . For example, 2+3 returns 5 so it's an expression. Whereas statements don't . Using let doesn't return anything thus it's a statement. Again , to make it clear ,

 **Statements don't return values whereas expressions do**. 