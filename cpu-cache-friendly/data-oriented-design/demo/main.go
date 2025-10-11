package main

const (
	// 将实体数量增加到 1M，确保工作集大于大多数 CPU 的 L3 缓存
	numEntities = 1024 * 1024
)

// --- AoS (Array of Structs): 缓存不友好 ---
type EntityAoS struct {
	// 假设这是一个更复杂的结构体
	ID       uint64
	Health   int
	Position [3]float64
	// ... 更多字段
}

func SumHealthAoS(entities []EntityAoS) int {
	var totalHealth int
	for i := range entities {
		// 每次循环，CPU 都必须加载整个庞大的 EntityAoS 结构体，
		// 即使我们只用到了 Health 这一个字段。
		totalHealth += entities[i].Health
	}
	return totalHealth
}

// --- SoA (Struct of Arrays): 缓存的挚友 ---
type WorldSoA struct {
	IDs       []uint64
	Healths   []int
	Positions [][3]float64
	// ... 更多字段的切片
}

func NewWorldSoA(n int) *WorldSoA {
	return &WorldSoA{
		IDs:       make([]uint64, n),
		Healths:   make([]int, n),
		Positions: make([][3]float64, n),
	}
}

func SumHealthSoA(world *WorldSoA) int {
	var totalHealth int
	// 这个循环只访问 Healths 切片，数据完美连续。
	for i := range world.Healths {
		totalHealth += world.Healths[i]
	}
	return totalHealth
}
