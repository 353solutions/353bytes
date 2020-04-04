During program life, it'll probably uses some resource. A resource is something
that you need to release when done using it. One resource, the memory, is
handled by Go using the garbage collector. However other resources such as
files, sockets, database connections, locks and others need to be manually
returned or released.

You might ask why do we care if we release (close) a file or a socket? The
answer is that there is a limit on the number of open files a process can hold.

    $ ulimit -a

You can see that I can have only 1024 file descriptors. These bugs of
forgetting to close a file are hard. Your server will work fine for a couple of
day and then start failing with the cryptic message of "too many open files".
Good luck finding the error.

Once you acquire a resource, you want to make sure it's released - even if
there was an error (panic) in your code. The `defer` statement does just that.

In other language we use try/finally to have the same effect. I find `defer`
much more elegant.

Let's have a look

    $ vim defer.go
    $ go run defer.go

You see that "defer" is printed after "done". Deferred are executed when a
function exists. And they work only at function level, don't use them in a
`for` loop for example.

Let's add a panic and run again

    panic("oops")
    $ go run defer.go

We see that "done" is not printed, but "defer" does print.

What happens if we have more than one defer?

    // panic("oops")
    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")

We see that "defer 2" is printed before "defer 1". Deffers are executed in
reversed order. This enables us to have a hierarchy of resources. Say open a
database connection and get some rows. We'd like to close the rows first and
then the database connection.

Here's a more realistic example:
    $ vim sha1.go
