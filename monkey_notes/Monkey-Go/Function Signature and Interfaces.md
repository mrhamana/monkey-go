
## Function Signature

A **function signature** defines how a function is called in a programming language. It typically includes:

- **Function Name**
- **Parameters and their types** (Note: Some languages do not require explicit types.)
- **Return Type** (Optional in some languages)

Note: In Golang the signature of the function requires all of these things . Here's an example.
 `func Myfunc(int a, int b) string`. This was important to define because of the upcoming definition.

## Contracts 

A **contract** in programming refers to a formal agreement between different components of a program — typically between a function, class, or interface.

In the context of interfaces:

- An interface defines a **blueprint** — a set of method signatures that any implementing class must fulfill.
- If a class claims to implement an interface but fails to define any of its methods, the compiler will throw an error.

#### What about variables inside an interface?
In most object-oriented languages:

- **Mutable variables are not allowed** inside interfaces.
- However, **`final`** and **`static`** variables **can** be declared.

Key points:

- **Static variables** are shared across all implementations and persist their values.
- **Final variables** must be assigned a value at declaration and cannot be modified later.
- Interfaces **do not have their own memory on the heap** — they serve only as templates.

## The Bingo Question (What's the need of this thing ?)

Interfaces help create a **generalized framework** for building software. They allow us to define **what** a class should do, not **how** it should do it.
#### Real-World Analogy:

Imagine a company gives:

- Laptops to some employees
- Desktops to others

Both laptops and desktops fall under the general category of **Computers**. Each can be used to perform tasks like coding.

We can define a general abstract class or interface like this:
`interface Computer {
    void code();
}`
 Then implement it like this :
`class Laptop implements Computer {
    public void code() {
    }
}`

`class Desktop implements Computer {
    public void code() {
    }
}`

This way, your program can work with the general type `Computer` and not worry about whether it’s a `Laptop` or `Desktop`.

### Summary

- **Function signatures** define how functions are called — especially important in strongly typed languages like Go.
- **Interfaces** provide a contract for classes to follow and enable abstraction and loose coupling.
- They help build flexible, scalable, and maintainable systems by enforcing consistency without dictating implementation.