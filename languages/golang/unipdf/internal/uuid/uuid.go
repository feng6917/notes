//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package uuid ;import (_ga "crypto/rand";_f "encoding/hex";_gg "io";);func _dbf (_fff []byte ,_gf UUID ){_f .Encode (_fff ,_gf [:4]);_fff [8]='-';_f .Encode (_fff [9:13],_gf [4:6]);_fff [13]='-';_f .Encode (_fff [14:18],_gf [6:8]);_fff [18]='-';_f .Encode (_fff [19:23],_gf [8:10]);
_fff [23]='-';_f .Encode (_fff [24:],_gf [10:]);};var _ggf =_ga .Reader ;func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_d :=_gg .ReadFull (_ggf ,uuid [:]);if _d !=nil {return _e ,_d ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;
};func (_de UUID )String ()string {var _fe [36]byte ;_dbf (_fe [:],_de );return string (_fe [:])};var Nil =_e ;func MustUUID ()UUID {uuid ,_b :=NewUUID ();if _b !=nil {panic (_b );};return uuid ;};var _e UUID ;type UUID [16]byte ;