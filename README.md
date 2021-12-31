gmm 是go maxmind的缩写，一个用golang封装的对libmaxmind的访问API。

故事是这样的，我找到了一个叫telize的项目， 是基于kore做的一个用来查询maxmind的IP数据的API项目，于是把它封装成了一个docker镜像，命名为 renlu/telize,发布在https://hub.docker.com/r/renlu/telize 上（Dockerfile在 https://github.com/xurenlu/withpush )。

kore是一个用来方便发布http的api接口的项目；这个telize就是基于kore来做的。我压测的时候发布在20~30个并发的情况下达到20000个请求左右时接口会意外报错；
我怀疑是kore本身在处理网络情请时有问题，决定用golang重新封装一下，这就是gmm的由来。

gmm引入了gin这个golang的web framework，你也可以修改一下，去掉这个依赖，直接用golang原生的net/http包来处理请求。

最后的最后，尴尬的是，重写的时候高并发的时候依然会卡住，qps高的有时会急剧降到200个左右。。。不过依然足以满足日常查询用了。

这个项目的docker镜像，我发布到了docker.io的renlu/gmm 镜像了,工作在7654端口下:
```
docker run -p7654:7654 renlu/gmm 
```
接下来请求一下试试：
```
curl  localhost:7654/location/218.79.21.57
```
返回结果示例:
```
{"ip":"218.79.21.57","continent_code":"AS","country":"China","latitude":31.0442,"longitude":121.4054,"country_code":"CN","country_code3":"CHN","string":"","asn":4812,"organization":"-"}
```
最后，再次感谢一下maxmind；如果是国内对精准IP数据有需求，推荐 https://www.ipip.net/ 。

## and more

我买了个域名，把它运行起来了: https://ip4/dev/
