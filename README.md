petdog 是方便新建go web项目的一个小工具, 尽量去掉一些重复的工作

###  new  新建一个GO项目
    petdog new projectname 
    
    默认生成conf/app.ini下[base]格式的目录结构如下:
    |projectname
    |----config
    |----controller
    |----models
    |----routes
    |----views
    |----static
    |----go.mod
    |----main.go

    以go modules 为软件包依赖管理模块, 
    项目下生成的目录可以在conf下app.ini文件中自定义
    petdog new modelname:projectname  使用自定义的模块构建项目结构, 在冒号前面区分modelname在app.ini的[]中定义，类似[base]


### make 生成一些模板文件
    petdog make  [ controller|models ] filename
