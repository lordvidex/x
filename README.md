# x
Highly opinionated lib that contains utility functions.
## Content
- **req/**: HTTP request utility
- **res/**: HTTP response utility
- **ptr/**: Conversion of object from and to pointer values (one-liners without nil checks)
- **auth/**: Passing tokens across app layers with `context.Context`
## Motivation
Some verbose operations in Golang can distract and take one's mind off the business logic and to solve this, I create util functions to `pkg/` to reduce repetition. This was okay, until I had to work on many HTTP projects.  
  
I have created this package to reduce repetitive copying of my `pkg/` helper functions from project to project.
 
## Inspiration
### Package Naming
- I've named the package `lordvidex/x` because many of the functionalities that are either upgrades or that the Go (Golang) maintainers wish were in the standard library are typically placed in the golang.org/x package.
- Package name `lordvidex/x` is relatively short, thereby easy to remember and import

### API and Design Goals
```go
err := req.I().Will().Bind(r, &v).Validate(v).Err()
```
- Functions should be designed in a way that it feels as natural as speaking English language.
- Functions should reduce boilerplate code required to parse and validate JSON.

## Future Plans
### Extensibility
- [ ] Use interfaces where necessary for `req/`
- [ ] More flexible response struct in `res/`
- [ ] Add more functions that are often repeated.
