FROM golang:1.22-bullseye

ENV CGO_ENABLED=0

RUN apt-get update -y
RUN apt-get install -y --no-install-recommends python3-pip=20.3.4-4+deb11u1


RUN pip --no-cache-dir install junit2html==30.1.3

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.18.0

# golang:1.22-bullseye contains an existing go.mod under /usr/local/go/src
# which is required to run "go install" or "go get" commands.
WORKDIR /usr/local/go/src

# New in golang 1.22, "go install" must be used to install go binaries.
# Temporarily setting a non-default go module cache directory so that
# go binaries and associated modules will be installed under /go/bin
# and /go/pkg/mod respectively. This will ensure the binaries are
# accessible in the default PATH env based on the golang:1.22+ image.

RUN export GOMODCACHE=/go/pkg/mod && go install github.com/jstemmer/go-junit-report@latest

RUN export GOMODCACHE=/go/pkg/mod && go install github.com/ains/go-test-html@latest

RUN export GOMODCACHE=/go/pkg/mod && go install github.com/axw/gocov/gocov@latest

RUN export GOMODCACHE=/go/pkg/mod && go install github.com/matm/gocov-html/cmd/gocov-html@latest

RUN export GOMODCACHE=/go/pkg/mod && go install github.com/wadey/gocovmerge@latest

RUN export GOMODCACHE=/go/pkg/mod && go install github.com/t-yuki/gocover-cobertura@latest

RUN export GOMODCACHE=/go/pkg/mod && go install golang.org/x/lint/golint@v0.0.0-20210508222113-6edffad5e616

RUN export GOMODCACHE=/go/pkg/mod && go install gotest.tools/gotestsum@latest

# The default go module cache directory is under /root/go/pkg/mod.
ENV GOPATH ""
ENV GO111MODULE=on
RUN for mod in \
    ariga.io/atlas-go-sdk@v0.5.3 \
    ariga.io/atlas-go-sdk@v0.5.6 \
    ariga.io/atlas-provider-gorm@v0.3.4 \
    ariga.io/atlas-provider-gorm@v0.5.0 \
    filippo.io/edwards25519@v1.1.0 \
    github.com/BurntSushi/toml@v1.3.2 \
    github.com/SentimensRG/ctx@v0.0.0-20180729130232-0bfd988c655d \
    github.com/alitto/pond@v1.8.3 \
    github.com/alitto/pond@v1.9.0 \
    github.com/antihax/optional@v1.0.0 \
    github.com/beevik/etree@v1.4.0 \
    github.com/c-seeger/mac-gen-go@v0.0.0-20210816124238-465118e656da \
    github.com/cespare/xxhash/v2@v2.3.0 \
    github.com/confluentinc/confluent-kafka-go@v1.9.2 \
    github.com/davecgh/go-spew@v1.1.2-0.20180830191138-d8f796af33cc \
    github.com/deckarep/golang-set/v2@v2.6.0 \
    github.com/dgryski/go-rendezvous@v0.0.0-20200823014737-9f7001d12a5f \
    github.com/fsnotify/fsnotify@v1.7.0 \
    github.com/gammazero/deque@v0.2.1 \
    github.com/gammazero/workerpool@v1.1.3 \
    github.com/go-jose/go-jose/v3@v3.0.3 \
    github.com/go-redsync/redsync/v4@v4.13.0 \
    github.com/go-sql-driver/mysql@v1.8.1 \
    github.com/go-stack/stack@v1.8.1 \
    github.com/gocarina/gocsv@v0.0.0-20240520201108-78e41c74b4b1 \
    github.com/golang-jwt/jwt/v4@v4.5.0 \
    github.com/golang-jwt/jwt@v3.2.2+incompatible \
    github.com/golang-sql/civil@v0.0.0-20220223132316-b832511892a9 \
    github.com/golang-sql/sqlexp@v0.1.0 \
    github.com/golang/glog@v1.2.2 \
    github.com/golang/protobuf@v1.5.4 \
    github.com/gomodule/redigo@v1.9.2 \
    github.com/google/go-cmp@v0.6.0 \
    github.com/google/uuid@v1.6.0 \
    github.com/gorilla/mux@v1.8.1 \
    github.com/gorilla/websocket@v1.5.3 \
    github.com/hashicorp/errwrap@v1.1.0 \
    github.com/hashicorp/go-multierror@v1.1.1 \
    github.com/hashicorp/hcl@v1.0.0 \
    github.com/jackc/pgpassfile@v1.0.0 \
    github.com/jackc/pgservicefile@v0.0.0-20231201235250-de7065d80cb9 \
    github.com/jackc/pgservicefile@v0.0.0-20240606120523-5a60cdf6a761 \
    github.com/jackc/pgx/v5@v5.5.5 \
    github.com/jackc/pgx/v5@v5.6.0 \
    github.com/jackc/puddle/v2@v2.2.1 \
    github.com/jinzhu/inflection@v1.0.0 \
    github.com/jinzhu/now@v1.1.5 \
    github.com/jonboulle/clockwork@v0.4.0 \
    github.com/korylprince/ipnetgen@v1.0.1 \
    github.com/kr/fs@v0.1.0 \
    github.com/kylelemons/godebug@v1.1.0 \
    github.com/lib/pq@v1.10.9 \
    github.com/magefile/mage@v1.15.0 \
    github.com/magiconair/properties@v1.8.7 \
    github.com/mattn/go-colorable@v0.1.13 \
    github.com/mattn/go-isatty@v0.0.20 \
    github.com/mattn/go-sqlite3@v1.14.22 \
    github.com/microsoft/go-mssqldb@v1.7.1 \
    github.com/microsoft/go-mssqldb@v1.7.2 \
    github.com/mitchellh/mapstructure@v1.5.0 \
    github.com/mmanda-extr/cdule@v0.0.0-20230922063427-fe29fb62e13e \
    github.com/nitishm/go-rejson/v4@v4.2.0 \
    github.com/openconfig/gnmi@v0.11.0 \
    github.com/openconfig/gnoi@v0.0.0-20230208173232-3df08ff73842 \
    github.com/openconfig/goyang@v1.5.0 \
    github.com/openconfig/ygot@v0.29.19 \
    github.com/pelletier/go-toml/v2@v2.2.2 \
    github.com/petermattis/goid@v0.0.0-20240503122002-4b96552b8156 \
    github.com/petermattis/goid@v0.0.0-20240607163614-bb94eb51e7a7 \
    github.com/pkg/sftp@v1.13.6 \
    github.com/playwright-community/playwright-go@v0.4401.1 \
    github.com/pmezard/go-difflib@v1.0.1-0.20181226105442-5d4384ee4fb2 \
    github.com/redis/go-redis/v9@v9.5.1 \
    github.com/redis/go-redis/v9@v9.5.3 \
    github.com/rifflock/lfshook@v0.0.0-20180920164130-b9218ef580f5 \
    github.com/robfig/cron/v3@v3.0.1 \
    github.com/rogpeppe/go-internal@v1.11.0 \
    github.com/rs/zerolog@v1.32.0 \
    github.com/rs/zerolog@v1.33.0 \
    github.com/sagikazarmark/locafero@v0.4.0 \
    github.com/sagikazarmark/locafero@v0.6.0 \
    github.com/sagikazarmark/slog-shim@v0.1.0 \
    github.com/sasha-s/go-deadlock@v0.3.1 \
    github.com/sirikothe/gotextfsm@v1.0.1-0.20210615061731-b060a6ccd955 \
    github.com/sirupsen/logrus@v1.9.3 \
    github.com/sourcegraph/conc@v0.3.0 \
    github.com/spf13/afero@v1.11.0 \
    github.com/spf13/cast@v1.6.0 \
    github.com/spf13/pflag@v1.0.5 \
    github.com/spf13/viper@v1.18.2 \
    github.com/spf13/viper@v1.19.0 \
    github.com/stretchr/testify@v1.9.0 \
    github.com/subosito/gotenv@v1.6.0 \
    github.com/svatantra/go-netconf@v0.0.0-20220211042307-2b1fe34afc5a \
    github.com/svatantra/lumberjack@v2.1.6+incompatible \
    github.com/ziutek/telnet@v0.0.0-20180329124119-c3b780dc415b \
    go.uber.org/multierr@v1.11.0 \
    golang.org/x/crypto@v0.23.0 \
    golang.org/x/crypto@v0.25.0 \
    golang.org/x/exp@v0.0.0-20240506185415-9bf2ced13842 \
    golang.org/x/exp@v0.0.0-20240707233637-46b078467d37 \
    golang.org/x/net@v0.25.0 \
    golang.org/x/net@v0.27.0 \
    golang.org/x/oauth2@v0.20.0 \
    golang.org/x/oauth2@v0.21.0 \
    golang.org/x/sync@v0.7.0 \
    golang.org/x/sys@v0.20.0 \
    golang.org/x/sys@v0.22.0 \
    golang.org/x/text@v0.15.0 \
    golang.org/x/text@v0.16.0 \
    google.golang.org/genproto/googleapis/rpc@v0.0.0-20240506185236-b8a5c65736ae \
    google.golang.org/genproto/googleapis/rpc@v0.0.0-20240709173604-40e1e62336c5 \
    google.golang.org/grpc@v1.63.2 \
    google.golang.org/grpc@v1.65.0 \
    google.golang.org/protobuf@v1.34.1 \
    google.golang.org/protobuf@v1.34.2 \
    gopkg.in/ini.v1@v1.67.0 \
    gopkg.in/natefinch/lumberjack.v2@v2.2.1 \
    gopkg.in/yaml.v2@v2.4.0 \
    gopkg.in/yaml.v3@v3.0.1 \
    gorm.io/driver/mysql@v1.5.6 \
    gorm.io/driver/mysql@v1.5.7 \
    gorm.io/driver/postgres@v1.5.7 \
    gorm.io/driver/postgres@v1.5.9 \
    gorm.io/driver/sqlite@v1.5.5 \
    gorm.io/driver/sqlite@v1.5.6 \
    gorm.io/driver/sqlserver@v1.5.3 \
    gorm.io/gorm@v1.25.10 \
    ; do go get -v $mod ; done

# Restore the default WORKDIR based on the golang:1.22-alpine image
WORKDIR /go
