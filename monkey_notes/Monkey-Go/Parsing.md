# Parsing the Code

Parser is a program used to process tokens made by the tokenizer and is used to make syntax tree (AST which stands for abstract syntax tree).
There are mainly two ways to parse code . Them being top-down approach and bottom-up approach.  And they have sub types of their own. The one we're implementing is called "Recursive Descent parser". 
Some details regarding this parser is mentioned below.
1) Founded by Vaughan Pratt that's why sometimes it is referred as "Pratt parser".
2) It's a top-bottom approach meaning we're analyzing the code from top and ending it at the bottom until we encounter EOF(End of file).
3) The top bottom parsers starts with constructing root node of the AST and then descends while the latter does it the other way around.
4) We'll be implementing Recursive Descent Parser

### The Infamous Let
The let statement in out code looks like this . "let **IDENTIFIER** = **EXPRESSION** "Before we move on , there's an important difference between statements and expressions. Expressions outputs something . For example, 2+3 returns 5 so it's an expression. Whereas statements don't . Using let doesn't return anything thus it's a statement. Again , to make it clear , __Statements don't return values whereas expressions do__.

## A Short Note for Interfaces in Golang

This took me soo long to understand just because it's not a common functionality and I am not familiar with the concept. An interface is a group of methods which is grouped to perform a specific task and they work together to achieve a certain goal. For example , i want to build a car . Then i make an interface **BuildCar** and inside that I'll add methods like __putwheel__ , __addengine__ , __assemblecar__ ,etc. All we're doing is grouping like method signature inside the interface .This helps to make the code modular and easy to write . Rather than write every functions differently , make an interface which groups like methods. Golang takes this as a datatype soo if i want to add an interface , i can just do 

`func buildMeACar(maker MakeCar) {`
`maker.makeEngine() 
`maker.putTire()
`maker.assembleCar()}`




