package model

type NvmeOfGatewayName []struct {
	Hostname    string `json:"hostname"`
	Daemon_name string `json:"daemon_name"`
}
type NvmeOfSubSystemList struct {
	ErrorMessage string `json:"error_message"`
	Subsystems   []struct {
		Nqn            string `json:"nqn"`
		SerialNumber   string `json:"serial_number"`
		ModelNumber    string `json:"model_number"`
		MinCntlid      int    `json:"min_cntlid"`
		MaxCntlid      int    `json:"max_cntlid"`
		NamespaceCount int    `json:"namespace_count"`
		Subtype        string `json:"subtype"`
		EnableHa       bool   `json:"enable_ha"`
	} `json:"subsystems"`
	Status int `json:"status"`
} //@name NvmeOfSubSystemList

type NvmeOfNameSpaceList struct {
	ErrorMessage string `json:"error_message"`
	SubsystemNqn string `json:"subsystem_nqn"`
	Namespaces   []struct {
		Nsid               int    `json:"nsid"`
		BdevName           string `json:"bdev_name"`
		RbdImageName       string `json:"rbd_image_name"`
		RbdPoolName        string `json:"rbd_pool_name"`
		BlockSize          int    `json:"block_size"`
		RbdImageSize       string `json:"rbd_image_size"`
		UUID               string `json:"uuid"`
		LoadBalancingGroup int    `json:"load_balancing_group"`
		RwIosPerSecond     string `json:"rw_ios_per_second"`
		RwMbytesPerSecond  string `json:"rw_mbytes_per_second"`
		RMbytesPerSecond   string `json:"r_mbytes_per_second"`
		WMbytesPerSecond   string `json:"w_mbytes_per_second"`
	} `json:"namespaces"`
	Status int `json:"status"`
} //@name NvmeOfNameSpaceList

type NvmeOfTargetVerify struct {
	Genctr  int `json:"genctr"`
	Records []struct {
		Trtype  string `json:"trtype"`
		Adrfam  string `json:"adrfam"`
		Subtype string `json:"subtype"`
		Treq    string `json:"treq"`
		Portid  int    `json:"portid"`
		Trsvcid string `json:"trsvcid"`
		Subnqn  string `json:"subnqn"`
		Traddr  string `json:"traddr"`
		Eflags  string `json:"eflags"`
		Sectype string `json:"sectype"`
	} `json:"records"`
} //@name NvmeOfTargetVerify

type NvmeOfList struct {
	Devices []struct {
		HostNQN    string `json:"HostNQN"`
		HostID     string `json:"HostID"`
		Subsystems []struct {
			Subsystem    string `json:"Subsystem"`
			SubsystemNQN string `json:"SubsystemNQN"`
			Controllers  []struct {
				Controller   string `json:"Controller"`
				SerialNumber string `json:"SerialNumber"`
				ModelNumber  string `json:"ModelNumber"`
				Firmware     string `json:"Firmware"`
				Transport    string `json:"Transport"`
				Address      string `json:"Address"`
				Namespaces   []any  `json:"Namespaces"`
				Paths        []struct {
					Path     string `json:"Path"`
					ANAState string `json:"ANAState"`
				} `json:"Paths"`
			} `json:"Controllers"`
			Namespaces []struct {
				NameSpace    string `json:"NameSpace"`
				Generic      string `json:"Generic"`
				Nsid         int    `json:"NSID"`
				UsedBytes    int64  `json:"UsedBytes"`
				MaximumLBA   int    `json:"MaximumLBA"`
				PhysicalSize int64  `json:"PhysicalSize"`
				SectorSize   int    `json:"SectorSize"`
			} `json:"Namespaces"`
		} `json:"Subsystems"`
	} `json:"Devices"`
} // @name NvmeOfList

type HostNvmeOfList struct {
	Hostname string `json:"Hostname"`
	Detail   NvmeOfList
} // @name HostNvmeOfList

type NvmeOfPath struct {
	Hostname string `json:"Hostname"`
	Devices  []struct {
		NameSpace    int    `json:"NameSpace"`
		DevicePath   string `json:"DevicePath"`
		GenericPath  string `json:"GenericPath"`
		Firmware     string `json:"Firmware"`
		ModelNumber  string `json:"ModelNumber"`
		SerialNumber string `json:"SerialNumber"`
		UsedBytes    int64  `json:"UsedBytes"`
		MaximumLBA   int    `json:"MaximumLBA"`
		PhysicalSize int64  `json:"PhysicalSize"`
		SectorSize   int    `json:"SectorSize"`
	} `json:"Devices"`
}

type NvmeOfServiceCreate struct {
	ServiceType string          `yaml:"service_type"`
	ServiceId   string          `yaml:"service_id"`
	Placement   NvmeOfPlacement `yaml:"placement"`
	Spec        NvmeOfSpec      `yaml:"spec"`
}
type NvmeOfPlacement struct {
	Hosts []string `yaml:"hosts"`
}
type NvmeOfSpec struct {
	Pool            string `yaml:"pool"`
	TgtCmdExtraArgs string `yaml:"tgt_cmd_extra_args"`
}
type NvmeOfConnection []struct {
	Cntlid      int    `json:"cntlid"`
	Hostnqn     string `json:"hostnqn"`
	Hostid      string `json:"hostid"`
	NumIoQpairs int    `json:"num_io_qpairs"`
}
type NvmeOfTarget []struct {
	Nqn             string `json:"nqn"`
	Subtype         string `json:"subtype"`
	ListenAddresses []struct {
		Transport string `json:"transport"`
		Trtype    string `json:"trtype"`
		Adrfam    string `json:"adrfam"`
		Traddr    string `json:"traddr"`
		Trsvcid   string `json:"trsvcid"`
	} `json:"listen_addresses"`
	AllowAnyHost  bool   `json:"allow_any_host"`
	Hosts         []any  `json:"hosts"`
	SerialNumber  string `json:"serial_number"`
	ModelNumber   string `json:"model_number"`
	MaxNamespaces int    `json:"max_namespaces"`
	MinCntlid     int    `json:"min_cntlid"`
	MaxCntlid     int    `json:"max_cntlid"`
	Namespaces    []struct {
		Nsid           int    `json:"nsid"`
		BdevName       string `json:"bdev_name"`
		Name           string `json:"name"`
		Nguid          string `json:"nguid"`
		UUID           string `json:"uuid"`
		Rbd_image_name string `json:"rbd_image_name"`
		Rbd_pool_name  string `json:"rbd_pool_name"`
		Block_size     int    `json:"block_size"`
		Rbd_image_size string `json:"rbd_image_size"`
	} `json:"namespaces"`
	Session int `json:"session"`
} //@name NvmeOfTarget
