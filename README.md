声明：本项目遵守原作者著作权，仅仅分析学习研究，勿用于其他用途，欢迎大神指导。QQ群：228347251，验证：Mattermost

如果需要最新产品，请移步：
- 查看[官方站](https://www.mattermost.org/)
- 查看[文档](http://docs.mattermost.com/)


## Mattermost简介
Mattermost 是一个 Slack 的开源替代品。Mattermost 采用 Go 语言开发，这是一个开源的团队通讯服务。为团队带来跨 PC 和移动设备的消息、文件分享，提供归档和搜索功能。

## 项目结构

```Go
├─api4					//API
├─app					//控制器
├─build					//docker
│  └─docker
├─cmd					//脚本
│  ├─mattermost
│  └─platform
├─config				//配置
├─doc
├─einterfaces			//业务接口
├─fonts
├─i18n
├─imports
├─jobs
├─manualtesting
├─migrations
├─mlog
├─model					//业务数据建模
├─plugin
├─scripts
├─services				//服务封装
│  ├─configservice
│  ├─filesstore
│  │  └─mocks
│  ├─httpservice
│  ├─imageproxy
│  ├─mailservice
│  ├─mfa
│  └─timezones
├─store					//@@@@
│  ├─sqlstore
│  └─storetest
│      └─mocks
├─templates
├─testlib
├─tests
├─utils
│  ├─fileutils
│  ├─jsonutils
│  ├─markdown
│  └─testutils
├─vendor				//一堆包
├─web					//前端应用业务
└─wsapi					//socket业务
```