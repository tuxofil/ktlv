"""
Simple serialize/deserialize library.
"""

import struct

# ----------------------------------------------------------------------
# Known data type IDs

BOOL = 0
UINT8 = 1
UINT16 = 2
UINT24 = 3
UINT32 = 4
UINT64 = 5
DOUBLE = 6
STRING = 7
BITMAP = 8
INT8 = 9
INT16 = 10
INT24 = 11
INT32 = 12
INT64 = 13
LIST_OF_STRING = 50
LIST_OF_UINT8 = 51
LIST_OF_UINT16 = 52
LIST_OF_UINT24 = 53
LIST_OF_UINT32 = 54
LIST_OF_UINT64 = 55
LIST_OF_DOUBLE = 56
LIST_OF_INT8 = 57
LIST_OF_INT16 = 58
LIST_OF_INT24 = 59
LIST_OF_INT32 = 60
LIST_OF_INT64 = 61


MIN_INT8 = -0x80
MIN_INT16 = -0x8000
MIN_INT24 = -0x800000
MIN_INT32 = -0x80000000
MIN_INT64 = -0x8000000000000000
MAX_INT8 = 0x7f
MAX_INT16 = 0x7fff
MAX_INT24 = 0x7fffff
MAX_INT32 = 0x7fffffff
MAX_INT64 = 0x7fffffffffffffff
MIN_UINT8 = 0
MIN_UINT16 = 0
MIN_UINT24 = 0
MIN_UINT32 = 0
MIN_UINT64 = 0
MAX_UINT8 = 0xff
MAX_UINT16 = 0xffff
MAX_UINT24 = 0xffffff
MAX_UINT32 = 0xffffffff
MAX_UINT64 = 0xffffffffffffffff

def enc(elements):
    """
    Encode (key,type,value) list to a byte array.

    :param elements: element list to encode.
    :type elements: list of (key,type,value).

    :rtype: string
    """
    encoded = ''
    for key, dtype, value in elements:
        binary = enc_elem(dtype, value)
        encoded += struct.pack('>HBH', key, dtype, len(binary)) + binary
    return encoded


def dec(binary):
    """
    Decode byte array to (key,type,value) list.

    :param binary: encoded object.
    :type binary: string

    :rtype: list of (key,type,value)
    """
    result = []
    while binary:
        (key, dtype, length) = struct.unpack('>HBH', binary[:5])
        value = dec_elem(dtype, binary[5:5 + length])
        if value is not None:
            result.append((key, dtype, value))
        binary = binary[5 + length:]
    return result


def decd(binary):
    """
    Decode byte array to key -> (type,value) dictionary.

    :param binary: encoded object.
    :type binary: string

    :rtype: dict of (key,(type,value))
    """
    result = {}
    while binary:
        (key, dtype, length) = struct.unpack('>HBH', binary[:5])
        value = dec_elem(dtype, binary[5:5 + length])
        if value is not None:
            result[key] = (dtype, value)
        binary = binary[5 + length:]
    return result

# ----------------------------------------------------------------------
# Internal functions
# ----------------------------------------------------------------------

def dec_elem(dtype, binary):
    """
    Decode next element from the stream.

    :param dtype: data type identifier
    :type dtype: integer

    :param binary: input stream
    :type binary: string

    :rtype: value
    """
    if dtype == BOOL:
        return struct.unpack('>?', binary)[0]
    elif dtype == UINT8:
        return struct.unpack('>B', binary)[0]
    elif dtype == UINT16:
        return struct.unpack('>H', binary)[0]
    elif dtype == UINT24:
        (major, minor) = struct.unpack('>BH', binary)
        return (major << 16) + minor
    elif dtype == UINT32:
        return struct.unpack('>I', binary)[0]
    elif dtype == UINT64:
        return struct.unpack('>Q', binary)[0]
    elif dtype == INT8:
        return struct.unpack('>b', binary)[0]
    elif dtype == INT16:
        return struct.unpack('>h', binary)[0]
    elif dtype == INT24:
        (major, minor) = struct.unpack('>hB', binary)
        return (major << 8) + minor
    elif dtype == INT32:
        return struct.unpack('>i', binary)[0]
    elif dtype == INT64:
        return struct.unpack('>q', binary)[0]
    elif dtype == DOUBLE:
        return struct.unpack('>d', binary)[0]
    elif dtype == STRING:
        return binary
    elif dtype == BITMAP:
        (unused,) = struct.unpack('>B', binary[0])
        binary = binary[1:]
        bitsize = len(binary) * 8 - unused
        if bitsize == 0:
            return []
        result = []
        b = None
        bitpointer = 0
        while bitsize > 0:
            if b is None:
                b = struct.unpack('>B', binary[0])[0]
                binary = binary[1:]
                bitpointer = 0
            bitpointer += 1
            if unused > 0:
                unused -= 1
                b <<= 1
                continue
            result.append((b & 0b10000000) >> 7)
            bitsize -= 1
            b <<= 1
            if bitpointer == 8:
                b = None
        return result
    elif dtype == LIST_OF_STRING:
        result = []
        while binary:
            length = struct.unpack('>H', binary[:2])[0]
            result.append(binary[2:2 + length])
            binary = binary[2 + length:]
        return result
    elif dtype == LIST_OF_UINT8:
        return list(struct.unpack('>' + 'B' * len(binary), binary))
    elif dtype == LIST_OF_UINT16:
        return list(struct.unpack('>' + 'H' * (len(binary) / 2), binary))
    elif dtype == LIST_OF_UINT24:
        unpacked = struct.unpack('>' + 'BH' * (len(binary) / 3), binary)
        result = []
        while unpacked:
            result.append((unpacked[0] << 16) + unpacked[1])
            unpacked = unpacked[2:]
        return result
    elif dtype == LIST_OF_UINT32:
        return list(struct.unpack('>' + 'I' * (len(binary) / 4), binary))
    elif dtype == LIST_OF_UINT64:
        return list(struct.unpack('>' + 'Q' * (len(binary) / 8), binary))
    elif dtype == LIST_OF_INT8:
        return list(struct.unpack('>' + 'b' * len(binary), binary))
    elif dtype == LIST_OF_INT16:
        return list(struct.unpack('>' + 'h' * (len(binary) / 2), binary))
    elif dtype == LIST_OF_INT24:
        unpacked = struct.unpack('>' + 'hB' * (len(binary) / 3), binary)
        result = []
        while unpacked:
            result.append((unpacked[0] << 8) + unpacked[1])
            unpacked = unpacked[2:]
        return result
    elif dtype == LIST_OF_INT32:
        return list(struct.unpack('>' + 'i' * (len(binary) / 4), binary))
    elif dtype == LIST_OF_INT64:
        return list(struct.unpack('>' + 'q' * (len(binary) / 8), binary))
    elif dtype == LIST_OF_DOUBLE:
        return list(struct.unpack('>' + 'd' * (len(binary) / 8), binary))


def enc_elem(dtype, val):
    """
    Encode element.

    :param dtype: data type identifier
    :type dtype: integer

    :param val: element value
    :type val: any

    :rtype: binary
    """
    if dtype == BOOL:
        return struct.pack('>?', val)
    elif dtype == UINT8:
        return struct.pack('>B', val)
    elif dtype == UINT16:
        return struct.pack('>H', val)
    elif dtype == UINT24:
        return struct.pack('>HB', *divmod(val, 0x100))
    elif dtype == UINT32:
        return struct.pack('>I', val)
    elif dtype == UINT64:
        return struct.pack('>Q', val)
    elif dtype == INT8:
        return struct.pack('>b', val)
    elif dtype == INT16:
        return struct.pack('>h', val)
    elif dtype == INT24:
        return struct.pack('>hB', *divmod(val, 0x100))
    elif dtype == INT32:
        return struct.pack('>i', val)
    elif dtype == INT64:
        return struct.pack('>q', val)
    elif dtype == DOUBLE:
        return struct.pack('>d', val)
    elif dtype == STRING:
        return val
    elif dtype == LIST_OF_STRING:
        return ''.join([struct.pack('>H', len(e)) + e for e in val])
    elif dtype == LIST_OF_UINT8:
        return struct.pack('>' + 'B' * len(val), *val)
    elif dtype == LIST_OF_UINT16:
        return struct.pack('>' + 'H' * len(val), *val)
    elif dtype == LIST_OF_UINT24:
        return ''.join([struct.pack('>HB', *divmod(e, 0x100)) for e in val])
    elif dtype == LIST_OF_UINT32:
        return struct.pack('>' + 'I' * len(val), *val)
    elif dtype == LIST_OF_UINT64:
        return struct.pack('>' + 'Q' * len(val), *val)
    elif dtype == LIST_OF_INT8:
        return struct.pack('>' + 'b' * len(val), *val)
    elif dtype == LIST_OF_INT16:
        return struct.pack('>' + 'h' * len(val), *val)
    elif dtype == LIST_OF_INT24:
        return ''.join([struct.pack('>hB', *divmod(e, 0x100)) for e in val])
    elif dtype == LIST_OF_INT32:
        return struct.pack('>' + 'i' * len(val), *val)
    elif dtype == LIST_OF_INT64:
        return struct.pack('>' + 'q' * len(val), *val)
    elif dtype == LIST_OF_DOUBLE:
        return struct.pack('>' + 'd' * len(val), *val)
    elif dtype == BITMAP:
        unused = 8 - len(val) % 8
        val = [0] * unused + val
        fun = lambda a, n: (a << 1) + n
        ints = [reduce(fun, val[i:i + 8], 0)
                for i in range(0, len(val), 8)]
        return struct.pack('>B' + 'B' * len(ints), unused, *ints)
