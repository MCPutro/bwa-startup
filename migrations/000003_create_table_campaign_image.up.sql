CREATE TABLE IF NOT EXISTS public.campaign_images
(
    id          bigserial NOT NULL,
    campaign_id int8      NOT NULL,
    filename    varchar NULL,
    is_primary  bool      NOT NULL DEFAULT false,
    created_at  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT campaign_image_pk PRIMARY KEY (id)
);