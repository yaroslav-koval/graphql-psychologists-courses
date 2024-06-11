CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW() AT TIME ZONE 'UTC';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS psychologist
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at  TIMESTAMP        DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at  TIMESTAMP        DEFAULT (NOW() AT TIME ZONE 'UTC'),
    name        VARCHAR(50) NOT NULL,
    description VARCHAR(500)
);

CREATE OR REPLACE TRIGGER updated_at_trigger
    BEFORE UPDATE
    ON psychologist
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS course
(
    id                  UUID      NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at          TIMESTAMP NOT NULL             DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at          TIMESTAMP NOT NULL             DEFAULT (NOW() AT TIME ZONE 'UTC'),
    name                VARCHAR(200) NOT NULL,
    description         TEXT,
    price               INT NOT NULL
);

CREATE OR REPLACE TRIGGER updated_at_trigger
    BEFORE UPDATE
    ON course
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS courses_psychologists
(
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at   TIMESTAMP        DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at   TIMESTAMP        DEFAULT (NOW() AT TIME ZONE 'UTC'),
    course       UUID NOT NULL
        constraint fk_c
            references course,
    psychologist UUID NOT NULL
        constraint fk_p
            references psychologist,
    unique (course, psychologist)
);

CREATE OR REPLACE TRIGGER updated_at_trigger
    BEFORE UPDATE
    ON courses_psychologists
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS lesson
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP        DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP        DEFAULT (NOW() AT TIME ZONE 'UTC'),
    name       VARCHAR(500) NOT NULL,
    number     INT  NOT NULL,
    course     UUID NOT NULL
        constraint fk_c
            references course
);

CREATE OR REPLACE TRIGGER updated_at_trigger
    BEFORE UPDATE
    ON lesson
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
