package api

import (
	"github.com/finomen/go-moonraker-api/jsonrpc"
	uuid "github.com/satori/go.uuid"
)

type ServerConnectionIdentityRequest struct {
	ClientName string `json:"client_name"`
	Version    string `json:"version"`
	Type       string `json:"type"`
	Url        string `json:"url"`
}

type ServerConnectionIdentityResponse struct {
	ConnectionId int `json:"connection_id"`
}

var ServerConnectionIdentity = jsonrpc.Method[ServerConnectionIdentityRequest, ServerConnectionIdentityResponse]{Name: "server.connection.identify"}

type GetWebsocketIdResponse struct {
	WebsocketId int `json:"websocket_id"`
}

var GetWebsocketId = jsonrpc.Method[struct{}, GetWebsocketIdResponse]{Name: "server.websocket.id"}

type PrinterInfoResponse struct {
	State           string `json:"state"`
	StateMessage    string `json:"state_message"`
	Hostname        string `json:"hostname"`
	SoftwareVersion string `json:"software_version"`
	CpuInfo         string `json:"cpu_info"`
	KlipperPath     string `json:"klipper_path"`
	PythonPath      string `json:"python_path"`
	LogFile         string `json:"log_file"`
	ConfigFile      string `json:"config_file"`
}

var PrinterInfo = jsonrpc.Method[struct{}, PrinterInfoResponse]{Name: "printer.info"}

var PrinterEmergencyStop = jsonrpc.Method[struct{}, struct{}]{Name: "printer.emergency_stop"}
var PrinterRestart = jsonrpc.Method[struct{}, struct{}]{Name: "printer.restart"}
var PrinterFirmwareRestart = jsonrpc.Method[struct{}, struct{}]{Name: "printer.firmware_restart"}

type PrinterObjectsListResponse struct {
	Objects []string `json:"objects"`
}

var PrinterObjectsList = jsonrpc.Method[struct{}, PrinterObjectsListResponse]{Name: "printer.objects.list"}

//TODO: check null deserialization!
type PrinterObjectsQueryRequest struct {
	Objects map[string][]string `json:"objects"`
}

//TODO: define objects https://moonraker.readthedocs.io/en/latest/printer_objects/
type PrinterObjectsQueryResponse struct {
	Eventtime float64                `json:"eventtime"`
	Status    map[string]interface{} `json:"status"`
}

var PrinterObjectsQuery = jsonrpc.Method[PrinterObjectsQueryRequest, PrinterObjectsQueryResponse]{Name: "printer.objects.query"}

var PrinterObjectsSubscribe = jsonrpc.Method[PrinterObjectsQueryRequest, PrinterObjectsQueryResponse]{Name: "printer.objects.subscribe"}

type EndstopState string

const (
	Open      EndstopState = "open"
	Triggered EndstopState = "TRIGGERED"
)

type PrinterQueryEndstopsStatusResponse map[string]EndstopState

var PrinterQueryEndstopsStatus = jsonrpc.Method[struct{}, PrinterQueryEndstopsStatusResponse]{Name: "printer.query_endstops.status"}

type ServerInfoResponse struct {
	KlippyConnected       bool     `json:"klippy_connected"`
	KlippyState           string   `json:"klippy_state"`
	Components            []string `json:"components"`
	FailedComponents      []string `json:"failed_components"`
	RegisteredDirectories []string `json:"registered_directories"`
	Warnings              []string `json:"warnings"`
	WebsocketCount        uint64   `json:"websocket_count"`
	MoonrakerVersion      string   `json:"moonraker_version"`
	ApiVersion            []int    `json:"api_version"`
	ApiVersionString      string   `json:"api_version_string"`
}

var Serverinfo = jsonrpc.Method[struct{}, ServerInfoResponse]{Name: "server.info"}

type ServerConfigResponse struct {
	Config struct {
		Server struct {
			Host               string      `json:"host"`
			Port               int         `json:"port"`
			SslPort            int         `json:"ssl_port"`
			EnableDebugLogging bool        `json:"enable_debug_logging"`
			EnableAsyncioDebug bool        `json:"enable_asyncio_debug"`
			KlippyUdsAddress   string      `json:"klippy_uds_address"`
			MaxUploadSize      int         `json:"max_upload_size"`
			SslCertificatePath interface{} `json:"ssl_certificate_path"`
			SslKeyPath         interface{} `json:"ssl_key_path"`
		} `json:"server"`
		DbusManager struct {
		} `json:"dbus_manager"`
		Database struct {
			DatabasePath        string `json:"database_path"`
			EnableDatabaseDebug bool   `json:"enable_database_debug"`
		} `json:"database"`
		FileManager struct {
			EnableObjectProcessing bool   `json:"enable_object_processing"`
			QueueGcodeUploads      bool   `json:"queue_gcode_uploads"`
			ConfigPath             string `json:"config_path"`
			LogPath                string `json:"log_path"`
		} `json:"file_manager"`
		KlippyApis struct {
		} `json:"klippy_apis"`
		Machine struct {
			Provider string `json:"provider"`
		} `json:"machine"`
		ShellCommand struct {
		} `json:"shell_command"`
		DataStore struct {
			TemperatureStoreSize int `json:"temperature_store_size"`
			GcodeStoreSize       int `json:"gcode_store_size"`
		} `json:"data_store"`
		ProcStats struct {
		} `json:"proc_stats"`
		JobState struct {
		} `json:"job_state"`
		JobQueue struct {
			LoadOnStartup       bool   `json:"load_on_startup"`
			AutomaticTransition bool   `json:"automatic_transition"`
			JobTransitionDelay  int    `json:"job_transition_delay"`
			JobTransitionGcode  string `json:"job_transition_gcode"`
		} `json:"job_queue"`
		HttpClient struct {
		} `json:"http_client"`
		Announcements struct {
			DevMode       bool          `json:"dev_mode"`
			Subscriptions []interface{} `json:"subscriptions"`
		} `json:"announcements"`
		Authorization struct {
			LoginTimeout   int      `json:"login_timeout"`
			ForceLogins    bool     `json:"force_logins"`
			CorsDomains    []string `json:"cors_domains"`
			TrustedClients []string `json:"trusted_clients"`
		} `json:"authorization"`
		Zeroconf struct {
		} `json:"zeroconf"`
		OctoprintCompat struct {
			EnableUfp     bool   `json:"enable_ufp"`
			FlipH         bool   `json:"flip_h"`
			FlipV         bool   `json:"flip_v"`
			Rotate90      bool   `json:"rotate_90"`
			StreamUrl     string `json:"stream_url"`
			WebcamEnabled bool   `json:"webcam_enabled"`
		} `json:"octoprint_compat"`
		History struct {
		} `json:"history"`
		Secrets struct {
			SecretsPath string `json:"secrets_path"`
		} `json:"secrets"`
		Mqtt struct {
			Address       string      `json:"address"`
			Port          int         `json:"port"`
			Username      string      `json:"username"`
			PasswordFile  interface{} `json:"password_file"`
			Password      string      `json:"password"`
			MqttProtocol  string      `json:"mqtt_protocol"`
			InstanceName  string      `json:"instance_name"`
			DefaultQos    int         `json:"default_qos"`
			StatusObjects struct {
				Webhooks       interface{} `json:"webhooks"`
				Toolhead       string      `json:"toolhead"`
				IdleTimeout    string      `json:"idle_timeout"`
				GcodeMacroM118 interface{} `json:"gcode_macro M118"`
			} `json:"status_objects"`
			ApiQos             int  `json:"api_qos"`
			EnableMoonrakerApi bool `json:"enable_moonraker_api"`
		} `json:"mqtt"`
		Template struct {
		} `json:"template"`
	} `json:"config"`
	Orig struct {
		DEFAULT struct {
		} `json:"DEFAULT"`
		Server struct {
			EnableDebugLogging string `json:"enable_debug_logging"`
			MaxUploadSize      string `json:"max_upload_size"`
		} `json:"server"`
		FileManager struct {
			ConfigPath             string `json:"config_path"`
			LogPath                string `json:"log_path"`
			QueueGcodeUploads      string `json:"queue_gcode_uploads"`
			EnableObjectProcessing string `json:"enable_object_processing"`
		} `json:"file_manager"`
		Machine struct {
			Provider string `json:"provider"`
		} `json:"machine"`
		Announcements struct {
		} `json:"announcements"`
		JobQueue struct {
			JobTransitionDelay string `json:"job_transition_delay"`
			JobTransitionGcode string `json:"job_transition_gcode"`
			LoadOnStartup      string `json:"load_on_startup"`
		} `json:"job_queue"`
		Authorization struct {
			TrustedClients string `json:"trusted_clients"`
			CorsDomains    string `json:"cors_domains"`
		} `json:"authorization"`
		Zeroconf struct {
		} `json:"zeroconf"`
		OctoprintCompat struct {
		} `json:"octoprint_compat"`
		History struct {
		} `json:"history"`
		Secrets struct {
			SecretsPath string `json:"secrets_path"`
		} `json:"secrets"`
		Mqtt struct {
			Address            string `json:"address"`
			Port               string `json:"port"`
			Username           string `json:"username"`
			Password           string `json:"password"`
			EnableMoonrakerApi string `json:"enable_moonraker_api"`
			StatusObjects      string `json:"status_objects"`
		} `json:"mqtt"`
	} `json:"orig"`
	Files []struct {
		Filename string   `json:"filename"`
		Sections []string `json:"sections"`
	} `json:"files"`
}

var ServerConfig = jsonrpc.Method[struct{}, ServerConfigResponse]{Name: "server.config"}

type ServerTemperatureStoreResponse map[string]struct {
	Temperatures []float64 `json:"temperatures"`
	Targets      []int     `json:"targets"`
	Powers       []int     `json:"powers"`
}

var ServerTemperatureStore = jsonrpc.Method[struct{}, ServerTemperatureStoreResponse]{Name: "server.temperature_store"}

type ServerGCodeStoreRequest struct {
	Count uint64 `json:"count"`
}

type ServerGCodeStoreResponse struct {
	GcodeStore []struct {
		Message string  `json:"message"`
		Time    float64 `json:"time"`
		Type    string  `json:"type"`
	} `json:"gcode_store"`
}

var ServerGCodeStore = jsonrpc.Method[ServerGCodeStoreRequest, ServerGCodeStoreResponse]{Name: "server.gcode_store"}

var ServerRestart = jsonrpc.Method[struct{}, struct{}]{Name: "server.restart"}

type PrinterGCodeScriptRequest struct {
	Script string `json:"script"`
}

var PrinterGCodeScript = jsonrpc.Method[PrinterGCodeScriptRequest, struct{}]{Name: "printer.gcode.script"}

type PrinterGCodeHelpResponse map[string]string

var PrinterGCodeHelp = jsonrpc.Method[struct{}, PrinterGCodeHelpResponse]{Name: "printer.gcode.help"}

type PrinterPrintStartRequest struct {
	Filename string `json:"filename"`
}

var PrinterPrintStart = jsonrpc.Method[PrinterPrintStartRequest, struct{}]{Name: "printer.print.start"}
var PrinterPrintPause = jsonrpc.Method[struct{}, struct{}]{Name: "printer.print.pause"}
var PrinterPrintResume = jsonrpc.Method[struct{}, struct{}]{Name: "printer.print.resume"}
var PrinterPrintCancel = jsonrpc.Method[struct{}, struct{}]{Name: "printer.print.cancel"}

type MachineSystemInfoResponse struct {
	SystemInfo struct {
		CpuInfo struct {
			CpuCount     int    `json:"cpu_count"`
			Bits         string `json:"bits"`
			Processor    string `json:"processor"`
			CpuDesc      string `json:"cpu_desc"`
			SerialNumber string `json:"serial_number"`
			HardwareDesc string `json:"hardware_desc"`
			Model        string `json:"model"`
			TotalMemory  int    `json:"total_memory"`
			MemoryUnits  string `json:"memory_units"`
		} `json:"cpu_info"`
		SdInfo struct {
			ManufacturerId   string `json:"manufacturer_id"`
			Manufacturer     string `json:"manufacturer"`
			OemId            string `json:"oem_id"`
			ProductName      string `json:"product_name"`
			ProductRevision  string `json:"product_revision"`
			SerialNumber     string `json:"serial_number"`
			ManufacturerDate string `json:"manufacturer_date"`
			Capacity         string `json:"capacity"`
			TotalBytes       int64  `json:"total_bytes"`
		} `json:"sd_info"`
		Distribution struct {
			Name         string `json:"name"`
			Id           string `json:"id"`
			Version      string `json:"version"`
			VersionParts struct {
				Major       string `json:"major"`
				Minor       string `json:"minor"`
				BuildNumber string `json:"build_number"`
			} `json:"version_parts"`
			Like     string `json:"like"`
			Codename string `json:"codename"`
		} `json:"distribution"`
		AvailableServices []string `json:"available_services"`
		ServiceState      struct {
			Klipper struct {
				ActiveState string `json:"active_state"`
				SubState    string `json:"sub_state"`
			} `json:"klipper"`
			KlipperMcu struct {
				ActiveState string `json:"active_state"`
				SubState    string `json:"sub_state"`
			} `json:"klipper_mcu"`
			Moonraker struct {
				ActiveState string `json:"active_state"`
				SubState    string `json:"sub_state"`
			} `json:"moonraker"`
		} `json:"service_state"`
		Virtualization struct {
			VirtType       string `json:"virt_type"`
			VirtIdentifier string `json:"virt_identifier"`
		} `json:"virtualization"`
		Python struct {
			Version       []interface{} `json:"version"`
			VersionString string        `json:"version_string"`
		} `json:"python"`
		Network struct {
			Wlan0 struct {
				MacAddress  string `json:"mac_address"`
				IpAddresses []struct {
					Family      string `json:"family"`
					Address     string `json:"address"`
					IsLinkLocal bool   `json:"is_link_local"`
				} `json:"ip_addresses"`
			} `json:"wlan0"`
		} `json:"network"`
	} `json:"system_info"`
}

var MachineSystemInfo = jsonrpc.Method[struct{}, MachineSystemInfoResponse]{Name: "machine.system_info"}
var MachineShutdown = jsonrpc.Method[struct{}, struct{}]{Name: "machine.shutdown"}
var MachineReboot = jsonrpc.Method[struct{}, struct{}]{Name: "machine.reboot"}

type MachineServicesRestartRequest struct {
	Service string `json:"service"`
}

var MachineServicesRestart = jsonrpc.Method[MachineServicesRestartRequest, struct{}]{Name: "machine.services.restart"}

type MachineServicesStopRequest struct {
	Service string `json:"service"`
}

var MachineServicesStop = jsonrpc.Method[MachineServicesRestartRequest, struct{}]{Name: "machine.services.stop"}

type MachineServicesStartRequest struct {
	Service string `json:"service"`
}

var MachineServicesStart = jsonrpc.Method[MachineServicesRestartRequest, struct{}]{Name: "machine.services.start"}

type MachineProcStatsResponse struct {
	MoonrakerStats []struct {
		Time     float64 `json:"time"`
		CpuUsage float64 `json:"cpu_usage"`
		Memory   int     `json:"memory"`
		MemUnits string  `json:"mem_units"`
	} `json:"moonraker_stats"`
	ThrottledState struct {
		Bits  int           `json:"bits"`
		Flags []interface{} `json:"flags"`
	} `json:"throttled_state"`
	CpuTemp float64 `json:"cpu_temp"`
	Network struct {
		Lo struct {
			RxBytes   int     `json:"rx_bytes"`
			TxBytes   int     `json:"tx_bytes"`
			Bandwidth float64 `json:"bandwidth"`
		} `json:"lo"`
		Wlan0 struct {
			RxBytes   float64 `json:"rx_bytes"`
			TxBytes   float64 `json:"tx_bytes"`
			Bandwidth float64 `json:"bandwidth"`
		} `json:"wlan0"`
	} `json:"network"`
	SystemCpuUsage struct {
		Cpu  float64 `json:"cpu"`
		Cpu0 float64 `json:"cpu0"`
		Cpu1 float64 `json:"cpu1"`
		Cpu2 float64 `json:"cpu2"`
		Cpu3 float64 `json:"cpu3"`
	} `json:"system_cpu_usage"`
	SystemUptime         float64 `json:"system_uptime"`
	WebsocketConnections int     `json:"websocket_connections"`
}

var MachineProcStats = jsonrpc.Method[struct{}, MachineProcStatsResponse]{Name: "machine.proc_stats"}

type ServerFilesListRequest struct {
	Root string `json:"root"`
}

type ServerFilesListResponse []struct {
	Path        string  `json:"path"`
	Modified    float64 `json:"modified"`
	Size        int     `json:"size"`
	Permissions string  `json:"permissions"`
}

var ServerFilesList = jsonrpc.Method[ServerFilesListRequest, ServerFilesListResponse]{Name: "server.files.list"}

type ServerFilesMetadataRequest struct {
	Filename string `json:"filename"`
}

type ServerFilesMetadataResponse struct {
	PrintStartTime   interface{} `json:"print_start_time"`
	JobId            interface{} `json:"job_id"`
	Size             int         `json:"size"`
	Modified         float64     `json:"modified"`
	Slicer           string      `json:"slicer"`
	SlicerVersion    string      `json:"slicer_version"`
	LayerHeight      float64     `json:"layer_height"`
	FirstLayerHeight float64     `json:"first_layer_height"`
	ObjectHeight     float64     `json:"object_height"`
	FilamentTotal    float64     `json:"filament_total"`
	EstimatedTime    int         `json:"estimated_time"`
	Thumbnails       []struct {
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		Size         int    `json:"size"`
		RelativePath string `json:"relative_path"`
	} `json:"thumbnails"`
	FirstLayerBedTemp  int    `json:"first_layer_bed_temp"`
	FirstLayerExtrTemp int    `json:"first_layer_extr_temp"`
	GcodeStartByte     int    `json:"gcode_start_byte"`
	GcodeEndByte       int    `json:"gcode_end_byte"`
	Filename           string `json:"filename"`
}

var ServerFilesMetadata = jsonrpc.Method[ServerFilesMetadataRequest, ServerFilesMetadataResponse]{Name: "server.files.metadata"}

type ServerFilesGetDirectoryRequest struct {
	Path     string `json:"path"`
	Extended bool   `json:"extended"`
}

type ServerFilesGetDirectoryResponse struct {
	Dirs []struct {
		Modified    float64 `json:"modified"`
		Size        int     `json:"size"`
		Permissions string  `json:"permissions"`
		Dirname     string  `json:"dirname"`
	} `json:"dirs"`
	Files []struct {
		Modified    float64 `json:"modified"`
		Size        int     `json:"size"`
		Permissions string  `json:"permissions"`
		Filename    string  `json:"filename"`
	} `json:"files"`
	DiskUsage struct {
		Total int64 `json:"total"`
		Used  int64 `json:"used"`
		Free  int64 `json:"free"`
	} `json:"disk_usage"`
	RootInfo struct {
		Name        string `json:"name"`
		Permissions string `json:"permissions"`
	} `json:"root_info"`
}

var ServerFilesGetDirectory = jsonrpc.Method[ServerFilesGetDirectoryRequest, ServerFilesGetDirectoryResponse]{Name: "server.files.get_directory"}

type ServerFilesPostDirectoryRequest struct {
	Path string `json:"path"`
}

type ServerFilesPostDirectoryResponse struct {
	Item struct {
		Path string `json:"path"`
		Root string `json:"root"`
	} `json:"item"`
	Action string `json:"action"`
}

var ServerFilesPostDirectory = jsonrpc.Method[ServerFilesPostDirectoryRequest, ServerFilesPostDirectoryResponse]{Name: "server.files.post_directory"}

type ServerFilesDeleteDirectoryRequest struct {
	Path  string `json:"path"`
	Force bool   `json:"force"`
}

type ServerFilesDeleteDirectoryResponse struct {
	Item struct {
		Path string `json:"path"`
		Root string `json:"root"`
	} `json:"item"`
	Action string `json:"action"`
}

var ServerFilesDeleteDirectory = jsonrpc.Method[ServerFilesDeleteDirectoryRequest, ServerFilesDeleteDirectoryResponse]{Name: "server.files.delete_directory"}

type ServerFilesMoveRequest struct {
	Source string `json:"source"`
	Dest   string `json:"dest"`
}

type ServerFilesMoveResponse struct {
	Item struct {
		Root string `json:"root"`
		Path string `json:"path"`
	} `json:"item"`
	SourceItem struct {
		Path string `json:"path"`
		Root string `json:"root"`
	} `json:"source_item"`
	Action string `json:"action"`
}

var ServerFilesMove = jsonrpc.Method[ServerFilesMoveRequest, ServerFilesMoveResponse]{Name: "server.files.move"}

type ServerFilesCopyRequest struct {
	Source string `json:"source"`
	Dest   string `json:"dest"`
}

type ServerFilesCopyResponse struct {
	Item struct {
		Root string `json:"root"`
		Path string `json:"path"`
	} `json:"item"`
	Action string `json:"action"`
}

var ServerFilesCopy = jsonrpc.Method[ServerFilesCopyRequest, ServerFilesCopyResponse]{Name: "server.files.copy"}

/// Start kcc extension methods

type CloudDownloadRequest struct {
	Root       string    `json:"root"`
	Path       string    `json:"path"`
	DownloadId uuid.UUID `json:"downloadId"`
}

type CloudDownloadResponse struct {
	Size        uint64 `json:"size"`
	ContentType string `json:"contentType"`
	Status      uint   `json:"status"`
}

var CloudDownload = jsonrpc.Method[CloudDownloadRequest, CloudDownloadResponse]{Name: "cloud.download"}

type CloudUploadRequest struct {
	Root     string `json:"root"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	Checksum string `json:"checksum"`

	DownloadId uuid.UUID `json:"downloadId"`
}

type CloudUploadResponse struct {
	Status uint `json:"status"`
}

var CloudUpload = jsonrpc.Method[CloudUploadRequest, CloudUploadResponse]{Name: "cloud.upload"}

/// End kcc extension methods

type ServerFilesDeleteFileRequest struct {
	Path string `json:"path"`
}

type ServerFilesDeleteFileResponse struct {
	Item struct {
		Path string `json:"path"`
		Root string `json:"root"`
	} `json:"item"`
	Action string `json:"action"`
}

var ServerFilesDeleteFile = jsonrpc.Method[ServerFilesDeleteFileRequest, ServerFilesDeleteFileResponse]{Name: "server.files.delete_file"}

// TODO: Authorization methods ?

type ServerDatabaseListRequest struct {
	Root string `json:"root,omitempty"`
}

type ServerDatabaseListResponse struct {
	Namespaces []string `json:"namespaces"`
}

var ServerDatabaseList = jsonrpc.Method[ServerDatabaseListRequest, ServerDatabaseListResponse]{Name: "server.database.list"}

type ServerDatabaseGetItemRequest struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
}

type ServerDatabaseGetItemResponse struct {
	Namespace string      `json:"namespace"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

var ServerDatabaseGetItem = jsonrpc.Method[ServerDatabaseGetItemRequest, ServerDatabaseGetItemResponse]{Name: "server.database.get_item"}

type ServerDatabasePostItemRequest struct {
	Namespace string      `json:"namespace"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type ServerDatabasePostItemResponse struct {
	Namespace string      `json:"namespace"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

var ServerDatabasePostItem = jsonrpc.Method[ServerDatabasePostItemRequest, ServerDatabasePostItemResponse]{Name: "server.database.post_item"}

type ServerDatabaseDeleteItemRequest struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
}

type ServerDatabaseDeleteItemResponse struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Value     int    `json:"value"`
}

var ServerDatabaseDeleteItem = jsonrpc.Method[ServerDatabaseDeleteItemRequest, ServerDatabaseDeleteItemResponse]{Name: "server.database.delete_item"}

type ServerJobQueueStatusResponse struct {
	QueuedJobs []struct {
		Filename    string  `json:"filename"`
		JobId       string  `json:"job_id"`
		TimeAdded   float64 `json:"time_added"`
		TimeInQueue float64 `json:"time_in_queue"`
	} `json:"queued_jobs"`
	QueueState string `json:"queue_state"`
}

var ServerJobQueueStatus = jsonrpc.Method[struct{}, ServerJobQueueStatusResponse]{Name: "server.job_queue.status"}

type ServerJobQueuePostJobRequest struct {
	Filenames []string `json:"filenames"`
}

type ServerJobQueuePostJobResponse struct {
	QueuedJobs []struct {
		Filename    string  `json:"filename"`
		JobId       string  `json:"job_id"`
		TimeAdded   float64 `json:"time_added"`
		TimeInQueue float64 `json:"time_in_queue"`
	} `json:"queued_jobs"`
	QueueState string `json:"queue_state"`
}

var ServerJobQueuePostJob = jsonrpc.Method[ServerJobQueuePostJobRequest, ServerJobQueuePostJobResponse]{Name: "server.job_queue.post_job"}

type ServerJobQueueDeleteJobRequest struct {
	JobIds []string `json:"job_ids"`
}

type ServerJobQueueDeleteJobResponse struct {
	QueuedJobs []struct {
		Filename    string  `json:"filename"`
		JobId       string  `json:"job_id"`
		TimeAdded   float64 `json:"time_added"`
		TimeInQueue float64 `json:"time_in_queue"`
	} `json:"queued_jobs"`
	QueueState string `json:"queue_state"`
}

var ServerJobQueueDeleteJob = jsonrpc.Method[ServerJobQueueDeleteJobRequest, ServerJobQueueDeleteJobResponse]{Name: "server.job_queue.delete_job"}

type ServerJobQueuePauseResponse struct {
	QueuedJobs []struct {
		Filename    string  `json:"filename"`
		JobId       string  `json:"job_id"`
		TimeAdded   float64 `json:"time_added"`
		TimeInQueue float64 `json:"time_in_queue"`
	} `json:"queued_jobs"`
	QueueState string `json:"queue_state"`
}

var ServerJobQueuePause = jsonrpc.Method[struct{}, ServerJobQueuePauseResponse]{Name: "server.job_queue.pause"}

type ServerJobQueueStartResponse struct {
	QueuedJobs []struct {
		Filename    string  `json:"filename"`
		JobId       string  `json:"job_id"`
		TimeAdded   float64 `json:"time_added"`
		TimeInQueue float64 `json:"time_in_queue"`
	} `json:"queued_jobs"`
	QueueState string `json:"queue_state"`
}

var ServerJobQueueStart = jsonrpc.Method[struct{}, ServerJobQueueStartResponse]{Name: "server.job_queue.pause"}

type ServerAnnouncementsListRequest struct {
	IncludeDismissed *bool `json:"include_dismissed,omitempty"`
}

type ServerAnnouncementsListResponse struct {
	Entries []struct {
		EntryId       string      `json:"entry_id"`
		Url           string      `json:"url"`
		Title         string      `json:"title"`
		Description   string      `json:"description"`
		Priority      string      `json:"priority"`
		Date          int         `json:"date"`
		Dismissed     bool        `json:"dismissed"`
		DateDismissed interface{} `json:"date_dismissed"`
		DismissWake   interface{} `json:"dismiss_wake"`
		Source        string      `json:"source"`
		Feed          string      `json:"feed"`
	} `json:"entries"`
	Feeds []string `json:"feeds"`
}

var ServerAnnouncementsList = jsonrpc.Method[ServerAnnouncementsListRequest, ServerAnnouncementsListResponse]{Name: "server.announcements.list"}

type ServerAnnouncementsUpdateResponse struct {
	Entries []struct {
		EntryId     string `json:"entry_id"`
		Url         string `json:"url"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    string `json:"priority"`
		Date        int    `json:"date"`
		Dismissed   bool   `json:"dismissed"`
		Source      string `json:"source"`
		Feed        string `json:"feed"`
	} `json:"entries"`
	Modified bool `json:"modified"`
}

var ServerAnnouncementsUpdate = jsonrpc.Method[struct{}, ServerAnnouncementsUpdateResponse]{Name: "server.announcements.update"}

type ServerAnnouncementsDismissRequest struct {
	EntryId  string  `json:"entry_id"`
	WakeTime float64 `json:"wake_time"`
}

type ServerAnnouncementsDismissResponse struct {
	EntryId string `json:"entry_id"`
}

var ServerAnnouncementsDismiss = jsonrpc.Method[ServerAnnouncementsDismissRequest, ServerAnnouncementsDismissResponse]{Name: "server.announcements.dismiss"}

type ServerAnnouncementsFeedsResponse struct {
	Feeds []string `json:"feeds"`
}

var ServerAnnouncementsFeeds = jsonrpc.Method[struct{}, ServerAnnouncementsFeedsResponse]{Name: "server.announcements.feeds"}

type ServerAnnouncementsPostFeedRequest struct {
	Name string `json:"name"`
}

type ServerAnnouncementsPostFeedResponse struct {
	Feed   string `json:"feed"`
	Action string `json:"action"`
}

var ServerAnnouncementsPostFeed = jsonrpc.Method[ServerAnnouncementsPostFeedRequest, ServerAnnouncementsPostFeedResponse]{Name: "server.announcements.post_feed"}

type ServerAnnouncementsDeleteFeedRequest struct {
	Name string `json:"name"`
}

type ServerAnnouncementsDeleteFeedResponse struct {
	Feed   string `json:"feed"`
	Action string `json:"action"`
}

var ServerAnnouncementsDeleteFeed = jsonrpc.Method[ServerAnnouncementsDeleteFeedRequest, ServerAnnouncementsDeleteFeedResponse]{Name: "server.announcements.delete_feed"}

type MachineUpdateStatusRequest struct {
	Refresh bool `json:"refresh"`
}

type MachineUpdateStatusResponse struct {
	Busy                    bool `json:"busy"`
	GithubRateLimit         int  `json:"github_rate_limit"`
	GithubRequestsRemaining int  `json:"github_requests_remaining"`
	GithubLimitResetTime    int  `json:"github_limit_reset_time"`
	VersionInfo             struct {
		System struct {
			PackageCount int      `json:"package_count"`
			PackageList  []string `json:"package_list"`
		} `json:"system"`
		Moonraker struct {
			Channel           string        `json:"channel"`
			DebugEnabled      bool          `json:"debug_enabled"`
			NeedChannelUpdate bool          `json:"need_channel_update"`
			IsValid           bool          `json:"is_valid"`
			ConfiguredType    string        `json:"configured_type"`
			InfoTags          []interface{} `json:"info_tags"`
			DetectedType      string        `json:"detected_type"`
			RemoteAlias       string        `json:"remote_alias"`
			Branch            string        `json:"branch"`
			Owner             string        `json:"owner"`
			RepoName          string        `json:"repo_name"`
			Version           string        `json:"version"`
			RemoteVersion     string        `json:"remote_version"`
			CurrentHash       string        `json:"current_hash"`
			RemoteHash        string        `json:"remote_hash"`
			IsDirty           bool          `json:"is_dirty"`
			Detached          bool          `json:"detached"`
			CommitsBehind     []interface{} `json:"commits_behind"`
			GitMessages       []interface{} `json:"git_messages"`
			FullVersionString string        `json:"full_version_string"`
			Pristine          bool          `json:"pristine"`
		} `json:"moonraker"`
		Mainsail struct {
			Name           string   `json:"name"`
			Owner          string   `json:"owner"`
			Version        string   `json:"version"`
			RemoteVersion  string   `json:"remote_version"`
			ConfiguredType string   `json:"configured_type"`
			Channel        string   `json:"channel"`
			InfoTags       []string `json:"info_tags"`
		} `json:"mainsail"`
		Fluidd struct {
			Name           string        `json:"name"`
			Owner          string        `json:"owner"`
			Version        string        `json:"version"`
			RemoteVersion  string        `json:"remote_version"`
			ConfiguredType string        `json:"configured_type"`
			Channel        string        `json:"channel"`
			InfoTags       []interface{} `json:"info_tags"`
		} `json:"fluidd"`
		Klipper struct {
			Channel           string        `json:"channel"`
			DebugEnabled      bool          `json:"debug_enabled"`
			NeedChannelUpdate bool          `json:"need_channel_update"`
			IsValid           bool          `json:"is_valid"`
			ConfiguredType    string        `json:"configured_type"`
			InfoTags          []interface{} `json:"info_tags"`
			DetectedType      string        `json:"detected_type"`
			RemoteAlias       string        `json:"remote_alias"`
			Branch            string        `json:"branch"`
			Owner             string        `json:"owner"`
			RepoName          string        `json:"repo_name"`
			Version           string        `json:"version"`
			RemoteVersion     string        `json:"remote_version"`
			CurrentHash       string        `json:"current_hash"`
			RemoteHash        string        `json:"remote_hash"`
			IsDirty           bool          `json:"is_dirty"`
			Detached          bool          `json:"detached"`
			CommitsBehind     []struct {
				Sha     string      `json:"sha"`
				Author  string      `json:"author"`
				Date    string      `json:"date"`
				Subject string      `json:"subject"`
				Message string      `json:"message"`
				Tag     interface{} `json:"tag"`
			} `json:"commits_behind"`
			GitMessages       []interface{} `json:"git_messages"`
			FullVersionString string        `json:"full_version_string"`
			Pristine          bool          `json:"pristine"`
		} `json:"klipper"`
	} `json:"version_info"`
}

var MachineUpdateStatus = jsonrpc.Method[MachineUpdateStatusRequest, MachineUpdateStatusResponse]{Name: "machine.update.status"}

var MachineUpdateFull = jsonrpc.Method[struct{}, struct{}]{Name: "machine.update.full"}
var MachineUpdateMoonraker = jsonrpc.Method[struct{}, struct{}]{Name: "machine.update.moonraker"}
var MachineUpdateKlipper = jsonrpc.Method[struct{}, struct{}]{Name: "machine.update.klipper"}

type MachineUpdateClientRequest struct {
	Name string `json:"name"`
}

var MachineUpdateClient = jsonrpc.Method[MachineUpdateClientRequest, struct{}]{Name: "machine.update.client"}

var MachineUpdateSystem = jsonrpc.Method[struct{}, struct{}]{Name: "machine.update.system"}

type MachineUpdateRecoverRequest struct {
	Name string `json:"name"`
	Hard bool   `json:"hard"`
}

var MachineUpdateRecover = jsonrpc.Method[MachineUpdateRecoverRequest, struct{}]{Name: "machine.update.recover"}

// TODO: Power apis
// TODO: Led apis
// TODO: Led apis

type ServerHistoryListRequest struct {
	Limit  int     `json:"limit"`
	Start  int     `json:"start"`
	Since  float64 `json:"since"`
	Before float64 `json:"before"`
	Order  string  `json:"order"`
}

type ServerHistoryListResponse struct {
	Count int `json:"count"`
	Jobs  []struct {
		JobId        string  `json:"job_id"`
		Exists       bool    `json:"exists"`
		EndTime      float64 `json:"end_time"`
		FilamentUsed float64 `json:"filament_used"`
		Filename     string  `json:"filename"`
		Metadata     struct {
		} `json:"metadata"`
		PrintDuration float64 `json:"print_duration"`
		Status        string  `json:"status"`
		StartTime     float64 `json:"start_time"`
		TotalDuration float64 `json:"total_duration"`
	} `json:"jobs"`
}

var ServerHistoryList = jsonrpc.Method[ServerHistoryListRequest, ServerHistoryListResponse]{Name: "server.history.list"}

type ServerHistoryTotalsResponse struct {
	JobTotals struct {
		TotalJobs         int     `json:"total_jobs"`
		TotalTime         float64 `json:"total_time"`
		TotalPrintTime    float64 `json:"total_print_time"`
		TotalFilamentUsed float64 `json:"total_filament_used"`
		LongestJob        float64 `json:"longest_job"`
		LongestPrint      float64 `json:"longest_print"`
	} `json:"job_totals"`
}

var ServerHistoryTotals = jsonrpc.Method[struct{}, ServerHistoryTotalsResponse]{Name: "server.history.totals"}

type ServerHistoryResetTotalsResponse struct {
	LastTotals struct {
		TotalJobs         int     `json:"total_jobs"`
		TotalTime         float64 `json:"total_time"`
		TotalPrintTime    float64 `json:"total_print_time"`
		TotalFilamentUsed float64 `json:"total_filament_used"`
		LongestJob        float64 `json:"longest_job"`
		LongestPrint      float64 `json:"longest_print"`
	} `json:"last_totals"`
}

var ServerHistoryResetTotals = jsonrpc.Method[struct{}, ServerHistoryResetTotalsResponse]{Name: "server.history.reset_totals"}

type ServerHistoryGetJobRequest struct {
	Uid string `json:"uid"`
}

type ServerHistoryGetJobResponse struct {
	Job struct {
		JobId        string  `json:"job_id"`
		Exists       bool    `json:"exists"`
		EndTime      float64 `json:"end_time"`
		FilamentUsed float64 `json:"filament_used"`
		Filename     string  `json:"filename"`
		Metadata     struct {
		} `json:"metadata"`
		PrintDuration float64 `json:"print_duration"`
		Status        string  `json:"status"`
		StartTime     float64 `json:"start_time"`
		TotalDuration float64 `json:"total_duration"`
	} `json:"job"`
}

var ServerHistoryGetJob = jsonrpc.Method[ServerHistoryGetJobRequest, ServerHistoryGetJobResponse]{Name: "server.history.get_job"}

type ServerHistoryDeleteJobRequest struct {
	Uid string `json:"uid"`
}

type ServerHistoryDeleteJobResponse []string

var ServerHistoryDeleteJob = jsonrpc.Method[ServerHistoryDeleteJobRequest, ServerHistoryDeleteJobResponse]{Name: "server.history.delete_job"}

// TODO: MQTT apis

//TODO:
type NotifyGCodeResponseRequest []interface{}

var NotifyGCodeResponse = jsonrpc.Notify[NotifyGCodeResponseRequest]{Name: "notify_gcode_response"}

//TODO:
type NotifyStatusUpdateRequest []interface{}

var NotifyStatusUpdate = jsonrpc.Notify[NotifyStatusUpdateRequest]{Name: "notify_status_update"}

var NotifyKlippyReady = jsonrpc.Notify[struct{}]{Name: "notify_klippy_ready"}
var NotifyKlippyShutdown = jsonrpc.Notify[struct{}]{Name: "notify_klippy_shutdown"}
var NotifyKlippyDisconnected = jsonrpc.Notify[struct{}]{Name: "notify_klippy_disconnected"}

//TODO:
type NotifyFileListChangedRequest []interface{}

var NotifyFileListChanged = jsonrpc.Notify[NotifyFileListChangedRequest]{Name: "notify_filelist_changed"}

//TODO:
type NotifyUpdateResponseRequest []interface{}

var NotifyUpdateResponse = jsonrpc.Notify[NotifyUpdateResponseRequest]{Name: "notify_update_response"}

//TODO:
type NotifyUpdateRefreshedRequest []interface{}

var NotifyUpdateRefreshed = jsonrpc.Notify[NotifyUpdateRefreshedRequest]{Name: "notify_update_refreshed"}

//TODO:
type NotifyCpuThrottledRequest []interface{}

var NotifyCpuThrottled = jsonrpc.Notify[NotifyCpuThrottledRequest]{Name: "notify_cpu_throttled"}

type NotifyProcStatUpdateRequest []struct {
	MoonrakerStats struct {
		Time     float64 `json:"time"`
		CpuUsage float64 `json:"cpu_usage"`
		Memory   int     `json:"memory"`
		MemUnits string  `json:"mem_units"`
	} `json:"moonraker_stats"`
	CpuTemp float64 `json:"cpu_temp"`
	Network struct {
		Lo struct {
			RxBytes   int     `json:"rx_bytes"`
			TxBytes   int     `json:"tx_bytes"`
			Bandwidth float64 `json:"bandwidth"`
		} `json:"lo"`
		Wlan0 struct {
			RxBytes   int     `json:"rx_bytes"`
			TxBytes   int     `json:"tx_bytes"`
			Bandwidth float64 `json:"bandwidth"`
		} `json:"wlan0"`
	} `json:"network"`
	SystemCpuUsage struct {
		Cpu  float64 `json:"cpu"`
		Cpu0 float64 `json:"cpu0"`
		Cpu1 float64 `json:"cpu1"`
		Cpu2 float64 `json:"cpu2"`
		Cpu3 float64 `json:"cpu3"`
	} `json:"system_cpu_usage"`
	WebsocketConnections int `json:"websocket_connections"`
}

var NotifyProcStatUpdate = jsonrpc.Notify[NotifyProcStatUpdateRequest]{Name: "notify_proc_stat_update"}

//TODO:
type NotifyHistoryChangedRequest []interface{}

var NotifyHistoryChanged = jsonrpc.Notify[NotifyHistoryChangedRequest]{Name: "notify_history_changed"}

//TODO:
type NotifyUserCreatedRequest []interface{}

var NotifyUserCreated = jsonrpc.Notify[NotifyUserCreatedRequest]{Name: "notify_user_created"}

//TODO:
type NotifyUserDeletedRequest []interface{}

var NotifyUserDeleted = jsonrpc.Notify[NotifyUserDeletedRequest]{Name: "notify_user_deleted"}

//TODO:
type NotifyServiceStateChangedRequest []interface{}

var NotifyServiceStateChanged = jsonrpc.Notify[NotifyServiceStateChangedRequest]{Name: "notify_service_state_changed"}

//TODO:
type NotifyJobQueueChangedRequest []interface{}

var NotifyJobQueueChanged = jsonrpc.Notify[NotifyJobQueueChangedRequest]{Name: "notify_job_queue_changed"}

// TODO: notify apis
