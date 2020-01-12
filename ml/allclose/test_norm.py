import numpy as np


def normalize(v, n=1.1):
    return v * n


def test_norm():
    v = np.array([0, 1, 1.1])
    expected = np.array([0, 1.1, 1.21])
    out = normalize(v)
    assert expected == out, 'bad normalize'
