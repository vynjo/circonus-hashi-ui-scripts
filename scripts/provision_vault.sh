#!/bin/bash
###### Vault Metrics
## Numeric


## CLient
create_cgc -query "vault*expire*num_leases" -title "Vault expire num_leases" -tags "creator:api,role:client,service:vault,data-type:counter,group:primary"
create_cgc -query "vault*runtime*alloc_bytes" -title "Vault runtime alloc_bytes" -tags "creator:api,role:client,service:vault,data-type:gauge,group:primary"
create_cgc -query "vault*runtime*free_count" -title "Vault runtime free_count" -tags "creator:api,role:client,service:vault,data-type:counter,group:primary"
create_cgc -query "vault*runtime*heap_objects" -title "Vault runtime heap_objects" -tags "creator:api,role:client,service:vault,data-type:counter,group:primary"
create_cgc -query "vault*runtime*malloc_count" -title "Vault runtime malloc_count" -tags "creator:api,role:client,service:vault,data-type:counter,group:primary"
create_cgc -query "vault*runtime*num_goroutines" -title "Vault runtime num_goroutines" -tags "creator:api,role:client,service:vault,data-type:gauge,group:primary"
create_cgc -query "vault*runtime*sys_bytes" -title "Vault runtime sys_bytes" -tags "creator:api,role:client,service:vault,data-type:gauge,group:primary"
create_cgc -query "vault*runtime*total_gc_pause_ns" -title "Vault runtime total_gc_pause_ns" -tags "creator:api,role:client,service:vault,data-type:counter,group:primary"
create_cgc -query "vault*runtime*total_gc_runs" -title "Vault runtime total_gc_runs" -tags "creator:api,role:client,service:vault,data-type:counter,group:primary"

## Histogram

## Server

create_cgc -query "vault*core*post_unseal" -title "Vault core post_unseal" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*core*pre_seal" -title "Vault core pre_seal" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*core*seal-internal" -title "Vault core seal-internal" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*core*seal-with-request" -title "Vault core seal-with-request" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*core*unseal" -title "Vault core unseal" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*expire*fetch-lease-times" -title "Vault expire fetch-lease-times" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*expire*fetch-lease-times-by-token" -title "Vault expire fetch-lease-times-by-token" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*expire*register" -title "Vault expire register" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*expire*register-auth" -title "Vault expire register-auth" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*expire*renew" -title "Vault expire renew" -tags "creator:api,role:server,service:vault,data-type:histogram,group:secondary"

## Client
create_cgc -query "vault*policy*get_policy" -title "Vault policy get_policy" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*policy*set_policy" -title "Vault policy set_policy" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*rollback*attempt*auth-token-" -title "Vault rollback attempt auth-token" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*rollback*attempt*cubbyhole-" -title "Vault rollback attempt cubbyhole" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*rollback*attempt*mysql-" -title "Vault rollback attempt mysql" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*rollback*attempt*secret-" -title "Vault rollback attempt secret" -tags "creator:api,role:client,service:vault,data-type:histogram,group:primary"
create_cgc -query "vault*rollback*attempt*sys-" -title "Vault rollback attempt sys" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*create*secret-" -title "Vault route create secret" -tags "creator:api,role:client,service:vault,data-type:histogram,group:primary"
create_cgc -query "vault*route*read*auth-token-" -title "Vault route read auth-token" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*read*mysql-" -title "Vault route read mysql" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*read*secret-" -title "Vault route read secret" -tags "creator:api,role:client,service:vault,data-type:histogram,group:primary"
create_cgc -query "vault*route*read*sys-" -title "Vault route read sys" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*renew*mysql-" -title "Vault route renew mysql" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*rollback*auth-token-" -title "Vault route rollback auth-token" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*rollback*cubbyhole-" -title "Vault route rollback cubbyhole " -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*rollback*mysql-" -title "Vault route rollback mysql" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*rollback*secret-" -title "Vault route rollback secret" -tags "creator:api,role:client,service:vault,data-type:histogram,group:primary"
create_cgc -query "vault*route*rollback*sys-" -title "Vault route rollback sys" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*update*auth-token-" -title "Vault route update auth-token" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*update*mysql-" -title "Vault route update mysql" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*route*update*secret-" -title "Vault route update secret" -tags "creator:api,role:client,service:vault,data-type:histogram,group:primary"
create_cgc -query "vault*route*update*sys-" -title "Vault route update sys" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*runtime*gc_pause_ns" -title "Vault runtime gc_pause_ns" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*token*create" -title "Vault token create" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*token*createAccessor" -title "Vault token createAccessor" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
create_cgc -query "vault*token*lookup" -title "Vault token lookup" -tags "creator:api,role:client,service:vault,data-type:histogram,group:secondary"
