// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: gateway/proto/plugin/plugin.proto

package plugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for ListPlugins RPC
type ListPluginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPluginsRequest) Reset() {
	*x = ListPluginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginsRequest) ProtoMessage() {}

func (x *ListPluginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginsRequest.ProtoReflect.Descriptor instead.
func (*ListPluginsRequest) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

// Response message for ListPlugins RPC
type ListPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map of plugin lists, keyed by an integer ID
	Plugins map[string]*PluginList `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListPluginsResponse) Reset() {
	*x = ListPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginsResponse) ProtoMessage() {}

func (x *ListPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginsResponse.ProtoReflect.Descriptor instead.
func (*ListPluginsResponse) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *ListPluginsResponse) GetPlugins() map[string]*PluginList {
	if x != nil {
		return x.Plugins
	}
	return nil
}

// Message representing a list of plugins
type PluginList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repeated field of plugin names
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *PluginList) Reset() {
	*x = PluginList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginList) ProtoMessage() {}

func (x *PluginList) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginList.ProtoReflect.Descriptor instead.
func (*PluginList) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *PluginList) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type Plugins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestInterceptor  []*Plugin `protobuf:"bytes,1,rep,name=request_interceptor,json=RequestInterceptor,proto3" json:"request_interceptor,omitempty"`
	ResponseInterceptor []*Plugin `protobuf:"bytes,2,rep,name=response_interceptor,json=ResponseInterceptor,proto3" json:"response_interceptor,omitempty"`
	Middleware          []*Plugin `protobuf:"bytes,3,rep,name=middleware,json=Middleware,proto3" json:"middleware,omitempty"`
}

func (x *Plugins) Reset() {
	*x = Plugins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugins) ProtoMessage() {}

func (x *Plugins) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugins.ProtoReflect.Descriptor instead.
func (*Plugins) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *Plugins) GetRequestInterceptor() []*Plugin {
	if x != nil {
		return x.RequestInterceptor
	}
	return nil
}

func (x *Plugins) GetResponseInterceptor() []*Plugin {
	if x != nil {
		return x.ResponseInterceptor
	}
	return nil
}

func (x *Plugins) GetMiddleware() []*Plugin {
	if x != nil {
		return x.Middleware
	}
	return nil
}

type PublishPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteHost         string `protobuf:"bytes,1,opt,name=route_host,json=routeHost,proto3" json:"route_host,omitempty"`
	RoutePath         string `protobuf:"bytes,2,opt,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	PluginName        string `protobuf:"bytes,3,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	PluginDescription string `protobuf:"bytes,4,opt,name=plugin_description,json=pluginDescription,proto3" json:"plugin_description,omitempty"`
	PluginConfig      []byte `protobuf:"bytes,5,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
	PluginType        string `protobuf:"bytes,6,opt,name=plugin_type,json=pluginType,proto3" json:"plugin_type,omitempty"`
	PluginPriority    int32  `protobuf:"varint,7,opt,name=plugin_priority,json=pluginPriority,proto3" json:"plugin_priority,omitempty"`
}

func (x *PublishPlugin) Reset() {
	*x = PublishPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishPlugin) ProtoMessage() {}

func (x *PublishPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishPlugin.ProtoReflect.Descriptor instead.
func (*PublishPlugin) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *PublishPlugin) GetRouteHost() string {
	if x != nil {
		return x.RouteHost
	}
	return ""
}

func (x *PublishPlugin) GetRoutePath() string {
	if x != nil {
		return x.RoutePath
	}
	return ""
}

func (x *PublishPlugin) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *PublishPlugin) GetPluginDescription() string {
	if x != nil {
		return x.PluginDescription
	}
	return ""
}

func (x *PublishPlugin) GetPluginConfig() []byte {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

func (x *PublishPlugin) GetPluginType() string {
	if x != nil {
		return x.PluginType
	}
	return ""
}

func (x *PublishPlugin) GetPluginPriority() int32 {
	if x != nil {
		return x.PluginPriority
	}
	return 0
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Config      []byte `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	Priority    uint32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Plugin) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Plugin) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Plugin) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type ValidatePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type is the type of the plugin
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Payload is the configuration for the plugin
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ValidatePluginRequest) Reset() {
	*x = ValidatePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePluginRequest) ProtoMessage() {}

func (x *ValidatePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePluginRequest.ProtoReflect.Descriptor instead.
func (*ValidatePluginRequest) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *ValidatePluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ValidatePluginRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ValidatePluginRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ValidatePluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidatePluginResponse) Reset() {
	*x = ValidatePluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePluginResponse) ProtoMessage() {}

func (x *ValidatePluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_plugin_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePluginResponse.ProtoReflect.Descriptor instead.
func (*ValidatePluginResponse) Descriptor() ([]byte, []int) {
	return file_gateway_proto_plugin_plugin_proto_rawDescGZIP(), []int{7}
}

var File_gateway_proto_plugin_plugin_proto protoreflect.FileDescriptor

var file_gateway_proto_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa9,
	0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x4e, 0x0a, 0x0c, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x0a, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xbd,
	0x01, 0x0a, 0x07, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x13, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x14, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x2e,
	0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x22, 0x8c,
	0x02, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x86, 0x01,
	0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x59, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x18, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa0, 0x01, 0x0a, 0x0d,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_gateway_proto_plugin_plugin_proto_rawDescOnce sync.Once
	file_gateway_proto_plugin_plugin_proto_rawDescData = file_gateway_proto_plugin_plugin_proto_rawDesc
)

func file_gateway_proto_plugin_plugin_proto_rawDescGZIP() []byte {
	file_gateway_proto_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_gateway_proto_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_plugin_plugin_proto_rawDescData)
	})
	return file_gateway_proto_plugin_plugin_proto_rawDescData
}

var file_gateway_proto_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gateway_proto_plugin_plugin_proto_goTypes = []interface{}{
	(*ListPluginsRequest)(nil),     // 0: plugin.ListPluginsRequest
	(*ListPluginsResponse)(nil),    // 1: plugin.ListPluginsResponse
	(*PluginList)(nil),             // 2: plugin.PluginList
	(*Plugins)(nil),                // 3: plugin.Plugins
	(*PublishPlugin)(nil),          // 4: plugin.PublishPlugin
	(*Plugin)(nil),                 // 5: plugin.Plugin
	(*ValidatePluginRequest)(nil),  // 6: plugin.ValidatePluginRequest
	(*ValidatePluginResponse)(nil), // 7: plugin.ValidatePluginResponse
	nil,                            // 8: plugin.ListPluginsResponse.PluginsEntry
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_gateway_proto_plugin_plugin_proto_depIdxs = []int32{
	8, // 0: plugin.ListPluginsResponse.plugins:type_name -> plugin.ListPluginsResponse.PluginsEntry
	5, // 1: plugin.Plugins.request_interceptor:type_name -> plugin.Plugin
	5, // 2: plugin.Plugins.response_interceptor:type_name -> plugin.Plugin
	5, // 3: plugin.Plugins.middleware:type_name -> plugin.Plugin
	2, // 4: plugin.ListPluginsResponse.PluginsEntry.value:type_name -> plugin.PluginList
	0, // 5: plugin.PluginService.ListPlugins:input_type -> plugin.ListPluginsRequest
	6, // 6: plugin.PluginService.ValidatePlugin:input_type -> plugin.ValidatePluginRequest
	1, // 7: plugin.PluginService.ListPlugins:output_type -> plugin.ListPluginsResponse
	9, // 8: plugin.PluginService.ValidatePlugin:output_type -> google.protobuf.Empty
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gateway_proto_plugin_plugin_proto_init() }
func file_gateway_proto_plugin_plugin_proto_init() {
	if File_gateway_proto_plugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_plugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugins); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishPlugin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePluginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_plugin_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePluginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_proto_plugin_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_proto_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_gateway_proto_plugin_plugin_proto_depIdxs,
		MessageInfos:      file_gateway_proto_plugin_plugin_proto_msgTypes,
	}.Build()
	File_gateway_proto_plugin_plugin_proto = out.File
	file_gateway_proto_plugin_plugin_proto_rawDesc = nil
	file_gateway_proto_plugin_plugin_proto_goTypes = nil
	file_gateway_proto_plugin_plugin_proto_depIdxs = nil
}
