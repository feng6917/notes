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

package decoder ;import (_b "golang/unipdf/internal/bitwise";_ef "golang/unipdf/internal/jbig2/bitmap";_a "golang/unipdf/internal/jbig2/document";_f "golang/unipdf/internal/jbig2/errors";_d "image";
);type Decoder struct{_bf *_b .Reader ;_ad *_a .Document ;_fb int ;_g Parameters ;};func (_ea *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _ea .decodePage (pageNumber )};func (_ae *Decoder )DecodePageImage (pageNumber int )(_d .Image ,error ){const _ff ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
_ed ,_ede :=_ae .decodePageImage (pageNumber );if _ede !=nil {return nil ,_f .Wrap (_ede ,_ff ,"");};return _ed ,nil ;};func (_de *Decoder )decodePage (_gf int )([]byte ,error ){const _dc ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";
if _gf < 0{return nil ,_f .Errorf (_dc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_gf );};if _gf > int (_de ._ad .NumberOfPages ){return nil ,_f .Errorf (_dc ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_gf );
};_aeg ,_db :=_de ._ad .GetPage (_gf );if _db !=nil {return nil ,_f .Wrap (_db ,_dc ,"");};_be ,_db :=_aeg .GetBitmap ();if _db !=nil {return nil ,_f .Wrap (_db ,_dc ,"");};_be .InverseData ();if !_de ._g .UnpaddedData {return _be .Data ,nil ;};return _be .GetUnpaddedData ();
};func (_gb *Decoder )decodePageImage (_ga int )(_d .Image ,error ){const _aea ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _ga < 0{return nil ,_f .Errorf (_aea ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ga );
};if _ga > int (_gb ._ad .NumberOfPages ){return nil ,_f .Errorf (_aea ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ga );
};_daf ,_gag :=_gb ._ad .GetPage (_ga );if _gag !=nil {return nil ,_f .Wrap (_gag ,_aea ,"");};_beg ,_gag :=_daf .GetBitmap ();if _gag !=nil {return nil ,_f .Wrap (_gag ,_aea ,"");};_beg .InverseData ();return _beg .ToImage (),nil ;};func (_da *Decoder )PageNumber ()(int ,error ){const _bg ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";
if _da ._ad ==nil {return 0,_f .Error (_bg ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");};return int (_da ._ad .NumberOfPages ),nil ;};func Decode (input []byte ,parameters Parameters ,globals *_a .Globals )(*Decoder ,error ){_ee :=_b .NewReader (input );
_fe ,_c :=_a .DecodeDocument (_ee ,globals );if _c !=nil {return nil ,_c ;};return &Decoder {_bf :_ee ,_ad :_fe ,_g :parameters },nil ;};type Parameters struct{UnpaddedData bool ;Color _ef .Color ;};func (_bc *Decoder )DecodeNextPage ()([]byte ,error ){_bc ._fb ++;
_eff :=_bc ._fb ;return _bc .decodePage (_eff );};