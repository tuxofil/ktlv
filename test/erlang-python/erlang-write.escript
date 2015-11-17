#!/usr/bin/env escript
%%! -pa ../../erlang/ebin

-include("ktlv.hrl").

main(_Args) ->
    ok = file:write_file(
           "object.bin",
           ktlv:enc(
             [{1, ?bool, 1},
              {2, ?uint8, 2},
              {3, ?uint16, 3},
              {4, ?uint24, 4},
              {5, ?uint32, 5},
              {6, ?uint64, 6},
              {7, ?double, 3.1415927},
              {8, ?string, <<"hello">>},
              {9, ?bitmap, [1,1,0,0,1,0,1,1,1,1]},
              {10, ?list_of_string, [<<"hello">>, <<"world">>, <<"!">>]},
              {11, ?list_of_uint8,
               [?min_uint8, ?max_uint8, ?min_uint8 - 1, ?max_uint8 + 1]},
              {12, ?list_of_uint16,
               [?min_uint16, ?max_uint16, ?min_uint16 - 1, ?max_uint16 + 1]},
              {13, ?list_of_uint24,
               [?min_uint24, ?max_uint24, ?min_uint24 - 1, ?max_uint24 + 1]},
              {14, ?list_of_uint32,
               [?min_uint32, ?max_uint32, ?min_uint32 - 1, ?max_uint32 + 1]},
              {15, ?list_of_uint64,
               [?min_uint64, ?max_uint64, ?min_uint64 - 1, ?max_uint64 + 1]},
              {16, ?list_of_double, [1.1, -2.2, 3.3]},
              {17, ?int8, -2},
              {18, ?int16, -3},
              {19, ?int24, -4},
              {20, ?int32, -5},
              {21, ?int64, -6},
              {22, ?list_of_int8,
               [0, ?min_int8, ?max_int8, ?min_int8 - 1, ?max_int8 + 1]},
              {23, ?list_of_int16,
               [0, ?min_int16, ?max_int16, ?min_int16 - 1, ?max_int16 + 1]},
              {24, ?list_of_int24,
               [0, ?min_int24, ?max_int24, ?min_int24 - 1, ?max_int24 + 1]},
              {25, ?list_of_int32,
               [0, ?min_int32, ?max_int32, ?min_int32 - 1, ?max_int32 + 1]},
              {26, ?list_of_int64,
               [0, ?min_int64, ?max_int64, ?min_int64 - 1, ?max_int64 + 1]}
             ])).
