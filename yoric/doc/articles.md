## Great articles about Golang

### GoLang, The Next Language to Learn for Developers

https://dev.to/dizdarevic/golang-the-next-language-to-learn-for-developers-pc5

Multi-threading at core

In one language:
- Efficient compilation
- Efficient execution
- Ease of programming

Concurrencency is dealing with many processes, rather than parallelism.

Channels are pipes that connect concurrent GoRoutines. You are able to send values and signals over Channels from GoRoutine to GoRoutine. This allows for synchronizing execution.

Usually, you would have to choose to either throw/return an exception / error, but because you can return tuples in Go, you can both throw an error and return a value when necessary.
This is important because Go doesn't have exceptions.

---


