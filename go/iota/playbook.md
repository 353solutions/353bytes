In some cases we'd like to have an enumeration of options. In Go we don't have
enum such as in other languages - we have iota

Say we'd like to enumerate logging levels. We can use strings but they take memory and comparing them is slower than comparing two integers.

    level.go

- iota (start 0)
- MaxLevel
- String() string (fmt.Stringer)
- 28 be careful not to use %s

In other cases, we pass a bitmask. It's a number where every bit in the number is an enumeration or an option. For example when looking at file permission on linux

    $ ls -l level.go

We can see read/write/execute for user, group and other. We have 6 bits and can
encode these permissions in a single byte.

We can use iota here as well.

    perm.go

- << operator
- MaxPerm
- String with compound
