
## The Function Signature

Signature of the function refers to how the function is called in a programming language. This includes, 
1) Function Name
2) Parameters and it's types (some languages do not require types of their parameter)
3) Return Type (not mandatory for  some languages)

Note: In Golang the signature of the function requires all of these things . Here's an example.
 `func Myfunc(int a, int b) string`. This was important to define because of the upcoming definition.

## Contracts 

Refers to the formal agreement between different parts of the program - usually between function , class or interface. It a blueprint for any class which should implement the methods declared in the interface. In other words, if an interface is inherited , then the methods should be declared . If not , the compiler gives an error.

#### What about variables inside an interface?
While mutable variables are not allowed in it, final and static variables can be declared . 
Static variables persists it's value and increment and decrement of the variable will will change it's value. The variable is final meaning it should be given a value to be declared.

**Note**: Interfaces doesn't have it's own memory in the heap.

## The Bingo Question (What's the need of this thing ?)

It gives us the general system or framework to work on. Let's say there's a company which gives Laptops to it's employees. Some other company may give Desktop to it's employees . Now the interesting thing is , both of these things fall under the class Computer. And both of them are provided to the employees to code on. So what we can do is , we can create an abstract class Computer with abstract method code() and we can inherit the class to Laptop and Desktop class. The interesting thing is , an abstract class with abstract method is called the ALMIGHTY interface.

**To sum up**, interfaces contains signature of abstract functions which is used to create an abstract layer /generalize a class .
