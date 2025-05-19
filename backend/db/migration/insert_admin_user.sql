-- Admin қолданушыны қосу (паролі bcrypt-пен хештелген деп есептейміз)
INSERT INTO users (username, password, role, created_at, updated_at)
VALUES (
           'admin',
           '$2a$10$zKXFx1e5t9hdJgOo3t6Tp.zF.yjQWvMfCniDiIwS0oM5RzHf6Ey8Gy', -- пароль: admin456
           'admin',
           NOW(),
           NOW()
       )
    ON CONFLICT (username) DO NOTHING;