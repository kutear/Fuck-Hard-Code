[![Build Status](https://travis-ci.org/Kutear/Fuck-Hard-Code.svg?branch=master)](https://travis-ci.org/Kutear/Fuck-Hard-Code)

# 中文文档
[中文](./README-ZH.md)

# Android hardcoding extraction tool

Extraction string and dimens from layout,Merge into strings.xml and dimens.xml

[Download](https://gobuilder.me/github.com/Kutear/Fuck-Hard-Code)

# Usage

![Example](http://pic.kutear.com/2017/01/22/b728daeb8969e99f3502ccf246f104fd.png)

> **Programs description**
>
> **Programs purpose**:  fixed Android project Hard Code Problems
> **Args**:  
>
> - layout:the path of $project/appmodule/src/main/res/layout
>  - config: configure file
>  - existPixels:the path of $project/appmodule/src/main/res/values/dimens.xml
>  - existStrings:the path of $project/appmodule/src/main/res/values/strings.xml
>  - scaleRatio:px to dp,default 3px == 1dp
>  - out:layout dir out path

after that you will see strings.xml and dimens.xml in `{layout output path}/out` and all xml file in `{layout output path}`

# Example

Create New Android Application and modify main layout

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

Now We Run This Tool,Over this,We can see the file in {out path/layout} with
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

And {out path/layout/out} has two file ,strings.xml and dimens.xml
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

Before Copy New file to Replace original File,Please make sure everything is OK!!!!



