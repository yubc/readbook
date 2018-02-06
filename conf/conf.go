package conf

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"readbook/models"

	"gopkg.in/yaml.v2"
)

var (
	fpath string
)

func init() {
	flag.StringVar(&fpath, "c", "./conf/config_local.yaml", "config file path")
	flag.Parse()
	if strings.TrimSpace(fpath) == "" {
		panic("config file path is null")
	}
	f, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(fmt.Sprintf("config file path error %s", err.Error()))
	}

	err = yaml.Unmarshal(f, &models.Conf)
	if err != nil {
		panic(fmt.Sprintf("config file nmarshal error %s", err.Error()))
	}

	models.Conf.WxTemplate = make([]models.Template, len(models.Conf.WxPushTemplate))
	for i := 0; i < len(models.Conf.WxPushTemplate); i++ {
		v := strings.SplitN(models.Conf.WxPushTemplate[i], ",", 3)
		if len(v) != 3 {
			panic("models.Conf file template length error")
		}
		models.Conf.WxTemplate[i].KeyType = v[0]
		models.Conf.WxTemplate[i].KeyId = v[1]
		models.Conf.WxTemplate[i].KeyURL = v[2]
	}
	// fmt.Println(fmt.Sprintf("%+v", models.Conf))
}
