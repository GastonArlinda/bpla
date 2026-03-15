CREATE OR REPLACE FUNCTION get_flight_statistics()
RETURNS TABLE (
    total_flights INTEGER,
    avg_distance_meters NUMERIC,
    max_flight_distance_meters NUMERIC,
    max_flight_duration_seconds INTEGER,
    flights_last_30sec INTEGER,
    max_speed_mps NUMERIC,
    avg_battery_drain_percent NUMERIC,
    total_distance_meters NUMERIC
) 
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
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
            COALESCE(SUM(segment_distance), 0) as total_distance,
            EXTRACT(EPOCH FROM (MAX(session_end) - MIN(session_start))) as duration,
            COALESCE(MAX(speed), 0) as max_speed,
            MIN(session_start) as start_time,
            COALESCE((MAX(first_battery) - MIN(last_battery)), 0) as battery_drain
        FROM segment_distances
        WHERE session_id IS NOT NULL
        GROUP BY session_id
    )
    SELECT 
        COUNT(DISTINCT session_id)::INTEGER as total_flights,
        ROUND(COALESCE(AVG(total_distance), 0)::NUMERIC, 2) as avg_distance_meters,
        ROUND(COALESCE(MAX(total_distance), 0)::NUMERIC, 2) as max_flight_distance_meters,
        COALESCE(MAX(duration), 0)::INTEGER as max_flight_duration_seconds,
        COUNT(DISTINCT CASE 
            WHEN start_time > NOW() - INTERVAL '30 seconds' 
            THEN session_id 
        END)::INTEGER as flights_last_30sec,
        ROUND(COALESCE(MAX(max_speed), 0)::NUMERIC, 2) as max_speed_mps,
        ROUND(COALESCE(AVG(battery_drain), 0)::NUMERIC, 2) as avg_battery_drain_percent,
        ROUND(COALESCE(SUM(total_distance), 0)::NUMERIC, 2) as total_distance_meters
    FROM session_metrics;
END;
$$;