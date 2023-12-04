package main

import (
	"log"
	"math"
)

type Cluster struct {
	Left     *Cluster
	Right    *Cluster
	Blog     *Blog
	Distance float64
	Parent   *Cluster
}

func initializeClusters(blogs []Blog) []*Cluster {
	clusters := make([]*Cluster, len(blogs))
	for i := range blogs {
			clusters[i] = &Cluster{Blog: &blogs[i]} // Directly reference the blog
	}
	return clusters
}

func mergeClusters(clusterA, clusterB *Cluster, distance float64) *Cluster {
	if clusterA == nil || clusterB == nil {
			log.Println("Attempted to merge nil clusters")
			return nil
	}

	// Creating a new cluster without a blog reference as it's a merged cluster
	newCluster := &Cluster{
			Left:     clusterA,
			Right:    clusterB,
			Distance: distance,
	}
	clusterA.Parent = newCluster
	clusterB.Parent = newCluster
	return newCluster
}

func findClosestClusters(clusters []*Cluster, merged map[*Cluster]bool) (*Cluster, *Cluster, float64) {
	closestDistance := math.MaxFloat64
	var clusterA, clusterB *Cluster

	for i, cA := range clusters {
			if merged[cA] {
					continue
			}
			for j, cB := range clusters {
					if merged[cB] || i == j {
							continue
					}

					distance := calculateClusterDistance(cA, cB)
					if distance < closestDistance {
							closestDistance = distance
							clusterA = cA
							clusterB = cB
					}
			}
	}

	if clusterA == nil || clusterB == nil {
			log.Println("findClosestClusters: No valid cluster pair found")
			return nil, nil, math.MaxFloat64
	}

	return clusterA, clusterB, closestDistance
}

func calculateClusterDistance(clusterA, clusterB *Cluster) float64 {
	// If both clusters are not merged, use direct blog comparison
	if clusterA.Blog != nil && clusterB.Blog != nil {
			return pearsonDistanceForHierarchical(clusterA.Blog, clusterB.Blog)
	}

	// Handle merged clusters by calculating average distance
	var totalDistance, count float64
	forEachBlogInCluster(clusterA, func(blogA *Blog) {
			forEachBlogInCluster(clusterB, func(blogB *Blog) {
					totalDistance += pearsonDistanceForHierarchical(blogA, blogB)
					count++
			})
	})

	if count == 0 {
			return 0
	}
	return totalDistance / count
}

func forEachBlogInCluster(cluster *Cluster, action func(*Blog)) {
	if cluster == nil {
			return
	}
	if cluster.Blog != nil {
			action(cluster.Blog)
	} else {
			forEachBlogInCluster(cluster.Left, action)
			forEachBlogInCluster(cluster.Right, action)
	}
}

func pearsonDistanceForHierarchical(blogA, blogB *Blog) float64 {
	if blogA == nil || blogB == nil {
		log.Println("Attempted to calculate Pearson distance with nil blog")
		return 0
	}

	sumA, sumB, sumAsq, sumBsq, pSum := 0.0, 0.0, 0.0, 0.0, 0.0
	n := float64(len(blogA.WordCounts))

	for i := 0; i < len(blogA.WordCounts); i++ {
		cntA := float64(blogA.WordCounts[i])
		cntB := float64(blogB.WordCounts[i])
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

func hierarchicalClustering(blogs []Blog) *Cluster {
	clusters := initializeClusters(blogs)
	merged := make(map[*Cluster]bool)

	for len(clusters) > 1 {
			clusterA, clusterB, distance := findClosestClusters(clusters, merged)


		if clusterA == nil || clusterB == nil || clusterA == clusterB {
			log.Println("No more valid clusters to merge, stopping the clustering")
			break
		}

		newCluster := mergeClusters(clusterA, clusterB, distance)
		if newCluster == nil {
			log.Println("Failed to merge clusters, stopping the clustering")
			break
		}

		// Mark the merged clusters
		merged[clusterA] = true
		merged[clusterB] = true

		// Add the new cluster
		clusters = append(clusters, newCluster)
	}

	// Find the root cluster (the one that hasn't been merged)
	for _, cluster := range clusters {
		if !merged[cluster] {
			return cluster
		}
	}

	return nil // In case no root cluster is found
}
