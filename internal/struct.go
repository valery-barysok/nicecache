package internal

//TestValue some test type
type TestValue struct {
	ID           string
	N            uint32
	Stat         uint32
	Published    bool
	Deprecated   *bool
	System       uint32
	Subsystem    uint32
	ParentID     string
	Name         string
	Name2        string
	Description  string
	Description2 string
	CreatedBy    uint32
	UpdatedBy    *uint32
	List1        []uint32
	List2        []uint32
	Items        []TestItem
}

//TestItem some test type to test cache in different package
type TestItem struct {
	ID           string
	N            uint32
	Stat         uint32
	Published    bool
	Deprecated   *bool
	System       uint32
	Subsystem    uint32
	ParentID     string
	Name         string
	Name2        string
	Description  string
	Description2 string
	CreatedBy    uint32
	UpdatedBy    *uint32
	List1        []uint32
	List2        []uint32
}
