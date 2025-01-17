package common

import (
	"github.com/Albert-Zhan/httpc"
	"github.com/unknwon/goconfig"
)

const (
	SoftName          = "jd_seckill"
	Version           = "0.2.1"
	DateTimeFormatStr = "2006-01-02 15:04:05"
	DateFormatStr     = "2006-01-02"
	IniFileContent    = `#默认配置
[config]
# 设置日志级别：debug=全部打印debug、info、warn、error级别日志(开发模式)；info=打印出info、warn、error级别日志（默认级别）；warn=只打印warn、error级别日志（生产模式）
log_level =
# 登录二维码展示方式：1.open=调用UI打开 2.print=控制台输出展示 3.dingtalk=钉钉机器人推送（[dingtalk]必填）
qrcode_show_type = open
# 第三方生成二维码API，仅钉钉机器人推送时有用
qrcode_create_api = https://api.pwmqr.com/qrcode/create/?url=
# eid, fp参数必须填写，具体请参考 wiki-常见问题
# 随意填写可能导致订单无法提交等问题
eid =
fp =
# 商品id
# 已经是茅台的sku_id了
sku_id = 100012043978
# 抢购数量
seckill_num = 2
# 抢购开始时间设定 2021-01-01 09:59:59 (PS.预约成功后会自动更新)
buy_time = 2021-01-01 09:59:59
# 抢购总时间，单位:分钟，默认两分钟
seckill_time =
# 抢购任务数量，默认5个
task_num =
# 每次抢购间隔时间，单位:毫秒，默认1500毫秒，每1000毫秒等于1秒
ticker_time =
# 默认UA
default_user_agent = Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36

#账号信息
[account]
# 支付密码
# 如果你的账户中有可用的京券（注意不是东券）或 在上次购买订单中使用了京豆，
# 那么京东可能会在下单时自动选择京券支付 或 自动勾选京豆支付。
# 此时下单会要求输入六位数字的支付密码。请在下方配置你的支付密码，如 123456 。
# 如果没有上述情况，下方请留空。
payment_pwd =

#消息推送
[messenger]
# 开启推送服务
enable = false
# 目前支持smtp邮箱推送和Server酱推送服务，选值smtp,wechat,dingtalk
type = none
# 邮箱推送消息接收人
email = 
#Server酱推送key，当type为wechat有效
server_chan_sckey =

#smtp配置
# 开启smtp消息推送必须填入 email_user、email_pwd，email_host，port 如何使用请自行百度。
[smtp]
# 邮箱域名 smtp.xx.com
email_host =
# 通信端口
port =
# 邮箱地址 xxxxxxxx@xx.com
email_user =
# 邮箱授权码（并不一定是邮箱密码） xxxxxxxxxxxxxxxx
email_pwd =

#钉钉机器人配置
[dingtalk]
# 机器人的Webhook地址(https://oapi.dingtalk.com/robot/send?access_token=XXXXXX)的access_token
access_token =
# 密钥，机器人安全设置页面，加签一栏下面显示的SEC开头的字符串
secret =
# 钉钉提醒：1.不提醒=none(或留空)；2.所有人=all；3.指定人=11位手机号_1,11位手机号_2 (多个人以半角逗号分隔)
at =
`
)

var (
	SoftDir string

	Client *httpc.HttpClient

	CookieJar *httpc.CookieJar

	Config *goconfig.ConfigFile

	SeckillStatus chan bool
)
