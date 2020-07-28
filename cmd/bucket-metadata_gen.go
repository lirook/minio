package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BucketMetadata) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Created":
			z.Created, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "Created")
				return
			}
		case "LockEnabled":
			z.LockEnabled, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "LockEnabled")
				return
			}
		case "PolicyConfigJSON":
			z.PolicyConfigJSON, err = dc.ReadBytes(z.PolicyConfigJSON)
			if err != nil {
				err = msgp.WrapError(err, "PolicyConfigJSON")
				return
			}
		case "NotificationConfigXML":
			z.NotificationConfigXML, err = dc.ReadBytes(z.NotificationConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "NotificationConfigXML")
				return
			}
		case "LifecycleConfigXML":
			z.LifecycleConfigXML, err = dc.ReadBytes(z.LifecycleConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "LifecycleConfigXML")
				return
			}
		case "ObjectLockConfigXML":
			z.ObjectLockConfigXML, err = dc.ReadBytes(z.ObjectLockConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "ObjectLockConfigXML")
				return
			}
		case "VersioningConfigXML":
			z.VersioningConfigXML, err = dc.ReadBytes(z.VersioningConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "VersioningConfigXML")
				return
			}
		case "EncryptionConfigXML":
			z.EncryptionConfigXML, err = dc.ReadBytes(z.EncryptionConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "EncryptionConfigXML")
				return
			}
		case "TaggingConfigXML":
			z.TaggingConfigXML, err = dc.ReadBytes(z.TaggingConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "TaggingConfigXML")
				return
			}
		case "QuotaConfigJSON":
			z.QuotaConfigJSON, err = dc.ReadBytes(z.QuotaConfigJSON)
			if err != nil {
				err = msgp.WrapError(err, "QuotaConfigJSON")
				return
			}
		case "ReplicationConfigXML":
			z.ReplicationConfigXML, err = dc.ReadBytes(z.ReplicationConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "ReplicationConfigXML")
				return
			}
		case "BucketTargetsConfigJSON":
			z.BucketTargetsConfigJSON, err = dc.ReadBytes(z.BucketTargetsConfigJSON)
			if err != nil {
				err = msgp.WrapError(err, "BucketTargetsConfigJSON")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketMetadata) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 13
	// write "Name"
	err = en.Append(0x8d, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Created"
	err = en.Append(0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Created)
	if err != nil {
		err = msgp.WrapError(err, "Created")
		return
	}
	// write "LockEnabled"
	err = en.Append(0xab, 0x4c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBool(z.LockEnabled)
	if err != nil {
		err = msgp.WrapError(err, "LockEnabled")
		return
	}
	// write "PolicyConfigJSON"
	err = en.Append(0xb0, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PolicyConfigJSON)
	if err != nil {
		err = msgp.WrapError(err, "PolicyConfigJSON")
		return
	}
	// write "NotificationConfigXML"
	err = en.Append(0xb5, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.NotificationConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "NotificationConfigXML")
		return
	}
	// write "LifecycleConfigXML"
	err = en.Append(0xb2, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.LifecycleConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "LifecycleConfigXML")
		return
	}
	// write "ObjectLockConfigXML"
	err = en.Append(0xb3, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.ObjectLockConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "ObjectLockConfigXML")
		return
	}
	// write "VersioningConfigXML"
	err = en.Append(0xb3, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.VersioningConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "VersioningConfigXML")
		return
	}
	// write "EncryptionConfigXML"
	err = en.Append(0xb3, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.EncryptionConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "EncryptionConfigXML")
		return
	}
	// write "TaggingConfigXML"
	err = en.Append(0xb0, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.TaggingConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "TaggingConfigXML")
		return
	}
	// write "QuotaConfigJSON"
	err = en.Append(0xaf, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.QuotaConfigJSON)
	if err != nil {
		err = msgp.WrapError(err, "QuotaConfigJSON")
		return
	}
	// write "ReplicationConfigXML"
	err = en.Append(0xb4, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.ReplicationConfigXML)
	if err != nil {
		err = msgp.WrapError(err, "ReplicationConfigXML")
		return
	}
	// write "BucketTargetsConfigJSON"
	err = en.Append(0xb7, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.BucketTargetsConfigJSON)
	if err != nil {
		err = msgp.WrapError(err, "BucketTargetsConfigJSON")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketMetadata) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 13
	// string "Name"
	o = append(o, 0x8d, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Created"
	o = append(o, 0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Created)
	// string "LockEnabled"
	o = append(o, 0xab, 0x4c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	o = msgp.AppendBool(o, z.LockEnabled)
	// string "PolicyConfigJSON"
	o = append(o, 0xb0, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e)
	o = msgp.AppendBytes(o, z.PolicyConfigJSON)
	// string "NotificationConfigXML"
	o = append(o, 0xb5, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.NotificationConfigXML)
	// string "LifecycleConfigXML"
	o = append(o, 0xb2, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.LifecycleConfigXML)
	// string "ObjectLockConfigXML"
	o = append(o, 0xb3, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.ObjectLockConfigXML)
	// string "VersioningConfigXML"
	o = append(o, 0xb3, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.VersioningConfigXML)
	// string "EncryptionConfigXML"
	o = append(o, 0xb3, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.EncryptionConfigXML)
	// string "TaggingConfigXML"
	o = append(o, 0xb0, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.TaggingConfigXML)
	// string "QuotaConfigJSON"
	o = append(o, 0xaf, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e)
	o = msgp.AppendBytes(o, z.QuotaConfigJSON)
	// string "ReplicationConfigXML"
	o = append(o, 0xb4, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x58, 0x4d, 0x4c)
	o = msgp.AppendBytes(o, z.ReplicationConfigXML)
	// string "BucketTargetsConfigJSON"
	o = append(o, 0xb7, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e)
	o = msgp.AppendBytes(o, z.BucketTargetsConfigJSON)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketMetadata) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Created":
			z.Created, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Created")
				return
			}
		case "LockEnabled":
			z.LockEnabled, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LockEnabled")
				return
			}
		case "PolicyConfigJSON":
			z.PolicyConfigJSON, bts, err = msgp.ReadBytesBytes(bts, z.PolicyConfigJSON)
			if err != nil {
				err = msgp.WrapError(err, "PolicyConfigJSON")
				return
			}
		case "NotificationConfigXML":
			z.NotificationConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.NotificationConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "NotificationConfigXML")
				return
			}
		case "LifecycleConfigXML":
			z.LifecycleConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.LifecycleConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "LifecycleConfigXML")
				return
			}
		case "ObjectLockConfigXML":
			z.ObjectLockConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.ObjectLockConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "ObjectLockConfigXML")
				return
			}
		case "VersioningConfigXML":
			z.VersioningConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.VersioningConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "VersioningConfigXML")
				return
			}
		case "EncryptionConfigXML":
			z.EncryptionConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.EncryptionConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "EncryptionConfigXML")
				return
			}
		case "TaggingConfigXML":
			z.TaggingConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.TaggingConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "TaggingConfigXML")
				return
			}
		case "QuotaConfigJSON":
			z.QuotaConfigJSON, bts, err = msgp.ReadBytesBytes(bts, z.QuotaConfigJSON)
			if err != nil {
				err = msgp.WrapError(err, "QuotaConfigJSON")
				return
			}
		case "ReplicationConfigXML":
			z.ReplicationConfigXML, bts, err = msgp.ReadBytesBytes(bts, z.ReplicationConfigXML)
			if err != nil {
				err = msgp.WrapError(err, "ReplicationConfigXML")
				return
			}
		case "BucketTargetsConfigJSON":
			z.BucketTargetsConfigJSON, bts, err = msgp.ReadBytesBytes(bts, z.BucketTargetsConfigJSON)
			if err != nil {
				err = msgp.WrapError(err, "BucketTargetsConfigJSON")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketMetadata) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.TimeSize + 12 + msgp.BoolSize + 17 + msgp.BytesPrefixSize + len(z.PolicyConfigJSON) + 22 + msgp.BytesPrefixSize + len(z.NotificationConfigXML) + 19 + msgp.BytesPrefixSize + len(z.LifecycleConfigXML) + 20 + msgp.BytesPrefixSize + len(z.ObjectLockConfigXML) + 20 + msgp.BytesPrefixSize + len(z.VersioningConfigXML) + 20 + msgp.BytesPrefixSize + len(z.EncryptionConfigXML) + 17 + msgp.BytesPrefixSize + len(z.TaggingConfigXML) + 16 + msgp.BytesPrefixSize + len(z.QuotaConfigJSON) + 21 + msgp.BytesPrefixSize + len(z.ReplicationConfigXML) + 24 + msgp.BytesPrefixSize + len(z.BucketTargetsConfigJSON)
	return
}
