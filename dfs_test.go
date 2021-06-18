package fastdfs

import (
	"testing"
)

var Domain string = "http://127.0.0.1:8091"
var Group string = "group1"

func TestGetPostResp(t *testing.T) {
	tests := []struct {
		url    string
		params []ParamKeyVal
	}{
		{"list_dir", []ParamKeyVal{}},                                                 //文件列表
		{"list_dir", []ParamKeyVal{{"dir", "default"}}},                               //文件夹下的文件列表
		{"stat", []ParamKeyVal{}},                                                     //文件统计信息
		{"reload", []ParamKeyVal{{"action", "get"}}},                                  //获取配置信息
		{"get_file_info", []ParamKeyVal{{"md5", "8326be0ad92d165b05f6736d4c10d4e3"}}}, //获取文件信息
	}
	dfs := NewGoFastDfs(Domain, Group)
	for _, tt := range tests {
		rawstr, resp, err := dfs.getPostResp(tt.url, tt.params...)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(rawstr)
			t.Log(resp)
		}
	}
}

func TestGetDirList(t *testing.T) {
	tests := []struct {
		dir string
	}{
		{"default"},
		{"default/20210210"},
	}
	dfs := NewGoFastDfs(Domain, Group)
	for _, tt := range tests {
		rawstr, resp, err := dfs.GetDirList(tt.dir)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(rawstr)
			for _, dir := range resp {
				t.Log(dir)
			}
		}
	}
}

func TestGetStat(t *testing.T) {
	dfs := NewGoFastDfs(Domain, Group)
	rawstr, resp, err := dfs.GetStat()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log(rawstr)
		for _, dir := range resp {
			t.Log(dir)
		}
	}
}

func TestGetGloablConfig(t *testing.T) {
	dfs := NewGoFastDfs(Domain, Group)
	rawstr, resp, err := dfs.GetGloablConfig()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log(rawstr)
		t.Log(resp)
	}
}

var testpath, testmd5 string

func TestUploadFile(t *testing.T) {
	tests := []struct {
		filepath string
		params   []ParamKeyVal
	}{
		{"D:/Code/python/micpy/data/tmpimg/pylearn_1_hist_raw_20210205170215.svg", []ParamKeyVal{{"scene", "default"}, {"output", "json"}}},
	}
	dfs := NewGoFastDfs(Domain, Group)
	for _, tt := range tests {
		rawstr, resp, err := dfs.UploadFile(tt.filepath, tt.params...)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(rawstr)
			t.Log(resp)
			testpath = resp.Path
			testmd5 = resp.Md5
		}
	}
}

func TestGetFileInfo(t *testing.T) {
	tests := []struct {
		path string
		md5  string
	}{
		{testpath, ""},
		{"", testmd5},
	}
	dfs := NewGoFastDfs(Domain, Group)
	for _, tt := range tests {
		rawstr, resp, err := dfs.GetFileInfo(tt.path, tt.md5)
		t.Log(rawstr)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(resp)
		}
	}
}

func TestDeleteFile(t *testing.T) {
	tests := []struct {
		path string
		md5  string
	}{
		{testpath, testmd5},
	}
	dfs := NewGoFastDfs(Domain, Group)
	for _, tt := range tests {
		rawstr, resp, err := dfs.DeleteFile(tt.path, tt.md5)
		t.Log(rawstr)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(resp)
		}
	}
}
