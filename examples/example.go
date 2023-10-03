package examples

//go:generate go run github.com/youta1119/zap-obj-marshaler-gen/cmd -f $GOFILE
import (
	"time"
)

type Example struct {
	Other         Other
	OtherPtr      *Other
	Bool          bool
	BoolPtr       *bool
	Complex128    complex128
	Complex128Ptr *complex128
	Complex64     complex64
	Complex64Ptr  *complex64
	Float64       float64
	Float64Ptr    *float64
	Float32       float32
	Float32Ptr    *float32
	Int           int
	IntPtr        *int
	Int64         int64
	Int64Ptr      *int64
	Int32         int32
	Int32Ptr      *int32
	Int16         int16
	Int16Ptr      *int16
	Int8          int8
	Int8Ptr       *int8
	String        string
	StringPtr     *string
	Uint          uint
	UintPtr       *uint
	Uint64        uint64
	Uint64Ptr     *uint64
	Uint32        uint32
	Uint32Ptr     *uint32
	Uint16        uint16
	Uint16Ptr     *uint16
	Uint8         uint8
	Uint8Ptr      *uint8
	Bytes         []byte
	UintPtrVal    uintptr
	UintPtrValPtr *uintptr
	Time          time.Time
	TimePtr       *time.Time
	Duration      time.Duration
	DurationPtr   *time.Duration
	CustomStr     string `zap:"custom_str"`
	CustomBool    bool   `zap:"custom_bool"`
}

type Slice struct {
	OtherSlice         []*Other
	BoolSlice          []bool
	BoolPtrSlice       []*bool
	Complex128Slice    []complex128
	Complex128PtrSlice []*complex128
	Complex64Slice     []complex64
	Complex64PtrSlice  []*complex64
	Float64Slice       []float64
	Float64PtrSlice    []*float64
	Float32Slice       []float32
	Float32PtrSlice    []*float32
	IntSlice           []int
	IntPtrSlice        []*int
	Int64Slice         []int64
	Int64PtrSlice      []*int64
	Int32Slice         []int32
	Int32PtrSlice      []*int32
	Int16Slice         []int16
	Int16PtrSlice      []*int16
	Int8Slice          []int8
	Int8PtrSlice       []*int8
	StringSlice        []string
	StringPtrSlice     []*string
	UintSlice          []uint
	UintPtrSlice       []*uint
	Uint64Slice        []uint64
	Uint64PtrSlice     []*uint64
	Uint32Slice        []uint32
	Uint32PtrSlice     []*uint32
	Uint16Slice        []uint16
	Uint16PtrSlice     []*uint16
	Uint8Slice         []uint8
	Uint8PtrSlice      []*uint8
	BytesSlice         [][]byte
	UintPtrValSlice    []uintptr
	UintPtrValPtrSlice []*uintptr
	TimeSlice          []time.Time
	TimePtrSlice       []*time.Time
	DurationSlice      []time.Duration
	DurationPtrSlice   []*time.Duration
}

type OmitEmpty struct {
	Other         *Other         `zap:",omitempty"`
	Bool          bool           `zap:",omitempty"`
	BoolPtr       *bool          `zap:",omitempty"`
	Complex128    complex128     `zap:",omitempty"`
	Complex128Ptr *complex128    `zap:",omitempty"`
	Complex64     complex64      `zap:",omitempty"`
	Complex64Ptr  *complex64     `zap:",omitempty"`
	Float64       float64        `zap:",omitempty"`
	Float64Ptr    *float64       `zap:",omitempty"`
	Float32       float32        `zap:",omitempty"`
	Float32Ptr    *float32       `zap:",omitempty"`
	Int           int            `zap:",omitempty"`
	IntPtr        *int           `zap:",omitempty"`
	Int64         int64          `zap:",omitempty"`
	Int64Ptr      *int64         `zap:",omitempty"`
	Int32         int32          `zap:",omitempty"`
	Int32Ptr      *int32         `zap:",omitempty"`
	Int16         int16          `zap:",omitempty"`
	Int16Ptr      *int16         `zap:",omitempty"`
	Int8          int8           `zap:",omitempty"`
	Int8Ptr       *int8          `zap:",omitempty"`
	String        string         `zap:",omitempty"`
	StringPtr     *string        `zap:",omitempty"`
	Uint          uint           `zap:",omitempty"`
	UintPtr       *uint          `zap:",omitempty"`
	Uint64        uint64         `zap:",omitempty"`
	Uint64Ptr     *uint64        `zap:",omitempty"`
	Uint32        uint32         `zap:",omitempty"`
	Uint32Ptr     *uint32        `zap:",omitempty"`
	Uint16        uint16         `zap:",omitempty"`
	Uint16Ptr     *uint16        `zap:",omitempty"`
	Uint8         uint8          `zap:",omitempty"`
	Uint8Ptr      *uint8         `zap:",omitempty"`
	Bytes         []byte         `zap:",omitempty"`
	UintPtrVal    uintptr        `zap:",omitempty"`
	UintPtrValPtr *uintptr       `zap:",omitempty"`
	Time          time.Time      `zap:",omitempty"`
	TimePtr       *time.Time     `zap:",omitempty"`
	Duration      time.Duration  `zap:",omitempty"`
	DurationPtr   *time.Duration `zap:",omitempty"`
}

type OmitEmptySlice struct {
	OtherSlice         []*Other         `zap:",omitempty"`
	BoolSlice          []bool           `zap:",omitempty"`
	BoolPtrSlice       []*bool          `zap:",omitempty"`
	Complex128Slice    []complex128     `zap:",omitempty"`
	Complex128PtrSlice []*complex128    `zap:",omitempty"`
	Complex64Slice     []complex64      `zap:",omitempty"`
	Complex64PtrSlice  []*complex64     `zap:",omitempty"`
	Float64Slice       []float64        `zap:",omitempty"`
	Float64PtrSlice    []*float64       `zap:",omitempty"`
	Float32Slice       []float32        `zap:",omitempty"`
	Float32PtrSlice    []*float32       `zap:",omitempty"`
	IntSlice           []int            `zap:",omitempty"`
	IntPtrSlice        []*int           `zap:",omitempty"`
	Int64Slice         []int64          `zap:",omitempty"`
	Int64PtrSlice      []*int64         `zap:",omitempty"`
	Int32Slice         []int32          `zap:",omitempty"`
	Int32PtrSlice      []*int32         `zap:",omitempty"`
	Int16Slice         []int16          `zap:",omitempty"`
	Int16PtrSlice      []*int16         `zap:",omitempty"`
	Int8Slice          []int8           `zap:",omitempty"`
	Int8PtrSlice       []*int8          `zap:",omitempty"`
	StringSlice        []string         `zap:",omitempty"`
	StringPtrSlice     []*string        `zap:",omitempty"`
	UintSlice          []uint           `zap:",omitempty"`
	UintPtrSlice       []*uint          `zap:",omitempty"`
	Uint64Slice        []uint64         `zap:",omitempty"`
	Uint64PtrSlice     []*uint64        `zap:",omitempty"`
	Uint32Slice        []uint32         `zap:",omitempty"`
	Uint32PtrSlice     []*uint32        `zap:",omitempty"`
	Uint16Slice        []uint16         `zap:",omitempty"`
	Uint16PtrSlice     []*uint16        `zap:",omitempty"`
	Uint8Slice         []uint8          `zap:",omitempty"`
	Uint8PtrSlice      []*uint8         `zap:",omitempty"`
	BytesSlice         [][]byte         `zap:",omitempty"`
	UintPtrValSlice    []uintptr        `zap:",omitempty"`
	UintPtrValPtrSlice []*uintptr       `zap:",omitempty"`
	TimeSlice          []time.Time      `zap:",omitempty"`
	TimePtrSlice       []*time.Time     `zap:",omitempty"`
	DurationSlice      []time.Duration  `zap:",omitempty"`
	DurationPtrSlice   []*time.Duration `zap:",omitempty"`
}
