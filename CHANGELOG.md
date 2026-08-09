# Changelog

## [1.23.0](https://github.com/mhughdo/solana-go/compare/v1.22.0...v1.23.0) (2026-08-09)


### Features

* actualize alut program ([9e89b53](https://github.com/mhughdo/solana-go/commit/9e89b532dd145484a04f002b7613bc52080f5718))
* actualize alut program ([4e693d6](https://github.com/mhughdo/solana-go/commit/4e693d69d5821daaf69e42eb4711c20f6c65de7b))
* actualize ata program ([1b2dbbd](https://github.com/mhughdo/solana-go/commit/1b2dbbdc16510da81bab559c7a6bbfd06216a837))
* actualize ata program ([983d6ef](https://github.com/mhughdo/solana-go/commit/983d6ef3520ed9ab7f0fc9c7824ce7922276db09))
* actualize rpc methods ([46f9602](https://github.com/mhughdo/solana-go/commit/46f9602d582a62a427a0053c0d9b887caefa0745))
* actualize rpc methods ([2e23eea](https://github.com/mhughdo/solana-go/commit/2e23eeaaec7550fadea505e754dd143c1606e2e2))
* actualize stake program ([c950396](https://github.com/mhughdo/solana-go/commit/c950396397f1898eebfa902a710b781233df6a55))
* actualize stake program ([419521a](https://github.com/mhughdo/solana-go/commit/419521ac60738b7f5c303d96a583dd40878d2656))
* actualize sysvar constants ([9345f26](https://github.com/mhughdo/solana-go/commit/9345f26e653ffe0ea93f5f71e424044d106daf41))
* actualize sysvar constants ([653c634](https://github.com/mhughdo/solana-go/commit/653c634d12c44727df41e04c89c7245ddf281124))
* Add `.Response()` and `.Err()` method for all subscriber ([#177](https://github.com/mhughdo/solana-go/issues/177)) ([a8900f8](https://github.com/mhughdo/solana-go/commit/a8900f8a1f2cddbd2dc09c81d036eeb956274ac8))
* add Associated Token Account Program: CreateIdempotent instruction ([#4](https://github.com/mhughdo/solana-go/issues/4)) ([3d55abb](https://github.com/mhughdo/solana-go/commit/3d55abb4cc8431c09eeaccc808563f0f30226ade))
* add getters to txn with meta ([48d196b](https://github.com/mhughdo/solana-go/commit/48d196b47f79212bd4e97fb61d35b0482b5ef8e9))
* add maxSupportedTransactionVersion to getParsedTransaction opt ([#143](https://github.com/mhughdo/solana-go/issues/143)) ([4493616](https://github.com/mhughdo/solana-go/commit/4493616ba98ffe31e39d376f8e5e2e5d925979d7))
* add nonce account support ([#456](https://github.com/mhughdo/solana-go/issues/456)) ([336881c](https://github.com/mhughdo/solana-go/commit/336881cfbe66902ee0ec07b25dc1912186508a72))
* Add SetAccounts implementation to be able to access accounts for Create instruction after decoding ([#337](https://github.com/mhughdo/solana-go/issues/337)) ([ac7125b](https://github.com/mhughdo/solana-go/commit/ac7125b9dcfed87fef00956799433b594aef74c0))
* add SetLoadedAccountsDataSizeLimitInstruction ([7926abb](https://github.com/mhughdo/solana-go/commit/7926abb88edeaf915adb5d682c485df528252f3a))
* add SetLoadedAccountsDataSizeLimitInstruction ([3433ac4](https://github.com/mhughdo/solana-go/commit/3433ac45ab7764687a1973be6447875d6d61f0a7))
* add slot field to TransactionSignatureResult ([#7](https://github.com/mhughdo/solana-go/issues/7)) ([abafd3e](https://github.com/mhughdo/solana-go/commit/abafd3e9c0d08b531801d2232d9939ba077924e9))
* add support for getRecentPrioritizationFees RPC method ([#140](https://github.com/mhughdo/solana-go/issues/140)) ([290a21a](https://github.com/mhughdo/solana-go/commit/290a21adc5d262d93baba0378ebf1dc9a5a1d21d))
* add sysvars ([#457](https://github.com/mhughdo/solana-go/issues/457)) ([f777896](https://github.com/mhughdo/solana-go/commit/f7778964b61f6ac2622a194fb034dd92da59c67f))
* add t22 program ([abdc7b7](https://github.com/mhughdo/solana-go/commit/abdc7b75909cd242dfe309b42fb6b4205300de91))
* add t22 program ([485ceb4](https://github.com/mhughdo/solana-go/commit/485ceb4de7d6ec657cd588fa3b0db4b2b1c0cc19))
* add token-2022 extensions ([04dfc79](https://github.com/mhughdo/solana-go/commit/04dfc794b4061522954fa3f10f5bfad2ff982728))
* add token-2022 extensions ([db2fdaa](https://github.com/mhughdo/solana-go/commit/db2fdaa42ed9062d23eb51ae278856197fb66633))
* add transaction signature subscribe ([#1](https://github.com/mhughdo/solana-go/issues/1)) ([eb17ac1](https://github.com/mhughdo/solana-go/commit/eb17ac13814f917cb4f9b3100a752e6784ef83b6))
* added maxSupportedTransactionVersion field to getBlock opts ([6ead48a](https://github.com/mhughdo/solana-go/commit/6ead48adf2b0313e7ef0aaf53bf0cd2e53285f9b))
* **address-lookup-table:** one-call resolve for message lookups (closes [#262](https://github.com/mhughdo/solana-go/issues/262)) ([#445](https://github.com/mhughdo/solana-go/issues/445)) ([34beab1](https://github.com/mhughdo/solana-go/commit/34beab1e231b6e85b7232361997d70bdd79825ef))
* allow apps to create solana client with custom http client ([#2](https://github.com/mhughdo/solana-go/issues/2)) ([7864e02](https://github.com/mhughdo/solana-go/commit/7864e02024ead7eeb1ccf216aa57bf9aae7b666b))
* **associated-token-account:** add token program aware ata helpers ([f0960f5](https://github.com/mhughdo/solana-go/commit/f0960f50b7cd3d991154778d50d1e8314c5e08ec))
* **associated-token-account:** add token program aware ata helpers ([3cb13b3](https://github.com/mhughdo/solana-go/commit/3cb13b392078f39c2683dc52d6725876ba339a72))
* change the return type of transactionSubscribe ([#3](https://github.com/mhughdo/solana-go/issues/3)) ([4b81b3e](https://github.com/mhughdo/solana-go/commit/4b81b3e180e83d03b4b4fa03488d6230b3a5bc50))
* handle HeliusWS error ([#6](https://github.com/mhughdo/solana-go/issues/6)) ([223c258](https://github.com/mhughdo/solana-go/commit/223c258d82ab22334c43bfc4c85ee7ff5edb9253))
* increase s.stream channel size ([8d5e96e](https://github.com/mhughdo/solana-go/commit/8d5e96ee88b37b73cd9aa73facd3230984b9c96f))
* Introduce  instruction to upgrade legacy nonce accounts. ([c9270cd](https://github.com/mhughdo/solana-go/commit/c9270cdb8d159673b424cadc221eec2daab17f5f))
* Introduce  instruction to upgrade legacy nonce accounts. ([da05bb9](https://github.com/mhughdo/solana-go/commit/da05bb93ecab7ea8fa4b5b78c9b63f93bdc21657))
* is token mint classifier ([4f72982](https://github.com/mhughdo/solana-go/commit/4f72982442c9b3c166b72dbb2de730f58b575539))
* is token mint classifier ([4920307](https://github.com/mhughdo/solana-go/commit/4920307384f2866ce4cafa10707653b56c13d16d))
* **jsonrpc:** add CustomHeader http.Header for multi-value headers ([20b37ba](https://github.com/mhughdo/solana-go/commit/20b37ba403c438ebe914b43ff7081f9598832d0c))
* reduce stream and err buffers ([55d5fba](https://github.com/mhughdo/solana-go/commit/55d5fba4013c3fce342d182fb8155e1b4e4cca2f))
* **rpc:** add getTransactionsForAddress client method (closes [#343](https://github.com/mhughdo/solana-go/issues/343)) ([#450](https://github.com/mhughdo/solana-go/issues/450)) ([9e538c8](https://github.com/mhughdo/solana-go/commit/9e538c84246ef1a3abfbcc96cac7a3196889b479))
* **rpc:** add NewWithCommitment / NewWithTimeout / NewWithTimeoutAndCommitment ([#436](https://github.com/mhughdo/solana-go/issues/436)) ([e93ff5e](https://github.com/mhughdo/solana-go/commit/e93ff5e937733daca5fed01c362961c4c8aead25)), closes [#414](https://github.com/mhughdo/solana-go/issues/414)
* **rpc:** forward minContextSlot in 4 remaining JSON-RPC endpoints ([#448](https://github.com/mhughdo/solana-go/issues/448)) ([c176402](https://github.com/mhughdo/solana-go/commit/c176402c339e25c704f4eba9be540d56f770feb0))
* **rpc:** forward MinContextSlot in getBalance/getLatestBlockhash/getSlot/getTokenAccountBalance ([#442](https://github.com/mhughdo/solana-go/issues/442)) ([b8e70e8](https://github.com/mhughdo/solana-go/commit/b8e70e8c5cdd1229c6126f00090527664aa8496c))
* **rpc:** forward MinContextSlot in getProgramAccounts and getTokenAccounts ([#431](https://github.com/mhughdo/solana-go/issues/431)) ([17984a5](https://github.com/mhughdo/solana-go/commit/17984a55c17ab0fc9f308872a43b737601d6a8da))
* **rpc:** support EncodingJSON in GetTransaction ([#420](https://github.com/mhughdo/solana-go/issues/420)) ([b906b70](https://github.com/mhughdo/solana-go/commit/b906b70527a5dfed358090e27dd7f4a7f12749c3))
* send tx V0 with AddressLookupTable ([bbd443b](https://github.com/mhughdo/solana-go/commit/bbd443bc28d6e3ac665667ce154771ee18695745))
* stake state types & ext tests ([6325515](https://github.com/mhughdo/solana-go/commit/6325515ae6686d10351a1649e3b46951811e1fbe))
* stake state types & ext tests ([ea77e31](https://github.com/mhughdo/solana-go/commit/ea77e318bb8158f391b349dd6c736fc457dbda46))
* support ComputeBudgetProgram instructions ([#132](https://github.com/mhughdo/solana-go/issues/132)) ([1985c48](https://github.com/mhughdo/solana-go/commit/1985c484726f29a9da2281db1de4447986c77442))
* support partial sign & marshal partial signed transaction ([5eb8de2](https://github.com/mhughdo/solana-go/commit/5eb8de2d6bd4f61d5c4855af4201f00e89d40b7a))
* support TritonRpc transactionSubscribe method ([#8](https://github.com/mhughdo/solana-go/issues/8)) ([a9d0d27](https://github.com/mhughdo/solana-go/commit/a9d0d27eb9dc7e08a9c1d4f8920e8c39ad32ce8d))
* vote program complete ([173d7f4](https://github.com/mhughdo/solana-go/commit/173d7f4059fd12e5d747d5b0ffea7850211086a2))
* vote program complete ([dcff584](https://github.com/mhughdo/solana-go/commit/dcff5845b15ca32c90a9b4336e8a902917132854))
* **wallet:** derive PrivateKey/Wallet from BIP-39 mnemonic ([#429](https://github.com/mhughdo/solana-go/issues/429)) ([89ef706](https://github.com/mhughdo/solana-go/commit/89ef706472ad49a9622a058497852711f7bd3771))
* **ws:** support dataSlice in AccountSubscribe ([#433](https://github.com/mhughdo/solana-go/issues/433)) ([fb31fb1](https://github.com/mhughdo/solana-go/commit/fb31fb13b42141bb6067c7447b8618e7e848b97b))
* **ws:** support dataSlice in ProgramSubscribe ([#434](https://github.com/mhughdo/solana-go/issues/434)) ([950b110](https://github.com/mhughdo/solana-go/commit/950b110b8f369de33143705cfba0b8da7d240d6f))
* **ws:** support enableReceivedNotification in SignatureSubscribe ([#432](https://github.com/mhughdo/solana-go/issues/432)) ([810f171](https://github.com/mhughdo/solana-go/commit/810f171ff933c1508e9526a2a536a287cac7c386))
* **zk:** add ElGamal & AES key derivation ([#413](https://github.com/mhughdo/solana-go/issues/413)) ([9fcbf0c](https://github.com/mhughdo/solana-go/commit/9fcbf0ced8af5b5ad975e45999478b5c36c3e65a))


### Bug Fixes

* actualize simulate transaction ([364496d](https://github.com/mhughdo/solana-go/commit/364496d1aa4ce17b47b5428033563c1d17ffbdea))
* actualize simulate transaction ([48975d7](https://github.com/mhughdo/solana-go/commit/48975d7dcd5602d11d9d8ef7db94831293d80966))
* Add `ShortID` option for some WebSocket RPC which not support int63/uint64 ID ([#179](https://github.com/mhughdo/solana-go/issues/179)) ([e1dd29f](https://github.com/mhughdo/solana-go/commit/e1dd29f38124ab3c61239d279ba3a1fe00629514))
* allign rpc client with agave ([49bc8d6](https://github.com/mhughdo/solana-go/commit/49bc8d6c44e53b9ab8710c8c333e479514678e40))
* allign rpc client with agave ([c985b99](https://github.com/mhughdo/solana-go/commit/c985b99fd333a46eaacd8b5dfab2db697bdab54c))
* check if the channel is closed before returning ws.result ([#198](https://github.com/mhughdo/solana-go/issues/198)) ([6beb7f8](https://github.com/mhughdo/solana-go/commit/6beb7f85b5389f7750504e32d6028853a688cc26))
* correct go.mod format and update CI matrix for Go 1.23+ ([eced08f](https://github.com/mhughdo/solana-go/commit/eced08ffe5294731f71aa96803e1ad483e6a846a))
* correct simulateTransaction response payload decoding ([06aad92](https://github.com/mhughdo/solana-go/commit/06aad92b02cc1a000ac255255e3d6775447b639e))
* correct the order of Slot field in TransactionSignatureResult ([535b1e1](https://github.com/mhughdo/solana-go/commit/535b1e1484cecbe332d278a8294a666716a4adda))
* correct TransactionSignatureResult struct ([bf6ee61](https://github.com/mhughdo/solana-go/commit/bf6ee6138e479827407f3338daef16c1fe5fa4a9))
* enhance getUint64 function to handle string inputs ([5309095](https://github.com/mhughdo/solana-go/commit/53090952ffc598c1870617b1727179135994ec65))
* Handle client.newRequest returning nil request ([7a0f053](https://github.com/mhughdo/solana-go/commit/7a0f05371ad251bd233f1d217329cada44215c4e))
* Handle client.newRequest returning nil request ([0e0d186](https://github.com/mhughdo/solana-go/commit/0e0d1866c493b4d23c6ceb2fec9bf710ddbb2cb0))
* keep websocket request IDs within JSON-safe range ([8ed3105](https://github.com/mhughdo/solana-go/commit/8ed31050f7af62f26b5615f40546bb498cab9219))
* lint error ([add9aa5](https://github.com/mhughdo/solana-go/commit/add9aa5b46d4b07bc1aba365eecdb241119ddd0e))
* lookup tables ordering ([1ba0d4b](https://github.com/mhughdo/solana-go/commit/1ba0d4bde87fad00340e8fdd311a4ea3811373ce))
* lookup tables ordering ([f304171](https://github.com/mhughdo/solana-go/commit/f3041712f82e780b1b3d9a61bccb566097b65186))
* memo program parity ([f832bd8](https://github.com/mhughdo/solana-go/commit/f832bd84b73eb5a3f2222072d45a8d8c5efeac56))
* memo program parity ([7550777](https://github.com/mhughdo/solana-go/commit/755077780d34ed78d3a8c00a9e7d54542e66d051))
* message bugs ([9ef2d24](https://github.com/mhughdo/solana-go/commit/9ef2d244324a77289615327315a96e00db6dcaa3))
* **message:** json version detection ([1fd2201](https://github.com/mhughdo/solana-go/commit/1fd2201431de71d9164d281eef2c62f858fb5016))
* **message:** json version detection ([7b75471](https://github.com/mhughdo/solana-go/commit/7b7547155906993e24f0bee1af0468a679dc00b3))
* **message:** surface typed `ErrAddressTablesNotSet` from AccountMetaList (closes [#280](https://github.com/mhughdo/solana-go/issues/280)) ([#441](https://github.com/mhughdo/solana-go/issues/441)) ([a87922d](https://github.com/mhughdo/solana-go/commit/a87922db64914ec510f0bd3994fa97f9e8cb41cc))
* **message:** use gojson ([8d211d5](https://github.com/mhughdo/solana-go/commit/8d211d5dc9e610b54fb84f662d83e2f55668e9d4))
* missing signer message ([285a29f](https://github.com/mhughdo/solana-go/commit/285a29f0f7fc7a95486aa3a4d2409de577dc1577))
* msg/txn inconsistencies with original SDK ([88a5b52](https://github.com/mhughdo/solana-go/commit/88a5b523d2c971514b12592e832ecc6c7554be50))
* nil pointer err in decodeResponseFromMessage ([bad6cfb](https://github.com/mhughdo/solana-go/commit/bad6cfbb1de15653e3fbefce509b15a83f43f682))
* panics in non-Must* functions ([bdd6a7d](https://github.com/mhughdo/solana-go/commit/bdd6a7dfab6554c5dd28070120f79196f7859402))
* panics in non-Must* functions ([e17cd73](https://github.com/mhughdo/solana-go/commit/e17cd730796b5b08ecb870406aa270144534878d))
* reject malformed ed25519 private keys in PrivateKeyFromBase58 ([edcedcc](https://github.com/mhughdo/solana-go/commit/edcedcc2ba5ebd01c65baf64b8a22bf879cb0d55))
* remove toolchain directive and update CI matrix for Go 1.23+ ([eaf3272](https://github.com/mhughdo/solana-go/commit/eaf3272ac9c6c9c394201a3364a936d54c97083a))
* **rpc:** default simulateTransaction Accounts.Encoding to base64 (closes [#446](https://github.com/mhughdo/solana-go/issues/446)) ([#447](https://github.com/mhughdo/solana-go/issues/447)) ([2697614](https://github.com/mhughdo/solana-go/commit/26976146120a5030f35cd108d6a1257554c51cfa))
* **rpc:** match ParsedTransactionMeta to TransactionMeta ([a0f95c2](https://github.com/mhughdo/solana-go/commit/a0f95c23eac6031c0f44e3095b763da531b8b2b7))
* **rpc:** match ParsedTransactionMeta to TransactionMeta ([60ccd62](https://github.com/mhughdo/solana-go/commit/60ccd6211e18215ec1b6d9d7b071529ffc11067f)), closes [#284](https://github.com/mhughdo/solana-go/issues/284)
* **rpc:** support EncodingJSON in GetBlockWithOpts ([#419](https://github.com/mhughdo/solana-go/issues/419)) ([eee363a](https://github.com/mhughdo/solana-go/commit/eee363a738642efc6006cdce863689d49afc712c))
* simulateTransaction response decoding ([3e39c80](https://github.com/mhughdo/solana-go/commit/3e39c80b7211c33b5a2b5af2eb9e5d5f66e56c3b))
* sol/wsol addresses mismatch ([ae53ea9](https://github.com/mhughdo/solana-go/commit/ae53ea96965cc526ecbccf26649edb0d8a81ccf3))
* sol/wsol addresses mismatch ([b4796db](https://github.com/mhughdo/solana-go/commit/b4796dbd64454cd98b27d14639ad9fb4e063352b))
* **token,token-2022:** Build() sets Impl to *T, matching DecodeInstruction (closes [#222](https://github.com/mhughdo/solana-go/issues/222)) ([#440](https://github.com/mhughdo/solana-go/issues/440)) ([38c57db](https://github.com/mhughdo/solana-go/commit/38c57dbfbc5f1b18fffa9bbc029f867f8a9116c5))
* **token,token-2022:** per-instruction ProgramID override (closes [#254](https://github.com/mhughdo/solana-go/issues/254)) ([#439](https://github.com/mhughdo/solana-go/issues/439)) ([f489aac](https://github.com/mhughdo/solana-go/commit/f489aaced7390fc7bd0fb2ceed2773f2414992d7))
* use bytes.Compare for deterministic account sort + update tests ([99a66fb](https://github.com/mhughdo/solana-go/commit/99a66fbc1aa5fb8049bdf3f38dd5fbf5915f48c0))
* use bytes.Compare for deterministic account sort + update tests ([cb87ac7](https://github.com/mhughdo/solana-go/commit/cb87ac7fd78ae5ba22bb73fcc292b0dfa4fe3d06))
* ws race and leak fixes ([4370b70](https://github.com/mhughdo/solana-go/commit/4370b709b0d128e28c6ed45c829bd2972ec89941))
* ws race and leak fixes ([6be3404](https://github.com/mhughdo/solana-go/commit/6be3404fe9ef3240eae80aa50d41896eb08771ba))
* **ws:** reject EncodingJSONParsed in BlockSubscribe ([#426](https://github.com/mhughdo/solana-go/issues/426)) ([bf130a2](https://github.com/mhughdo/solana-go/commit/bf130a2a69b0a3f0462f8119c6b03dd1e9282cf8))
* **ws:** surface subscription request errors to Recv (closes [#175](https://github.com/mhughdo/solana-go/issues/175)) ([#449](https://github.com/mhughdo/solana-go/issues/449)) ([725147a](https://github.com/mhughdo/solana-go/commit/725147a6691fa580d392ec7ccf844da41d4ddafe))
* **ws:** use spec "showRewards" key in blockSubscribe params ([#430](https://github.com/mhughdo/solana-go/issues/430)) ([6969f12](https://github.com/mhughdo/solana-go/commit/6969f121e5700803befeb089e9dc4bbecfdb5f89))
* **ws:** use uint64 for params.Subscription in incoming notifications ([#427](https://github.com/mhughdo/solana-go/issues/427)) ([427de1a](https://github.com/mhughdo/solana-go/commit/427de1a9f438b658dd649ba6f13ba81558192ee1))


### Performance Improvements

* **json:** swap encoding/json and jsoniter for goccy/go-json ([c445f76](https://github.com/mhughdo/solana-go/commit/c445f76c249d944731983fd720c2a9e6a874dc62))
* **json:** swap encoding/json and jsoniter for goccy/go-json ([2afb213](https://github.com/mhughdo/solana-go/commit/2afb213f730f3a5a594b449a86eee7cd9ce10235))
* **message:** eliminate complex scans, struct copies, and redundant allocs ([aea7d1f](https://github.com/mhughdo/solana-go/commit/aea7d1f61da18b240d741fb003e83182c968e701))
* **message:** eliminate complex scans, struct copies, and redundant allocs ([adbb10e](https://github.com/mhughdo/solana-go/commit/adbb10e200c3979dafdee38fb2e13dbe356ab58d))
* migrate and use `fd_base58` algo ([801a013](https://github.com/mhughdo/solana-go/commit/801a013d59cf553e0c9e2a635fd1376e96459253))
* migrate to curve25519-voi for ed25519 operations ([20713fb](https://github.com/mhughdo/solana-go/commit/20713fbbe52d4d20cab792a702771790346f19c7))
* **transaction:** add cap hints and use pk instead of str ([91e8cec](https://github.com/mhughdo/solana-go/commit/91e8cec9785fccd2663f28e61c8cc5353f38c419))
* **transaction:** add cap hints and use pk instead of str ([05c352e](https://github.com/mhughdo/solana-go/commit/05c352e3eff441cb696539f84ce2158989c64e61))
* use fd_base58 algo ([689623d](https://github.com/mhughdo/solana-go/commit/689623db1477519056d79fbd6b32ec463f106673))

## [1.22.0](https://github.com/solana-foundation/solana-go/compare/v1.21.0...v1.22.0) (2026-06-30)


### Features

* add nonce account support ([#456](https://github.com/solana-foundation/solana-go/issues/456)) ([336881c](https://github.com/solana-foundation/solana-go/commit/336881cfbe66902ee0ec07b25dc1912186508a72))
* add sysvars ([#457](https://github.com/solana-foundation/solana-go/issues/457)) ([f777896](https://github.com/solana-foundation/solana-go/commit/f7778964b61f6ac2622a194fb034dd92da59c67f))
* **address-lookup-table:** one-call resolve for message lookups (closes [#262](https://github.com/solana-foundation/solana-go/issues/262)) ([#445](https://github.com/solana-foundation/solana-go/issues/445)) ([34beab1](https://github.com/solana-foundation/solana-go/commit/34beab1e231b6e85b7232361997d70bdd79825ef))
* **rpc:** add getTransactionsForAddress client method (closes [#343](https://github.com/solana-foundation/solana-go/issues/343)) ([#450](https://github.com/solana-foundation/solana-go/issues/450)) ([9e538c8](https://github.com/solana-foundation/solana-go/commit/9e538c84246ef1a3abfbcc96cac7a3196889b479))
* **rpc:** forward minContextSlot in 4 remaining JSON-RPC endpoints ([#448](https://github.com/solana-foundation/solana-go/issues/448)) ([c176402](https://github.com/solana-foundation/solana-go/commit/c176402c339e25c704f4eba9be540d56f770feb0))
* **rpc:** forward MinContextSlot in getBalance/getLatestBlockhash/getSlot/getTokenAccountBalance ([#442](https://github.com/solana-foundation/solana-go/issues/442)) ([b8e70e8](https://github.com/solana-foundation/solana-go/commit/b8e70e8c5cdd1229c6126f00090527664aa8496c))
* **zk:** add ElGamal & AES key derivation ([#413](https://github.com/solana-foundation/solana-go/issues/413)) ([9fcbf0c](https://github.com/solana-foundation/solana-go/commit/9fcbf0ced8af5b5ad975e45999478b5c36c3e65a))


### Bug Fixes

* **message:** surface typed `ErrAddressTablesNotSet` from AccountMetaList (closes [#280](https://github.com/solana-foundation/solana-go/issues/280)) ([#441](https://github.com/solana-foundation/solana-go/issues/441)) ([a87922d](https://github.com/solana-foundation/solana-go/commit/a87922db64914ec510f0bd3994fa97f9e8cb41cc))
* **rpc:** default simulateTransaction Accounts.Encoding to base64 (closes [#446](https://github.com/solana-foundation/solana-go/issues/446)) ([#447](https://github.com/solana-foundation/solana-go/issues/447)) ([2697614](https://github.com/solana-foundation/solana-go/commit/26976146120a5030f35cd108d6a1257554c51cfa))
* **token,token-2022:** Build() sets Impl to *T, matching DecodeInstruction (closes [#222](https://github.com/solana-foundation/solana-go/issues/222)) ([#440](https://github.com/solana-foundation/solana-go/issues/440)) ([38c57db](https://github.com/solana-foundation/solana-go/commit/38c57dbfbc5f1b18fffa9bbc029f867f8a9116c5))
* **token,token-2022:** per-instruction ProgramID override (closes [#254](https://github.com/solana-foundation/solana-go/issues/254)) ([#439](https://github.com/solana-foundation/solana-go/issues/439)) ([f489aac](https://github.com/solana-foundation/solana-go/commit/f489aaced7390fc7bd0fb2ceed2773f2414992d7))
* **ws:** surface subscription request errors to Recv (closes [#175](https://github.com/solana-foundation/solana-go/issues/175)) ([#449](https://github.com/solana-foundation/solana-go/issues/449)) ([725147a](https://github.com/solana-foundation/solana-go/commit/725147a6691fa580d392ec7ccf844da41d4ddafe))

## [1.21.0](https://github.com/solana-foundation/solana-go/compare/v1.20.0...v1.21.0) (2026-05-25)


### Features

* **rpc:** add NewWithCommitment / NewWithTimeout / NewWithTimeoutAndCommitment ([#436](https://github.com/solana-foundation/solana-go/issues/436)) ([e93ff5e](https://github.com/solana-foundation/solana-go/commit/e93ff5e937733daca5fed01c362961c4c8aead25)), closes [#414](https://github.com/solana-foundation/solana-go/issues/414)
* **rpc:** forward MinContextSlot in getProgramAccounts and getTokenAccounts ([#431](https://github.com/solana-foundation/solana-go/issues/431)) ([17984a5](https://github.com/solana-foundation/solana-go/commit/17984a55c17ab0fc9f308872a43b737601d6a8da))
* **rpc:** support EncodingJSON in GetTransaction ([#420](https://github.com/solana-foundation/solana-go/issues/420)) ([b906b70](https://github.com/solana-foundation/solana-go/commit/b906b70527a5dfed358090e27dd7f4a7f12749c3))
* **wallet:** derive PrivateKey/Wallet from BIP-39 mnemonic ([#429](https://github.com/solana-foundation/solana-go/issues/429)) ([89ef706](https://github.com/solana-foundation/solana-go/commit/89ef706472ad49a9622a058497852711f7bd3771))
* **ws:** support dataSlice in AccountSubscribe ([#433](https://github.com/solana-foundation/solana-go/issues/433)) ([fb31fb1](https://github.com/solana-foundation/solana-go/commit/fb31fb13b42141bb6067c7447b8618e7e848b97b))
* **ws:** support dataSlice in ProgramSubscribe ([#434](https://github.com/solana-foundation/solana-go/issues/434)) ([950b110](https://github.com/solana-foundation/solana-go/commit/950b110b8f369de33143705cfba0b8da7d240d6f))
* **ws:** support enableReceivedNotification in SignatureSubscribe ([#432](https://github.com/solana-foundation/solana-go/issues/432)) ([810f171](https://github.com/solana-foundation/solana-go/commit/810f171ff933c1508e9526a2a536a287cac7c386))


### Bug Fixes

* **rpc:** support EncodingJSON in GetBlockWithOpts ([#419](https://github.com/solana-foundation/solana-go/issues/419)) ([eee363a](https://github.com/solana-foundation/solana-go/commit/eee363a738642efc6006cdce863689d49afc712c))
* **ws:** reject EncodingJSONParsed in BlockSubscribe ([#426](https://github.com/solana-foundation/solana-go/issues/426)) ([bf130a2](https://github.com/solana-foundation/solana-go/commit/bf130a2a69b0a3f0462f8119c6b03dd1e9282cf8))
* **ws:** use spec "showRewards" key in blockSubscribe params ([#430](https://github.com/solana-foundation/solana-go/issues/430)) ([6969f12](https://github.com/solana-foundation/solana-go/commit/6969f121e5700803befeb089e9dc4bbecfdb5f89))
* **ws:** use uint64 for params.Subscription in incoming notifications ([#427](https://github.com/solana-foundation/solana-go/issues/427)) ([427de1a](https://github.com/solana-foundation/solana-go/commit/427de1a9f438b658dd649ba6f13ba81558192ee1))

## [1.20.0](https://github.com/solana-foundation/solana-go/compare/v1.19.0...v1.20.0) (2026-05-08)


### Features

* **jsonrpc:** add CustomHeader http.Header for multi-value headers ([20b37ba](https://github.com/solana-foundation/solana-go/commit/20b37ba403c438ebe914b43ff7081f9598832d0c))


### Performance Improvements

* migrate to curve25519-voi for ed25519 operations ([20713fb](https://github.com/solana-foundation/solana-go/commit/20713fbbe52d4d20cab792a702771790346f19c7))

## [1.19.0](https://github.com/solana-foundation/solana-go/compare/v1.18.0...v1.19.0) (2026-04-23)


### Features

* is token mint classifier ([4f72982](https://github.com/solana-foundation/solana-go/commit/4f72982442c9b3c166b72dbb2de730f58b575539))


### Bug Fixes

* enhance getUint64 function to handle string inputs ([5309095](https://github.com/solana-foundation/solana-go/commit/53090952ffc598c1870617b1727179135994ec65))
* keep websocket request IDs within JSON-safe range ([8ed3105](https://github.com/solana-foundation/solana-go/commit/8ed31050f7af62f26b5615f40546bb498cab9219))
* **message:** json version detection ([1fd2201](https://github.com/solana-foundation/solana-go/commit/1fd2201431de71d9164d281eef2c62f858fb5016))
* **message:** use gojson ([8d211d5](https://github.com/solana-foundation/solana-go/commit/8d211d5dc9e610b54fb84f662d83e2f55668e9d4))
* reject malformed ed25519 private keys in PrivateKeyFromBase58 ([edcedcc](https://github.com/solana-foundation/solana-go/commit/edcedcc2ba5ebd01c65baf64b8a22bf879cb0d55))
* **rpc:** match ParsedTransactionMeta to TransactionMeta ([a0f95c2](https://github.com/solana-foundation/solana-go/commit/a0f95c23eac6031c0f44e3095b763da531b8b2b7))

### Performance Improvements

* **json:** swap encoding/json and jsoniter for goccy/go-json ([c445f76](https://github.com/solana-foundation/solana-go/commit/c445f76c249d944731983fd720c2a9e6a874dc62))
* **transaction:** add cap hints and use pk instead of str ([91e8cec](https://github.com/solana-foundation/solana-go/commit/91e8cec9785fccd2663f28e61c8cc5353f38c419))


## [1.18.0](https://github.com/solana-foundation/solana-go/compare/v1.17.0...v1.18.0) (2026-04-16)


### Features

* add getters to txn with meta
* add token-2022 extensions
* stake state types & ext tests
* vote program complete

### Bug Fixes

* allign rpc client with agave
* memo program parity

### Performance Improvements

* **message:** eliminate complex scans, struct copies, and redundant allocs
