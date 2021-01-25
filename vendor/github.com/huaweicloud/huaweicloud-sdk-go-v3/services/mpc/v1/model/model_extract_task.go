/*
 * MPC
 *
 * # 接口调用方法  ## 服务使用方法   公有云API符合RESTful API的设计理论。   REST从资源的角度来观察整个网络，分布在各处的资源由URI（Uniform Resource Identifier）确定，而客户端的应用通过URL（Unified Resource Locator）来获取资源。   URL的一般格式为：https://Endpoint/uri   URL中的参数说明如表1所示。    **表1** URL中的参数说明    参数 | 描述   ------------ | ------------   Endpoint | Web服务入口点的URL，从地区和终端节点[http://developer.huaweicloud.com/endpoint.html](http://developer.huaweicloud.com/endpoint.html)获取。   uri | 资源路径，也即API访问路径。从具体接口的URI模块获取，例如“v3/auth/tokens”。  ## 请求方法   在HTTP协议中，请求可以使用多种请求方法例如GET、PUT、POST、DELETE，用于指明以何种方式来访问指定的资源，目前媒体处理服务提供的REST接口支持的请求方法如下表所示：    方法 | 说明   ------------ | -------------   GET | 请求服务器返回指定资源   PUT | 请求服务器更新指定资源   POST | 请求服务器新增资源或执行特殊操作   DELETE | 请求服务器删除指定资源，如删除对象等  ## 公共参数  ### 公共请求消息头参数    指每一个接口都需要使用到的参数：    消息头名称 | 描述 | 是否必选   ------------ | ------------- | -------------   Host | 主机地址。媒体处理服务的Host为： {endpoint} | 是   Content-Type | 发送的实体的MIME类型。默认值：application/json | 是   Content-Length | 资源内容的长度。 | 否   X-sdk-date | 请求的发生时间，格式为(YYYYMMDD'T'HHMMSS'Z')。取值为当前系统的GMT时间 | 否   Authorization | 请求消息中携带的鉴权信息。 | 否   X-Auth-Token | 如果使用Token认证的方式，此字段携带认证密钥。类型：字符串。默认值：无 | 否    > 使用AK/SK方式认证时，上述X-sdk-date、Authorization字段必填。使用Token方式认证时，X-Auth-Token字段必填。  ### API接口中的公共参数   API接口的uri里要求填写project_id，project_id指的是在公有云上注册成功的租户拥有的项目编号，如创建转码接口：/v1/{project_id}/transcodings。  ### 公共响应消息头参数   指每一个接口都会返回的参数：    消息头名称 | 描述 | 是否必选   ------------ | ------------- | -------------   Content-Type | 发送给接收者的实体正文的媒体类型。类型：字符串。默认值：application/json; charset=UTF-8。 | 是   X-request-id | 此字段携带请求ID号，以便任务跟踪。类型：字符串。request_id-timestamp-hostname(request_id在服务器端生成UUID， timestamp为当前时间戳，hostname为处理当前接口的服务器名称)。默认值：无。 | 是   X-ratelimit | 此字段携带总计流控请求数。类型：整型。默认值：无。 | 否   X-ratelimit-used | 此字段携带剩下请求数。类型：整型。默认值：无。 | 否   X-ratelimit-window | 此字段携带流控单位，有分钟、小时、天三种。类型：字符串。默认值：小时。 | 否  ## 返回参数   调用API服务后返回数据采用统一格式。   返回2xx的HTTP状态码代表调用成功，返回4xx或5xx的HTTP状态码代表调用失败。   统一返回json格式的响应。  ## 请求认证方式  ### 生成AK、SK   请按照如下步骤创建AK、SK：   1. 登录[华为云服务](https://www.huaweicloud.com/)。   2. 在华为云服务首页上面单击“控制台”进入控制台管理页面。   3. 单击右上方登录的用户，进入“账号中心”页面。   4. 在“基本信息”页面，单击“管理我的凭证”。   5. 在“我的凭证”页面，单击“管理访问密钥”，单击“新增访问秘钥”，即可创建AK、SK。  ### 请求签名流程    * **签名前准备**   1. 下载API网关签名工具。   下载地址：http://esdk.huawei.com/ilink/esdk/download/HW_456706   2. 解压下载的压缩包。   创建java工程，将解压出来的jar引用到依赖路径中。    * **签名过程**   1.  创建用于签名的请求com.cloud.sdk.DefaultRequest(JAVA)。   2.  设置DefaultRequest的目标API URL、HTTPS方法、内容等信息。   3.  对DefaultRequest进行签名：   调用SignerFactory.getSigner(String serviceName, String regionName)获取一个签名工具实现的实例。   调用Signer.sign(Request<?> request, Credentials credentials)对步骤1创建的请求进行签名。   以下代码展示了这个步骤：       ```java       //选用签名算法，对请求进行签名       Signer signer = SignerFactory.getSigner(serviceName, region);       //对请求进行签名，request会发生改变       signer.sign(request, new BasicCredentials(this.ak, this.sk));       ```       把签名产生的request转换为一个适合发送的请求，并将签名后request中的header信息放入新的request中。       以Apache HttpClient为例，需要把DefaultRequest转换为HttpRequestBase，把签名后的DefaultRequest的header信息放入HttpRequestBase中。  ### 示例代码   下面代码展示了如何对一个请求进行签名，并通过HTTP Client发送一个HTTPS请求的过程。   代码分成三个类进行演示：   * AccessService：抽象类，将GET/POST/PUT/DELETE归一成access方法。   * Demo：运行入口，模拟用户进行GET/POST/PUT/DELETE请求。   * AccessServiceImpl：实现access方法，具体与API网关通信的代码都在access方法中。  ## 查看项目ID   在调用接口的时候，部分URL中需要填入项目编号(project_id)，所以需要先在管理控制台上获取到项目ID。   项目ID获取步骤如下：   1. 登录[华为云服务](https://www.huaweicloud.com/)。   2. 在华为云服务首页上面单击“控制台”进入控制台管理页面。   3. 单击右上方登录的用户，进入“账号中心”页面。   4. 在“基本信息”页面，单击“管理我的凭证”。   5. 在“我的凭证”页面的“项目列表”页签查看项目ID。  ## API概览  转码服务对应的接口列表如下：   接口 | 说明   ------------ | -------------   POST /v1/{project_id}/transcodings | 创建转码任务   DELETE /v1/{project_id}/transcodings{?task_id} | 删除转码任务   GET /v1/{project_id}/transcodings{?task_id} | 查询转码任务   GET /v1/{project_id}/transcodings/detail{?task_id} | 查询转码任务详情   POST /v1/{project_id}/template/transcodings | 创建自定义转码模板   DELETE /v1/{project_id}/template/transcodings{?temp_id} | 删除自定义转码模板   PUT /v1/{project_id}/template/transcodings | 修改自定义转码模板   GET /v1/{project_id}/template/transcodings{?temp_id} | 查询自定义转码模板   POST /v1/{project_id}/thumbnails | 创建下发截图任务   DELETE /v1/{project_id}/thumbnails{?task_id} | 取消已下发截图任务   GET /v1/{project_id}/thumbnails{?task_id} | 查询截图任务状态   PUT /v1/{project_id}/notification | 配置转码服务端事件通知   GET /v1/{project_id}/notification | 查询转码服务端事件通知   GET /v1/{project_id}/notification/template |  查询转码服务端事件通知模板(待废弃)   GET /v1/{project_id}/notification/event | 查询转码服务所有事件   DELETE /v1/{project_id}/transcodings/task{?task_id} | 删除转码任务记录   PUT /v1/{project_id}/transcodings/task  | 重试任务   GET /v1/{project_id}/buckets | 查询桶列表   GET /v1.0-ext/{project_id}/objects{?bucket,prefix,type} | 查询对象列表   PUT /v1/{project_id}/authority | 桶授权或取消授权   GET /v1/{project_id}/transcodings/summary{?range} |  查询转码概览信息   POST /v1/{project_id}/extract_audio | 创建下发音频处理任务   GET /v1/{project_id}/extract_audio | 查询音频处理任务接口   DELETE /v1/{project_id}/extract_audio | 取消音频处理任务接口   GET /v1/{project_id}/transcodings/summaries{?start_time,end_time,stat_type} |  查询转码概览信息   POST /v1/{project_id}/agencies |  创建/取消委托授权任务   GET  /v1/{project_id}/agencies |  查询委托授权任务   POST /v1/{project_id}/encryptions | 创建独立加密任务   DELETE /v1/{project_id}/encryptions{?task_id} | 删除独立加密任务   GET /v1/{project_id}/encryptions{?task_id} | 查询独立加密任务   POST /v1/{project_id}/enhancements | 创建视频增强任务   DELETE /v1/{project_id}/enhancements{?task_id} | 删除视频增强任务   GET /v1/{project_id}/enhancements{?task_id} | 查询视频增强任务   POST /v1/{project_id}/template/qualityenhance | 创建视频增强配置模板   DELETE /v1/{project_id}/template/qualityenhance{?template_id} | 删除视频增强配置模板   PUT /v1/{project_id}/template/qualityenhance | 修改视频增强配置模板   GET /v1/{project_id}/template/qualityenhance{?template_id} | 查询视频增强配置模板   GET /v1/{project_id}/template/qualityenhance/default | 查询视频增强配置预置模板   POST /v1/{project_id}/remux |  创建下发转封装任务   GET /v1/{project_id}/remux |  查询转封装任务   DELETE /v1/{project_id}/remux | 取消转封装任务   POST /v1/{project_id}/animated-graphics |  创建下发动图任务   GET /v1/{project_id}/animated-graphics |  查询动图任务   DELETE /v1/{project_id}/animated-graphics | 取消动图任务   POST /v1/{project_id}/extract-metadata |  创建解析任务   GET /v1/{project_id}/extract-metadata |  查询解析任务   DELETE /v1/{project_id}/extract-metadata | 取消解析任务   POST /v1/ctm/notification | ctm通知mpc转码结果   GET /health-check | 容器健康检查   DELETE /v1/{project_id}/mpc/cleanup | 注销账号资源清理接口   POST /v1/{project_id}/audio/services/merge_channels/task  | 合并音频多声道文件   DELETE /v1/{project_id}/audio/services/merge_channels/task |  取消合并音频多声道文件   GET /v1/{project_id}/audio/services/merge_channels/task |  查询合并音频多声道文件任务   POST /v1/{project_id}/audio/services/reset_tracks/task  |  重置音频文件声轨   DELETE /v1/{project_id}/audio/services/reset_tracks/task | 取消重置音频文件声轨任务   GET /v1/{project_id}/audio/services/reset_tracks/task |  查询重置音频文件声轨任务   PUT /v1/mediabox/tasks/report |  合并音频多声道文件任务、重置音频文件声轨任务上报结果接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ExtractTask struct {
	// 任务ID
	TaskId *string `json:"task_id,omitempty"`
	// 任务状态。  取值如下： - INIT：初始状态。 - WAITING：等待启动。 - PROCESSING：处理中。 - SUCCEED：处理成功。 - FAILED：处理失败。 - CANCELED：已取消。
	Status *string `json:"status,omitempty"`
	// 任务创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 任务启动时间
	StartTime *string `json:"start_time,omitempty"`
	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 错误描述
	Description *string     `json:"description,omitempty"`
	Input       *ObsObjInfo `json:"input,omitempty"`
	Output      *ObsObjInfo `json:"output,omitempty"`
	// 用户数据。
	UserData *string   `json:"user_data,omitempty"`
	Metadata *MetaData `json:"metadata,omitempty"`
}

func (o ExtractTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExtractTask struct{}"
	}

	return strings.Join([]string{"ExtractTask", string(data)}, " ")
}
