#gopdf
功能点：
1.支持中文
2.支持自动换行
3.支持分页，并计算页码
4.支持页眉、页脚
5.支持无边框table，可用作等分
6.支持设置字体设置，包括大小，颜色、加粗等设置
7.页面最后一行因为行合并跨页可以放到下一页
8.支持纸张A3、A4、三联等，可以定制页面尺寸
9.支持二维码、条形码生成，可以通过div浮动到指定位置
10.可以通过合并单元格来达到设置列宽的目的

后端返回文件流
c.Writer.Header().Set("Content-type", "application/pdf")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Disposition")
	filename := url.QueryEscape(fmt.Sprintf("出库单%v.pdf", req.ID))
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", filename))
	c.String(http.StatusOK, string(res))


example  div + table

![image](https://github.com/MonkeyCodeHoward/gopdf/assets/22990697/6583b857-439d-4e50-b261-a337a4fbeb80)


