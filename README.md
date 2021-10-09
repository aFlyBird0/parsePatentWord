# 专利审查指南解析
本项目代码地址：[aFlyBird0/parsePatentWord: 利用 Golang 解析《专利审查指南 2020...doc》，并将标题、段落层级关系存至数据库](https://github.com/aFlyBird0/parsePatentWord)
## 〇、本项目的由来、意义
### 由来
导师布置了个任务，提取《专利审查指南...》，做成电子书。所以第一步是把标题层级和段落内容提取到数据库里面。  
师兄说建议用 Java 的 `Apache POI`，但我刚学了 Golang 一星期，就想用 Golang 做！！！死磕~~~（其实中间反复过两次，真的很想用 Java 或 Python 的 API！还好坚持下来了，收获了很多）
### 意义
1. 第一个真正意义上的自己的 Go 项目，检验、巩固了前几天的学习
2. 第一次这么认真看源码，因为没有文档？？！极大锻炼了自己看源码的能力，还挺有趣？
3. 和室友拥有了一次愉快的半夜一起写代码的奇妙而又愉快的体验。他看我在弄，来了兴趣，现学 C# 和我并行尝试。中间不少思路我们都想到一块了，比如发现文档没有标题格式全是正文，思考大纲转标题，然后又都决定找框架的「大纲」接口
4. 提前一星期完成了老师的任务，后面一星期可以专心学 Go 了。
5. 利用这次项目，对 vim 比以前更习惯了一点，总算「存活」下来了。

## 一、本项目内容
项目大背景：把《专利审查指南...》做成网页版的电子书，加入高级搜索等功能。
本项目工作：提取《专利审查...》内容，抽取标题层级和段落内容，存入数据库中。
项目进度：
* 2021年9月11日 3点52分 目前已完成标题层级与文档内容抽取为 `Golang Slice` 功能。
* 2021年9月13日：标题与段落内容加上层级关系；数据入库
* 2021年9月15日：完成了从数据库中查询目录，生成树结构传给前端，以及查询文章内容。（在另一个项目中）

## 二、抽取过程概述
### 2.1 按段落与大纲初步抽取
阅读 `gooxml` 源码，按段落抽取 Word 文档，每段带有大纲级别。该步骤得到的内容格式如下：

![image-20210915223953307](https://bird-notes.oss-cn-hangzhou.aliyuncs.com/img/image-20210915223953307.png)

其中，`4` 代表4级目录，`0` 代表文本内容

### 2.2 合并近邻的同大纲内容
同一目录下可能有多个段落，但向前端返回的文本，应该是一整段，所以合并紧邻的同大纲内容。

### 2.3 将层级目录与文本内容存至数据库中
** 这步为本项目的难点 **  
源码见 `/build/build.go`
上一步获得的内容是这样的：
```yml
1   这是一级标题 a
1   这是一级标题 b
2 这是一级标题 b 下的二级标题 c
3 这是二级标题 c 下的三级标题 d
0 这是三级标题下的文本 e
2 这是一级标题 a 下的二级标题 f
...
```
难点在于：**既要在数据库中指明段落文本属于哪个标题（目录），又要指明每个标题的父标题是什么**  
这里用到了 **栈**，来保存最近的一个标题在数据库中的 `id` 以及 当前的最深的大纲级别 `maxLvl`（代码中是 `lastLvl`）
逐个遍历第二步得到的带有大纲级别的段落：
* 如果当前内容是文本（大纲级别是0），之前链接此文本内容到栈顶的标题。因为是文本内容一定附属在最近的标题下。
* 如果是当前内容是标题（大纲级别大于0）
    * 如果当前内容的大纲级别大于 `maxLvl` ，就说明目录变深了，压栈，`maxLvl` + 1，并将当前标题的父标题指向原来栈顶的 `id`
    * 如果相同，出栈，保证栈内不存在同级大纲。这样出栈后，当前栈顶元素一定是当前内容的父标题。 `maxLvl`不变，标题父目录指向出栈后的 `id`。
    * 如果小于。出栈多次。因为变深只能一次深一级，如从二级目录变成三级目录，而变浅可以直接从三级目录变成一级目录。计算变浅的差值，决定 `pop` 的次数，以及 `maxLvl` 减少的层级即可。

## 三、开发问题与解决方案记录
### 1.文件相对绝对路径
Go 的环境路径和项目路径不一致，所以要获取代码文件所在的路径，自己写了个工具类，即 `/util/fileUtil.go` 的 `GetRunPath()`
### 2. gooxml document 不支持 doc
doc 转存成 docx

### 3. gooxml 和 unidoc/unioffice 弃用
gooxml 太旧了，文档也 404  
unioffice 应该是基于 gooxml 写的，但是要授权，只能免费用 14 天  
所以都弃用，但是真就没办法了吗？

### 4. 切换解析 word 方法
受某大佬启发，先将 word 转成 html 的形式，再读取 html，原因如下：  
我要实现的目标是读取 word，所以可以转换成等效格式，再提取。如果是编辑 word，那可能要直接操作 word 了  
之所以选择 html ，是因为我之前对爬虫比较熟，用 xpath 有信心。分析 html 源码后，确定能获取各元素属性

### 4. 如何解析标题
`mso-outline-level` 属性，揭示了大纲级别，`1` 为「第一部分 初步审查」，以此类推

### 5. 如何解析文本（段落）
转化成的 html ，一段为一个 `div`， 同一标题下的多段没有共同的 `div`，所以只能一段一段提取。  
目前拟采用 `mso-char-indent-count:2.0`，因为每段前都有缩进。

### 6. 编码与文件转换问题
先把 doc 转 docx，再另存为 html，再用 vscode 打开（编码 GBK），然后保存成 UTF-8

### 7. htmlQuery 获取 html node 的字符串
查看源代码，引用 `Data` 属性就行  
注：因为 Word 转 html 后有很多格式错乱(例如左侧的「细则」会整段混入到右侧的文本中），导致虽然能识别大部分文档，但是会有很多乱七八糟的东西，所以还是决定再试试 gooxml

### 8. 利用 gooxml/document 设置 word 样式
[gooxml/main.go at master · carmel/gooxml](https://github.com/carmel/gooxml/blob/master/_examples/document/simple/main.go)
利用 para.SetStyle() 属性

### 9. 如何利用 gooxml/document 获取段落属性
这里没有文档，是个难点，只能肝源码  
先看看 `(p Paragraph) SetStyle(s string)` 的函数定义：
```Golang
func (p Paragraph) SetStyle(s string) {
    p.ensurePPr()
    if s == "" {
        p.x.PPr.PStyle = nil
    } else {
        p.x.PPr.PStyle = wml.NewCT_String()
        p.x.PPr.PStyle.ValAttr = s
    }
}
```
可以得知，关键样式属性在于 `Paragraph.p.x.PPr.PStyle.ValAttr`, 但这是私有属性。  
继续看 `Paragraph` 所在的 `paragraph.go` 文件，找到了 `(p Paragraph) Style() string` 函数
```Golang

func (p Paragraph) Style() string {
    if p.x.PPr != nil && p.x.PPr.PStyle != nil {
        return p.x.PPr.PStyle.ValAttr
    }
    return ""
}
```
Bingo！

### 10. 《审查指南...》中，没有「标题」样式，全是正文
正常的格式应该是 `Heading1` 这种（经由 `/try/parseTry.go` 中的 `TestNewDocGooxml` 测试，确保正确）  
但是，获取了同一个 `style`（即 `x.PPr.PStyle.ValAttr`）属性，发现：  
** 《审查指南..》中的样式，要么是空，要么是类似 `1`、`2`、`40`、`50` 这样的数字  
并且，这些数字和大纲级别没有严格对应之处

### 11. 尝试用属性获取
`paragraph.go` 源码文件，有个 `Properties` 函数，或许有用
```Golang 
func (p Paragraph) Properties() ParagraphProperties {
	p.ensurePPr()
	return ParagraphProperties{p.d, p.x.PPr}
}
```
进入 `PragraphProperties.go`，得到：
```Golang

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	d *Document
	x *wml.CT_PPr
}

...
...

// Style returns the style for a paragraph, or an empty string if it is unset.
func (p ParagraphProperties) Style() string {
    if p.x.PStyle != nil {
        return p.x.PStyle.ValAttr
    }
    return ""
}
```
经过测试，发现这里的 `style` 和前面那个 `style` 指向的是同一个东西，线索中断。

### 12. 能否通过将大纲转成标题来实现提取？
《审查指南...》虽然没有标题级别，但是有严格的大纲级别，能否通过外部手段，不通过 `Golang`，先把大纲转成 `style` 的 `Heading`？

### 13. 找到了 gooxml 中的 「大纲属性」
先去看 `ParagraphProperties.go` 文件，找到类定义：
```Golang
// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	d *Document
	x *wml.CT_PPr
}
```
然后看看 `wml.CT_PPr` 到底是怎么定义的？
`CT_PPr.go` 中定义了该结构体，其中有一个属性是这个：
```Golang
// Associated Outline Level
type CT_PPr struct {
	...
	OutlineLvl *CT_DecimalNumber
	...
}
```
再看看 `CT_DecimalNumber` 是什么：
```Golang
type CT_DecimalNumber struct {
	// Decimal Number Value
	ValAttr int64
}
```
其实就是个结构体指针，里面有个 `int64`   
那么怎么获取这个属性(`x *wml.CT_PPr`)呢？回到 `ParagraphProperties.go` 中：
```Golang
// X returns the inner wrapped XML type.
func (p ParagraphProperties) X() *wml.CT_PPr {
    return p.x
}
```
总结一下，获取 Word 大纲级别的方式如下，这里要防止空指针：
```Golang
for _, para := range doc.Paragraphs() {
    var outlineLvl int64
    if outlineLvlStruct := para.Properties().X().OutlineLvl; outlineLvlStruct!=nil{
        outlineLvl = outlineLvlStruct.ValAttr
    }else {
        outlineLvl = 0
    }
}
```



## 文档
* [unidoc/unioffice 另外一个可以处理 office 的文档](https://github.com/unidoc/unioffice)
* [gooxml 官方文档](https://pkg.go.dev/github.com/baliance/gooxml#section-readme) 这个文档应该是 `unioffice` 的前身
