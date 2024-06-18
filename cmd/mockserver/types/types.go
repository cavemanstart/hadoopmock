package types

import "time"

type AdjustParam string // 带宽处理
type (
	//计量数据格式

	// HadoopResp 大数据接口公共 resp
	HadoopResp struct {
		Code  int         `json:"code"`
		Data  interface{} `json:"data"`
		Error string      `json:"error"`
	}

	// MeasureCommonData 计量公用的节点级别的 data
	MeasureCommonDataNodes struct {
		NodeList []*MeasureCommonDataPerNode `bson:"nodeList,omitempty"  json:"nodeList"`
	}

	MeasureCommonDataPerNode struct {
		NodeId   string                        `bson:"nodeId" json:"nodeId"`
		DayPeak  map[string]*MeasureCommonUnit `bson:"dayPeak" json:"dayPeak"`   // 日带宽值(一天中的95值或峰值)
		BillPeak MeasureCommonUnit             `bson:"billPeak" json:"billPeak"` // 计费天数的区间带宽值(区间内的95值或峰值)，目前账单计费用此字段
		Peak     MeasureCommonUnit             `bson:"peak" json:"peak"`         // 区间带宽值(区间内的95值或峰值)
	}
	// MeasureCommonData 计量公用 data
	MeasureCommonData struct {
		DayPeak  map[string]*MeasureCommonUnit `bson:"dayPeak,omitempty" json:"dayPeak"`   // 日带宽值(一天中的95值或峰值或晚高峰95)
		BillPeak MeasureCommonUnit             `bson:"billPeak,omitempty" json:"billPeak"` // 计费天数的区间带宽值(区间内的95值或峰值)，目前账单计费用此字段
		Peak     MeasureCommonUnit             `bson:"peak,omitempty" json:"peak"`         // 区间带宽值(区间内的95值或峰值)
	}
	MeasureCommonUnit struct {
		Bandwidth int64 `json:"bandwidth"` // 计量值，单位：bps
		Time      int64 `json:"time"`      // 当前计量值所处时刻，时间戳，秒
	}
	MeasureCommonUnitList struct {
		MeasureCommonUnitList []*MeasureCommonUnit `bson:"measureCommonUnitList" json:"measureCommonUnitList"`
	}
	LineBandWidthSeries struct {
		LineBandWidth map[string][]*MeasureCommonUnit `bson:"lineBandWidth" json:"lineBandWidth"`
	}
	LineCommonUint struct {
		Line      string `bson:"line" json:"line"`
		Bandwidth int64  `bson:"bandwidth" json:"bandwidth"`
	}
	LineMomentBandWidth struct {
		LineMoment map[string][]*LineCommonUint `bson:"lineCommonUint" json:"lineCommonUint"`
	}
	GetMomentNode struct {
		NodeId string `json:"nodeId"` // 节点 Id
		Time   int64  `json:"time"`   // 时刻值，时间戳，秒
	}
	NodePingLossRatioUnit struct {
		NodeId string  `bson:"nodeId" json:"nodeId"`
		Ratio  float64 `bson:"ratio" json:"ratio"`
	}
	PingLossNodeRatioData struct {
		PingLossRatio map[string][]*NodePingLossRatioUnit `bson:"pingLossRatio" json:"pingLossRatio"`
	}
	// NodeMomentData 节点时刻公用 data
	NodeMomentData struct {
		NodeId    string `bson:"nodeId" json:"nodeId"`
		Bandwidth int64  `bson:"bandwidth" json:"bandwidth"` // 计量值，单位：bps
		Time      int64  `bson:"time" json:"time"`           // 时刻值，时间戳，秒
	}

	NodeMomentDataList struct {
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
		NodeId string                 `json:"nodeId"`           // 节点 Id
		Ranges []*NodeMetricInfoRange `json:"ranges,omitempty"` // 计量排除时间范围
	}

	// NodeMetricInfoRange 时间范围
	NodeMetricInfoRange struct {
		Start int64 `json:"start"` // 范围起期，时间戳，秒
		End   int64 `json:"end"`   // 范围止期，时间戳，秒
	}
)

type (
	VendorNodeMetric struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonData
	}
	CustomerNodeMetric struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonData
	}
	VendorEveningMetric struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonData
	}
	CustomerEveningMetric struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonData
	}
	VendorNode5Min struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonUnitList
	}
	VendorNodeBw struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		NodeMomentDataList
	}
	CustomerNode5Min struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonUnitList
	}
	CustomerNodeBw struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		NodeMomentDataList
	}
	VendorPerNodeMetric struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonDataNodes
	}
	CustomerPerNodeMetric struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		MeasureCommonDataNodes
	}
	LineBwSeries struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		LineBandWidthSeries
	}
	LineMomentBw struct {
		Id     string `bson:"_id,omitempty"  json:"id"`
		Filter string `bson:"filter" json:"filter"`
		LineMomentBandWidth
	}
	PingLossNodeRatio struct {
		Id     string `bson:"_id" json:"id"`
		Filter string `bson:"filter" json:"filter"`
		PingLossNodeRatioData
	}
)

// ========================================================
type (
	// PostCustomerMeasureReq 业务侧计量接口参数
	PostCustomerMeasureReq struct {
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
	PostCustomerPerNodeMetricReq struct {
		Uid           string               `json:"uid"`           // 业务方 Id
		NodeType      string               `json:"nodeType"`      // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		Start         string               `json:"start"`         // 开始时间，格式：2006-01-02，东八区
		End           string               `json:"end"`           // 结束时间，格式：2006-01-02，东八区
		BillDays      int                  `json:"billDays"`      // 计费天数
		MeasureType   string               `json:"measureType"`   // 计量算法，peak：峰值 | peak95: 95峰值
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		DataSource    string               `json:"dataSource"`    // 计量数据源。process_ds：进程；netcard_ds：网卡。如果不传值，大数据使用网卡数据源
	}

	// PostCustomerNodeBwReq 业务侧获取时刻带宽值参数
	PostCustomerNodeBwReq struct {
		Uid           string           `json:"uid,optional"`  // 接入业务 快手、b站等业务id。注意这里uid就是指的业务id，因为uid很早就被定义为业务id
		NodeType      string           `json:"nodeType"`      // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes         []*GetMomentNode `json:"nodes"`         // 批量查询参数
		BandwidthType string           `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		DataSource    string           `json:"dataSource"`    // 计量数据源，process：进程；netcard：网卡
	}

	// PostCustomerNode5MinReq 业务侧计量 5min 数据明细参数
	PostCustomerNode5MinReq struct {
		SignatoryIds  []string             `json:"signatoryIds"`  // 签约方 Ids
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		NodeType      string               `json:"nodeType"`      // 节点类型，node | switch
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		Start         int64                `json:"start"`         // 开始时间，时间戳，秒
		End           int64                `json:"end"`           // 结束时间，时间戳，秒
		DataSource    string               `json:"dataSource"`    // 计量数据源。process：进程；netcard：网卡。如果不传值，大数据使用网卡数据源
	}

	// GetCustomerTimeQuantumSortReq 业务侧时间段内指定峰值
	PostCustomerEveningMetricReq struct {
		Quantum       TimeQuantum          `json:"eveningPeak"`   // 每天高峰时间区间
		Nodes         []*MeasureMetricNode `json:"nodes"`         // 节点列表
		NodeType      string               `json:"nodeType"`      // 节点类型，node | switch
		Uid           string               `json:"uid"`           // 业务方 Id
		BandwidthType string               `json:"bandwidthType"` // 带宽方向，上行减下行：out_minus_in / 上行：out；目前仅爱奇艺为上行减下行，其他为上行
		Start         string               `json:"start"`         // 开始时间，格式：2006-01-02
		End           string               `json:"end"`           // 结束时间，格式：2006-01-02
	}
	PostPingLossNodeRateReq struct {
		Quantum  TimeQuantum `json:"eveningPeak"` // 每天高峰时间区间
		NodeIds  []string    `json:"nodeIds"`
		VendorId uint32      `json:"vendorId"`
		NodeType string      `json:"nodeType"`
		Start    string      `json:"start"`
		End      string      `json:"end"`
	}
	// ========================================================

	// PostVendorPerNodeMetricReq 供应侧计量接口参数,用于分节点返回数据使用，不合并
	PostVendorPerNodeMetricReq struct {
		Vendor       string                  `json:"vendor"`                // 业务方 Id
		NodeType     string                  `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes        []*MeasureMetricPerNode `json:"nodes"`                 // 节点列表
		Start        string                  `json:"start"`                 // 开始时间，格式：2006-01-02，东八区
		End          string                  `json:"end"`                   // 结束时间，格式：2006-01-02，东八区
		BillDays     int64                   `json:"billDays"`              // 计费天数
		MeasureType  string                  `json:"measureType"`           // 计量算法，peak：峰值 | peak95: 95峰值
		AdjustParams []AdjustParam           `json:"adjustParams,optional"` // 带宽处理
	}

	// PostVendorNodeMetricReq 供应侧计量接口参数
	PostVendorNodeMetricReq struct {
		Vendor       string               `json:"vendor"`                // 业务方 Id
		NodeType     string               `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes        []*MeasureMetricNode `json:"nodes"`                 // 节点列表
		Start        string               `json:"start"`                 // 开始时间，格式：2006-01-02，东八区
		End          string               `json:"end"`                   // 结束时间，格式：2006-01-02，东八区
		BillDays     int64                `json:"billDays"`              // 计费天数
		MeasureType  string               `json:"measureType"`           // 计量算法，peak：峰值 | peak95: 95峰值
		AdjustParams []AdjustParam        `json:"adjustParams,optional"` // 带宽处理
		DataSource   string               `json:"dataSource,optional"`   // 计量数据源
	}
	PostGenerateMeasureReq struct {
		VendorId   string    `json:"vendorId"`
		NodeId     string    `json:"nodeId"` // 节点id
		CustomerId string    `json:"customerId"`
		Start      time.Time `json:"start"` // 开始时间
		End        time.Time `json:"end"`   // 结束时间
		Flow       int64     `json:"flow"`
	}

	// PostVendorNodeBwReq 供应侧获取时刻带宽值参数
	PostVendorNodeBwReq struct {
		VendorId     string           `json:"vendor"`                // 供应商 Id
		NodeType     string           `json:"nodeType"`              // 节点类型，node：单机 | switch：交换机，默认单机
		Nodes        []*GetMomentNode `json:"nodes"`                 // 批量查询参数
		AdjustParams []AdjustParam    `json:"adjustParams,optional"` // 带宽处理
	}

	// PostVendorNode5MinReq 供应侧计量 5min 数据明细参数
	PostVendorNode5MinReq struct {
		Nodes        []*MeasureMetricNode `json:"nodes"`                 // 节点列表
		NodeType     string               `json:"nodeType"`              // 节点类型，node | switch
		Vendor       string               `json:"vendor"`                // 供应商 Id
		AdjustParams []AdjustParam        `json:"adjustParams,optional"` // 带宽处理
		Start        int64                `json:"start"`                 // 开始时间，时间戳，秒
		End          int64                `json:"end"`                   // 结束时间，时间戳，秒
	}

	// GetVendorTimeQuantumSortReq 供应侧时间段内指定峰值
	PostVendorEveningMetricReq struct {
		Quantum      TimeQuantum          `json:"eveningPeak"`           // 每天高峰时间区间
		Nodes        []*MeasureMetricNode `json:"nodes"`                 // 节点列表
		NodeType     string               `json:"nodeType"`              // 节点类型，node | switch
		Uid          string               `json:"vendor"`                // 供应商 Id
		AdjustParams []AdjustParam        `json:"adjustParams,optional"` // 带宽处理
		Start        string               `json:"start"`                 // 开始时间，格式：2006-01-02
		End          string               `json:"end"`                   // 结束时间，格式：2006-01-02
	}
	PostLineBwSeriesReq struct {
		VendorId uint32 `json:"vendorId,optional"` // 供应商 Id
		NodeId   string `json:"nodeId"`            // 节点id
		Start    int64  `json:"start"`             // 开始时间，闭区间，时间戳，单位：秒
		End      int64  `json:"end"`               // 结束时间，开区间，时间戳，单位：秒
	}

	PostLineMomentBwReq struct {
		VendorId uint32   `json:"vendorId,optional"` // 供应商 Id
		NodeIds  []string `json:"nodeIds"`           // 多个节点id。大数据接口nodeIds size最大为100
		Time     int64    `json:"time"`              // 时间戳，单位：秒
	}
)
