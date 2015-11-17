%%% @doc
%%% Simple and efficient serialize/deserialize library.
%%%
%%% KTLV is acronym for Key-Type-Length-Value which mean
%%% a method which used to encode and decode your data.
%%%
%%% KTLV was inspired by [https://thrift.apache.org/](Apache Thrift).
%%% Thrift is more mature and more wide used, but it's encoder has
%%% a couple of disadvantages comparing with KTLV:
%%%
%%% <ul>
%%%  <li>it is 7 times slower in average;</li>
%%%  <li>it produces a slightly bigger binary (~8%);</li>
%%%  <li>it's documentation is very, very poor and the
%%%    source code is undocumented and contains
%%%    no specifications.</li>
%%% </ul>
%%%
%%% On the other hand, KTLV provides only object encoding and
%%% decoding features and it is not kinda protocol or transport
%%% as Thrift is. Also there is no any
%%% [https://thrift.apache.org/docs/idl](IDL) and IDL compiler
%%% in KTLV. You should implement your own codec for the
%%% programming language X by yourself instead.
%%%
%%% The one of the most useful features KTLV taken from Thrift
%%% is keyed elements, which can help to achieve backward and
%%% forward compatibility between codecs.

%%% @author Aleksey Morarash <aleksey.morarash@gmail.com>
%%% @since 10 Nov 2015
%%% @copyright 2015, Aleksey Morarash <aleksey.morarash@gmail.com>

-module(ktlv).

%% API exports
-export([enc/1, dec/1, decd/1]).

-include("ktlv.hrl").

-ifdef(TEST).
-include_lib("eunit/include/eunit.hrl").
-endif.

%% ----------------------------------------------------------------------
%% Type definitions
%% ----------------------------------------------------------------------

-export_type(
   [key/0,
    type/0,
    value/0,
    element/0,
    objectd/0,
    objectd_element/0,
    value_bool/0,
    value_uint8/0,
    value_uint16/0,
    value_uint24/0,
    value_uint32/0,
    value_uint64/0,
    value_int8/0,
    value_int16/0,
    value_int24/0,
    value_int32/0,
    value_int64/0,
    value_double/0,
    value_string/0
   ]).

-type key() :: 0..16#ffff.

-type type() :: ?bool | ?uint8 | ?uint16 | ?uint32 | ?uint64 | ?double |
                ?string | ?bitmap | ?int8 | ?int16 | ?int32 | ?int64 |
                ?list_of_uint8 | ?list_of_uint16 | ?list_of_uint24 |
                ?list_of_uint32 | ?list_of_uint64 | ?list_of_double |
                ?list_of_string | ?list_of_int8 | ?list_of_int16 |
                ?list_of_int24 | ?list_of_int32 | ?list_of_int64.

-type element() ::
        {key(), ?bool, value_bool()} |
        {key(), ?uint8, value_uint8()} |
        {key(), ?uint16, value_uint16()} |
        {key(), ?uint24, value_uint24()} |
        {key(), ?uint32, value_uint32()} |
        {key(), ?uint64, value_uint64()} |
        {key(), ?int8, value_int8()} |
        {key(), ?int16, value_int16()} |
        {key(), ?int24, value_int24()} |
        {key(), ?int32, value_int32()} |
        {key(), ?int64, value_int64()} |
        {key(), ?double, value_double()} |
        {key(), ?string, value_string()} |
        {key(), ?bitmap, [value_bool()]} |
        {key(), ?list_of_string, [value_string()]} |
        {key(), ?list_of_uint8, [value_uint8()]} |
        {key(), ?list_of_uint16, [value_uint16()]} |
        {key(), ?list_of_uint24, [value_uint24()]} |
        {key(), ?list_of_uint32, [value_uint32()]} |
        {key(), ?list_of_uint64, [value_uint64()]} |
        {key(), ?list_of_int8, [value_int8()]} |
        {key(), ?list_of_int16, [value_int16()]} |
        {key(), ?list_of_int24, [value_int24()]} |
        {key(), ?list_of_int32, [value_int32()]} |
        {key(), ?list_of_int64, [value_int64()]} |
        {key(), ?list_of_double, [value_double()]}.

-type objectd() :: dict:dict(key(), objectd_element()).
-type objectd_element() ::
        {?bool, value_bool()} |
        {?uint8, value_uint8()} |
        {?uint16, value_uint16()} |
        {?uint24, value_uint24()} |
        {?uint32, value_uint32()} |
        {?uint64, value_uint64()} |
        {?double, value_double()} |
        {?string, value_string()} |
        {?bitmap, [value_bool()]} |
        {?list_of_string, [value_string()]} |
        {?list_of_uint8, [value_uint8()]} |
        {?list_of_uint16, [value_uint16()]} |
        {?list_of_uint24, [value_uint24()]} |
        {?list_of_uint32, [value_uint32()]} |
        {?list_of_uint64, [value_uint64()]} |
        {?list_of_int8, [value_int8()]} |
        {?list_of_int16, [value_int16()]} |
        {?list_of_int24, [value_int24()]} |
        {?list_of_int32, [value_int32()]} |
        {?list_of_int64, [value_int64()]} |
        {?list_of_double, [value_double()]}.

-type value_bool() :: 0 | 1.
-type value_uint8() :: ?min_uint8..?max_uint8.
-type value_uint16() :: ?min_uint16..?max_uint16.
-type value_uint24() :: ?min_uint24..?max_uint24.
-type value_uint32() :: ?min_uint32..?max_uint32.
-type value_uint64() :: ?min_uint64..?max_uint64.
-type value_int8() :: ?min_int8..?max_int8.
-type value_int16() :: ?min_int16..?max_int16.
-type value_int24() :: ?min_int24..?max_int24.
-type value_int32() :: ?min_int32..?max_int32.
-type value_int64() :: ?min_int64..?max_int64.
-type value_double() :: float().
-type value_string() :: binary().

-type value() :: value_bool() | value_uint8() | value_uint16() |
                 value_uint24() | value_uint32() | value_uint64() |
                 value_double() | value_string() | [value_bool()] |
                 [value_uint8()] | [value_uint16()] | [value_uint24()] |
                 [value_uint32()] | [value_uint64()] | [value_double()] |
                 [value_string()] | [value_int8()] | [value_int16()] |
                 [value_int24()] | [value_int32()] | [value_int64()].

%% ----------------------------------------------------------------------
%% API functions
%% ----------------------------------------------------------------------

%% @doc Encode {Key,Type,Value} list to binary.
-spec enc([element()]) -> binary().
enc(List) ->
    << <<Key:16/unsigned-big, Type:8/unsigned-big, (enc(Type, Value))/binary>>
       || {Key, Type, Value} <- List>>.

%% @doc Decode binary to {Key,Type,Value} list.
-spec dec(binary()) -> [element()].
dec(<<>>) -> [];
dec(<<Key:16/unsigned-big, Type:8/unsigned-big, Len:16/unsigned-big, Tail/binary>>) ->
    case dec(Type, Len, Tail) of
        {Value, Tail2} ->
            [{Key, Type, Value} | dec(Tail2)];
        {Tail2} ->
            dec(Tail2)
    end.

%% @doc Decode binary to Key -> {Type,Value} dictionary.
-spec decd(binary()) -> objectd().
decd(Binary) ->
    decd(Binary, _Accum = dict:new()).

%% ----------------------------------------------------------------------
%% Internal functions
%% ----------------------------------------------------------------------

%% @doc Encode element.
-spec enc(type(), value()) -> binary().
enc(?bool, V) -> <<1:16/unsigned-big, 0:7, V:1/unsigned-big>>;
enc(?uint8, V) -> <<1:16/unsigned-big, V:8/unsigned-big>>;
enc(?uint16, V) -> <<2:16/unsigned-big, V:16/unsigned-big>>;
enc(?uint32, V) -> <<4:16/unsigned-big, V:32/unsigned-big>>;
enc(?uint24, V) -> <<3:16/unsigned-big, V:24/unsigned-big>>;
enc(?uint64, V) -> <<8:16/unsigned-big, V:64/unsigned-big>>;
enc(?int8, V) -> <<1:16/unsigned-big, V:8/signed-big>>;
enc(?int16, V) -> <<2:16/unsigned-big, V:16/signed-big>>;
enc(?int32, V) -> <<4:16/unsigned-big, V:32/signed-big>>;
enc(?int24, V) -> <<3:16/unsigned-big, V:24/signed-big>>;
enc(?int64, V) -> <<8:16/unsigned-big, V:64/signed-big>>;
enc(?double, V) -> <<8:16/unsigned-big, V:64/float-big>>;
enc(?string, V) -> <<(size(V)):16/unsigned-big, V/binary>>;
enc(?bitmap, M) ->
    BitString = << <<I:1>> || I <- M>>,
    BitSize = bit_size(BitString),
    ByteSize = byte_size(BitString),
    Unused = ByteSize * 8 - BitSize,
    Len = ByteSize + 1,
    <<Len:16/unsigned-big, Unused:8/unsigned-big, 0:Unused/unsigned-big,
      BitString/bits>>;
enc(?list_of_string, V) ->
    Encoded = << <<(size(I)):16/unsigned-big, I/binary>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_uint8, V) ->
    Encoded = << <<I:8/unsigned-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_uint16, V) ->
    Encoded = << <<I:16/unsigned-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_uint24, V) ->
    Encoded = << <<I:24/unsigned-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_uint32, V) ->
    Encoded = << <<I:32/unsigned-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_uint64, V) ->
    Encoded = << <<I:64/unsigned-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_int8, V) ->
    Encoded = << <<I:8/signed-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_int16, V) ->
    Encoded = << <<I:16/signed-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_int24, V) ->
    Encoded = << <<I:24/signed-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_int32, V) ->
    Encoded = << <<I:32/signed-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_int64, V) ->
    Encoded = << <<I:64/signed-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(?list_of_double, V) ->
    Encoded = << <<I:64/float-big>> || I <- V>>,
    <<(size(Encoded)):16/unsigned-big, Encoded/binary>>;
enc(16#ff, Binary) ->
    %% only for testing purposes
    Size = size(Binary),
    <<Size:16/unsigned-big, Binary/binary>>.

%% @doc Decode next element from the stream.
%% Return tuple of two element on successful decode and
%% one element tuple if type of the element is unknown.
-spec dec(type(), Bytes :: non_neg_integer(), binary()) ->
                 {value(), binary()} | {binary()}.
dec(?bool, _1, <<0:7, V:1/unsigned-big, Tail/binary>>) -> {V, Tail};
dec(?uint8, _1, <<V:8/unsigned-big, Tail/binary>>) -> {V, Tail};
dec(?uint16, _2, <<V:16/unsigned-big, Tail/binary>>) -> {V, Tail};
dec(?uint24, _3, <<V:24/unsigned-big, Tail/binary>>) -> {V, Tail};
dec(?uint32, _4, <<V:32/unsigned-big, Tail/binary>>) -> {V, Tail};
dec(?uint64, _8, <<V:64/unsigned-big, Tail/binary>>) -> {V, Tail};
dec(?int8, _1, <<V:8/signed-big, Tail/binary>>) -> {V, Tail};
dec(?int16, _2, <<V:16/signed-big, Tail/binary>>) -> {V, Tail};
dec(?int24, _3, <<V:24/signed-big, Tail/binary>>) -> {V, Tail};
dec(?int32, _4, <<V:32/signed-big, Tail/binary>>) -> {V, Tail};
dec(?int64, _8, <<V:64/signed-big, Tail/binary>>) -> {V, Tail};
dec(?double, _8, <<V:64/float-big, Tail/binary>>) -> {V, Tail};
dec(?string, Len, Tail) -> split_binary(Tail, Len);
dec(?bitmap, Len, <<Unused:8/unsigned-big, Tail/binary>>) ->
    ByteSize = Len - 1,
    BitSize = ByteSize * 8 - Unused,
    <<0:Unused/unsigned-big, BitString:BitSize/bits, Tail2/binary>> = Tail,
    {[I || <<I:1>> <= BitString], Tail2};
dec(?list_of_string, Len, Tail) ->
    {Encoded, Tail2} = split_binary(Tail, Len),
    {dec_str_list_loop(Encoded), Tail2};
dec(?list_of_uint8, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:8/unsigned-big>> <= EncodedList], Tail2};
dec(?list_of_uint16, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:16/unsigned-big>> <= EncodedList], Tail2};
dec(?list_of_uint24, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:24/unsigned-big>> <= EncodedList], Tail2};
dec(?list_of_uint32, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:32/unsigned-big>> <= EncodedList], Tail2};
dec(?list_of_uint64, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:64/unsigned-big>> <= EncodedList], Tail2};
dec(?list_of_int8, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:8/signed-big>> <= EncodedList], Tail2};
dec(?list_of_int16, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:16/signed-big>> <= EncodedList], Tail2};
dec(?list_of_int24, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:24/signed-big>> <= EncodedList], Tail2};
dec(?list_of_int32, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:32/signed-big>> <= EncodedList], Tail2};
dec(?list_of_int64, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:64/signed-big>> <= EncodedList], Tail2};
dec(?list_of_double, Len, Tail) ->
    {EncodedList, Tail2} = split_binary(Tail, Len),
    {[I || <<I:64/float-big>> <= EncodedList], Tail2};
dec(_UnknownType, Len, Tail) ->
    {_Unknown, Tail2} = split_binary(Tail, Len),
    {Tail2}.

%% @doc Decode list of strings.
-spec dec_str_list_loop(Encoded :: binary()) -> [binary()].
dec_str_list_loop(<<>>) -> [];
dec_str_list_loop(<<ItemLen:16/unsigned-big, Tail/binary>>) ->
    {Value, Tail2} = split_binary(Tail, ItemLen),
    [Value | dec_str_list_loop(Tail2)].

%% @doc Helper for the decd/1 function.
%% Decode object and fill a dictionary with the object elements.
-spec decd(Encoded :: binary(), Accum :: objectd()) -> objectd().
decd(<<>>, Accum) -> Accum;
decd(<<K:16/unsigned-big, T:8/unsigned-big, L:16/unsigned-big, Tail/binary>>, Accum) ->
    case dec(T, L, Tail) of
        {V, Tail2} ->
            decd(Tail2, dict:store(K, {T, V}, Accum));
        {Tail2} ->
            decd(Tail2, Accum)
    end.

%% ----------------------------------------------------------------------
%% Unit tests
%% ----------------------------------------------------------------------

-ifdef(TEST).

-spec encdec(type(), value()) -> value().
encdec(Type, Term) ->
    <<Len:16/unsigned-big, Binary/binary>> = iolist_to_binary(enc(Type, Term)),
    {Value, <<>>} = dec(Type, Len, Binary),
    Value.

bool_test_() ->
    [?_assertMatch(1, encdec(?bool, 1)),
     ?_assertMatch(0, encdec(?bool, 0)),
     ?_assertMatch(1, encdec(?bool, 5)),
     ?_assertMatch(1, encdec(?bool, 55555555555))
    ].

uint8_test_() ->
    [?_assertMatch(?min_uint8, encdec(?uint8, ?min_uint8)),
     ?_assertMatch(?max_uint8, encdec(?uint8, ?max_uint8)),
     ?_assertMatch(?max_uint8, encdec(?uint8, ?min_uint8 - 1)),
     ?_assertMatch(?min_uint8, encdec(?uint8, ?max_uint8 + 1))
    ].

uint16_test_() ->
    [?_assertMatch(?min_uint16, encdec(?uint16, ?min_uint16)),
     ?_assertMatch(?max_uint16, encdec(?uint16, ?max_uint16)),
     ?_assertMatch(?max_uint16, encdec(?uint16, ?min_uint16 - 1)),
     ?_assertMatch(?min_uint16, encdec(?uint16, ?max_uint16 + 1))
    ].

uint24_test_() ->
    [?_assertMatch(?min_uint24, encdec(?uint24, ?min_uint24)),
     ?_assertMatch(?max_uint24, encdec(?uint24, ?max_uint24)),
     ?_assertMatch(?max_uint24, encdec(?uint24, ?min_uint24 - 1)),
     ?_assertMatch(?min_uint24, encdec(?uint24, ?max_uint24 + 1))
    ].

uint32_test_() ->
    [?_assertMatch(?min_uint32, encdec(?uint32, ?min_uint32)),
     ?_assertMatch(?max_uint32, encdec(?uint32, ?max_uint32)),
     ?_assertMatch(?max_uint32, encdec(?uint32, ?min_uint32 - 1)),
     ?_assertMatch(?min_uint32, encdec(?uint32, ?max_uint32 + 1))
    ].

uint64_test_() ->
    [?_assertMatch(?min_uint64, encdec(?uint64, ?min_uint64)),
     ?_assertMatch(?max_uint64, encdec(?uint64, ?max_uint64)),
     ?_assertMatch(?max_uint64, encdec(?uint64, ?min_uint64 - 1)),
     ?_assertMatch(?min_uint64, encdec(?uint64, ?max_uint64 + 1))
    ].

int8_test_() ->
    [?_assertMatch(?min_int8, encdec(?int8, ?min_int8)),
     ?_assertMatch(0, encdec(?int8, 0)),
     ?_assertMatch(?max_int8, encdec(?int8, ?max_int8)),
     ?_assertMatch(?max_int8, encdec(?int8, ?min_int8 - 1)),
     ?_assertMatch(?min_int8, encdec(?int8, ?max_int8 + 1))
    ].

int16_test_() ->
    [?_assertMatch(?min_int16, encdec(?int16, ?min_int16)),
     ?_assertMatch(0, encdec(?int16, 0)),
     ?_assertMatch(?max_int16, encdec(?int16, ?max_int16)),
     ?_assertMatch(?max_int16, encdec(?int16, ?min_int16 - 1)),
     ?_assertMatch(?min_int16, encdec(?int16, ?max_int16 + 1))
    ].

int24_test_() ->
    [?_assertMatch(?min_int24, encdec(?int24, ?min_int24)),
     ?_assertMatch(0, encdec(?int24, 0)),
     ?_assertMatch(?max_int24, encdec(?int24, ?max_int24)),
     ?_assertMatch(?max_int24, encdec(?int24, ?min_int24 - 1)),
     ?_assertMatch(?min_int24, encdec(?int24, ?max_int24 + 1))
    ].

int32_test_() ->
    [?_assertMatch(?min_int32, encdec(?int32, ?min_int32)),
     ?_assertMatch(0, encdec(?int32, 0)),
     ?_assertMatch(?max_int32, encdec(?int32, ?max_int32)),
     ?_assertMatch(?max_int32, encdec(?int32, ?min_int32 - 1)),
     ?_assertMatch(?min_int32, encdec(?int32, ?max_int32 + 1))
    ].

int64_test_() ->
    [?_assertMatch(?min_int64, encdec(?int64, ?min_int64)),
     ?_assertMatch(0, encdec(?int64, 0)),
     ?_assertMatch(?max_int64, encdec(?int64, ?max_int64)),
     ?_assertMatch(?max_int64, encdec(?int64, ?min_int64 - 1)),
     ?_assertMatch(?min_int64, encdec(?int64, ?max_int64 + 1))
    ].

double_test_() ->
    [?_assertMatch(0.0, encdec(?double, 0.0)),
     ?_assertMatch(-0.0, encdec(?double, -0.0)),
     ?_assertMatch(-1.0, encdec(?double, -1.0)),
     ?_assertMatch(1.0, encdec(?double, 1.0))
    ].

string_test_() ->
    [?_assertMatch(<<>>, encdec(?string, <<>>)),
     ?_assertMatch(<<"a">>, encdec(?string, <<"a">>)),
     ?_assertMatch(<<"abc">>, encdec(?string, <<"abc">>))
    ].

bitmap_test_() ->
    [?_assertMatch([], encdec(?bitmap, [])),
     ?_assertMatch([0], encdec(?bitmap, [0])),
     ?_assertMatch([1], encdec(?bitmap, [1])),
     ?_assertMatch([1,1], encdec(?bitmap, [1,1])),
     ?_assertMatch([1,1,0], encdec(?bitmap, [1,1,0])),
     ?_assertMatch([1,1,0,1,1,0,0,1,1], encdec(?bitmap, [1,1,0,1,1,0,0,1,1]))
    ].

list_of_string_test_() ->
    [?_assertMatch([], encdec(?list_of_string, [])),
     ?_assertMatch([<<>>], encdec(?list_of_string, [<<>>])),
     ?_assertMatch([<<>>, <<>>], encdec(?list_of_string, [<<>>, <<>>])),
     ?_assertMatch([<<$a>>, <<$b>>], encdec(?list_of_string, [<<$a>>, <<$b>>]))
    ].

list_of_uint8_test_() ->
    [?_assertMatch([], encdec(?list_of_uint8, [])),
     ?_assertMatch([0], encdec(?list_of_uint8, [0])),
     ?_assertMatch([1], encdec(?list_of_uint8, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_uint8, [1, 1])),
     ?_assertMatch([1, 2, 3], encdec(?list_of_uint8, [1, 2, 3])),
     ?_assertMatch(
        [?min_uint8, ?max_uint8, ?max_uint8, ?min_uint8],
        encdec(?list_of_uint8,
               [?min_uint8, ?max_uint8, ?min_uint8 - 1, ?max_uint8 + 1]))
    ].

list_of_uint16_test_() ->
    [?_assertMatch([], encdec(?list_of_uint16, [])),
     ?_assertMatch([0], encdec(?list_of_uint16, [0])),
     ?_assertMatch([1], encdec(?list_of_uint16, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_uint16, [1, 1])),
     ?_assertMatch([1, 2, 3], encdec(?list_of_uint16, [1, 2, 3])),
     ?_assertMatch(
        [?min_uint16, ?max_uint16, ?max_uint16, ?min_uint16],
        encdec(?list_of_uint16,
               [?min_uint16, ?max_uint16, ?min_uint16 - 1, ?max_uint16 + 1]))
    ].

list_of_uint24_test_() ->
    [?_assertMatch([], encdec(?list_of_uint24, [])),
     ?_assertMatch([0], encdec(?list_of_uint24, [0])),
     ?_assertMatch([1], encdec(?list_of_uint24, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_uint24, [1, 1])),
     ?_assertMatch([1, 2, 3], encdec(?list_of_uint24, [1, 2, 3])),
     ?_assertMatch(
        [?min_uint24, ?max_uint24, ?max_uint24, ?min_uint24],
        encdec(?list_of_uint24,
               [?min_uint24, ?max_uint24, ?min_uint24 - 1, ?max_uint24 + 1]))
    ].

list_of_uint32_test_() ->
    [?_assertMatch([], encdec(?list_of_uint32, [])),
     ?_assertMatch([0], encdec(?list_of_uint32, [0])),
     ?_assertMatch([1], encdec(?list_of_uint32, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_uint32, [1, 1])),
     ?_assertMatch([1, 2, 3], encdec(?list_of_uint32, [1, 2, 3])),
     ?_assertMatch(
        [?min_uint32, ?max_uint32, ?max_uint32, ?min_uint32],
        encdec(?list_of_uint32,
               [?min_uint32, ?max_uint32, ?min_uint32 - 1, ?max_uint32 + 1]))
    ].

list_of_uint64_test_() ->
    [?_assertMatch([], encdec(?list_of_uint64, [])),
     ?_assertMatch([0], encdec(?list_of_uint64, [0])),
     ?_assertMatch([1], encdec(?list_of_uint64, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_uint64, [1, 1])),
     ?_assertMatch([1, 2, 3], encdec(?list_of_uint64, [1, 2, 3])),
     ?_assertMatch(
        [?min_uint64, ?max_uint64, ?max_uint64, ?min_uint64],
        encdec(?list_of_uint64,
               [?min_uint64, ?max_uint64, ?min_uint64 - 1, ?max_uint64 + 1]))
    ].

list_of_int8_test_() ->
    [?_assertMatch([], encdec(?list_of_int8, [])),
     ?_assertMatch([0], encdec(?list_of_int8, [0])),
     ?_assertMatch([1], encdec(?list_of_int8, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_int8, [1, 1])),
     ?_assertMatch([1, -2, 3], encdec(?list_of_int8, [1, -2, 3])),
     ?_assertMatch(
        [0, ?min_int8, ?max_int8, ?max_int8, ?min_int8],
        encdec(?list_of_int8,
               [0, ?min_int8, ?max_int8, ?min_int8 - 1, ?max_int8 + 1]))
    ].

list_of_int16_test_() ->
    [?_assertMatch([], encdec(?list_of_int16, [])),
     ?_assertMatch([0], encdec(?list_of_int16, [0])),
     ?_assertMatch([1], encdec(?list_of_int16, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_int16, [1, 1])),
     ?_assertMatch([1, -2, 3], encdec(?list_of_int16, [1, -2, 3])),
     ?_assertMatch(
        [0, ?min_int16, ?max_int16, ?max_int16, ?min_int16],
        encdec(?list_of_int16,
               [0, ?min_int16, ?max_int16, ?min_int16 - 1, ?max_int16 + 1]))
    ].

list_of_int24_test_() ->
    [?_assertMatch([], encdec(?list_of_int24, [])),
     ?_assertMatch([0], encdec(?list_of_int24, [0])),
     ?_assertMatch([1], encdec(?list_of_int24, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_int24, [1, 1])),
     ?_assertMatch([1, -2, 3], encdec(?list_of_int24, [1, -2, 3])),
     ?_assertMatch(
        [0, ?min_int24, ?max_int24, ?max_int24, ?min_int24],
        encdec(?list_of_int24,
               [0, ?min_int24, ?max_int24, ?min_int24 - 1, ?max_int24 + 1]))
    ].

list_of_int32_test_() ->
    [?_assertMatch([], encdec(?list_of_int32, [])),
     ?_assertMatch([0], encdec(?list_of_int32, [0])),
     ?_assertMatch([1], encdec(?list_of_int32, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_int32, [1, 1])),
     ?_assertMatch([1, -2, 3], encdec(?list_of_int32, [1, -2, 3])),
     ?_assertMatch(
        [0, ?min_int32, ?max_int32, ?max_int32, ?min_int32],
        encdec(?list_of_int32,
               [0, ?min_int32, ?max_int32, ?min_int32 - 1, ?max_int32 + 1]))
    ].

list_of_int64_test_() ->
    [?_assertMatch([], encdec(?list_of_int64, [])),
     ?_assertMatch([0], encdec(?list_of_int64, [0])),
     ?_assertMatch([1], encdec(?list_of_int64, [1])),
     ?_assertMatch([1, 1], encdec(?list_of_int64, [1, 1])),
     ?_assertMatch([1, -2, 3], encdec(?list_of_int64, [1, -2, 3])),
     ?_assertMatch(
        [0, ?min_int64, ?max_int64, ?max_int64, ?min_int64],
        encdec(?list_of_int64,
               [0, ?min_int64, ?max_int64, ?min_int64 - 1, ?max_int64 + 1]))
    ].

main_test_() ->
    [?_assertMatch([{1, ?bool, 1}], dec(enc([{1, ?bool, 1}]))),
     ?_assertMatch([{1, ?bool, 1},
                    {2, ?uint8, 2},
                    {3, ?uint16, 3},
                    {4, ?uint32, 4},
                    {5, ?uint64, 5},
                    {6, ?double, 3.1415927},
                    {7, ?string, <<"str">>},
                    {8, ?list_of_string, [<<"a">>, <<"b">>]},
                    {9, ?bitmap, [1,1,0,1,1,1]},
                    {10, ?list_of_uint8, [1, 2, 3]}],
                   dec(enc([{1, ?bool, 1},
                            {2, ?uint8, 2},
                            {3, ?uint16, 3},
                            {4, ?uint32, 4},
                            {5, ?uint64, 5},
                            {11, 255, <<"hello">>},
                            {6, ?double, 3.1415927},
                            {7, ?string, <<"str">>},
                            {8, ?list_of_string, [<<"a">>, <<"b">>]},
                            {9, ?bitmap, [1,1,0,1,1,1]},
                            {10, ?list_of_uint8, [1, 2, 3]}])))
    ].

dict_test_() ->
    D = decd(enc([{1, ?bool, 1},
                  {2, ?uint8, 2},
                  {3, ?uint16, 3},
                  {4, ?uint32, 4},
                  {5, ?uint64, 5},
                  {6, ?double, 3.1415927},
                  {7, ?string, <<"str">>},
                  {8, ?list_of_string, [<<"a">>, <<"b">>]},
                  {9, ?bitmap, [1,1,0,1,1,1]},
                  {10, ?list_of_uint8, [1, 2, 3]}])),
    [?_assertMatch({?bool, 1}, dict:fetch(1, D)),
     ?_assertMatch({?uint8, 2}, dict:fetch(2, D)),
     ?_assertMatch({?uint16, 3}, dict:fetch(3, D)),
     ?_assertMatch({?uint32, 4}, dict:fetch(4, D)),
     ?_assertMatch({?uint64, 5}, dict:fetch(5, D)),
     ?_assertMatch({?double, 3.1415927}, dict:fetch(6, D)),
     ?_assertMatch({?string, <<"str">>}, dict:fetch(7, D)),
     ?_assertMatch({?list_of_string, [<<"a">>, <<"b">>]}, dict:fetch(8, D)),
     ?_assertMatch({?bitmap, [1,1,0,1,1,1]}, dict:fetch(9, D)),
     ?_assertMatch({?list_of_uint8, [1, 2, 3]}, dict:fetch(10, D))
    ].

-endif.
