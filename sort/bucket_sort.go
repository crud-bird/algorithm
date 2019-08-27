package sort

import (
	"fmt"
	"math/rand"
	"time"
)

type buckets struct {
	Buckets map[int][]int64
	m       int
}
type bucket struct {
	Bucket []int64
}

//BucketSort 桶排序,桶内使用快排时，时间复杂度O(n)+m*(n/m)*log(n/m)
func BucketSort(a []int64, m int) []int64 {
	bs := initBucket(m)
	//分桶
	capa := len(a)/m + 1
	for _, v := range a {
		bucketNO := len(a)/capa + 1
		bs.Add(bucketNO, v)
	}

	return bs.Sort()
}

func initBucket(m int) *buckets {
	bs := make(map[int][]int64)
	for i := 0; i < m; i++ {
		bs[i+1] = []int64{}
	}
	return &buckets{
		Buckets: bs,
		m:       m,
	}
}

func (bs *buckets) Add(key int, v int64) {
	tmp := bs.Buckets[key]
	bs.Buckets[key] = append(tmp, v)
}

func (bs *buckets) Sort() []int64 {
	res := make([]int64, 0)
	for i := 0; i < bs.m; i++ {
		QuickSort(bs.Buckets[i+1], 0, len(bs.Buckets[i+1])-1)
		res = append(res, bs.Buckets[i+1]...)
	}
	return res
}

//TestBucketSort ..
func TestBucketSort() {
	rand.Seed(time.Now().Unix())
	a := make([]int64, 100)
	for i := 0; i < 100; i++ {
		a[i] = rand.Int63n(10000)
	}
	b := BucketSort(a, 10)
	fmt.Println(a)
	fmt.Println(b)
}
