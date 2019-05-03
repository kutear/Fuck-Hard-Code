[![Build Status](https://travis-ci.org/Kutear/Fuck-Hard-Code.svg?branch=master)](https://travis-ci.org/Kutear/Fuck-Hard-Code)

# Document for English
[English](./README.md)

# Android 硬编码提取工具

提取Android Layout硬编码的string和dimens出来

[点击下载|Download](https://gobuilder.me/github.com/Kutear/fuck-hard-code)

# 用法
![Example](http://pic.kutear.com/2017/01/22/b728daeb8969e99f3502ccf246f104fd.png)

>
> **程序说明**
> **目的**:  解决Android项目中遗留的硬编码问题
> **参数**:  
> - layout:Android 项目中的layout目录
> - config: json文件,内部表示具体哪些字段需要被替换
> - existPixels:通常为/values/dimens.xml
> - existStrings:通常为/values/strings.xml
> - scaleRatio:dp转化px的比例值 默认为3
> - out:输出修改后layout目录,注意不要与输入layout的一样


执行之后你可以在这个目录`{layout output path}/out/`看见两个文件`strings.xml` 和 `dimens.xml`以及其他所有的layout文件在`{layout output path}/`

# 实例

找到自己App的Layout目录(或新建Application),
下面展示其中一个文件

```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout
	xmlns:android="http://schemas.android.com/apk/res/android"
	android:layout_width="match_parent"
	android:layout_height="wrap_content"
	android:orientation="horizontal">
	<TextView
		android:id="@+id/content"
		android:layout_width="wrap_content"
		android:layout_height="40dp"
		android:text="Hello Word"
		android:textSize="20sp" />
</LinearLayout>
```

运行工具

```
> Fuck-Hard-Code -input="{project path/app/src/main/res/layout}" -output="{out path/layout}"
```

之后我们可以在`{out path/layout}`下看到对应文件的内容被修改为

```xml
<?xml version="1.0" encoding="utf-8"?>

<LinearLayout
	xmlns:android="http://schemas.android.com/apk/res/android"
	android:layout_width="match_parent"
	android:layout_height="wrap_content"
	android:orientation="horizontal"
	android:gravity="top">
	<TextView
		android:id="@+id/content"
		android:layout_width="wrap_content"
		android:layout_height="@dimen/dp_0040_0"
		android:text="@string/strings_0"
		android:textSize="@dimen/sp_0020_0" />
</LinearLayout>
```

`{out path/layout/out}` 中有两个文件,strings.xml 和 dimens.xml,内容分别为

```xml
<!--strings.xml-->
<?xml version="1.0" encoding="utf-8"?>
<resources>
	<string	name="strings_0">Hello Word</string>
</resources>
```

```xml
<!--dimens.xml-->
<?xml version="1.0" encoding="utf-8"?>

<resources>
	<dimen	name="dp_0040_0">40.0dp</dimen>
	<dimen	name="sp_0020_0">20.0sp</dimen>
</resources>
```

这样就替换了布局中所有的硬编码格式。可以使用文件比较工具查看之后，没有问题就替换掉以前的布局代码

