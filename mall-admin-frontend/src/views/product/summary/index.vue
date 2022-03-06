<template>
  <div class="summary-chart">
    <ve-histogram :data="chartData" :extend="chartExtend"> </ve-histogram>
    <!-- <div id="main" style="width: 100%;height: 400px;"></div>
    <div style="margin: 30px 0;"></div>
    <div id="lineStack" style="width: 100%;height: 400px;"></div> -->
  </div>
</template>

<script>
import ProductService from "@/service/product.service.js"

export default {
  name: "Echarts",
  data() {
    this.chartExtend = {
      // 柱子宽度
      series: {
        barWidth: 20,
      },
      // x轴字体斜显示
      xAxis: {
        axisLabel: {
          margin: 15,
          interval: 0,
          rotate: 30,
        },
        triggerEvent: true,
      },
    };
    return {
      extend: {
        // x轴的文字倾斜
        "xAxis.0.axisLabel.rotate": 45,
        yAxis: {
          //是否显示y轴线条
          axisLine: {
            show: false,
          },
          // 纵坐标网格线设置，同理横坐标
          splitLine: {
            show: true,
          },
          // 线条位置
          position: "left",
        },
        xAxis: {
          axisLine: {
            show: true,
          },
          axisLabel: {
            margin: 15,
            interval: 0,
            rotate: 30,
            formatter: name => {
              // eslint-disable-next-line
              return echarts.format.truncateText(name, 100, '14px Microsoft Yahei', '…')
            }
          },
          triggerEvent: true
        },
        series: {
          type: "bar",
          label: { show: true, position: "top" },
          barMinHeight: 400
        },
        //数据过多时出现横向滚动条
        // dataZoom: [
        //   {
        //     show: true,
        //     realtime: true,
        //     start: 0,
        //     end: 50,
        //   },
        //   {
        //     type: "inside",
        //     realtime: true,
        //     start: 0,
        //     end: 50,
        //   },
        // ],
      },
      chartData: {
        columns: ["name", "count"],
        rows: [],
      },
      chartSet: {
        labelMap: {
          name: "访问用户",
        },
        max: [1],
      },
    };
  },
  methods: {
    myEcharts() {
      var myChart = this.$echarts.init(document.getElementById("main"));
      //配置图表
      var option = {
        title: {
          text: "加购数量排名前十的商品",
        },
        tooltip: {},
        legend: {
          data: ["销量"],
        },
        xAxis: {
          data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"],
        },
        yAxis: {},
        series: [
          {
            name: "销量",
            type: "bar",
            data: [5, 20, 36, 10, 10, 20],
          },
        ],
      };
      myChart.setOption(option);

      var lineStack = this.$echarts.init(document.getElementById("lineStack"));
      var lineStackOption = {
        title: {
          text: "商品加购数量走势图",
        },
        tooltip: {
          trigger: "axis",
        },
        legend: {
          data: ["Email", "Union Ads", "Video Ads", "Direct", "Search Engine"],
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        toolbox: {
          feature: {
            saveAsImage: {},
          },
        },
        xAxis: {
          type: "category",
          boundaryGap: false,
          data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            name: "Email",
            type: "line",
            stack: "Total",
            data: [120, 132, 101, 134, 90, 230, 210],
          },
          {
            name: "Union Ads",
            type: "line",
            stack: "Total",
            data: [220, 182, 191, 234, 290, 330, 310],
          },
          {
            name: "Video Ads",
            type: "line",
            stack: "Total",
            data: [150, 232, 201, 154, 190, 330, 410],
          },
          {
            name: "Direct",
            type: "line",
            stack: "Total",
            data: [320, 332, 301, 334, 390, 330, 320],
          },
          {
            name: "Search Engine",
            type: "line",
            stack: "Total",
            data: [820, 932, 901, 934, 1290, 1330, 1320],
          },
        ],
      };
      lineStack.setOption(lineStackOption);
    },
  },
  mounted() {
    // this.myEcharts();
    ProductService.productMostAdd().then(res => {
      const { data } = res
      this.chartData.rows = data
      this.chartData.rows.forEach(item => item.count = parseInt(item.count))
    })
  },
  created() {},
};
</script>

<style>
.summary-chart {
  padding: 20px;
}
</style>
