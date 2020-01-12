There's a term in computer science GIGO - garbage in, garbage out

It crucial to look at the raw data before you start to process, your algorithm depends on it being correct.

    $ curl -o taxi.csv.bz2 http://j.mp/1Mtaxi

```
In [1]: csv_file = 'taxi.csv.bz2'

In [2]: from pathlib import Path

In [3]: csv_file = Path('taxi.csv.bz2')

In [4]: csv_file.stat().st_size
Out[4]: 9669573

In [5]: csv_file.stat().st_size/(2**20)
Out[5]: 9.221623420715332

In [6]: !ls -lh $csv_file
-rw-r--r-- 1 miki miki 9.3M Jan 12 18:27 taxi.csv.bz2

In [7]: import bz2

In [8]: with bz2.open(csv_file) as fp:
   ...:     print(sum(1 for _ in fp))
   ...:
1000002

In [9]: !bzcat $csv_file | wc -l
1000002

In [11]: with bz2.open(csv_file, 'rt') as fp:
    ...:     for lnum, line in enumerate(fp, 1):
    ...:         if lnum == 10:
    ...:             break
    ...:         print(line, end='')
    ...:
VendorID,tpep_pickup_datetime,tpep_dropoff_datetime,passenger_count,trip_distance,RatecodeID,store_and_fwd_flag,PULocationID,DOLocationID,payment_type,fare_amount,extra,mta_tax,tip_amount,tolls_amount,improvement_surcharge,total_amount

1,2018-11-01 00:51:36,2018-11-01 00:52:36,1,.00,1,N,145,145,2,2.5,0.5,0.5,0,0,0.3,3.8
1,2018-11-01 00:07:47,2018-11-01 00:21:43,1,2.30,1,N,142,164,1,11,0.5,0.5,2.45,0,0.3,14.75
1,2018-11-01 00:24:27,2018-11-01 00:34:29,1,1.80,1,N,164,48,2,8.5,0.5,0.5,0,0,0.3,9.8
1,2018-11-01 00:35:27,2018-11-01 00:47:02,1,2.30,1,N,48,107,1,10,0.5,0.5,3.35,0,0.3,14.65
1,2018-11-01 00:16:46,2018-11-01 00:22:50,1,1.00,1,N,163,170,2,6,0.5,0.5,0,0,0.3,7.3
1,2018-11-01 00:23:57,2018-11-01 00:34:29,1,2.10,1,N,170,79,2,9,0.5,0.5,0,0,0.3,10.3
1,2018-11-01 00:57:07,2018-11-01 01:05:27,1,1.60,1,N,148,79,2,7.5,0.5,0.5,0,0,0.3,8.8

In [12]:

In [12]: !bzcat $csv_file | head -10
VendorID,tpep_pickup_datetime,tpep_dropoff_datetime,passenger_count,trip_distance,RatecodeID,store_and_fwd_flag,PULocationID,DOLocationID,payment_type,fare_amount,extra,mta_tax,tip_amount,tolls_amount,improvement_surcharge,total_amount

1,2018-11-01 00:51:36,2018-11-01 00:52:36,1,.00,1,N,145,145,2,2.5,0.5,0.5,0,0,0.3,3.8
1,2018-11-01 00:07:47,2018-11-01 00:21:43,1,2.30,1,N,142,164,1,11,0.5,0.5,2.45,0,0.3,14.75
1,2018-11-01 00:24:27,2018-11-01 00:34:29,1,1.80,1,N,164,48,2,8.5,0.5,0.5,0,0,0.3,9.8
1,2018-11-01 00:35:27,2018-11-01 00:47:02,1,2.30,1,N,48,107,1,10,0.5,0.5,3.35,0,0.3,14.65
1,2018-11-01 00:16:46,2018-11-01 00:22:50,1,1.00,1,N,163,170,2,6,0.5,0.5,0,0,0.3,7.3
1,2018-11-01 00:23:57,2018-11-01 00:34:29,1,2.10,1,N,170,79,2,9,0.5,0.5,0,0,0.3,10.3
1,2018-11-01 00:57:07,2018-11-01 01:05:27,1,1.60,1,N,148,79,2,7.5,0.5,0.5,0,0,0.3,8.8
2,2018-11-01 01:03:28,2018-11-01 01:32:10,2,22.09,2,N,132,162,2,52,0,0.5,0,5.76,0.3,58.56

In [13]: import pandas as pd

In [14]: df = pd.read_csv(csv_file, nrows=1000)

In [15]: df.dtypes
Out[15]:
VendorID                   int64
tpep_pickup_datetime      object
tpep_dropoff_datetime     object
passenger_count            int64
trip_distance            float64
RatecodeID                 int64
store_and_fwd_flag        object
PULocationID               int64
DOLocationID               int64
payment_type               int64
fare_amount              float64
extra                    float64
mta_tax                  float64
tip_amount               float64
tolls_amount             float64
improvement_surcharge    float64
total_amount             float64
dtype: object

In [16]: df.memory_usage()
Out[16]:
Index                     128
VendorID                 8000
tpep_pickup_datetime     8000
tpep_dropoff_datetime    8000
passenger_count          8000
trip_distance            8000
RatecodeID               8000
store_and_fwd_flag       8000
PULocationID             8000
DOLocationID             8000
payment_type             8000
fare_amount              8000
extra                    8000
mta_tax                  8000
tip_amount               8000
tolls_amount             8000
improvement_surcharge    8000
total_amount             8000
dtype: int64

In [17]: df.memory_usage().sum()
Out[17]: 136128

In [18]: df.memory_usage().sum() / (2**20)
Out[18]: 0.12982177734375

In [19]: time_cols = ['tpep_pickup_datetime', 'tpep_dropoff_datetime']

In [20]: df = pd.read_csv(csv_file, nrows=1000, parse_dates=time_cols)

In [21]: df.memory_usage().sum() / (2**20)
Out[21]: 0.12982177734375

In [22]: df.dtypes
Out[22]:
VendorID                          int64
tpep_pickup_datetime     datetime64[ns]
tpep_dropoff_datetime    datetime64[ns]
passenger_count                   int64
trip_distance                   float64
RatecodeID                        int64
store_and_fwd_flag               object
PULocationID                      int64
DOLocationID                      int64
payment_type                      int64
fare_amount                     float64
extra                           float64
mta_tax                         float64
tip_amount                      float64
tolls_amount                    float64
improvement_surcharge           float64
total_amount                    float64
dtype: object

In [23]: df = pd.read_csv(csv_file, parse_dates=time_cols)

In [24]: df.memory_usage().sum() / (2**20)
Out[24]: 129.6998291015625
```



If you're interested in these command line tools and other, I recommend reading
[Data Science at the Command
Line](https://www.datascienceatthecommandline.com/)
