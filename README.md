# Simple and efficient serialize/deserialize library.

KTLV is acronym for Key-Type-Length-Value which mean
a method which used to encode and decode your data.

KTLV was inspired by [Apache Thrift](https://thrift.apache.org/).
Thrift is more mature and more wide used, but it's Erlang encoder
has a couple of disadvantages comparing with KTLV:

* it is 7 times slower in average;
* it produces a slightly bigger binary (near 8%);
* it's documentation is very, very poor and the source code is
 undocumented and contains no specifications.

On the other hand, KTLV provides only object encoding and
decoding features and it is not kinda protocol or transport
as Thrift is. Also there is no any
[Thrift IDL](https://thrift.apache.org/docs/idl) and IDL compiler
in KTLV. You should implement your own codec for the
programming language X by yourself instead.

The one of the most useful features KTLV taken from Thrift
is keyed elements, which can help to achieve backward and
forward compatibility between codecs.
