-- teswir --

CREATE TABLE tbL_user (
    id        uuid PRIMARY KEY        DEFAULT gen_random_uuid(),
    username  VARCHAR(64) NOT NULL,
    firstname VARCHAR(64) NOT NULL,
    lastname  VARCHAR(64) NOT NULL,
    user_role VARCHAR(20) NOT NULL,
    create_ts timestamp without time zone default current_timestamp,
    update_ts timestamp without time zone default current_timestamp
);

