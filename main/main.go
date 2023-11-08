// This file is auto-generated, don't edit it. Thanks.
package main

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

/**
* 使用AK&SK初始化账号Client
* @param accessKeyId
* @param accessKeySecret
* @return Client
* @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *domain20180129.Client, _err error) {
	//	client, _err := CreateClient(tea.String("LTAI5tF2YuFLw8BsE3peimBV"), tea.String("2SLj0xxkOTFvwz7b7EP6IfD1Ahqmib"))

	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Domain
	config.Endpoint = tea.String("domain.aliyuncs.com")
	_result = &domain20180129.Client{}
	_result, _err = domain20180129.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String("****uFLw8BsE3peimBV"), tea.String("xxxxx:vwz7b7EP6IfD1Ahqmib"))
	if _err != nil {
		return _err
	}

	queryAdvancedDomainListRequest := &domain20180129.QueryAdvancedDomainListRequest{
		PageNum:  tea.Int32(1),
		PageSize: tea.Int32(2),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		resp, _err := client.QueryAdvancedDomainListWithOptions(queryAdvancedDomainListRequest, runtime)

		if _err != nil {
			return _err
		}

		console.Log(util.ToJSONString(resp))
		if _err != nil {
			return _err
		}
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}

//
//package main
//
//import (
//	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
//	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
//	console "github.com/alibabacloud-go/tea-console/client"
//	util "github.com/alibabacloud-go/tea-utils/v2/service"
//	"github.com/alibabacloud-go/tea/tea"
//	"os"
//)
//
///**
//* 使用AK&SK初始化账号Client
//* @param accessKeyId
//* @param accessKeySecret
//* @return Client
//* @throws Exception
// */
//func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *openapi.Client, _err error) {
//	config := &openapi.Config{
//		// 必填，您的 AccessKey ID
//		AccessKeyId: accessKeyId,
//		// 必填，您的 AccessKey Secret
//		AccessKeySecret: accessKeySecret,
//	}
//	// Endpoint 请参考 https://api.aliyun.com/product/Domain
//	config.Endpoint = tea.String("domain.aliyuncs.com")
//	_result = &openapi.Client{}
//	_result, _err = openapi.NewClient(config)
//	return _result, _err
//}
//
///**
//* 使用STS鉴权方式初始化账号Client，推荐此方式。
//* @param accessKeyId
//* @param accessKeySecret
//* @param securityToken
//* @return Client
//* @throws Exception
// */
//func CreateClientWithSTS(accessKeyId *string, accessKeySecret *string, securityToken *string) (_result *openapi.Client, _err error) {
//	config := &openapi.Config{
//		// 必填，您的 AccessKey ID
//		AccessKeyId: accessKeyId,
//		// 必填，您的 AccessKey Secret
//		AccessKeySecret: accessKeySecret,
//		// 必填，您的 Security Token
//		SecurityToken: securityToken,
//		// 必填，表明使用 STS 方式
//		Type: tea.String("sts"),
//	}
//	// Endpoint 请参考 https://api.aliyun.com/product/Domain
//	config.Endpoint = tea.String("domain.aliyuncs.com")
//	_result = &openapi.Client{}
//	_result, _err = openapi.NewClient(config)
//	return _result, _err
//}
//
///**
//* API 相关
//* @param path params
//* @return OpenApi.Params
// */
//func CreateApiInfo() (_result *openapi.Params) {
//	params := &openapi.Params{
//		// 接口名称
//		Action: tea.String("QueryDomainList"),
//		// 接口版本
//		Version: tea.String("2018-01-29"),
//		// 接口协议
//		Protocol: tea.String("HTTPS"),
//		// 接口 HTTP 方法
//		Method:   tea.String("POST"),
//		AuthType: tea.String("AK"),
//		Style:    tea.String("RPC"),
//		// 接口 PATH
//		Pathname: tea.String("/"),
//		// 接口请求体内容格式
//		ReqBodyType: tea.String("json"),
//		// 接口响应体内容格式
//		BodyType: tea.String("json"),
//	}
//	_result = params
//	return _result
//}
//
//func _main(args []*string) (_err error) {
//	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
//	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
//	//LTAI5tF2YuFLw8BsE3peimBV
//	//2SLj0xxkOTFvwz7b7EP6IfD1Ahqmib
//	client, _err := CreateClient(tea.String("LTAI5tF2YuFLw8BsE3peimBV"), tea.String("2SLj0xxkOTFvwz7b7EP6IfD1Ahqmib"))
//	if _err != nil {
//		return _err
//	}
//	queryAdvancedDomainListRequest := &domain20180129.QueryAdvancedDomainListRequest{
//		PageNum:  tea.Int32(1),
//		PageSize: tea.Int32(2),
//	}
//	params := CreateApiInfo()
//	// runtime options
//	runtime := &util.RuntimeOptions{}
//	request := &openapi.OpenApiRequest{}
//	// 复制代码运行请自行打印 API 的返回值
//	// 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
//	resp, _err := client.CallApi(params, request, runtime)
//	if _err != nil {
//		return _err
//	}
//
//	console.Log(util.ToJSONString(resp))
//	return _err
//}
//
//func main() {
//	err := _main(tea.StringSlice(os.Args[1:]))
//	if err != nil {
//		panic(err)
//	}
//}
