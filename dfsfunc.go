package fastdfs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/client/httplib"
)

//新建结构
func NewGoFastDfs(host, group string) *GoFastDfs {
	dfs := new(GoFastDfs)
	dfs.DfsHost = host
	dfs.Group = group
	dfs.MakeHost()
	return dfs
}

func (dfs *GoFastDfs) MakeHost() {
	dfs.host = dfs.DfsHost
	if len(dfs.Group) > 0 {
		dfs.host = fmt.Sprintf("%s/%s", dfs.DfsHost, dfs.Group)
	}
}

//POST请求
//输入:requrl string:请求url,是group/之后的部分
//    params ...ParamKeyVal,参数键值对
//输出:
//    string:请求返回的原始字符串
//    *FastDfsResponse:请求转换为结构体之后的信息
//    error:错误信息. 如果 status!="ok",也会产生错误信息
func (dfs *GoFastDfs) getPostResp(requrl string, params ...ParamKeyVal) (string, *JsonResult, error) {
	url := fmt.Sprintf("%s/%s", dfs.host, requrl)
	req := httplib.Post(url)
	for _, param := range params { //添加参数
		req.Param(param.Key, fmt.Sprint(param.Value))
	}
	res := new(JsonResult)
	respstr, err := req.String()
	if err != nil {
		return respstr, res, err
	}
	err = json.Unmarshal([]byte(respstr), res)
	if err != nil {
		return respstr, res, err
	}
	if res.Status != "ok" {
		err = fmt.Errorf(res.Message)
	}
	return respstr, res, err
}

//上传文件
//可选的params参数：
//scene string:场景,默认 default
//output string:输出格式,json|text,默认text
//path string:自定义路径,默认根据日期自动创建
//filename string:重命名文件
//code string:goolge认证码
//auth_token string:自定义认证码
//md5 string:用于秒传的文件md5摘要
func (dfs *GoFastDfs) UploadFile(filepath string, params ...ParamKeyVal) (string, *FileResult, error) {
	url := fmt.Sprintf("%s/upload", dfs.host)
	req := httplib.Post(url)
	for _, param := range params { //添加参数
		req.Param(param.Key, fmt.Sprint(param.Value))
	}
	req.PostFile("file", filepath)
	res := new(FileResult)
	respstr, err := req.String()
	if err != nil {
		return respstr, res, err
	}
	err = json.Unmarshal([]byte(respstr), res)
	return respstr, res, err
}

//数据接口转[]*FileInfoResult结构体
func DataToStruct(data interface{}, result interface{}) error {
	//var dirlist []*FileInfoResult
	if msg, err := json.Marshal(data); err == nil {
		if err := json.Unmarshal(msg, result); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

//获取文件列表
func (dfs *GoFastDfs) GetDirList(dir ...string) (*JsonResult, []*DirInfo, error) {
	var params []ParamKeyVal
	if len(dir) > 0 {
		var param ParamKeyVal
		param.Key = "dir"
		param.Value = dir[0]
		params = append(params, param)
	}
	_, resp, err := dfs.getPostResp("list_dir", params...)
	if err != nil {
		return resp, nil, err
	}
	var res []*DirInfo
	err = DataToStruct(resp.Data, &res)
	resp.Data = res
	return resp, res, err
}

//获取文件信息
//path和md5二选一,不选的设置为 ""
func (dfs *GoFastDfs) GetFileInfo(path, md5 string) (*JsonResult, *FileInfo, error) {
	var key, val string
	if len(path) > 0 {
		key = "path"
		val = path
	} else {
		key = "md5"
		val = md5
	}
	_, resp, err := dfs.getPostResp("get_file_info", ParamKeyVal{key, val})
	if err != nil {
		return resp, nil, err
	}
	res := new(FileInfo)
	err = DataToStruct(resp.Data, res)
	resp.Data = *res
	return resp, res, err
}

//在服务器上删除文件
//path和md5二选一,不选的设置为 ""
func (dfs *GoFastDfs) DeleteFile(path, md5 string) (*JsonResult, bool, error) {
	var key, val string
	if len(path) > 0 {
		key = "path"
		val = path
	} else {
		key = "md5"
		val = md5
	}
	_, resp, err := dfs.getPostResp("delete", ParamKeyVal{key, val})
	if err != nil {
		return resp, false, err
	}
	if resp.Status == "ok" {
		return resp, true, nil
	}
	return resp, false, fmt.Errorf(resp.Message)
}

//获取文件统计信息
func (dfs *GoFastDfs) GetStat() (*JsonResult, []*StatDateFileInfo, error) {
	_, resp, err := dfs.getPostResp("stat")
	if err != nil {
		return resp, nil, err
	}
	var res []*StatDateFileInfo
	err = DataToStruct(resp.Data, &res)
	resp.Data = res
	return resp, res, err
}

//获取配置信息
func (dfs *GoFastDfs) GetGloablConfig() (*JsonResult, *GloablConfig, error) {
	_, resp, err := dfs.getPostResp("reload", ParamKeyVal{"action", "get"})
	if err != nil {
		return resp, nil, err
	}
	res := new(GloablConfig)
	err = DataToStruct(resp.Data, res)
	resp.Data = *res
	return resp, res, err
}

//文件信息转换为文件上传结果
func (dfs *GoFastDfs) FileInfo2FileResult(finfo *FileInfo, domain, group string) *FileResult {
	fres := new(FileResult)
	fres.Domain = domain
	fres.Md5 = finfo.Md5
	fres.ModTime = finfo.TimeStamp
	paths := strings.Split(finfo.Path, "/")
	if len(paths) > 0 {
		paths[0] = group
	}
	paths = append(paths, finfo.ReName)
	fres.Path = "/"
	for i, s := range paths {
		if i < len(paths)-1 {
			fres.Path = fres.Path + s + "/"
		} else {
			fres.Path = fres.Path + s
		}
	}
	fres.Scene = finfo.Scene
	fres.Size = finfo.Size
	fres.Url = domain + fres.Path

	return fres
}
