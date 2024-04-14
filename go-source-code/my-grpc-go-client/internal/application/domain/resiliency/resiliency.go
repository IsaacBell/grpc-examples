package resiliency

const (
	OK                 uint32 = 0
	CANCELLED          uint32 = 1
	UNKNOWN            uint32 = 2
	INVALID_ARGUMENT   uint32 = 3
	DEADLINE_EXCEEDED  uint32 = 4
	NOT_FOUND          uint32 = 5
	ALREADY_EXISTS     uint32 = 6
	PERMISSION_DENIED  uint32 = 7
	RESOURCE_EXHAUSTED uint32 = 8
)
