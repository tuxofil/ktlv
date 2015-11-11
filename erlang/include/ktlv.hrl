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

-define(list_of_string, 50).
-define(list_of_uint8, 51).
-define(list_of_uint16, 52).
-define(list_of_uint24, 53).
-define(list_of_uint32, 54).
-define(list_of_uint64, 55).
-define(list_of_double, 56).

-endif.
