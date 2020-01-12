Boolean indexing is a way to select DataFrame rows based on some conditions

```
import pandas as pd
In [2]: df = pd.read_csv('stocks.csv', parse_dates=['Date'])
In [3]: df
Out[3]:
         Date Symbol  Volume        Price
0  2014-02-06   AAPL      42    72.544289
1  2014-12-09   AAPL      95   109.349998
2  2015-07-28   INTC      30    28.440001
3  2015-09-09   AAPL      76   109.769997
4  2015-10-05   AAPL      21   109.070000
5  2016-01-12   GOOG      43   717.317017
6  2016-03-17   GOOG      37   736.000000
7  2016-08-22   INTC      64    35.119999
8  2017-05-01   GOOG      68   901.450012
9  2017-08-23   AMZN      34   954.200012
10 2017-11-03   AAPL      21   171.119995
11 2018-03-01   AMZN      10  1465.000000
12 2018-10-23   QCOM      53    63.570000
In [8]: df['Symbol'] == 'AAPL'
Out[8]:
0      True
1      True
2     False
3      True
4      True
5     False
6     False
7     False
8     False
9     False
10     True
11    False
12    False
Name: Symbol, dtype: bool
In [9]:mask = df['Symbol'] == 'AAPL'
In [10]: df[mask]
Out[10]:
         Date Symbol  Volume       Price
0  2014-02-06   AAPL      42   72.544289
1  2014-12-09   AAPL      95  109.349998
3  2015-09-09   AAPL      76  109.769997
4  2015-10-05   AAPL      21  109.070000
10 2017-11-03   AAPL      21  171.119995
In [11]: df[df['Symbol'] == 'AAPL']
Out[11]:
         Date Symbol  Volume       Price
0  2014-02-06   AAPL      42   72.544289
1  2014-12-09   AAPL      95  109.349998
3  2015-09-09   AAPL      76  109.769997
4  2015-10-05   AAPL      21  109.070000
10 2017-11-03   AAPL      21  171.119995
In [12]:df.query('Symbol == "AAPL"')
Out[12]:
         Date Symbol  Volume       Price
0  2014-02-06   AAPL      42   72.544289
1  2014-12-09   AAPL      95  109.349998
3  2015-09-09   AAPL      76  109.769997
4  2015-10-05   AAPL      21  109.070000
10 2017-11-03   AAPL      21  171.119995
In [15]: mask = (df['Symbol'] == 'AAPL') & (df['Volume'] > 30)
In [17]: df[~mask]
Out[17]:
         Date Symbol  Volume        Price
2  2015-07-28   INTC      30    28.440001
4  2015-10-05   AAPL      21   109.070000
5  2016-01-12   GOOG      43   717.317017
6  2016-03-17   GOOG      37   736.000000
7  2016-08-22   INTC      64    35.119999
8  2017-05-01   GOOG      68   901.450012
9  2017-08-23   AMZN      34   954.200012
10 2017-11-03   AAPL      21   171.119995
11 2018-03-01   AMZN      10  1465.000000
12 2018-10-23   QCOM      53    63.570000
In [18]: df[(df['Symbol'] == 'AAPL') | (df['Symbol'] == 'AMZN')]
Out[18]:
         Date Symbol  Volume        Price
0  2014-02-06   AAPL      42    72.544289
1  2014-12-09   AAPL      95   109.349998
3  2015-09-09   AAPL      76   109.769997
4  2015-10-05   AAPL      21   109.070000
9  2017-08-23   AMZN      34   954.200012
10 2017-11-03   AAPL      21   171.119995
11 2018-03-01   AMZN      10  1465.000000
```

Here's an exercise for you

    open snail.jpg


```
In [1]: import matplotlib.pyplot as plt
In [2]: img = plt.imread('snail.jpg').copy()
In [3]: img.shape
Out[3]: (590, 590, 3)
```

Draw a green (color=[0, 0xFF, 0]) square on the image,
- line width=5
- top left:     (217,  42)
- bottom right: (525, 275)

Pause the video and try to solve, I'll wait.

    snail.py
