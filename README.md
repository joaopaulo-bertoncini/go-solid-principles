# Go Solid Principles

## What is SOLID?

These principles establish practices that lend to developing software with considerations for maintaining and extending as the project grows. Adopting these practices can also contribute to avoiding code smells, refactoring code, and Agile or Adaptive software development.

- Single responsability:
“A module should have one, and only one reason to change”

- Open-Close principle:
“A software artifact should be open for extension but closed for modifications”

- Liskov substitution principle
“What is wanted here is something like the following substitution property: If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T”

- Interface segregation
“Clients should not be forced to depend on methods they don't use”

- Dependency inversion
“High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should not depend on abstractions”