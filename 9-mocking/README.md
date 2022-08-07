# Mocking

This continues on the last section, dependency injection. We've demonstrated how to create a spy.

# Spy

A spy is a stub that records some information on how they were called.

# Importance of Mocking

1. Useful for testing important areas of the code. It's useful to get the system in a particular state. Without mocking, it's difficult.
2. Without mocks, you may have to setup databases, and other third parties things just to test simple business rules.
3. By having to spin up a database or webservice to test something you're likely to have fragile tests due to the unreliabilty of such services.
4. Be mindful of the value of your tests - don't overdo it, and think about the impact they would have in future of refactoring.

# Test Doubles

A test double is a generic term for any case where you replace a production object for testing purposes.
