module github.com/myxtype/filecoin-client

go 1.15

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/davecgh/go-spew v1.1.1
	github.com/filecoin-project/go-address v0.0.5-0.20201103152444-f2023ef3f5bb
	github.com/filecoin-project/go-state-types v0.0.0-20201102161440-c8033295a1fc
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-cid v0.0.7
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/shopspring/decimal v1.2.0
	github.com/supranational/blst v0.0.0-00010101000000-000000000000
	github.com/whyrusleeping/cbor-gen v0.0.0-20200826160007-0b9f6c5fb163
)

replace github.com/supranational/blst => ./extern/blst
