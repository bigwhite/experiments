package main

import "testing"

func BenchmarkAoSvsSoA(b *testing.B) {
	b.Run("AoS (Sum Health) - Large", func(b *testing.B) {
		entities := make([]EntityAoS, numEntities)
		for i := range entities {
			entities[i].Health = i
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SumHealthAoS(entities)
		}
	})

	b.Run("SoA (Sum Health) - Large", func(b *testing.B) {
		world := NewWorldSoA(numEntities)
		for i := range world.Healths {
			world.Healths[i] = i
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SumHealthSoA(world)
		}
	})
}
