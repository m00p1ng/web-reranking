package pagerank

import (
	"log"
	"math"
)

const (
	alpha   = 0.85
	epsilon = 1e-5
)

// Compute -- compute pagerank
func Compute(wg [][]int) []float64 {
	log.Println("Initializing PageRank...")
	r0 := initValue(len(wg), 1)
	r1 := make([]float64, len(wg))
	tp := initValue(len(wg), 1-alpha)
	t := 0

	vecA := calcVectA(wg)

	log.Println("Calculating PageRank...")

	cvg := 0.0
	cnt := 0
	for {
		t++
		r1 = mulMatrix(vecA, r0)
		r1 = sumMatrix(r1, tp)
		dist := distance(r0, r1)
		sumr := sumRank(r1)
		log.Printf("Round %d (err=%.5f, sumR=%.5f)", t, dist, sumr)

		ncvg := math.Floor(dist / epsilon)
		if ncvg == cvg {
			cnt++
		} else {
			cvg = ncvg
			cnt = 0
		}
		if dist < epsilon || cnt == 10 {
			for i := range r1 {
				r1[i] = r1[i] / sumr
			}
			break
		}
		r0 = r1
	}
	log.Println("Finished")

	return r1
}

func initValue(dim int, mul float64) []float64 {
	a := make([]float64, dim)
	for i := range a {
		a[i] = mul / float64(dim)
	}
	return a
}

func distance(a []float64, b []float64) float64 {
	sum := 0.0

	for i := range a {
		sum += (a[i] - b[i]) * (a[i] - b[i])
	}

	return math.Sqrt(sum)
}

func calcVectA(wg [][]int) [][]float64 {
	sz := len(wg)

	log.Println("Initializing VectorA...")
	vecA := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		vecA[i] = initValue(sz, 0)
	}
	log.Println("VectorA was Initialized")

	log.Println("Calculating VectorA...")
	for i, row := range wg {
		var val float64
		if len(row) > 0 {
			val = alpha / float64(len(row))
		} else {
			val = 0
		}
		for _, idx := range row {
			vecA[i][idx-1] += val
		}
	}

	vecA = transposeMatrix(vecA)
	log.Println("VectorA was calculated")

	return vecA
}

func transposeMatrix(m [][]float64) [][]float64 {
	sz := len(m)
	for i := 0; i < sz; i++ {
		for j := 0; j < i; j++ {
			temp := m[j][i]
			m[j][i] = m[i][j]
			m[i][j] = temp
		}
	}

	return m
}

func mulMatrix(a [][]float64, b []float64) []float64 {
	sz := len(b)
	result := make([]float64, sz)

	for i, row := range a {
		sum := 0.0
		for j := range row {
			sum += b[j] * row[j]
		}
		result[i] = sum
	}

	return result
}

func sumRank(a []float64) float64 {
	sum := 0.0
	for _, t := range a {
		sum += t
	}
	return sum
}

func sumMatrix(a []float64, b []float64) []float64 {
	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}
	return result
}
