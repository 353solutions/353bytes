If you noticed, numpy functions can work both on scalar (numbers) and on arrays.

```
In [1]: import numpy as np

In [2]: np.sin(90)
Out[2]: 0.8939966636005579  # why?

In [ ]: np.sin?

In [3]: np.sin(np.arange(10, 30, 5))
Out[3]: array([-0.54402111,  0.65028784,  0.91294525, -0.13235175])
```

It'll be cool to have this in functions we write. Let's try

```
In [1]: %run relu.py
In [2]: relu(30)
Out[2]: 30
In [3]: relu(-30)
Out[3]: 0
In [4]: import numpy as np
In [5]: values = np.array([-3, 7, 12, 0])
In [6]: values
Out[6]: array([-3,  7, 12,  0])
In [7]: relu(values)
... error
In [8]: if values:
   ...:     print('yes')
   ...:
```

To solve this we can use np.vectorize decorator

```
import numpy as np

@np.vectorize
def relu(n):
...
```

``
In [10]: relu(values)
Out[10]: array([ 0,  7, 12,  0])

In [11]: relu(5)
Out[11]: array(5)
```

Even though it looks funny, it's a number
```
In [11]: relu(5) * 3.32
```
