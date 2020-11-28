module 01first_go

go 1.15

require "test" v0.0.0
replace "test" => "../02test_go" //“=>”需要空格 而且test的指明路径,
//最后只需要写到文件夹即可，应该会搜索该文件夹的所有方法
//"../02test_go" 其中test是package名，可见最不重要的就是文件名
//指向说明：文件夹路径，包名，方法名，其中文件名不重要,很有可能将文件名替换为包名
