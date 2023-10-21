# x
Highly opinionated lib that contains utility functions.
## Content
- **req/**: HTTP request utility
- **res/**: HTTP response utility
- **ptr/**: Conversion of object from and to pointer values (one-liners without nil checks)
- **auth/**: Passing tokens across app layers with `context.Context`
## Motivation
Some actions are repeated **too** often that I've decided to extract them into a separate repository.
> Though, I try to keep pkg level code for all projects in the `pkg/` directory as advised in many project structure layout guides in golang, the repetition got too much.
 
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
