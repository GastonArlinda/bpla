CREATE OR REPLACE FUNCTION calculate_distance(
    lat1 NUMERIC, lon1 NUMERIC,
    lat2 NUMERIC, lon2 NUMERIC
) RETURNS NUMERIC AS $$

DECLARE
    R NUMERIC := 6371000;
    phi1 NUMERIC;
    phi2 NUMERIC;
    dphi NUMERIC;
    dlambda NUMERIC;
    a NUMERIC;
    c NUMERIC;

BEGIN
    phi1 := radians(lat1);
    phi2 := radians(lat2);
    dphi := radians(lat2 - lat1);
    dlambda := radians(lon2 - lon1);
    
    a := sin(dphi/2) * sin(dphi/2) +
         cos(phi1) * cos(phi2) *
         sin(dlambda/2) * sin(dlambda/2);
    c := 2 * atan2(sqrt(a), sqrt(1-a));
    
    RETURN R * c;
END;

$$ LANGUAGE plpgsql;