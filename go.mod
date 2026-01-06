module device-simulator/v2

go 1.24.4

toolchain go1.24.5

require (
	github.com/aarondl/sqlboiler/v4 v4.19.7
	github.com/aarondl/strmangle v0.0.9
	github.com/eliona-smart-building-assistant/app-integration-tests v1.1.6
	github.com/eliona-smart-building-assistant/backend-frm v0.0.3
	github.com/eliona-smart-building-assistant/go-eliona-api-client/v3 v3.0.6
	github.com/eliona-smart-building-assistant/go-eliona/v2 v2.0.5
	github.com/eliona-smart-building-assistant/go-utils v1.1.10
	github.com/friendsofgo/errors v0.9.2
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	gopkg.in/yaml.v3 v3.0.1
)

// Bugfix see: https://github.com/volatiletech/sqlboiler/blob/91c4f335dd886d95b03857aceaf17507c46f9ec5/README.md
// decimal library showing errors like: pq: encode: unknown type types.NullDecimal is a result of a too-new and broken version of the github.com/ericlargergren/decimal package, use the following version in your go.mod: github.com/ericlagergren/decimal v0.0.0-20181231230500-73749d4874d5
replace github.com/ericlagergren/decimal => github.com/ericlagergren/decimal v0.0.0-20181231230500-73749d4874d5

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.20.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.13.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.2 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.6.0 // indirect
	github.com/ProtonMail/go-crypto v1.3.0 // indirect
	github.com/ProtonMail/gopenpgp/v3 v3.3.0 // indirect
	github.com/aarondl/inflect v0.0.2 // indirect
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eliona-smart-building-assistant/go-eliona v1.9.39 // indirect
	github.com/eliona-smart-building-assistant/go-eliona-api-client/v2 v2.6.12 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.4 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/jackc/pgx/v5 v5.7.6 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.9.2 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/sqlboiler/v4 v4.19.1 // indirect
	github.com/volatiletech/strmangle v0.0.8 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/image v0.19.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
)
