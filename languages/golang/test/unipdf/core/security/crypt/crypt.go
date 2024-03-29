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

package crypt ;import (_c "crypto/aes";_ce "crypto/cipher";_ec "crypto/md5";_f "crypto/rand";_a "crypto/rc4";_e "fmt";_fe "lgo/test/unipdf/common";_dd "lgo/test/unipdf/core/security";_g "io";);func init (){_cbg ("\u0041\u0045\u0053V\u0032",_ca )};


// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_gd ,_ef :=_ca (FilterDict {});if _ef !=nil {_fe .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ef );
return filterAESV2 {};};return _gd ;};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};func _gdd (_dac ,_cg uint32 ,_eae []byte ,_bd bool )([]byte ,error ){_fee :=make ([]byte ,len (_eae )+5);copy (_fee ,_eae );for _ace :=0;_ace < 3;_ace ++{_fg :=byte ((_dac >>uint32 (8*_ace ))&0xff);
_fee [_ace +len (_eae )]=_fg ;};for _bcd :=0;_bcd < 2;_bcd ++{_bef :=byte ((_cg >>uint32 (8*_bcd ))&0xff);_fee [_bcd +len (_eae )+3]=_bef ;};if _bd {_fee =append (_fee ,0x73);_fee =append (_fee ,0x41);_fee =append (_fee ,0x6C);_fee =append (_fee ,0x54);
};_efc :=_ec .New ();_efc .Write (_fee );_ddb :=_efc .Sum (nil );if len (_eae )+5< 16{return _ddb [0:len (_eae )+5],nil ;};return _ddb ,nil ;};func _caf (_ab FilterDict )(Filter ,error ){if _ab .Length %8!=0{return nil ,_e .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_ab .Length );
};if _ab .Length < 5||_ab .Length > 16{if _ab .Length ==40||_ab .Length ==64||_ab .Length ==128{_fe .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_ab .Length );
_ab .Length /=8;}else {return nil ,_e .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_ab .Length );
};};return filterV2 {_fgf :_ab .Length },nil ;};func _b (_fbf FilterDict )(Filter ,error ){if _fbf .Length ==256{_fe .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_fbf .Length );
_fbf .Length /=8;};if _fbf .Length !=0&&_fbf .Length !=32{return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_fbf .Length );
};return filterAESV3 {},nil ;};type filterAESV3 struct{filterAES };

// PDFVersion implements Filter interface.
func (_dg filterV2 )PDFVersion ()[2]int {return [2]int {}};

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bbf ,_gce :=_a .NewCipher (okey );if _gce !=nil {return nil ,_gce ;};_fe .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );
_bbf .XORKeyStream (buf ,buf );_fe .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };func init (){_cbg ("\u0041\u0045\u0053V\u0033",_b )};


// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};var _ Filter =filterAESV3 {};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bea ,_ea :=_c .NewCipher (okey );if _ea !=nil {return nil ,_ea ;};if len (buf )< 16{_fe .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );
return buf ,_e .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_bb :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_fe .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_bb ),_bb );
_fe .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_e .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));
};_ae :=_ce .NewCBCDecrypter (_bea ,_bb );_fe .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_fe .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );
_ae .CryptBlocks (buf ,buf );_fe .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_fe .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");
return buf ,nil ;};_eg :=int (buf [len (buf )-1]);if _eg > len (buf ){_fe .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_eg ,len (buf ));
return buf ,_e .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_eg ];return buf ,nil ;};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };var _ Filter =filterAESV2 {};func _fa (_gccc string )(filterFunc ,error ){_egb :=_dfa [_gccc ];if _egb ==nil {return nil ,_e .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_gccc );
};return _egb ,nil ;};

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_ge ,_bc :=_caf (FilterDict {Length :length });if _bc !=nil {_fe .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_bc );
return filterV2 {_fgf :length };};return _ge ;};func _ca (_fb FilterDict )(Filter ,error ){if _fb .Length ==128{_fe .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_fb .Length );
_fb .Length /=8;};if _fb .Length !=0&&_fb .Length !=16{return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_fb .Length );
};return filterAESV2 {},nil ;};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};type filterFunc func (_fd FilterDict )(Filter ,error );func init (){_cbg ("\u0056\u0032",_caf )};var _ Filter =filterV2 {};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};

// Filter is a common interface for crypt filter methods.
type Filter interface{

// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
Name ()string ;

// KeyLength returns a length of the encryption key in bytes.
KeyLength ()int ;

// PDFVersion reports the minimal version of PDF document that introduced this filter.
PDFVersion ()[2]int ;

// HandlerVersion reports V and R parameters that should be used for this filter.
HandlerVersion ()(V ,R int );

// MakeKey generates a object encryption key based on file encryption key and object numbers.
// Used only for legacy filters - AESV3 doesn't change the key for each object.
MakeKey (_abg ,_cgb uint32 ,_bbb []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_egg []byte ,_db []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_ga []byte ,_dbf []byte )([]byte ,error );};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _gdd (objNum ,genNum ,ekey ,true );};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_gg ,_cf :=_a .NewCipher (okey );if _cf !=nil {return nil ,_cf ;};_fe .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_gg .XORKeyStream (buf ,buf );
_fe .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ac ,_cag :=_c .NewCipher (okey );if _cag !=nil {return nil ,_cag ;};_fe .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
const _gc =_c .BlockSize ;_ba :=_gc -len (buf )%_gc ;for _be :=0;_be < _ba ;_be ++{buf =append (buf ,byte (_ba ));};_fe .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_dfc :=make ([]byte ,_gc +len (buf ));
_da :=_dfc [:_gc ];if _ ,_gca :=_g .ReadFull (_f .Reader ,_da );_gca !=nil {return nil ,_gca ;};_dag :=_ce .NewCBCEncrypter (_ac ,_da );_dag .CryptBlocks (_dfc [_gc :],buf );buf =_dfc ;_fe .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
return buf ,nil ;};func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};type filterV2 struct{_fgf int };

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };


// MakeKey implements Filter interface.
func (_cba filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _gdd (objNum ,genNum ,ekey ,false );};var (_dfa =make (map[string ]filterFunc ););func (filterIdentity )KeyLength ()int {return 0};

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_dc ,_df :=_b (FilterDict {});if _df !=nil {_fe .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_df );
return filterAESV3 {};};return _dc ;};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _dd .AuthEvent ;Length int ;};type filterIdentity struct{};

// KeyLength implements Filter interface.
func (_fc filterV2 )KeyLength ()int {return _fc ._fgf };func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_fed ,_gfa :=_fa (d .CFM );if _gfa !=nil {return nil ,_gfa ;};_ddc ,_gfa :=_fed (d );if _gfa !=nil {return nil ,_gfa ;};return _ddc ,nil ;};type filterAES struct{};type filterAESV2 struct{filterAES };

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};func _cbg (_abb string ,_gdc filterFunc ){if _ ,_dbd :=_dfa [_abb ];_dbd {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_dfa [_abb ]=_gdc ;};

// HandlerVersion implements Filter interface.
func (_aeb filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};