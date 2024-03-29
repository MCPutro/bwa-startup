CREATE TABLE IF NOT EXISTS public.transactions(
    id bigserial        NOT NULL,
    campaign_id int8    NOT NULL,
    user_id int8        NOT NULL,
    amount int8         DEFAULT 0 NULL,
    status              varchar NULL,
    code                varchar NULL,
    created_at          timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at          timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT transactions_pk PRIMARY KEY (id),
    CONSTRAINT transactions_campaigns_id_fk FOREIGN KEY (campaign_id) REFERENCES public.campaigns(id),
    CONSTRAINT transactions_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id)

create index transactions_campaign_id_idx
    on transactions (campaign_id);

create index transactions_campaign_n_user_id_idx
    on transactions (campaign_id, user_id);

create index transactions_user_id_idx
    on transactions (user_id);

