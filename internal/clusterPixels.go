package picToBric

import (
	"image"
	"log"
	"math"
	"math/rand"
)

type Pixel struct {
	r       int16
	g       int16
	b       int16
	cluster int
}

func squared(i int16) float64 {
	return float64(i * i)
}

func (px Pixel) distance(other Pixel) float64 {
	return math.Sqrt(squared(px.r-other.r) + squared(px.g-other.g) + squared(px.b-other.b))
}

func (px Pixel) nearestCluster(centres []Pixel) int {
	distance := math.MaxFloat64
	var clusterId int

	for _, c := range centres {
		if px.distance(c) < distance {
			clusterId = c.cluster
		}
	}
	return clusterId
}

func imgToPixels(img image.Image) [][]Pixel {
	pixels := make([][]Pixel, img.Bounds().Max.Y)

	for i := 0; i < img.Bounds().Max.Y; i++ {
		pixelRow := make([]Pixel, img.Bounds().Max.X)
		for j := 0; j < img.Bounds().Max.X; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			pixelRow[j] = Pixel{int16(r), int16(g), int16(b), -1}
		}
		pixels[i] = pixelRow
	}

	return pixels
}

func createClusters(n int) []Pixel {
	if n < 1 {
		log.Fatal("Cannot create non positive number of clusters")
	}
	clusters := make([]Pixel, n)
	for i := 0; i < n; i++ {
		clusters[i-1] = Pixel{int16(rand.Intn(255)), int16(rand.Intn(255)), int16(rand.Intn(255)), i}
	}
	return clusters
}

func kmeans(n int, pixels [][]Pixel) []Pixel {
	converged := false
	clusters := createClusters(n)

	for !converged {
		clusterR := make([]int, n)
		clusterG := make([]int, n)
		clusterB := make([]int, n)
		clusterN := make([]int, n)
		// Assign all pixels to a cluster centre
		for i := 0; i < len(pixels); i++ {
			for j := 0; j < len(pixels[i]); j++ {
				pixels[i][j].cluster = pixels[i][j].nearestCluster(clusters)
				clusterR[pixels[i][j].cluster] += int(pixels[i][j].r)
				clusterG[pixels[i][j].cluster] += int(pixels[i][j].g)
				clusterB[pixels[i][j].cluster] += int(pixels[i][j].b)
				clusterN[pixels[i][j].cluster] += 1
			}
		}
		// Calculate the new centres and test for convergence

	}
	return clusters
}
