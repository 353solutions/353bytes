def add(x, y):
    """Addition"""
    return x + y


plus = add


def make_adder(n):
    """Returns a functions that add n to the argument"""
    def adder(val):
        return val + n

    return adder


add7 = make_adder(7)
