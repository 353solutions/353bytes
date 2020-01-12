```
In [1]: import pandas as pd

In [2]: df = pd.read_csv('cart.csv', parse_dates=['Date'])

In [3]: df
Out[3]:
        Date Customer    Item  Amount  Item Price
0 2020-01-02     Rick     Eel     231         3.2
1 2020-01-07    Morty  Carrot       7         0.3
2 2020-01-07     Beth   Vodka       3        23.5
3 2020-01-10     Rick    Acid      33         4.5
4 2020-01-12    Morty     Egg      12         5.2

In [4]: df[2]
---------------------------------------------------------------------------
KeyError                                  Traceback (most recent call last)
...
KeyError: 2

In [5]: df.loc[2]
Out[5]:
Date          2020-01-07 00:00:00
Customer                     Beth
Item                        Vodka
Amount                          3
Item Price                   23.5
Name: 2, dtype: object

In [5]: df.iloc[2]
Out[5]:
Date          2020-01-07 00:00:00
Customer                     Beth
Item                        Vodka
Amount                          3
Item Price                   23.5
Name: 2, dtype: object

In [6]: df[2:4]
Out[6]:
        Date Customer   Item  Amount  Item Price
2 2020-01-07     Beth  Vodka       3        23.5
3 2020-01-10     Rick   Acid      33         4.5

In [7]: df.iloc[2:4]
Out[7]:
        Date Customer   Item  Amount  Item Price
2 2020-01-07     Beth  Vodka       3        23.5
3 2020-01-10     Rick   Acid      33         4.5

In [8]: df.loc[2:4]
Out[8]:
        Date Customer   Item  Amount  Item Price
2 2020-01-07     Beth  Vodka       3        23.5
3 2020-01-10     Rick   Acid      33         4.5
4 2020-01-12    Morty    Egg      12         5.2

In [9]: df.index
Out[9]: RangeIndex(start=0, stop=5, step=1)

In [10]: df = pd.read_csv('cart.csv', parse_dates=['Date'], index_col='Date')

In [11]: df.loc[2]
---------------------------------------------------------------------------
TypeError                                 Traceback (most recent call last)
...
TypeError: cannot do index indexing on <class 'pandas.core.indexes.datetimes.DatetimeIndex'> with these indexers [2] of <class 'int'>

In [12]: df.index
Out[12]:
DatetimeIndex(['2020-01-02', '2020-01-07', '2020-01-07', '2020-01-10',
               '2020-01-12'],
              dtype='datetime64[ns]', name='Date', freq=None)

In [13]: df.loc['2020-01-07']
Out[13]:
           Customer    Item  Amount  Item Price
Date
2020-01-07    Morty  Carrot       7         0.3
2020-01-07     Beth   Vodka       3        23.5

In [14]: df.iloc[2]
Out[14]:
Customer       Beth
Item          Vodka
Amount            3
Item Price     23.5
Name: 2020-01-07 00:00:00, dtype: object

In [15]: df.loc['2020-01']
Out[15]:
           Customer    Item  Amount  Item Price
Date
2020-01-02     Rick     Eel     231         3.2
2020-01-07    Morty  Carrot       7         0.3
2020-01-07     Beth   Vodka       3        23.5
2020-01-10     Rick    Acid      33         4.5
2020-01-12    Morty     Egg      12         5.2
```
