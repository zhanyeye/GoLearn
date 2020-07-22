### Wrapping up
#### Some of Go's syntax around
+ Writing tests
+ Declaring functions, with arguments and return types
+ `if`, `const` and `switch`
+ Declaring variables and constants
#### The TDD process and why the steps are important
+ Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
+ Writing the smallest amount of code to make it pass so we know we have working software
+ Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with

In our case we've gone from `Hello()` to `Hello("name")`, to `Hello("name", "French")` in small, easy to understand steps.

This is of course trivial compared to "real world" software but the principles still stand. TDD is a skill that needs practice to develop but by being able to break problems down into smaller components that you can test you will have a much easier time writing software.