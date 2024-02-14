CREATE TABLE public.campaigns
(
    id               bigserial NOT NULL,
    user_id          int8 NULL,
    name             varchar NULL,
    sort_description varchar NULL,
    description      varchar NULL,
    perks            varchar NULL,
    backer_count     int8 NULL,
    goal_amount      int8 NULL,
    current_amount   int8 NULL,
    slug             varchar NULL,
    created_at       timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT campaign_pk PRIMARY KEY (id)
);