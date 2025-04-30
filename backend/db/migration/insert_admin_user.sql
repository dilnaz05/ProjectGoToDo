-- Admin қолданушыны қосу (паролі bcrypt-пен хештелген деп есептейміз)
INSERT INTO users (username, password, role, created_at, updated_at)
VALUES (
           'admin',
           '$2a$10$Xh8Y9yOZk5uL5uEtx4uL/O5kzbEOGzDR3yRpM0KUmzR9HgqfQ8uEi', -- пароль: admin123
           'admin',
           NOW(),
           NOW()
       )
    ON CONFLICT (username) DO NOTHING;