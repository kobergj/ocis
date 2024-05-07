// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ocis/messages/search/v0/search.proto

package v0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	OpaqueId  string `protobuf:"bytes,2,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
	SpaceId   string `protobuf:"bytes,3,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
}

func (x *ResourceID) Reset() {
	*x = ResourceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_messages_search_v0_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceID) ProtoMessage() {}

func (x *ResourceID) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_messages_search_v0_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceID.ProtoReflect.Descriptor instead.
func (*ResourceID) Descriptor() ([]byte, []int) {
	return file_ocis_messages_search_v0_search_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceID) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *ResourceID) GetOpaqueId() string {
	if x != nil {
		return x.OpaqueId
	}
	return ""
}

func (x *ResourceID) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId *ResourceID `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Path       string      `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_messages_search_v0_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_messages_search_v0_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_ocis_messages_search_v0_search_proto_rawDescGZIP(), []int{1}
}

func (x *Reference) GetResourceId() *ResourceID {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *Reference) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Audio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album             *string `protobuf:"bytes,1,opt,name=album,proto3,oneof" json:"album,omitempty"`
	AlbumArtist       *string `protobuf:"bytes,2,opt,name=albumArtist,proto3,oneof" json:"albumArtist,omitempty"`
	Artist            *string `protobuf:"bytes,3,opt,name=artist,proto3,oneof" json:"artist,omitempty"`
	Bitrate           *int64  `protobuf:"varint,4,opt,name=bitrate,proto3,oneof" json:"bitrate,omitempty"`
	Composers         *string `protobuf:"bytes,5,opt,name=composers,proto3,oneof" json:"composers,omitempty"`
	Copyright         *string `protobuf:"bytes,6,opt,name=copyright,proto3,oneof" json:"copyright,omitempty"`
	Disc              *int32  `protobuf:"varint,7,opt,name=disc,proto3,oneof" json:"disc,omitempty"`
	DiscCount         *int32  `protobuf:"varint,8,opt,name=discCount,proto3,oneof" json:"discCount,omitempty"`
	Duration          *int64  `protobuf:"varint,9,opt,name=duration,proto3,oneof" json:"duration,omitempty"`
	Genre             *string `protobuf:"bytes,10,opt,name=genre,proto3,oneof" json:"genre,omitempty"`
	HasDrm            *bool   `protobuf:"varint,11,opt,name=hasDrm,proto3,oneof" json:"hasDrm,omitempty"`
	IsVariableBitrate *bool   `protobuf:"varint,12,opt,name=isVariableBitrate,proto3,oneof" json:"isVariableBitrate,omitempty"`
	Title             *string `protobuf:"bytes,13,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Track             *int32  `protobuf:"varint,14,opt,name=track,proto3,oneof" json:"track,omitempty"`
	TrackCount        *int32  `protobuf:"varint,15,opt,name=trackCount,proto3,oneof" json:"trackCount,omitempty"`
	Year              *int32  `protobuf:"varint,16,opt,name=year,proto3,oneof" json:"year,omitempty"`
}

func (x *Audio) Reset() {
	*x = Audio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_messages_search_v0_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audio) ProtoMessage() {}

func (x *Audio) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_messages_search_v0_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audio.ProtoReflect.Descriptor instead.
func (*Audio) Descriptor() ([]byte, []int) {
	return file_ocis_messages_search_v0_search_proto_rawDescGZIP(), []int{2}
}

func (x *Audio) GetAlbum() string {
	if x != nil && x.Album != nil {
		return *x.Album
	}
	return ""
}

func (x *Audio) GetAlbumArtist() string {
	if x != nil && x.AlbumArtist != nil {
		return *x.AlbumArtist
	}
	return ""
}

func (x *Audio) GetArtist() string {
	if x != nil && x.Artist != nil {
		return *x.Artist
	}
	return ""
}

func (x *Audio) GetBitrate() int64 {
	if x != nil && x.Bitrate != nil {
		return *x.Bitrate
	}
	return 0
}

func (x *Audio) GetComposers() string {
	if x != nil && x.Composers != nil {
		return *x.Composers
	}
	return ""
}

func (x *Audio) GetCopyright() string {
	if x != nil && x.Copyright != nil {
		return *x.Copyright
	}
	return ""
}

func (x *Audio) GetDisc() int32 {
	if x != nil && x.Disc != nil {
		return *x.Disc
	}
	return 0
}

func (x *Audio) GetDiscCount() int32 {
	if x != nil && x.DiscCount != nil {
		return *x.DiscCount
	}
	return 0
}

func (x *Audio) GetDuration() int64 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *Audio) GetGenre() string {
	if x != nil && x.Genre != nil {
		return *x.Genre
	}
	return ""
}

func (x *Audio) GetHasDrm() bool {
	if x != nil && x.HasDrm != nil {
		return *x.HasDrm
	}
	return false
}

func (x *Audio) GetIsVariableBitrate() bool {
	if x != nil && x.IsVariableBitrate != nil {
		return *x.IsVariableBitrate
	}
	return false
}

func (x *Audio) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Audio) GetTrack() int32 {
	if x != nil && x.Track != nil {
		return *x.Track
	}
	return 0
}

func (x *Audio) GetTrackCount() int32 {
	if x != nil && x.TrackCount != nil {
		return *x.TrackCount
	}
	return 0
}

func (x *Audio) GetYear() int32 {
	if x != nil && x.Year != nil {
		return *x.Year
	}
	return 0
}

type GeoCoordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Altitude  *float64 `protobuf:"fixed64,1,opt,name=altitude,proto3,oneof" json:"altitude,omitempty"`
	Latitude  *float64 `protobuf:"fixed64,2,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude *float64 `protobuf:"fixed64,3,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
}

func (x *GeoCoordinates) Reset() {
	*x = GeoCoordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_messages_search_v0_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCoordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCoordinates) ProtoMessage() {}

func (x *GeoCoordinates) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_messages_search_v0_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoCoordinates.ProtoReflect.Descriptor instead.
func (*GeoCoordinates) Descriptor() ([]byte, []int) {
	return file_ocis_messages_search_v0_search_proto_rawDescGZIP(), []int{3}
}

func (x *GeoCoordinates) GetAltitude() float64 {
	if x != nil && x.Altitude != nil {
		return *x.Altitude
	}
	return 0
}

func (x *GeoCoordinates) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *GeoCoordinates) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref              *Reference             `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Id               *ResourceID            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Etag             string                 `protobuf:"bytes,4,opt,name=etag,proto3" json:"etag,omitempty"`
	Size             uint64                 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	LastModifiedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_modified_time,json=lastModifiedTime,proto3" json:"last_modified_time,omitempty"`
	MimeType         string                 `protobuf:"bytes,7,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Permissions      string                 `protobuf:"bytes,8,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Type             uint64                 `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Deleted          bool                   `protobuf:"varint,10,opt,name=deleted,proto3" json:"deleted,omitempty"`
	ShareRootName    string                 `protobuf:"bytes,11,opt,name=shareRootName,proto3" json:"shareRootName,omitempty"`
	ParentId         *ResourceID            `protobuf:"bytes,12,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Tags             []string               `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	Highlights       string                 `protobuf:"bytes,14,opt,name=highlights,proto3" json:"highlights,omitempty"`
	Audio            *Audio                 `protobuf:"bytes,15,opt,name=audio,proto3" json:"audio,omitempty"`
	Location         *GeoCoordinates        `protobuf:"bytes,16,opt,name=location,proto3" json:"location,omitempty"`
	RemoteItemId     *ResourceID            `protobuf:"bytes,17,opt,name=remote_item_id,json=remoteItemId,proto3" json:"remote_item_id,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_messages_search_v0_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_messages_search_v0_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_ocis_messages_search_v0_search_proto_rawDescGZIP(), []int{4}
}

func (x *Entity) GetRef() *Reference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *Entity) GetId() *ResourceID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Entity) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Entity) GetLastModifiedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModifiedTime
	}
	return nil
}

func (x *Entity) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Entity) GetPermissions() string {
	if x != nil {
		return x.Permissions
	}
	return ""
}

func (x *Entity) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Entity) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Entity) GetShareRootName() string {
	if x != nil {
		return x.ShareRootName
	}
	return ""
}

func (x *Entity) GetParentId() *ResourceID {
	if x != nil {
		return x.ParentId
	}
	return nil
}

func (x *Entity) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Entity) GetHighlights() string {
	if x != nil {
		return x.Highlights
	}
	return ""
}

func (x *Entity) GetAudio() *Audio {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *Entity) GetLocation() *GeoCoordinates {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Entity) GetRemoteItemId() *ResourceID {
	if x != nil {
		return x.RemoteItemId
	}
	return nil
}

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the matched entity
	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// the match score
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_messages_search_v0_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_messages_search_v0_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_ocis_messages_search_v0_search_proto_rawDescGZIP(), []int{5}
}

func (x *Match) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Match) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_ocis_messages_search_v0_search_proto protoreflect.FileDescriptor

var file_ocis_messages_search_v0_search_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x63, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76,
	0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xcf, 0x05, 0x0a,
	0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x41,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x70, 0x79, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x09, 0x63, 0x6f,
	0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x69,
	0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x04, 0x64, 0x69, 0x73, 0x63,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x07, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x44, 0x72, 0x6d, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x0a, 0x52, 0x06, 0x68, 0x61, 0x73, 0x44, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12,
	0x31, 0x0a, 0x11, 0x69, 0x73, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x69, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0b, 0x52, 0x11, 0x69, 0x73,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x0c, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0d, 0x52, 0x05,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0e, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0f, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x61, 0x73, 0x44, 0x72, 0x6d, 0x42, 0x14, 0x0a, 0x12,
	0x5f, 0x69, 0x73, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x69, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x22, 0x9d,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x6c, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xc8,
	0x05, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x03, 0x72, 0x65, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12,
	0x33, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x63,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x48, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x44, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x43, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x63,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49,
	0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x52, 0x0c, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x05, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x37, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x63,
	0x69, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocis_messages_search_v0_search_proto_rawDescOnce sync.Once
	file_ocis_messages_search_v0_search_proto_rawDescData = file_ocis_messages_search_v0_search_proto_rawDesc
)

func file_ocis_messages_search_v0_search_proto_rawDescGZIP() []byte {
	file_ocis_messages_search_v0_search_proto_rawDescOnce.Do(func() {
		file_ocis_messages_search_v0_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocis_messages_search_v0_search_proto_rawDescData)
	})
	return file_ocis_messages_search_v0_search_proto_rawDescData
}

var file_ocis_messages_search_v0_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ocis_messages_search_v0_search_proto_goTypes = []interface{}{
	(*ResourceID)(nil),            // 0: ocis.messages.search.v0.ResourceID
	(*Reference)(nil),             // 1: ocis.messages.search.v0.Reference
	(*Audio)(nil),                 // 2: ocis.messages.search.v0.Audio
	(*GeoCoordinates)(nil),        // 3: ocis.messages.search.v0.GeoCoordinates
	(*Entity)(nil),                // 4: ocis.messages.search.v0.Entity
	(*Match)(nil),                 // 5: ocis.messages.search.v0.Match
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_ocis_messages_search_v0_search_proto_depIdxs = []int32{
	0, // 0: ocis.messages.search.v0.Reference.resource_id:type_name -> ocis.messages.search.v0.ResourceID
	1, // 1: ocis.messages.search.v0.Entity.ref:type_name -> ocis.messages.search.v0.Reference
	0, // 2: ocis.messages.search.v0.Entity.id:type_name -> ocis.messages.search.v0.ResourceID
	6, // 3: ocis.messages.search.v0.Entity.last_modified_time:type_name -> google.protobuf.Timestamp
	0, // 4: ocis.messages.search.v0.Entity.parent_id:type_name -> ocis.messages.search.v0.ResourceID
	2, // 5: ocis.messages.search.v0.Entity.audio:type_name -> ocis.messages.search.v0.Audio
	3, // 6: ocis.messages.search.v0.Entity.location:type_name -> ocis.messages.search.v0.GeoCoordinates
	0, // 7: ocis.messages.search.v0.Entity.remote_item_id:type_name -> ocis.messages.search.v0.ResourceID
	4, // 8: ocis.messages.search.v0.Match.entity:type_name -> ocis.messages.search.v0.Entity
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_ocis_messages_search_v0_search_proto_init() }
func file_ocis_messages_search_v0_search_proto_init() {
	if File_ocis_messages_search_v0_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocis_messages_search_v0_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceID); i {
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
		file_ocis_messages_search_v0_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_ocis_messages_search_v0_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audio); i {
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
		file_ocis_messages_search_v0_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCoordinates); i {
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
		file_ocis_messages_search_v0_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_ocis_messages_search_v0_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
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
	file_ocis_messages_search_v0_search_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_ocis_messages_search_v0_search_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocis_messages_search_v0_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocis_messages_search_v0_search_proto_goTypes,
		DependencyIndexes: file_ocis_messages_search_v0_search_proto_depIdxs,
		MessageInfos:      file_ocis_messages_search_v0_search_proto_msgTypes,
	}.Build()
	File_ocis_messages_search_v0_search_proto = out.File
	file_ocis_messages_search_v0_search_proto_rawDesc = nil
	file_ocis_messages_search_v0_search_proto_goTypes = nil
	file_ocis_messages_search_v0_search_proto_depIdxs = nil
}
