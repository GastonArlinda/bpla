CREATE OR REPLACE PROCEDURE get_flight_statistics(
    OUT total_flights INTEGER,
    OUT avg_distance_meters NUMERIC,
    OUT max_distance_meters NUMERIC,
    OUT max_flight_duration_seconds INTEGER,
    OUT flights_last_30sec INTEGER,
    OUT max_speed_mps NUMERIC,
    OUT avg_battery_drain_percent NUMERIC,
    OUT total_distance_meters NUMERIC
)
LANGUAGE plpgsql
AS $$
BEGIN
    WITH segment_distances AS (
        SELECT 
            session_id,
            timestamp,
            calculate_distance(
                LAG(latitude) OVER (PARTITION BY session_id ORDER BY timestamp),
                LAG(longitude) OVER (PARTITION BY session_id ORDER BY timestamp),
                latitude,
                longitude
            ) as segment_distance,
            speed,
            battery,
            FIRST_VALUE(battery) OVER (PARTITION BY session_id ORDER BY timestamp) as first_battery,
            LAST_VALUE(battery) OVER (PARTITION BY session_id ORDER BY timestamp) as last_battery,
            MIN(timestamp) OVER (PARTITION BY session_id) as session_start,
            MAX(timestamp) OVER (PARTITION BY session_id) as session_end
        FROM session
    ),
    session_metrics AS (
        SELECT 
            session_id,
            SUM(segment_distance) as total_distance,
            MAX(segment_distance) as max_segment_distance,
            EXTRACT(EPOCH FROM (MAX(session_end) - MIN(session_start))) as duration,
            MAX(speed) as max_speed,
            MIN(session_start) as start_time,
            (MAX(first_battery) - MIN(last_battery)) as battery_drain
        FROM segment_distances
        GROUP BY session_id
    )
    SELECT 
        COUNT(DISTINCT session_id),
        COALESCE(AVG(total_distance), 0),
        COALESCE(MAX(max_segment_distance), 0),
        COALESCE(MAX(duration)::INTEGER, 0),
        COUNT(DISTINCT CASE WHEN start_time > NOW() - INTERVAL '30 seconds' THEN session_id END),
        COALESCE(MAX(max_speed), 0),
        COALESCE(AVG(battery_drain), 0),
        COALESCE(SUM(total_distance), 0)
    INTO 
        total_flights,
        avg_distance_meters,
        max_distance_meters,
        max_flight_duration_seconds,
        flights_last_30sec,
        max_speed_mps,
        avg_battery_drain_percent,
        total_distance_meters
    FROM session_metrics;
END;
$$;