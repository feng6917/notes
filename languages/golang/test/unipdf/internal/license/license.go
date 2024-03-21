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

package license

import (
	_fb "bytes"
	_a "compress/gzip"
	_g "crypto"
	_cd "crypto/aes"
	_de "crypto/cipher"
	_fdad "crypto/hmac"
	_ee "crypto/rand"
	_bf "crypto/rsa"
	_eed "crypto/sha256"
	_cf "crypto/sha512"
	_fe "crypto/x509"
	_ba "encoding/base64"
	_fa "encoding/hex"
	_fda "encoding/json"
	_fd "encoding/pem"
	_eg "errors"
	_gg "fmt"
	_cdd "lgo/test/unipdf/common"
	_d "io"
	_ccb "io/ioutil"
	_ae "net"
	_cc "net/http"
	_b "os"
	_f "path/filepath"
	_dd "sort"
	_db "strings"
	_e "sync"
	_gd "time"
)

func Track(docKey string, useKey string) error { return _gca(docKey, useKey, !_bag._ff) }
func _ded() *meteredClient {
	_ccd := meteredClient{_fce: "h\u0074\u0074\u0070\u0073\u003a\u002f/\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069d\u006f\u0063\u002ei\u006f/\u0061\u0070\u0069", _ab: &_cc.Client{Timeout: 30 * _gd.Second}}
	if _cdf := _b.Getenv("\u0055N\u0049\u0044\u004f\u0043_\u004c\u0049\u0043\u0045\u004eS\u0045_\u0053E\u0052\u0056\u0045\u0052\u005f\u0055\u0052L"); _db.HasPrefix(_cdf, "\u0068\u0074\u0074\u0070") {
		_ccd._fce = _cdf
	}
	return &_ccd
}
func (_af *LicenseKey) isExpired() bool { return _af.getExpiryDateToCompare().After(*_af.ExpiresAt) }
func GetLicenseKey() *LicenseKey {
	if _bag == nil {
		return nil
	}
	_bcfe := *_bag
	return &_bcfe
}

type meteredClient struct {
	_fce string
	_dgg string
	_ab  *_cc.Client
}

func _gbb() string {
	_agf := _b.Getenv("\u0048\u004f\u004d\u0045")
	if len(_agf) == 0 {
		_agf, _ = _b.UserHomeDir()
	}
	return _agf
}

type meteredStatusForm struct{}

var _gee map[string]int

func (_fag *meteredClient) checkinUsage(_aef meteredUsageCheckinForm) (meteredUsageCheckinResp, error) {
	_aef.Package = "\u0075\u006e\u0069\u0070\u0064\u0066"
	_aef.PackageVersion = _cdd.Version
	var _aefa meteredUsageCheckinResp
	_bfe := _fag._fce + "\u002f\u006d\u0065\u0074er\u0065\u0064\u002f\u0075\u0073\u0061\u0067\u0065\u005f\u0063\u0068\u0065\u0063\u006bi\u006e"
	_dga, _gddb := _fda.Marshal(_aef)
	if _gddb != nil {
		return _aefa, _gddb
	}
	_bec, _gddb := _fac(_dga)
	if _gddb != nil {
		return _aefa, _gddb
	}
	_gda, _gddb := _cc.NewRequest("\u0050\u004f\u0053\u0054", _bfe, _bec)
	if _gddb != nil {
		return _aefa, _gddb
	}
	_gda.Header.Add("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e")
	_gda.Header.Add("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_gda.Header.Add("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_gda.Header.Add("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y", _fag._dgg)
	_egg, _gddb := _fag._ab.Do(_gda)
	if _gddb != nil {
		return _aefa, _gddb
	}
	defer _egg.Body.Close()
	if _egg.StatusCode != 200 {
		_gbd, _fcd := _bab(_egg)
		if _fcd != nil {
			return _aefa, _fcd
		}
		_fcd = _fda.Unmarshal(_gbd, &_aefa)
		if _fcd != nil {
			return _aefa, _fcd
		}
		return _aefa, _gg.Errorf("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064", _egg.StatusCode)
	}
	_cde := _egg.Header.Get("\u0058\u002d\u0055\u0043\u002d\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065")
	_ggg := _abd(_aef.MacAddress, string(_dga))
	if _ggg != _cde {
		_cdd.Log.Error("I\u006e\u0076\u0061l\u0069\u0064\u0020\u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u002c\u0020\u0073\u0065t\u0020\u0074\u0068e\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0073\u0065\u0072\u0076e\u0072\u0020\u0074\u006f \u0068\u0074\u0074\u0070s\u003a\u002f\u002f\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069o\u002f\u0061\u0070\u0069")
		return _aefa, _eg.New("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0065\u0072\u0076\u0065\u0072 \u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065")
	}
	_bfec, _gddb := _bab(_egg)
	if _gddb != nil {
		return _aefa, _gddb
	}
	_gddb = _fda.Unmarshal(_bfec, &_aefa)
	if _gddb != nil {
		return _aefa, _gddb
	}
	return _aefa, nil
}
func SetMeteredKey(apiKey string) error {
	if len(apiKey) == 0 {
		_cdd.Log.Error("\u004d\u0065\u0074\u0065\u0072e\u0064\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0041\u0050\u0049 \u004b\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079")
		_cdd.Log.Error("\u002d\u0020\u0047\u0072\u0061\u0062\u0020\u006f\u006e\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0072\u0065\u0065\u0020\u0054\u0069\u0065\u0072\u0020\u0061t\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063\u006c\u006fud\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _gg.Errorf("\u006de\u0074\u0065\u0072e\u0064\u0020\u006ci\u0063en\u0073\u0065\u0020\u0061\u0070\u0069\u0020k\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u003a\u0020\u0063\u0072\u0065\u0061\u0074\u0065 o\u006ee\u0020\u0061\u0074\u0020\u0068\u0074t\u0070\u0073\u003a\u002f\u002fc\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063.\u0069\u006f")
	}
	if _bag != nil && (_bag._bge || _bag.Tier != LicenseTierUnlicensed) {
		_cdd.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0043\u0061\u006e\u006eo\u0074 \u0073\u0065\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0074\u0077\u0069c\u0065\u0020\u002d\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u006a\u0075\u0073\u0074\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069z\u0065\u0020\u006f\u006e\u0063\u0065")
		return _eg.New("\u006c\u0069\u0063en\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0061\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0073\u0065\u0074")
	}
	_fbc := _ded()
	_fbc._dgg = apiKey
	_gbg, _dgc := _fbc.getStatus()
	if _dgc != nil {
		return _dgc
	}
	if !_gbg.Valid {
		return _eg.New("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064")
	}
	_fef := &LicenseKey{_bge: true, _ebe: apiKey, _ff: true}
	_bag = _fef
	return nil
}
func GetMeteredState() (MeteredStatus, error) {
	if _bag == nil {
		return MeteredStatus{}, _eg.New("\u006c\u0069\u0063\u0065ns\u0065\u0020\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	if !_bag._bge || len(_bag._ebe) == 0 {
		return MeteredStatus{}, _eg.New("\u0061p\u0069 \u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_adaa, _eggb := _fdc.loadState(_bag._ebe)
	if _eggb != nil {
		_cdd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _eggb)
		return MeteredStatus{}, _eggb
	}
	if _adaa.Docs > 0 {
		_ebg := _gca("", "", true)
		if _ebg != nil {
			return MeteredStatus{}, _ebg
		}
	}
	_aga.Lock()
	defer _aga.Unlock()
	_ddb := _ded()
	_ddb._dgg = _bag._ebe
	_dfg, _eggb := _ddb.getStatus()
	if _eggb != nil {
		return MeteredStatus{}, _eggb
	}
	if !_dfg.Valid {
		return MeteredStatus{}, _eg.New("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064")
	}
	_fde := MeteredStatus{OK: true, Credits: _dfg.OrgCredits, Used: _dfg.OrgUsed}
	return _fde, nil
}
func _aea(_bdbg, _egge []byte) ([]byte, error) {
	_edfd, _abcf := _cd.NewCipher(_bdbg)
	if _abcf != nil {
		return nil, _abcf
	}
	_adgc := make([]byte, _cd.BlockSize+len(_egge))
	_dcc := _adgc[:_cd.BlockSize]
	if _, _bgff := _d.ReadFull(_ee.Reader, _dcc); _bgff != nil {
		return nil, _bgff
	}
	_eaa := _de.NewCFBEncrypter(_edfd, _dcc)
	_eaa.XORKeyStream(_adgc[_cd.BlockSize:], _egge)
	_gdac := make([]byte, _ba.URLEncoding.EncodedLen(len(_adgc)))
	_ba.URLEncoding.Encode(_gdac, _adgc)
	return _gdac, nil
}

type reportState struct {
	Instance      string         `json:"inst"`
	Next          string         `json:"n"`
	Docs          int64          `json:"d"`
	NumErrors     int64          `json:"e"`
	LimitDocs     bool           `json:"ld"`
	RemainingDocs int64          `json:"rd"`
	LastReported  _gd.Time       `json:"lr"`
	LastWritten   _gd.Time       `json:"lw"`
	Usage         map[string]int `json:"u"`
}
type meteredUsageCheckinForm struct {
	Instance          string         `json:"inst"`
	Next              string         `json:"next"`
	UsageNumber       int            `json:"usage_number"`
	NumFailed         int64          `json:"num_failed"`
	Hostname          string         `json:"hostname"`
	LocalIP           string         `json:"local_ip"`
	MacAddress        string         `json:"mac_address"`
	Package           string         `json:"package"`
	PackageVersion    string         `json:"package_version"`
	Usage             map[string]int `json:"u"`
	IsPersistentCache bool           `json:"is_persistent_cache"`
	Timestamp         int64          `json:"timestamp"`
}

func (_gddg *LicenseKey) IsLicensed() bool { return _gddg.Tier != LicenseTierUnlicensed || _gddg._bge }

var _eea = _gd.Date(2019, 6, 6, 0, 0, 0, 0, _gd.UTC)
var _da = _gd.Date(2020, 1, 1, 0, 0, 0, 0, _gd.UTC)

type meteredStatusResp struct {
	Valid        bool  `json:"valid"`
	OrgCredits   int64 `json:"org_credits"`
	OrgUsed      int64 `json:"org_used"`
	OrgRemaining int64 `json:"org_remaining"`
}
type MeteredStatus struct {
	OK      bool
	Credits int64
	Used    int64
}

func (_bbf *LicenseKey) TypeToString() string {
	if _bbf._bge {
		return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e"
	}
	if _bbf.Tier == LicenseTierUnlicensed {
		return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	}
	if _bbf.Tier == LicenseTierCommunity {
		return "\u0041\u0047PL\u0076\u0033\u0020O\u0070\u0065\u006e\u0020Sou\u0072ce\u0020\u0043\u006f\u006d\u006d\u0075\u006eit\u0079\u0020\u004c\u0069\u0063\u0065\u006es\u0065"
	}
	if _bbf.Tier == LicenseTierIndividual || _bbf.Tier == "\u0069\u006e\u0064i\u0065" {
		return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	}
	return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073"
}
func _cee() (_ae.IP, error) {
	_bbec, _cbaf := _ae.Dial("\u0075\u0064\u0070", "\u0038\u002e\u0038\u002e\u0038\u002e\u0038\u003a\u0038\u0030")
	if _cbaf != nil {
		return nil, _cbaf
	}
	defer _bbec.Close()
	_edf := _bbec.LocalAddr().(*_ae.UDPAddr)
	return _edf.IP, nil
}

type defaultStateHolder struct{}

var _aga = &_e.Mutex{}

func _cab() ([]string, []string, error) {
	_bgf, _agb := _ae.Interfaces()
	if _agb != nil {
		return nil, nil, _agb
	}
	var _eag []string
	var _ggb []string
	for _, _fff := range _bgf {
		if _fff.Flags&_ae.FlagUp == 0 || _fb.Equal(_fff.HardwareAddr, nil) {
			continue
		}
		_adg, _ffa := _fff.Addrs()
		if _ffa != nil {
			return nil, nil, _ffa
		}
		_dc := 0
		for _, _fad := range _adg {
			var _fbbc _ae.IP
			switch _dda := _fad.(type) {
			case *_ae.IPNet:
				_fbbc = _dda.IP
			case *_ae.IPAddr:
				_fbbc = _dda.IP
			}
			if _fbbc.IsLoopback() {
				continue
			}
			if _fbbc.To4() == nil {
				continue
			}
			_ggb = append(_ggb, _fbbc.String())
			_dc++
		}
		_gfb := _fff.HardwareAddr.String()
		if _gfb != "" && _dc > 0 {
			_eag = append(_eag, _gfb)
		}
	}
	return _eag, _ggb, nil
}

const _ddag = "\u0055\u004e\u0049\u0050DF\u005f\u004c\u0049\u0043\u0045\u004e\u0053\u0045\u005f\u0050\u0041\u0054\u0048"

func _agcd(_gddd, _cbc []byte) ([]byte, error) {
	_faa := make([]byte, _ba.URLEncoding.DecodedLen(len(_cbc)))
	_eacf, _ccbcc := _ba.URLEncoding.Decode(_faa, _cbc)
	if _ccbcc != nil {
		return nil, _ccbcc
	}
	_faa = _faa[:_eacf]
	_afb, _ccbcc := _cd.NewCipher(_gddd)
	if _ccbcc != nil {
		return nil, _ccbcc
	}
	if len(_faa) < _cd.BlockSize {
		return nil, _eg.New("c\u0069p\u0068\u0065\u0072\u0074\u0065\u0078\u0074\u0020t\u006f\u006f\u0020\u0073ho\u0072\u0074")
	}
	_dbf := _faa[:_cd.BlockSize]
	_faa = _faa[_cd.BlockSize:]
	_gbed := _de.NewCFBDecrypter(_afb, _dbf)
	_gbed.XORKeyStream(_faa, _faa)
	return _faa, nil
}

const _faga = "U\u004eI\u0050\u0044\u0046\u005f\u0043\u0055\u0053\u0054O\u004d\u0045\u0052\u005fNA\u004d\u0045"

func (_ddf *meteredClient) getStatus() (meteredStatusResp, error) {
	var _dgga meteredStatusResp
	_bd := _ddf._fce + "\u002fm\u0065t\u0065\u0072\u0065\u0064\u002f\u0073\u0074\u0061\u0074\u0075\u0073"
	var _gcd meteredStatusForm
	_bfb, _beeb := _fda.Marshal(_gcd)
	if _beeb != nil {
		return _dgga, _beeb
	}
	_fabde, _beeb := _fac(_bfb)
	if _beeb != nil {
		return _dgga, _beeb
	}
	_gdb, _beeb := _cc.NewRequest("\u0050\u004f\u0053\u0054", _bd, _fabde)
	if _beeb != nil {
		return _dgga, _beeb
	}
	_gdb.Header.Add("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e")
	_gdb.Header.Add("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_gdb.Header.Add("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_gdb.Header.Add("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y", _ddf._dgg)
	_bcc, _beeb := _ddf._ab.Do(_gdb)
	if _beeb != nil {
		return _dgga, _beeb
	}
	defer _bcc.Body.Close()
	if _bcc.StatusCode != 200 {
		return _dgga, _gg.Errorf("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064", _bcc.StatusCode)
	}
	_cbe, _beeb := _bab(_bcc)
	if _beeb != nil {
		return _dgga, _beeb
	}
	_beeb = _fda.Unmarshal(_cbe, &_dgga)
	if _beeb != nil {
		return _dgga, _beeb
	}
	return _dgga, nil
}

const _efd = "\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n"

type meteredUsageCheckinResp struct {
	Instance      string `json:"inst"`
	Next          string `json:"next"`
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	RemainingDocs int    `json:"rd"`
	LimitDocs     bool   `json:"ld"`
}

func _bab(_bgg *_cc.Response) ([]byte, error) {
	var _agag []byte
	_ccbc, _agba := _beef(_bgg)
	if _agba != nil {
		return _agag, _agba
	}
	return _ccb.ReadAll(_ccbc)
}
func _adad(_bg string) (LicenseKey, error) {
	var _ec LicenseKey
	_gdg, _fcc := _cbd(_ga, _cca, _bg)
	if _fcc != nil {
		return _ec, _fcc
	}
	_aeg, _fcc := _fdg(_efd, _gdg)
	if _fcc != nil {
		return _ec, _fcc
	}
	_fcc = _fda.Unmarshal(_aeg, &_ec)
	if _fcc != nil {
		return _ec, _fcc
	}
	_ec.CreatedAt = _gd.Unix(_ec.CreatedAtInt, 0)
	if _ec.ExpiresAtInt > 0 {
		_df := _gd.Unix(_ec.ExpiresAtInt, 0)
		_ec.ExpiresAt = &_df
	}
	return _ec, nil
}
func _beef(_cced *_cc.Response) (_d.ReadCloser, error) {
	var _aege error
	var _ffag _d.ReadCloser
	switch _db.ToLower(_cced.Header.Get("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")) {
	case "\u0067\u007a\u0069\u0070":
		_ffag, _aege = _a.NewReader(_cced.Body)
		if _aege != nil {
			return _ffag, _aege
		}
		defer _ffag.Close()
	default:
		_ffag = _cced.Body
	}
	return _ffag, nil
}
func _abd(_aefd, _afg string) string {
	_ffc := []byte(_aefd)
	_cdg := _fdad.New(_eed.New, _ffc)
	_cdg.Write([]byte(_afg))
	return _ba.StdEncoding.EncodeToString(_cdg.Sum(nil))
}

var _eeb = _gd.Date(2010, 1, 1, 0, 0, 0, 0, _gd.UTC)

func (_ecg *LicenseKey) Validate() error {
	if _ecg._bge {
		return nil
	}
	if len(_ecg.LicenseId) < 10 {
		return _gg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064")
	}
	if len(_ecg.CustomerId) < 10 {
		return _gg.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064")
	}
	if len(_ecg.CustomerName) < 1 {
		return _gg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065")
	}
	if _eeb.After(_ecg.CreatedAt) {
		return _gg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _ecg.ExpiresAt == nil {
		_bee := _ecg.CreatedAt.AddDate(1, 0, 0)
		if _da.After(_bee) {
			_bee = _da
		}
		_ecg.ExpiresAt = &_bee
	}
	if _ecg.CreatedAt.After(*_ecg.ExpiresAt) {
		return _gg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074")
	}
	if _ecg.isExpired() {
		return _gg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020l\u0069\u0063\u0065ns\u0065\u003a\u0020\u0054\u0068\u0065 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064")
	}
	if len(_ecg.CreatorName) < 1 {
		return _gg.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065")
	}
	if len(_ecg.CreatorEmail) < 1 {
		return _gg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c")
	}
	if _ecg.CreatedAt.After(_eea) {
		if !_ecg.UniPDF {
			return _gg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020\u0054\u0068\u0069\u0073\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020k\u0065\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d \u0066\u006f\u0072\u0020\u0055\u006e\u0069\u0050\u0044\u0046")
		}
	}
	return nil
}

const (
	LicenseTierUnlicensed = "\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	LicenseTierCommunity  = "\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y"
	LicenseTierIndividual = "\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	LicenseTierBusiness   = "\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073"
)

func MakeUnlicensedKey() *LicenseKey {
	_ef := LicenseKey{}
	_ef.CustomerName = "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	_ef.Tier = LicenseTierUnlicensed
	_ef.CreatedAt = _gd.Now().UTC()
	_ef.CreatedAtInt = _ef.CreatedAt.Unix()
	return &_ef
}
func (_bfbc defaultStateHolder) loadState(_edb string) (reportState, error) {
	_ce := _gbb()
	if len(_ce) == 0 {
		return reportState{}, _eg.New("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_daf := _f.Join(_ce, "\u002eu\u006e\u0069\u0064\u006f\u0063")
	_gad := _b.MkdirAll(_daf, 0777)
	if _gad != nil {
		return reportState{}, _gad
	}
	if len(_edb) < 20 {
		return reportState{}, _eg.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079")
	}
	_bbb := []byte(_edb)
	_eca := _cf.Sum512_256(_bbb[:20])
	_dad := _fa.EncodeToString(_eca[:])
	_agc := _f.Join(_daf, _dad)
	_gga, _gad := _ccb.ReadFile(_agc)
	if _gad != nil {
		if _b.IsNotExist(_gad) {
			return reportState{}, nil
		}
		_cdd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gad)
		return reportState{}, _eg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061")
	}
	const _cea = "\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057"
	_gga, _gad = _agcd([]byte(_cea), _gga)
	if _gad != nil {
		return reportState{}, _gad
	}
	var _adag reportState
	_gad = _fda.Unmarshal(_gga, &_adag)
	if _gad != nil {
		_cdd.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0076", _gad)
		return reportState{}, _eg.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061")
	}
	return _adag, nil
}
func _fdg(_bcd string, _bcf string) ([]byte, error) {
	var (
		_ge  int
		_gge string
	)
	for _, _gge = range []string{"\u000a\u002b\u000a", "\u000d\u000a\u002b\r\u000a", "\u0020\u002b\u0020"} {
		if _ge = _db.Index(_bcf, _gge); _ge != -1 {
			break
		}
	}
	if _ge == -1 {
		return nil, _gg.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072")
	}
	_fc := _bcf[:_ge]
	_bb := _ge + len(_gge)
	_ebf := _bcf[_bb:]
	if _fc == "" || _ebf == "" {
		return nil, _gg.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065")
	}
	_gdd, _ada := _ba.StdEncoding.DecodeString(_fc)
	if _ada != nil {
		return nil, _gg.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c")
	}
	_cg, _ada := _ba.StdEncoding.DecodeString(_ebf)
	if _ada != nil {
		return nil, _gg.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_fab, _ := _fd.Decode([]byte(_bcd))
	if _fab == nil {
		return nil, _gg.Errorf("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_fbd, _ada := _fe.ParsePKIXPublicKey(_fab.Bytes)
	if _ada != nil {
		return nil, _ada
	}
	_fbb := _fbd.(*_bf.PublicKey)
	if _fbb == nil {
		return nil, _gg.Errorf("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_ca := _cf.New()
	_ca.Write(_gdd)
	_bce := _ca.Sum(nil)
	_ada = _bf.VerifyPKCS1v15(_fbb, _g.SHA512, _bce, _cg)
	if _ada != nil {
		return nil, _ada
	}
	return _gdd, nil
}
func _cbd(_egf string, _bbg string, _dbe string) (string, error) {
	_gef := _db.Index(_dbe, _egf)
	if _gef == -1 {
		return "", _gg.Errorf("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_fabe := _db.Index(_dbe, _bbg)
	if _fabe == -1 {
		return "", _gg.Errorf("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ag := _gef + len(_egf) + 1
	return _dbe[_ag : _fabe-1], nil
}
func SetMeteredKeyPersistentCache(val bool) { _bag._ff = val }
func _fac(_efa []byte) (_d.Reader, error) {
	_adcd := new(_fb.Buffer)
	_gbeg := _a.NewWriter(_adcd)
	_gbeg.Write(_efa)
	_afa := _gbeg.Close()
	if _afa != nil {
		return nil, _afa
	}
	return _adcd, nil
}

var _fdc stateLoader = defaultStateHolder{}

const (
	_ga  = "\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
	_cca = "\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
)

func (_gcc *LicenseKey) getExpiryDateToCompare() _gd.Time {
	if _gcc.Trial {
		return _gd.Now().UTC()
	}
	return _cdd.ReleasedAt
}

var _bcg map[string]struct{}
var _bag = MakeUnlicensedKey()

func init() {
	_ebb := _b.Getenv(_ddag)
	_baagc := _b.Getenv(_faga)
	if len(_ebb) == 0 || len(_baagc) == 0 {
		return
	}
	_ccbg, _beg := _ccb.ReadFile(_ebb)
	if _beg != nil {
		_cdd.Log.Error("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0072\u0065ad \u006cic\u0065\u006e\u0073\u0065\u0020\u0063\u006fde\u0020\u0066\u0069\u006c\u0065\u003a\u0020%\u0076", _beg)
		return
	}
	_beg = SetLicenseKey(string(_ccbg), _baagc)
	if _beg != nil {
		_cdd.Log.Error("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u006c\u006f\u0061\u0064\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _beg)
		return
	}
}
func _gca(_bcdf string, _gdda string, _ggge bool) error {
	if _bag == nil {
		return _eg.New("\u006e\u006f\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u0020\u006b\u0065\u0079")
	}
	if !_bag._bge || len(_bag._ebe) == 0 {
		return nil
	}
	if len(_bcdf) == 0 && !_ggge {
		return _eg.New("\u0064\u006f\u0063\u004b\u0065\u0079\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	_aga.Lock()
	defer _aga.Unlock()
	if _bcg == nil {
		_bcg = map[string]struct{}{}
	}
	if _gee == nil {
		_gee = map[string]int{}
	}
	_fbba := 0
	_, _bdb := _bcg[_bcdf]
	if !_bdb {
		_bcg[_bcdf] = struct{}{}
		_fbba++
	}
	if _fbba == 0 {
		return nil
	}
	_gee[_gdda]++
	_fabf := _gd.Now()
	_cce, _bdg := _fdc.loadState(_bag._ebe)
	if _bdg != nil {
		_cdd.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bdg)
		return _bdg
	}
	if _cce.Usage == nil {
		_cce.Usage = map[string]int{}
	}
	for _ebag, _acg := range _gee {
		_cce.Usage[_ebag] += _acg
	}
	_gee = nil
	const _egd = 24 * _gd.Hour
	const _ebfc = 3 * 24 * _gd.Hour
	if len(_cce.Instance) == 0 || _fabf.Sub(_cce.LastReported) > _egd || (_cce.LimitDocs && _cce.RemainingDocs <= _cce.Docs+int64(_fbba)) || _ggge {
		_dee, _efe := _b.Hostname()
		if _efe != nil {
			return _efe
		}
		_edc := _cce.Docs
		_bbd, _fgg, _efe := _cab()
		if _efe != nil {
			_cdd.Log.Debug("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u006c\u006fc\u0061\u006c\u0020\u0061\u0064\u0064\u0072\u0065\u0073\u0073\u003a\u0020\u0025\u0073", _efe.Error())
			_bbd = append(_bbd, "\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce")
			_fgg = append(_fgg, "\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce")
		} else {
			_dd.Strings(_fgg)
			_dd.Strings(_bbd)
			_ebaf, _eegg := _cee()
			if _eegg != nil {
				return _eegg
			}
			_gce := false
			for _, _cggd := range _fgg {
				if _cggd == _ebaf.String() {
					_gce = true
				}
			}
			if !_gce {
				_fgg = append(_fgg, _ebaf.String())
			}
		}
		_gdag := _ded()
		_gdag._dgg = _bag._ebe
		_edc += int64(_fbba)
		_cba := meteredUsageCheckinForm{Instance: _cce.Instance, Next: _cce.Next, UsageNumber: int(_edc), NumFailed: _cce.NumErrors, Hostname: _dee, LocalIP: _db.Join(_fgg, "\u002c\u0020"), MacAddress: _db.Join(_bbd, "\u002c\u0020"), Package: "\u0075\u006e\u0069\u0070\u0064\u0066", PackageVersion: _cdd.Version, Usage: _cce.Usage, IsPersistentCache: _bag._ff, Timestamp: _fabf.Unix()}
		if len(_bbd) == 0 {
			_cba.MacAddress = "\u006e\u006f\u006e\u0065"
		}
		_bfd := int64(0)
		_aad := _cce.NumErrors
		_ccbd := _fabf
		_eaca := 0
		_ccc := _cce.LimitDocs
		_cbg, _efe := _gdag.checkinUsage(_cba)
		if _efe != nil {
			if _fabf.Sub(_cce.LastReported) > _ebfc {
				if !_cbg.Success {
					return _eg.New(_cbg.Message)
				}
				return _eg.New("\u0074\u006f\u006f\u0020\u006c\u006f\u006e\u0067\u0020\u0073\u0069\u006e\u0063\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0073\u0075\u0063\u0063e\u0073\u0073\u0066\u0075\u006c \u0063\u0068e\u0063\u006b\u0069\u006e")
			}
			_bfd = _edc
			_aad++
			_ccbd = _cce.LastReported
		} else {
			_ccc = _cbg.LimitDocs
			_eaca = _cbg.RemainingDocs
			_aad = 0
		}
		if len(_cbg.Instance) == 0 {
			_cbg.Instance = _cba.Instance
		}
		if len(_cbg.Next) == 0 {
			_cbg.Next = _cba.Next
		}
		_efe = _fdc.updateState(_gdag._dgg, _cbg.Instance, _cbg.Next, int(_bfd), _ccc, _eaca, int(_aad), _ccbd, nil)
		if _efe != nil {
			return _efe
		}
		if !_cbg.Success {
			return _gg.Errorf("\u0065r\u0072\u006f\u0072\u003a\u0020\u0025s", _cbg.Message)
		}
	} else {
		_bdg = _fdc.updateState(_bag._ebe, _cce.Instance, _cce.Next, int(_cce.Docs)+_fbba, _cce.LimitDocs, int(_cce.RemainingDocs), int(_cce.NumErrors), _cce.LastReported, _cce.Usage)
		if _bdg != nil {
			return _bdg
		}
	}
	return nil
}
func SetLicenseKey(content string, customerName string) error {
	_cae, _edd := _adad(content)
	if _edd != nil {
		_cdd.Log.Error("\u004c\u0069c\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u0072\u0072\u006f\u0072: \u0025\u0076", _edd)
		return _edd
	}
	if !_db.EqualFold(_cae.CustomerName, customerName) {
		_cdd.Log.Error("L\u0069ce\u006es\u0065 \u0063\u006f\u0064\u0065\u0020i\u0073\u0073\u0075e\u0020\u002d\u0020\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006da\u0074\u0063\u0068, e\u0078\u0070\u0065\u0063\u0074\u0065d\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067o\u0074 \u0027\u0025\u0073\u0027", _cae.CustomerName, customerName)
		return _gg.Errorf("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'", _cae.CustomerName, customerName)
	}
	_edd = _cae.Validate()
	if _edd != nil {
		_cdd.Log.Error("\u004c\u0069\u0063\u0065\u006e\u0073e\u0020\u0063\u006f\u0064\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074i\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _edd)
		return _edd
	}
	_bag = &_cae
	return nil
}
func TrackUse(useKey string) {
	if _bag == nil {
		return
	}
	if !_bag._bge || len(_bag._ebe) == 0 {
		return
	}
	if len(useKey) == 0 {
		return
	}
	_aga.Lock()
	defer _aga.Unlock()
	if _gee == nil {
		_gee = map[string]int{}
	}
	_gee[useKey]++
}

type stateLoader interface {
	loadState(_adc string) (reportState, error)
	updateState(_cfgf, _eba, _bde string, _beb int, _dfgb bool, _geg int, _ea int, _ccdc _gd.Time, _bbe map[string]int) error
}

func (_gde defaultStateHolder) updateState(_bebc, _aa, _dfa string, _fge int, _ac bool, _cdfg int, _bded int, _ffb _gd.Time, _fca map[string]int) error {
	_cgd := _gbb()
	if len(_cgd) == 0 {
		return _eg.New("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_adcb := _f.Join(_cgd, "\u002eu\u006e\u0069\u0064\u006f\u0063")
	_eeg := _b.MkdirAll(_adcb, 0777)
	if _eeg != nil {
		return _eeg
	}
	if len(_bebc) < 20 {
		return _eg.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079")
	}
	_caf := []byte(_bebc)
	_eac := _cf.Sum512_256(_caf[:20])
	_cda := _fa.EncodeToString(_eac[:])
	_dfab := _f.Join(_adcb, _cda)
	var _abc reportState
	_abc.Docs = int64(_fge)
	_abc.NumErrors = int64(_bded)
	_abc.LimitDocs = _ac
	_abc.RemainingDocs = int64(_cdfg)
	_abc.LastWritten = _gd.Now().UTC()
	_abc.LastReported = _ffb
	_abc.Instance = _aa
	_abc.Next = _dfa
	_abc.Usage = _fca
	_gbe, _eeg := _fda.Marshal(_abc)
	if _eeg != nil {
		return _eeg
	}
	const _cgg = "\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057"
	_gbe, _eeg = _aea([]byte(_cgg), _gbe)
	if _eeg != nil {
		return _eeg
	}
	_eeg = _ccb.WriteFile(_dfab, _gbe, 0600)
	if _eeg != nil {
		return _eeg
	}
	return nil
}

type LicenseKey struct {
	LicenseId    string    `json:"license_id"`
	CustomerId   string    `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Tier         string    `json:"tier"`
	CreatedAt    _gd.Time  `json:"-"`
	CreatedAtInt int64     `json:"created_at"`
	ExpiresAt    *_gd.Time `json:"-"`
	ExpiresAtInt int64     `json:"expires_at"`
	CreatedBy    string    `json:"created_by"`
	CreatorName  string    `json:"creator_name"`
	CreatorEmail string    `json:"creator_email"`
	UniPDF       bool      `json:"unipdf"`
	UniOffice    bool      `json:"unioffice"`
	UniHTML      bool      `json:"unihtml"`
	Trial        bool      `json:"trial"`
	_bge         bool
	_ebe         string
	_ff          bool
}

func (_fabd *LicenseKey) ToString() string {
	if _fabd._bge {
		return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e"
	}
	_gf := _gg.Sprintf("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _fabd.LicenseId)
	_gf += _gg.Sprintf("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _fabd.CustomerId)
	_gf += _gg.Sprintf("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a", _fabd.CustomerName)
	_gf += _gg.Sprintf("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n", _fabd.Tier)
	_gf += _gg.Sprintf("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _cdd.UtcTimeFormat(_fabd.CreatedAt))
	if _fabd.ExpiresAt == nil {
		_gf += "\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041t\u003a\u0020N\u0065\u0076\u0065\u0072\u000a"
	} else {
		_gf += _gg.Sprintf("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _cdd.UtcTimeFormat(*_fabd.ExpiresAt))
	}
	_gf += _gg.Sprintf("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a", _fabd.CreatorName, _fabd.CreatorEmail)
	return _gf
}
func _bc(_ad string, _eb []byte) (string, error) {
	_ed, _ := _fd.Decode([]byte(_ad))
	if _ed == nil {
		return "", _gg.Errorf("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	_cb, _gb := _fe.ParsePKCS1PrivateKey(_ed.Bytes)
	if _gb != nil {
		return "", _gb
	}
	_gc := _cf.New()
	_gc.Write(_eb)
	_dg := _gc.Sum(nil)
	_be, _gb := _bf.SignPKCS1v15(_ee.Reader, _cb, _g.SHA512, _dg)
	if _gb != nil {
		return "", _gb
	}
	_gaa := _ba.StdEncoding.EncodeToString(_eb)
	_gaa += "\u000a\u002b\u000a"
	_gaa += _ba.StdEncoding.EncodeToString(_be)
	return _gaa, nil
}