CREATE TABLE models (
    model JSONB
);

CREATE INDEX idx_model ON models USING gin ((model->'order_uid'))