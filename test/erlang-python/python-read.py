#!/usr/bin/env python

import sys
import ktlv


LIST_MODEL = [(1, ktlv.BOOL, True),
              (2, ktlv.UINT8, 2),
              (3, ktlv.UINT16, 3),
              (4, ktlv.UINT24, 4),
              (5, ktlv.UINT32, 5),
              (6, ktlv.UINT64, 6),
              (7, ktlv.DOUBLE, 3.1415927),
              (8, ktlv.STRING, 'hello'),
              (9, ktlv.BITMAP, [1,1,0,0,1,0,1,1,1,1]),
              (10, ktlv.LIST_OF_STRING, ['hello', 'world', '!']),
              (11, ktlv.LIST_OF_UINT8, [1, 2, 3]),
              (12, ktlv.LIST_OF_UINT16, [4, 5, 6]),
              (13, ktlv.LIST_OF_UINT24, [7, 8, 9]),
              (14, ktlv.LIST_OF_UINT32, [10, 11, 12]),
              (15, ktlv.LIST_OF_UINT64, [13, 14, 15]),
              (16, ktlv.LIST_OF_DOUBLE, [1.1, 2.2, 3.3]),
              (17, ktlv.INT8, -2),
              (18, ktlv.INT16, -3),
              (19, ktlv.INT24, -4),
              (20, ktlv.INT32, -5),
              (21, ktlv.INT64, -6),
              (22, ktlv.LIST_OF_INT8, [1, -2, 3]),
              (23, ktlv.LIST_OF_INT16, [4, -5, 6]),
              (24, ktlv.LIST_OF_INT24, [7, -8, 9]),
              (25, ktlv.LIST_OF_INT32, [10, -11, 12]),
              (26, ktlv.LIST_OF_INT64, [13, -14, 15])]
DICT_MODEL = {k: (t, v) for k, t, v in LIST_MODEL}


if __name__ == '__main__':
    with open('object.bin') as fdescr:
        binary = fdescr.read()
    elements = ktlv.dec(binary)
    if elements == LIST_MODEL:
        print 'OK'
    else:
        print 'FAILED'
        print 'GOT:\n\t%r\n\nBUT EXPECT:\n\t%r\n' % (elements, LIST_MODEL)
        sys.exit(1)

    with open('object.bin') as fdescr:
        binary = fdescr.read()
    elements = ktlv.decd(binary)
    if elements == DICT_MODEL:
        print 'OK'
    else:
        print 'FAILED'
        print 'GOT:\n\t%r\n\nBUT EXPECT:\n\t%r\n' % (elements, DICT_MODEL)
        sys.exit(1)
