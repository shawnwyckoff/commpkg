package htmls

import (
	"github.com/shawnwyckoff/gpkg/dsa/decimals"
	"github.com/shawnwyckoff/gpkg/dsa/num"
	"github.com/shawnwyckoff/gpkg/encoding/color"
	"github.com/shawnwyckoff/gpkg/sys/gclock"
)

type (
	// 一个点+文字注释
	Tip struct {
		X     string      // X坐标，即时间
		Y     float64     // Y坐标
		Text  string      // 内容
		Color color.Color // 样式
	}

	// 矩形色块
	Rect struct {
		X1    string           // 矩形第一个点的X坐标，即时间
		Y1    decimals.Decimal // 矩形第一个点的Y坐标
		X2    string           // 矩形第二个点的X坐标，即时间
		Y2    decimals.Decimal // 矩形第二个点的Y坐标
		Color color.Color      // 样式，默认半透明的
	}

	// 直线
	StraightLine struct {
		X1    string           // 直线第一个点的X坐标，即时间
		Y1    decimals.Decimal // 直线第一个点的Y坐标
		X2    string           // 直线第二个点的X坐标
		Y2    decimals.Decimal // 直线第二个点的Y坐标
		Color color.Color      // 样式
	}

	// 蜡烛图
	CandleStick struct {
		Name          string                // 值名称，光标放在绘图区时在绘图区上沿要显示值名称和值
		Ohlc          [4][]num.ElegantFloat // Open,High,Low,Close
		DownColor     color.Color           // 阴线样式，支持背景颜色
		UpColor       color.Color           // 阳线样式，支持背景颜色
		Tips          []Tip                 // 文字注释
		Rects         []Rect                // 矩形色块
		StraightLines []StraightLine        // 直线，比如压力位、之字形折线、趋势线
	}

	// 柱状图
	Bar struct {
		Name   string             // 值名称，光标放在绘图区时在绘图区上沿要显示值名称和值
		Data   []num.ElegantFloat // 数据
		Colors []color.Color      // 每根bar的颜色
	}

	// 散点连线
	Line struct {
		Name  string              // 值名称，光标放在绘图区时在绘图区上沿要显示值名称和值
		Data  []*num.ElegantFloat // 数据，空指针表示不画点，在eCharts输入数据中就是null
		Color color.Color         // 样式，支持背景颜色
	}

	// 一个绘图区域，CandleStick和Bar不可以同时绘制，Lines可以有多个
	Series struct {
		Name        string       // 绘图区左上角显示的名字，比如"601398（工商银行）"
		CandleStick *CandleStick // 时间轴线条类型1：蜡烛图, 每个Series最多一个蜡烛图
		BarStick    *Bar         // 时间轴线条类型2：柱状图, 每个Series最多一个柱状图
		Lines       []Line       // 时间轴线条类型3：曲线图或者折线图
	}

	// 一个绘图模板，包含了整个页面需要的数据，包括一个共享的时间轴和多个绘图区域，多种数据显示样式
	ChartTemplate struct {
		Title   string               // 整个图表的大标题
		Times   []gclock.ElegantTime // 时间轴
		Series  []Series             // 所有子图表
		Heights []int                // 比如有4个Series，那么就用[60,10,10,20]表示每个Series的显示百分比
	}
)

// TODO: implement
// 要求光标同步
func ServeChart(listen string, tpl ChartTemplate) {

}
