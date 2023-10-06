// Code generated with zap-obj-marshaler-gen, DO NOT EDIT.
package examples

import "go.uber.org/zap/zapcore"

func (l *Example) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	var err error

	if v, ok := any(l.Other).(zapcore.ObjectMarshaler); ok {
		err = enc.AddObject("Other", v)
		if err != nil {
			return err
		}
	} else {
		err = enc.AddReflected("Other", l.Other)
		if err != nil {
			return err
		}
	}

	if v, ok := any(l.OtherPtr).(zapcore.ObjectMarshaler); ok {
		err = enc.AddObject("OtherPtr", v)
		if err != nil {
			return err
		}
	} else {
		err = enc.AddReflected("OtherPtr", l.OtherPtr)
		if err != nil {
			return err
		}
	}
	enc.AddBool("Bool", l.Bool)
	if l.BoolPtr != nil {
		enc.AddBool("BoolPtr", *l.BoolPtr)
	} else {
		err = enc.AddReflected("BoolPtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddComplex128("Complex128", l.Complex128)
	if l.Complex128Ptr != nil {
		enc.AddComplex128("Complex128Ptr", *l.Complex128Ptr)
	} else {
		err = enc.AddReflected("Complex128Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddComplex64("Complex64", l.Complex64)
	if l.Complex64Ptr != nil {
		enc.AddComplex64("Complex64Ptr", *l.Complex64Ptr)
	} else {
		err = enc.AddReflected("Complex64Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddFloat64("Float64", l.Float64)
	if l.Float64Ptr != nil {
		enc.AddFloat64("Float64Ptr", *l.Float64Ptr)
	} else {
		err = enc.AddReflected("Float64Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddFloat32("Float32", l.Float32)
	if l.Float32Ptr != nil {
		enc.AddFloat32("Float32Ptr", *l.Float32Ptr)
	} else {
		err = enc.AddReflected("Float32Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddInt("Int", l.Int)
	if l.IntPtr != nil {
		enc.AddInt("IntPtr", *l.IntPtr)
	} else {
		err = enc.AddReflected("IntPtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddInt64("Int64", l.Int64)
	if l.Int64Ptr != nil {
		enc.AddInt64("Int64Ptr", *l.Int64Ptr)
	} else {
		err = enc.AddReflected("Int64Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddInt32("Int32", l.Int32)
	if l.Int32Ptr != nil {
		enc.AddInt32("Int32Ptr", *l.Int32Ptr)
	} else {
		err = enc.AddReflected("Int32Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddInt16("Int16", l.Int16)
	if l.Int16Ptr != nil {
		enc.AddInt16("Int16Ptr", *l.Int16Ptr)
	} else {
		err = enc.AddReflected("Int16Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddInt8("Int8", l.Int8)
	if l.Int8Ptr != nil {
		enc.AddInt8("Int8Ptr", *l.Int8Ptr)
	} else {
		err = enc.AddReflected("Int8Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddString("String", l.String)
	if l.StringPtr != nil {
		enc.AddString("StringPtr", *l.StringPtr)
	} else {
		err = enc.AddReflected("StringPtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddUint("Uint", l.Uint)
	if l.UintPtr != nil {
		enc.AddUint("UintPtr", *l.UintPtr)
	} else {
		err = enc.AddReflected("UintPtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddUint64("Uint64", l.Uint64)
	if l.Uint64Ptr != nil {
		enc.AddUint64("Uint64Ptr", *l.Uint64Ptr)
	} else {
		err = enc.AddReflected("Uint64Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddUint32("Uint32", l.Uint32)
	if l.Uint32Ptr != nil {
		enc.AddUint32("Uint32Ptr", *l.Uint32Ptr)
	} else {
		err = enc.AddReflected("Uint32Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddUint16("Uint16", l.Uint16)
	if l.Uint16Ptr != nil {
		enc.AddUint16("Uint16Ptr", *l.Uint16Ptr)
	} else {
		err = enc.AddReflected("Uint16Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddUint8("Uint8", l.Uint8)
	if l.Uint8Ptr != nil {
		enc.AddUint8("Uint8Ptr", *l.Uint8Ptr)
	} else {
		err = enc.AddReflected("Uint8Ptr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddBinary("Bytes", l.Bytes)
	enc.AddUintptr("UintPtrVal", l.UintPtrVal)
	if l.UintPtrValPtr != nil {
		enc.AddUintptr("UintPtrValPtr", *l.UintPtrValPtr)
	} else {
		err = enc.AddReflected("UintPtrValPtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddTime("Time", l.Time)
	if l.TimePtr != nil {
		enc.AddTime("TimePtr", *l.TimePtr)
	} else {
		err = enc.AddReflected("TimePtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddDuration("Duration", l.Duration)
	if l.DurationPtr != nil {
		enc.AddDuration("DurationPtr", *l.DurationPtr)
	} else {
		err = enc.AddReflected("DurationPtr", nil)
		if err != nil {
			return err
		}
	}
	enc.AddString("custom_str", l.CustomStr)
	enc.AddBool("custom_bool", l.CustomBool)
	return err
}
func (l *OmitEmpty) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	var err error

	if v, ok := any(l.Other).(zapcore.ObjectMarshaler); ok {
		err = enc.AddObject("Other", v)
		if err != nil {
			return err
		}
	} else {
		err = enc.AddReflected("Other", l.Other)
		if err != nil {
			return err
		}
	}
	if l.Bool != false {
		enc.AddBool("Bool", l.Bool)
	}
	if l.BoolPtr != nil {
		enc.AddBool("BoolPtr", *l.BoolPtr)
	}
	if l.Complex128 != 0 {
		enc.AddComplex128("Complex128", l.Complex128)
	}
	if l.Complex128Ptr != nil {
		enc.AddComplex128("Complex128Ptr", *l.Complex128Ptr)
	}
	if l.Complex64 != 0 {
		enc.AddComplex64("Complex64", l.Complex64)
	}
	if l.Complex64Ptr != nil {
		enc.AddComplex64("Complex64Ptr", *l.Complex64Ptr)
	}
	if l.Float64 != 0 {
		enc.AddFloat64("Float64", l.Float64)
	}
	if l.Float64Ptr != nil {
		enc.AddFloat64("Float64Ptr", *l.Float64Ptr)
	}
	if l.Float32 != 0 {
		enc.AddFloat32("Float32", l.Float32)
	}
	if l.Float32Ptr != nil {
		enc.AddFloat32("Float32Ptr", *l.Float32Ptr)
	}
	if l.Int != 0 {
		enc.AddInt("Int", l.Int)
	}
	if l.IntPtr != nil {
		enc.AddInt("IntPtr", *l.IntPtr)
	}
	if l.Int64 != 0 {
		enc.AddInt64("Int64", l.Int64)
	}
	if l.Int64Ptr != nil {
		enc.AddInt64("Int64Ptr", *l.Int64Ptr)
	}
	if l.Int32 != 0 {
		enc.AddInt32("Int32", l.Int32)
	}
	if l.Int32Ptr != nil {
		enc.AddInt32("Int32Ptr", *l.Int32Ptr)
	}
	if l.Int16 != 0 {
		enc.AddInt16("Int16", l.Int16)
	}
	if l.Int16Ptr != nil {
		enc.AddInt16("Int16Ptr", *l.Int16Ptr)
	}
	if l.Int8 != 0 {
		enc.AddInt8("Int8", l.Int8)
	}
	if l.Int8Ptr != nil {
		enc.AddInt8("Int8Ptr", *l.Int8Ptr)
	}
	if l.String != "" {
		enc.AddString("String", l.String)
	}
	if l.StringPtr != nil {
		enc.AddString("StringPtr", *l.StringPtr)
	}
	if l.Uint != 0 {
		enc.AddUint("Uint", l.Uint)
	}
	if l.UintPtr != nil {
		enc.AddUint("UintPtr", *l.UintPtr)
	}
	if l.Uint64 != 0 {
		enc.AddUint64("Uint64", l.Uint64)
	}
	if l.Uint64Ptr != nil {
		enc.AddUint64("Uint64Ptr", *l.Uint64Ptr)
	}
	if l.Uint32 != 0 {
		enc.AddUint32("Uint32", l.Uint32)
	}
	if l.Uint32Ptr != nil {
		enc.AddUint32("Uint32Ptr", *l.Uint32Ptr)
	}
	if l.Uint16 != 0 {
		enc.AddUint16("Uint16", l.Uint16)
	}
	if l.Uint16Ptr != nil {
		enc.AddUint16("Uint16Ptr", *l.Uint16Ptr)
	}
	if l.Uint8 != 0 {
		enc.AddUint8("Uint8", l.Uint8)
	}
	if l.Uint8Ptr != nil {
		enc.AddUint8("Uint8Ptr", *l.Uint8Ptr)
	}
	enc.AddBinary("Bytes", l.Bytes)
	enc.AddUintptr("UintPtrVal", l.UintPtrVal)
	if l.UintPtrValPtr != nil {
		enc.AddUintptr("UintPtrValPtr", *l.UintPtrValPtr)
	}
	enc.AddTime("Time", l.Time)
	if l.TimePtr != nil {
		enc.AddTime("TimePtr", *l.TimePtr)
	}
	enc.AddDuration("Duration", l.Duration)
	if l.DurationPtr != nil {
		enc.AddDuration("DurationPtr", *l.DurationPtr)
	}
	return err
}
func (l *OmitEmptySlice) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	var err error
	if len(l.OtherSlice) > 0 {
		err = enc.AddArray("OtherSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.OtherSlice {

				if v, ok := any(v).(zapcore.ObjectMarshaler); ok {
					err = enc.AppendObject(v)
					if err != nil {
						return err
					}
				} else {
					err = enc.AppendReflected(v)
					if err != nil {
						return err
					}
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.BoolSlice) > 0 {
		err = enc.AddArray("BoolSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.BoolSlice {
				enc.AppendBool(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.BoolPtrSlice) > 0 {
		err = enc.AddArray("BoolPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.BoolPtrSlice {
				if l.BoolPtrSlice != nil {
					enc.AppendBool(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Complex128Slice) > 0 {
		err = enc.AddArray("Complex128Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Complex128Slice {
				enc.AppendComplex128(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Complex128PtrSlice) > 0 {
		err = enc.AddArray("Complex128PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Complex128PtrSlice {
				if l.Complex128PtrSlice != nil {
					enc.AppendComplex128(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Complex64Slice) > 0 {
		err = enc.AddArray("Complex64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Complex64Slice {
				enc.AppendComplex64(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Complex64PtrSlice) > 0 {
		err = enc.AddArray("Complex64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Complex64PtrSlice {
				if l.Complex64PtrSlice != nil {
					enc.AppendComplex64(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Float64Slice) > 0 {
		err = enc.AddArray("Float64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Float64Slice {
				enc.AppendFloat64(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Float64PtrSlice) > 0 {
		err = enc.AddArray("Float64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Float64PtrSlice {
				if l.Float64PtrSlice != nil {
					enc.AppendFloat64(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Float32Slice) > 0 {
		err = enc.AddArray("Float32Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Float32Slice {
				enc.AppendFloat32(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Float32PtrSlice) > 0 {
		err = enc.AddArray("Float32PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Float32PtrSlice {
				if l.Float32PtrSlice != nil {
					enc.AppendFloat32(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.IntSlice) > 0 {
		err = enc.AddArray("IntSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.IntSlice {
				enc.AppendInt(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.IntPtrSlice) > 0 {
		err = enc.AddArray("IntPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.IntPtrSlice {
				if l.IntPtrSlice != nil {
					enc.AppendInt(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int64Slice) > 0 {
		err = enc.AddArray("Int64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int64Slice {
				enc.AppendInt64(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int64PtrSlice) > 0 {
		err = enc.AddArray("Int64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int64PtrSlice {
				if l.Int64PtrSlice != nil {
					enc.AppendInt64(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int32Slice) > 0 {
		err = enc.AddArray("Int32Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int32Slice {
				enc.AppendInt32(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int32PtrSlice) > 0 {
		err = enc.AddArray("Int32PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int32PtrSlice {
				if l.Int32PtrSlice != nil {
					enc.AppendInt32(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int16Slice) > 0 {
		err = enc.AddArray("Int16Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int16Slice {
				enc.AppendInt16(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int16PtrSlice) > 0 {
		err = enc.AddArray("Int16PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int16PtrSlice {
				if l.Int16PtrSlice != nil {
					enc.AppendInt16(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int8Slice) > 0 {
		err = enc.AddArray("Int8Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int8Slice {
				enc.AppendInt8(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Int8PtrSlice) > 0 {
		err = enc.AddArray("Int8PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Int8PtrSlice {
				if l.Int8PtrSlice != nil {
					enc.AppendInt8(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.StringSlice) > 0 {
		err = enc.AddArray("StringSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.StringSlice {
				enc.AppendString(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.StringPtrSlice) > 0 {
		err = enc.AddArray("StringPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.StringPtrSlice {
				if l.StringPtrSlice != nil {
					enc.AppendString(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.UintSlice) > 0 {
		err = enc.AddArray("UintSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.UintSlice {
				enc.AppendUint(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.UintPtrSlice) > 0 {
		err = enc.AddArray("UintPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.UintPtrSlice {
				if l.UintPtrSlice != nil {
					enc.AppendUint(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint64Slice) > 0 {
		err = enc.AddArray("Uint64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint64Slice {
				enc.AppendUint64(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint64PtrSlice) > 0 {
		err = enc.AddArray("Uint64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint64PtrSlice {
				if l.Uint64PtrSlice != nil {
					enc.AppendUint64(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint32Slice) > 0 {
		err = enc.AddArray("Uint32Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint32Slice {
				enc.AppendUint32(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint32PtrSlice) > 0 {
		err = enc.AddArray("Uint32PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint32PtrSlice {
				if l.Uint32PtrSlice != nil {
					enc.AppendUint32(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint16Slice) > 0 {
		err = enc.AddArray("Uint16Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint16Slice {
				enc.AppendUint16(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint16PtrSlice) > 0 {
		err = enc.AddArray("Uint16PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint16PtrSlice {
				if l.Uint16PtrSlice != nil {
					enc.AppendUint16(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint8Slice) > 0 {
		err = enc.AddArray("Uint8Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint8Slice {
				enc.AppendUint8(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.Uint8PtrSlice) > 0 {
		err = enc.AddArray("Uint8PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.Uint8PtrSlice {
				if l.Uint8PtrSlice != nil {
					enc.AppendUint8(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.BytesSlice) > 0 {
		err = enc.AddArray("BytesSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.BytesSlice {
				enc.AppendByteString(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.UintPtrValSlice) > 0 {
		err = enc.AddArray("UintPtrValSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.UintPtrValSlice {
				enc.AppendUintptr(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.UintPtrValPtrSlice) > 0 {
		err = enc.AddArray("UintPtrValPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.UintPtrValPtrSlice {
				if l.UintPtrValPtrSlice != nil {
					enc.AppendUintptr(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.TimeSlice) > 0 {
		err = enc.AddArray("TimeSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.TimeSlice {
				enc.AppendTime(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.TimePtrSlice) > 0 {
		err = enc.AddArray("TimePtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.TimePtrSlice {
				if l.TimePtrSlice != nil {
					enc.AppendTime(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.DurationSlice) > 0 {
		err = enc.AddArray("DurationSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.DurationSlice {
				enc.AppendDuration(v)
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	if len(l.DurationPtrSlice) > 0 {
		err = enc.AddArray("DurationPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, v := range l.DurationPtrSlice {
				if l.DurationPtrSlice != nil {
					enc.AppendDuration(*v)
				}
			}
			return nil
		}))
		if err != nil {
			return err
		}
	}
	return err
}
func (l *Slice) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	var err error
	err = enc.AddArray("OtherSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.OtherSlice {

			if v, ok := any(v).(zapcore.ObjectMarshaler); ok {
				err = enc.AppendObject(v)
				if err != nil {
					return err
				}
			} else {
				err = enc.AppendReflected(v)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("BoolSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.BoolSlice {
			enc.AppendBool(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("BoolPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.BoolPtrSlice {
			if l.BoolPtrSlice != nil {
				enc.AppendBool(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Complex128Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Complex128Slice {
			enc.AppendComplex128(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Complex128PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Complex128PtrSlice {
			if l.Complex128PtrSlice != nil {
				enc.AppendComplex128(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Complex64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Complex64Slice {
			enc.AppendComplex64(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Complex64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Complex64PtrSlice {
			if l.Complex64PtrSlice != nil {
				enc.AppendComplex64(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Float64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Float64Slice {
			enc.AppendFloat64(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Float64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Float64PtrSlice {
			if l.Float64PtrSlice != nil {
				enc.AppendFloat64(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Float32Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Float32Slice {
			enc.AppendFloat32(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Float32PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Float32PtrSlice {
			if l.Float32PtrSlice != nil {
				enc.AppendFloat32(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("IntSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.IntSlice {
			enc.AppendInt(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("IntPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.IntPtrSlice {
			if l.IntPtrSlice != nil {
				enc.AppendInt(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int64Slice {
			enc.AppendInt64(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int64PtrSlice {
			if l.Int64PtrSlice != nil {
				enc.AppendInt64(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int32Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int32Slice {
			enc.AppendInt32(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int32PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int32PtrSlice {
			if l.Int32PtrSlice != nil {
				enc.AppendInt32(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int16Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int16Slice {
			enc.AppendInt16(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int16PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int16PtrSlice {
			if l.Int16PtrSlice != nil {
				enc.AppendInt16(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int8Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int8Slice {
			enc.AppendInt8(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Int8PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Int8PtrSlice {
			if l.Int8PtrSlice != nil {
				enc.AppendInt8(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("StringSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.StringSlice {
			enc.AppendString(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("StringPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.StringPtrSlice {
			if l.StringPtrSlice != nil {
				enc.AppendString(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("UintSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.UintSlice {
			enc.AppendUint(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("UintPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.UintPtrSlice {
			if l.UintPtrSlice != nil {
				enc.AppendUint(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint64Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint64Slice {
			enc.AppendUint64(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint64PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint64PtrSlice {
			if l.Uint64PtrSlice != nil {
				enc.AppendUint64(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint32Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint32Slice {
			enc.AppendUint32(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint32PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint32PtrSlice {
			if l.Uint32PtrSlice != nil {
				enc.AppendUint32(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint16Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint16Slice {
			enc.AppendUint16(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint16PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint16PtrSlice {
			if l.Uint16PtrSlice != nil {
				enc.AppendUint16(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint8Slice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint8Slice {
			enc.AppendUint8(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("Uint8PtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.Uint8PtrSlice {
			if l.Uint8PtrSlice != nil {
				enc.AppendUint8(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("BytesSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.BytesSlice {
			enc.AppendByteString(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("UintPtrValSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.UintPtrValSlice {
			enc.AppendUintptr(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("UintPtrValPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.UintPtrValPtrSlice {
			if l.UintPtrValPtrSlice != nil {
				enc.AppendUintptr(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("TimeSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.TimeSlice {
			enc.AppendTime(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("TimePtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.TimePtrSlice {
			if l.TimePtrSlice != nil {
				enc.AppendTime(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("DurationSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.DurationSlice {
			enc.AppendDuration(v)
		}
		return nil
	}))
	if err != nil {
		return err
	}
	err = enc.AddArray("DurationPtrSlice", zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
		for _, v := range l.DurationPtrSlice {
			if l.DurationPtrSlice != nil {
				enc.AppendDuration(*v)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	return err
}
