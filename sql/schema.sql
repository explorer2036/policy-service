CREATE TABLE IF NOT EXISTS policy (
    id UUID NOT NULL PRIMARY KEY,
    name varchar(64) UNIQUE NOT NULL,
    state varchar(16) NOT NULL DEFAULT 'active',
    provider UUID NOT NULL, -- provider id
    resource_type varchar(255),
    resources_evaluated varchar(255),
    tags UUID[] NOT NULL,
    steampipe varchar(256) NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tag (
    UNIQUE (type, key),

    id UUID NOT NULL PRIMARY KEY,
    type varchar(32) NOT NULL,
    key varchar(64) NOT NULL,
    value varchar(255) NOT NULL,
    state varchar(16) NOT NULL DEFAULT 'active',
    provider UUID NOT NULL, -- provider id
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS provider_type (
    id UUID NOT NULL PRIMARY KEY,
    name varchar(32) UNIQUE NOT NULL,
    state varchar(16) NOT NULL DEFAULT 'active',
    description varchar(128),
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS provider (
    id UUID NOT NULL PRIMARY KEY,
    name varchar(32) UNIQUE NOT NULL,
    url varchar(64) NOT NULL,
    provider_type UUID NOT NULL, -- provider type id
    state varchar(16) NOT NULL DEFAULT 'active',
    description varchar(128),
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS benchmark (
    id UUID NOT NULL PRIMARY KEY,
    name varchar(64) UNIQUE NOT NULL,
    state varchar(16) NOT NULL DEFAULT 'active',
    provider UUID NOT NULL, -- provider id
    resource_type varchar(255),
    resources_evaluated varchar(255),
    tags UUID [] NOT NULL, -- tags id
    policies UUID [] NOT NULL, -- policies id
    description varchar(128),
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now()
);