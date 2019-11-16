package wxpay

const (
	Fail                       = "FAIL"
	Success                    = "SUCCESS"
	HMACSHA256                 = "HMAC-SHA256"
	MD5                        = "MD5"
	Sign                       = "sign"
	MicroPayUrl                = "https://api.mch.weixin.qq.com/pay/micropay"
	UnifiedOrderUrl            = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryUrl              = "https://api.mch.weixin.qq.com/pay/orderquery"
	ReverseUrl                 = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	CloseOrderUrl              = "https://api.mch.weixin.qq.com/pay/closeorder"
	RefundUrl                  = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	RefundQueryUrl             = "https://api.mch.weixin.qq.com/pay/refundquery"
	DownloadBillUrl            = "https://api.mch.weixin.qq.com/pay/downloadbill"
	DownloadFundFlowUrl        = "https://api.mch.weixin.qq.com/pay/downloadfundflow"
	ReportUrl                  = "https://api.mch.weixin.qq.com/payitil/report"
	ShortUrl                   = "https://api.mch.weixin.qq.com/tools/shorturl"
	AuthCodeToOpenidUrl        = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	SandboxMicroPayUrl         = "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay"
	SandboxUnifiedOrderUrl     = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	SandboxOrderQueryUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
	SandboxReverseUrl          = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/reverse"
	SandboxCloseOrderUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/closeorder"
	SandboxRefundUrl           = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
	SandboxRefundQueryUrl      = "https://api.mch.weixin.qq.com/sandboxnew/pay/refundquery"
	SandboxDownloadBillUrl     = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
	SandboxDownloadFundFlowUrl = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadfundflow"
	SandboxReportUrl           = "https://api.mch.weixin.qq.com/sandboxnew/payitil/report"
	SandboxShortUrl            = "https://api.mch.weixin.qq.com/sandboxnew/tools/shorturl"
	SandboxAuthCodeToOpenidUrl = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"
	// 请求单次分账
	UrlProfitingSharding = "https://api.mch.weixin.qq.com/secapi/pay/profitsharing"
	// 请求多次分账
	UrlMultiProfitSharding = "https://api.mch.weixin.qq.com/secapi/pay/multiprofitsharing"
	// 分账结果查询
	UrlProfitShardingQuery = "https://api.mch.weixin.qq.com/pay/profitsharingquery"
	// 添加分账接收方
	UrlAddProfitReceiver = "https://api.mch.weixin.qq.com/pay/profitsharingaddreceiver"
	// 删除分账接收方
	UrlRemoveProfitReceiver = "https://api.mch.weixin.qq.com/pay/profitsharingremovereceiver"
	// 完结分账
	UrlProfitShardingFinish = "https://api.mch.weixin.qq.com/secapi/pay/profitsharingfinish"
	// 分账回退
	UrlProfitShardingReturn = "https://api.mch.weixin.qq.com/secapi/pay/profitsharingreturn"
	// 分账回退结果查询
	UrlProfitShardingReturnQuery = "https://api.mch.weixin.qq.com/pay/profitsharingreturnquery"
)
