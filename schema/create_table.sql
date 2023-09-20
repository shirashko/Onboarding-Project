CREATE TABLE IF NOT EXISTS userProfiles (
                                            id SERIAL PRIMARY KEY,
                                            username VARCHAR(255) UNIQUE,
                                            full_name VARCHAR(255),
                                            bio TEXT,
                                            profile_pic_url VARCHAR(255)
);
