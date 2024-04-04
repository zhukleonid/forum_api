CREATE TABLE IF NOT EXISTS posts_likes (
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    status BOOLEAN  NOT NULL, 
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);