package auth

import (
	"gorm.io/gorm"
	"time"

	casbin "github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"log" // 添加日志库
)

// Authz 定义了一个授权器，提供授权功能.
type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz 创建一个使用 casbin 完成授权的授权器.
func NewAuthz(db *gorm.DB, autoLoadInterval time.Duration) (*Authz, error) {
	a, err := adapter.NewAdapterByDB(db)
	if err != nil {
		log.Printf("Failed to create adapter: %v", err)
		return nil, err
	}

	e, err := casbin.NewSyncedEnforcer("./rbac_model.conf", a)
	if err != nil {
		log.Printf("Failed to create enforcer: %v", err)
		return nil, err
	}

	e.StartAutoLoadPolicy(autoLoadInterval)
	log.Println("Policy loaded successfully")

	az := &Authz{e}
	return az, nil
}

// Authorize 用来进行授权.
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}
