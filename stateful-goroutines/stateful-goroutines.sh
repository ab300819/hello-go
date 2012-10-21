# Running our program shows that the goroutine-based
# state management example achieves about 800,000
# operations per second.
$ go run stateful-goroutines.go
ops: 807434

# For this particular case the goroutine-based approach
# was a bit more involved than the mutex-based one. It
# might be useful in certain cases though, for example
# where you have other channels involved or when managing
# multiple such mutexes would be error-prone. The right
# approach for your program will depend on its particular
# details. You should use whatever approach feels most
# natural, especially with respect to understanding the
# correctness of your program.
