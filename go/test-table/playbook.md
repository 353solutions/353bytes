Sometimes we'd like to test the same function with different input. The common
way to do it is with table driven testing.

Say we have a square root function we'd like to test

    sqrt.go

And we'd like to test it with 0, 2 maybe with negative number to see we get an
error...

    sqrt_test.go

- using testify
- anonymous struct
- t.Run, see which value cause the test
    
    $ go test -v

You can move this even one step further and store the test cases in a file.
This has the advantage that non-programmers (such product owners) can add
tests.

Let's move the tests to a YAML file.

    sqrt_cases.yml

And now the test looks like

    sqrt_yaml_test.go

- Exported
- Pass t, use Fatal/require
