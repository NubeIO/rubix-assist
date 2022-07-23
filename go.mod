module github.com/NubeIO/rubix-assist

go 1.17

//replace github.com/NubeIO/edge => /home/aidan/code/go/nube/edge
//replace github.com/NubeIO/rubix-automater => /home/aidan/code/go/nube/rubix-automater
//replace github.com/NubeIO/lib-schema => /home/aidan/code/go/nube/lib/lib-schema
replace github.com/NubeIO/lib-rubix-installer => /home/aidan/code/go/nube/lib/lib-rubix-installer

require (
	github.com/NubeIO/nubeio-rubix-lib-helpers-go v0.2.7
	github.com/NubeIO/nubeio-rubix-lib-rest-go v1.0.8
	github.com/appleboy/gin-jwt/v2 v2.8.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.8.0
	github.com/jinzhu/copier v0.3.4
	github.com/melbahja/goph v1.3.0
	github.com/pkg/sftp v1.13.4 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.11.0
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	gorm.io/driver/sqlite v1.3.2
	gorm.io/gorm v1.23.5
)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect; indirect``
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0
	github.com/golang-jwt/jwt/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/helloyi/go-sshclient v1.1.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/NubeIO/edge v0.0.8
	github.com/NubeIO/flow-framework v0.5.6
	github.com/NubeIO/lib-bus v0.0.1
	github.com/NubeIO/lib-command v0.0.2
	github.com/NubeIO/lib-date v0.0.1
	github.com/NubeIO/lib-dirs v0.0.2
	github.com/NubeIO/lib-networking v0.0.4
	github.com/NubeIO/lib-redis v0.0.3
	github.com/NubeIO/lib-rubix-installer v0.0.1
	github.com/NubeIO/lib-schema v0.0.3
	github.com/NubeIO/lib-uuid v0.0.2
	github.com/NubeIO/nubeio-rubix-lib-models-go v1.3.0
	github.com/NubeIO/rubix-automater v0.0.6
	github.com/mustafaturan/bus/v3 v3.0.3
	gorm.io/datatypes v1.0.6
)

require (
	github.com/NubeIO/git v0.0.3 // indirect
	github.com/NubeIO/lib-store v0.0.1 // indirect
	github.com/NubeIO/lib-systemctl-go v0.0.5 // indirect
	github.com/THREATINT/go-net v1.2.10 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/google/go-github/v32 v32.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmattheis/go-timemath v1.0.1 // indirect
	github.com/mcnijman/go-emailaddress v1.1.0 // indirect
	github.com/mustafaturan/monoton/v2 v2.0.2 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
)
