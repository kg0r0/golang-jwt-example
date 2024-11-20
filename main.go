package main

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt"
)

const (
	keypair = `{
    "p": "0yCVZVOikdhUA5m6OlPYvYps9vBJx4ilyDKs95jNQGMg_GuD8yvkx1Efas7yBsRfA3Fn6AxiaLsMAZUtfimKAdbMtP_tltV6kH_Zx9E9OXxkRFUI7KsvhzBeHgoz5vTGg-WbDuuJ2WPn8wOM2IGtkULQhN6Gz3sHjCHOxiHcBB8",
    "kty": "RSA",
    "q": "xrT4BAAB5r2SiCQZUgNsuRFAi_sitAJzX_rrQP1OCPmOwdTXy_vcJtFYlhnaf3u9OYT4nNQXX6duIjJUbJgMisVRGj0AB-ijcpsVr1cr_Y_JcSilyALeDES7SaaVUiZ8L1rTbDka-VhpcKVdzPpl8Pkxg7z1IRl9w08ayuw4gbM",
    "d": "a2W_OGFsxEb2fBXEFs9ZKKOx9fwcCJhM_XlJT6EaCgIiTcxruLXOkRl5hIL7JkmlG3rTq_Dciawh7BjhwHZyRTi1tNJMoG541zYbt_w_78RcFUiQsd2G7dph2FcxtY5B-SxM00rFaFNEfiVkmdLpgc-BC3mqm8AOMgAguOVs40YYun1OTz1G_2waA4pKM9hyRXuhWpUhWYXODIkDrYX5-K2lJdS7ZlYTy_JiM_rRUF_Wg6XpZy1SubxS-CrcaWWdJJAGE2bsRc70oS2RJyjIkTkJcfu41deyH_iFI5Dxd5XtyybfCiKLiBbSNya3EsHZx4b7nUgFqkuzPZFdCKyKKQ",
    "e": "AQAB",
    "use": "sig",
    "kid": "1",
    "qi": "zBSeR0YayX7KawDPTw9Qfjk3JOo-pNCE6w1eTSN2e90KvVJiSbp4rVO8NuNJww_htK__aniDMT-rJDDP25UTDjCHoNz2GtBM3XCSGMjvEdSOqSkbxNo-BxxhxDutRReipn5Aa3LWJ2SMfKnFSECcAPNwmBHTRbvVtU0s1jKc4vA",
    "dp": "0Bov5752YbePqDTgwRlgbAODwCu9LXZdomWA5FSzC6IqI2R-nTRIvsYRZ6AwI8dvt98SgkGixoSIIw891jtvkrx87nPNZn1p4ACFU1XFOWKJGmmO8GkT4fck7gs0eZQQEHZDToOQTr0RJhH7xHSd9q6bBjypON2V5OR2Agnh6hU",
    "alg": "RS256",
    "dq": "M1mA3kfCNga3X0c04-TOq-SxcXsstKgNeLg3I0xSZi9XnO-L9MLZWY6v_doghOFNPRgHxz9n6ugxpdSrzIReeV4UX1t0LpcH5g39xJoaXCRUQlHmxZE4IKOCYr4RyHD5lqM6D7WSKu2WEe4qF1Z-EY_UI98o2azkuxwuKFJzJ9E",
    "n": "o-BzAEcsmyLcbzkQ60pHT3UdKA3vgXlmO1TnprQWDFncc5_hIbuwdSiCvFoqzam8ubnsm__P5xDLRodJ58ETAwKIYi4YNzAn5awqkeMUX3CjnIM2aMYIoIKOIu8rBtqyZ2Qcdut9CEPh0ypAfLBmjB-hjpnvLRzGzdzlvpdbi1fG07FkrkxINwd3xZ_Yk1IKXdHpJF9beu23APILuz_2x0S3g1Q1SQWGBKfj9LCpRN-jTLvvLnRJYumaksrjTqXDcl1dAlTRHM_QZQ9zd3tGAh1w5IGylEiJSXX9J5DU7UpKQmcIGeOGWcU84ZT19Yf9jUDEk6NYAdrZ-w39BbKArQ"
	}`

	jwkPub = `{
    "kty": "RSA",
    "e": "AQAB",
    "use": "sig",
    "kid": "1",
    "alg": "RS256",
    "n": "o-BzAEcsmyLcbzkQ60pHT3UdKA3vgXlmO1TnprQWDFncc5_hIbuwdSiCvFoqzam8ubnsm__P5xDLRodJ58ETAwKIYi4YNzAn5awqkeMUX3CjnIM2aMYIoIKOIu8rBtqyZ2Qcdut9CEPh0ypAfLBmjB-hjpnvLRzGzdzlvpdbi1fG07FkrkxINwd3xZ_Yk1IKXdHpJF9beu23APILuz_2x0S3g1Q1SQWGBKfj9LCpRN-jTLvvLnRJYumaksrjTqXDcl1dAlTRHM_QZQ9zd3tGAh1w5IGylEiJSXX9J5DU7UpKQmcIGeOGWcU84ZT19Yf9jUDEk6NYAdrZ-w39BbKArQ"
  }`

	pemPub = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo+BzAEcsmyLcbzkQ60pH
T3UdKA3vgXlmO1TnprQWDFncc5/hIbuwdSiCvFoqzam8ubnsm//P5xDLRodJ58ET
AwKIYi4YNzAn5awqkeMUX3CjnIM2aMYIoIKOIu8rBtqyZ2Qcdut9CEPh0ypAfLBm
jB+hjpnvLRzGzdzlvpdbi1fG07FkrkxINwd3xZ/Yk1IKXdHpJF9beu23APILuz/2
x0S3g1Q1SQWGBKfj9LCpRN+jTLvvLnRJYumaksrjTqXDcl1dAlTRHM/QZQ9zd3tG
Ah1w5IGylEiJSXX9J5DU7UpKQmcIGeOGWcU84ZT19Yf9jUDEk6NYAdrZ+w39BbKA
rQIDAQAB
-----END PUBLIC KEY-----
	`

	jwtStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.lzkUHOacEdJn6iJKscImCXRObu8ceBwF3dj2qQArT8Skj4Fnh6RO91Ji1CIY4eSJIA-6CYPynv-FNfp_3EiH44rZw4KYPWVbSWwTECg1Et41izTqzGy3aUK0DLRqwZXbV0vxWhhS4S5jaD2aAM2t0_UeQtF-sZIbdrjsRbaNBiOp7nX7O1xgb3F0iGcdLUmeg0B4xPaXnqmPxnlatTuAvtvo7HNkYE1d30HeLTiQlmG2LIdwA6l9HKGglU8oqEY2fFzO1QpD19KNyB-WyF_uGs3hFsLstdYhp5jy-voMzHXSjFXXPV1zu0KqXV-MEP4BnczuM42XzSzURMZWXDhLwQ`
)

func main() {
	token, err := jwt.ParseWithClaims(jwtStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(pemPub))
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", token)
}
