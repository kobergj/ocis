module github.com/owncloud/ocis/v2

go 1.22.0

require (
	dario.cat/mergo v1.0.1
	github.com/CiscoM31/godata v1.0.10
	github.com/KimMachineGun/automemlimit v0.6.1
	github.com/Masterminds/semver v1.5.0
	github.com/MicahParks/keyfunc/v2 v2.1.0
	github.com/Nerzal/gocloak/v13 v13.9.0
	github.com/bbalet/stopwords v1.0.0
	github.com/beevik/etree v1.4.1
	github.com/blevesearch/bleve/v2 v2.4.2
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/coreos/go-oidc/v3 v3.11.0
	github.com/cs3org/go-cs3apis v0.0.0-20240724121416-062c4e3046cb
	github.com/cs3org/reva/v2 v2.24.2-0.20240918135034-5a494dc2ebf7
	github.com/dhowden/tag v0.0.0-20230630033851-978a0926ee25
	github.com/dutchcoders/go-clamd v0.0.0-20170520113014-b970184f4d9e
	github.com/egirna/icap-client v0.1.1
	github.com/gabriel-vasile/mimetype v1.4.4
	github.com/ggwhite/go-masker v1.1.0
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/render v1.0.3
	github.com/go-ldap/ldap/v3 v3.4.8
	github.com/go-ldap/ldif v0.0.0-20200320164324-fd88d9b715b3
	github.com/go-micro/plugins/v4/client/grpc v1.2.1
	github.com/go-micro/plugins/v4/logger/zerolog v1.2.0
	github.com/go-micro/plugins/v4/registry/consul v1.2.1
	github.com/go-micro/plugins/v4/registry/etcd v1.2.0
	github.com/go-micro/plugins/v4/registry/kubernetes v1.1.2
	github.com/go-micro/plugins/v4/registry/mdns v1.2.0
	github.com/go-micro/plugins/v4/registry/memory v1.2.0
	github.com/go-micro/plugins/v4/registry/nats v1.2.2
	github.com/go-micro/plugins/v4/server/grpc v1.2.0
	github.com/go-micro/plugins/v4/server/http v1.2.2
	github.com/go-micro/plugins/v4/store/nats-js-kv v0.0.0-20231226212146-94a49ba3e06e
	github.com/go-micro/plugins/v4/wrapper/breaker/gobreaker v1.2.0
	github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus v1.2.0
	github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry v1.2.0
	github.com/go-playground/validator/v10 v10.22.1
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/protobuf v1.5.4
	github.com/google/go-cmp v0.6.0
	github.com/google/go-tika v0.3.1
	github.com/google/uuid v1.6.0
	github.com/gookit/config/v2 v2.2.5
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0
	github.com/invopop/validation v0.8.0
	github.com/jellydator/ttlcache/v2 v2.11.1
	github.com/jellydator/ttlcache/v3 v3.3.0
	github.com/jinzhu/now v1.1.5
	github.com/justinas/alice v1.2.0
	github.com/kovidgoyal/imaging v1.6.3
	github.com/leonelquinteros/gotext v1.6.1
	github.com/libregraph/idm v0.5.0
	github.com/libregraph/lico v0.62.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mna/pigeon v1.2.1
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/nats-io/nats-server/v2 v2.10.20
	github.com/nats-io/nats.go v1.37.0
	github.com/oklog/run v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/ginkgo/v2 v2.20.2
	github.com/onsi/gomega v1.34.1
	github.com/open-policy-agent/opa v0.68.0
	github.com/orcaman/concurrent-map v1.0.0
	github.com/owncloud/libre-graph-api-go v1.0.5-0.20240820135012-5fac8096ce9c
	github.com/pkg/errors v0.9.1
	github.com/pkg/xattr v0.4.10
	github.com/prometheus/client_golang v1.20.2
	github.com/r3labs/sse/v2 v2.10.0
	github.com/riandyrn/otelchi v0.9.0
	github.com/rogpeppe/go-internal v1.12.0
	github.com/rs/cors v1.11.0
	github.com/rs/zerolog v1.33.0
	github.com/shamaton/msgpack/v2 v2.2.0
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/afero v1.11.0
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.9.0
	github.com/test-go/testify v1.1.4
	github.com/thejerf/suture/v4 v4.0.5
	github.com/tidwall/gjson v1.17.3
	github.com/tus/tusd/v2 v2.5.0
	github.com/unrolled/secure v1.14.0
	github.com/urfave/cli/v2 v2.27.4
	github.com/xhit/go-simple-mail/v2 v2.16.0
	go-micro.dev/v4 v4.11.0
	go.etcd.io/bbolt v1.3.11
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.55.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0
	go.opentelemetry.io/contrib/zpages v0.54.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.30.0
	go.opentelemetry.io/otel/sdk v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
	golang.org/x/crypto v0.27.0
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56
	golang.org/x/image v0.20.0
	golang.org/x/net v0.29.0
	golang.org/x/oauth2 v0.23.0
	golang.org/x/sync v0.8.0
	golang.org/x/term v0.24.0
	golang.org/x/text v0.18.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1
	google.golang.org/grpc v1.66.1
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.5.1
	stash.kopano.io/kgol/rndm v1.1.2
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.2 // indirect
	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371 // indirect
	github.com/RoaringBitmap/roaring v1.9.3 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/alexedwards/argon2id v1.0.0 // indirect
	github.com/amoghe/go-crypt v0.0.0-20220222110647-20eada5f5964 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go v1.45.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bits-and-blooms/bitset v1.12.0 // indirect
	github.com/blevesearch/bleve_index_api v1.1.10 // indirect
	github.com/blevesearch/geo v0.1.20 // indirect
	github.com/blevesearch/go-faiss v1.0.20 // indirect
	github.com/blevesearch/go-porterstemmer v1.0.3 // indirect
	github.com/blevesearch/gtreap v0.1.1 // indirect
	github.com/blevesearch/mmap-go v1.0.4 // indirect
	github.com/blevesearch/scorch_segment_api/v2 v2.2.15 // indirect
	github.com/blevesearch/segment v0.9.1 // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/blevesearch/upsidedown_store_api v1.0.2 // indirect
	github.com/blevesearch/vellum v1.0.10 // indirect
	github.com/blevesearch/zapx/v11 v11.3.10 // indirect
	github.com/blevesearch/zapx/v12 v12.3.10 // indirect
	github.com/blevesearch/zapx/v13 v13.3.10 // indirect
	github.com/blevesearch/zapx/v14 v14.3.10 // indirect
	github.com/blevesearch/zapx/v15 v15.3.13 // indirect
	github.com/blevesearch/zapx/v16 v16.1.5 // indirect
	github.com/bluele/gcache v0.0.2 // indirect
	github.com/bombsimon/logrusr/v3 v3.1.0 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/ceph/go-ceph v0.18.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73 // indirect
	github.com/cilium/ebpf v0.9.1 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/containerd/cgroups/v3 v3.0.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cornelk/hashmap v1.0.8 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/crewjam/httperr v0.2.0 // indirect
	github.com/crewjam/saml v0.4.14 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/egirna/icap v0.0.0-20181108071049-d5ee18bd70bc // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/emvi/iso-639-1 v1.1.0 // indirect
	github.com/evanphx/json-patch/v5 v5.5.0 // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/frankban/quicktest v1.14.6 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gdexlab/go-render v1.0.1 // indirect
	github.com/go-acme/lego/v4 v4.4.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.5 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-git/go-git/v5 v5.11.0 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.3 // indirect
	github.com/go-jose/go-jose/v4 v4.0.2 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-micro/plugins/v4/events/natsjs v1.2.2-0.20231215124540-f7f8d3274bf9 // indirect
	github.com/go-micro/plugins/v4/store/nats-js v1.2.1-0.20231129143103-d72facc652f0 // indirect
	github.com/go-micro/plugins/v4/store/redis v1.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/go-test/deep v1.1.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.2.1 // indirect
	github.com/goccy/go-yaml v1.11.2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/geo v0.0.0-20210211234256-740aa86cb551 // indirect
	github.com/golang/glog v1.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/google/flatbuffers v2.0.8+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20240827171923-fa2c70bbbfe5 // indirect
	github.com/google/renameio/v2 v2.0.0 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/gookit/goutil v0.6.15 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/schema v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/consul/api v1.25.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v1.1.5 // indirect
	github.com/hashicorp/go-plugin v1.6.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.6.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juliangruber/go-intersect v1.1.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/libregraph/oidc-go v1.1.0 // indirect
	github.com/longsleep/go-metrics v1.0.0 // indirect
	github.com/longsleep/rndm v1.2.0 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/maxymania/go-system v0.0.0-20170110133659-647cc364bf0b // indirect
	github.com/mendsley/gojwk v0.0.0-20141217222730-4d5ec6e58103 // indirect
	github.com/miekg/dns v1.1.57 // indirect
	github.com/mileusna/useragent v1.3.4 // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.0.66 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/jwt/v2 v2.5.8 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/opencontainers/runtime-spec v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pablodz/inotifywaitgo v0.0.7 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/pquerna/cachecontrol v0.2.0 // indirect
	github.com/prometheus/alertmanager v0.26.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/prometheus/statsd_exporter v0.22.8 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/russellhaering/goxmldsig v1.4.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/sercand/kuberesolver/v5 v5.1.1 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/sethvargo/go-password v0.2.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/skeema/knownhosts v1.2.1 // indirect
	github.com/sony/gobreaker v0.5.0 // indirect
	github.com/spacewander/go-suffix-tree v0.0.0-20191010040751-0865e368c784 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/studio-b12/gowebdav v0.0.0-20221015232716-17255f2e7423 // indirect
	github.com/tchap/go-patricia/v2 v2.3.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/trustelem/zxcvbn v1.0.1 // indirect
	github.com/wk8/go-ordered-map v1.0.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	github.com/yashtewari/glob-intersection v0.2.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/v3 v3.5.12 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20240823204242-4ba0660f739c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	gopkg.in/cenkalti/backoff.v1 v1.1.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace github.com/studio-b12/gowebdav => github.com/aduffeck/gowebdav v0.0.0-20231215102054-212d4a4374f6

replace github.com/egirna/icap-client => github.com/fschade/icap-client v0.0.0-20240802074440-aade4a234387

replace github.com/unrolled/secure => github.com/DeepDiver1975/secure v0.0.0-20240611112133-abc838fb797c

replace github.com/go-micro/plugins/v4/store/nats-js-kv => github.com/kobergj/plugins/v4/store/nats-js-kv v0.0.0-20240807130109-f62bb67e8c90

// exclude the v2 line of go-sqlite3 which was released accidentally and prevents pulling in newer versions of go-sqlite3
// see https://github.com/mattn/go-sqlite3/issues/965 for more details
exclude github.com/mattn/go-sqlite3 v2.0.3+incompatible
