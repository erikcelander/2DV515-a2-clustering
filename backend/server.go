package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"


	
)

type CombinedClusterResult struct {
	KMeans       ClusterResult   `json:"kmeans"`
	Hierarchical []ClusterNode   `json:"hierarchical"`
}

type ClusterResult struct {
	Cluster1 []string `json:"cluster1"`
	Cluster2 []string `json:"cluster2"`
	Cluster3 []string `json:"cluster3"`
	Cluster4 []string `json:"cluster4"`
	Cluster5 []string `json:"cluster5"`
}

type ClusterNode struct {
	Name     string        `json:"name,omitempty"`
	Children []ClusterNode `json:"children,omitempty"`
}



func main() {
	http.HandleFunc("/api/clusters", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		blogs, err := readBlogsFromFile("/Users/erik/Desktop/school/web-intelligence-2DV515/a2-clustering/backend/blogdata.txt")
		if err != nil {
			http.Error(w, "Error reading blogs: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// K-means clustering
		kmeansStart := time.Now()
		centroids := initializeCentroids(numClusters)
		for i := 0; i < maxIterations; i++ {
			clearAssignments(centroids)
			assignBlogsToCentroids(blogs, centroids)

		
			if checkConvergence(centroids) {
				break
			}
            
			updateCentroids(centroids)
		}
		kmeansResult := prepareClusterResult(centroids)
		kmeansTime := time.Since(kmeansStart)
		log.Printf("K-means clustering took %v ms", kmeansTime.Milliseconds())


		// Hierarchical clustering
		hierarchicalStart := time.Now()
		hierarchicalRoot := hierarchicalClustering(blogs)
		hierarchicalResult := prepareHierarchicalResult(hierarchicalRoot)
		hierarchicalTime := time.Since(hierarchicalStart)
		log.Printf("Hierarchical clustering took %v ms", hierarchicalTime.Milliseconds())

		combinedResult := CombinedClusterResult{
			KMeans:       kmeansResult,
			Hierarchical: hierarchicalResult,
		}

		jsonResult, err := json.MarshalIndent(combinedResult, "", "    ")
		if err != nil {
			http.Error(w, "Error generating JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}


		totalTime := time.Since(startTime)
		log.Printf("Total time: %v ms", totalTime.Milliseconds())

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	})

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func prepareClusterResult(centroids []*Centroid) ClusterResult {
    result := ClusterResult{}
    for i, c := range centroids {
        for _, b := range c.Assignments {
            switch i {
            case 0:
                result.Cluster1 = append(result.Cluster1, b.Name)
            case 1:
                result.Cluster2 = append(result.Cluster2, b.Name)
            case 2:
                result.Cluster3 = append(result.Cluster3, b.Name)
            case 3:
                result.Cluster4 = append(result.Cluster4, b.Name)
            case 4:
                result.Cluster5 = append(result.Cluster5, b.Name)
            }
        }
    }
    return result
}


// prepareHierarchicalResult converts the hierarchical cluster tree to a slice of ClusterNode for JSON serialization.
func prepareHierarchicalResult(cluster *Cluster) []ClusterNode {
	if cluster == nil {
		return nil
	}

	var result []ClusterNode

	// Recursive function to traverse the cluster tree.
	var traverse func(c *Cluster) ClusterNode
	traverse = func(c *Cluster) ClusterNode {
		if c.Blog != nil {
			// Leaf node with a blog.
			return ClusterNode{Name: c.Blog.Name}
		}

		// Non-leaf node, recursively process children.
		children := []ClusterNode{}
		if c.Left != nil {
			children = append(children, traverse(c.Left))
		}
		if c.Right != nil {
			children = append(children, traverse(c.Right))
		}

		return ClusterNode{Children: children}
	}

	// Start the traversal from the root cluster.
	rootNode := traverse(cluster)
	result = append(result, rootNode)

	return result
}
