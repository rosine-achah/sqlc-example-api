ALTER TABLE message
ADD COLUMN thread_id INT NOT NULL,
ADD CONSTRAINT fk_thread
    FOREIGN KEY(thread_id)
    REFERENCES thread(id)
    ON DELETE CASCADE;
