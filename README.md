## LBS Emergency
-------------------

### 简介
-------------------

LBS-Emergency是提供基于地理位置的求助服务。应用场景有：例如路上遇到交通事故的时候，使用电话紧急求助110/120甚至保险公司都需要语言对当前位置进行详尽描述。在一些特殊位置、或者比较紧急的情况下，想要简短地描述清楚当前位置会比较困难。利用LBS-Emergency可以一键将当前位置的经纬度、街道等信息发送至求助地，以达到减少求助时间，更快地获得帮助的目的。


### 当前版本v1.0
-------------------

支持手机端浏览器访问[测试地址](http://123.206.58.28:11000)，能够显示当前街道、静态地图以及可缩放的动态地图。

详情见[CHANGELOG](https://github.com/maxwell92/LBS-Emergency/CHANGELOG.md)


### 计划：
-------------------

1. 客户端能够利用百度地图API进行定位、静态地图显示以及缩放。

2. 客户端能够利用浏览器、IP等进行定位，并将地址转换为地图供显示，同时在求助发出的时候将经纬度地址、街道名等以Json的形式发送至服务器端。

3. 服务器端容器化，能接受到客户端发来的地址(json形式)并结合百度地图API成功定位，并能计算出路线。

4. 大量请求时的同步处理以及水平扩展。

5. 前面的都是集中式，后面可采用分布式


TODO:

1. 修正通过浏览器定位的精度，现在偏差太大

2. 使用微信公众号，并采用已有的腾讯云做后台主机提供服务
