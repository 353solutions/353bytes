Let's look at a simple test

    $ vim test_norm.py
    $ python -m pytest -v test_norm.py
    error

Why?
    
```
In [2]: import numpy as np
In [3]: v = np.array([0, 1, 1.1])
In [4]: v == v
Out[4]: array([ True,  True,  True])
In [5]: bool(Out[4])
---------------------------------------------------------------------------
ValueError                                Traceback (most recent call last)
<ipython-input-5-5d2478ae24c4> in <module>
----> 1 bool(Out[4])
ValueError: The truth value of an array with more than one element is ambiguous. Use a.any() or a.all()
In [6]: v.all()
Out[6]: False
In [7]: res = (v == v)
In [8]: res
Out[8]: array([ True,  True,  True])
In [9]: res.all()
Out[9]: True
In [10]: res.any()
Out[10]: True
```

Let's fix the test
```
    assert (expected == out).all(), 'bad normalize'
```

Hmm, why is that

```
In [1]: 1.1 * 1.1 == 1.21
Out[1]: False
```

floating points are inexact. Way reduce accuracy to gain speed.
```
In [1]: 1.1 * 1.1
Out[1]: 1.2100000000000002
```

To solve this, we have 

    np.allclose?

rotl, atol

Let's fix
    assert np.allclose(expected, out), 'bad normalize'
    $ python -m pytest -v test_norm.py

Now the test pass.

Another funny number we need to consider is np.nan - which is not a number

    v = np.array([0, 1, 1.1, np.nan])
    expected = np.array([0, 1.1, 1.21, np.nan])

    $ python -m pytest -v test_norm.py

fails.

    In [1]: np.nan == np.nan
    Out[1]: False

If we look at np.allclose doc again
    np.allclose?

We'll see we have an option for `equal_nan`. This depends on your application
logic if you consider np.nan equal or not.

    assert np.allclose(expected, out, equal_nan=True), 'bad normalize'

