package healthcheck

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/jinzhu/gorm"
)

func Serve(db *gorm.DB) {
	if os.Getenv("ONE_NODE") == "true" {
		return
	}

	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(300))
	// health.AddReadinessCheck("upstream-dep-dns", healthcheck.DNSResolveCheck("ya.ru", 1*time.Second))

	if db != nil {
		health.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db.DB(), 1*time.Second))
	}

	go log.Printf("ListenAndServe:%v", http.ListenAndServe("0.0.0.0:8086", health))
}
