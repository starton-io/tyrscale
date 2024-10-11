package dto

type ListUpstreamReq struct {
	Uuid      *string `query:"uuid" validate:"omitempty,uuid"`
	RouteUuid string  `query:"route_uuid" validate:"required,uuid"`
}

type UpstreamCreateReq struct {
	RouteUuid string  `json:"route_uuid"`
	Host      string  `json:"host"`
	Port      int     `json:"port"`
	Path      string  `json:"path"`
	Scheme    string  `json:"scheme"`
	Weight    float64 `json:"weight"`
}

type Upstream struct {
	Uuid             string                    `json:"uuid" validate:"omitempty,uuid"`
	RpcUuid          string                    `json:"rpc_uuid" validate:"excluded_with=Host Port Path Scheme"`
	Host             string                    `json:"host" validate:"required_without=RpcUuid"`
	Port             int32                     `json:"port" validate:"required_without=RpcUuid,gte=0,lte=65535"`
	Path             string                    `json:"path" validate:"required_without=RpcUuid,omitempty"`
	Scheme           string                    `json:"scheme" validate:"required_without=RpcUuid"`
	Weight           float64                   `json:"weight" validate:"required,gte=0,lte=100"`
	FastHTTPSettings *UpstreamFastHTTPSettings `json:"fasthttp_settings"`
}

type UpstreamFastHTTPSettings struct {
	ProxyHost string `json:"proxy_host" validate:"omitempty,regexp=^(?:([a-zA-Z0-9._-]+):([a-zA-Z0-9._-]+)@)?([a-zA-Z0-9]{1}[a-zA-Z0-9_-]+)(\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]+)*?:([0-9]+)$"`
}

type UpstreamUpdateReq struct {
	Uuid      string  `json:"uuid" validate:"required,uuid"`
	RouteUuid string  `json:"route_uuid" validate:"required,uuid"`
	Weight    float64 `json:"weight" validate:"required,gte=0,lte=100"`
}

type ListUpstreamRes struct {
	Upstreams []*Upstream `json:"items"`
}

type UpstreamDeleteReq struct {
	Uuid      string `json:"uuid"`
	RouteUuid string `json:"route_uuid"`
}

type UpstreamUpsertRes struct {
	Uuid string `json:"uuid"`
}
