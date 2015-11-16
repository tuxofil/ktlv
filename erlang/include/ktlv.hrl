%%% ---------------------------------------------------------------------
%%% File        : ktlv.hrl
%%% Author      : Aleksey Morarash <aleksey.morarash@gmail.com>
%%% Description : ktlv definitions file
%%% Created     : 10 Nov 2015
%%% ---------------------------------------------------------------------

-ifndef(_KTLV).
-define(_KTLV, true).

%% ----------------------------------------------------------------------
%% Data type IDs

-define(bool, 0).
-define(uint8, 1).
-define(uint16, 2).
-define(uint24, 3).
-define(uint32, 4).
-define(uint64, 5).
-define(double, 6).
-define(string, 7).
-define(bitmap, 8).
-define(int8, 9).
-define(int16, 10).
-define(int24, 11).
-define(int32, 12).
-define(int64, 13).

-define(list_of_string, 50).
-define(list_of_uint8, 51).
-define(list_of_uint16, 52).
-define(list_of_uint24, 53).
-define(list_of_uint32, 54).
-define(list_of_uint64, 55).
-define(list_of_double, 56).
-define(list_of_int8, 57).
-define(list_of_int16, 58).
-define(list_of_int24, 59).
-define(list_of_int32, 60).
-define(list_of_int64, 61).

-define(min_int8, -16#80).
-define(min_int16, -16#8000).
-define(min_int24, -16#800000).
-define(min_int32, -16#80000000).
-define(min_int64, -16#8000000000000000).
-define(max_int8, 16#7f).
-define(max_int16, 16#7fff).
-define(max_int24, 16#7fffff).
-define(max_int32, 16#7fffffff).
-define(max_int64, 16#7fffffffffffffff).
-define(min_uint8, 0).
-define(min_uint16, 0).
-define(min_uint24, 0).
-define(min_uint32, 0).
-define(min_uint64, 0).
-define(max_uint8, 16#ff).
-define(max_uint16, 16#ffff).
-define(max_uint24, 16#ffffff).
-define(max_uint32, 16#ffffffff).
-define(max_uint64, 16#ffffffffffffffff).

-endif.
