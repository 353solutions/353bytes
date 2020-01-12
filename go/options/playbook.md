Dave Cheney has a great article (one of many) about [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) I think it's worth knowing about.

Say you have a server. It starts simple with listening on a port, then you add verbosity, then SSL and for SSL you need the certificate file location or ... You see where this is going. The call to `NewServer` will have many parameters and users will use only a small subset of them.

You can say, "OK - I'll create an options struct". The problem here is that you expose how you store options internally to the user and you can't refactor the options sturct without breaking changes - which is bad.

The solution that Dave found, is that an option is a function working on the server. And you expose these option functions. Combined with variable number of arguments this creates a nice solution.

    options.go

https://godoc.org/google.golang.org/grpc#Dial
