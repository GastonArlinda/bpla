CREATE TABLE session
(
  session_id UUID NOT NULL,
  timestamp TIMESTAMPTZ NOT NULL,
  latitude NUMERIC NOT NULL,
  longitude NUMERIC NOT NULL,
  altitude NUMERIC NOT NULL,
  speed NUMERIC NOT NULL,
  roll NUMERIC NOT NULL,
  battery NUMERIC NOT NULL
);