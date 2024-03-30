CREATE TABLE IF NOT EXISTS public.campaign_images
(
    id          bigserial NOT NULL,
    campaign_id int8      NOT NULL,
    image       varchar NULL,
    is_primary  bool      NOT NULL DEFAULT false,
    created_at  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT campaign_image_pk PRIMARY KEY (id),
    CONSTRAINT campaign_images_campaigns_id_fk FOREIGN KEY (campaign_id) REFERENCES public.campaigns(id)
);