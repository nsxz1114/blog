package flags

import (
	"github.com/nsxz1114/blog/global"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

//type Option struct {
//	DB     bool   // 建表
//	User   string // -u admin  -u user
//	Load   string // 导入数据库文件
//	Dump   bool   // 导出数据库
//	Es     bool   // 创建索引
//	ESDump bool   // 导出es索引
//	ESLoad string // 导入es索引
//}
//
//// Parse 解析命令行参数
//func Parse() (option *Option) {
//	option = new(Option)
//	flag.StringVar(&option.User, "u", "", "创建用户")
//	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
//	flag.BoolVar(&option.Es, "es", false, "创建索引")
//	flag.BoolVar(&option.Dump, "dump", false, "导出sql数据库")
//	flag.StringVar(&option.Load, "load", "", "导入sql数据库")
//	flag.BoolVar(&option.ESDump, "esdump", false, "导出es索引")
//	flag.StringVar(&option.ESLoad, "esload", "", "导入es索引")
//	flag.Parse()
//	return option
//}
//
//// Run 根据命令执行不同的函数
//func (option Option) Run() bool {
//	if option.DB {
//		DB()
//		return true
//	}
//	return false
//}

func Newflags() {
	var app = cli.NewApp()
	app.Name = "github.com/nsxz1114/blog"
	app.Usage = "github.com/nsxz1114/blog"
	app.Authors = []*cli.Author{
		{
			Name:  "Axios",
			Email: "1790146932@qq.com",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "create table",
			Action:  DB,
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "create a user",
		},
		{
			Name:    "admin",
			Aliases: []string{"a"},
			Usage:   "create a admin",
		},
		{
			Name:    "export-mysql",
			Aliases: []string{"e-m"},
			Usage:   "export mysql data",
			Action:  MysqlExport,
		},
		{
			Name:    "import-mysql",
			Aliases: []string{"i-m"},
			Usage:   "import mysql data",
			Action:  MysqlImport,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "path",
				},
			},
		},
		{
			Name:    "elasticsearch",
			Aliases: []string{"es"},
			Usage:   "create a elasticsearch index",
			Action:  EsIndexCreate,
		},
		{
			Name:    "export-es",
			Aliases: []string{"e-e"},
			Usage:   "export elasticsearch data",
			Action:  EsExport,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "index",
				},
			}},
		{
			Name:    "import-es",
			Aliases: []string{"i-e"},
			Usage:   "import elasticsearch data",
			Action:  EsImport,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "path",
				},
			},
		},
	}
	if len(os.Args) > 1 {
		err := app.Run(os.Args)
		if err != nil {
			global.Log.Error("init cmd error", zap.Error(err))
			return
		}
		os.Exit(0)
	}
}
