CREATE TABLE IF NOT EXISTS policy (
    id UUID NOT NULL PRIMARY KEY,
    name varchar(64) UNIQUE NOT NULL,
    state varchar(16) NOT NULL DEFAULT 'active',
    provider UUID NOT NULL,
    resource_type varchar(32),
    resources_evaluated varchar(64),
    tags UUID NOT NULL,
    steampipe varchar(256) NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tags (
    UNIQUE (type, key),

    id UUID NOT NULL PRIMARY KEY,
    type varchar(32) NOT NULL,
    key varchar(64) NOT NULL,
    value varchar(255) NOT NULL,
    state varchar(16) NOT NULL DEFAULT 'active',
    provider UUID NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);