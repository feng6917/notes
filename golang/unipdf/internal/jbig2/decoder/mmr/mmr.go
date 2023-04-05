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

package mmr ;import (_g "errors";_f "fmt";_e "golang/unipdf/common";_a "golang/unipdf/internal/bitwise";_ce "golang/unipdf/internal/jbig2/bitmap";_cf "io";);func (_aaa *Decoder )UncompressMMR ()(_db *_ce .Bitmap ,_dc error ){_db =_ce .New (_aaa ._ff ,_aaa ._bg );
_ffe :=make ([]int ,_db .Width +5);_aab :=make ([]int ,_db .Width +5);_aab [0]=_db .Width ;_ec :=1;var _gbd int ;for _cba :=0;_cba < _db .Height ;_cba ++{_gbd ,_dc =_aaa .uncompress2d (_aaa ._gb ,_aab ,_ec ,_ffe ,_db .Width );if _dc !=nil {return nil ,_dc ;
};if _gbd ==EOF {break ;};if _gbd > 0{_dc =_aaa .fillBitmap (_db ,_cba ,_ffe ,_gbd );if _dc !=nil {return nil ,_dc ;};};_aab ,_ffe =_ffe ,_aab ;_ec =_gbd ;};if _dc =_aaa .detectAndSkipEOL ();_dc !=nil {return nil ,_dc ;};_aaa ._gb .align ();return _db ,nil ;
};type Decoder struct{_ff ,_bg int ;_gb *runData ;_bfa []*code ;_acb []*code ;_ga []*code ;};var (_ab =[][3]int {{4,0x1,int (_ge )},{3,0x1,int (_gg )},{1,0x1,int (_eg )},{3,0x3,int (_fge )},{6,0x3,int (_fb )},{7,0x3,int (_be )},{3,0x2,int (_gf )},{6,0x2,int (_bc )},{7,0x2,int (_aa )},{10,0xf,int (_fgg )},{12,0xf,int (_gd )},{12,0x1,int (EOL )}};
_ag =[][3]int {{4,0x07,2},{4,0x08,3},{4,0x0B,4},{4,0x0C,5},{4,0x0E,6},{4,0x0F,7},{5,0x12,128},{5,0x13,8},{5,0x14,9},{5,0x1B,64},{5,0x07,10},{5,0x08,11},{6,0x17,192},{6,0x18,1664},{6,0x2A,16},{6,0x2B,17},{6,0x03,13},{6,0x34,14},{6,0x35,15},{6,0x07,1},{6,0x08,12},{7,0x13,26},{7,0x17,21},{7,0x18,28},{7,0x24,27},{7,0x27,18},{7,0x28,24},{7,0x2B,25},{7,0x03,22},{7,0x37,256},{7,0x04,23},{7,0x08,20},{7,0xC,19},{8,0x12,33},{8,0x13,34},{8,0x14,35},{8,0x15,36},{8,0x16,37},{8,0x17,38},{8,0x1A,31},{8,0x1B,32},{8,0x02,29},{8,0x24,53},{8,0x25,54},{8,0x28,39},{8,0x29,40},{8,0x2A,41},{8,0x2B,42},{8,0x2C,43},{8,0x2D,44},{8,0x03,30},{8,0x32,61},{8,0x33,62},{8,0x34,63},{8,0x35,0},{8,0x36,320},{8,0x37,384},{8,0x04,45},{8,0x4A,59},{8,0x4B,60},{8,0x5,46},{8,0x52,49},{8,0x53,50},{8,0x54,51},{8,0x55,52},{8,0x58,55},{8,0x59,56},{8,0x5A,57},{8,0x5B,58},{8,0x64,448},{8,0x65,512},{8,0x67,640},{8,0x68,576},{8,0x0A,47},{8,0x0B,48},{9,0x01,_ca },{9,0x98,1472},{9,0x99,1536},{9,0x9A,1600},{9,0x9B,1728},{9,0xCC,704},{9,0xCD,768},{9,0xD2,832},{9,0xD3,896},{9,0xD4,960},{9,0xD5,1024},{9,0xD6,1088},{9,0xD7,1152},{9,0xD8,1216},{9,0xD9,1280},{9,0xDA,1344},{9,0xDB,1408},{10,0x01,_ca },{11,0x01,_ca },{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560}};
_fba =[][3]int {{2,0x02,3},{2,0x03,2},{3,0x02,1},{3,0x03,4},{4,0x02,6},{4,0x03,5},{5,0x03,7},{6,0x04,9},{6,0x05,8},{7,0x04,10},{7,0x05,11},{7,0x07,12},{8,0x04,13},{8,0x07,14},{9,0x01,_ca },{9,0x18,15},{10,0x01,_ca },{10,0x17,16},{10,0x18,17},{10,0x37,0},{10,0x08,18},{10,0x0F,64},{11,0x01,_ca },{11,0x17,24},{11,0x18,25},{11,0x28,23},{11,0x37,22},{11,0x67,19},{11,0x68,20},{11,0x6C,21},{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560},{12,0x24,52},{12,0x27,55},{12,0x28,56},{12,0x2B,59},{12,0x2C,60},{12,0x33,320},{12,0x34,384},{12,0x35,448},{12,0x37,53},{12,0x38,54},{12,0x52,50},{12,0x53,51},{12,0x54,44},{12,0x55,45},{12,0x56,46},{12,0x57,47},{12,0x58,57},{12,0x59,58},{12,0x5A,61},{12,0x5B,256},{12,0x64,48},{12,0x65,49},{12,0x66,62},{12,0x67,63},{12,0x68,30},{12,0x69,31},{12,0x6A,32},{12,0x6B,33},{12,0x6C,40},{12,0x6D,41},{12,0xC8,128},{12,0xC9,192},{12,0xCA,26},{12,0xCB,27},{12,0xCC,28},{12,0xCD,29},{12,0xD2,34},{12,0xD3,35},{12,0xD4,36},{12,0xD5,37},{12,0xD6,38},{12,0xD7,39},{12,0xDA,42},{12,0xDB,43},{13,0x4A,640},{13,0x4B,704},{13,0x4C,768},{13,0x4D,832},{13,0x52,1280},{13,0x53,1344},{13,0x54,1408},{13,0x55,1472},{13,0x5A,1536},{13,0x5B,1600},{13,0x64,1664},{13,0x65,1728},{13,0x6C,512},{13,0x6D,576},{13,0x72,896},{13,0x73,960},{13,0x74,1024},{13,0x75,1088},{13,0x76,1152},{13,0x77,1216}};
);func (_bee *Decoder )fillBitmap (_eeb *_ce .Bitmap ,_dae int ,_eea []int ,_efb int )error {var _bbc byte ;_gga :=0;_bgf :=_eeb .GetByteIndex (_gga ,_dae );for _abe :=0;_abe < _efb ;_abe ++{_cd :=byte (1);_gbe :=_eea [_abe ];if (_abe &1)==0{_cd =0;};for _gga < _gbe {_bbc =(_bbc <<1)|_cd ;
_gga ++;if (_gga &7)==0{if _age :=_eeb .SetByte (_bgf ,_bbc );_age !=nil {return _age ;};_bgf ++;_bbc =0;};};};if (_gga &7)!=0{_bbc <<=uint (8-(_gga &7));if _af :=_eeb .SetByte (_bgf ,_bbc );_af !=nil {return _af ;};};return nil ;};func _ee (_cc [3]int )*code {return &code {_ced :_cc [0],_d :_cc [1],_ceb :_cc [2]}};
func _ac (_cfb ,_ef int )int {if _cfb < _ef {return _ef ;};return _cfb ;};func New (r *_a .Reader ,width ,height int ,dataOffset ,dataLength int64 )(*Decoder ,error ){_bb :=&Decoder {_ff :width ,_bg :height };_bgg ,_cbg :=r .NewPartialReader (int (dataOffset ),int (dataLength ),false );
if _cbg !=nil {return nil ,_cbg ;};_fe ,_cbg :=_gcg (_bgg );if _cbg !=nil {return nil ,_cbg ;};_ ,_cbg =r .Seek (_bgg .RelativePosition (),_cf .SeekCurrent );if _cbg !=nil {return nil ,_cbg ;};_bb ._gb =_fe ;if _dg :=_bb .initTables ();_dg !=nil {return nil ,_dg ;
};return _bb ,nil ;};func (_bgag *runData )fillBuffer (_bcf int )error {_bgag ._eag =_bcf ;_ ,_aeg :=_bgag ._gba .Seek (int64 (_bcf ),_cf .SeekStart );if _aeg !=nil {if _aeg ==_cf .EOF {_e .Log .Debug ("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046");
_bgag ._bca =-1;}else {return _aeg ;};};if _aeg ==nil {_bgag ._bca ,_aeg =_bgag ._gba .Read (_bgag ._gfd );if _aeg !=nil {if _aeg ==_cf .EOF {_e .Log .Trace ("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046");_bgag ._bca =-1;}else {return _aeg ;};};};
if _bgag ._bca > -1&&_bgag ._bca < 3{for _bgag ._bca < 3{_gdf ,_fbgf :=_bgag ._gba .ReadByte ();if _fbgf !=nil {if _fbgf ==_cf .EOF {_bgag ._gfd [_bgag ._bca ]=0;}else {return _fbgf ;};}else {_bgag ._gfd [_bgag ._bca ]=_gdf &0xFF;};_bgag ._bca ++;};};_bgag ._bca -=3;
if _bgag ._bca < 0{_bgag ._gfd =make ([]byte ,len (_bgag ._gfd ));_bgag ._bca =len (_bgag ._gfd )-3;};return nil ;};func _b (_cb ,_bf int )int {if _cb > _bf {return _bf ;};return _cb ;};func (_ba *Decoder )createLittleEndianTable (_cec [][3]int )([]*code ,error ){_aba :=make ([]*code ,_aca +1);
for _cbd :=0;_cbd < len (_cec );_cbd ++{_cge :=_ee (_cec [_cbd ]);if _cge ._ced <=_cfbb {_bcd :=_cfbb -_cge ._ced ;_eee :=_cge ._d <<uint (_bcd );for _ggd :=(1<<uint (_bcd ))-1;_ggd >=0;_ggd --{_cag :=_eee |_ggd ;_aba [_cag ]=_cge ;};}else {_ega :=_cge ._d >>uint (_cge ._ced -_cfbb );
if _aba [_ega ]==nil {var _fd =_ee ([3]int {});_fd ._cg =make ([]*code ,_gdb +1);_aba [_ega ]=_fd ;};if _cge ._ced <=_cfbb +_da {_ea :=_cfbb +_da -_cge ._ced ;_eec :=(_cge ._d <<uint (_ea ))&_gdb ;_aba [_ega ]._dd =true ;for _dbd :=(1<<uint (_ea ))-1;_dbd >=0;
_dbd --{_aba [_ega ]._cg [_eec |_dbd ]=_cge ;};}else {return nil ,_g .New ("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072");
};};};return _aba ,nil ;};type code struct{_ced int ;_d int ;_ceb int ;_cg []*code ;_dd bool ;};const (_ge mmrCode =iota ;_gg ;_eg ;_fge ;_fb ;_be ;_gf ;_bc ;_aa ;_fgg ;_gd ;);const (_ffd int =1024<<7;_cfbg int =3;_fbe uint =24;);func (_ae *Decoder )detectAndSkipEOL ()error {for {_de ,_fbg :=_ae ._gb .uncompressGetCode (_ae ._ga );
if _fbg !=nil {return _fbg ;};if _de !=nil &&_de ._ceb ==EOL {_ae ._gb ._baa +=_de ._ced ;}else {return nil ;};};};func (_gc *Decoder )uncompress2d (_dcd *runData ,_gad []int ,_fea int ,_eda []int ,_dbg int )(int ,error ){var (_cga int ;_eaa int ;_acc int ;
_cae =true ;_dbdb error ;_aef *code ;);_gad [_fea ]=_dbg ;_gad [_fea +1]=_dbg ;_gad [_fea +2]=_dbg +1;_gad [_fea +3]=_dbg +1;_aac :for _acc < _dbg {_aef ,_dbdb =_dcd .uncompressGetCode (_gc ._ga );if _dbdb !=nil {return EOL ,nil ;};if _aef ==nil {_dcd ._baa ++;
break _aac ;};_dcd ._baa +=_aef ._ced ;switch mmrCode (_aef ._ceb ){case _eg :_acc =_gad [_cga ];case _fge :_acc =_gad [_cga ]+1;case _gf :_acc =_gad [_cga ]-1;case _gg :for {var _ggdg []*code ;if _cae {_ggdg =_gc ._bfa ;}else {_ggdg =_gc ._acb ;};_aef ,_dbdb =_dcd .uncompressGetCode (_ggdg );
if _dbdb !=nil {return 0,_dbdb ;};if _aef ==nil {break _aac ;};_dcd ._baa +=_aef ._ced ;if _aef ._ceb < 64{if _aef ._ceb < 0{_eda [_eaa ]=_acc ;_eaa ++;_aef =nil ;break _aac ;};_acc +=_aef ._ceb ;_eda [_eaa ]=_acc ;_eaa ++;break ;};_acc +=_aef ._ceb ;};
_gfc :=_acc ;_fef :for {var _fead []*code ;if !_cae {_fead =_gc ._bfa ;}else {_fead =_gc ._acb ;};_aef ,_dbdb =_dcd .uncompressGetCode (_fead );if _dbdb !=nil {return 0,_dbdb ;};if _aef ==nil {break _aac ;};_dcd ._baa +=_aef ._ced ;if _aef ._ceb < 64{if _aef ._ceb < 0{_eda [_eaa ]=_acc ;
_eaa ++;break _aac ;};_acc +=_aef ._ceb ;if _acc < _dbg ||_acc !=_gfc {_eda [_eaa ]=_acc ;_eaa ++;};break _fef ;};_acc +=_aef ._ceb ;};for _acc < _dbg &&_gad [_cga ]<=_acc {_cga +=2;};continue _aac ;case _ge :_cga ++;_acc =_gad [_cga ];_cga ++;continue _aac ;
case _fb :_acc =_gad [_cga ]+2;case _bc :_acc =_gad [_cga ]-2;case _be :_acc =_gad [_cga ]+3;case _aa :_acc =_gad [_cga ]-3;default:if _dcd ._baa ==12&&_aef ._ceb ==EOL {_dcd ._baa =0;if _ ,_dbdb =_gc .uncompress1d (_dcd ,_gad ,_dbg );_dbdb !=nil {return 0,_dbdb ;
};_dcd ._baa ++;if _ ,_dbdb =_gc .uncompress1d (_dcd ,_eda ,_dbg );_dbdb !=nil {return 0,_dbdb ;};_fgd ,_fgde :=_gc .uncompress1d (_dcd ,_gad ,_dbg );if _fgde !=nil {return EOF ,_fgde ;};_dcd ._baa ++;return _fgd ,nil ;};_acc =_dbg ;continue _aac ;};if _acc <=_dbg {_cae =!_cae ;
_eda [_eaa ]=_acc ;_eaa ++;if _cga > 0{_cga --;}else {_cga ++;};for _acc < _dbg &&_gad [_cga ]<=_acc {_cga +=2;};};};if _eda [_eaa ]!=_dbg {_eda [_eaa ]=_dbg ;};if _aef ==nil {return EOL ,nil ;};return _eaa ,nil ;};func (_cgee *runData )uncompressGetCodeLittleEndian (_gag []*code )(*code ,error ){_ad ,_aaf :=_cgee .uncompressGetNextCodeLittleEndian ();
if _aaf !=nil {_e .Log .Debug ("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076",_aaf );
return nil ,_aaf ;};_ad &=0xffffff;_dca :=_ad >>(_fbe -_cfbb );_cfba :=_gag [_dca ];if _cfba !=nil &&_cfba ._dd {_dca =(_ad >>(_fbe -_cfbb -_da ))&_gdb ;_cfba =_cfba ._cg [_dca ];};return _cfba ,nil ;};func (_cebg *runData )align (){_cebg ._baa =((_cebg ._baa +7)>>3)<<3};
func _gcg (_acg *_a .Reader )(*runData ,error ){_deb :=&runData {_gba :_acg ,_baa :0,_eeef :1};_daf :=_b (_ac (_cfbg ,int (_acg .Length ())),_ffd );_deb ._gfd =make ([]byte ,_daf );if _fee :=_deb .fillBuffer (0);_fee !=nil {if _fee ==_cf .EOF {_deb ._gfd =make ([]byte ,10);
_e .Log .Debug ("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076",_fee );}else {return nil ,_fee ;};};return _deb ,nil ;};func (_bbe *Decoder )uncompress1d (_ccc *runData ,_edd []int ,_gef int )(int ,error ){var (_bd =true ;
_ede int ;_cgf *code ;_ddc int ;_bfe error ;);_bbb :for _ede < _gef {_feb :for {if _bd {_cgf ,_bfe =_ccc .uncompressGetCode (_bbe ._bfa );if _bfe !=nil {return 0,_bfe ;};}else {_cgf ,_bfe =_ccc .uncompressGetCode (_bbe ._acb );if _bfe !=nil {return 0,_bfe ;
};};_ccc ._baa +=_cgf ._ced ;if _cgf ._ceb < 0{break _bbb ;};_ede +=_cgf ._ceb ;if _cgf ._ceb < 64{_bd =!_bd ;_edd [_ddc ]=_ede ;_ddc ++;break _feb ;};};};if _edd [_ddc ]!=_gef {_edd [_ddc ]=_gef ;};_bga :=EOL ;if _cgf !=nil &&_cgf ._ceb !=EOL {_bga =_ddc ;
};return _bga ,nil ;};func (_aff *runData )uncompressGetCode (_acae []*code )(*code ,error ){return _aff .uncompressGetCodeLittleEndian (_acae );};const (EOF =-3;_ca =-2;EOL =-1;_cfbb =8;_aca =(1<<_cfbb )-1;_da =5;_gdb =(1<<_da )-1;);func (_fg *code )String ()string {return _f .Sprintf ("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_fg ._ced ,_fg ._d ,_fg ._ceb );
};func (_ead *runData )uncompressGetNextCodeLittleEndian ()(int ,error ){_bbed :=_ead ._baa -_ead ._eeef ;if _bbed < 0||_bbed > 24{_cdc :=(_ead ._baa >>3)-_ead ._eag ;if _cdc >=_ead ._bca {_cdc +=_ead ._eag ;if _eaf :=_ead .fillBuffer (_cdc );_eaf !=nil {return 0,_eaf ;
};_cdc -=_ead ._eag ;};_gcc :=(uint32 (_ead ._gfd [_cdc ]&0xFF)<<16)|(uint32 (_ead ._gfd [_cdc +1]&0xFF)<<8)|(uint32 (_ead ._gfd [_cdc +2]&0xFF));_edg :=uint32 (_ead ._baa &7);_gcc <<=_edg ;_ead ._ecf =int (_gcc );}else {_adg :=_ead ._eeef &7;_cebf :=7-_adg ;
if _bbed <=_cebf {_ead ._ecf <<=uint (_bbed );}else {_eagf :=(_ead ._eeef >>3)+3-_ead ._eag ;if _eagf >=_ead ._bca {_eagf +=_ead ._eag ;if _fbad :=_ead .fillBuffer (_eagf );_fbad !=nil {return 0,_fbad ;};_eagf -=_ead ._eag ;};_adg =8-_adg ;for {_ead ._ecf <<=uint (_adg );
_ead ._ecf |=int (uint (_ead ._gfd [_eagf ])&0xFF);_bbed -=_adg ;_eagf ++;_adg =8;if !(_bbed >=8){break ;};};_ead ._ecf <<=uint (_bbed );};};_ead ._eeef =_ead ._baa ;return _ead ._ecf ,nil ;};func (_ffea *Decoder )initTables ()(_gaa error ){if _ffea ._bfa ==nil {_ffea ._bfa ,_gaa =_ffea .createLittleEndianTable (_ag );
if _gaa !=nil {return ;};_ffea ._acb ,_gaa =_ffea .createLittleEndianTable (_fba );if _gaa !=nil {return ;};_ffea ._ga ,_gaa =_ffea .createLittleEndianTable (_ab );if _gaa !=nil {return ;};};return nil ;};type runData struct{_gba *_a .Reader ;_baa int ;
_eeef int ;_ecf int ;_gfd []byte ;_eag int ;_bca int ;};type mmrCode int ;