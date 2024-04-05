alter table device_profile
    add column save_gw_rx_on_join_req boolean default false,
    add column sync_sec_ctx_on_join_req boolean default false;
