package wxpay

type Account struct {
	appID     string
	subAppId  string
	mchID     string
	subMchId  string
	apiKey    string
	cert      string
	key       string
	isSandbox bool
}

// 创建微信支付账号
func NewAccount(appID, mchID, apiKey, cert, key string, isSanbox bool) *Account {
	return &Account{
		appID:     appID,
		mchID:     mchID,
		apiKey:    apiKey,
		isSandbox: isSanbox,
		cert:      cert,
		key:       key,
	}
}
