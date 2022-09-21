CREATE TABLE public.registered_paths (
    user_id uuid,
    paths jsonb DEFAULT '[]'::jsonb NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    id uuid DEFAULT extensions.uuid_generate_v4() NOT NULL
);


CREATE TABLE public.tips (
    user_id uuid,
    data jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    id character varying DEFAULT extensions.uuid_generate_v4() NOT NULL
);
