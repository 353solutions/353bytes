We're going to write a HTTP client using sockets. I find that tinkering with the low level helps me understand concepts better. In real life you'll use the net/http package.

    httpc.go

- crypto/tls since most use https
- reqTemplate
    - Close connection (http 2)
- tls.Dial nil - default options
- io.Copy IRL limit size  
- Talk on response

https://www.flickr.com/photos/girliemac/sets/72157628409467125/
https://http.cat/
