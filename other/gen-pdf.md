- 需求：本地存在大量PDF文件，需要在所有PDF上附上水印，用于云存储上传。

- 分析：经过对市面上程序包考察，最终选定unipdf程序包，选择的原因，开源，操作简单，速度够快。

> 添加水印示例：https://github.com/unidoc/unipdf-examples/blob/master/image/pdf_watermark_image.go

- 遇到问题：使用程序包处理PDF时需要获取证书及生成新的PDF时会出现程序包logo。

- 分析问题：首先明确一点，程序包是开源的，源码可以拉取下来，根据经验来猜测一下程序的处理PDF步骤：1. 获取服务器证书校验。 2.处理PDF。 3. 添加水印。

- 处理问题：首先拉取源码，然后在本地debug模式下运行代码，一步步走下去，跟分析时猜测一般无二，接下来就很简单了，只需要把校验证书及添加程序包自身logo位置的代码注释了就可以了。

<img width="972" alt="image" src="https://user-images.githubusercontent.com/82997695/212447654-705bfc45-b927-4e1f-95d2-685a46d306a3.png">
<img width="957" alt="image" src="https://user-images.githubusercontent.com/82997695/212447712-d5dc178c-08a4-4391-96fc-49cb43ee86a4.png">



- 复盘：处理问题需要先理出逻辑，理清思绪，这样在解决问题时如有神助。
- 备注：建议通过 https://cloud.unidoc.io 获取证书，仅个人使用，如有侵权，联系即删。