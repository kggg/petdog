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
    petdog new templatename:projectname  使用自定义的模块构建项目结构, 在冒号前面区分templatename在app.ini的[]中定义，类似[base]


### make 生成一些模板文件
    petdog make  tempname:directoryname filename

    1, 设置basedir的方式需要改进


### 模板说明
    模板的定义在conf目录下以[templatename].ini格式命名的文件, 模板文件中可以定义需要在项目目录中设置生成哪些目录，和哪些文件。

    如base.ini中：
       dirs = config, controller, models,routes,views,static  //新建项目时需要生成的子目录
       files = main.go    //新建项目时需要生成的文件
       templates = main.go.tpl  //设置生成文件时，在文件中要插入的内容， 对应的文件是上面files中列的文件名加tpl结尾， 模板文件存放在template目录下面对应的templatename目录里面， templatename下面的文件和内容对应新生成项目目录和文件


### run 运行项目
