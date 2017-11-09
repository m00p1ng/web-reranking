package pagerank

import (
	"log"
	"math"
	"time"
)

const (
	alpha   = 0.85
	epsilon = 1e-5
)

// Compute -- compute pagerank
func Compute(wg [][]int) []float64 {
	log.Println("Initializing PageRank...")
	x0 := initValue(len(wg), 1)
	x1 := make([]float64, len(wg))
	tp := initValue(len(wg), 1-alpha)
	round := 0

	matP := calcMatPTrans(wg)

	log.Println("Calculating PageRank...")

	for {
		round++
		start := time.Now()
		x1 = mulMatrix(matP, x0)
		x1 = sumMatrix(x1, tp)
		dist := distance(x0, x1)
		sumr := sumRank(x1)
		elapsed := time.Now().Sub(start).Seconds()
		log.Printf("Round %d (err=%.5f, time=%.5f, sum_rank=%.5f)", round, dist, elapsed, sumr)

		if dist < epsilon {
			break
		}
		x0 = x1
	}
	log.Println("Finished")

	return x1
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

func calcMatPTrans(wg [][]int) [][]float64 {
	sz := len(wg)

	log.Println("Initializing mat(P)...")
	matP := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		matP[i] = initValue(sz, 0)
	}
	log.Println("mat(P) was Initialized")

	log.Println("Calculating mat(p)...")

	totalDoc := len(wg)
	for i, row := range wg {
		var val float64
		outLink := len(row)

		if outLink > 0 {
			val = alpha / float64(outLink)
			for _, j := range row {
				matP[i][j-1] += val
			}
		} else {
			val = alpha / float64(totalDoc)
			for j := 0; j < totalDoc; j++ {
				matP[i][j] = val
			}
		}

	}

	matPTrans := transposeMatrix(matP)
	log.Println("mat(p).transpose() was calculated")

	return matPTrans
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
