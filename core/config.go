package core

import "fmt"

// the units of below is pixels.
type Config struct {
	startX, startY float64 // PDF page start position
	endX, endY     float64 // PDF page end postion
	width, height  float64 // PDF page width and height

	contentWidth, contentHeight float64 // PDF page content width and height
}

// Params width, height is gopdf page width and height
// Params padingH, padingV is gopdf horizontal and vertical padding
// The units of the above parameters are pixels.
// Params width must more than 2*padingH, and height must more 2*padingV
func NewConfig(width, height float64, padingH, padingV float64) (*Config, error) {
	if width <= 0 || height <= 0 || padingH < 0 || padingV < 0 {
		return nil, fmt.Errorf("params must more than zero")
	}

	if width <= 2*padingH || height <= 2*padingV {
		return nil, fmt.Errorf("this config params invalid")
	}

	c := &Config{
		width:  width,
		height: height,

		startX: padingH,
		startY: padingV,

		contentWidth:  width - 2*padingH,
		contentHeight: height - 2*padingV,
	}

	c.endX = c.startX + c.contentWidth
	c.endY = c.startY + c.contentHeight

	return c, nil
}

func (config *Config) GetWidthAndHeight() (width, height float64) {
	return config.width, config.width
}

// Get gopdf page start position, from the position you can write the gopdf body content.
func (config *Config) GetStart() (x, y float64) {
	return config.startX, config.startY
}

func (config *Config) GetEnd() (x, y float64) {
	return config.endX, config.endY
}

var defaultConfigs map[string]*Config // page -> config

/*
*************************************
A0 ~ A5 page width and height config:

	'A0': [2383.94, 3370.39],
	'A1': [1683.78, 2383.94],
	'A2': [1190.55, 1683.78],
	'A3': [841.89, 1190.55],
	'A4': [595.28, 841.89],
	'A5': [419.53, 595.28],

**************************************
*/
func init() {
	defaultConfigs = make(map[string]*Config)

	defaultConfigs["A3"] = &Config{
		startX:        90.14,
		startY:        72.00,
		endX:          751.76,
		endY:          1118.55,
		width:         841.89,
		height:        1190.55,
		contentWidth:  661.62,
		contentHeight: 1046.55,
	}

	// todo 三联 宽城高  24.3*27.5  a4 21* 29.7
	//三联  w688 h780  a4  w595.28 h841.89

	defaultConfigs["A4"] = &Config{
		//startX: 90.14, //原值
		//startY: 72.00,
		//endX:   505.14, //原值
		//endY:          769.89, //画布大小pagesize
		//width:         595.28, //纸张
		//height:        841.89, //纸张
		//contentWidth:  415,
		//contentHeight: 697.89,

		//startX:        60.00,
		//startY:        50.00,
		//endX:          720,
		//endY:          514,
		//width:         780, //纸张
		//height:        595, //纸张
		//contentWidth:  670,
		//contentHeight: 480,
		//
		//startX:        60.00,
		//startY:        60.00,
		//endX:          628,
		//endY:          710,
		//width:         688, //纸张
		//height:        780, //纸张
		//contentWidth:  568,
		//contentHeight: 660,

		width:         688,   //纸张
		height:        780,   //纸张
		startX:        40.00, //起始位置
		startY:        40.00,
		endX:          640, //左40 右48
		endY:          720,
		contentWidth:  600,
		contentHeight: 680,
	}

	//三联 27*21  765 *595
	defaultConfigs["Triplet"] = &Config{
		startX:        80.14,
		startY:        72.00,
		endX:          750,
		endY:          550,
		width:         765, //纸张
		height:        595, //纸张
		contentWidth:  580,
		contentHeight: 450,
	}

	defaultConfigs["LTR"] = &Config{
		startX:        90.14,
		startY:        72.00,
		endX:          521.86,
		endY:          720,
		width:         612,
		height:        792,
		contentWidth:  431.72,
		contentHeight: 648,
	}
}

// Register create self gopdf config
func Register(size string, config *Config) {
	if _, ok := defaultConfigs[size]; ok {
		panic("config size has exist")
	}

	defaultConfigs[size] = config
}
