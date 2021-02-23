package Common

import (
	"net/http"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var accessKey = "mOpOT-dsE6wvyV3wHNdP7mLomqHJPLuj-NxitEdu"
var secretKey = "HkUy43SMlZsE4ImFPIAgEps9DXPoujxjV4l2vKtP"

func QiNiuToken(bucket string) string {
	if bucket == "" {
		bucket = "xc-year"
	}

	http.Get()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	putPolicy.Expires = 86400 //示例24小时有效期单位秒

	mac := qbox.NewMac(accessKey, secretKey)

	return putPolicy.UploadToken(mac)
}
