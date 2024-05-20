package hadoop

import (
	"encoding/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type (
	MeasureType   string // 计量类型
	BandwidthType string // 带宽方向
	Granularity   string // 带宽粒度
	AdjustParam   string // 带宽处理
)

const (
	MeasureTypePeak95 MeasureType = "peak95" // 95峰值
	MeasureTypePeak   MeasureType = "peak"   // 最高峰值

	BandwidthTypeOutMinusIn BandwidthType = "out_minus_in" // 上行减下行
	BandwidthTypeOut        BandwidthType = "out"          // 上行

	Granularity5Min Granularity = "5min" // 5分钟
	GranularityHour Granularity = "hour" // 小时

	AdjustParam302PeakClipping AdjustParam = "302PeakClipping" // 302削峰处理
	AdjustParamNetbench        AdjustParam = "netbench"        // 压测处理
	AdjustParamScale           AdjustParam = "scale"           // 进制处理
	AdjustParamMonitor         AdjustParam = "monitor"         // 分析数据 90% 处理，仅在分析数据接口生效（目前只有前端调用，OpenAPI 分析数据需要剔除这个处理）
	AdjustParamAll             AdjustParam = "all"             // 所有处理皆生效
)

var (
	// 带宽不做任何处理
	AdjustParamsEmpty = []AdjustParam{}
	// 带宽所有处理皆生效
	AdjustParamsAll = []AdjustParam{AdjustParamAll}
	// 除分析数据 90% 处理，其他皆生效
	AdjustParamsNoMonitor = []AdjustParam{AdjustParam302PeakClipping, AdjustParamNetbench, AdjustParamScale}
	// 带宽除进制外，其他处理皆生效
	AdjustParamsNoScale = []AdjustParam{AdjustParam302PeakClipping, AdjustParamNetbench}
)

func (t MeasureType) String() string {
	return string(t)
}
func (t BandwidthType) String() string {
	return string(t)
}
func (t Granularity) String() string {
	return string(t)
}

type (
	// HadoopResp 大数据接口公共 resp
	HadoopResp[T any] struct {
		Code  int    `json:"code"`
		Data  T      `json:"data"`
		Error string `json:"error"`
	}

	// HadoopResp 大数据接口公共 resp
	HadoopResponse struct {
		Code  int              `json:"code"`
		Data  []BaseConfigItem `json:"data"`
		Error string           `json:"error"`
	}

	// MeasureCommonData 计量公用的节点级别的 data
	MeasureCommonDataPerNode struct {
		NodeId   string                       `json:"nodeId"`
		DayPeak  map[string]MeasureCommonUnit `json:"dayPeak"`  // 日带宽值(一天中的95值或峰值)
		BillPeak MeasureCommonUnit            `json:"billPeak"` // 计费天数的区间带宽值(区间内的95值或峰值)，目前账单计费用此字段
		Peak     MeasureCommonUnit            `json:"peak"`     // 区间带宽值(区间内的95值或峰值)
	}

	// MeasureCommonData 计量公用 data
	MeasureCommonData struct {
		Id       string                       `bson:"_id,omitempty"  json:"id"`
		DayPeak  map[string]MeasureCommonUnit `bson:"dayPeak,omitempty" json:"dayPeak"`   // 日带宽值(一天中的95值或峰值或晚高峰95)
		BillPeak MeasureCommonUnit            `bson:"billPeak,omitempty" json:"billPeak"` // 计费天数的区间带宽值(区间内的95值或峰值)，目前账单计费用此字段
		Peak     MeasureCommonUnit            `bson:"peak,omitempty" json:"peak"`         // 区间带宽值(区间内的95值或峰值)
	}
	MeasureCommonUnit struct {
		Bandwidth int64 `json:"bandwidth"` // 计量值，单位：bps
		Time      int64 `json:"time"`      // 当前计量值所处时刻，时间戳，秒
	}
	MeasureCommonUnitList struct {
		Id                    string               `bson:"_id" json:"id"`
		MeasureCommonUnitList []*MeasureCommonUnit `bson:"measureCommonUnitList" json:"measureCommonUnitList"`
	}

	GetMomentNode struct {
		NodeId string `json:"nodeId"` // 节点 Id
		Time   int64  `json:"time"`   // 时刻值，时间戳，秒
	}

	// NodeMomentData 节点时刻公用 data
	NodeMomentData struct {
		NodeId    string `bson:"nodeId" json:"nodeId"`
		Bandwidth int64  `bson:"bandwidth" json:"bandwidth"` // 计量值，单位：bps
		Time      int64  `bson:"time" json:"time"`           // 时刻值，时间戳，秒
	}

	NodeMomentDataList struct {
		Id                 string            `bson:"_id" json:"id"`
		NodeMomentDataList []*NodeMomentData `bson:"nodeMomentDataList" json:"nodeMomentDataList,omitempty"`
	}

	TimeQuantum struct {
		Position  int `json:"position"`  // 取第多少个点
		StartHour int `json:"startHour"` // 区间开始时间，例如：18
		EndHour   int `json:"endHour"`   // 区间结束时间，例如：23(23:59 在 23 范围内，闭区间)
	}

	// MeasureMetricNode 业务侧计量节点详情
	MeasureMetricNode struct {
		NodeId string                 `json:"nodeId"`           // 节点 Id
		Ranges []*NodeMetricInfoRange `json:"ranges,omitempty"` // 计量排除时间范围
	}

	// MeasureMetricPerNode 业务侧计量节点详情, 用于分节点返回数据的请求使用
	MeasureMetricPerNode struct {
		BillDays int64                  `json:"billDays"`         // 计费天数
		NodeId   string                 `json:"nodeId"`           // 节点 Id
		Ranges   []*NodeMetricInfoRange `json:"ranges,omitempty"` // 计量排除时间范围
	}

	// NodeMetricInfoRange 时间范围
	NodeMetricInfoRange struct {
		Start int64 `json:"start"` // 范围起期，时间戳，秒
		End   int64 `json:"end"`   // 范围止期，时间戳，秒
	}
)

func (m *MeasureCommonData) DayPeak95Ts() (map[int64]uint64, error) {
	dm := make(map[int64]uint64)
	for day, peak95 := range m.DayPeak {
		d, err := time.ParseInLocation("2006-01-02", day, time.Local)
		if err != nil {
			logx.Errorf("parse api result err: %v", err)
			return nil, err
		}
		dm[d.Unix()] = uint64(peak95.Bandwidth)
	}
	return dm, nil
}

// ========================================================
type (
	// GetBaseConfigWhiteListReq 查询供应侧进制白名单参数
	GetBaseConfigWhiteListReq struct {
		VendorId uint32 `form:"vendorId,omitempty"` // 供应商 Id
	}

	// BaseConfigWhiteListItem 进制白名单
	BaseConfigWhiteListItem struct {
		Id         int    `json:"id"`         // db_id，仅作为查询时，存在值，update 时可为空。
		VendorId   uint32 `json:"vendorId"`   // 供应商 Id，upsert 时，vendorId 存在白名单中，即为更新，不在白名单中即为新增
		Day        string `json:"day"`        // 生效时间，格式：2006-01-02
		CreateUser string `json:"createUser"` // 创建人
		UpdateUser string `json:"updateUser"` // 修改人
		CreateTime string `json:"createTime"` // 创建时间，格式：2006-01-02T15:04:05.000+0000
		UpdateTime string `json:"updateTime"` // 修改时间，格式：2006-01-02T15:04:05.000+0000
	}

	// UpsertBaseConfigWhiteListReq 修改供应侧进制白名单参数
	UpsertBaseConfigWhiteListReq struct {
		Operator    string                    `json:"operator"`         // 操作人
		BaseConfigs []BaseConfigWhiteListItem `json:"whitelistConfigs"` // 修改参数
	}

	// DeleteBaseConfigWhiteListReq 删除供应商进制白名单参数
	DeleteBaseConfigWhiteListReq struct {
		Operator  string   `json:"operator"`  // 操作人
		VendorIds []uint32 `json:"vendorIds"` // 供应商 Id
	}
)

// =====================================

// BaseConfigItems 供应侧进制配置
type (
	BaseConfigItems struct {
		Data []BaseConfigItem
	}

	BaseConfigItem struct {
		RecordId    int64   `json:"id"`          // db_id
		CustomerId  uint32  `json:"uid"`         // 签约方 Id
		VendorBase  int64   `json:"vendorBase"`  // 进制，格式为 bps 转换到 Gbps 所需进制
		Day         string  `json:"day"`         // 生效时间，格式：2006-01-02
		Coefficient float64 `json:"coefficient"` // 系数，示例： 1.0
	}

	// GetBaseConfig 查询供应侧进制配置参数
	GetBaseConfigReq struct {
		Uid   uint32 `form:"uid"`   // 签约方 Id
		Start string `form:"start"` // 开始时间，格式：2006-01-02，东八区
		End   string `form:"end"`   // 结束时间，格式：2006-01-02，东八区
	}

	// UpsertBaseConfigReq 修改供应侧进制配置参数
	UpsertBaseConfigReq struct {
		BaseConfigs []BaseConfigItem `json:"baseConfigs"`
	}

	// DeleteBaseConfigReq 删除供应侧进制配置参数
	DeleteBaseConfigReq []int64

	VendorBaseConfigItem struct {
		VendorId    uint32  `json:"vendorId"`    // 供应商 Id
		Base        int64   `json:"base"`        // 进制，格式为 bps 转换到 Gbps 所需进制
		Coefficient float64 `json:"coefficient"` // 系数，示例：1.0
		EffectDay   string  `json:"effectDay"`   // 生效时间，格式：2006-01-02
		ExpiryDay   string  `json:"expiryDay"`   // 失效时间，格式：2006-01-02
	}
	GetVendorBaseConfigListReq struct {
		VendorIds []uint32 `json:"vendorIds,optional"` // 供应商 Ids
	}
	UpdateBaseVendorConfigReq struct {
		BaseConfigs []*UpdateBaseVendorConfigItem `json:"baseConfigs"`
		Operator    string                        `json:"operator"` // 操作人
	}
	UpdateBaseVendorConfigItem struct {
		VendorId    uint32  `json:"vendorId"`             // 供应商 Id
		Base        int64   `json:"base,optional"`        // 进制，格式为 bps 转换到 Gbps 所需进制
		Coefficient float64 `json:"coefficient,optional"` // 系数，示例：1.0
		EffectDay   string  `json:"effectDay"`            // 生效时间，格式：2006-01-02
		ExpiryDay   string  `json:"expiryDay,optional"`   // 失效时间，格式：2006-01-02
	}
	GetVendorBaseConfigHistoryReq struct {
		VendorIds []uint32 `json:"vendorIds,optional"` // 供应商 Ids
	}
	VendorBaseConfigItemHistory struct {
		VendorId    uint32  `json:"vendorId"`    // 供应商 Id
		Base        int64   `json:"base"`        // 进制，格式为 bps 转换到 Gbps 所需进制
		Coefficient float64 `json:"coefficient"` // 系数，示例：1.0
		EffectDay   string  `json:"effectDay"`   // 生效时间，格式：2006-01-02
		ExpiryDay   string  `json:"expiryDay"`   // 失效时间，格式：2006-01-02
		CreateTime  string  `json:"createTime"`  // 创建时间，格式：2006-01-02T15:04:05.000+0000
		UpdateTime  string  `json:"updateTime"`  // 更新时间，格式：2006-01-02T15:04:05.000+0000
		CreateUser  string  `json:"createUser"`  // 创建人
		UpdateUser  string  `json:"updateUser"`  // 更新人
	}

	NodeBaseConfigItem struct {
		VendorId    uint32  `json:"vendorId"`    // 供应商 Id
		NodeId      string  `json:"nodeId"`      // 节点 Id
		Base        int64   `json:"base"`        // 进制，格式为 bps 转换到 Gbps 所需进制
		EffectDay   string  `json:"effectDay"`   // 生效时间，格式：2006-01-02
		ExpiryDay   string  `json:"expiryDay"`   // 失效时间，格式：2006-01-02
		Coefficient float64 `json:"coefficient"` // 系数，示例：1.0
	}
	GetBaseNodeConfigListReq struct {
		VendorIds []uint32 `json:"vendorIds,optional"` // 供应商 Ids
		NodeIds   []string `json:"nodeIds,optional"`   // 节点 Ids
	}
	UpdateBaseNodeConfigReq struct {
		BaseConfigs []*UpdateBaseNodeConfigItem `json:"baseConfigs"`
		Operator    string                      `json:"operator"` // 操作人
	}
	UpdateBaseNodeConfigItem struct {
		VendorId    uint32  `json:"vendorId"`             // 供应商 Id
		NodeId      string  `json:"nodeId"`               // 节点 Id
		Base        int64   `json:"base,optional"`        // 进制，格式为 bps 转换到 Gbps 所需进制
		Coefficient float64 `json:"coefficient,optional"` // 系数，示例：1.0
		EffectDay   string  `json:"effectDay"`            // 生效时间，格式：2006-01-02
		ExpiryDay   string  `json:"expiryDay,optional"`   // 失效时间，格式：2006-01-02
	}

	GetBaseNodeConfigHistoryReq struct {
		VendorIds []uint32 `json:"vendorIds,optional"` // 供应商 Ids
		NodeIds   []string `json:"nodeIds,optional"`   // 节点 Ids
	}
	NodeBaseConfigItemHistory struct {
		VendorId    uint32  `json:"vendorId"`    // 供应商 Id
		NodeId      string  `json:"nodeId"`      // 节点 Id
		Base        int64   `json:"base"`        // 进制，格式为 bps 转换到 Gbps 所需进制
		EffectDay   string  `json:"effectDay"`   // 生效时间，格式：2006-01-02
		ExpiryDay   string  `json:"expiryDay"`   // 失效时间，格式：2006-01-02
		Coefficient float64 `json:"coefficient"` // 系数，示例：1.0
		CreateTime  string  `json:"createTime"`  // 创建时间，格式：2006-01-02T15:04:05.000+0000
		UpdateTime  string  `json:"updateTime"`  // 更新时间，格式：2006-01-02T15:04:05.000+0000
		CreateUser  string  `json:"createUser"`  // 创建人
		UpdateUser  string  `json:"updateUser"`  // 更新人
	}
)

func (b *BaseConfigItems) MarshalBinary() (data []byte, err error) {
	if b == nil {
		return []byte{}, nil
	}
	return json.Marshal(b)
}
func (b *BaseConfigItems) UnMarshalBinary(data []byte) error {
	return json.Unmarshal(data, b)
}

// ========================================================
type (
	// GetCustomerMeasureReq 业务侧计量接口参数
	GetCustomerMeasureReq struct {
		SignatoryIds  []uint32             `json:"signatoryIds"`  // 签约方 Ids
		NodeType      string               `json:"nodeType"`      // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		Start         string               `json:"start"`         // 开始时间，格式：2006-01-02，东八区
		End           string               `json:"end"`           // 结束时间，格式：2006-01-02，东八区
		BillDays      int                  `json:"billDays"`      // 计费天数
		MeasureType   string               `json:"measureType"`   // 计量算法，peak：峰值 | peak95: 95峰值
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		DataSource    string               `json:"dataSource"`    // 计量数据源。process：进程；netcard：网卡。如果不传值，大数据使用网卡数据源
	}

	// GetCustomerMeasurePerNodeReq 业务侧计量接口参数,用于分节点返回数据使用，不合并
	GetCustomerMeasurePerNodeReq struct {
		Uid           uint32               `json:"uid"`           // 业务方 Id
		NodeType      string               `json:"nodeType"`      // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		Start         string               `json:"start"`         // 开始时间，格式：2006-01-02，东八区
		End           string               `json:"end"`           // 结束时间，格式：2006-01-02，东八区
		BillDays      int                  `json:"billDays"`      // 计费天数
		MeasureType   string               `json:"measureType"`   // 计量算法，peak：峰值 | peak95: 95峰值
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		DataSource    string               `json:"dataSource"`    // 计量数据源。process_ds：进程；netcard_ds：网卡。如果不传值，大数据使用网卡数据源
	}

	// GetCustomerMomentReq 业务侧获取时刻带宽值参数
	GetCustomerMomentReq struct {
		SignatoryIds  []uint32         `json:"signatoryIds"`  // 签约方 Ids
		NodeType      string           `json:"nodeType"`      // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes         []*GetMomentNode `json:"nodes"`         // 批量查询参数
		BandwidthType string           `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		DataSource    string           `json:"dataSource"`    // 计量数据源，process：进程；netcard：网卡
	}

	// GetCustomerMeasure5MinReq 业务侧计量 5min 数据明细参数
	GetCustomerMeasure5MinReq struct {
		SignatoryIds  []uint32             `json:"signatoryIds"`  // 签约方 Ids
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		NodeType      string               `json:"nodeType"`      // 节点类型，node | switch
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		Start         int64                `json:"start"`         // 开始时间，时间戳，秒
		End           int64                `json:"end"`           // 结束时间，时间戳，秒
		DataSource    string               `json:"dataSource"`    // 计量数据源。process：进程；netcard：网卡。如果不传值，大数据使用网卡数据源
	}

	// GetCustomerTimeQuantumSortReq 业务侧时间段内指定峰值
	GetCustomerTimeQuantumSortReq struct {
		Quantum       TimeQuantum          `json:"eveningPeak"`   // 每天高峰时间区间
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		NodeType      string               `json:"nodeType"`      // 节点类型，node | switch
		Uid           uint32               `json:"uid"`           // 业务方 Id
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		Start         string               `json:"start"`         // 开始时间，格式：2006-01-02
		End           string               `json:"end"`           // 结束时间，格式：2006-01-02
	}
	GetCustomerTimeQuantumSortResp struct {
		DayQuantumMap map[string]*MeasureCommonUnit `json:"dayPeak"` // 返回天粒度的区间指定峰值
	}

	// ========================================================

	// GetVendorMeasurePerNodeReq 供应侧计量接口参数,用于分节点返回数据使用，不合并
	GetVendorMeasurePerNodeReq struct {
		Vendor       uint32                  `json:"vendor"`                // 业务方 Id
		NodeType     string                  `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes        []*MeasureMetricPerNode `json:"nodes"`                 // 节点列表
		Start        string                  `json:"start"`                 // 开始时间，格式：2006-01-02，东八区
		End          string                  `json:"end"`                   // 结束时间，格式：2006-01-02，东八区
		MeasureType  string                  `json:"measureType"`           // 计量算法，peak：峰值 | peak95: 95峰值
		AdjustParams []AdjustParam           `json:"adjustParams,optional"` // 带宽处理
	}

	// GetVendorMeasureReq 供应侧计量接口参数
	GetVendorMeasureReq struct {
		Vendor       uint32               `json:"vendor"`                // 业务方 Id
		NodeType     string               `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes        []*MeasureMetricNode `json:"nodes"`                 // 节点列表
		Start        string               `json:"start"`                 // 开始时间，格式：2006-01-02，东八区
		End          string               `json:"end"`                   // 结束时间，格式：2006-01-02，东八区
		BillDays     int64                `json:"billDays"`              // 计费天数
		MeasureType  string               `json:"measureType"`           // 计量算法，peak：峰值 | peak95: 95峰值
		AdjustParams []AdjustParam        `json:"adjustParams,optional"` // 带宽处理
		DataSource   string               `json:"dataSource,optional"`   // 计量数据源
	}

	// VendorNodeMetricInfoRange 供应侧节点范围
	VendorNodeMetricInfoRange struct {
		Start int64 `json:"start"` // 范围起期，时间戳，秒
		End   int64 `json:"end"`   // 范围止期，时间戳，秒
	}

	// GetVendorMomentReq 供应侧获取时刻带宽值参数
	GetVendorMomentReq struct {
		VendorId     uint32           `json:"vendor"`                // 供应商 Id
		NodeType     string           `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes        []*GetMomentNode `json:"nodes"`                 // 批量查询参数
		AdjustParams []AdjustParam    `json:"adjustParams,optional"` // 带宽处理
	}

	// GetVendorMeasure5MinReq 供应侧计量 5min 数据明细参数
	GetVendorMeasure5MinReq struct {
		Nodes        []*MeasureMetricNode `json:"nodes"`                 // 节点列表
		NodeType     string               `json:"nodeType"`              // 节点类型，node | switch
		Vendor       uint32               `json:"vendor"`                // 供应商 Id
		AdjustParams []AdjustParam        `json:"adjustParams,optional"` // 带宽处理
		Start        int64                `json:"start"`                 // 开始时间，时间戳，秒
		End          int64                `json:"end"`                   // 结束时间，时间戳，秒
		DataSource   string               `json:"dataSource,optional"`   // 计量数据源
	}

	// GetVendorTimeQuantumSortReq 供应侧时间段内指定峰值
	GetVendorTimeQuantumSortReq struct {
		Quantum      TimeQuantum          `json:"eveningPeak"`           // 每天高峰时间区间
		Nodes        []*MeasureMetricNode `json:"nodes"`                 // 节点列表
		NodeType     string               `json:"nodeType"`              // 节点类型，node | switch
		Uid          uint32               `json:"vendor"`                // 供应商 Id
		AdjustParams []AdjustParam        `json:"adjustParams,optional"` // 带宽处理
		Start        string               `json:"start"`                 // 开始时间，格式：2006-01-02
		End          string               `json:"end"`                   // 结束时间，格式：2006-01-02
	}
	GetVendorTimeQuantumSortResp struct {
		DayQuantumMap map[string]*MeasureCommonUnit `json:"dayPeak"` // 返回天粒度的区间指定峰值
	}

	// GetNodeAnalyzeBandwithReq 查询节点 5min 分析数据
	GetNodeAnalyzeBandwithReq struct {
		VendorId     uint32        `json:"vendorId"`              // 供应商 Id
		NodeId       string        `json:"nodeId"`                // 节点id
		NodeType     string        `json:"nodeType"`              // node：单机 / switch：交换机 不传值默认为单机节点
		Granularity  Granularity   `json:"granularity"`           // 粒度 5min/hour
		AdjustParams []AdjustParam `json:"adjustParams,optional"` // 带宽处理
		Start        string        `json:"start"`                 // yyyyMMddHHmm
		End          string        `json:"end"`                   // yyyyMMddHHmm
	}
	AnalyzeBandwithItem struct {
		Time int64 `json:"time"` // 时间戳，秒
		Up   int64 `json:"up"`   // out 带宽，单位：bps
		Down int64 `json:"down"` // in 带宽，单位：bps
	}

	GetNodesAnalyzeBandwithReq struct {
		VendorIds   []string    `json:"vendors,optional"`    // 供应商 Id
		CustomerId  string      `json:"customerId,optional"` // 接入客户 快手、b站等客户id。注意这里customerId就是指的客户id，因为uid很早就被定义为业务id
		Uid         string      `json:"uid,optional"`        // 接入业务 快手、b站等业务id。注意这里uid就是指的业务id，因为uid很早就被定义为业务id
		NodeIds     []string    `json:"nodeIds"`             // 多个节点id
		NodeType    string      `json:"nodeType"`            // node：单机 / switch：交换机 不传值默认为单机节点
		Granularity Granularity `json:"granularity"`         // 粒度 5min/hour
		Start       string      `json:"start"`               // 开始时间 yyyyMMddHHmm 闭区间
		End         string      `json:"end"`                 // 结束时间 yyyyMMddHHmm 开区间
		DataSource  string      `json:"dataSource"`          // 计量数据源。如果不传值，大数据使用网卡数据源
	}
	AnalyzeNodeBandwithItem struct {
		NodeId         string                 `json:"nodeId"`    // 节点id
		BandwidthItems []*AnalyzeBandwithItem `json:"bandwidth"` // 带宽明细
	}

	GetNodeLineBandwithSeriesReq struct {
		VendorId uint32 `json:"vendorId,optional"` // 供应商 Id
		NodeId   string `json:"nodeId"`            // 节点id
		Start    int64  `json:"start"`             // 开始时间，闭区间，时间戳，单位：秒
		End      int64  `json:"end"`               // 结束时间，开区间，时间戳，单位：秒
	}

	GetNodeLineMomentBandwithReq struct {
		VendorId uint32   `json:"vendorId,optional"` // 供应商 Id
		NodeIds  []string `json:"nodeIds"`           // 多个节点id。大数据接口nodeIds size最大为100
		Time     int64    `json:"time"`              // 时间戳，单位：秒
	}
	NodeLineBandwithItem struct {
		Line      string `json:"line"`      // 线路名
		Bandwidth int64  `json:"bandwidth"` // 带宽，单位：bps
	}

	// NodeAnalyseMetricReq 节点分析数据区间指标接口参数
	NodeAnalyseMetricReq struct {
		Start        string        `json:"start"`                 // 开始时间 yyyyMMddHHmm 闭区间
		End          string        `json:"end"`                   // 结束时间 yyyyMMddHHmm 开区间
		NodeId       string        `json:"nodeId"`                // 节点id
		NodeType     string        `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		AdjustParams []AdjustParam `json:"adjustParams,optional"` // 带宽处理
	}
	NodeAnalyseMetricResp struct {
		Peak95 int64 `json:"peak95"` // 分析数据 区间 95值
	}

	GetNodesDatesAnalyseBandwidthReq struct {
		NodeIds []string `json:"nodeIds"` //节点id列表
		Start   int64    `json:"start"`   //开始时间，秒时间戳，闭区间
		End     int64    `json:"end"`     //结束时间，秒时间戳，开区间
	}

	GetNodesDatesAnalyseBandwidthItem struct {
		NodeId string `json:"nodeId"`
		Peak95 int64  `json:"peak95"`
		Time   int64  `json:"time"`
	}

	// GetDataSourceConfig 查询签约方计量数据源生效日列表
	GetDataSourceConfigReq struct {
		Uids []uint32 `json:"uids,optional"`
	}

	DataSourceDetail struct {
		DataSourceConfig
		Id         int    `json:"id"` //数据库表自增id
		CreateTime string `json:"createTime"`
		UpdateTime string `json:"updateTime"`
		CreateUser string `json:"createUser"`
		UpdateUser string `json:"updateUser"`
	}

	DataSourceConfig struct {
		Uid        uint32 `json:"uid"`        //签约方id
		Day        string `json:"day"`        //生效日 包含当天，格式2022-01-01
		DataSource string `json:"dataSource"` // 计量数据源 3种枚举值 netcard：网卡 / process：进程 / deretran_process：进程去重传
	}

	// 添加or更新签约方计量数据源生效日配置
	UpdateDataSourceConfigReq struct {
		Operator          string             `json:"operator"`
		DataSourceConfigs []DataSourceConfig `json:"dataSourceConfigs"`
	}

	// 查询节点指定时间的各业务方的带宽值
	GetNodesMomentsCustomerBandwidthReq struct {
		NodeType string                               `json:"nodeType"` // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes    []*NodesMomentsCustomerBandwidthItem `json:"nodes"`
	}
	NodesMomentsCustomerBandwidthItem struct {
		NodeId string `json:"nodeId"` // 节点 Id
		Time   int64  `json:"time"`   // 时间戳，秒，5min 粒度
	}

	GetNodesMomentsCustomerBandwidthResult struct {
		NodeId            string              `json:"nodeId"`        // 节点 Id
		Time              int64               `json:"time"`          // 时间戳，对应 request 中的 Time
		CustomerBandwiths []*MomentCustomerBw `json:"cidBandwidths"` // 业务方的带宽值
	}
	MomentCustomerBw struct {
		CustomerId string `json:"customerId"` // 业务方 Id
		Bandwith   int64  `json:"bandwidth"`  // 带宽值，bps
	}
)
