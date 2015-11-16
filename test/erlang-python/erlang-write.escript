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
              {11, ?list_of_uint8, [1, 2, 3]},
              {12, ?list_of_uint16, [4, 5, 6]},
              {13, ?list_of_uint24, [7, 8, 9]},
              {14, ?list_of_uint32, [10, 11, 12]},
              {15, ?list_of_uint64, [13, 14, 15]},
              {16, ?list_of_double, [1.1, 2.2, 3.3]},
              {17, ?int8, -2},
              {18, ?int16, -3},
              {19, ?int24, -4},
              {20, ?int32, -5},
              {21, ?int64, -6},
              {22, ?list_of_int8, [1, -2, 3]},
              {23, ?list_of_int16, [4, -5, 6]},
              {24, ?list_of_int24, [7, -8, 9]},
              {25, ?list_of_int32, [10, -11, 12]},
              {26, ?list_of_int64, [13, -14, 15]}
             ])).
