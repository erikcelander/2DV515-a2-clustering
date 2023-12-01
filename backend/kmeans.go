package main

import (
	"fmt"
	"math"
	"math/rand"
)


type Centroid struct {
	WordCounts []float64
	Assignments []*Blog
}

const (
	NumWords         = 706 // Total number of words
	NumClusters      = 5   // Number of clusters
	MaxIterations    = 10  // Maximum iterations for K-means
)

func main() {
	blogs, err := readBlogsFromFile("./blogdata.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	centroids := initializeCentroids(NumClusters)

	// K-means Clustering
	for i := 0; i < MaxIterations; i++ {
		clearAssignments(centroids)
		assignBlogsToCentroids(blogs, centroids)
		updateCentroids(centroids)
	}

	printClusters(centroids)
}

func initializeCentroids(k int) []*Centroid {
	centroids := make([]*Centroid, k)

	for i := range centroids {
		centroid := &Centroid{WordCounts: make([]float64, NumWords)}
		for j := range centroid.WordCounts {
			centroid.WordCounts[j] = rand.Float64()
		}
		centroids[i] = centroid
	}

	return centroids
}

func clearAssignments(centroids []*Centroid) {
	for _, c := range centroids {
		c.Assignments = nil
	}
}

func assignBlogsToCentroids(blogs []Blog, centroids []*Centroid) {
	for i, blog := range blogs {
		var bestCentroid *Centroid
		minDistance := math.MaxFloat64

		for _, centroid := range centroids {
			distance := pearsonDistance(centroid, &blog)
			if distance < minDistance {
				minDistance = distance
				bestCentroid = centroid
			}
		}

		blogCopy := blogs[i]
		bestCentroid.Assignments = append(bestCentroid.Assignments, &blogCopy)
	}
}


func findCentroidIndex(centroid *Centroid, centroids []*Centroid) int {
	for i, c := range centroids {
		if c == centroid {
			return i
		}
	}
	return -1
}

func updateCentroids(centroids []*Centroid) {
	for _, c := range centroids {
		if len(c.Assignments) == 0 {
			continue
		}

		sumWordCounts := make([]float64, NumWords)
		for _, blog := range c.Assignments {
			for j, count := range blog.WordCounts {
				sumWordCounts[j] += float64(count)
			}
		}

		for j := range c.WordCounts {
			c.WordCounts[j] = sumWordCounts[j] / float64(len(c.Assignments))
		}

	}
}


func pearsonDistance(centroid *Centroid, blog *Blog) float64 {
	sumA, sumB, sumAsq, sumBsq, pSum := 0.0, 0.0, 0.0, 0.0, 0.0
	n := float64(len(centroid.WordCounts)) 

	for i := 0; i < len(centroid.WordCounts); i++ {
		cntA := centroid.WordCounts[i]
		cntB := float64(blog.WordCounts[i])
		sumA += cntA
		sumB += cntB
		sumAsq += cntA * cntA
		sumBsq += cntB * cntB
		pSum += cntA * cntB
	}

	num := pSum - (sumA * sumB / n)
	den := math.Sqrt((sumAsq - (sumA * sumA / n)) * (sumBsq - (sumB * sumB / n)))

	if den == 0 {
		return 0
	}

	return 1.0 - num/den
}

func printClusters(centroids []*Centroid) {
	for i, c := range centroids {
		fmt.Printf("Cluster %d:\n", i)
		for _, b := range c.Assignments {
			fmt.Printf("  Blog: %s\n", b.Name)
		}
	}
}
