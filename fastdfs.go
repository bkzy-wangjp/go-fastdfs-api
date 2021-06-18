package fastdfs

//主结构体
type GoFastDfs struct {
	DfsHost string
	Group   string
	host    string
}

//http请求参数结构
type ParamKeyVal struct {
	Key   string
	Value interface{}
}

//响应信息结构体
type JsonResult struct {
	Status  string      `json:"status"`  //状态,"ok"为正常
	Message string      `json:"message"` //信息
	Data    interface{} `json:"data"`
}

//文件列表返回信息 FileInfoResult
type DirInfo struct {
	IsDir   bool   `json:"is_dir"` //是否文件夹
	Md5     string `json:"md5"`    //Md5摘要
	Modtime int64  `json:"mtime"`  //创建时间UTC秒
	Name    string `json:"name"`   //名称
	Path    string `json:"path"`   //路径
	Size    int64  `json:"size"`   //大小,字节
}

//上传文件返回信息
type FileResult struct {
	Url     string `json:"url"`    //存储全路径,Domain+Src
	Md5     string `json:"md5"`    //文件的MD5摘要
	Path    string `json:"path"`   //存储路径
	Domain  string `json:"domain"` //域名
	Scene   string `json:"scene"`  //场景
	Size    int64  `json:"size"`   //文件大小,字节
	ModTime int64  `json:"mtime"`  //创建时间,UTC秒
	//Just for Compatibility
	//Scenes  string `json:"scenes"` //场景
	//Retmsg  string `json:"retmsg"`
	//Retcode int    `json:"retcode"`
	//Src     string `json:"src"` //存储路径
}

//文件信息
type FileInfo struct {
	Md5       string   `json:"md5"`       //Md5摘要
	Name      string   `json:"name"`      //名称
	Offset    int64    `json:"offset"`    //
	Path      string   `json:"path"`      //路径
	Peers     []string `json:"peers"`     //所在集群IP列表
	ReName    string   `json:"rename"`    //重命名
	Scene     string   `json:"scene"`     //场景
	Size      int64    `json:"size"`      //大小(字节)
	TimeStamp int64    `json:"timeStamp"` //创建时间(UTC秒)
}

//文件统计信息
type StatDateFileInfo struct {
	Date      string `json:"date"`      //日期
	FileCount int64  `json:"fileCount"` //文件数量
	TotalSize int64  `json:"totalSize"` //文件总大小(字节)
}

type Mail struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
type GloablConfig struct {
	Addr                 string   `json:"addr"`
	Peers                []string `json:"peers"`
	EnableHttps          bool     `json:"enable_https"`
	Group                string   `json:"group"`
	RenameFile           bool     `json:"rename_file"`
	ShowDir              bool     `json:"show_dir"`
	Extensions           []string `json:"extensions"`
	RefreshInterval      int      `json:"refresh_interval"`
	EnableWebUpload      bool     `json:"enable_web_upload"`
	DownloadDomain       string   `json:"download_domain"`
	EnableCustomPath     bool     `json:"enable_custom_path"`
	Scenes               []string `json:"scenes"`
	AlarmReceivers       []string `json:"alarm_receivers"`
	DefaultScene         string   `json:"default_scene"`
	Mail                 Mail     `json:"mail"`
	AlarmUrl             string   `json:"alarm_url"`
	DownloadUseToken     bool     `json:"download_use_token"`
	DownloadTokenExpire  int      `json:"download_token_expire"`
	QueueSize            int      `json:"queue_size"`
	AutoRepair           bool     `json:"auto_repair"`
	Host                 string   `json:"host"`
	FileSumArithmetic    string   `json:"file_sum_arithmetic"`
	PeerId               string   `json:"peer_id"`
	SupportGroupManage   bool     `json:"support_group_manage"`
	AdminIps             []string `json:"admin_ips"`
	EnableMergeSmallFile bool     `json:"enable_merge_small_file"`
	EnableMigrate        bool     `json:"enable_migrate"`
	EnableDistinctFile   bool     `json:"enable_distinct_file"`
	ReadOnly             bool     `json:"read_only"`
	EnableCrossOrigin    bool     `json:"enable_cross_origin"`
	EnableGoogleAuth     bool     `json:"enable_google_auth"`
	AuthUrl              string   `json:"auth_url"`
	EnableDownloadAuth   bool     `json:"enable_download_auth"`
	DefaultDownload      bool     `json:"default_download"`
	EnableTus            bool     `json:"enable_tus"`
	SyncTimeout          int64    `json:"sync_timeout"`
	EnableFsnotify       bool     `json:"enable_fsnotify"`
	EnableDiskCache      bool     `json:"enable_disk_cache"`
	ConnectTimeout       bool     `json:"connect_timeout"`
	ReadTimeout          int      `json:"read_timeout"`
	WriteTimeout         int      `json:"write_timeout"`
	IdleTimeout          int      `json:"idle_timeout"`
	ReadHeaderTimeout    int      `json:"read_header_timeout"`
	SyncWorker           int      `json:"sync_worker"`
	UploadWorker         int      `json:"upload_worker"`
	UploadQueueSize      int      `json:"upload_queue_size"`
	RetryCount           int      `json:"retry_count"`
	SyncDelay            int64    `json:"sync_delay"`
	WatchChanSize        int      `json:"watch_chan_size"`
}
