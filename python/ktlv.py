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


def enc(elements):
    """
    Encode (key,type,value) list to a byte array.

    :param elements: element list to encode.
    :type elements: list of (key,type,value).

    :rtype: string
    """
    raise NotImplemented


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
        return major * 256 + minor
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
        return major * 256 + minor
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
            result.append(unpacked[0] * 256 + unpacked[1])
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
            result.append(unpacked[0] * 256 + unpacked[1])
            unpacked = unpacked[2:]
        return result
    elif dtype == LIST_OF_INT32:
        return list(struct.unpack('>' + 'i' * (len(binary) / 4), binary))
    elif dtype == LIST_OF_INT64:
        return list(struct.unpack('>' + 'q' * (len(binary) / 8), binary))
    elif dtype == LIST_OF_DOUBLE:
        return list(struct.unpack('>' + 'd' * (len(binary) / 8), binary))
