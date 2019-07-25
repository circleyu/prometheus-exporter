/**
 *  Author: SongLee24
 *  Email: lisong.shine@qq.com
 *  Date: 2018-08-15
 *
 *
 *  prometheus.Desc是指标的描述符，用于实现对指标的管理
 *
 */

 package collector

 import (
	 "math/rand"
	 "sync"
 
	 "github.com/prometheus/client_golang/prometheus"
 )
 
 // Metrics 指標結構
 type Metrics struct {
	 metrics map[string]*prometheus.Desc
	 mutex   sync.Mutex
 }
 
 // newGlobalMetric 創建指標描述符
 func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	 return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
 }
 
 // NewMetrics 初始化指標訊息，即Metrics結構體
 func NewMetrics(namespace string) *Metrics {
	 return &Metrics{
		 metrics: map[string]*prometheus.Desc{
			 "mock_ccu_metric":    newGlobalMetric(namespace, "mock_ccu_metric", "The description of mock_ccu_metric", []string{"srevice"}),
			 "mock_battle_metric": newGlobalMetric(namespace, "mock_battle_metric", "The description of mock_battle_metric", []string{"srevice"}),
		 },
	 }
 }
 
 // Describe 傳遞結構中的指標描述符到channel
 func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	 for _, m := range c.metrics {
		 ch <- m
	 }
 }
 
 // Collect 抓取最新數據，傳遞給channel
 func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	 c.mutex.Lock() // 上鎖
	 defer c.mutex.Unlock()
 
	 mockCCUMetricData, mockBattleMetricData := c.GenerateMockData()
	 for host, currentValue := range mockCCUMetricData {
		 ch <- prometheus.MustNewConstMetric(c.metrics["mock_ccu_metric"], prometheus.GaugeValue, float64(currentValue), host)
	 }
	 for host, currentValue := range mockBattleMetricData {
		 ch <- prometheus.MustNewConstMetric(c.metrics["mock_battle_metric"], prometheus.CounterValue, float64(currentValue), host)
	 }
 }
 
 // GenerateMockData 生成模擬數據
 func (c *Metrics) GenerateMockData() (mockCCUMetricData map[string]int, mockBattleMetricData map[string]int) {
	 mockCCUMetricData = map[string]int{
		 "Login":  int(rand.Int31n(1000)),
		 "Lobby":  int(rand.Int31n(1000)),
		 "Battle": int(rand.Int31n(1000)),
	 }
	 mockBattleMetricData = map[string]int{
		 "Battle1": int(rand.Int31n(10)),
		 "Battle2": int(rand.Int31n(10)),
		 "Battle3": int(rand.Int31n(10)),
	 }
	 return
 }
 