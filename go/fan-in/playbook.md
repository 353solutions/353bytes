Go's select statement has a fixed number of `case` clauses. Sometimes we don't
know how many channels we'd like to listen on. One solution is to do a "fan in"
and consume all the channels into a single channel.

    fanin.go

- Start with producer and main
- time.Sleep & conversion from `rand.Intn->int` to time.Duration
- <- in chan 

    $ go run fanin.go

- deadlock
- "All problems in computer science can be solved by another level of
  indirection" - David Wheeler 
    All concurrency problems can be solved with another goroutine
- how can we know when all are done? sync.WaitGroup

    fanin_close.go
