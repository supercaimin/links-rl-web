package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"
	"github.com/qor/links-rl-web/config"
	"github.com/qor/links-rl-web/models"
	"github.com/qor/media/media_library"
	"github.com/qor/session/manager"
	"github.com/qor/transition"
)

var (
	DB, _ = gorm.Open("mysql", "root:root@tcp(localhost:3306)/qor_example?charset=utf8")
	Admin = admin.New(&admin.AdminConfig{DB: DB, Auth: config.AdminAuth})
	Auth  = auth.New(&auth.Config{DB: DB})
)

func main() {
	//DB.AutoMigrate()
	DB.AutoMigrate(&models.User)
	DB.AutoMigrate(&transition.StateChangeLog{})
	DB.AutoMigrate(&auth_identity.AuthIdentity{}, &auth_identity.AuthIdentity{}) //自动创建表。
	Auth.RegisterProvider(password.New(&password.Config{}))
	Admin.AddResource(&auth_identity.AuthIdentity{})

	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	var I18n *i18n.I18n

	I18n = i18n.New(database.New(DB), yaml.New(filepath.Join(os.Getenv("GOPATH")+"/src/github.com/qor/links-rl-web", "config/locales")))
	I18n.AddTranslation(&i18n.Translation{Key: "hello-world", Locale: "zh-CN", Value: "你好啊"})

	// Add Translations
	Admin.AddResource(I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: -1})
	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	mux.Handle("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9090", manager.SessionManager.Middleware(mux))

}
