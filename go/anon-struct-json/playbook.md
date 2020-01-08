Some APIs return complex data structure. Let's see an example, we'll ask
StockTwits about latest Netflix tweets.

    $ curl https://api.stocktwits.com/api/2/streams/symbol/NFLX.json

It's hard to read. When I usually do it indent the JSON and save it to a file.
I'm going to work in the terminal and use the
[jq](https://stedolan.github.io/jq/) utility. You can do the same with UI tools
such as [Postman](https://www.getpostman.com/) or [Insomnia](https://insomnia.rest/).


    $ curl https://api.stocktwits.com/api/2/streams/symbol/NFLX.json | jq > out.json
    $ vim out.json

We have `messages` array and inside every message object there's a `symbols`
array each having a `symbol` field.

Out task is, for a given symbol (such as `NFLX`). Print the list of other
symbols mentioned with it and count for each one.
    
Let's have a look at the code

    :e related.go

And now let's run this
    
    $ go run related.go

Sometimes to avoid hitting the API too much when writing the code. Instead of
making an HTTP request I save the output ones to a file and read the data from
the file.
