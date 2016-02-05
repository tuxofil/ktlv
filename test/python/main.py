"""
Unit test for Python KTLV Codec.
"""

import struct
import unittest

from ktlv import *


def encdec(dtype, value):
    encoded = enc_elem(dtype, value)
    return dec_elem(dtype, encoded)


class KTLVTest(unittest.TestCase):
    """
    Unit test for Python KTLV Codec.
    """

    def enc_dec(self, pattern, dtype, value):
        encoded = enc_elem(dtype, value)
        self.assertEqual(pattern, dec_elem(dtype, encoded))

    def test_bool(self):
        self.enc_dec(1, BOOL, 1)
        self.enc_dec(0, BOOL, 0)
        self.enc_dec(1, BOOL, 5)
        self.enc_dec(1, BOOL, 5555555555)

    def test_uint8(self):
        self.enc_dec(MIN_UINT8, UINT8, MIN_UINT8)
        self.enc_dec(MAX_UINT8, UINT8, MAX_UINT8)

    def test_uint16(self):
        self.enc_dec(MIN_UINT16, UINT16, MIN_UINT16)
        self.enc_dec(MAX_UINT16, UINT16, MAX_UINT16)

    def test_uint24(self):
        self.enc_dec(MIN_UINT24, UINT24, MIN_UINT24)
        self.enc_dec(MAX_UINT24, UINT24, MAX_UINT24)

    def test_uint32(self):
        self.enc_dec(MIN_UINT32, UINT32, MIN_UINT32)
        self.enc_dec(MAX_UINT32, UINT32, MAX_UINT32)

    def test_uint64(self):
        self.enc_dec(MIN_UINT64, UINT64, MIN_UINT64)
        self.enc_dec(MAX_UINT64, UINT64, MAX_UINT64)

    def test_int8(self):
        self.enc_dec(MIN_INT8, INT8, MIN_INT8)
        self.enc_dec(0, INT8, 0)
        self.enc_dec(MAX_INT8, INT8, MAX_INT8)

    def test_int16(self):
        self.enc_dec(MIN_INT16, INT16, MIN_INT16)
        self.enc_dec(0, INT16, 0)
        self.enc_dec(MAX_INT16, INT16, MAX_INT16)

    def test_int24(self):
        self.enc_dec(MIN_INT24, INT24, MIN_INT24)
        self.enc_dec(0, INT24, 0)
        self.enc_dec(MAX_INT24, INT24, MAX_INT24)

    def test_int32(self):
        self.enc_dec(MIN_INT32, INT32, MIN_INT32)
        self.enc_dec(0, INT32, 0)
        self.enc_dec(MAX_INT32, INT32, MAX_INT32)

    def test_int64(self):
        self.enc_dec(MIN_INT64, INT64, MIN_INT64)
        self.enc_dec(0, INT64, 0)
        self.enc_dec(MAX_INT64, INT64, MAX_INT64)

    def test_double(self):
        self.enc_dec(0.0, DOUBLE, 0.0)
        self.enc_dec(-0.0, DOUBLE, -0.0)
        self.enc_dec(-1.0, DOUBLE, -1.0)
        self.enc_dec(1.0, DOUBLE, 1.0)

    def test_string(self):
        self.enc_dec('', STRING, '')
        self.enc_dec('a', STRING, 'a')
        self.enc_dec('abc', STRING, 'abc')

    def test_bitmap(self):
        self.enc_dec([], BITMAP, [])
        self.enc_dec([0], BITMAP, [0])
        self.enc_dec([1], BITMAP, [1])
        self.enc_dec([1,1], BITMAP, [1,1])
        self.enc_dec([1,1,0], BITMAP, [1,1,0])
        self.enc_dec([1,1,0,1,1,0,0,1,1], BITMAP, [1,1,0,1,1,0,0,1,1])

    def test_list_of_string(self):
        self.enc_dec([], STRING, [])
        self.enc_dec([''], STRING, [''])
        self.enc_dec(['', ''], STRING, ['', ''])
        self.enc_dec(['a', 'b'], STRING, ['a', 'b'])
        self.enc_dec(['a', 'bc'], STRING, ['a', 'bc'])

    def test_list_of_uint8(self):
        self.enc_dec([], LIST_OF_UINT8, [])
        self.enc_dec([MIN_UINT8, MAX_UINT8], LIST_OF_UINT8,
                     [MIN_UINT8, MAX_UINT8])

    def test_list_of_uint16(self):
        self.enc_dec([], LIST_OF_UINT16, [])
        self.enc_dec([MIN_UINT16, MAX_UINT16], LIST_OF_UINT16,
                     [MIN_UINT16, MAX_UINT16])

    def test_list_of_uint24(self):
        self.enc_dec([], LIST_OF_UINT24, [])
        self.enc_dec([MIN_UINT24, MAX_UINT24], LIST_OF_UINT24,
                     [MIN_UINT24, MAX_UINT24])

    def test_list_of_uint32(self):
        self.enc_dec([], LIST_OF_UINT32, [])
        self.enc_dec([MIN_UINT32, MAX_UINT32], LIST_OF_UINT32,
                     [MIN_UINT32, MAX_UINT32])

    def test_list_of_uint64(self):
        self.enc_dec([], LIST_OF_UINT64, [])
        self.enc_dec([MIN_UINT64, MAX_UINT64], LIST_OF_UINT64,
                     [MIN_UINT64, MAX_UINT64])

    def test_list_of_int8(self):
        self.enc_dec([], LIST_OF_INT8, [])
        self.enc_dec([MIN_INT8, 0, MAX_INT8], LIST_OF_INT8,
                     [MIN_INT8, 0, MAX_INT8])

    def test_list_of_int16(self):
        self.enc_dec([], LIST_OF_INT16, [])
        self.enc_dec([MIN_INT16, 0, MAX_INT16], LIST_OF_INT16,
                     [MIN_INT16, 0, MAX_INT16])

    def test_list_of_int24(self):
        self.enc_dec([], LIST_OF_INT24, [])
        self.enc_dec([MIN_INT24, 0, MAX_INT24], LIST_OF_INT24,
                     [MIN_INT24, 0, MAX_INT24])

    def test_list_of_int32(self):
        self.enc_dec([], LIST_OF_INT32, [])
        self.enc_dec([MIN_INT32, 0, MAX_INT32], LIST_OF_INT32,
                     [MIN_INT32, 0, MAX_INT32])

    def test_list_of_int64(self):
        self.enc_dec([], LIST_OF_INT64, [])
        self.enc_dec([MIN_INT64, 0, MAX_INT64], LIST_OF_INT64,
                     [MIN_INT64, 0, MAX_INT64])

    def test_list_of_double(self):
        self.enc_dec([], LIST_OF_DOUBLE, [])
        self.enc_dec([-1.0, 0.0, 1.0], LIST_OF_DOUBLE, [-1.0, 0.0, 1.0])

    def test_main(self):
        self.assertEqual([], dec(enc([])))
        self.assertEqual([(1, BOOL, 1)], dec(enc([(1, BOOL, 1)])))
        self.assertEqual(
            [(1, BOOL, 1),
             (2, UINT8, 2),
             (3, UINT16, 3),
             (4, UINT24, 4),
             (5, UINT32, 5),
             (6, UINT64, 6),
             (7, DOUBLE, 3.1415927),
             (8, INT8, -8),
             (9, INT16, -9),
             (10, INT24, -10),
             (11, INT32, -11),
             (12, INT64, -12),
             (13, STRING, 'str'),
             (14, LIST_OF_STRING, ['a', 'b']),
             (15, BITMAP, [1,1,0,1,1,1]),
             (16, LIST_OF_UINT8, [1, 2, 3]),
             (17, LIST_OF_UINT16, [1, 2, 3]),
             (18, LIST_OF_UINT24, [1, 2, 3]),
             (19, LIST_OF_UINT32, [1, 2, 3]),
             (20, LIST_OF_UINT64, [1, 2, 3]),
             (21, LIST_OF_INT8, [1, 2, 3]),
             (22, LIST_OF_INT16, [1, 2, 3]),
             (23, LIST_OF_INT24, [1, 2, 3]),
             (24, LIST_OF_INT32, [1, 2, 3]),
             (25, LIST_OF_INT64, [1, 2, 3]),
             (26, LIST_OF_DOUBLE, [-1.0, 0.0, 1.0])],
            dec(enc([(1, BOOL, 1),
                     (2, UINT8, 2),
                     (3, UINT16, 3),
                     (4, UINT24, 4),
                     (5, UINT32, 5),
                     (6, UINT64, 6),
                     (7, DOUBLE, 3.1415927),
                     (8, INT8, -8),
                     (9, INT16, -9),
                     (10, INT24, -10),
                     (11, INT32, -11),
                     (12, INT64, -12),
                     (13, STRING, 'str'),
                     (14, LIST_OF_STRING, ['a', 'b']),
                     (15, BITMAP, [1,1,0,1,1,1]),
                     (16, LIST_OF_UINT8, [1, 2, 3]),
                     (17, LIST_OF_UINT16, [1, 2, 3]),
                     (18, LIST_OF_UINT24, [1, 2, 3]),
                     (19, LIST_OF_UINT32, [1, 2, 3]),
                     (20, LIST_OF_UINT64, [1, 2, 3]),
                     (21, LIST_OF_INT8, [1, 2, 3]),
                     (22, LIST_OF_INT16, [1, 2, 3]),
                     (23, LIST_OF_INT24, [1, 2, 3]),
                     (24, LIST_OF_INT32, [1, 2, 3]),
                     (25, LIST_OF_INT64, [1, 2, 3]),
                     (26, LIST_OF_DOUBLE, [-1.0, 0.0, 1.0])])))
