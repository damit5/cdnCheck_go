package util

import (
	_ "embed"
	"encoding/json"
	"github.com/oschwald/geoip2-golang"
	"github.com/wxnacy/wgo/arrays"
	"net"
	"reflect"
	"strings"
)

// 初始化变量
var (
	ALL_CNAME []string
	CDN_CIDR []string
	ASNS = []uint {
		10576, 10762, 11748, 131099, 132601, 133496, 134409, 135295, 136764, 137187, 13777, 13890,
		14103, 14520, 17132, 199251, 200013, 200325, 200856, 201263, 202294, 203075, 203139, 204248,
		204286, 204545, 206227, 206734, 206848, 206986, 207158, 208559, 209403, 21030, 21257, 23327,
		23393, 23637, 23794, 24997, 26492, 268843, 28709, 29264, 30282, 30637, 328126, 36408,
		38107, 397192, 40366, 43303, 44907, 46071, 46177, 47542, 49287, 49689, 51286, 55082,
		55254, 56636, 57363, 58127, 59730, 59776, 60068, 60626, 60922, 61107, 61159, 62026, 62229,
		63062, 64232, 8868, 9053, 55770, 49846, 49249, 48163, 45700, 43639, 39836, 393560, 393234,
		36183, 35994, 35993, 35204, 34850, 34164, 33905, 32787, 31377, 31110, 31109, 31108, 31107,
		30675, 24319, 23903, 23455, 23454, 22207, 21399, 21357, 21342, 20940, 20189, 18717, 18680,
		17334, 16702, 16625, 12222, 209101, 201585, 135429, 395747, 394536, 209242, 203898, 202623,
		14789, 133877, 13335, 132892, 21859, 6185, 47823, 30148,
	}
)
var DB *geoip2.Reader
var err error

//go:embed GeoLite2-ASN.mmdb
var getlite2_asn []byte

// 来源：https://raw.githubusercontent.com/mabangde/cdncheck_cn/main/sources_data.json
//go:embed sources_data.json
var data string

func init() {
	DB, err = geoip2.FromBytes(getlite2_asn)
	if err != nil {
		return 
	}

	// 定义JSON结构
	type DataStruct struct {
		CDN struct{
			Baidu      []string `json:"Baidu-加速乐"`
			Cloudfront []string `json:"cloudfront"`
			Fastly     []string `json:"fastly"`
			Google     []string `json:"google"`
			Leaseweb   []string `json:"leaseweb"`
			Stackpath  []string `json:"stackpath"`
			YunDun     []string `json:"云盾CDN"`
			BaiDuZNY   []string `json:"百度智能云CDN"`
			WangXiu    []string `json:"网宿 CDN"`
			Wangshen   []string `json:"网神CDN"`
			TencentYun []string `json:"腾讯云CDN"`
			LanXun     []string `json:"蓝讯"`
			Aliyun     []string `json:"阿里云 CDN"`
		} `json:"cdn"`
		WAF struct{
			Akamai     []string `json:"akamai"`
			Cloudflare []string `json:"cloudflare"`
			Incapsula  []string `json:"incapsula"`
		} `json:"waf"`
		Common struct{
			Qihoo360CdnOperatedByQihoo360 []string `json:"360 云 CDN (由奇安信运营)"`
			Qihoo360CdnOperatedByQihoo3602 []string `json:"360 云 CDN (由奇虎 360 运营)"`
			AttContentDeliveryNetwork     []string `json:"AT&T Content Delivery Network"`
			AwsCloud                       []string `json:"AWS Cloud"`
			AwsCloudFront                  []string `json:"AWS CloudFront"`
			AkamaiCdn                      []string `json:"Akamai CDN"`
			AzionTechEdgeComputingPlatform []string `json:"Azion Tech | Edge Computing Platform"`
			BaiduJiasule                    []string `json:"Baidu-加速乐"`
			BelugaCDN                       []string `json:"BelugaCDN"`
			BilibiliBusinessGSLB            []string `json:"Bilibili 业务 GSLB"`
			BilibiliHighAvailabilityRegionLoadBalancing []string `json:"Bilibili 高可用地域负载均衡"`
			BilibiliHighAvailabilityLoadBalancing    []string `json:"Bilibili 高可用负载均衡"`
			BunnyCDN                               []string `json:"Bunny CDN"`
			CDNDotNET                              []string `json:"CDN.NET"`
			CDNDotNETCDNSUNONAPP                   []string `json:"CDN.NET / CDNSUN / ONAPP"`
			CDN77                                  []string `json:"CDN77"`
			CDNIFY                                 []string `json:"CDNIFY"`
			CDNSUN                                 []string `json:"CDNSUN"`
			CDNetworks                             []string `json:"CDNetworks"`
			CacheFlyCDN                            []string `json:"CacheFly CDN"`
			CedexisGSLB                            []string `json:"Cedexis GSLB"`
			CedexisGSLBForChina                    []string `json:"Cedexis GSLB (For China)"`
			CenturyLinkCDNOriginalLevel3           []string `json:"CenturyLink CDN (原 Level 3)"`
			CloudXNS                               []string `json:"CloudXNS"`
			Cloudflare                              []string `json:"Cloudflare"`
			CodingPages                             []string `json:"Coding Pages"`
			ConversantSwiftServeCDN                 []string `json:"Conversant - SwiftServe CDN"`
			Fastly                                  []string `json:"Fastly"`
			FlexBalancerSmartTrafficRouting          []string `json:"FlexBalancer - Smart Traffic Routing"`
			GCoreLabs                               []string `json:"G - Core Labs"`
			GitHubPages                              []string `json:"GitHub Pages"`
			GitLabPages                              []string `json:"GitLab Pages"`
			GoogleCloudStorage                       []string `json:"Google Cloud Storage"`
			GoogleWebBusiness                        []string `json:"Google Web 业务"`
			HerokuSaaS                                []string `json:"Heroku SaaS"`
			IncapsulaCDN                             []string `json:"Incapsula CDN"`
			InstartCDN                               []string `json:"Instart CDN"`
			InternapCDN                              []string `json:"Internap CDN"`
			KeyCDN                                   []string `json:"KeyCDN"`
			LeaseWebCDN                              []string `json:"LeaseWeb CDN"`
			LimelightNetwork                         []string `json:"Limelight Network"`
			Medianova                                []string `json:"Medianova"`
			MicrosoftAzure                           []string `json:"Microsoft Azure"`
			MicrosoftAzureAppService                 []string `json:"Microsoft Azure App Service"`
			MicrosoftAzureCDN                        []string `json:"Microsoft Azure CDN"`
			MicrosoftAzureTrafficManager             []string `json:"Microsoft Azure Traffic Manager"`
			Netlify                                  []string `json:"Netlify"`
			NodeCache                                []string `json:"NodeCache"`
			OracleDynWebApplicationSecuritySuite     []string `json:"Oracle Dyn Web Application Security suite (原 Zenedge CDN)"`
			QUANTILNetEaseWangsu                     []string `json:"QUANTIL (网宿)"`
			QUICCloud                                []string `json:"QUIC.Cloud"`
			RelectedNetworks                         []string `json:"Relected Networks"`
			SpeedyCloudCDN                          []string `json:"SpeedyCloud CDN"`
			StackpathOriginalHighwinds              []string `json:"Stackpath (原 Highwinds)"`
			StackpathOriginalMaxCDN                 []string `json:"Stackpath (原 MaxCDN)"`
			StackpathCDN                             []string `json:"Stackpath CDN"`
			StatusPageIO                             []string `json:"StatusPage.io"`
			TAN14CDN                                []string `json:"TAN14 CDN"`
			TataCommunicationsCDN                  []string `json:"Tata communications CDN"`
			TurboBytesMultiCDN                     []string `json:"TurboBytes Multi-CDN"`
			UCloudCDN                              []string `json:"UCloud CDN"`
			UCloudRomeGlobalNetworkAcceleration    []string `json:"UCloud 罗马 Rome 全球网络加速"`
			VerizonCDNEdgecast                     []string `json:"Verizon CDN (Edgecast)"`
			VeryCloudCloudDistribution             []string `json:"VeryCloud 云分发"`
			WebLukerBlueFlood                      []string `json:"WebLuker (蓝汛)"`
			ZEITNowSmartCDN                        []string `json:"ZEIT Now Smart CDN"`
			ZenlayerCDN                            []string `json:"Zenlayer CDN"`
			Amazon                                 []string `json:"amazon"`
			CloudflareCNAME                        []string `json:"cloudflare"`
			Edgecast                               []string `json:"edgecast"`
			FastlyCNAME                            []string `json:"fastly"`
			Incapsula                              []string `json:"incapsula"`
			MmTrixPerformanceMagicCube             []string `json:"mmTrix性能魔方（高升控股旗下）"`
			QiniuCloud                             []string `json:"七牛云"`
			ShanghaiYundunCDN                      []string `json:"上海云盾 CDN"`
			CenturyInterconnectCloudExpressService []string `json:"世纪互联云快线业务"`
			CenturyInterconnectShanghaiLanyunAzure []string `json:"世纪互联旗下上海蓝云（承载 Azure 中国）"`
			ChinaTelecomTianyiCloudCDN             []string `json:"中国电信天翼云CDN"`
			ZhonglianDataZhonglianLiXin            []string `json:"中联数据（中联利信）"`
			YunfanAcceleratedCDN                   []string `json:"云帆加速CDN"`
			CloudBrainFusionCDN                    []string `json:"云端智度融合 CDN"`
			CloudBrainNetwork                        []string `json:"云端网络"`
			JingdongCloudCDN                        []string `json:"京东云 CDN"`
			YisuyunCDN                              []string `json:"亿速云 CDN"`
			QuansuyunWangsuCloudEdgeCloudAcceleration []string `json:"全速云（网宿）CloudEdge 云加速"`
			ChuangshiyunFusionCDN                    []string `json:"创世云融合 CDN"`
			DonglizaixianCDN                         []string `json:"动力在线CDN"`
			BeijingTongxingwandianNetworkTechnology []string `json:"北京同兴万点网络技术"`
			HuaweiCloudCDN                          []string `json:"华为云 CDN"`
			HuaweiCloudWAFHighDefenseShield         []string `json:"华为云WAF高防云盾"`
			Youpaiyun                                []string `json:"又拍云"`
			KuakeCloudCDN                            []string `json:"可靠云 CDN (贴图库)"`
			QiAnXinWangshenCDN                      []string `json:"奇安信网神CDN"`
			QiAnXinWebsiteGuard                     []string `json:"奇安信网站卫士"`
			ByteDanceCDN                             []string `json:"字节跳动 CDN"`
			ByteDanceSubsidiaryVolcanoEngine         []string `json:"字节跳动旗下火山引擎"`
			AnhengXuanwuShieldWAF                    []string `json:"安恒玄武盾 （WAF）"`
			BaotengHulianShanghaiWangenNetworkCDN    []string `json:"宝腾互联旗下上海万根网络（CDN 联盟）"`
			BaotengHulianShanghaiWangenNetworkYaoCDN []string `json:"宝腾互联旗下上海万根网络（YaoCDN）"`
			DilianCDN                                []string `json:"帝联 CDN"`
			GuangdongWangdiCDN                       []string `json:"广东网堤CDN"`
			KuaiwangCDN                              []string `json:"快网 CDN"`
			SouhuYuntaiCDN                           []string `json:"搜狐云台CDN"`
			NewLeShiYunLianCDN                       []string `json:"新乐视云联（原乐视云）CDN"`
			New1CloudCDN                             []string `json:"新壹云-NEW1CLOUD"`
			XinLiuYunCDN                             []string `json:"新流云（新流万联）"`
			SinaCloudCDN                             []string `json:"新浪云 CDN"`
			SinaCloudSAEEngine                       []string `json:"新浪云 SAE 云引擎"`
			SinaTechFusionCDNLoadBalancer            []string `json:"新浪科技融合CDN负载均衡"`
			SinaStaticDomain                         []string `json:"新浪静态域名"`
			YiTongRuiJinAkamaiChinaByWangsu          []string `json:"易通锐进（Akamai 中国）由网宿承接"`
			XingYuYunP2PCDN                          []string `json:"星域云P2P CDN"`
			JiYuYunAnquanYiYun                       []string `json:"极御云安全（浙江壹云云计算有限公司）"`
			JiSuDun                                  []string `json:"极速盾"`
			ShenXinFuYunDun                          []string `json:"深信服云盾"`
			NiuDunYunAnquan                          []string `json:"牛盾云安全"`
			MaoYunRongHeCDN                          []string `json:"猫云融合 CDN"`
			BaiShanYunCDN                            []string `json:"白山云 CDN"`
			BaiduCloudCDN                                      []string `json:"百度云 CDN"`
			BaiduCloudAcceleration                             []string `json:"百度云加速"`
			BaiduSubsidiaryBusinessRegionalLoadBalancingSystem []string `json:"百度旗下业务地域负载均衡系统"`
			BaiduIntelligentCloudCDN                           []string `json:"百度智能云CDN"`
			KnownSecYunAnquanCDN                               []string `json:"知道创宇云安全 CDN"`
			KnownSecYunAnquanChuangYuDunGovernment             []string `json:"知道创宇云安全创宇盾（政务专用）"`
			KnownSecYunAnquanJiaSuLeCDN                        []string `json:"知道创宇云安全加速乐CDN"`
			GreenShieldYunWAF                                  []string `json:"绿盟云 WAF"`
			WangsuCDN                                          []string `json:"网宿 CDN"`
			WangsuWAFCDN                                       []string `json:"网宿 WAF CDN"`
			NetEaseCloudCDN                                    []string `json:"网易云 CDN"`
			MeituanCloudCDN                                    []string `json:"美团云 CDN"`
			MeituanCloudSanKuaiTechnologyLoadBalancer          []string `json:"美团云（三快科技）负载均衡"`
			MeiChengHuiLianCDN                                 []string `json:"美橙互联CDN"`
			MeiChengHuiLianSubsidiaryBuildStar                 []string `json:"美橙互联旗下建站之星"`
			TengZhengAnQuanJiaSu15CDN                          []string `json:"腾正安全加速（原 15CDN）"`
			TencentCloudAPIGateway                             []string `json:"腾讯云 API 网关"`
			TencentCloudDDoSProtection                           []string `json:"腾讯云 DDoS 防护"`
			TencentCloudCDN                                      []string `json:"腾讯云CDN"`
			TencentCloudGlobalApplicationAcceleration            []string `json:"腾讯云全球应用加速"`
			TencentCloudDaYuBGPHighDefense                       []string `json:"腾讯云大禹 BGP 高防"`
			TencentCloudObjectStorage                            []string `json:"腾讯云对象存储"`
			TencentCloudLiveCDN                                  []string `json:"腾讯云直播 CDN"`
			TencentCloudVideoCDN                                 []string `json:"腾讯云视频 CDN"`
			TencentSubsidiaryBusinessRegionalLoadBalancingSystem []string `json:"腾讯旗下业务地域负载均衡系统"`
			LanXunCDN                                            []string `json:"蓝汛 CDN"`
			LanDunYunCDN                                         []string `json:"蓝盾云CDN"`
			LanShiYunCDN                                         []string `json:"蓝视云 CDN"`
			AntFinancialSubsidiaryBusinessRegionalLoadBalancingSystem []string `json:"蚂蚁金服旗下业务地域负载均衡系统"`
			ManManYunCDNZhongLianLiXin                                []string `json:"蛮蛮云 CDN（中联利信）"`
			XiBuShuMa                                                 []string `json:"西部数码"`
			XiBuShuMaCDN                                              []string `json:"西部数码CDN"`
			YiYunKeJiYunJiaSuCDN                                      []string `json:"逸云科技云加速 CDN"`
			JinShanYunCDN                                             []string `json:"金山云 CDN"`
			RuiSuYunCDN                                               []string `json:"锐速云 CDN"`
			AliyunCDN                                                 []string `json:"阿里云 CDN"`
			AliyunGlobalTrafficManagement                             []string `json:"阿里云全局流量管理"`
			AliyunDunHighDefense                                      []string `json:"阿里云盾高防"`
			QueNiuYunCDNAliyun                                        []string `json:"雀牛云CDN 阿里云"`
			QingCloudCDN []string `json:"青云 CDN"`
			QingYeYunCDN []string `json:"青叶云 CDN"`
			LingZhiYunCDNHangzhouLingZhiYunHua []string `json:"领智云 CDN（杭州领智云画）"`
			ElemeStaticDomainRegionalLoadBalancing []string `json:"饿了么静态域名与地域负载均衡"`
			GaoShengKongGuCDNTechnology []string `json:"高升控股CDN技术"`
			MoMenYunCDN []string `json:"魔门云 CDN"`
		} `json:"common"`
	}

	// 反序列化
	var cdnData DataStruct
	json.Unmarshal([]byte(data), &cdnData)

	// 通过反射获取所有字段的数据
	// 初始化 CIDR 部分
	cdnReflectValue := reflect.ValueOf(cdnData.CDN)
	for i := 0; i < cdnReflectValue.NumField(); i++ {
		field := cdnReflectValue.Field(i)
		fieldValue := field.Interface().([]string)
		CDN_CIDR = append(CDN_CIDR, fieldValue...)
	}
	wafReflectValue := reflect.ValueOf(cdnData.WAF)
	for i := 0; i < wafReflectValue.NumField(); i++ {
		field := wafReflectValue.Field(i)
		fieldValue := field.Interface().([]string)
		CDN_CIDR = append(CDN_CIDR, fieldValue...)
	}

	// 初始化 CNAME 部分
	cnameReflectValue := reflect.ValueOf(cdnData.Common)
	for i := 0; i < cnameReflectValue.NumField(); i++ {
		field := cnameReflectValue.Field(i)
		fieldValue := field.Interface().([]string)
		ALL_CNAME = append(ALL_CNAME, fieldValue...)
	}
}

// CheckCNAME 是否有CDN， true 有, false 没有
func CheckCNAME(domain string) bool {
	found := false

	cname, err := net.LookupCNAME(domain)
	if err != nil {	// 解析错误也会返回为有CDN
		return true
	}

	for _, cdncname := range ALL_CNAME {
		if strings.Contains(cname, cdncname) {
			found = true
			break
		}
	}
	return found
}

// CheckCIDR 通过CIDR来判断， true 有, false 没有
func CheckCIDR(ip net.IP) bool {
	found := false
	for _, cidr := range CDN_CIDR {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			return false
		}

		if ipnet.Contains(ip) {
			found = true
			break
		}
	}
	return found
}

// CheckASN true 有, false 没有
// https://dev.maxmind.com/geoip/geolite2-free-geolocation-data
func CheckASN(ip net.IP) bool {
	found := false
	asn, err := DB.ASN(ip)
	if err != nil {
		return false
	}
	contains := arrays.Contains(ASNS, asn.AutonomousSystemNumber)
	if contains != -1 {
		found = true
	}
	return found
}