package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type ClusterResult struct {
    Cluster1 []string `json:"cluster1"`
    Cluster2 []string `json:"cluster2"`
    Cluster3 []string `json:"cluster3"`
    Cluster4 []string `json:"cluster4"`
    Cluster5 []string `json:"cluster5"`
}

func main() {
    blogs, err := readBlogsFromFile("./blogdata.txt")
    if err != nil {
        log.Fatal("Error reading blogs:", err)
    }

    centroids := initializeCentroids(numClusters)
    for i := 0; i < maxIterations; i++ {
        clearAssignments(centroids)
        assignBlogsToCentroids(blogs, centroids)
        updateCentroids(centroids)
    }

    StartServer(centroids)
}

func StartServer(centroids []*Centroid) {
	http.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			result := prepareClusterResult(centroids)
			
			jsonResult, err := json.MarshalIndent(result, "", "    ")
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}

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
