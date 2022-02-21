package main

const (
	VendorOsTypeAndroid = 2
	VendorOsTypeIOS     = 1
)
const (
	FancyOsTypeAndroid = "android"
	FancyOsTypeIOS     = "ios"
)

var (
	OsTypeMapping = map[int]string{
		VendorOsTypeAndroid: FancyOsTypeAndroid,
		VendorOsTypeIOS:     FancyOsTypeIOS,
	}
)

var (
	ToOsTypeMapping = map[string]int{
		FancyOsTypeAndroid: VendorOsTypeAndroid,
		FancyOsTypeIOS:     VendorOsTypeIOS,
	}
)

const (
	VendorSlotTypeImage = "image"
	VendorSlotTypeVideo = "video"
)
const (
	FancySlotTypeImage = "image/mp4"
	FancySlotTypeVideo = "video/mp4"
)

var (
	SlotTypeMapping = map[string]string{
		VendorSlotTypeImage: FancySlotTypeImage,
		VendorSlotTypeVideo: FancySlotTypeVideo,
	}
)
