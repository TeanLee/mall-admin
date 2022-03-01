<template>
  <div class="summary-chart">
    <div id="main" style="width: 100%;height: 400px;"></div>
    <div style="margin: 30px 0;"></div>
    <div id="lineStack" style="width: 100%;height: 400px;"></div>
  </div>
</template>
 
<script>
export default {
  name: 'Echarts',
  methods: {
    myEcharts(){
      var myChart = this.$echarts.init(document.getElementById('main'));
      //配置图表
      var option = {
        title: {
          text: '加购数量排名前十的商品',
        },
        tooltip: {},
        legend: {
          data: ['销量']
        },
        xAxis: {
          data: ['衬衫','羊毛衫','雪纺衫','裤子','高跟鞋','袜子']
        },
        yAxis: {},
        series: [{
          name: '销量',
          type: 'bar',
          data: [5,20,36,10,10,20]
        }]
      };
      myChart.setOption(option);

      var lineStack = this.$echarts.init(document.getElementById('lineStack'))
      var lineStackOption = {
        title: {
          text: '商品加购数量走势图'
        },
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          data: ['Email', 'Union Ads', 'Video Ads', 'Direct', 'Search Engine']
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        toolbox: {
          feature: {
            saveAsImage: {}
          }
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: 'Email',
            type: 'line',
            stack: 'Total',
            data: [120, 132, 101, 134, 90, 230, 210]
          },
          {
            name: 'Union Ads',
            type: 'line',
            stack: 'Total',
            data: [220, 182, 191, 234, 290, 330, 310]
          },
          {
            name: 'Video Ads',
            type: 'line',
            stack: 'Total',
            data: [150, 232, 201, 154, 190, 330, 410]
          },
          {
            name: 'Direct',
            type: 'line',
            stack: 'Total',
            data: [320, 332, 301, 334, 390, 330, 320]
          },
          {
            name: 'Search Engine',
            type: 'line',
            stack: 'Total',
            data: [820, 932, 901, 934, 1290, 1330, 1320]
          }
        ]
      }
      lineStack.setOption(lineStackOption)
    }
  },
  mounted(){
    this.myEcharts();
  }
}
</script>
 
<style>
.summary-chart {
  padding: 20px;
}
</style>